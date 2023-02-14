# Energy is the framework for Go to build desktop applications based on CEF

[中文](README.zh_CN.md) |
English

---
[![github](https://img.shields.io/github/last-commit/energye/energy/main.svg?logo=github&logoColor=green&label=commit)](https://github.com/energye/energy)
[![release](https://img.shields.io/github/v/release/energye/energy?logo=git&logoColor=green)](https://github.com/energye/energy/releases)
[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=red)](http://www.apache.org/licenses/LICENSE-2.0)
![downloads](https://img.shields.io/github/downloads/energye/energy/total?logo=git&logoColor=green)
![repo](https://img.shields.io/github/repo-size/energye/energy.svg?logo=github&logoColor=green&label=repo-size)

![go-version](https://img.shields.io/github/go-mod/go-version/energye/energy?logo=git&logoColor=green)
---

### [Project Introduction](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/6350f94ca749ba0318943f25)
Energy is a framework developed by Golang based on CEF(Chromium Embedded Framework), embedded with [CEF](https://bitbucket.org/chromiumembedded/cef) binary
> [energy](https://github.com/energye/energy) is a framework developed by Golang based on CEF(Chromium Embedded Framework), embedded with [CEF](https://bitbucket.org/chromiumembedded/cef) binary
>
> Use Go and Web technology (HTML+CSS+JavaScript) to build cross-platform desktop applications that support Windows, Linux and MacOS
>
> Based on Go and CEF cross-platform features

### Go interacts with the Web

> The Go and Web technologies are based on IPC communication.  Data and event triggering can be exchanged between Go and Web without using the http interface, which is as simple as calling the functions of the language itself
>
> Define JS binding type variables in Go and provide them to JS on the Web side to realize data synchronization of Go variables or structural objects
>
> Listen for events in JS and trigger JS events in Go to achieve Go calling JS functions and passing parameter
>
> Listen for events in Go and trigger Go events in JS to achieve JS calling Go functions and passing parameters
>

### Built-in dependency&integration

- [![golcl](https://img.shields.io/badge/Golcl-green)](https://github.com/energye/golcl)
- [![golcl](https://img.shields.io/badge/Liblcl-green)](https://github.com/energye/liblcl)
- [![CEFDelphi](https://img.shields.io/badge/CEFDelphi4-green)](https://github.com/salvadordf/CEF4Delphi)
- [![CEF](https://img.shields.io/badge/CEF(Chromium%20Embedded%20Framework)-green)](https://bitbucket.org/chromiumembedded/cef)

### [Development environment](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/63511b14a749ba0318943f3a)

> Install automatically using the energy command line tool
#### Basic needs
> - golang >= 1.18
>
> - energy development environment
> 
> Use the energy command line tool to automatically install the development environment
>
> Get [energy](https://github.com/energye/energy)
> project, or use the precompiled command line tool directly [Download address](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/63511b14a749ba0318943f3a)
> 1. go get github.com/energye/energy
>
> Enter the  [energy](https://github.com/energye/energy) command line directory
> 2. cd energy/cmd/energy
>
> Install command line tools
> 3. go install
>
> Execute the installation command
> 4. energy install .

### Getting Started Guide - [Transfer gate](https://energy.yanghy.cn)

* [course](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/6350f94ca749ba0318943f25)
* [example](https://energy.yanghy.cn/#/example/6342d986401bfe4d0cdf6067/634d3bd5a749ba0318943eb6)
* [document](https://energy.yanghy.cn/#/document/6342d9a4401bfe4d0cdf6069/0)

### Quick Get Start
> Use [energy](https://github.com/energye/energy) Command line tool automatic installation environment dependency `energy install .`
>
> Take example/simple as an example
>
> Update latest release dependency
>
> `go mod tidy`
>
> Run `simple` in the IDE or `go run simple.go`

### example/simple Code

```go
package main

import (
	"github.com/energye/energy/cef"
)

func main() {
	//Global initialization must be called by every application
	cef.GlobalInit(nil, nil)
	//Create application
	cefApp := cef.NewApplication(nil)
	//Set URL
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	//Run App
	cef.Run(cefApp)
}
```
---

### Pack Project [reference](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/636e397ba749ba01d04ff595)
1. Compile: Go program compilation `go build xxx. go` If you use resource built-in (HTML, CSS, JavaScript, Image, etc.), the resource will be compiled into the execution file
2. Copy: copy the execution file to the CEF directory of the ENERGY environment
3. Packaging: use the installation package tool to make it as an installation package, Or refer to the production of installation package for each system platform
4. Finally: the compiled program or installation package and CEF directory no longer need to configure the environment, and can be run directly in the CEF root directory

#### Go编译命令
1. Compress and hide the cmd window `go build -ldflags "-H windowsgui -s -w"`, Note: Windows needs to hide the CMD window
2. Compress `go build -ldflags "-s -w"`
2. Uncompressed `go build`

---

### System support

![Windows](https://img.shields.io/badge/windows-supported-success.svg?logo=Windows&logoColor=blue)
![MacOSX](https://img.shields.io/badge/MacOSX-supported-success.svg?logo=MacOS)
![Linux](https://img.shields.io/badge/Linux-supported-success.svg?logo=Linux&logoColor=red)


|           | 32 Bit                                                                                     | 64 Bit                                                                                     | Test System Version                |
|-----------|--------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------|------------------------------------|
| Windows   | ![Windows](https://img.shields.io/badge/supported-success.svg?logo=Windows&logoColor=blue) | ![Windows](https://img.shields.io/badge/supported-success.svg?logo=Windows&logoColor=blue) | Windows 7、Windows 10               |
| MacOSX    | ![MacOSX](https://img.shields.io/badge/N/A-inactive.svg?logo=MacOS)                        | ![MacOSX](https://img.shields.io/badge/supported-success.svg?logo=MacOS)                   | MacOSX 10.15                       |
| Linux     | ![Linux](https://img.shields.io/badge/SelfCompila-supported-success.svg?logo=Linux)    | ![Linux](https://img.shields.io/badge/supported-success.svg?logo=Linux&logoColor=red)      | Deepin20.8、Ubuntu18.04、LinuxMint21 |
| Linux ARM | ![Linux ARM](https://img.shields.io/badge/SelfCompila-supported-success.svg?logo=Linux)            | ![Linux ARM](https://img.shields.io/badge/SelfCompila-supported-success.svg?logo=Linux)            | Kylin-V10-SP1-2107                 |

---

### QQ Group

[![QQGroup](https://img.shields.io/badge/QQ-541258627-green.svg?logo=tencentqq&logoColor=blue)](https://jq.qq.com/?_wv=1027&k=YgFjCGJX)

<img src="https://assets.yanghy.cn/energy-doc/qq-group.jpg" width="300">
---

### Project screenshot
##### Windows-10
<img src="https://assets.yanghy.cn/energy-doc/frameless-windows-10.png" width="600">

##### Windows-7-32
<img src="https://assets.yanghy.cn/energy-doc/frameless-windows-7-32.png" width="600">

##### Windows-7-64
<img src="https://assets.yanghy.cn/energy-doc/frameless-windows-7-64.png" width="600">

##### Ubuntu-18.04.6
<img src="https://assets.yanghy.cn/energy-doc/frameless-ubuntu-18.04.6.png" width="600">

##### MacOSX
<img src="https://assets.yanghy.cn/energy-doc/frameless-macos.png" width="600">


----

### Public License

[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=green)](http://www.apache.org/licenses/LICENSE-2.0)



