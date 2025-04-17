package main

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/exception"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/cef/ipc/context"
	ipcTypes "github.com/cyber-xxm/energy/v2/cef/ipc/types"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/examples/common"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"path/filepath"
	"strings"
	"time"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	exception.SetOnException(func(message string) {
		fmt.Println("message", message)
	})
	rootCache := filepath.Join(consts.CurrentExecuteDir, "rootcache", "gotojs")
	//创建应用
	app := cef.NewApplication()
	app.SetRootCache(rootCache)
	app.SetCache(filepath.Join(rootCache, "cache"))
	app.SetUseMockKeyChain(true)
	app.SetUseMockKeyChain(true)
	port := common.Port()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = fmt.Sprintf("http://localhost:%d/go-to-js.html", port)
	cef.BrowserWindow.Config.Title = "Energy - go on event - js emit event"

	//在go中监听一个事件, 不带返回值
	//使用上下文获取参数
	ipc.On("go-on-event-demo", func(context context.IContext) {
		fmt.Println("go-on-event-demo event run")
		//js 中传递的数据
		//虽然 Arguments 结构支持多个数据类型，但在js和go的对应中，只保留了 string, integer, double, boolean 的对应关系，其它类型在 go 和 js数据传递时不支持
		arguments := context.ArgumentList()
		fmt.Println("参数个数:", arguments.Size())
		//参数是以js调用时传递的参数下标位置开始计算，从0开始表示第1个参数
		p1 := arguments.GetStringByIndex(0)
		fmt.Println("参数1:", p1)
		context.Result(time.Now().String())
	} /*, ipcTypes.OnOptions{Mode: ipcTypes.MAsync}*/)

	//带有返回值的事件
	//使用上下文获取参数
	ipc.On("go-on-event-demo-return", func(context context.IContext) {
		fmt.Println("go-on-event-demo-return event run")
		//js 中传递的数据
		//虽然 Arguments 结构支持多个数据类型，但在js和go的对应中，只保留了 string, integer, double, boolean 的对应关系，其它类型在 go 和 js数据传递时不支持
		arguments := context.ArgumentList()
		fmt.Println("参数个数:", arguments.Size())
		//参数是以js调用时传递的参数下标位置开始计算，从0开始表示第1个参数
		p1 := arguments.GetStringByIndex(0)
		p2 := arguments.GetIntByIndex(1)
		p3 := arguments.GetBoolByIndex(2)
		p4 := arguments.GetFloatByIndex(3)
		p5 := arguments.GetStringByIndex(4)
		fmt.Println("\t参数1-length:", len(p1), p1)
		//fmt.Println("\t参数1:", p1)
		fmt.Println("\t参数2:", p2)
		fmt.Println("\t参数3:", p3)
		fmt.Println("\t参数4:", p4)
		fmt.Println("\t参数5:", p5)
		//返回给JS数据, 通过 context.Result()
		var buf = bytes.Buffer{}
		for i := 0; i < 100000; i++ {
			buf.WriteString(fmt.Sprintf("[%d]-", i))
		}
		var data = "这是在GO中监听事件返回给JS的数据:" + buf.String()
		fmt.Println("返回给JS数据 - length:", strings.Count(data, "")-1)
		context.Result(data)
	})

	// 在Go中监听一个事件, 不带返回值
	// 使用形参接收参数
	// 在JS中入参类型必须相同
	ipc.On("go-on-event-demo-argument", func(param1 int, param2 string, param3 float64, param4 bool, param5 string) {
		fmt.Println("param1:", param1)
		fmt.Println("param2:", param2)
		fmt.Println("param3:", param3)
		fmt.Println("param4:", param4)
		fmt.Println("param5:", param5)
	})

	// 在Go中监听一个事件, 带返回值
	// 使用形参接收参数
	// 在JS中入参类型必须相同
	// 返回参数可以同时返回多个, 在JS接收时同样使用回调函数方式以多个入参形式接收
	// 注意：这是一个异步事件执行, 在此函数内大多数情况无法正常Debug调试它
	ipc.On("go-on-event-demo-argument-return", func(param1 int, param2 string, param3 float64, param4 bool, param5 string) string {
		fmt.Println("param1:", param1)
		fmt.Println("param2:", param2)
		fmt.Println("param3:", param3)
		fmt.Println("param4:", param4)
		fmt.Println("param5:", param5)
		return fmt.Sprintf("%d-%v-%v-%v-%v", param1, param2, param3, param4, param5)
	}, ipcTypes.OnOptions{Mode: ipcTypes.MAsync})

	ipc.On("ipc-emit-wait", func(data string) string {
		fmt.Println("data:", data)
		return fmt.Sprintf("Go Result: %v - %v", data, time.Now().String())
	})

	ipc.On("ipc-emit-wait-2000", func() string {
		fmt.Println("run task 4 second")
		time.Sleep(time.Second * 4)
		fmt.Println("run task end")
		return "这个字符串返回也没用了"
	})

	//内置http服务链接安全配置
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = port
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(app)
}
