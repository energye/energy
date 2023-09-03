//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package command

import "runtime"

const (
	domain             = "https://energy.yanghy.cn"
	DownloadVersionURL = domain + "/api/cmd/edv"
	DownloadInstallURL = domain + "/api/cmd/install"
	DownloadExtractURL = domain + "/api/cmd/extract"
	CheckUpgradeURL    = domain + "/api/cmd/upgrade"
	CheckCreateURL     = domain + "/api/cmd/create"
)
const (
	CefKey         = "cef"
	LiblclKey      = "liblcl"
	FrameworkCache = "EnergyFrameworkDownloadCache"
	EnergyHomeKey  = "ENERGY_HOME"
)

const (
	IsWindows = runtime.GOOS == "windows" //support
	IsLinux   = runtime.GOOS == "linux"   //support
	IsDarwin  = runtime.GOOS == "darwin"  //support
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
