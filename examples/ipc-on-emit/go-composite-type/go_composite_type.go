package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/examples/ipc-on-emit/go-composite-type/src"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"strconv"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	cefApp := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/go-composite-type.html"
	cef.BrowserWindow.Config.Title = "Energy - go-composite-type"

	// 在这个示例中演示了复合类型参数
	// 包含 struct slice map
	var userInfo = func(v string) *src.DemoUser {
		result := &src.DemoUser{
			Name:   "张三" + v,
			Age:    66,
			Income: 99988.0009,
			Sex:    true,
			UserInfo: src.DemoUserInfo{
				Phone:       "888-999-000",
				Addr:        "银河系-猎户座旋臂(离中心远,离边缘近)-太阳系第三环总体位置,离银棒(中心)",
				HeadPicture: "https://www.demo.com/head.png",
				Height:      800,
				Weight:      800,
			},
		}
		return result
	}

	// 获取用户信息-返回结构体
	ipc.On("userInfo-struct", func() (*src.DemoUser, src.DemoUser) {
		// 返回用户信息
		// 2种方式，指针和非指针方式
		result := userInfo("")
		return result, *result
	})

	// 获取用户信息-返回map
	ipc.On("userInfo-map", func() map[string]*src.DemoUser {
		result := make(map[string]*src.DemoUser)
		result["zhangsan-1"] = userInfo("1")
		result["zhangsan-2"] = userInfo("2")
		result["zhangsan-3"] = userInfo("3")
		return result
	})

	// 获取用户信息-返回切片|数组
	ipc.On("userInfo-slice", func() ([]*src.DemoUser, []src.DemoUser) {
		// 返回用户信息
		result1 := make([]*src.DemoUser, 10)
		result2 := make([]src.DemoUser, 10)
		for i := 0; i < len(result1); i++ {
			result1[i] = userInfo(strconv.Itoa(i))
			result2[i] = *userInfo(strconv.Itoa(i))
		}
		return result1, result2
	})

	// 设置用户信息
	// js中 需要传递 json 格式, 字段和数据类型要正确对应
	// map 的key必须为string类型, value 不允许为指针, 否则参数接收失败
	// 结构类型不允许为指针, 否则参数接收失败
	ipc.On("setUserInfo", func(user src.DemoUser, info src.DemoUserInfo, userMap map[string]src.DemoUser) {
		fmt.Printf("user: %+v\n", user)
		fmt.Printf("info: %+v\n", info)
		fmt.Printf("userMap: %+v\n", userMap)
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
	cef.Run(cefApp)
}
