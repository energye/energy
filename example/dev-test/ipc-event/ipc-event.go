package main

import (
	"embed"
	"errors"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/example/dev-test/ipc-event/src"
	"github.com/energye/energy/ipc"
	"net/http"
	_ "net/http/pprof"
)

//go:embed resources
var resources embed.FS
var cefApp *cef.TCEFApplication

func main() {
	go func() {
		fmt.Println("pprof")
		http.ListenAndServe(":9999", nil)
	}()
	//logger.SetEnable(true)
	//logger.SetLevel(logger.CefLog_Debug)
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	cefApp = cef.NewApplication()
	//cefApp.SetSingleProcess(true)
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
	ipc.OnArguments("testResultArgs", func(args1 int) (r1 map[string]string, err error) {
		fmt.Println("args1", args1)
		r1 = make(map[string]string)
		r1["key1"] = "value1"
		r1["key2"] = "value2"
		return r1, errors.New("错误返回值")
	})

	ipc.OnArguments("testInArgs", func(args1 string, args2 int, args3 bool, args4 []any, args5 map[string]any, args6 src.TestInArgs, args7 map[string]src.TestInArgs,
		args8 map[string]string, args9 float64, args10 []float64) (string, int, bool) {
		fmt.Println("args1", args1)
		fmt.Println("args2", args2)
		fmt.Println("args3", args3)
		fmt.Println("args4", args4)
		fmt.Println("args5", args5)
		fmt.Println("args6", args6, args6.SubObj)
		fmt.Println("args7", args7)
		fmt.Println("args8", args8)
		fmt.Println("args9", args9)
		fmt.Println("args10", args10)
		return "result testInArgs", 10, true
	})
	var num = 0
	ipc.On("testEmitName", func(context ipc.IContext) {
		num++
		argument := context.ArgumentList()
		fmt.Println("testEmitName", argument.Size(), context.BrowserId(), context.FrameId(), num)
		fmt.Println("data:", argument.ToJSONString())
		for i := 0; i < argument.Size(); i++ {
			value := argument.GetByIndex(i)
			fmt.Println(i, "type:", value.Type(), "isInt:", value.IsInt())
		}
		var r5 = make([]string, 3, 3)
		r5[0] = "Array数组值1"
		r5[1] = "Array数组值2"
		r5[2] = "Array数组值3"
		var r6 = make([]*src.StructVarDemo, 4, 4)
		r6[0] = &src.StructVarDemo{StringField: "StringField1字符串1"}
		r6[1] = &src.StructVarDemo{StringField: "StringField2字符串2", IntField: 111, BoolField: true, FloatField: 999.99, SubStructObj: &src.SubStructObj{StringField: "子对象String值", StructVarDemo: &src.StructVarDemo{StringField: "嵌套了嵌套了"}}}
		var r7 = make([]src.StructVarDemo, 4, 4)
		r7[0] = src.StructVarDemo{StringField: "r7参数字符串1"}
		r7[1] = src.StructVarDemo{StringField: "r7参数字符串2"}
		var r8 = map[string]string{}
		r8["r8key1"] = "r8key1"
		r8["r8key2"] = "r8key2"
		var r9 = map[string]interface{}{}
		r9["r9keyr6"] = r6
		r9["r9keyr61"] = r6[1]
		r9["r9keyr7"] = r7[1]
		r9["r9keystrValue"] = "stringValue"
		r9["r9keyintValue"] = 50000
		r9["r9keyboolValue"] = true
		r9["r9keyfloatValue"] = 5555555.99999
		r9["r9keystrArrr5"] = r5
		var objMapArr = make([]map[string]interface{}, 3)
		objMapArr[0] = r9
		objMapArr[1] = r9
		objMapArr[2] = r9
		err := &MyError{error: "返回值"}
		context.Result("asdfsadf", 123123, true, err /*, r5, r6, r7, r8, r9*/)
	})

	cef.VariableBind.Bind("funcName", func(intVar int, stringVar string, doubleVar float64) (string, int, bool) {
		return "StringValue", 100000111, true
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

type MyError struct {
	error string
}

func (m *MyError) Error() string {
	return m.error
}
