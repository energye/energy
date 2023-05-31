package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/cef/process"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
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
	cef.BrowserWindow.Config.Url = "http://localhost:22022/bind.html"
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
	fmt.Println("main-ProcessType:", process.Args.ProcessType())

	cefApp.SetOnContextCreated(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, context *cef.ICefV8Context) bool {
		fmt.Println("iCefV8ContextPtr", context, "Global.IsValid:", context.Global().IsValid(), context.Global().IsUndefined(), context.Global().GetDateValue())
		fmt.Println("iCefV8ContextPtr GetStringValuer", context.Global().GetStringValue())
		fmt.Println("iCefV8ContextPtr GetValueByIndex", context.Global().GetValueByIndex(0).IsValid())
		fmt.Println("iCefV8ContextPtr GetValueByIndex", context.Global().GetValueByIndex(1).IsValid())
		fmt.Println("iCefV8ContextPtr GetValueByKey", context.Global().GetValueByKey("name").IsValid())
		fmt.Println("iCefV8ContextPtr SetValueByAccessor", context.Global().SetValueByAccessor("nametest", consts.V8_ACCESS_CONTROL_DEFAULT, consts.V8_PROPERTY_ATTRIBUTE_NONE))
		fmt.Println("iCefV8ContextPtr GetExternallyAllocatedMemory", context.Global().GetExternallyAllocatedMemory())
		fmt.Println("iCefV8ContextPtr AdjustExternallyAllocatedMemory", context.Global().AdjustExternallyAllocatedMemory(0))
		fmt.Println("iCefV8ContextPtr GetExternallyAllocatedMemory.GetValueByKey", context.Global().GetValueByKey("name").GetExternallyAllocatedMemory())
		fmt.Println("iCefV8ContextPtr GetFunctionName", context.Global().GetFunctionName())
		fmt.Println("V8ValueRef IsValid", cef.V8ValueRef.NewUndefined().IsValid())
		handler := cef.V8HandlerRef.New()
		accessor := cef.V8AccessorRef.New()
		fmt.Println("handler-accessor:", handler, accessor)
		accessor.Get(func(name string, object *cef.ICefV8Value, retVal *cef.ResultV8Value, exception *cef.ResultString) bool {
			fmt.Println("accessor get name", name)
			retVal.SetResult(cef.V8ValueRef.NewString("这能返回？"))
			return true
		})
		accessor.Set(func(name string, object *cef.ICefV8Value, value *cef.ICefV8Value, exception *cef.ResultString) bool {
			fmt.Println("accessor set name", name, "object.IsValid", object.IsValid(), object.IsObject(), object.IsString(), "value.IsValid", value.IsValid(), value.IsString(), value.IsObject())
			return true
		})
		handler.Execute(func(name string, object *cef.ICefV8Value, arguments *cef.TCefV8ValueArray, retVal *cef.ResultV8Value, exception *cef.ResultString) bool {
			fmt.Println("handler.Execute", arguments.Size(), arguments.Get(3))
			fmt.Println(arguments.Get(0).IsValid(), arguments.Get(0).GetStringValue())
			fmt.Println(arguments.Get(1).IsValid(), arguments.Get(1).GetIntValue())
			retVal.SetResult(cef.V8ValueRef.NewString("函数返回值？"))
			val, ex, ok := cef.V8ContextRef.Current().Eval("fntest()", "", 0)
			fmt.Println("? = Execute eval fntest:", val, ex, ok)
			if ok {
				fmt.Println("OK = Execute eval fntest-return-value:", val.GetStringValue())
			}
			val, ex, ok = cef.V8ContextRef.Current().Eval("errtest();", "", 0)
			fmt.Println("? = Execute errtest:", val, ex, ok)
			fmt.Println("OK = Execute errtest-error:", ex.Message(), ex.LineNumber())
			fmt.Println("V8ContextRef.Current()", cef.V8ContextRef.Current().Global().GetValueByKey("arrBuf").IsArrayBuffer())
			fmt.Println("V8ContextRef.Entered()", cef.V8ContextRef.Entered().Global().GetValueByKey("arrBuf").IsArrayBuffer())
			return true
		})
		object := cef.V8ValueRef.NewObject(accessor)
		function := cef.V8ValueRef.NewFunction("testfn", handler)
		fmt.Println("V8ValueRef NewObject", object, object.IsValid())
		object.SetValueByAccessor("testcca", consts.V8_ACCESS_CONTROL_DEFAULT, consts.V8_PROPERTY_ATTRIBUTE_NONE)
		object.SetValueByKey("testcca", cef.V8ValueRef.NewObject(accessor), consts.V8_PROPERTY_ATTRIBUTE_NONE)
		object.SetValueByKey("testcca", cef.V8ValueRef.NewObject(accessor), consts.V8_PROPERTY_ATTRIBUTE_NONE)
		object.SetValueByKey("testfn", function, consts.V8_PROPERTY_ATTRIBUTE_NONE)
		fmt.Println("Global.SetValueByKey", context.Global().SetValueByKey("testset", object, consts.V8_PROPERTY_ATTRIBUTE_READONLY))
		fmt.Println("GetFunctionHandler", function.GetFunctionHandler())
		val, exception, ok := context.Eval("console.log('evalaa');", "", 0)
		fmt.Println("eval:", val, exception, ok)
		fmt.Println("eval-return-value:", val.GetStringValue())
		array := cef.V8ValueRef.NewArray(1024)
		fmt.Println("数据有效", array.IsValid())
		fmt.Println("array-GetArrayLength", array.GetArrayLength())
		array.SetValueByIndex(0, cef.V8ValueRef.NewString("数组里的值"))
		fmt.Println("array-GetValueByIndex", array.GetValueByIndex(0).GetStringValue())
		buf := make([]byte, 256)
		for i := 0; i < len(buf); i++ {
			buf[i] = byte(i)
		}
		fmt.Println(buf)
		callback := cef.V8ArrayBufferReleaseCallbackRef.New()
		callback.ReleaseBuffer(func(buffer uintptr) bool {
			fmt.Println("释放？")
			return true
		})
		buffer := cef.V8ValueRef.NewArrayBuffer(buf, callback)
		fmt.Println("ArrayBuffer IsValid", buffer.IsValid())
		context.Global().SetValueByKey("arrBuf", buffer, consts.V8_PROPERTY_ATTRIBUTE_NONE)
		return false
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {

		event.SetOnBeforeResourceLoad(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, request *cef.ICefRequest, callback *cef.ICefCallback, result *consts.TCefReturnValue) {
			fmt.Println("SetOnBeforeResourceLoad:", request.URL(), request.Method(), "headerMap:", request.GetHeaderMap().GetSize())
			headerMap := request.GetHeaderMap()
			fmt.Println("\t", request.GetHeaderByName("energy"), headerMap.GetEnumerate("energy", 1), "size:", headerMap.GetSize())
			for i := 0; i < int(headerMap.GetSize()); i++ {
				fmt.Println("\tkey:", headerMap.GetKey(uint32(i)), "value:", headerMap.GetValue(uint32(i)))
			}
			multiMap := cef.StringMultiMapRef.New()
			fmt.Println("multiMap.GetSize()", multiMap.GetSize())
			multiMap.Append("key1", "value1")
			fmt.Println("multiMap.GetSize()", multiMap.GetSize())
			//postData := cef.PostDataRef.New()
			//postData.
			fmt.Println("GetPostData().GetElementCount", request.GetPostData().IsValid())
			if !request.GetPostData().IsValid() {
				data := cef.PostDataRef.New()
				postDataElement := cef.PostDataElementRef.New()
				//postDataElement.SetToFile("7e9fac0f30c829738cc3ad8a69da97ba.txt")
				data.AddElement(postDataElement)
				postDataElement = cef.PostDataElementRef.New()
				bytes := make([]byte, 256, 256)
				for i := 0; i < len(bytes); i++ {
					bytes[i] = byte(i)
				}
				fmt.Println("postDataElement.SetToBytes", bytes)
				postDataElement.SetToBytes(bytes)
				data.AddElement(postDataElement)
				request.SetPostData(data)
				fmt.Println("\tGetPostData GetElementCount:", request.GetPostData().IsValid(), request.GetPostData().GetElementCount(), data.IsReadOnly())
				fmt.Println("\tGetElements Size:", request.GetPostData().GetElements().Size())
				fmt.Println("\tGetElements GetFile:", request.GetPostData().GetElements().Get(0).GetFile())
				fmt.Println("\tGetElements GetBytesCount:", request.GetPostData().GetElements().Get(0).GetBytesCount())
				postDataElement = request.GetPostData().GetElements().Get(1)
				fmt.Println("\tGetElements GetBytesCount:", postDataElement.GetBytesCount())
				bytes, count := postDataElement.GetBytes()
				fmt.Println("\tGetElements bytes,count:", bytes, count)
			}
		})
	})

	//运行应用
	cef.Run(cefApp)
}
