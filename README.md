# Energy 是Go基于CEF的构建桌面应用的框架
中文 |
[English](https://github.com/energye/energy/blob/main/README.en-US.md)

---
### [简介](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/6350f94ca749ba0318943f25)
> [Energy](https://github.com/energye/energy) 使用JavaScript、HTML和CSS 构建桌面应用的框架, 是Golang基于 [CEF]() 和 [Golcl](https://github.com/energye/golcl) 开发的框架，内嵌 [Chromium CEF](https://bitbucket.org/chromiumembedded/cef) 二进制
>
> 可以让你使用 [Web]() 前端技术构建在Windows, Linux, MacOS跨平台的应用
>
>> 也可以使用 [Energy](https://github.com/energye/energy) 和 [Golcl](https://github.com/energye/golcl) 单独构建小巧的系统UI跨平台应用程序
>>
>
> 支持 Windows_32、64 bits, Linux_x86_64 bits, MacOS_x86_64 bits
>
> 在Go和Web端技术中，基于IPC通信，你可以很方便的在Go和Web端交互数据,函数调用以及事件调用, 不需要Web Service接口，就像调用语言本身函数一样简单
> 
> 在Go中还可以定义JS变量，提供给Web端JS使用，实现Go变量或结构对象数据同步
> 
> 在Go中调用JS函数、JS事件监听，实现GO与JS的功能交互
> 
> 在JS中调用Go函数，Go的监听事件，实现JS与GO的功能交互

### [![windows 32 bits](https://img.shields.io/badge/Downloads-green)](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/6364c5c2a749ba01d04ff485) 
| CEF(105.3.39)  | Energy(1.1.0)                                                                                                                                                                                |
|-----------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Windows 32 bits](https://cef-builds.spotifycdn.com/cef_binary_105.3.39%2Bg2ec21f9%2Bchromium-105.0.5195.127_windows32.tar.bz2)   | [![windows 32 bits](https://img.shields.io/badge/downloads-Windows%2032%20bits-brightgreen)](https://gitee.com/energye/energy/releases/download/v1.1.0/Windows%2032%20bits.zip)              |
| [Windows 64 bits](https://cef-builds.spotifycdn.com/cef_binary_105.3.39%2Bg2ec21f9%2Bchromium-105.0.5195.127_windows64.tar.bz2)   | [![windows 64 bits](https://img.shields.io/badge/downloads-Windows%2064%20bits-brightgreen)](https://gitee.com/energye/energy/releases/download/v1.1.0/Windows%2064%20bits.zip)              |
| [Linux x86 64 bits](https://cef-builds.spotifycdn.com/cef_binary_105.3.39%2Bg2ec21f9%2Bchromium-105.0.5195.127_linux64.tar.bz2)   | [![linux x86 64 bits](https://img.shields.io/badge/downloads-Linux%20x86%2064%20bits-brightgreen)](https://gitee.com/energye/energy/releases/download/v1.1.0/Linux%20x86%2064%20bits.zip)    |
| [MacOSX x86 64 bits](https://cef-builds.spotifycdn.com/cef_binary_105.3.39%2Bg2ec21f9%2Bchromium-105.0.5195.127_macosx64.tar.bz2) | [![macOSX x86 64 bits](https://img.shields.io/badge/downloads-MacOSX%20x86%2064%20bits-brightgreen)](https://gitee.com/energye/energy/releases/download/v1.1.0/MacOSX%20x86%2064%20bits.zip) |

#### [CEF和Energy框架压缩包使用说明](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/63511b14a749ba0318943f3a)

### [入门指南-网址](https://energy.yanghy.cn)
* [教程](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/6350f94ca749ba0318943f25)
* [示例](https://energy.yanghy.cn/#/example/6342d986401bfe4d0cdf6067/634d3bd5a749ba0318943eb6)
* [文档](https://energy.yanghy.cn/#/document/6342d9a4401bfe4d0cdf6069/0)

### 快速入门
#### 基本需求
> golang >= 1.9.2
>
> 下载对应平台的CEF和Energy的动态链接库.并将其解压至任意目录.
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
	cfg.SetFrameworkDirPath("/xxxx/xxxx/ChromiumDemo")
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
	//创建窗口时的回调函数 对浏览器事件设置，和窗口属性组件等创建和修改
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
### 安装
* 下载对应平台CEF和Energy的二进制压缩包
* 安装energy依赖 go get github.com/energye/energy 或者使用 go mod init, go mod tidy
* 创建GO程序应用, 参考入门指南和example示例

----
### 授权
### [![License GPL 3.0](https://img.shields.io/badge/License%20GPL3.0-green)](https://opensource.org/licenses/GPL-3.0)
