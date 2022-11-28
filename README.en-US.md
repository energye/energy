# Energy is Go's framework for building desktop applications based on CEF
[中文 ](https://github.com/energye/energy/blob/main/README.md) |
English

---
### [Introduction](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/6350f94ca749ba0318943f25)
> [Energy](https://github.com/energye/energy) Is a framework for building desktop applications using JavaScript,HTML, and CSS based on [Golcl](https://github.com/energye/golcl) and [CEF4Delphi](https://github.com/salvadordf/CEF4Delphi) Pure Go language development framework, embedded [Chromium CEF](https://bitbucket.org/chromiumembedded/cef) binary
>
> Allows you to use [Web]() front-end technology to build cross-platform applications on Windows, Linux, and MacOS
>
>> You can use the [Energy](https://github.com/energye/energy) and [Golcl](https://github.com/energye/golcl) to build compact system UI cross-platform application separately
>
> Supports Windows_32, 64 bits, Linux_x86_64 bits, MacOS_x86_64 bits
> 
> In Go and Web technologies, based on IPC communication, you can easily interact data, function calls, and event calls between Go and Web without the need for a Web Service interface, just as easily as calling functions in the language itself
> 
> In Go, you can also define JS variables, which can be used by Web end JS to realize Go variable or structure object data synchronization
> 
> Call JS function and JS event listener in Go to realize functional interaction between GO and JS
> 
> Call Go function in JS, listen to events of Go, and realize functional interaction between JS and GO

### [![windows 32 bits](https://img.shields.io/badge/Downloads-green)](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/6364c5c2a749ba01d04ff485)

#### [Installation Environment](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/63511b14a749ba0318943f3a)
> Install it automatically using the energy command line tool
>
> 1. go get github.com/energye/energy
>
> 2. cd energy/cmd/energy
>
> 3. go install
>
> 4. energy install .

### [Getting started guide](https://energy.yanghy.cn)
* [tutorial](https://energy.yanghy.cn/#/course/6342d92c401bfe4d0cdf6065/6350f94ca749ba0318943f25)
* [sample](https://energy.yanghy.cn/#/example/6342d986401bfe4d0cdf6067/634d3bd5a749ba0318943eb6)
* [document](https://energy.yanghy.cn/#/document/6342d9a4401bfe4d0cdf6069/0)

### Quick start
#### Basic needs
> golang >= 1.9.2
>
> 1. Use the energy command line tool to automatically install environment dependencies
> 
> 2. Configure the ENERGY_HOME environment variable, ENERGY_HOME= framework root
>
> Example/simple example
>
> go mod tidy
>
> IDE run simple
>
> Currently, the command line tools do not support packaged applications, Windows you can use (MSI or Inno Setup) and other green packaging tools, Deb installation package in Linux, MacOS generates. App packages by default or custom. App packages

##### example/simple code
```go
package main

import (
	"github.com/energye/energy/cef"
)

func main() {
	//Global initialization, which every application must call
	cef.GlobalCEFInit(nil, nil)
	//Creating an Application
	cefApp := cef.NewApplication(nil)
	//Specify a URL address, or a local html file directory
	cef.BrowserWindow.Config.DefaultUrl = "https://energy.yanghy.cn"
	//Running the Application
	cef.Run(cefApp)
}

```

----
### [License GPL 3.0](https://opensource.org/licenses/GPL-3.0)