<p align="center">
   <img src="https://assets.yanghy.cn/energy-icon.png">
   <br>
   <strong style="font-size: 24px">3.0 dev</strong>
</p>

<p align="center" style="font-size: 24px;">
    <strong>
        是Go基于 LCL & CEF & Webview2 构建桌面应用的框架
    </strong>
</p>

中文 |
[English](README.md)

---
![go-version](https://img.shields.io/github/go-mod/go-version/energye/energy?logo=git&logoColor=green)
[![github](https://img.shields.io/github/last-commit/energye/energy/main.svg?logo=github&logoColor=green&label=commit)](https://github.com/energye/energy)
[![release](https://img.shields.io/github/v/release/energye/energy?logo=git&logoColor=green)](https://github.com/energye/energy/releases)
![Build Status](https://github.com/energye/energy/actions/workflows/go.yml/badge.svg)
![repo](https://img.shields.io/github/repo-size/energye/energy.svg?logo=github&logoColor=green&label=repo-size)
[![Go Report](https://goreportcard.com/badge/github.com/energye/energy)](https://goreportcard.com/report/github.com/energye/energy/v2)
[![Go Reference](https://pkg.go.dev/badge/github.com/energye/energy)](https://pkg.go.dev/github.com/energye/energy/v2)
[![codecov](https://codecov.io/gh/energye/energy/graph/badge.svg?token=H370UFUF12)](https://codecov.io/gh/energye/energy)
[![contributors](https://img.shields.io/github/contributors/energye/energy)](https://github.com/energye/energy/graphs/contributors)
[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=red)](http://www.apache.org/licenses/LICENSE-2.0)
---

### [项目简介](https://energy.yanghy.cn/course/100/6350f94ca749ba0318943f25)

> [energy](https://github.com/energye/energy) 
> 是 Go 基于
> [LCL](https://www.lazarus-ide.org/)、
> [CEF](https://bitbucket.org/chromiumembedded/cef)、
> [Webview2](https://learn.microsoft.com/en-us/microsoft-edge/webview2/)、
> [Webkit Gtk](https://webkitgtk.org/)、
> [Webkit Apple](https://developer.apple.com/documentation/webkit/)
> 开发的框架
>
> 
>> LCL - 基础库, 图形用户界面(GUI)组件库, 提供了非常丰富的系统原生控件
>>
>> CEF - 浏览器组件库 [CEF4Delphi](https://github.com/salvadordf/CEF4Delphi), 在LCL基础上封装的CEF3库
>> 
>> Webview2 - 浏览器组件库 [WebView4Delphi](https://github.com/salvadordf/WebView4Delphi), 在LCL基础上封装的Webview2库
>> 
>> Webkit - 浏览器组件库, webkit webview, 在LCL基础上封装的Webkit库
>> 
>> 使用 Go 和 Web 端技术 ( HTML + CSS + JavaScript ) 构建支持Windows, Linux, MacOS跨平台桌面应用。
>> 将web内容无缝集成到应用程序中，并自定义内容交互以满足应用程序的需求。
> 
> 构建&使用
> 
>> LCL 单独使用, 开发原生图形用户界面(GUI)应用. 轻量级, 丰富的系统原生控件
> 
>> LCL + CEF 混合使用, 开发原生图形用户界面(GUI)和浏览器应用. 重量级, 全量chromium API
>
>> LCL + Webview2 混合使用, 开发原生图形用户界面(GUI)和浏览器应用. 轻量级, 全量webview2 API
>
>> LCL + Webkit 混合使用, 开发原生图形用户界面(GUI)和浏览器应用. 轻量级, 基础webkit API



### 特性

> - 具有丰富的浏览器 API 和 LCL 系统原生小部件
> - 开发环境简单,编译速度快,仅需Go和Energy依赖的CEF二进制框架
> - 跨平台: 一套代码可以打包成 Windows, 国产UOS、Deepin、Kylin, MacOS, Linux
>> - Golang: 窗口管理、CEF API封装&配置、功能实现、各种UI组件创建、系统低层调用和JavaScript处理不了的功能，如: 文件流、安全加密、高性能处理等
>> - Web: HTML + CSS + JavaScript 负责客户端界面的功能, 做出任意你想要的界面
> - 前端技术: 支持主流前端框架。例如：Vue、React、Angular 和 原生HTML+CSS+JS等
> - 事件驱动: 高性能事件驱动, 基于IPC通信，实现Go和Web端迅捷调用及数据交互
> - 资源加载: 可无需http服务支撑，直接读取本地资源或内置到执行文件的资源, 也支持http服务加载资源

### 内置依赖&集成

- [![LCL](https://img.shields.io/badge/LCL-green)](https://github.com/energye/golcl)
- [![CEF-CEF4Delphi](https://img.shields.io/badge/CEF(Chromium%20Embedded%20Framework)%20CEF4Delphi-green)](https://github.com/salvadordf/CEF4Delphi)

### [开发环境](https://energy.yanghy.cn/course/100/63511b14a749ba0318943f3a)

#### 基本需求

> - Golang >= 1.16
> - Energy 开发环境(CEF, liblcl)

#### 环境安装

- 自动安装开发环境

> 使用 energy 命令行工具自动安装完整开发环境 [命令行工具下载地址](https://energy.yanghy.cn/course/100/1694511322285207)

### 入门指南 - [传送门](https://energy.yanghy.cn)

* [教程](https://energy.yanghy.cn/course/100/0)
* [示例](https://energy.yanghy.cn/example/200/0)
* [文档](https://energy.yanghy.cn/document/300/0)

### 快速入门

> 使用 [energy](https://energy.yanghy.cn/course/100/1694511322285207) 命令行工具自动安装完整开发环境

### 三个步骤运行一个简单应用

1. 安装开发环境: `energy install .`
2. 初始化应用: `energy init .`
3. 运行应用: `go run main.go`

### 示例代码

main.go

```go
package main

import (
	"github.com/energye/energy/v2/cef"
)

func main() {
	//全局初始化
	cef.GlobalInit(nil, nil)
	//创建应用
	app := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	//运行应用
	cef.Run(app)
}
```

### 应用打包
1. 编译：`energy build .`
2. 打包：`energy package .`
3. 自动生成的安装包
   - windows  使用`nsis`工具生成exe安装包
   - linux    使用`dpkg`制作deb安装包
   - macos    生成`xxx.app`

### 系统支持

![Windows](https://img.shields.io/badge/windows-supported-success.svg?logo=Windows&logoColor=blue)
![MacOS](https://img.shields.io/badge/MacOS-supported-success.svg?logo=MacOS)
![Linux](https://img.shields.io/badge/Linux-supported-success.svg?logo=Linux&logoColor=red)


|             | 32位                                                                                        | 64位                                                                                        | 测试系统版本                             |
|-------------|--------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------|------------------------------------|
| Windows     | ![Windows](https://img.shields.io/badge/supported-success.svg?logo=Windows&logoColor=blue) | ![Windows](https://img.shields.io/badge/supported-success.svg?logo=Windows&logoColor=blue) | Windows XP SP3、 7、10、11            |
| MacOSX      | ![MacOSX](https://img.shields.io/badge/N/A-inactive.svg?logo=MacOS)                        | ![MacOSX](https://img.shields.io/badge/supported-success.svg?logo=MacOS)                   | MacOSX 10.15                       |
| MacOS M1 M2 | ![MacOS](https://img.shields.io/badge/N/A-inactive.svg?logo=MacOS)                         | ![MacOS](https://img.shields.io/badge/supported-success.svg?logo=MacOS)                    | MacOS M2                           |
| Linux       | ![Linux](https://img.shields.io/badge/自编译-supported-success.svg?logo=Linux)                | ![Linux](https://img.shields.io/badge/supported-success.svg?logo=Linux&logoColor=red)      | Deepin20.8、Ubuntu18.04、LinuxMint21 |
| Linux ARM   | ![Linux ARM](https://img.shields.io/badge/自编译-supported-success.svg?logo=Linux)            | ![Linux ARM](https://img.shields.io/badge/supported-success.svg?logo=Linux)                | Kylin-V10-SP1-2107                 |

### 相关项目
* [CEF](https://github.com/chromiumembedded/cef)
* [CEF4Delphi](https://github.com/salvadordf/CEF4Delphi)
* [CefSharp](https://github.com/cefsharp/CefSharp)
* [Java-CEF](https://bitbucket.org/chromiumembedded/java-cef)
* [cefpython](https://github.com/cztomczak/cefpython)
* [Chromium](https://chromium.googlesource.com/chromium/src/)

---

### 欢迎加入
energy扔处于建设的过程中，有很多的事情无法独自完成，如果有感兴趣的同学想参与energy的实现或学习，可通过微信或QQ联系我。

如果你觉得此项目对你有帮助，请点亮 Star

---

### ENERGY QQ交流群 & 微信

<p align="center">
    <img src="https://assets.yanghy.cn/qq-group.jpg" width="250" title="QQ交流群: 541258627" alt="QQ交流群: 541258627">
    <img src="https://assets.yanghy.cn/we-chat.jpg" width="250" title="微信: sniawmdf" alt="微信: sniawmdf" style="margin-left: 30px;">
</p>

---

### 鸣谢 Jetbrains

<p align="center">
    <a href="https://www.jetbrains.com?from=energy">
        <img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg" alt="JetBrains Logo (Main) logo.">
    </a>
</p>

---

### 项目截图
##### Windows-10
<img src="https://assets.yanghy.cn/frameless-windows-10.png">

##### Windows-7 32 & 64
<img src="https://assets.yanghy.cn/frameless-windows-7-64.png">
<img src="https://assets.yanghy.cn/frameless-windows-7-32.png">

##### Windows-XP-SP3

<img src="https://assets.yanghy.cn/windows-xp-sp3.png">

##### Linux - Deepin
<img src="https://assets.yanghy.cn/frameless-deepin-20.8.png">
<img src="https://assets.yanghy.cn/frameless-deepin-hide-20.8.png">

##### Linux - Kylin ARM
<img src="https://assets.yanghy.cn/frameless-kylin-arm-V10-SP1.png">
<img src="https://assets.yanghy.cn/frameless-kylin-arm-hide-V10-SP1.png">

##### Linux - Ubuntu
<img src="https://assets.yanghy.cn/frameless-ubuntu-18.04.6.png">
<img src="https://assets.yanghy.cn/frameless-ubuntu-hide-18.04.6.png">

##### MacOSX
<img src="https://assets.yanghy.cn/frameless-macos.png">


----

### 开源协议

[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=green)](http://www.apache.org/licenses/LICENSE-2.0)

### 贡献者
<a href="https://github.com/energye/energy/graphs/contributors">
    <img src="https://opencollective.com/energy/contributors.svg?width=890&button=false" />
</a>