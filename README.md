<p align="center">
    <img src="https://energye.github.io/imgs/energy.png">
</p>

<p align="center" style="font-size: 24px;">
    <strong>
        Energy is a GUI framework developed by Go, based on LCL and CEF
    </strong>
</p>

[中文](README.zh_CN.md) |
English

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

### [Project Introduction](https://energye.github.io/en/course/what-is-energy)

- Energy is a GUI framework developed in Go, based on [LCL](https://gitlab.com/freepascal.org/lazarus/lazarus) and [CEF](https://bitbucket.org/chromiumembedded/cef)(Chromium Embedded Framework), which is used to develop cross-platform desktop applications for Windows, MacOS, and Linux.
- LCL alone can create lightweight native applications, or can be combined with CEF to create more feature-complete applications.
- Construct cross-platform desktop applications that support Windows, Linux, and MacOS using Go and web technologies (HTML + CSS + JavaScript).

### Characteristic

- Rich CEF API and LCL system native widgets
- Development environment is simple and the compilation speed is fast. Only need Go and Energy.
- Cross-Platform: You can package your code as Windows, domestic UOS, Deepin, Kylin, MacOS and Linux
- Responsibilities
  - Go: Go is responsible for creating Windows, configuring CEF and implementing functions, creating various UI components, calling the low level of the system and some functions that JS cannot handle, such as: file flow, security encryption, high-performance processing, etc.
  - Web: HTML + CSS + JavaScript is responsible for the function of the client interface, you can make any interface you want.
- Front-end technology: Support mainstream frameworks, such as Vue, React, Angular or pure HTML+CSS+JS.
- Event driven: High performance event driven communication based on IPC allowing intercommunication between Go and Web.
- Resource loading: You can either read local resources or resources built into execution files either with or without HTTP services.

### Built-in dependency&integration

- [![LCL](https://img.shields.io/badge/LCL-green)](https://github.com/energye/golcl)
- [![CEF-CEF4Delphi](https://img.shields.io/badge/CEF(Chromium%20Embedded%20Framework)%20CEF4Delphi-green)](https://github.com/salvadordf/CEF4Delphi)

### [Development Environment](https://energye.github.io/en/course/getting-started)

#### Basic needs

- Golang >= 1.18
- Energy (CEF, liblcl)

#### Environmental installation

Automatic installation development environment using the energy [command-line tool](https://energye.github.io/en/course/cli-use/)

### Guide to Start - [Link](https://energye.github.io/en/course/getting-started)

- [Guide](https://energye.github.io/en/course/getting-started)
- [Example](https://energye.github.io/en/examples)
- [Document](https://energye.github.io/en/document)

### Quick Start

- [Getting Started](https://energye.github.io/en/course/getting-started)

### Run a simple application by three steps

1. Install development environment: `energy install`
2. Initiate an Application: `energy init`
3. Run the Application: `go run main.go`

### sample code

main.go

```go
package main

import (
    "github.com/cyber-xxm/energy/v2/cef"
)

func main() {
     //Global initialization
     cef.GlobalInit(nil, nil)
     //Create an application
     app := cef.NewApplication()
     //Specify a URL address or local HTML file directory
     cef.BrowserWindow.Config.Url = "https://energye.github.io"
     //Run Application
     cef.Run(app)
}
```

### Build

1. Build：`energy build`
2. Package：`energy package`
3. Package Type：
    - windows: Using `nsis` create exe installation package
    - linux: Using `dpkg` create deb installation package
    - macos: Generate `xxx.app`

### [System support](https://energye.github.io/en/document/version-details)

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

### v3.0 Related Projects
- [LCL](https://github.com/energye/lcl)
- [CEF](https://github.com/energye/cef)
- [Webview2 Webkit2](https://github.com/energye/wv)

### Other Projects

- [CEF(Chromium Embedded Framework)](https://github.com/chromiumembedded/cef)
- [CEF4Delphi](https://github.com/salvadordf/CEF4Delphi)
- [CefSharp](https://github.com/cefsharp/CefSharp)
- [Java-CEF](https://bitbucket.org/chromiumembedded/java-cef)
- [cefpython](https://github.com/cztomczak/cefpython)
- [Chromium](https://chromium.googlesource.com/chromium/src/)

---

### ENERGY QQ Group & WeChat

**Innovative functions break through difficulties. Light up the `star` and strive forward together!**

---

<p align="center">
    <img src="https://energye.github.io/imgs/assets/qq-group.jpg" width="250" title="QQ Group: 541258627" alt="QQ Group: 541258627">
    <img src="https://energye.github.io/imgs/assets/we-chat.jpg" width="250" title="WeChat: sniawmdf" alt="WeChat: sniawmdf" style="margin-left: 30px;">
</p>

---

### Thanks Jetbrains

<p align="center">
    <a href="https://www.jetbrains.com?from=energy">
        <img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg" alt="JetBrains Logo (Main) logo.">
    </a>
</p>

---

### Project screenshot

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

##### MacOSX

<img src="https://energye.github.io/imgs/readme/frameless-macos.png">

----

### Public License

[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=green)](http://www.apache.org/licenses/LICENSE-2.0)

### Contributors

<a href="https://github.com/energye/energy/graphs/contributors">
    <img src="https://opencollective.com/energy/contributors.svg?width=890&button=false" />
</a>
