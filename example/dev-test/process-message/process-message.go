package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"time"
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
	cefApp = cef.NewApplication()
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
		go func() {
			var i = 0
			for {
				time.Sleep(time.Second)
				fmt.Println("GetWindowInfos:", len(cef.BrowserWindow.GetWindowInfos()), cefApp.ProcessTypeValue())
				gwi := cef.BrowserWindow.GetWindowInfo(1)
				if gwi == nil {
					continue
				}
				browser := gwi.Browser()
				if browser == nil {
					continue
				}
				frame := browser.MainFrame()
				if frame == nil {
					continue
				}
				fmt.Println("browser", browser.Identifier(), frame.Instance())
				sendBrowserProcessMsg := cef.ProcessMessageRef.New("testName")
				sendBrowserProcessMsg.ArgumentList().SetString(0, fmt.Sprintf("发送给渲染进程, 测试值 %d ==== %d", time.Now().Second(), i))
				frame.SendProcessMessage(consts.PID_RENDER, sendBrowserProcessMsg)
				//主进程, 执行指定窗口的JS
				cef.BrowserWindow.GetWindowInfo(1).Chromium().ExecuteJavaScript("console.log('1111111');", "", 0)
				i++
			}
		}()
	})
	cefApp.SetOnWebKitInitialized(func() {
		fmt.Println("SetOnWebKitInitialized")
		v8Handler := cef.V8HandlerRef.New()
		v8Handler.Execute(func(name string, object *cef.ICefV8Value, arguments *cef.TCefV8ValueArray, retVal *cef.ResultV8Value, exception *cef.ResultString) bool {
			fmt.Println("v8Handler.Execute", name)
			return true
		})
		//注册js
		var jsCode = `
	let codeobj = {};
	codeobj.test = function(){
		document.getElementById("test").style.color="red";
	}
`
		cef.RegisterExtension("v8/codeobj", jsCode, v8Handler)
	})
	cefApp.SetOnContextCreated(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, context *cef.ICefV8Context) bool {
		handler := cef.V8HandlerRef.New()
		fmt.Println("handler:", handler, "frameId:", frame.Identifier())
		handler.Execute(func(name string, object *cef.ICefV8Value, arguments *cef.TCefV8ValueArray, retVal *cef.ResultV8Value, exception *cef.ResultString) bool {
			fmt.Println("handler.Execute", name)
			retVal.SetResult(cef.V8ValueRef.NewString("函数返回值？"))
			message := cef.ProcessMessageRef.New("testname")
			fmt.Println("ProcessMessageRef IsValid", message.IsValid(), message.Name())
			list := message.ArgumentList()
			list.SetString(0, "值？")
			fmt.Println("list IsValid", list.IsValid(), list.Size(), list.GetString(0))
			listCopy := list.Copy()
			fmt.Println("listCopy GetString", listCopy.IsValid(), list.Size(), list.GetString(0), list.GetValue(0).GetType())
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
			fmt.Println("DictionaryValueRef IsValid", dictionaryValue.IsValid(), dictionaryValue.Size(), dictionaryValue.GetString("strdicttest"), dictionaryValue.GetDouble("doubledicttest"))
			listCopy.SetDictionary(2, dictionaryValue)
			dictionaryValue = listCopy.GetDictionary(2)
			fmt.Println("DictionaryValueRef IsValid", dictionaryValue.IsValid(), dictionaryValue.Size(), dictionaryValue.GetDouble("doubledicttest"))
			//list.SetDictionary()
			//测试 - 给 browser 程发送消息
			fmt.Println("Execute.frameId", cef.V8ContextRef.Current().Frame().Identifier())
			sendBrowserProcessMsg := cef.ProcessMessageRef.New("testName")
			sendBrowserProcessMsg.ArgumentList().SetString(0, "发送给主进程, 测试值")
			//Execute 函数实际已脱离上下文，不能直接使用上下文传入的frame.
			//通过V8Context获取当前Frame, 或 browser.MainFrame()
			cef.V8ContextRef.Current().Frame().SendProcessMessage(consts.PID_BROWSER, sendBrowserProcessMsg)
			return true
		})
		object := cef.V8ValueRef.NewObject(nil)
		function := cef.V8ValueRef.NewFunction("testfn", handler)
		object.SetValueByKey("testfn", function, consts.V8_PROPERTY_ATTRIBUTE_NONE)
		context.Global().SetValueByKey("testset", object, consts.V8_PROPERTY_ATTRIBUTE_READONLY)
		return false
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		event.SetOnBrowseProcessMessageReceived(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, sourceProcess consts.CefProcessId, message *cef.ICefProcessMessage) bool {
			fmt.Println("browser 进程接收消息", message.Name(), message.ArgumentList().GetString(0), frame.Instance())
			//测试 - 给 render 进程发送消息
			//sendBrowserProcessMsg := cef.ProcessMessageRef.New("rnderTestName")
			//sendBrowserProcessMsg.ArgumentList().SetString(0, "发送给渲染进程, 测试值")
			//frame.SendProcessMessage(consts.PID_RENDER, sendBrowserProcessMsg)
			//sendBrowserProcessMsg.Free()
			return false
		})
	})
	cefApp.SetOnProcessMessageReceived(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, sourceProcess consts.CefProcessId, message *cef.ICefProcessMessage) bool {
		fmt.Println("render 进程接收消息", message.Name(), message.ArgumentList().GetString(0), frame.Instance())
		sendBrowserProcessMsg := cef.ProcessMessageRef.New("browserTestName")
		sendBrowserProcessMsg.ArgumentList().SetString(0, "发送给主进程, 测试值")
		frame.SendProcessMessage(consts.PID_BROWSER, sendBrowserProcessMsg)
		sendBrowserProcessMsg.Free()
		return false
	})
	//运行应用
	cef.Run(cefApp)
}
