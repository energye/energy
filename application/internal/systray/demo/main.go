package main

import (
	"fmt"
	"github.com/energye/energy/v3/application/internal/systray"
	"github.com/energye/energy/v3/application/internal/systray/demo/icon"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var quitChan = make(chan struct{})

func main() {
	MainRun()
}

func MainRun() {
	systray.NativeStart()
	onReady()
	<-quitChan
	systray.NativeEnd()
}

func addQuitItem() {
	mQuit := systray.AddMenuItem("Quit(退出)", "Quit the whole app")
	mQuit.Enable()
	mQuit.Click(func() {
		fmt.Println("Requesting quit")
		quitChan <- struct{}{}
		fmt.Println("Finished quitting")
	})
}

func onReady() {
	fmt.Println("systray.onReady")
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Energy Sys Tray")
	systray.SetTooltip("Energy tooltip")
	systray.SetOnClick(func() {
		fmt.Println("SetOnClick")
	})
	systray.SetOnDClick(func() {
		fmt.Println("SetOnDClick")
	})
	addQuitItem()

	systray.SetTemplateIcon(icon.Data, icon.Data)
	mChange := systray.AddMenuItem("Change Me", "Change Me")
	mChecked := systray.AddMenuItemCheckbox("Checked", "Check Me", true)
	mEnabled := systray.AddMenuItem("Enabled", "Enabled")
	// Sets the icon of a menu item. Only available on Mac.
	mEnabled.SetTemplateIcon(icon.Data, icon.Data)

	Ignored := systray.AddMenuItem("Ignored", "Ignored")
	Ignored.Click(func() {
		fmt.Println("Ignored click")
	})
	ccdd := Ignored.AddSubMenuItem("选中checked", "bbbb")
	ccdd.Click(func() {
		ccdd.SetChecked(!ccdd.Checked())
	})
	Ignored.AddSeparator()
	abab := Ignored.AddSubMenuItem("清空Ignored", "bbbb")
	abab.Click(func() {
		Ignored.Clear()
	})

	subMenuTop := systray.AddMenuItem("SubMenuTop", "SubMenu Test (top)")
	subMenuMiddle := subMenuTop.AddSubMenuItem("SubMenuMiddle", "SubMenu Test (middle)")
	subMenuBottom := subMenuMiddle.AddSubMenuItemCheckbox("SubMenuBottom - Toggle Panic!", "SubMenu Test (bottom) - Hide/Show Panic!", false)
	subMenuBottom2 := subMenuMiddle.AddSubMenuItem("SubMenuBottom - Panic!", "SubMenu Test (bottom)")
	subMenuBottom2.SetIcon(icon.Data)
	systray.AddSeparator()
	mToggle := systray.AddMenuItem("Toggle", "Toggle some menu items")
	shown := true
	toggle := func() {
		if shown {
			subMenuBottom.SetChecked(true)
			subMenuBottom2.Hide()
			mEnabled.Hide()
			shown = false
			mEnabled.Disable()
		} else {
			subMenuBottom.SetChecked(false)
			subMenuBottom2.Show()
			mEnabled.Show()
			mEnabled.Enable()
			shown = true
		}
	}
	mReset := systray.AddMenuItem("Reset", "Reset all items")

	mChange.Click(func() {
		mChange.SetTitle("I've Changed")
	})
	mChecked.Click(func() {
		if mChecked.Checked() {
			mChecked.SetChecked(false)
			mChecked.SetTitle("Unchecked")
		} else {
			mChecked.SetChecked(true)
			mChecked.SetTitle("Checked")
		}
	})
	mEnabled.Click(func() {
		mEnabled.SetTitle("Disabled")
		fmt.Println("mEnabled.Disabled()", mEnabled.Disabled())
		mEnabled.Disable()
	})
	subMenuBottom2.Click(func() {
		panic("panic button pressed")
	})
	subMenuBottom.Click(func() {
		toggle()
	})
	mReset.Click(func() {
		systray.ResetMenu()
		addQuitItem()
	})
	mToggle.Click(func() {
		toggle()
	})
	// tray icon switch
	go func() {
		var b bool
		// demo: to png full path
		wd, _ := os.Getwd()
		fmt.Println("wd", wd) // /to/icon/path/icon.png, logo.png
		wd = filepath.Join(wd, "application", "internal", "systray", "demo", "icon")
		var ext = ".png"
		if runtime.GOOS == "windows" {
			ext = ".ico" // windows .ico
		}
		icoData, _ := ioutil.ReadFile(filepath.Join(wd, "icon"+ext))
		logoData, _ := ioutil.ReadFile(filepath.Join(wd, "logo"+ext))
		_, _ = icoData, logoData
		for true {
			time.Sleep(time.Second * 1)
			b = !b
			if b {
				systray.SetIcon(logoData)
			} else {
				systray.SetIcon(icoData)
			}
		}
	}()
}
