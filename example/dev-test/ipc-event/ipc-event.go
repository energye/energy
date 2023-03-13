package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/example/dev-test/ipc-event/src"
	"github.com/energye/energy/ipc"
	//_ "net/http/pprof"
)

//go:embed resources
var resources embed.FS
var cefApp *cef.TCEFApplication

func main() {
	//go func() {
	//	fmt.Println("pprof")
	//	http.ListenAndServe(":9999", nil)
	//}()
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

	ipc.OnArguments("testResultArgs", func(args1 int) (string, int, float64, bool, *MyError, []string, []*src.StructVarDemo, []src.StructVarDemo, map[string]string, map[string]interface{}, []map[string]interface{}) {
		fmt.Println("args1", args1)
		var r0 = "字符串{}{}{}字符串[][]字符串"
		var r1 = 1000011
		var r2 = 66666611.0123
		var r3 = true
		var r4 = &MyError{error: "返回值"}
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
		var r10 = make([]map[string]interface{}, 3)
		r10[0] = r9
		r10[1] = r9
		r10[2] = r9
		return r0, r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
	})

	ipc.OnArguments("testInArgs", func(in1 string, in2 int, in3 float64, in4 bool, in5 []string, in6 []any, in7 map[string]any, in8 src.TestInArgs, in9 map[string]src.TestInArgs) (string, int, bool) {
		fmt.Println("in1: ", in1)
		fmt.Println("in2: ", in2)
		fmt.Println("in3: ", in3)
		fmt.Println("in4: ", in4)
		fmt.Println("in5: ", in5)
		fmt.Println("in6: ", in6)
		fmt.Println("in7: ", in7)
		fmt.Println("in8: ", in8, "in8.SubObj: ", in8.SubObj)
		fmt.Println("in9: ", in9)
		return "result testInArgs", 10, true //没有回调函数这些参数不会返回
	})
	var num = 0
	ipc.On("testEmitName", func(context ipc.IContext) {
		num++
		argument := context.ArgumentList()
		fmt.Println("testEmitName", argument.Size(), context.BrowserId(), context.FrameId(), num)
		fmt.Println("data:", argument.GetByIndex(1).Data())
		for i := 0; i < argument.Size(); i++ {
			value := argument.GetByIndex(i)
			fmt.Println(i, "type:", value.Type(), "isInt:", value.IsInt())
		}
		var r0 = "字符串{}{}{}字符串[][]字符串"
		var r1 = 1000011
		var r2 = 66666611.0123
		var r3 = true
		var r4 = &MyError{error: "返回值"}
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
		var r10 = make([]map[string]interface{}, 3)
		r10[0] = r9
		r10[1] = r9
		r10[2] = r9
		context.Result(r0, r1, r2, r3, r4, r5, r6, r7, r8, r9, r10)
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
