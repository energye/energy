# Energy 是Go基于CEF构建桌面应用的框架

中文 |
[English](README.md)

---
[![github](https://img.shields.io/github/last-commit/energye/energy/main.svg?logo=github&logoColor=green&label=commit)](https://github.com/energye/energy)
[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=red)](http://www.apache.org/licenses/LICENSE-2.0)
[![release](https://img.shields.io/github/v/release/energye/energy?logo=git&logoColor=green)](https://github.com/energye/energy/releases)
[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=red)](http://www.apache.org/licenses/LICENSE-2.0)
![downloads](https://img.shields.io/github/downloads/energye/energy/total?logo=git&logoColor=green)
![repo](https://img.shields.io/github/repo-size/energye/energy.svg?logo=github&logoColor=green&label=repo-size)

![go-version](https://img.shields.io/github/go-mod/go-version/energye/energy?logo=git&logoColor=green)
---

### [项目简介](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/6350f94ca749ba0318943f25)

> [energy](https://github.com/energye/energy) 是 Go 基于 CEF(Chromium Embedded Framework)
> 开发的框架，内嵌 [CEF](https://bitbucket.org/chromiumembedded/cef) 二进制
>
> 使用 Go 和 Web 端技术 ( HTML + CSS + JavaScript ) 构建支持Windows, Linux, MacOS跨平台桌面应用
>
> 基于Go和CEF跨平台特性

### Go和Web交互

> 在Go和Web技术基于IPC通信，可以在Go和Web交互数据、事件触发, 可以不使用 http 接口，就像调用语言本身函数一样简单
>
> 在Go中定义JS绑定类型变量，提供给Web端JS使用，实现Go变量或结构对象数据同步
>
> 在JS监听事件，在Go中触发JS事件，以达到Go调用JS函数和传递参数
>
> 在Go监听事件，在JS中触发Go事件，以达到JS调用Go函数和传递参数
>

### 内置依赖&集成

- [![golcl](https://img.shields.io/badge/Golcl-green)](https://github.com/energye/golcl)
- [![golcl](https://img.shields.io/badge/Liblcl-green)](https://github.com/energye/liblcl)
- [![CEFDelphi](https://img.shields.io/badge/CEFDelphi4-green)](https://github.com/salvadordf/CEF4Delphi)
- [![CEF](https://img.shields.io/badge/CEF(Chromium%20Embedded%20Framework)-green)](https://bitbucket.org/chromiumembedded/cef)

### [开发环境](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/63511b14a749ba0318943f3a)

#### 基本需求

> - golang >= 1.18
>
> - energy 开发环境
>
> 使用 energy 命令行工具自动安装开发环境
>
> 获取 [energy](https://github.com/energye/energy)
> 项目，或直接使用预编译命令行工具 [下载地址](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/63511b14a749ba0318943f3a)
> 1. go get github.com/energye/energy
>
> 进入 [energy](https://github.com/energye/energy) 命令行目录
> 2. cd energy/cmd/energy
>
> 安装命令行工具
> 3. go install
>
> 执行安装命令
> 4. energy install .

### 入门指南 - [传送门](https://energy.yanghy.cn)

* [教程](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/6350f94ca749ba0318943f25)
* [示例](https://energy.yanghy.cn/#/example/6342d986401bfe4d0cdf6067/634d3bd5a749ba0318943eb6)
* [文档](https://energy.yanghy.cn/#/document/6342d9a4401bfe4d0cdf6069/0)

### 快速入门

> 使用 [energy](https://github.com/energye/energy) 命令行工具自动安装环境依赖 `energy install .`
>
> 以example/simple示例为例
>
> 更新最新发布版本依赖
>
> 2. go mod tidy
>
> 在IDE中运行simple 或 go run simple.go

### example/simple 示例代码

```go
package main

import (
	"github.com/energye/energy/cef"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	cefApp := cef.NewApplication(nil)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	//运行应用
	cef.Run(cefApp)
}
```
---

### 项目打包 [参考](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/636e397ba749ba01d04ff595)
1. 编译：Go 程序编译`go build xxx.go` 如果使用资源内置(HTML、CSS、JavaScript、Image等等)会把资源编译到执行文件内
2. 复制：把执行文件复制到ENERGY环境的CEF目录中即可
3. 打包：使用制作安装包工具将其制作为安装包, 查阅各系统平台安装包制作
4. 最后：编译后的程序或安装包和CEF目录不再需要配置环境,在CEF根目录可直接运行

#### Go编译命令
1. 压缩并隐藏cmd窗口  `go build -ldflags "-H windowsgui -s -w"`, 注意: windows需要隐藏CMD窗口
2. 会压缩 不隐藏cmd窗口 `go build -ldflags "-s -w"`
2. 不压缩 不隐藏cmd窗口 `go build`

---

### 系统支持

![Windows](https://img.shields.io/badge/windows-supported-success.svg?logo=Windows&logoColor=blue)
![MacOSX](https://img.shields.io/badge/MacOSX-supported-success.svg?logo=MacOS)
![Linux](https://img.shields.io/badge/Linux-supported-success.svg?logo=Linux&logoColor=red)


|           | 32位                                                                             | 64位                                                                                        | 测试系统版本                               |
|-----------|---------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------|--------------------------------------|
| Windows   | ![Windows](https://img.shields.io/badge/supported-success.svg?logo=Windows&logoColor=blue)     | ![Windows](https://img.shields.io/badge/supported-success.svg?logo=Windows&logoColor=blue) | Windows 7、Windows 10                 |
| MacOSX    | ![MacOSX](https://img.shields.io/badge/N/A-inactive.svg?logo=MacOS)             | ![MacOSX](https://img.shields.io/badge/supported-success.svg?logo=MacOS)                   | MacOSX 10.15                         |
| Linux     | ![Linux](https://img.shields.io/badge/自编译-supported-success.svg?logo=Linux)     | ![Linux](https://img.shields.io/badge/supported-success.svg?logo=Linux&logoColor=red)      | Deepin20.8、Ubuntu18.04、LinuxMint21 |
| Linux ARM | ![Linux ARM](https://img.shields.io/badge/自编译-supported-success.svg?logo=Linux) | ![Linux ARM](https://img.shields.io/badge/自编译-supported-success.svg?logo=Linux)            | Kylin-V10-SP1-2107              |

---

### QQ交流群

[![QQGroup](https://img.shields.io/badge/QQ-541258627-green.svg?logo=tencentqq&logoColor=blue)](https://jq.qq.com/?_wv=1027&k=YgFjCGJX)

<img src="docs/qq-group.jpg" width="300">
---

### 项目截图
##### Windows-10
<img src="docs/frameless-windows-10.png" width="600">

##### Windows-7-32
<img src="docs/frameless-windows-7-32.png" width="600">

##### Windows-7-64
<img src="docs/frameless-windows-7-64.png" width="600">

##### Ubuntu-18.04.6
<img src="https://raw.githubusercontent.com/energye/energy/dev/docs/frameless-ubuntu-18.04.6.png" width="600">

##### MacOSX
<img src="docs/frameless-macos.png" width="600">


----

### 开源协议

[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=green)](http://www.apache.org/licenses/LICENSE-2.0)



