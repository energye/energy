<p align="center">
    <img src="https://assets.yanghy.cn/energy-doc/energy-icon.png">
</p>

<p align="center" style="font-size: 24px;">
    <strong>
        Energy是Go基于CEF构建桌面应用的框架
    </strong>
</p>

中文 |
[English](README.md)

---

[![github](https://img.shields.io/github/last-commit/energye/energy/main.svg?logo=github&logoColor=green&label=commit)](https://github.com/energye/energy)
[![release](https://img.shields.io/github/v/release/energye/energy?logo=git&logoColor=green)](https://github.com/energye/energy/releases)
[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=red)](http://www.apache.org/licenses/LICENSE-2.0)
![repo](https://img.shields.io/github/repo-size/energye/energy.svg?logo=github&logoColor=green&label=repo-size)

![go-version](https://img.shields.io/github/go-mod/go-version/energye/energy?logo=git&logoColor=green)
---

### [项目简介](https://energy.yanghy.cn/course/100/6350f94ca749ba0318943f25)

> [energy](https://github.com/energye/energy) 是 Go 基于 CEF(Chromium Embedded Framework)
> 开发的框架，内嵌 [CEF](https://bitbucket.org/chromiumembedded/cef) 二进制
>
> 使用 Go 和 Web 端技术 ( HTML + CSS + JavaScript ) 构建支持Windows, Linux, MacOS跨平台桌面应用
> 
> 需要会前端技术栈和略懂Go语言

### 特性

> - 具有丰富的CEF API 和 LCL 系统原生小部件
> - 开发环境简单,编译速度快,只需要Go和Energy依赖的CEF二进制框架
> - 跨平台: 一套代码可以打包成 Windows, 国产UOS、Deepin、Kylin, MacOS, Linux
> - 语言职责
>> - Go: Go负责窗口创建、CEF配置和功能实现、各种UI组件创建、系统低层调用，和JS处理不了的功能如: 文件流、安全加密、高性能处理等等
>> - Web: HTML + CSS + JavaScript 负责客户端界面的功能, 做出任意你想要的界面
> - 前端技术: 支持主流前端框架例如：Vue、React、Angular, 或纯HTML+CSS等等
> - 事件驱动: 高性能事件驱动, 基于IPC通信，实现Go和Web端很方便功能调用以及数据交互
> - 资源加载: 支持本地或内置到执行文件, 不需要http服务支撑, 多种选择

### 内置依赖&集成

- [![LCL](https://img.shields.io/badge/LCL-green)](https://github.com/energye/golcl)
- [![CEF-CEF4Delphi](https://img.shields.io/badge/CEF(Chromium%20Embedded%20Framework)%20CEF4Delphi-green)](https://github.com/salvadordf/CEF4Delphi)

### [开发环境](https://energy.yanghy.cn/course/100/63511b14a749ba0318943f3a)

#### 基本需求

> - Golang >= 1.18
> - Energy 开发环境(CEF, liblcl)

#### 环境安装

- 自动安装开发环境

> 使用 energy 命令行工具自动安装完整开发环境 [下载地址](https://energy.yanghy.cn/course/100/1694511322285207)
> 
> 此安装过程有选择性的从网络下载以下框架和工具

| 名称          | 平台             | 说明             |
|-------------|----------------|----------------|
| Golang      | ALL            | Go语言开发环境       |
| CEF, liblcl | ALL            | CEF框架          |
| NSIS        | Windows        | Windows安装包制作工具 |
| UPX         | Windows, Linux | 执行文件压缩工具       |
| 7z          | Windows, Linux | CEF框架压缩工具      |

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

### 项目打包
1. 编译：`energy build .`
2. 打包：`energy package .`
3. 最后自动生的安装包
   - windows  使用`nsis`工具生成exe安装包
   - linux    使用`dpkg`制作deb安装包
   - macos    生成`xxx.app`

### 系统支持

![Windows](https://img.shields.io/badge/windows-supported-success.svg?logo=Windows&logoColor=blue)
![MacOS](https://img.shields.io/badge/MacOS-supported-success.svg?logo=MacOS)
![Linux](https://img.shields.io/badge/Linux-supported-success.svg?logo=Linux&logoColor=red)


|             | 32位                                                                                        | 64位                                                                                        | 测试系统版本                             |
|-------------|--------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------|------------------------------------|
| Windows     | ![Windows](https://img.shields.io/badge/supported-success.svg?logo=Windows&logoColor=blue) | ![Windows](https://img.shields.io/badge/supported-success.svg?logo=Windows&logoColor=blue) | Windows 7、Windows 10、Windows 11    |
| MacOSX      | ![MacOS](https://img.shields.io/badge/N/A-inactive.svg?logo=MacOS)                         | ![MacOS](https://img.shields.io/badge/supported-success.svg?logo=MacOS)                    | MacOSX 10.15                       |
| MacOS M1 M2 | ![MacOS](https://img.shields.io/badge/N/A-inactive.svg?logo=MacOS)                         | ![MacOS](https://img.shields.io/badge/supported-success.svg?logo=MacOS)                    | MacOS M2, Rosetta2 AMD             |
| Linux       | ![Linux](https://img.shields.io/badge/自编译-supported-success.svg?logo=Linux)                | ![Linux](https://img.shields.io/badge/supported-success.svg?logo=Linux&logoColor=red)      | Deepin20.8、Ubuntu18.04、LinuxMint21 |
| Linux ARM   | ![Linux ARM](https://img.shields.io/badge/自编译-supported-success.svg?logo=Linux)            | ![Linux ARM](https://img.shields.io/badge/自编译-supported-success.svg?logo=Linux)            | Kylin-V10-SP1-2107                 |

---

### 鸣谢 Jetbrains

<p align="center">
    <a href="https://www.jetbrains.com?from=energy">
        <img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg" alt="JetBrains Logo (Main) logo.">
    </a>
</p>


感谢您给项目点个Star

---

### QQ交流群

[![QQGroup](https://img.shields.io/badge/QQ-541258627-green.svg?logo=tencentqq&logoColor=blue)](https://jq.qq.com/?_wv=1027&k=YgFjCGJX)

<p align="center">
    <img src="https://assets.yanghy.cn/energy-doc/qq-group.jpg" width="250">
</p>

---

### 项目截图
##### Windows-10
<img src="https://assets.yanghy.cn/energy-doc/frameless-windows-10.png">

##### Windows-7 32 & 64
<img src="https://assets.yanghy.cn/energy-doc/frameless-windows-7-64.png">
<img src="https://assets.yanghy.cn/energy-doc/frameless-windows-7-32.png">

##### Linux - 国产 Deepin
<img src="https://assets.yanghy.cn/energy-doc/frameless-deepin-20.8.png">
<img src="https://assets.yanghy.cn/energy-doc/frameless-deepin-hide-20.8.png">

##### Linux - 国产 Kylin ARM
<img src="https://assets.yanghy.cn/energy-doc/frameless-kylin-arm-V10-SP1.png">
<img src="https://assets.yanghy.cn/energy-doc/frameless-kylin-arm-hide-V10-SP1.png">

##### Linux - Ubuntu
<img src="https://assets.yanghy.cn/energy-doc/frameless-ubuntu-18.04.6.png">
<img src="https://assets.yanghy.cn/energy-doc/frameless-ubuntu-hide-18.04.6.png">

##### MacOSX
<img src="https://assets.yanghy.cn/energy-doc/frameless-macos.png">


----

### 开源协议

[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=green)](http://www.apache.org/licenses/LICENSE-2.0)

### 贡献者
<a href="https://github.com/energye/energy/graphs/contributors">
    <img src="https://opencollective.com/energy/contributors.svg?width=890&button=false" />
</a>