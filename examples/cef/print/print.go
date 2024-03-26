package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/ipc/callback"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/pkgs/assetserve"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	cefApp := cef.NewApplication()
	//cefApp.SetEnablePrintPreview(false)
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
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		bw := window.AsLCLBrowserWindow().BrowserWindow()
		printDialog := lcl.NewPrintDialog(bw)
		ipc.On("printList", func() []string {
			var result []string
			for i := 0; i < int(lcl.Printer.Printers().Count()); i++ {
				result = append(result, lcl.Printer.Printers().Strings(int32(i)))
			}
			return result
		})
		// lcl.Printer.Printing()// 检查是否打印完成
		//监听事件
		ipc.On("print", func(channel callback.IChannel, type_ int, index int32, text string) {
			fmt.Println("type", type_, index, text)
			if type_ == 3 {
				if printDialog.Execute() {
				}
			} else if type_ == 2 {
				//状态
				lcl.Printer.SetPrinterIndex(index)       // 选择打印机
				lcl.Printer.BeginDoc()                   //开始打印
				lcl.Printer.Canvas().Font().SetSize(100) //设置一些参数
				lcl.Printer.Canvas().TextOut(100, 100, text)
				lcl.Printer.EndDoc() //结束打印
			} else if type_ == 4 {

			} else if type_ == 1 {

			}
		})
	})
	//运行应用
	cef.Run(cefApp)
}
