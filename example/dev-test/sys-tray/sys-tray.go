package sys_tray

import (
	"energye/systray"
	"energye/systray/icon"
	"fmt"
	"time"
)

var start func()
var end func()

func TrayMain() {
	onExit := func() {
		now := time.Now()
		fmt.Println("Exit at", now.String())
	}

	//go systray.Run(onReady, onExit)
	//systray.Register(onReady, onExit)//windows
	start, end = systray.RunWithExternalLoop(onReady, onExit) //windows/linux/macos
	start()
}

func addQuitItem() {
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	mQuit.Enable()
	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		//systray.Quit()
		//systray.Quit()// macos error
		//end() // macos error
		fmt.Println("Finished quitting")
	}()
}

func onReady() {
	fmt.Println("systray.onReady")
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Lantern")
	addQuitItem()

	// We can manipulate the systray in other goroutines
	go func() {
		systray.SetTemplateIcon(icon.Data, icon.Data)
		systray.SetTitle("Awesome App")
		systray.SetTooltip("Pretty awesome棒棒嗒")
		mChange := systray.AddMenuItem("Change Me", "Change Me")
		mChecked := systray.AddMenuItemCheckbox("Checked", "Check Me", true)
		mEnabled := systray.AddMenuItem("Enabled", "Enabled")
		// Sets the icon of a menu item. Only available on Mac.
		mEnabled.SetTemplateIcon(icon.Data, icon.Data)

		systray.AddMenuItem("Ignored", "Ignored")

		subMenuTop := systray.AddMenuItem("SubMenuTop", "SubMenu Test (top)")
		subMenuMiddle := subMenuTop.AddSubMenuItem("SubMenuMiddle", "SubMenu Test (middle)")
		subMenuBottom := subMenuMiddle.AddSubMenuItemCheckbox("SubMenuBottom - Toggle Panic!", "SubMenu Test (bottom) - Hide/Show Panic!", false)
		subMenuBottom2 := subMenuMiddle.AddSubMenuItem("SubMenuBottom - Panic!", "SubMenu Test (bottom)")

		systray.AddSeparator()
		mToggle := systray.AddMenuItem("Toggle", "Toggle some menu items")
		shown := true
		toggle := func() {
			if shown {
				subMenuBottom.Check()
				subMenuBottom2.Hide()
				mEnabled.Hide()
				shown = false
			} else {
				subMenuBottom.Uncheck()
				subMenuBottom2.Show()
				mEnabled.Show()
				shown = true
			}
		}
		mReset := systray.AddMenuItem("Reset", "Reset all items")

		for {
			select {
			case <-mChange.ClickedCh:
				mChange.SetTitle("I've Changed")
			case <-mChecked.ClickedCh:
				if mChecked.Checked() {
					mChecked.Uncheck()
					mChecked.SetTitle("Unchecked")
				} else {
					mChecked.Check()
					mChecked.SetTitle("Checked")
				}
			case <-mEnabled.ClickedCh:
				mEnabled.SetTitle("Disabled")
				mEnabled.Disable()
			case <-subMenuBottom2.ClickedCh:
				panic("panic button pressed")
			case <-subMenuBottom.ClickedCh:
				toggle()
			case <-mReset.ClickedCh:
				systray.ResetMenu()
				addQuitItem()
			case <-mToggle.ClickedCh:
				toggle()
			}
		}
	}()
}
