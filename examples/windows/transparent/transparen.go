//go:build windows
// +build windows

package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/winapi"
	"github.com/cyber-xxm/energy/v2/examples/common"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl/types/colors"
	"github.com/energye/golcl/lcl/win"
)

func main() {
	cef.GlobalInit(nil, common.ResourcesFS())
	cefApp := cef.NewApplication()
	cef.BrowserWindow.Config.Url = "http://localhost:22022/transparent_index.html"
	cef.BrowserWindow.Config.EnableHideCaption = true
	cef.BrowserWindow.Config.AlwaysOnTop = true
	cef.BrowserWindow.Config.Width = 400
	cef.BrowserWindow.Config.Height = 450
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		server.Assets = common.ResourcesFS()
		go server.StartHttpServer()
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		WindowTransparent(types.HWND(window.Handle()))
	})
	cef.Run(cefApp)
}

// WindowTransparent 窗口透明组件不透明设置
func WindowTransparent(hWnd types.HWND) {
	//SetWindowLong（Handle， GWL_EXSTYLE， GetWindowLong（Handle， GWL_EXSTYLE） or WS_EX_LAYERED）
	//SetLayeredWindowAttributes（Handle，clWhite，255，LWA_COLORKEY）;
	exStyle := winapi.GetWindowLong(hWnd, win.GWL_EXSTYLE)
	exStyle = exStyle | win.WS_EX_LAYERED //win.WS_EX_LAYERED&^win.WS_EX_TRANSPARENT // or WS_EX_TRANSPARENT;
	winapi.SetWindowLong(hWnd, win.GWL_EXSTYLE, exStyle)
	win.SetLayeredWindowAttributes(hWnd.ToPtr(), //指定分层窗口句柄
		colors.ClWhite,   //crKey指定需要透明的背景颜色值，可用RGB()宏  0-255
		255,              //bAlpha设置透明度，0表示完全透明，255表示不透明
		win.LWA_COLORKEY) //LWA_ALPHA: crKey无效，bAlpha有效；
	//win.LWA_ALPHA|win.LWA_COLORKEY) //LWA_ALPHA: crKey无效，bAlpha有效；
	//LWA_COLORKEY：窗体中的所有颜色为crKey的地方全透明，bAlpha无效。
	//LWA_ALPHA | LWA_COLORKEY：crKey的地方全透明，其它地方根据bAlpha确定透明度
}
