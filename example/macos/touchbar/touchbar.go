package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/example/macos/touchbar/bar"
	"github.com/energye/energy/v2/pkgs/touchbar"
	"github.com/energye/energy/v2/pkgs/touchbar/barbuilder"
	"github.com/energye/energy/v2/pkgs/touchbar/barutils"
	"os"
)

func main() {
	cef.GlobalInit(nil, nil)

	//create application
	app := cef.NewApplication()
	app.SetUseMockKeyChain(true)
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
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
		//bw := window.AsLCLBrowserWindow().BrowserWindow()

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
					switcher.Update()
				},
			}
			minBtn = &barbuilder.Button{
				Title: "最小化",
				OnClick: func() {
					if isMin {
						window.Restore()
						minBtn.Title = "最小化"
					} else {
						window.Minimize()
						minBtn.Title = "还原"
					}
					isMin = !isMin
					switcher.Update()
				},
			}
			fullScreenBtn = &barbuilder.Button{
				Title: "全屏",
				OnClick: func() {
					if isFull {
						window.ExitFullScreen()
						fullScreenBtn.Title = "全屏"
					} else {
						window.FullScreen()
						fullScreenBtn.Title = "退出全屏"
					}
					isFull = !isFull
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
						freeTb()
						os.Exit(0) // 在这里关闭时失败, 所以这样退出
					},
				},
			}
		})

		err := tb.Install(config)
		fmt.Println("install err", err)
		if err != nil {
			panic(err)
		}

	})
	//run application
	cef.Run(app)
	freeTb()
}
