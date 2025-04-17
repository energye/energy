//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/common"
	demoCommon "github.com/cyber-xxm/energy/v2/examples/common"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"time"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, demoCommon.ResourcesFS())
	//创建应用
	app := cef.NewApplication()
	if common.IsDarwin() {
		app.SetUseMockKeyChain(true)
	}
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/webkit-register.html"
	cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	cef.BrowserWindow.Config.Title = "Energy webkit-register"

	// webkit 初始化时注册JS函数
	// 使用 RegisterExtension
	// V8Handler 接收执行时回调
	app.SetOnWebKitInitialized(func() {
		var myparamValue string
		v8Handler := cef.V8HandlerRef.New()
		v8Handler.Execute(func(name string, object *cef.ICefV8Value, arguments *cef.TCefV8ValueArray, retVal *cef.ResultV8Value, exception *cef.ResultString) bool {
			fmt.Println("v8Handler.Execute", name)
			var result bool
			if name == "GetMyParam" {
				result = true
				myparamValue = myparamValue + " " + time.Now().String()
				retVal.SetResult(cef.V8ValueRef.NewString(myparamValue))
			} else if name == "SetMyParam" {
				if arguments.Size() > 0 {
					newValue := arguments.Get(0)
					fmt.Println("value is string:", newValue.IsString())
					fmt.Println("value:", newValue.GetStringValue())
					myparamValue = newValue.GetStringValue()
					newValue.Free()
				}
				result = true
			}
			return result
		})
		//注册js
		var jsCode = `
            let test;
            if (!test) {
                test = {};
            }
            (function () {
                test.__defineGetter__('myparam', function () {
                    native function GetMyParam();
                    return GetMyParam();
                });
                test.__defineSetter__('myparam', function (b) {
                    native function SetMyParam();
					b = b + ' TEST';
                    if (b) SetMyParam(b);
                });
            })();
`
		// 注册JS 和v8处理器
		cef.RegisterExtension("v8/test", jsCode, v8Handler)
	})

	//在主进程启动成功之后执行
	//在这里启动内置http服务
	//内置http服务需要使用 go:embed resources 内置资源到执行程序中
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		server.Assets = demoCommon.ResourcesFS()
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(app)
}
