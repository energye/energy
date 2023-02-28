package main

import (
	"github.com/energye/energy/cef"
	"os"
	"path"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	cefApp := cef.NewApplication()
	wd, _ := os.Getwd()
	indexHtmlPath := path.Join(wd, "example", "browser-load-html-url", "resources", "index.html")
	println("indexHtmlPath", indexHtmlPath)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = indexHtmlPath
	cef.BrowserWindow.Config.Title = "Energy 本地加载html"
	//运行应用
	cef.Run(cefApp)
}
