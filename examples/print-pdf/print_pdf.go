package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/cef/ipc/context"
	"github.com/cyber-xxm/energy/v2/consts"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"path"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	cefApp := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.BrowserWindow.Config.Title = "Energy 打印PFD预览"
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = resources
		go server.StartHttpServer()
	})
	wd := consts.CurrentExecuteDir
	//监听事件
	ipc.On("print-pdf", func(context context.IContext) {
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		savePath := path.Join(wd, "examples", "print-pdf", "test.pdf")
		fmt.Println("当前页面保存为PDF", savePath)
		//bw.Chromium().PrintToPDF(savePath)
		settings := cef.TCefPdfPrintSettings{
			DisplayHeaderFooter: 1,
			PaperWidth:          10,
			PaperHeight:         10,
			PreferCssPageSize:   100,
			FooterTemplate:      "FooterTemplate",
			HeaderTemplate:      `HeaderTemplate`,
		}
		callback := cef.PdfPrintCallbackRef.New()
		callback.OnPdfPrintFinished(func(path string, ok bool) {
			fmt.Println("path:", path, "ok:", ok)
			callback.Free()
		})
		bw.Chromium().Browser().PrintToPdf(savePath, settings, callback)
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		window.Chromium().SetOnPdfPrintFinished(func(sender lcl.IObject, ok bool) {
			fmt.Println("OnPdfPrintFinished:", ok)
		})
	})
	//运行应用
	cef.Run(cefApp)
}
