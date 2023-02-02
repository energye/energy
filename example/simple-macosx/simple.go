package main

import (
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common"
	"github.com/energye/golcl/pkgs/macapp"
)

func main() {
	//开发环境中 MacOSX平台必须在"GlobalInit"之前设置CEF
	macapp.MacApp.IsCEF(common.IsDarwin())
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	cefApp := cef.NewApplication(nil)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	//运行应用
	cef.Run(cefApp)
}
