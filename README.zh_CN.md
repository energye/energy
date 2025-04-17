<p align="center">
    <img src="https://energye.github.io/imgs/energy.png">
</p>

<p align="center" style="font-size: 24px;">
    <strong>
        Energy是Go基于LCL和CEF开发的GUI框架
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
[![Go Report](https://goreportcard.com/badge/github.com/energye/energy)](https://goreportcard.com/report/github.com/cyber-xxm/energy/v2)
[![Go Reference](https://pkg.go.dev/badge/github.com/energye/energy)](https://pkg.go.dev/github.com/cyber-xxm/energy/v2)
[![Discord](https://img.shields.io/discord/1303173966747271209)](https://discord.gg/TejjxKz9)
[![codecov](https://codecov.io/gh/energye/energy/graph/badge.svg?token=H370UFUF12)](https://codecov.io/gh/energye/energy)
[![contributors](https://img.shields.io/github/contributors/energye/energy)](https://github.com/energye/energy/graphs/contributors)
[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=red)](http://www.apache.org/licenses/LICENSE-2.0)
---

### [项目简介](https://energye.github.io/course/what-is-energy)

> Energy 是 Go 基于 [LCL](https://gitlab.com/freepascal.org/lazarus/lazarus) 和 [CEF](https://bitbucket.org/chromiumembedded/cef)(Chromium Embedded Framework) 开发的 GUI 框架, 用于开发Windows、MacOS 和 Linux 跨平台桌面应用.
>
> 可仅用 LCL 创建轻量级原生应用，或结合 LCL 与 CEF 打造功能更全的应用程序
>
> 使用 Go 和 Web 端技术 ( HTML + CSS + JavaScript ) 构建支持Windows, Linux, MacOS跨平台桌面应用

### 特性

> - 具有丰富的 CEF API 和 丰富的 LCL 系统原生控件
> - 开发环境简单, 编译速度快, 仅需Go和Energy依赖的CEF二进制框架
> - 跨平台: 一套代码可以打包成 Windows, 国产UOS、Deepin、Kylin, MacOS, Linux
>>
>> - Golang: 窗口管理、CEF API封装&配置、功能实现、各种UI组件创建、系统低层调用和JavaScript处理不了的功能，如: 文件流、安全加密、高性能处理等
>> - Web: HTML + CSS + JavaScript 负责客户端界面的功能, 做出任意你想要的界面
>>
> - 前端技术: 支持主流前端框架。例如：Vue、React、Angular 和 原生HTML+CSS+JS等
> - 事件驱动: 高性能事件驱动, 基于IPC通信，实现Go和Web端迅捷调用及数据交互
> - 资源加载: 可无需http服务支撑，直接读取本地资源或内置到执行文件的资源, 也支持http服务加载资源

### 内置依赖&集成

- [![LCL](https://img.shields.io/badge/LCL-green)](https://github.com/energye/golcl)
- [![CEF-CEF4Delphi](https://img.shields.io/badge/CEF(Chromium%20Embedded%20Framework)%20CEF4Delphi-green)](https://github.com/salvadordf/CEF4Delphi)

### [开发环境](https://energye.github.io/course/getting-started)

#### 基本需求

> - Golang >= 1.18
> - Energy 开发环境(CEF, liblcl)

#### 环境安装

> 使用 [energy cli](https://energye.github.io/course/cli-use/) 自动安装完整开发环境

### 入门指南 - [传送门](https://energye.github.io/course/getting-started)

- [教程](https://energye.github.io/course/getting-started)
- [示例](https://energye.github.io/examples)
- [文档](https://energye.github.io/document)

### 快速入门

> [快速入门](https://energye.github.io/course/getting-started)

### 三个步骤运行一个简单应用

1. 安装开发环境: `energy install`
2. 初始化应用: `energy init`
3. 运行应用: `go run main.go`

### 示例代码

main.go

```go
package main

import (
    "github.com/cyber-xxm/energy/v2/cef"
)

func main() {
   //全局初始化
   cef.GlobalInit(nil, nil)
   //创建应用
   app := cef.NewApplication()
   //指定一个URL地址，或本地html文件目录
   cef.BrowserWindow.Config.Url = "https://energye.github.io"
   //运行应用
   cef.Run(app)
}
```

### 应用打包

1. 编译：`energy build`
2. 打包：`energy package`
3. 自动生成的安装包
   - windows  使用`nsis`工具生成exe安装包
   - linux    使用`dpkg`制作deb安装包
   - macos    生成`xxx.app`

### [系统支持](https://energye.github.io/document/version-details)

![Windows](https://img.shields.io/badge/Windows-✔️-success.svg?logo=Windows&logoColor=blue)
![MacOS](https://img.shields.io/badge/MacOS-✔️-success.svg?logo=MacOS)
![Linux](https://img.shields.io/badge/Linux-✔️-success.svg?logo=Linux&logoColor=red)

| OS                    | 32-bit | 64-bit |
|-----------------------|--------|--------|
| Windows               | ️✔️    | ️✔️    |
| MacOSX (Intel)        | ❌      | ️✔️    |
| MacOS (Apple Silicon) | ❌      | ️✔️    |
| Linux                 | ️✔️    | ️✔️    |
| Linux ARM             | ️✔️    | ️✔️    |


### v3.0 相关项目
- [LCL](https://github.com/energye/lcl)
- [CEF](https://github.com/energye/cef)
- [Webview2 Webkit2](https://github.com/energye/wv)

### 其它项目
- [CEF(Chromium Embedded Framework)](https://github.com/chromiumembedded/cef)
- [CEF4Delphi](https://github.com/salvadordf/CEF4Delphi)
- [CefSharp](https://github.com/cefsharp/CefSharp)
- [Java-CEF](https://bitbucket.org/chromiumembedded/java-cef)
- [cefpython](https://github.com/cztomczak/cefpython)
- [Chromium](https://chromium.googlesource.com/chromium/src/)

---

### ENERGY QQ交流群 & 微信

**创新功能破难题，点亮 `star` 共奋进！**

---

<p align="center">
    <img src="https://energye.github.io/imgs/assets/qq-group.jpg" width="250" title="QQ交流群: 541258627" alt="QQ交流群: 541258627">
    <img src="https://energye.github.io/imgs/assets/we-chat.jpg" width="250" title="微信: sniawmdf" alt="微信: sniawmdf" style="margin-left: 30px;">
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

<img src="https://energye.github.io/imgs/readme/frameless-windows-10.png">

##### Windows-7 32 & 64

<img src="https://energye.github.io/imgs/readme/frameless-windows-7-64.png">
<img src="https://energye.github.io/imgs/readme/frameless-windows-7-32.png">

##### Windows-XP-SP3

<img src="https://energye.github.io/imgs/readme/windows-xp-sp3.png">

##### Linux - Deepin

<img src="https://energye.github.io/imgs/readme/frameless-deepin-20.8.png">
<img src="https://energye.github.io/imgs/readme/frameless-deepin-hide-20.8.png">

##### Linux - Kylin ARM

<img src="https://energye.github.io/imgs/readme/frameless-kylin-arm-V10-SP1.png">
<img src="https://energye.github.io/imgs/readme/frameless-kylin-arm-hide-V10-SP1.png">

##### Linux - Ubuntu

<img src="https://energye.github.io/imgs/readme/frameless-ubuntu-18.04.6.png">
<img src="https://energye.github.io/imgs/readme/frameless-ubuntu-hide-18.04.6.png">

##### MacOS

<img src="https://energye.github.io/imgs/readme/frameless-macos.png">

----

### 开源协议

[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=green)](http://www.apache.org/licenses/LICENSE-2.0)

### 贡献者

<a href="https://github.com/energye/energy/graphs/contributors">
    <img src="https://opencollective.com/energy/contributors.svg?width=890&button=false" />
</a>

