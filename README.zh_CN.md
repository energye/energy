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

> [energy](https://github.com/energye/energy) 
> 是 Go 使用
 [LCL](https://www.lazarus-ide.org/),
 [CEF](https://bitbucket.org/chromiumembedded/cef),
 [Webview2](https://learn.microsoft.com/en-us/microsoft-edge/webview2/)
 [Webkit2Gtk](https://webkitgtk.org/)
 [WKWebView](https://developer.apple.com/documentation/webkit/)
 开发的框架
>
> 
>> LCL - 原生UI组件库
>>
>> CEF - 浏览器组件 CEF3-[CEF4Delphi](https://github.com/salvadordf/CEF4Delphi)
>> 
>> Webview2 - 浏览器组件 Webview2-[WebView4Delphi](https://github.com/salvadordf/WebView4Delphi) 系统运行时框架
>> 
>> Webkit2(Webkit2Gtk/WKWebView) - 浏览器组件 Webkit2 系统运行时框架


### 特性

> - 原生 和 Web 可混合开发或独立使用
> - 基于框架自举 Energy Designer GUI 设计器
> - 上百种原生控件, 具有丰富的 CEF 框架 API, 轻量级Webview系统运行时框架
> - 开发环境简单, 仅需Go和Energy所需的渲染运行时框架
> - Go后端: 窗口管理、CEF API封装&配置、功能实现、各种UI组件创建、系统低层调用和JavaScript处理不了的功能，如: 文件流、安全加密、高性能处理等
> - Web前端: HTML + CSS + JavaScript 负责客户端界面的功能, 做出任意你想要的界面
> - 前端技术: 支持主流前端框架
> - 事件驱动: 高性能事件驱动, 基于IPC通信，实现Go和Web端迅捷调用及数据交互
> - 资源加载: 可无需http服务支撑，直接读取本地资源或内置到执行文件的资源, 也支持http服务加载资源

### 开发环境

> - Golang >= 1.20
> - Energy 开发环境([CEF or System Runtime], libenergy runtime)

1. 安装 [Golang](https://golang.google.cn/dl/)
2. 从 [Energy Designer](https://github.com/energye/designer) [Releases](https://github.com/energye/designer/releases) 创建项目


## 1 分钟创建一个应用

1. 启动 ENERGY Designer
2. 新建项目
3. 拖拽组件到画布
4. 配置属性和事件
5. 点击运行预览

### NO CGO

> 可选纯 `Go` 开发, 无需 `CGO` 编译

### [示例](https://github.com/energye/examples/tree/main)


### 系统支持

![Windows](https://img.shields.io/badge/windows-supported-success.svg?logo=Windows&logoColor=blue)
![MacOS](https://img.shields.io/badge/MacOS-supported-success.svg?logo=MacOS)
![Linux](https://img.shields.io/badge/Linux-supported-success.svg?logo=Linux&logoColor=red)


|             | 32位                                                                                        | 64位                                                                                        | 测试系统版本                             |
|-------------|--------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------|------------------------------------|
| Windows     | ![Windows](https://img.shields.io/badge/supported-success.svg?logo=Windows&logoColor=blue) | ![Windows](https://img.shields.io/badge/supported-success.svg?logo=Windows&logoColor=blue) | 7、10、11                            |
| MacOSX      | ![MacOSX](https://img.shields.io/badge/N/A-inactive.svg?logo=MacOS)                        | ![MacOSX](https://img.shields.io/badge/supported-success.svg?logo=MacOS)                   | MacOSX Intel x86                   |
| MacOS M1 M2 | ![MacOS](https://img.shields.io/badge/N/A-inactive.svg?logo=MacOS)                         | ![MacOS](https://img.shields.io/badge/supported-success.svg?logo=MacOS)                    | MacOS Apple Silicon                |
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

💖💖💖如果你觉得此项目对你有帮助，请点亮 Star ✨✨✨

---

### ENERGY QQ交流群 & 微信

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