package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/cef/ipc/callback"
	"github.com/cyber-xxm/energy/v2/cef/ipc/types"
	"github.com/cyber-xxm/energy/v2/cef/process"
	"github.com/cyber-xxm/energy/v2/consts"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	//_ "net/http/pprof"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	var app = cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/dom-visitor.html"
	cef.BrowserWindow.Config.Title = "Energy - dom-visitor"

	// DomVisitor 必须在渲染进程中执行
	// 示例
	//	1. JS 通过 ipc 消息发送到主进程
	//  2. 在主进程中发送渲染进程消息
	//  3. 渲染进程中使用 DomVisitor
	// CEF DomVisitor 不如JS提供的功能函数全面
	ipc.On("dom-visitor", func(channel callback.IChannel) {
		// 此时在主进程中,我们将消息通过 frame.SendProcessMessage 发送到渲染进程
		wi := cef.BrowserWindow.GetWindowInfo(channel.BrowserId())
		fmt.Println("wi", wi, wi.Browser().MainFrame().Identifier())
		// 发送渲染进程消息
		// data 不能为空, 且大于 0 个
		wi.Chromium().SendProcessMessageForJSONBytes("dom-visitor-test", consts.PID_RENDER, []byte("test data"))
	})

	// 仅渲染(子)进程监听事件
	// 示例
	//	1. JS 通过 ipc 配置规则 发送消息到当前子进程
	//  2. 当前子进程接收到消息后创建 DomVisitor
	//  3. 渲染进程中使用 DomVisitor
	//  4. 在渲染进程的IPC消息事件中可以通过 V8ContextRef.Current() 获得 Browser 和 Frame 对象
	// CEF DomVisitor 不如JS提供的功能函数全面
	ipc.On("render-dom-visitor", func(channel callback.IChannel, args1 string) string {
		fmt.Println("render-dom-visitor", "channelId", channel.ChannelId(), "current-Id:", cef.V8ContextRef.Current().Frame().Identifier(), "args:", args1)
		// 创建 dom visitor
		visitor := cef.DomVisitorRef.New()
		// 监听事件
		// 这个事件在渲染进程中才会执行
		visitor.SetOnVisit(func(document *cef.ICefDomDocument) {
			fmt.Println("title:", document.GetTitle())
			body := document.GetBody()
			fmt.Println("body-InnerText:", body.GetElementInnerText())
			fmt.Println("GetNodeType:", body.GetNodeType())
			fmt.Println("button-domVisitor:", body.GetDocument().GetElementById("domVisitor").GetElementInnerText())
			fmt.Println("button-attrs:", body.GetDocument().GetElementById("domVisitor").GetElementAttributes())
			body.GetDocument().GetElementById("domVisitor").SetElementAttribute("id", "modify") // 把button按钮的id属性值更改了
		})
		fmt.Println("visitor:", visitor)
		// 只能在渲染进程中的IPC消息事件中使用 V8ContextRef.Current()
		cef.V8ContextRef.Current().Frame().VisitDom(visitor)
		fmt.Println("visitor-browserId:", cef.V8ContextRef.Current().Browser().Identifier())
		fmt.Println("visitor-frameId:", cef.V8ContextRef.Current().Frame().Identifier())
		fmt.Println("visitor-frameId:", process.FrameId())

		return process.FrameId()
		// OtSub 仅子进程监听该事件
	}, types.OnOptions{OnType: types.OtSub})

	// 监听渲染进程消息-在这里获取dom元素
	app.SetOnProcessMessageReceived(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, sourceProcess consts.CefProcessId, message *cef.ICefProcessMessage) bool {
		if message.Name() == "dom-visitor-test" {
			// 读取 data 数据 也可以不读取
			buf := message.ArgumentList().GetBinary(0) // 因为传递的是字节数组
			data := make([]byte, buf.GetSize())        // data 缓存区
			buf.GetData(data, 0)                       //读取数据
			buf.Free()                                 // 读取完释放掉

			// 创建 dom visitor
			visitor := cef.DomVisitorRef.New()
			// 监听事件
			// 这个事件在渲染进程中才会执行
			visitor.SetOnVisit(func(document *cef.ICefDomDocument) {
				fmt.Println("title:", document.GetTitle())
				body := document.GetBody()
				fmt.Println("body-InnerText:", body.GetElementInnerText())
				fmt.Println("GetNodeType:", body.GetNodeType())
				fmt.Println("button-domVisitor:", body.GetDocument().GetElementById("domVisitor").GetElementInnerText())
				fmt.Println("button-attrs:", body.GetDocument().GetElementById("domVisitor").GetElementAttributes())
				body.GetDocument().GetElementById("domVisitor").SetElementAttribute("id", "modify")
			})
			fmt.Println("visitor:", visitor)
			frame.VisitDom(visitor)
			return true // 接收消息已处理
		}
		return false // 接收不需要处理的消息 返回 false
	})
	//内置http服务链接安全配置
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(app)
}
