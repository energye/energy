package main

import (
	"embed"
	"github.com/energye/energy/v2/cef"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/types"
	"path"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	cefApp := cef.NewApplication()
	wd := types.CurrentExecuteDir
	indexHtmlPath := path.Join(wd, "examples", "load-html-url", "resources", "index.html")
	println("indexHtmlPath", indexHtmlPath)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = indexHtmlPath
	cef.BrowserWindow.Config.Title = "Energy 本地加载html"
	//运行应用
	cef.Run(cefApp)
}
