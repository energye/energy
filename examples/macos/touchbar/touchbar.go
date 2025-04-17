//go:build darwin
// +build darwin

package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/examples/macos/touchbar/bar"
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar"
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barbuilder"
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barutils"
	"os"
)

//go:embed assets
var assets embed.FS

func main() {
	cef.GlobalInit(nil, assets)

	//create application
	app := cef.NewApplication()
	if common.IsDarwin() {
		app.SetUseMockKeyChain(true)
	}
	cef.BrowserWindow.Config.LocalResource(cef.LocalLoadConfig{
		ResRootDir: "assets",
		FS:         assets,
		Home:       "touchbar.html",
	}.Build())
	cef.BrowserWindow.Config.Width = 400
	cef.BrowserWindow.Config.Height = 600

	var tb barbuilder.TouchBar
	var freeTb = func() {
		if tb != nil {
			fmt.Println("tb Uninstall")
			// end
			err := tb.Uninstall()
			if err != nil {
				panic(err)
			}
		}
	}
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		tb = touchbar.New(barbuilder.Options{
			EventErrorLogger: func(err error) {
				fmt.Println("EventErrorLogger", err)
			},
		})

		config := barutils.MakeStackableBar(tb, func(switcher barutils.Switcher) []barbuilder.Item {
			var (
				isMax  = false
				isMin  = false
				isFull = false
			)
			var (
				maxBtn        *barbuilder.Button
				minBtn        *barbuilder.Button
				fullScreenBtn *barbuilder.Button
			)
			maxBtn = &barbuilder.Button{
				Title: "最大化",
				OnClick: func() {
					if isMax {
						maxBtn.Title = "最大化"
					} else {
						maxBtn.Title = "还原"
					}
					window.Maximize()
					isMax = !isMax
					ipc.Emit("touchbar", 1, "touch bar: max button")
					switcher.Update()
				},
			}
			minBtn = &barbuilder.Button{
				Title: "最小化",
				OnClick: func() {
					if isMin {
						window.Restore()
						minBtn.Title = "最小化"
						fullScreenBtn.Disabled = false
						maxBtn.Disabled = false
					} else {
						window.Minimize()
						minBtn.Title = "还原"
						fullScreenBtn.Disabled = true
						maxBtn.Disabled = true
					}
					isMin = !isMin
					ipc.Emit("touchbar", 2, "touch bar: min button")
					switcher.Update()
				},
			}
			fullScreenBtn = &barbuilder.Button{
				Title: "全屏",
				OnClick: func() {
					if isFull {
						window.ExitFullScreen()
						fullScreenBtn.Title = "全屏"
						minBtn.Disabled = false
						maxBtn.Disabled = false
					} else {
						window.FullScreen()
						fullScreenBtn.Title = "退出全屏"
						minBtn.Disabled = true
						maxBtn.Disabled = true
					}
					isFull = !isFull
					ipc.Emit("touchbar", 3, "touch bar: full screen button")
					switcher.Update()
				},
			}
			return []barbuilder.Item{
				&barbuilder.Label{
					Content: &barbuilder.ContentLabel{
						Text: "Go Touch Bar",
					},
				},
				&barbuilder.SpaceLarge{},
				bar.MakeDemo(switcher),
				&barbuilder.SpaceSmall{},
				bar.MakeCatalog(switcher),
				&barbuilder.SpaceFlexible{},
				maxBtn,
				minBtn,
				fullScreenBtn,
				&barbuilder.Button{
					Title: "关闭",
					OnClick: func() {
						window.CloseBrowserWindow()
						freeTb()   // free
						os.Exit(0) // 在这里关闭时失败, 所以这样退出
					},
				},
			}
		})
		err := tb.Install(config)
		if err != nil {
			panic(err)
		}

	})
	//run application
	cef.Run(app)
	// free
	freeTb()
}
