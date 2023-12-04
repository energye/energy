package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/example/macos/touchbar/bar"
	"github.com/energye/energy/v2/pkgs/touchbar"
	"github.com/energye/energy/v2/pkgs/touchbar/barbuilder"
	"github.com/energye/energy/v2/pkgs/touchbar/barutils"
)

func main() {
	cef.GlobalInit(nil, nil)

	//create application
	app := cef.NewApplication()
	app.SetUseMockKeyChain(true)
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	var tb barbuilder.TouchBar
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		tb = touchbar.New(barbuilder.Options{
			EventErrorLogger: func(err error) {
				fmt.Println("EventErrorLogger", err)
			},
		})

		makeUpdater := func(switcher barutils.Switcher) func() {
			return func() {
				fmt.Println("makeUpdater")
				err := switcher.Update()
				if err != nil {
					fmt.Printf("could not update: %v\n", err)
				}
			}
		}

		config := barutils.MakeStackableBar(tb, func(switcher barutils.Switcher) []barbuilder.Item {
			update := makeUpdater(switcher)
			fmt.Println("MakeStackableBar")
			return []barbuilder.Item{
				&barbuilder.Label{
					Content: &barbuilder.ContentLabel{
						Text: "Go Touch Bar",
					},
				},
				&barbuilder.SpaceLarge{},
				bar.MakeDemo(update),
				&barbuilder.SpaceSmall{},
				bar.MakeCatalog(switcher, update),
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
	if tb != nil {
		fmt.Println("tb Uninstall")
		// end
		err := tb.Uninstall()
		if err != nil {
			panic(err)
		}
	}
}
