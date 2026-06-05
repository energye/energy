<p align="center">
   <img src="https://energye.github.io/imgs/energy.png">
   <strong style="font-size: 24px">3.0</strong>
</p>

<p align="center" style="font-size: 24px;">
    <strong>
        A Go framework for building cross-platform desktop applications based on LCL & CEF & Webview (Webview2, Webkit2)
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
[![Go Report](https://goreportcard.com/badge/github.com/energye/energy)](https://goreportcard.com/report/github.com/energye/energy/v2)
[![Go Reference](https://pkg.go.dev/badge/github.com/energye/energy)](https://pkg.go.dev/github.com/energye/energy/v2)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/energye/energy)
[![Discord](https://img.shields.io/discord/1303173966747271209)](https://discord.gg/TejjxKz9)
[![codecov](https://codecov.io/gh/energye/energy/graph/badge.svg?token=H370UFUF12)](https://codecov.io/gh/energye/energy)
[![contributors](https://img.shields.io/github/contributors/energye/energy)](https://github.com/energye/energy/graphs/contributors)
[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=red)](http://www.apache.org/licenses/LICENSE-2.0)
---

### [Introduction](https://energye.github.io/course/what-is-energy)

> [Energy](https://github.com/energye/energy) is a framework developed in Go using:
> [LCL](https://www.lazarus-ide.org/),
> [CEF](https://bitbucket.org/chromiumembedded/cef),
> [Webview2](https://learn.microsoft.com/en-us/microsoft-edge/webview2/),
> [Webkit2Gtk](https://webkitgtk.org/),
> [WKWebView](https://developer.apple.com/documentation/webkit/)
>
>> LCL — Native UI component library
>>
>> CEF — Browser component, CEF3 via [CEF4Delphi](https://github.com/salvadordf/CEF4Delphi)
>>
>> Webview2 — Browser component, Webview2 via [WebView4Delphi](https://github.com/salvadordf/WebView4Delphi); uses the system runtime
>>
>> Webkit2 (Webkit2Gtk / WKWebView) — Browser component, Webkit2; uses the system runtime


### Features

> - Native and Web can be used together (hybrid) or independently
> - Self-hosted GUI designer: Energy Designer
> - Hundreds of native controls, rich CEF API, and a lightweight Webview system runtime
> - Simple development environment — only Go and the Energy rendering runtime are required
> - Go backend: window management, CEF API encapsulation & configuration, feature implementation, various UI component creation, low-level system calls, and anything JavaScript cannot handle (e.g. file streams, encryption, high-performance processing)
> - Web frontend: HTML + CSS + JavaScript handles the client UI — build any interface you want
> - Frontend tech: supports mainstream frontend frameworks
> - Event-driven: high-performance event-driven IPC communication for fast invocation and data exchange between Go and the Web layer
> - Resource loading: load local resources or resources embedded into the executable directly, without an HTTP server; HTTP server loading is also supported

### Development Environment

> - Golang >= 1.20
> - Energy development environment ([CEF or System Runtime], libenergy runtime)

1. Install [Golang](https://go.dev/dl/)
2. Create a project from [Energy Designer](https://github.com/energye/designer) [Releases](https://github.com/energye/designer/releases)


## Create an App in 1 Minute

1. Launch Energy Designer
2. Create a new project
3. Drag components onto the canvas
4. Configure properties and events
5. Click Run to preview

### NO CGO

> Optionally develop in pure `Go` — no `CGO` compilation required

### [Examples](https://github.com/energye/examples/tree/main)


### Platform Support

![Windows](https://img.shields.io/badge/windows-supported-success.svg?logo=Windows&logoColor=blue)
![MacOS](https://img.shields.io/badge/MacOS-supported-success.svg?logo=MacOS)
![Linux](https://img.shields.io/badge/Linux-supported-success.svg?logo=Linux&logoColor=red)


|               | 32-bit                                                                                      | 64-bit                                                                                      | Tested OS Versions                  |
|---------------|---------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------|-------------------------------------|
| Windows       | ![Windows](https://img.shields.io/badge/supported-success.svg?logo=Windows&logoColor=blue)  | ![Windows](https://img.shields.io/badge/supported-success.svg?logo=Windows&logoColor=blue)  | 7, 10, 11                           |
| MacOSX        | ![MacOSX](https://img.shields.io/badge/N/A-inactive.svg?logo=MacOS)                         | ![MacOSX](https://img.shields.io/badge/supported-success.svg?logo=MacOS)                    | MacOSX Intel x86                    |
| MacOS M1/M2   | ![MacOS](https://img.shields.io/badge/N/A-inactive.svg?logo=MacOS)                          | ![MacOS](https://img.shields.io/badge/supported-success.svg?logo=MacOS)                     | MacOS Apple Silicon                 |
| Linux         | ![Linux](https://img.shields.io/badge/build%20yourself-supported-success.svg?logo=Linux)    | ![Linux](https://img.shields.io/badge/supported-success.svg?logo=Linux&logoColor=red)       | Deepin 20.8, Ubuntu 18.04, LinuxMint 21 |
| Linux ARM     | ![Linux ARM](https://img.shields.io/badge/build%20yourself-supported-success.svg?logo=Linux)| ![Linux ARM](https://img.shields.io/badge/supported-success.svg?logo=Linux)                 | Kylin-V10-SP1-2107                  |

### Related Projects
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

If you find this project helpful, please give us a Star!

---

### Screenshots

[image](https://github.com/energye/designer/tree/main/docs/image)

----

### License

[![license](https://img.shields.io/github/license/energye/energy.svg?logo=git&logoColor=green)](http://www.apache.org/licenses/LICENSE-2.0)

### Contributors
<a href="https://github.com/energye/energy/graphs/contributors">
    <img src="https://opencollective.com/energy/contributors.svg?width=890&button=false" />
</a>
