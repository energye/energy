package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
)

//go:embed resources
var resources embed.FS

func main() {
	cef.GlobalInit(nil, nil)
	cefApp := cef.NewApplication()
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		server.Assets = &resources
		go server.StartHttpServer()
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		//WindowTransparent(types.HWND(window.Handle()))
		ipc.On("testRunFile", func() {
			acceptFilters := lcl.NewStringList()
			acceptFilters.Add(".png")
			callback := cef.RunFileDialogCallbackRef.New()
			callback.SetOnFileDialogDismissed(func(filePaths *lcl.TStrings) {
				for i := 0; i < int(filePaths.Count()); i++ {
					path := filePaths.Strings(int32(i))
					fmt.Println(path)
				}
			})
			window.Chromium().Browser().RunFileDialog(consts.FILE_DIALOG_SAVE, "打开窗口", "", nil, callback)
			fmt.Println("RunFileDialog end")
		})
	})
	cef.Run(cefApp)
}
