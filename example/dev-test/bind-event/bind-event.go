package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/cef/bind"
	"github.com/energye/energy/common/assetserve"
	//_ "net/http/pprof"
)

//go:embed resources
var resources embed.FS
var cefApp *cef.TCEFApplication

func main() {
	//go func() {
	//	http.ListenAndServe(":10000", nil)
	//}()
	//logger.SetEnable(true)
	//logger.SetLevel(logger.CefLog_Debug)
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	cefApp = cef.NewApplication()
	//cefApp.SetLogSeverity(consts.LOGSEVERITY_DEBUG)
	//cefApp.SetSingleProcess(true)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/bind-event.html"
	cef.BrowserWindow.Config.Title = "Energy - bind-event"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"

	stringKey := bind.NewString("stringKey", "字符串值")
	fmt.Println("stringKey", stringKey, stringKey.IsString())
	stringKey.SetValue(9999)
	fmt.Println("stringKey", stringKey, stringKey.IsString(), stringKey.IsInteger())
	integerKey := bind.NewInteger("integerKey", 1000)
	fmt.Println("integerKey", integerKey)
	doubleKey := bind.NewDouble("doubleKey", 1000.001)
	fmt.Println("doubleKey", doubleKey)
	booleanKey := bind.NewBoolean("booleanKey", true)
	fmt.Println("booleanKey", booleanKey.Value())
	nullKey := bind.NewNull("nullKey")
	fmt.Println("nullKey", nullKey.Value())
	undefinedKey := bind.NewUndefined("undefinedKey")
	fmt.Println("undefinedKey", undefinedKey.Value(), undefinedKey.IsUndefined())
	funcKey := bind.NewFunction("funcKey", func() {

	})
	fmt.Println("funcKey", funcKey)

	type objectDemo1 struct {
		Key1 string
		Key2 string
	}
	type objectDemo2 struct {
		//Key1 string
		//Key2 string
		//Key3 int
		Key4 *objectDemo1
		Key5 []*objectDemo1
	}
	type object struct {
		//Key1 string
		//Key2 string
		//Key3 int
		//Key4 float64
		//Key5 bool
		//Key6 *objectDemo1
		Key7  objectDemo2
		Key8  []string
		Key9  []objectDemo2
		Key10 *objectDemo2
	}
	var testObj = &object{
		//Key1: "value1",
		//Key2: "value2",
		//Key3: 333,
		//Key4: 555.3,
		//Key5: true,
		//Key6: &objectDemo1{},
	}
	bind.NewObject(testObj)
	bind.NewArray("arrayKey", "字符串", 100001, 22222.333, true, testObj)

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
	//运行应用
	cef.Run(cefApp)
}
