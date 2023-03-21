package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/example/dev-test/ipc-event/src"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl"
	"time"
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
	cefApp.SetLogSeverity(consts.LOGSEVERITY_DEBUG)
	//cefApp.SetSingleProcess(true)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/ipc-event.html"
	//cef.BrowserWindow.Config.Url = "https://map.baidu.com/"
	cef.BrowserWindow.Config.Title = "Energy - ipc-event"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"

	// 测试用的入参 和 出参
	var r0 = "字符串{}{}{}字符串[][]字符串"
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
	var tm = time.Now().Second()
	var testGoEmit = 0
	var testGoEmitAndCallback = 0
	var testEmitName = 0
	var testResultArgs = 0
	var onTestName1Emit = 0

	//监听事件，js触发，之后再触发js监听的事件
	ipc.On("testGoEmit", func(context ipc.IContext) {
		testGoEmit++
		args := context.ArgumentList().JSONArray()
		if tm > 58 {
			tm = time.Now().Second()
		}
		if time.Now().Second() >= tm+1 {
			fmt.Println("GetIntByIndex", args.GetIntByIndex(0), "testGoEmit:", testGoEmit, "testGoEmitAndCallback:", testGoEmitAndCallback, "testEmitName:", testEmitName, "testResultArgs:", testResultArgs, "onTestName1Emit:", onTestName1Emit)
			tm = time.Now().Second()
		}
		//触发JS监听的事件，并传入参数
		//ipc.Emit("onTestName1", r0, testGoEmit, r2, r3, r4, r5, r6, r7, r8, r9, r10)
		ipc.EmitAndCallback("onTestName2", []interface{}{r0, testGoEmit, r2, r3, r4, r5, r6, r7, r8, r9, r10}, func(r1 string, r2 int, r3 float64, r4 bool) {
			fmt.Println("onTestName1 callback", r1, r2, r3, r4)
		})
	})
	ipc.On("testGoEmitAndCallback", func() {
		testGoEmitAndCallback++
		fmt.Println("testGoEmitAndCallback")
		//触发JS监听的函数，并传入参数
		//ipc.EmitAndCallback("onTestName2", []any{r0, r1 + count, r2, r3, r4, r5, r6, r7, r8, r9, r10}, func(r1 string) {
		//	//fmt.Println("onTestName2 r1: ", r1)
		//})
	})

	ipc.On("testResultArgs", func(args1 int) (string, int, float64, bool, *MyError, []string, []*src.StructVarDemo, []src.StructVarDemo, map[string]string, map[string]interface{}, []map[string]interface{}) {
		testResultArgs++
		return r0, testResultArgs, r2, r3, r4, r5, r6, r7, r8, r9, r10
	})

	ipc.On("testInArgs", func(in1 string, in2 int, in3 float64, in4 bool, in5 []string, in6 []any, in7 map[string]any, in8 src.TestInArgs, in9 map[string]src.TestInArgs) (string, int, bool) {
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
	ipc.On("testEmitName", func(context ipc.IContext) {
		testEmitName++
		argument := context.ArgumentList()
		fmt.Println("testEmitName", argument.Size(), context.BrowserId(), context.FrameId(), testEmitName)
		//fmt.Println("data:", argument.GetByIndex(1).Data())
		for i := 0; i < argument.Size(); i++ {
			argument.GetByIndex(i)
			//fmt.Println(i, "type:", value.Type(), "isInt:", value.IsInt())
		}

		context.Result(r0, testEmitName, r2, r3, r4, r5, r6, r7, r8, r9, r10)
	})

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		if window.IsLCL() {
			window.AsLCLBrowserWindow().BrowserWindow().SetOnConstrainedResize(func(sender lcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
				//Browser是在chromium加载完之后创建, 窗口创建时该对象还不存在
				if window.AsLCLBrowserWindow().Browser() != nil {
					fmt.Println("SetOnConstrainedResize Identifier", window.AsLCLBrowserWindow().Browser().Identifier())
					fmt.Println("SetOnConstrainedResize MainFrame", window.AsLCLBrowserWindow().Browser().MainFrame())
				}
			})
		}
	})
	//内置http服务链接安全配置
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = &resources
		go server.StartHttpServer()
		//go func() {
		//	var i = 0
		//	time.Sleep(time.Second * 5)
		//	fmt.Println("5000")
		//	for true {
		//		i++
		//		time.Sleep(time.Second / 1000)
		//		//ipc.Emit("onTestName1", r0, i, r2, r3, r4, r5, r6, r7, r8, r9, r10)
		//		ipc.EmitAndCallback("onTestName2", []any{r0, i, r2, r3, r4, r5, r6, r7, r8, r9, r10}, func(r1 string) {
		//			fmt.Println("onTestName2 r1: ", r1)
		//		})
		//	}
		//}()
	})
	//运行应用
	cef.Run(cefApp)
}

type MyError struct {
	error string
}

func (m *MyError) Error() string {
	return m.error
}
