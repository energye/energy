package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/consts"
)

//go:embed resources
var resources embed.FS
var cefApp *cef.TCEFApplication

func main() {
	//logger.SetEnable(true)
	//logger.SetLevel(logger.CefLog_Debug)
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	cefApp = cef.NewApplication(nil)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/process-message.html"
	cef.BrowserWindow.Config.Title = "Energy - execute-javascript"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
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
	cefApp.SetOnContextCreated(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, context *cef.ICefV8Context) bool {
		handler := cef.V8HandlerRef.New()
		fmt.Println("handler:", handler)
		handler.Execute(func(name string, object *cef.ICefV8Value, arguments *cef.TCefV8ValueArray, retVal *cef.ResultV8Value, exception *cef.Exception) bool {
			fmt.Println("handler.Execute", name)
			retVal.SetResult(cef.V8ValueRef.NewString("函数返回值？"))
			message := cef.ProcessMessage.New("testname")
			fmt.Println("ProcessMessageRef IsValid", message.IsValid(), message.Name())
			list := message.ArgumentList()
			list.SetString(0, "值？")
			fmt.Println("list IsValid", list.IsValid(), list.GetSize(), list.GetString(0))
			listCopy := list.Copy()
			fmt.Println("listCopy GetString", listCopy.IsValid(), list.GetSize(), list.GetString(0), list.GetValue(0).GetType())
			listCopy.SetDouble(1, 112211.22)
			fmt.Println("listCopy GetDouble", listCopy.GetDouble(1), listCopy.GetType(1))
			data := make([]byte, 1024, 1024)
			for i := 0; i < len(data); i++ {
				data[i] = byte(i % 255)
			}
			value := cef.BinaryValueRef.New(data)
			fmt.Println("BinaryValueNew IsValid", value.IsValid())
			fmt.Println("BinaryValueNew size", value.GetSize())
			buf := make([]byte, 300)
			fmt.Println("BinaryValueNew GetData", value.GetData(buf, 0))
			fmt.Println("BinaryValueNew GetData buf", buf)
			dictionaryValue := cef.DictionaryValueRef.New()
			dictionaryValue.SetString("strdicttest", "字符串？")
			dictionaryValue.SetDouble("doubledicttest", 9999.666)
			fmt.Println("DictionaryValueRef IsValid", dictionaryValue.IsValid(), dictionaryValue.GetSize(), dictionaryValue.GetString("strdicttest"), dictionaryValue.GetDouble("doubledicttest"))
			listCopy.SetDictionary(2, dictionaryValue)
			dictionaryValue = listCopy.GetDictionary(2)
			fmt.Println("DictionaryValueRef IsValid", dictionaryValue.IsValid(), dictionaryValue.GetSize(), dictionaryValue.GetDouble("doubledicttest"))
			//list.SetDictionary()
			return true
		})
		//
		object := cef.V8ValueRef.NewObject(nil)
		function := cef.V8ValueRef.NewFunction("testfn", handler)
		object.SetValueByKey("testfn", function, consts.V8_PROPERTY_ATTRIBUTE_NONE)
		context.Global.SetValueByKey("testset", object, consts.V8_PROPERTY_ATTRIBUTE_READONLY)
		return false
	})

	//运行应用
	cef.Run(cefApp)
}
