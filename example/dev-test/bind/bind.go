package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/example/dev-test/bind/src"
	"github.com/energye/energy/logger"
)

//go:embed resources
var resources embed.FS
var cefApp *cef.TCEFApplication

func main() {
	logger.SetEnable(true)
	logger.SetLevel(logger.CefLog_Debug)
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
		_ = bind.NewFunction("JSFunc", src.JSFunc)
		//cef.VariableBind.Bind("varStr", &varStr)

		bind.NewObjects()
	})
	fmt.Println("TestFunc", cef.VariableBind.Bind("TestFunc", TestFunc), "ProcessType", common.Args.ProcessType(), common.Args.IsMain() || common.Args.IsRender())

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
