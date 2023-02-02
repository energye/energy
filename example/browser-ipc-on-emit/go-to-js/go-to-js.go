package main

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/ipc"
	"strings"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	cefApp := cef.NewApplication(nil)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/go-to-js.html"
	cef.BrowserWindow.Config.Title = "Energy - go on event - js emit event"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"

	ipc.IPC.Browser().SetOnEvent(func(event ipc.IEventOn) {
		//在go中监听一个事件
		event.On("go-on-event-demo", func(context ipc.IIPCContext) {
			fmt.Println("go-on-event-demo event run")
			//js 中传递的数据
			//虽然 Arguments 结构支持多个数据类型，但在js和go的对应中，只保留了 string, integer, double, boolean 的对应关系，其它类型在 go 和 js数据传递时不支持
			arguments := context.Arguments()
			fmt.Println("参数个数:", arguments.Size())
			//参数是以js调用时传递的参数下标位置开始计算，从0开始表示第1个参数
			p1 := arguments.GetString(0)
			fmt.Println("参数1:", p1)
		})
		//带有返回值的事件
		event.On("go-on-event-demo-return", func(context ipc.IIPCContext) {
			fmt.Println("go-on-event-demo-return event run")
			//js 中传递的数据
			//虽然 Arguments 结构支持多个数据类型，但在js和go的对应中，只保留了 string, integer, double, boolean 的对应关系，其它类型在 go 和 js数据传递时不支持
			arguments := context.Arguments()
			fmt.Println("参数个数:", arguments.Size())
			//参数是以js调用时传递的参数下标位置开始计算，从0开始表示第1个参数
			p1 := arguments.GetString(0)
			p2 := arguments.GetInt32(1)
			p3 := arguments.GetBool(2)
			p4 := arguments.GetFloat64(3)
			fmt.Println("\t参数1-length:", len(p1))
			//fmt.Println("\t参数1:", p1)
			fmt.Println("\t参数2:", p2)
			fmt.Println("\t参数3:", p3)
			fmt.Println("\t参数4:", p4)
			//返回给JS数据, 通过 context.Result()
			var buf = bytes.Buffer{}
			for i := 0; i < 100000; i++ {
				buf.WriteString(fmt.Sprintf("[%d]-", i))
			}
			var data = "这是在GO中监听事件返回给JS的数据:" + buf.String()
			fmt.Println("返回给JS数据 - length:", strings.Count(data, "")-1)
			context.Result().SetString(data)
		})
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
	})
	//运行应用
	cef.Run(cefApp)
}
