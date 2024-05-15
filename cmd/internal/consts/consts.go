//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package consts

import "runtime"

const (
	ENERGY = "energy"
)

const (
	domain             = "https://energy.yanghy.cn"
	DownloadVersionURL = domain + "/api/cmd/edv"
	DownloadInstallURL = domain + "/api/cmd/install"
	DownloadExtractURL = domain + "/api/cmd/extract"
	CheckUpgradeURL    = domain + "/api/cmd/upgrade"
	CheckCreateURL     = domain + "/api/cmd/create"
	LatestVersionURL   = domain + "/api/upgrade/latest/version"
)

const (
	// GolangDownloadURL 下载地址, 默认 1.18.10
	// https://golang.google.cn/dl/go1.18.10.darwin-amd64.tar.gz
	// https://golang.google.cn/dl/go1.18.10.darwin-arm64.tar.gz
	// https://golang.google.cn/dl/go1.18.10.linux-amd64.tar.gz
	// https://golang.google.cn/dl/go1.18.10.linux-arm64.tar.gz
	// https://golang.google.cn/dl/go1.18.10.windows-386.zip
	// https://golang.google.cn/dl/go1.18.10.windows-amd64.zip
	// https://golang.google.cn/dl/go1.18.10.windows-arm64.zip
	// https://golang.google.cn
	GolangDownloadSource = "https://golang.google.cn"
	GolangDownloadURL    = "%s/dl/%s"
	GolangDefaultVersion = "1.18.10"
	GolanHomeKey         = "GOROOT"
)

const (
	// NSISDownloadURL 下载地址，默认 3.09
	// https://gitee.com/energye/assets/releases/download/environment/nsis.windows.386-3.09.zip
	NSISDownloadURL     = "https://gitee.com/energye/assets/releases/download/environment/%s"
	NSISDownloadVersion = "3.09"
	NSISHomeKey         = "NSIS_HOME"

	// NSIS7z
	// https://gitee.com/energye/assets/releases/download/environment/nsis7z.windows.386-19.00.zip
	NSIS7zDownloadURL     = "https://gitee.com/energye/assets/releases/download/environment/%s"
	NSIS7zDownloadVersion = "19.00"
)

const (
	// https://gitee.com/energye/assets/releases/download/environment/7za.windows.all-2301.zip
	Z7ZDownloadURL     = "https://gitee.com/energye/assets/releases/download/environment/%s"
	Z7ZDownloadVersion = "2301"
	Z7ZHomeKey         = "Z7Z_HOME"
)

const (
	UPXHomeKey = "UPX_HOME"
)

const (
	CefKey              = "cef"
	LiblclKey           = "liblcl"
	FrameworkCache      = "EnergyFrameworkDownloadCache"
	EnergyHomeKey       = "ENERGY_HOME"
	EnergyProjectConfig = "energy.json"
)

const (
	IsWindows = runtime.GOOS == "windows" //support
	IsLinux   = runtime.GOOS == "linux"   //support
	IsDarwin  = runtime.GOOS == "darwin"  //support
	IsAMD64   = runtime.GOARCH == "amd64"
	IsARM64   = runtime.GOARCH == "arm64"
	Is386     = runtime.GOARCH == "386"
)

const (
	CefEmpty = ""
	Cef109   = "109" // CEF 109.1.18
	Cef106   = "106" // CEF 106.1.1
	Cef87    = "87"  // CEF 87.1.14
)

const (
	Windows64      = "Windows64"
	Windows32      = "Windows32"
	WindowsARM64   = "WindowsARM64"
	MacOSX64       = "MacOSX64"
	MacOSARM64     = "MacOSARM64"
	Linux64        = "Linux64"
	Linux64GTK2    = "Linux64GTK2"
	Linux64GTK3    = "Linux64GTK3"
	LinuxARM64     = "LinuxARM64"
	LinuxARM64GTK2 = "LinuxARM64GTK2"
	LinuxARM64GTK3 = "LinuxARM64GTK3"
)
