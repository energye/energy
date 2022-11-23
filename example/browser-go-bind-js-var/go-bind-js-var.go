package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/example/browser-go-bind-js-var/src"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalCEFInit(nil, &resources)
	//创建应用
	cefApp := cef.NewApplication(nil)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.DefaultUrl = "http://localhost:22022/go-bind-js-var.html"
	cef.BrowserWindow.Config.Title = "Energy - execute-javascript"
	cef.BrowserWindow.Config.Icon = "resources/icon.ico"
	//内置http服务链接安全配置
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = &resources
		go server.StartHttpServer()
	})
	//变量创建回调-这个函数在主进程(browser)和子进程(render)中执行
	//变量的值绑定到主进程
	cef.VariableBind.VariableCreateCallback(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, bind cef.IProvisionalBindStorage) {
		//初始化要绑定的变量
		//结构类型
		src.JSStructVarDemo = &src.StructVarDemo{}
		src.JSStructVarDemo.StringField = "初始的字符串值"
		bind.NewObjects(src.JSStructVarDemo)
		//通用类型
		src.JSString = bind.NewString("JSString", "初始的字符串值")
		src.JSInt = bind.NewInteger("JSInt", 0)
		src.JSBool = bind.NewBoolean("JSBool", false)
		src.JSDouble = bind.NewDouble("JSDouble", 0.0)
		_ = bind.NewFunction("JSFunc", src.JSFunc)
	})
	//运行应用
	cef.Run(cefApp)
}
