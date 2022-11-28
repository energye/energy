# Energy 是Go基于CEF的构建桌面应用的框架
中文 |
[English](https://github.com/energye/energy/blob/main/README.en-US.md)

---
### [简介](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/6350f94ca749ba0318943f25)
> [Energy](https://github.com/energye/energy) 使用JavaScript、HTML和CSS 构建桌面应用的框架, 是Golang基于 [CEF4Delphi](https://github.com/salvadordf/CEF4Delphi) 和 [Golcl](https://github.com/energye/golcl) 开发的框架，内嵌 [Chromium CEF](https://bitbucket.org/chromiumembedded/cef) 二进制
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

#### [安装环境](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/63511b14a749ba0318943f3a)
> 使用energy命令行工具自动安装
>
> 1. go get github.com/energye/energy
>
> 2. cd energy/cmd/energy
>
> 3. go install
>
> 4. energy install .

### [入门指南-网址](https://energy.yanghy.cn)
* [教程](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/6350f94ca749ba0318943f25)
* [示例](https://energy.yanghy.cn/#/example/6342d986401bfe4d0cdf6067/634d3bd5a749ba0318943eb6)
* [文档](https://energy.yanghy.cn/#/document/6342d9a4401bfe4d0cdf6069/0)

### 快速入门
#### 基本需求
> golang >= 1.9.2
>
> 1. 使用energy命令行工具自动安装环境依赖
> 
> 2. 配置ENERGY_HOME环境变量, ENERGY_HOME=框架根目录
>
> 以example/simple示例为例
>
> go mod tidy
>
> 在IDE中运行simple
>
> 目前命令行工具不支持打包应用程序, windows你可以使用(MSI或Inno Setup)和其它绿色打包工具, linux下deb安装包等, MacOS默认开发时会生成.app包或者自行定制.app包

##### example/simple 示例代码
```go
package main

import (
	"github.com/energye/energy/cef"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalCEFInit(nil, nil)
	//创建应用
	cefApp := cef.NewApplication(nil)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.DefaultUrl = "https://energy.yanghy.cn"
	//运行应用
	cef.Run(cefApp)
}
```

----
### [License GPL 3.0](https://opensource.org/licenses/GPL-3.0)
