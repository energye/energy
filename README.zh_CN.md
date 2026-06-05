<p align="center">
   <img src="https://energye.github.io/imgs/energy.png">
   <strong style="font-size: 24px">3.0</strong>
</p>

<p align="center" style="font-size: 24px;">
    <strong>
        是Go基于 LCL & CEF & Webview(Webview2,Webkit2) 构建跨平台桌面应用的框架
    </strong>
</p>

中文 |
[English](README.md)

---
![go-version](https://img.shields.io/github/go-mod/go-version/energye/energy?logo=git&logoColor=green)
[![github](https://img.shields.io/github/last-commit/energye/energy/main.svg?logo=github&logoColor=green&label=commit)](https://github.com/energye/energy)
[![release](https://img.shields.io/github/v/release/energye/energy?logo=git&logoColor=green)](https://github.com/energye/energy/releases)
![Build Status](https://github.com/energye/energy/actions/workflows/build-test.yml/badge.svg)
![repo](https://img.shields.io/github/repo-size/energye/energy.svg?logo=github&logoColor=green&label=repo-size)
[![Go Report](https://goreportcard.com/badge/github.com/energye/energy)](https://goreportcard.com/report/github.com/energye/energy/v2)
[![Go Reference](https://pkg.go.dev/badge/github.com/energye/energy)](https://pkg.go.dev/github.com/energye/energy/v2)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/energye/energy)
[![Discord](https://img.shields.io/discord/1303173966747271209)](https://discord.gg/TejjxKz9)
[![codecov](https://codecov.io/gh/energye/energy/graph/badge.svg?token=H370UFUF12)](https://codecov.io/gh/energye/energy)
[![contributors](https://img.shields.io/github/contributors/energye/energy)](https://github.com/energye/energy/graphs/contributors)
[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=red)](http://www.apache.org/licenses/LICENSE-2.0)
---

### [项目简介](https://energye.github.io/course/what-is-energy)

> [Energy](https://github.com/energye/energy) 是 Go 使用
> [LCL](https://www.lazarus-ide.org/)、
> [CEF](https://bitbucket.org/chromiumembedded/cef)、
> [Webview2](https://learn.microsoft.com/en-us/microsoft-edge/webview2/)、
> [Webkit2Gtk](https://webkitgtk.org/)、
> [WKWebView](https://developer.apple.com/documentation/webkit/)
> 开发的框架
>
>> LCL — 原生 UI 组件库
>>
>> CEF — 浏览器组件，CEF3 via [CEF4Delphi](https://github.com/salvadordf/CEF4Delphi)
>>
>> Webview2 — 浏览器组件，Webview2 via [WebView4Delphi](https://github.com/salvadordf/WebView4Delphi)，使用系统运行时
>>
>> Webkit2（Webkit2Gtk / WKWebView）— 浏览器组件，Webkit2，使用系统运行时


### 特性

> - **三种渲染引擎可选**：LCL 原生控件（轻量、系统风格）、Webview 系统运行时（WebView2 / WebKit2，无需额外依赖）、CEF 完整 Chromium（全量浏览器能力）。三者可混合使用或独立使用
> - **可视化设计器**：Energy Designer 支持拖拽组件、配置属性和事件、一键运行预览，生成可维护的 Go 源码
> - **上百种原生控件**：基于 LCL 提供 100+ 系统原生 GUI 组件，包括按钮、表格、树形视图、列表、网格、对话框、代码编辑器等
> - **Go + Web 混合架构**：Go 后端负责窗口管理、系统调用、文件流、安全加密等；Web 前端（HTML + CSS + JavaScript）负责界面展示，支持 Vue、React、Angular 等主流前端框架
> - **IPC 双向通信**：基于事件驱动的高性能 IPC，Go 与 Web 端可双向调用和传递数据，支持回调和自动类型转换
> - **本地资源加载**：支持自定义协议直接读取本地文件或 Go embed 嵌入资源，无需 HTTP 服务器；也支持 HTTP 服务加载
> - **跨平台**：支持 Windows（7+）、macOS（Intel / Apple Silicon）、Linux（x86 / ARM）
> - **NO CGO**：可选纯 Go 开发，无需 CGO 编译，降低环境配置门槛

### 开发环境

> - Golang >= 1.20
> - Energy 开发环境（[CEF or System Runtime], libenergy runtime）

1. 安装 [Golang](https://golang.google.cn/dl/)
2. 从 [Energy Designer](https://github.com/energye/designer) [Releases](https://github.com/energye/designer/releases) 创建项目


## 快速开始

### LCL 原生模式

```go
package main

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/lcl"
)

type TMainForm struct {
	lcl.TEngForm
}

var MainForm TMainForm

func main() {
	lcl.Init()
	lcl.RunApp(&MainForm)
}

func (m *TMainForm) FormCreate(sender lcl.IObject) {
	m.SetCaption("Hello Energy")
	m.SetWidth(400)
	m.SetHeight(300)
	m.WorkAreaCenter()

	btn := lcl.NewButton(m)
	btn.SetParent(m)
	btn.SetCaption("点击我")
	btn.SetLeft(50)
	btn.SetTop(50)
	btn.SetOnClick(func(sender lcl.IObject) {
		api.ShowMessage("Hello from Go!")
	})
}
```

### Webview 混合模式

```go
package main

import (
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/window"
	"github.com/energye/energy/v3/wv"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
)

type TForm1 struct {
	window.TWindow
	Webview1 wv.IWebview
}

var Form1 TForm1

func (m *TForm1) FormCreate(sender lcl.IObject) {
	m.TWindow.InternalBeforeFormCreate() // window create before call
	m.SetCaption("Hello Energy")
	m.SetWidth(800)
	m.SetHeight(600)

	m.Webview1 = wv.NewWebview(m)
	m.Webview1.SetAlign(types.AlClient)
	m.Webview1.SetParent(m)

	m.TWindow.FormCreate(sender) // window widget create after call

	m.Webview1.SetWindow(m)
}

func main() {
	wvApp := wv.Init()
	wvApp.SetOptions(application.Options{
		DefaultURL: "https://energye.github.io",
	})
	wv.Run(&Form1)
}
```

### 使用 Energy Designer

1. 启动 [Energy Designer](https://github.com/energye/designer)
2. 新建项目，拖拽组件到画布
3. 配置属性和事件
4. 点击运行预览

### NO CGO

> 可选纯 `Go` 开发，无需 `CGO` 编译

### [示例](https://github.com/energye/examples/tree/main)

### 文档

- [项目简介](https://energye.github.io/course/what-is-energy) — 什么是 Energy
- [Go Reference](https://pkg.go.dev/github.com/energye/energy/v2) — API 文档
- [DeepWiki](https://deepwiki.com/energye/energy) — 项目 Wiki


### 系统支持

![Windows](https://img.shields.io/badge/windows-supported-success.svg?logo=Windows&logoColor=blue)
![MacOS](https://img.shields.io/badge/MacOS-supported-success.svg?logo=MacOS)
![Linux](https://img.shields.io/badge/Linux-supported-success.svg?logo=Linux&logoColor=red)


|             | 32位                                                                                       | 64位                                                                                       | 测试系统版本                             |
|-------------|--------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------|------------------------------------|
| Windows     | ![Windows](https://img.shields.io/badge/supported-success.svg?logo=Windows&logoColor=blue) | ![Windows](https://img.shields.io/badge/supported-success.svg?logo=Windows&logoColor=blue) | 7、10、11                            |
| MacOSX      | ![MacOSX](https://img.shields.io/badge/N/A-inactive.svg?logo=MacOS)                        | ![MacOSX](https://img.shields.io/badge/supported-success.svg?logo=MacOS)                   | MacOSX Intel x86                   |
| MacOS M1/M2 | ![MacOS](https://img.shields.io/badge/N/A-inactive.svg?logo=MacOS)                         | ![MacOS](https://img.shields.io/badge/supported-success.svg?logo=MacOS)                    | MacOS Apple Silicon                |
| Linux       | ![Linux](https://img.shields.io/badge/自编译-supported-success.svg?logo=Linux)                | ![Linux](https://img.shields.io/badge/supported-success.svg?logo=Linux&logoColor=red)      | Deepin20.8、Ubuntu18.04、LinuxMint21 |
| Linux ARM   | ![Linux ARM](https://img.shields.io/badge/自编译-supported-success.svg?logo=Linux)            | ![Linux ARM](https://img.shields.io/badge/supported-success.svg?logo=Linux)                | Kylin-V10-SP1-2107                 |

### 相关项目
* [Energy Designer](https://github.com/energye/designer)
* [Energy](https://github.com/energye/energy)
* [LCL](https://github.com/energye/lcl)
* [Webview](https://github.com/energye/wv)
* [CEF](https://github.com/energye/cef)
* [WebView4Delphi](https://github.com/salvadordf/WebView4Delphi)
* [CEF](https://github.com/chromiumembedded/cef)
* [CEF4Delphi](https://github.com/salvadordf/CEF4Delphi)
* [CefSharp](https://github.com/cefsharp/CefSharp)
* [Java-CEF](https://bitbucket.org/chromiumembedded/java-cef)
* [cefpython](https://github.com/cztomczak/cefpython)
* [Chromium](https://chromium.googlesource.com/chromium/src/)

---

如果你觉得此项目对你有帮助，请点亮 Star！

---

### ENERGY QQ 交流群 & 微信

<p align="center">
    <img src="https://energye.github.io/imgs/assets/qq-group.jpg" width="250" title="QQ交流群: 541258627" alt="QQ交流群: 541258627">
    <img src="https://energye.github.io/imgs/assets/we-chat.jpg" width="250" title="微信: sniawmdf" alt="微信: sniawmdf" style="margin-left: 30px;">
</p>

---

### 项目截图

[image](https://github.com/energye/designer/tree/main/docs/image)

----

### 开源协议

[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=green)](http://www.apache.org/licenses/LICENSE-2.0)

### 贡献者
<a href="https://github.com/energye/energy/graphs/contributors">
    <img src="https://opencollective.com/energy/contributors.svg?width=890&button=false" />
</a>
