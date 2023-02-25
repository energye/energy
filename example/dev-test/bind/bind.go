package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/example/dev-test/bind/src"
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
	//变量创建回调-这个函数在主进程(browser)和子进程(render)中执行
	//变量的值绑定到主进程
	cef.VariableBind.VariableCreateCallback(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, bind cef.IProvisionalBindStorage) {
		//初始化要绑定的变量
		//结构类型
		//src.JSStructVarDemo = &src.StructVarDemo{}
		//src.JSStructVarDemo.StringField = "初始的字符串值"
		//bind.NewObjects(src.JSStructVarDemo)
		////通用类型
		//src.JSString = bind.NewString("JSString", "初始的字符串值")
		//src.JSInt = bind.NewInteger("JSInt", 0)
		//src.JSBool = bind.NewBoolean("JSBool", false)
		src.JSDouble = bind.NewDouble("JSDouble", 0.0)
		bind.NewFunction("JSFunc", src.JSFunc)
		//cef.VariableBind.Bind("varStr", &varStr)
		bind.NewObjects()
		cef.VariableBind.Bind("TestFunc", TestFunc)
	})
	fmt.Println("main-ProcessType:", common.Args.ProcessType())
	//fmt.Println("TestFunc", cef.VariableBind.Bind("TestFunc", TestFunc), "ProcessType", common.Args.ProcessType(), common.Args.IsMain() || common.Args.IsRender())

	//fmt.Println("TestFunc", cef.VariableBind.Bind("TestFunc", TestFunc))
	//var err error
	//err = cef.VariableBind.Bind("bindStr", "直接传字符串")
	//fmt.Println("bindStr:", err)
	//err = cef.VariableBind.Bind("varStr", &varStr)
	//fmt.Println("varStr:", err)
	//err = cef.VariableBind.Bind("varInt", &varInt)
	//fmt.Println("varInt:", err)
	//err = cef.VariableBind.Bind("varInt8", &varInt8)
	//fmt.Println("varInt8:", err)
	//err = cef.VariableBind.Bind("varFloat64", &varFloat64)
	//fmt.Println("varFloat64:", err)
	//err = cef.VariableBind.Bind("varFloat32", &varFloat32)
	//fmt.Println("varFloat32:", err)
	//err = cef.VariableBind.Bind("varBool", &varBool)
	//fmt.Println("varBool:", err)

	//Test(&varStr)
	//fmt.Println(varStr)
	//varStr = "asdfadsf"
	//Test(&varStr)
	//fmt.Println(varStr)

	cefApp.SetOnContextCreated(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, context *cef.ICefV8Context) bool {
		fmt.Println("iCefV8ContextPtr", context, "Global.IsValid:", context.Global.IsValid(), context.Global.IsUndefined(), context.Global.GetDateValue())
		fmt.Println("iCefV8ContextPtr GetStringValuer", context.Global.GetStringValue())
		fmt.Println("iCefV8ContextPtr GetValueByIndex", context.Global.GetValueByIndex(0).IsValid())
		fmt.Println("iCefV8ContextPtr GetValueByIndex", context.Global.GetValueByIndex(1).IsValid())
		fmt.Println("iCefV8ContextPtr GetValueByKey", context.Global.GetValueByKey("name").IsValid())
		fmt.Println("iCefV8ContextPtr SetValueByAccessor", context.Global.SetValueByAccessor("nametest", consts.V8_ACCESS_CONTROL_DEFAULT, consts.V8_PROPERTY_ATTRIBUTE_NONE))
		fmt.Println("iCefV8ContextPtr GetExternallyAllocatedMemory", context.Global.GetExternallyAllocatedMemory())
		fmt.Println("iCefV8ContextPtr AdjustExternallyAllocatedMemory", context.Global.AdjustExternallyAllocatedMemory(0))
		fmt.Println("iCefV8ContextPtr GetExternallyAllocatedMemory.GetValueByKey", context.Global.GetValueByKey("name").GetExternallyAllocatedMemory())
		fmt.Println("iCefV8ContextPtr GetFunctionName", context.Global.GetFunctionName())
		fmt.Println("V8ValueRef IsValid", cef.V8ValueRef.NewUndefined().IsValid())
		handler := cef.V8HandlerRef.New()
		accessor := cef.V8AccessorRef.New()
		fmt.Println("handler-accessor:", handler, accessor)
		accessor.Get(func(name string, object *cef.ICefV8Value, retVal *cef.ResultV8Value, exception *cef.Exception) bool {
			retVal.SetResult(cef.V8ValueRef.NewString("这能返回？"))
			return true
		})
		accessor.Set(func(name string, object *cef.ICefV8Value, value *cef.ICefV8Value, exception *cef.Exception) bool {
			fmt.Println("name", name, "object.IsValid", object.IsValid(), object.IsObject(), object.IsString(), "value.IsValid", value.IsValid(), value.IsString(), value.IsObject())
			return true
		})
		handler.Execute(func(name string, object *cef.ICefV8Value, arguments *cef.TCefV8ValueArray, retVal *cef.ResultV8Value, exception *cef.Exception) bool {
			fmt.Println("handler", arguments.Size(), arguments.Get(3))
			fmt.Println(arguments.Get(0).IsValid(), arguments.Get(0).GetStringValue())
			fmt.Println(arguments.Get(1).IsValid(), arguments.Get(1).GetIntValue())
			retVal.SetResult(cef.V8ValueRef.NewString("函数返回值？"))
			val, ex, ok := context.Eval("fntest();", "", 0)
			fmt.Println("Execute eval fntest:", val, ex, ok)
			fmt.Println("Execute eval fntest-return-value:", val.GetStringValue())
			if ok {
				fmt.Println(val.GetStringValue())
			}
			val, ex, ok = context.Eval("errtest();", "", 0)
			fmt.Println("Execute errtest:", val, ex, ok)
			fmt.Println("Execute errtest-error:", ex.Message(), ex.LineNumber())
			return true
		})
		object := cef.V8ValueRef.NewObject(accessor, nil)
		function := cef.V8ValueRef.NewFunction("testfn", handler)
		fmt.Println("V8ValueRef NewObject", object, object.IsValid())
		object.SetValueByAccessor("testcca", consts.V8_ACCESS_CONTROL_DEFAULT, consts.V8_PROPERTY_ATTRIBUTE_NONE)
		object.SetValueByKey("testcca", cef.V8ValueRef.NewObject(accessor, nil), consts.V8_PROPERTY_ATTRIBUTE_NONE)
		object.SetValueByKey("testcca", cef.V8ValueRef.NewObject(accessor, nil), consts.V8_PROPERTY_ATTRIBUTE_NONE)
		object.SetValueByKey("testfn", function, consts.V8_PROPERTY_ATTRIBUTE_NONE)
		fmt.Println("Global.SetValueByKey", context.Global.SetValueByKey("testset", object, consts.V8_PROPERTY_ATTRIBUTE_READONLY))
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
		context.Global.SetValueByKey("arrBuf", buffer, consts.V8_PROPERTY_ATTRIBUTE_NONE)
		return false
	})

	//运行应用
	cef.Run(cefApp)
}

var (
	varStr             = "字符串变量"
	varInt     int64   = 21
	varInt8    int8    = 21
	varFloat64         = 333.55
	varFloat32 float32 = 222.11
	varBool            = true
)

func TestFunc() string {
	fmt.Println("TestFunc")
	return "test" + "afdasdfsadf: "
}

type TestStruct struct {
	Name string
	Age  int
}

func Test(bind interface{}) {
	var s = bind.(*string)
	println("====", *s)
	*s = "123"
}
