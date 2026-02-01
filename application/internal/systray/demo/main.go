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

var (
	quitChan = make(chan struct{})
	tray     *systray.Tray
)

func main() {
	tray = systray.NativeStart()
	onReady()
	<-quitChan
	tray.NativeEnd()
}

func addQuitItem() {
	mQuit := tray.Menu().AddMenuItem(tray, "Quit(退出)", "Quit the whole app")
	mQuit.Enable()
	mQuit.Click(func() {
		fmt.Println("Requesting quit")
		quitChan <- struct{}{}
		fmt.Println("Finished quitting")
	})
}

func onReady() {
	fmt.Println("systray.onReady")
	tray.SetIcon(icon.Data)
	tray.SetTitle("Energy Sys Tray")
	tray.SetTooltip("Energy tooltip")
	tray.SetOnClick(func() {
		fmt.Println("SetOnClick")
	})
	tray.SetOnDClick(func() {
		fmt.Println("SetOnDClick")
	})
	addQuitItem()
	mChange := tray.Menu().AddMenuItem(tray, "Change Me", "Change Me")
	mChecked := tray.Menu().AddMenuItem(tray, "Checked", "Check Me")
	mChecked.SetChecked(tray, true)
	mEnabled := tray.Menu().AddMenuItem(tray, "Enabled", "Enabled")
	// Sets the icon of a menu item. Only available on Mac.
	mEnabled.SetIcon(tray, icon.Data)

	Ignored := tray.Menu().AddMenuItem(tray, "Ignored", "Ignored")
	Ignored.Click(func() {
		fmt.Println("Ignored click")
	})
	ccdd := Ignored.AddMenuItem(tray, "选中checked", "bbbb")
	ccdd.Click(func() {
		ccdd.SetChecked(tray, !ccdd.Checked())
	})
	Ignored.AddSeparator(tray)
	abab := Ignored.AddMenuItem(tray, "清空Ignored", "bbbb")
	abab.Click(func() {
		Ignored.Clear(tray)
	})

	subMenuTop := tray.Menu().AddMenuItem(tray, "SubMenuTop", "SubMenu Test (top)")
	subMenuMiddle := subMenuTop.AddMenuItem(tray, "SubMenuMiddle", "SubMenu Test (middle)")
	subMenuBottom := subMenuMiddle.AddMenuItem(tray, "SubMenuBottom - Toggle Panic!", "SubMenu Test (bottom) - Hide/Show Panic!")
	subMenuBottom.SetChecked(tray, true)
	subMenuBottom2 := subMenuMiddle.AddMenuItem(tray, "SubMenuBottom - Panic!", "SubMenu Test (bottom)")
	subMenuBottom2.SetIcon(tray, icon.Data)
	tray.Menu().AddSeparator(tray)
	mToggle := tray.Menu().AddMenuItem(tray, "Toggle", "Toggle some menu items")
	shown := true
	toggle := func() {
		if shown {
			subMenuBottom.SetChecked(tray, true)
			subMenuBottom2.Hide(tray)
			mEnabled.Hide(tray)
			shown = false
			mEnabled.SetEnable(tray, true)
		} else {
			subMenuBottom.SetChecked(tray, false)
			subMenuBottom2.Show(tray)
			mEnabled.Show(tray)
			mEnabled.SetEnable(tray, false)
			shown = true
		}
	}
	mReset := tray.Menu().AddMenuItem(tray, "Reset", "Reset all items")

	mChange.Click(func() {
		mChange.SetTitle(tray, "I've Changed")
	})
	mChecked.Click(func() {
		if mChecked.Checked() {
			mChecked.SetChecked(tray, false)
			mChecked.SetTitle(tray, "Unchecked")
		} else {
			mChecked.SetChecked(tray, true)
			mChecked.SetTitle(tray, "Checked")
		}
	})
	mEnabled.Click(func() {
		mEnabled.SetTitle(tray, "Disabled")
		fmt.Println("mEnabled.Disabled()", mEnabled.Enable())
		mEnabled.SetEnable(tray, false)
	})
	subMenuBottom2.Click(func() {
		panic("panic button pressed")
	})
	subMenuBottom.Click(func() {
		toggle()
	})
	mReset.Click(func() {
		tray.ResetMenu()
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
				tray.SetIcon(logoData)
			} else {
				tray.SetIcon(icoData)
			}
		}
	}()
}
