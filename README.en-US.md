# Energy is Go's framework for building desktop applications based on CEF
[中文 ](https://github.com/energye/energy/blob/main/README.md) |
English

---
### Introduction
> [Energy](https://github.com/energye/energy) Is a framework for building desktop applications using JavaScript,HTML, and CSS based on [Golcl](https://github.com/energye/golcl) and [CEF4Delphi](https://patreon.com/salvadordf) Pure Go language development framework, embedded [Chromium CEF](https://bitbucket.org/chromiumembedded/cef) binary
>
> Lets you use Web front-end development techniques to build cross-platform applications on Windows, Linux, and MacOS
>
> You can also use the [Energy](https://github.com/energye/golcl) and [Golcl](https://github.com/energye/energy) Build applications that use system-level native UI and cross platform with Web front-end technology
>
> Supports Windows_32, 64 bits, Linux_x86_64 bits, MacOS_x86_64 bits

### CEF(105.3.39)Binary download
* [Windows 32 bits](https://cef-builds.spotifycdn.com/cef_binary_105.3.39%2Bg2ec21f9%2Bchromium-105.0.5195.127_windows32.tar.bz2)
* [Windows 64 bits](https://cef-builds.spotifycdn.com/cef_binary_105.3.39%2Bg2ec21f9%2Bchromium-105.0.5195.127_windows64.tar.bz2)
* [Linux x86 64 bits](https://cef-builds.spotifycdn.com/cef_binary_105.3.39%2Bg2ec21f9%2Bchromium-105.0.5195.127_linux64.tar.bz2)
* [MacOS x86 64 bits](https://cef-builds.spotifycdn.com/cef_binary_105.3.39%2Bg2ec21f9%2Bchromium-105.0.5195.127_macosx64.tar.bz2)

### Getting started guide
* [tutorial]()
* [sample]()
* [document]()

### Quick start
#### Basic needs
> golang >= 1.9.2
>
> Download the CEF binary package for the corresponding platform and version, and decompress it to the directory.
>
> Example/simple example
>
> Install Energy dependencies go get github.com/energye/energy
>
> Or use: go mod init, go mod tidy
>
> run simple
>
> The packaging application Energy does not have packaging modules, Windows you can use (MSI or Inno Setup) and other green packaging tools, Deb installation package in Linux, MacOS generates. App packages by default or custom. App packages

##### example/simple code
```go
package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/golcl/lcl"
)

//go:embed resources
var resources embed.FS

//这是一个简单的窗口创建示例
func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalCEFInit(nil, &resources)
	//可选的应用配置
	cfg := cef.NewApplicationConfig()
	//指定chromium的二进制包框架根目录, 不指定为当前程序执行目录
	cfg.SetFrameworkDirPath("/xxxx/xxxx/chromium")
	//创建应用
	cefApp := cef.NewApplication(cfg)
	//主窗口的配置
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.DefaultUrl = "https://energy.yanghy.cn"
	//窗口的标题
	cef.BrowserWindow.Config.Title = "energy - 这是一个简单的窗口示例"
	//窗口宽高
	cef.BrowserWindow.Config.Width = 1024
	cef.BrowserWindow.Config.Height = 768
	//chromium配置
	cef.BrowserWindow.Config.SetChromiumConfig(cef.NewChromiumConfig())
	//通过创建窗口时的回调函数 对浏览器事件设置，和窗口属性组件等创建和修改
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, browserWindow *cef.TCefWindowInfo) {
		//设置应用图标 这里加载的图标是内置到执行程序里的资源文件
		lcl.Application.Icon().LoadFromFSFile("resources/icon.ico")
		fmt.Println("SetBrowserInit")
	})
	//创建窗口之后对对主窗口的属性、组件或子窗口的创建
	cef.BrowserWindow.SetBrowserInitAfter(func(browserWindow *cef.TCefWindowInfo) {
		fmt.Println("SetBrowserInitAfter")
	})
	//运行应用
	cef.Run(cefApp)
}
```
### Installation
* Download the binary package for the corresponding platform and energy version
* Install Energy dependencies go get github.com/energye/energy Or use: go mod init, go mod tidy
* Create a GO application, see the Getting Started Guide and Example
