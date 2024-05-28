package main

import (
	"fmt"
	"github.com/energye/energy/v2/api/libname"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
	"github.com/energye/energy/v2/types/messages"
	"github.com/energye/energy/v2/wv"
	"path/filepath"
)

var mainForm TMainForm
var load wv.IWVLoader

func main() {
	fmt.Println("Go ENERGY Run Main")
	wv.Init(nil, nil)
	// GlobalWebView2Loader
	load = wv.GlobalWebView2Loader(nil)
	liblcl := libname.LibName
	webView2Loader, _ := filepath.Split(liblcl)
	webView2Loader = filepath.Join(webView2Loader, "WebView2Loader.dll")
	fmt.Println("当前目录:", types.CurrentExecuteDir)
	fmt.Println("liblcl.dll目录:", liblcl)
	fmt.Println("WebView2Loader.dll目录:", webView2Loader)
	fmt.Println("用户缓存目录:", filepath.Join(types.CurrentExecuteDir, "EnergyCache"))
	load.SetUserDataFolder(filepath.Join(types.CurrentExecuteDir, "EnergyCache"))
	load.SetLoaderDllPath(webView2Loader)
	r := load.StartWebView2()
	fmt.Println("StartWebView2", r)
	// 底层库全局异常
	lcl.Application.SetOnException(func(sender lcl.IObject, e lcl.IException) {
		fmt.Println("全局-底层库异常:", e.ToString())
	})
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&mainForm)
	lcl.Application.Run()
}

type TMainForm struct {
	lcl.TForm
	windowParent wv.IWVWindowParent
	browser      wv.IWVBrowser
}

func (m *TMainForm) FormCreate(sender lcl.IObject) {
	m.SetCaption("html5test")
	m.SetPosition(types.PoScreenCenter)
	m.SetWidth(1024)
	m.SetHeight(768)
	m.SetDoubleBuffered(true)

	m.windowParent = wv.NewWVWindowParent(m)
	m.windowParent.SetParent(m)
	m.windowParent.SetAlign(types.AlClient)

	m.browser = wv.NewWVBrowser(m)
	m.browser.SetDefaultURL("https://html5test.opensuse.org")
	m.browser.SetOnAfterCreated(func(sender lcl.IObject) {
		m.windowParent.UpdateSize()
	})
	// 设置browser到window parent
	m.windowParent.SetBrowser(m.browser)

	// 窗口显示时创建browser
	m.SetOnShow(func(sender lcl.IObject) {
		if load.InitializationError() {
			fmt.Println("回调函数 => SetOnShow 初始化失败")
		} else {
			if load.Initialized() {
				fmt.Println("回调函数 => SetOnShow 初始化成功")
				m.browser.CreateBrowser(m.windowParent.Handle(), true)
			}
		}
	})
	m.SetOnWndProc(func(msg *types.TMessage) {
		m.InheritedWndProc(msg)
		switch msg.Msg {
		case messages.WM_SIZE, messages.WM_MOVE, messages.WM_MOVING:
			m.browser.NotifyParentWindowPositionChanged()
		}
	})
	m.SetOnDestroy(m.OnFormDestroy)
}

func (m *TMainForm) OnFormDestroy(sender lcl.IObject) {
	fmt.Println("OnFormDestroy")
}
