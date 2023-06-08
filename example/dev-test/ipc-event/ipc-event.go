package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/i18n"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/ipc/callback"
	"github.com/energye/energy/v2/cef/ipc/context"
	"github.com/energye/energy/v2/cef/ipc/target"
	"github.com/energye/energy/v2/cef/ipc/types"
	"github.com/energye/energy/v2/cef/process"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/example/dev-test/ipc-event/src"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/energy/v2/pkgs/json"
	"github.com/energye/golcl/lcl"
	lclTypes "github.com/energye/golcl/lcl/types"
	"time"
	//_ "net/http/pprof"
)

//go:embed resources
var resources embed.FS
var cefApp *cef.TCEFApplication

func main() {
	//logger.SetEnable(true)
	//logger.SetLevel(logger.CefLog_Debug)
	//go func() {
	//	http.ListenAndServe(":10000", nil)
	//}()
	//logger.SetEnable(true)
	//logger.SetLevel(logger.CefLog_Debug)
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	cefApp = cef.NewApplication()
	fmt.Println("TotalSystemMemory", cefApp.TotalSystemMemory(), cefApp.UsedMemory(), cefApp.SystemMemoryLoad())
	fmt.Println("CEFVersion", cefApp.LibCefVersion(), cefApp.ChromeVersion(), cefApp.ApiHashUniversal())
	fmt.Println("LibVersion:", cefApp.LibVersion(), "LibBuildVersion:", cefApp.LibBuildVersion())
	i18n.SetLocalFS(&resources, "resources")
	i18n.Switch(consts.LANGUAGE_zh_CN)
	//cefApp.SetEnableGPU(true)
	//cefApp.SetLogSeverity(consts.LOGSEVERITY_DEBUG)
	//cefApp.SetSingleProcess(true)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/ipc-event.html"
	//cef.BrowserWindow.Config.Url = "https://map.baidu.com/"
	//cef.BrowserWindow.Config.Url = "https://www.csdn.net/"
	cef.BrowserWindow.Config.Title = "Energy - ipc-event"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	//cef.BrowserWindow.Config.ChromiumConfig().SetEnableWindowPopup(false)
	//cef.BrowserWindow.Config.ChromiumConfig().SetEnableDevTools(false)
	//cef.BrowserWindow.Config.ChromiumConfig().SetEnableMenu(false)
	//cef.BrowserWindow.Config.ChromiumConfig().SetEnableViewSource(false)
	//cef.BrowserWindow.Config.ChromiumConfig().SetEnabledJavascript(false)
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
	var testEmitSync = 0

	var ctime = 5
	//监听事件，js触发，之后再触发js监听的事件
	ipc.On("testGoEmit", func(context context.IContext) {
		testGoEmit++
		args := context.ArgumentList().JSONArray()
		if tm >= 59-ctime {
			tm = time.Now().Second()
		}
		if time.Now().Second() >= tm+ctime {
			fmt.Println("GetIntByIndex", args.GetIntByIndex(0), "testGoEmit:", testGoEmit, "testGoEmitAndCallback:", testGoEmitAndCallback, "testEmitName:", testEmitName, "testEmitSync:", testEmitSync, "testResultArgs:", testResultArgs, "onTestName1Emit:", onTestName1Emit)
			tm = time.Now().Second()
		}
		//触发JS监听的事件，并传入参数, 只主进程能接收
		ipc.Emit("onTestName1", r0, testGoEmit, r2, r3, r4, r5, r6, r7, r8, r9, r10)

		//给所有窗口发送消息
		//多窗口通信 CEF 不宜频率过高, 会导致内存泄露来不及释放, 一秒钟总共230次左右, 这个频率和 JS setInterval 一样
		for _, wi := range cef.BrowserWindow.GetWindowInfos() {
			//fmt.Println("wi", wi.Browser().Identifier())
			//给指定目标窗口发送消息
			//target := target.NewTarget(bsrId, wi.Browser().FrameId())
			target := wi.Browser().Target(0)
			//target := target.NewTarget(bsrId, 4444)
			ipc.EmitTarget("onTestName1", target, r0, testGoEmit, r2, r3, r4, r5, r6, r7, r8, r9, r10)
			ipc.EmitTargetAndCallback("onTestName2", target, []interface{}{r0, testGoEmit, r2, r3, r4, r5, r6, r7, r8, r9, r10}, func(channel callback.IChannel, r1 string, r2 int, r3 float64, r4 bool) {
				fmt.Println("EmitTargetAndCallback - onTestName1 callback", r1, r2, r3, r4, "frameId:", process.FrameId(), "channel:", channel)
			})
			ipc.EmitAndCallback("onTestName2", []interface{}{r0, testGoEmit, r2, r3, r4, r5, r6, r7, r8, r9, r10}, func(channel callback.IChannel, r1 string, r2 int, r3 float64, r4 bool) {
				fmt.Println("EmitAndCallback - onTestName1 callback", r1, r2, r3, r4, "frameId:", process.FrameId(), "channel:", channel)
			})
		}
	})
	ipc.On("testGoEmitAndCallback", func() {
		testGoEmitAndCallback++
		//fmt.Println("testGoEmitAndCallback")
		//触发JS监听的函数，并传入参数
		ipc.EmitAndCallback("onTestName2", []interface{}{r0, testGoEmit, r2, r3, r4, r5, r6, r7, r8, r9, r10}, func(r1 string, r2 int, r3 float64, r4 bool) {
			//fmt.Println("onTestName1 callback", r1, r2, r3, r4)
		})
	})

	ipc.On("testResultArgs", func(args1 int) (string, int, float64, bool, *MyError, []string, []*src.StructVarDemo, []src.StructVarDemo, map[string]string, map[string]interface{}, []map[string]interface{}) {
		testResultArgs++
		//fmt.Println("args1", args1)
		return r0, testResultArgs, r2, r3, r4, r5, r6, r7, r8, r9, r10
	})

	ipc.On("testInArgs", func(in1 string, in2 int, in3 float64, in4 bool, in5 []string, in6 []any, in7 map[string]any, in8 src.TestInArgs, in9 map[string]src.TestInArgs) (string, int, bool) {
		fmt.Println("testInArgs in1: ", in1)
		fmt.Println("testInArgs in2: ", in2)
		fmt.Println("testInArgs in3: ", in3)
		fmt.Println("testInArgs in4: ", in4)
		fmt.Println("testInArgs in5: ", in5)
		fmt.Println("testInArgs in6: ", in6)
		fmt.Println("testInArgs in7: ", json.NewJSONObject(in7).ToJSONString())
		fmt.Println("testInArgs in8: ", json.NewJSONObject(in8).ToJSONString())
		fmt.Println("testInArgs in9: ", json.NewJSONObject(in9).ToJSONString())
		return "result testInArgs", 10, true //JS 没有回调函数这些参数不会返回
	})
	ipc.On("testNotInArgs", func() {
		fmt.Println("无入参，无出参")
	})
	ipc.On("testEmitSync", func(channel callback.IChannel) (any, int) {
		testEmitSync++
		fmt.Println("无入参，无出参", channel)
		return r9, testEmitSync
	})

	ipc.On("testEmitName", func(context context.IContext) {
		testEmitName++
		argument := context.ArgumentList()
		//fmt.Println("testEmitName", argument.Size(), context.BrowserId(), context.FrameId(), testEmitName)
		fmt.Println("args1", argument.GetStringByIndex(0))
		fmt.Println("args2", argument.GetStringByIndex(1))
		fmt.Println("args3:", argument.GetByIndex(2).Data())
		fmt.Println("args4:", argument.GetByIndex(3).Data())
		fmt.Println("args5:", argument.GetByIndex(4).Data())
		fmt.Println("args6:", argument.GetByIndex(5).Data())
		for i := 0; i < argument.Size(); i++ {
			_ = argument.GetByIndex(i)
			//fmt.Println(i, "type:", value.Type(), "data:", value.Data())
		}
		fmt.Println("context.BrowserId():", context.BrowserId(), "context.FrameId():", context.FrameId(), "GetWindowInfos:", len(cef.BrowserWindow.GetWindowInfos()))
		if wi := cef.BrowserWindow.GetWindowInfo(context.BrowserId()); wi != nil {
			wi.Id()
		} else {
			fmt.Println("为获取到 windowInfo:", context.BrowserId())
		}
		context.Result(r0, testEmitName, r2, r3, r4, r5, r6, r7, r8, r9, r10)
	})
	ipc.On("testGoToJSEvent", func() {
		fmt.Println("testGoToJSEvent")
		ipc.Emit("notInArgs")
		ipc.Emit("notInArgs", "param1", "param2", 6666, "这些参数没接收")
		ipc.EmitAndCallback("notInArgs", []any{"param1", "param2", 6666, "这些参数没接收"}, func() {
			fmt.Println("有回调函数，但没有返回值")
		})
		ipc.EmitAndCallback("notInArgs", []any{"param1", "param2", 6666, "这些参数没接收"}, func(in1 int, in2 float64) {
			fmt.Println("有回调函数，有1个返回值 in1:", in1, "in2:", in2)
		})
	})

	ipc.On("testSendTargetGo", func(channel callback.IChannel, in0 int64, in1 string, channelId string) {
		fmt.Println("testSendTargetGo", channel, "in0:", in0, "in1:", in1, "channelId:", channelId, "当前通道ID:", process.FrameId())
		target := target.NewTarget(channel.BrowserId(), common.StrToInt64(channelId), target.TgGo)
		//cef.BrowserWindow.GetBrowser(channel.BrowserId()).FrameId()
		//ipc.EmitTarget("testOnTargetGo", target, 100001)

		ipc.EmitTargetAndCallback("testOnTargetGo", target, []any{2000002, channelId}, func(r1 string) {
			//fmt.Println("testOnTargetGo 返回值？r1:", r1, "当前通道ID:", process.FrameId())
		})

	})

	//子进程监听
	ipc.On("testOnTargetGo", func(channel callback.IChannel, in0 int, in1ChannelId string) string {
		fmt.Println("testOnTargetGo channel:", channel, " in0:", in0, "in1ChannelId:", in1ChannelId, "当前通道ID:", process.FrameId())

		return "返回了？" + in1ChannelId
	}, types.OnOptions{OnType: types.OtSub})

	ipc.On("reload", func(channel callback.IChannel) {
		if wi := cef.BrowserWindow.GetWindowInfo(channel.BrowserId()); wi != nil {
			wi.Chromium().ReloadIgnoreCache()
		}
	})

	ipc.On("check-param", func(pstr1 string, pint2 int, pint3 int, pfloat4 int, pany5 int, pstr6 string, pmap7 map[string]any, pstrarr8 []string, pintarr9 []int, pbool10 bool) {
		//fmt.Println("check-param：", pstr1, pint2, pint3, pfloat4, pany5, pstr6, pmap7, pstrarr8, pintarr9, pbool10)
		p1 := pstr1 != "参数一1验证字符串参数一1验证字符串参数一1验证字符串参数一1验证字符串"
		p2 := pint2 != 1212222
		p3 := pint3 != 44444444
		p4 := pfloat4 != 8888888
		p5 := pany5 != 666666
		p6 := pstr6 != "参数六6验证字符串参数验证字符串参数验证字符串参数验证字符串参数验证字符串参数"
		p7Key1 := pmap7["key1"] != "key1StringValue"
		p7Key2 := pmap7["key2"] != "888888"
		p7Key3 := pmap7["key3"] != "999999.88"
		p8Len := len(pstrarr8) != 5
		p9Len := len(pintarr9) != 3
		p10 := pbool10 != true
		if p1 || p2 || p3 || p4 || p5 || p6 || p7Key1 || p7Key2 || p7Key3 || p8Len || p9Len || p10 {
			fmt.Println("失败 ", "p1:", p1, "p2:", p2, "p3:", p3, "p4:", p4, "p5:", p5, "p6:", p6, "p7Key1:", p7Key1, "p7Key2:", p7Key2, "p7Key3:", p7Key3, "p8Len:", p8Len, "p9Len:", p9Len, "p10:", p10)
		}
	})

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		if window.IsLCL() && common.IsWindows() {
			window.AsLCLBrowserWindow().BrowserWindow().SetOnShow(func(sender lcl.IObject) bool {
				fmt.Println("窗口 显示/隐藏")
				return false
			})

			ipc.On("SendMouseWheelEvent", func() {
				browser := window.Chromium().Browser()
				// 滚轮
				browser.SendMouseWheelEvent(&cef.TCefMouseEvent{X: 100, Y: 100}, 200, 200)
				// 左键 按下
				browser.SendMouseClickEvent(&cef.TCefMouseEvent{X: 15, Y: 106}, 0, false, 1)
				// 左键 抬起
				browser.SendMouseClickEvent(&cef.TCefMouseEvent{X: 15, Y: 106}, 0, true, 1)
			})

			//window.Chromium().SetProxy(&cef.TCefProxy{
			//	ProxyType:              consts.PtAutodetect,
			//	ProxyScheme:            consts.PsSOCKS4,
			//	ProxyServer:            "127.0.0.1",
			//	ProxyPort:              8080,
			//	ProxyUsername:          "snsn",
			//	ProxyPassword:          "pwd",
			//	ProxyScriptURL:         "scriptURL",
			//	ProxyByPassList:        "aaa,bbb,ddd",
			//	MaxConnectionsPerProxy: 100,
			//})
			//window.Chromium().UpdatePreferences()
			event.SetOnBeforeBrowser(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, request *cef.ICefRequest, userGesture, isRedirect bool) bool {

				//settings := cef.PrintSettingsRef.New()
				//fmt.Println("settings-Instance:", settings.Instance())
				//fmt.Println("\tGetPageRangesCount:", settings.GetPageRangesCount())
				//ranges := make([]cef.TCefRange, 3, 3)
				//ranges[0] = cef.TCefRange{From: 33, To: 55}
				//ranges[1] = cef.TCefRange{From: 66, To: 77}
				//ranges[2] = cef.TCefRange{From: 88, To: 99}
				//settings.SetPageRanges(ranges)
				//fmt.Println("\tGetPageRangesCount:", settings.GetPageRangesCount())
				//pageRanges := settings.GetPageRanges()
				//fmt.Println("\tGetPageRanges:", pageRanges)

				fmt.Println("OnBeforeBrowser-Identifier:", browser.Identifier(), "userGesture:", userGesture, "isRedirect:", isRedirect)
				requestContext := browser.GetRequestContext()
				proxyDict := cef.DictionaryValueRef.New()
				proxyDict.SetString("mode", "fixed_servers")
				proxyDict.SetString("server", "192.168.1.31")
				proxy := cef.ValueRef.New()
				proxy.SetDictionary(proxyDict)
				errMsg, ok := requestContext.SetPreference("proxy", proxy)
				fmt.Println("\tproxy errMsg:", errMsg, "ok:", ok)
				//callbackRef := cef.CompletionCallbackRef.New()
				//callbackRef.OnComplete(func() {
				//	fmt.Println("callbackRef.OnComplete")
				//})
				//cookie := browser.GetRequestContext().GetCookieManager(callbackRef)
				//completionCallback := cef.CompletionCallbackRef.New()
				//completionCallback.OnComplete(func() {
				//	fmt.Println("CloseAllConnections.OnComplete")
				//})
				//browser.GetRequestContext().CloseAllConnections(completionCallback)
				//browser.GetRequestContext().LoadExtension("", nil, nil)
				//fmt.Println("cookie", cookie)
				return false
			})

			window.Chromium().SetOnGetAuthCredentials(func(sender lcl.IObject, browser *cef.ICefBrowser, originUrl string, isProxy bool, host string, port int32, realm, scheme string, callback *cef.ICefAuthCallback) bool {
				fmt.Println("onGetAuthCredentials:", originUrl, isProxy, host, port)
				return false
			})
			browserWindow := window.AsLCLBrowserWindow().BrowserWindow()
			parent := browserWindow.WindowParent()
			parent.RevertCustomAnchors() //恢复到自定义定位

			parent.SetAnchors(parent.Anchors().Include(lclTypes.AkTop, lclTypes.AkLeft, lclTypes.AkBottom, lclTypes.AkRight))
			//自定义定位
			rect := parent.BoundsRect()
			rect.Left = 50
			rect.Top = 100
			rect.SetSize(400, 600)
			//rect.Bottom = 50
			//rect.Right = 100
			parent.SetBoundsRect(rect)
			fmt.Println(parent.Width())
			//parent.SetHeight(600)
			//创建自定义系统UI组件
			var left int32 = 10
			btn := lcl.NewButton(browserWindow)
			btn.SetParent(browserWindow)
			btn.SetWidth(150)
			btn.SetCaption("HTML-窗口右移")
			btn.SetOnClick(func(sender lcl.IObject) {
				ipc.EmitAndCallback("onTestName2", []interface{}{r0, testGoEmit, r2, r3, r4, r5, r6, r7, r8, r9, r10}, func(r1 string, r2 int, r3 float64, r4 bool) {
					fmt.Println("onTestName1 callback", r1, r2, r3, r4)
				})
				rect.Left += left
				rect.SetWidth(rect.Width() + left)
				parent.SetBoundsRect(rect)
			})
			lbtn := lcl.NewButton(browserWindow)
			lbtn.SetParent(browserWindow)
			lbtn.SetTop(30)
			lbtn.SetWidth(150)
			lbtn.SetCaption("HTML-窗口左移")
			lbtn.SetOnClick(func(sender lcl.IObject) {
				ipc.EmitAndCallback("onTestName2", []interface{}{r0, testGoEmit, r2, r3, r4, r5, r6, r7, r8, r9, r10}, func(r1 string, r2 int, r3 float64, r4 bool) {
					fmt.Println("onTestName1 callback", r1, r2, r3, r4)
				})
				rect.Left -= left
				rect.SetWidth(rect.Width() - left)
				parent.SetBoundsRect(rect)
			})
			alignbtn := lcl.NewButton(browserWindow)
			alignbtn.SetParent(browserWindow)
			alignbtn.SetTop(60)
			alignbtn.SetWidth(150)
			alignbtn.SetCaption("HTML-窗口自适应/自定义")
			var align = false
			alignbtn.SetOnClick(func(sender lcl.IObject) {
				align = !align
				if align {
					parent.DefaultAnchors()
				} else {
					parent.RevertCustomAnchors() //恢复到自定义定位
					parent.SetBoundsRect(rect)
				}
				parent.UpdateSize()
			})
			browserWindow.SetOnActivate(func(sender lcl.IObject) bool {
				fmt.Println("窗口激活")
				return false
			})
			browserWindow.SetOnShow(func(sender lcl.IObject) bool {
				fmt.Println("在主窗口显示时创建第2个chromium")
				//return false
				// chromium 2
				chromium := cef.NewChromiumBrowser(browserWindow, nil)
				rect2 := chromium.WindowParent().BoundsRect()
				rect2.Left = 550
				rect2.Top = 100
				rect2.SetSize(400, 600)
				chromium.WindowParent().SetBoundsRect(rect2)
				//chromium.WindowParent().SetAnchors(chromium.WindowParent().Anchors().Include(lclTypes.AkRight))
				chromium.Chromium().SetDefaultURL("https://www.baidu.com")
				chromium.Chromium().SetOnBeforeBrowser(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, request *cef.ICefRequest, userGesture, isRedirect bool) bool {
					fmt.Println("窗口2")
					chromium.WindowParent().UpdateSize()
					return false
				})
				chromium.Chromium().SetOnGetAuthCredentials(func(sender lcl.IObject, browser *cef.ICefBrowser, originUrl string, isProxy bool, host string, port int32, realm, scheme string, callback *cef.ICefAuthCallback) bool {
					fmt.Println("onGetAuthCredentials:", originUrl, isProxy, host, port)
					return false
				})
				chromium.CreateBrowser()

				txt := lcl.NewEdit(browserWindow)
				txt.SetParent(browserWindow)
				txt.SetTop(10)
				txt.SetLeft(180)
				txt.SetWidth(200)
				txt.SetText("https://cn.bing.com/")
				goBtn := lcl.NewButton(browserWindow)
				goBtn.SetParent(browserWindow)
				goBtn.SetTop(10)
				goBtn.SetLeft(390)
				goBtn.SetCaption("GO")
				goBtn.SetOnClick(func(sender lcl.IObject) {
					chromium.Chromium().LoadUrl(txt.Text())
				})

				return false
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
		//server.SSL = &assetserve.SSL{}
		go server.StartHttpServer()
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
