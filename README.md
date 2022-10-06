# Energy 是Go基于CEF的构建桌面应用的框架
[中文 ](https://github.com/energye/energy/blob/main/README.md)
[English](https://github.com/energye/energy/blob/main/README.en-US.md)

---
### 简介
> [Energy](https://github.com/energye/energy) 是使用JavaScript,HTML,CSS 构建桌面应用的框架, 是基于 [Golcl](https://github.com/energye/golcl) 和 [CEF4Delphi](https://patreon.com/salvadordf) 纯Go语言开发的框架,内嵌 [Chromium CEF](https://bitbucket.org/chromiumembedded/cef) 二进制
>
> 可以让你使用web前端开发技术构建在Windows, Linux, MacOS跨平台的应用
>
> 也可以使用 [Energy](https://github.com/energye/energy) 和 [Golcl](https://github.com/energye/golcl) 构建使用系统级本地UI和使用Web前端技术跨平台的应用
>
> 支持 Windows_32、64 bits, Linux_x86_64 bits, MacOS_x86_64 bits

### CEF(105.3.39)二进制下载
* [Windows 32 bits](https://cef-builds.spotifycdn.com/cef_binary_105.3.39%2Bg2ec21f9%2Bchromium-105.0.5195.127_windows32.tar.bz2)
* [Windows 64 bits](https://cef-builds.spotifycdn.com/cef_binary_105.3.39%2Bg2ec21f9%2Bchromium-105.0.5195.127_windows64.tar.bz2)
* [Linux x86 64 bits](https://cef-builds.spotifycdn.com/cef_binary_105.3.39%2Bg2ec21f9%2Bchromium-105.0.5195.127_linux64.tar.bz2)
* [MacOS x86 64 bits](https://cef-builds.spotifycdn.com/cef_binary_105.3.39%2Bg2ec21f9%2Bchromium-105.0.5195.127_macosx64.tar.bz2)

### 入门指南
* [教程]()
* [示例]()
* [文档]()

### 快速入门
#### 基本需求
> golang >= 1.9.2
>
> 下载对应平台和版本的CEF二进制包,解压到目录.
>
> 以example/simple示例为例
>
> 安装energy依赖 go get github.com/energye/energy
>
> 或者使用 go mod init, go mod tidy
>
> 运行simple
>
> 打包应用程序 Energy 没有打包模块, windows你可以使用(MSI或Inno Setup)和其它绿色打包工具, linux下deb安装包等, MacOS默认开发时会生成.app包或者自行定制.app包

##### example/simple 示例代码
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
	//通过创建窗口之后对对主窗口的属性、组件或子创建的创建
	cef.BrowserWindow.SetBrowserInitAfter(func(browserWindow *cef.TCefWindowInfo) {
		fmt.Println("SetBrowserInitAfter")
	})
	//运行应用
	cef.Run(cefApp)
}
```
### 安装
* 下载对应平台和对应energy版本的二进制压缩包
* 安装energy依赖 go get github.com/energye/energy 或者使用 go mod init, go mod tidy
* 创建GO程序应用, 参考入门指南和example示例


### 授权

**保持跟golcl采用相同的授权协议: [golcl.LICENSE](https://github.com/energye/golcl/blob/main/LICENSE)**  