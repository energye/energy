package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/example/dev-test/ipc-event/src"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/types"
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
	cef.BrowserWindow.Config.Url = "http://localhost:22022/ipc-event.html"
	//cef.BrowserWindow.Config.Url = "https://map.baidu.com/"
	cef.BrowserWindow.Config.Title = "Energy - ipc-event"
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
	ipc.On("testEmitName", func(context ipc.IContext) {
		argument := context.ArgumentList()
		fmt.Println("testEmitName", argument.Size(), context.BrowserId(), context.FrameId())
		for i := 0; i < int(argument.Size()); i++ {
			value := argument.GetIValue(types.NativeUInt(i))
			fmt.Println("\tGetType:", i, value.GetType())
			switch value.GetType() {
			case consts.VTYPE_NULL:
				//null
			case consts.VTYPE_BOOL:
				value.GetBool()
			case consts.VTYPE_INT:
				value.GetInt()
			case consts.VTYPE_DOUBLE:
				value.GetDouble()
			case consts.VTYPE_STRING:
				value.GetString()
			case consts.VTYPE_DICTIONARY: // object
				object := value.GetIObject()
				fmt.Println("object", object.GetIKeys().Count())
			case consts.VTYPE_LIST: // array
				value.GetIArray()
			}
		}
	})

	//cefApp.SetSingleProcess(true)
	//cefApp.SetEnableGPU(true)
	cef.VariableBind.Bind("funcName", func() string {
		return ""
	})
	var stringField = "stringField"
	cef.VariableBind.Bind("stringField", &stringField)
	var intField = 100
	cef.VariableBind.Bind("intField", &intField)
	var doubleField = 900.001
	cef.VariableBind.Bind("doubleField", &doubleField)
	var boolField = true
	cef.VariableBind.Bind("boolField", &boolField)
	cef.VariableBind.Bind("structField", src.StructField)
	//运行应用
	cef.Run(cefApp)
}
