//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package internal

import "runtime"

const (
	//domain             = "https://energy.yanghy.cn"
	domain             = "http://localhost:8080"
	DownloadVersionURL = domain + "/api/cmd/edv"
	DownloadInstallURL = domain + "/api/cmd/install"
	DownloadExtractURL = domain + "/api/cmd/extract"
	CheckUpgradeURL    = domain + "/api/cmd/upgrade"
	CheckCreateURL     = domain + "/api/cmd/create"
)
const (
	cefKey         = "cef"
	liblclKey      = "liblcl"
	frameworkCache = "EnergyFrameworkDownloadCache"
	EnergyHomeKey  = "ENERGY_HOME"
)

const (
	isWindows = runtime.GOOS == "windows" //support
	isLinux   = runtime.GOOS == "linux"   //support
	isDarwin  = runtime.GOOS == "darwin"  //support
)

const (
	CefEmpty = ""
	CefWin7  = "windows7" // CEF 109.1.18
	Cef109   = "cef-109"  // CEF 109.1.18
	CefGtk2  = "gtk2"     // CEF 106.1.1
	Cef106   = "cef-106"  // CEF 106.1.1
	CefFlash = "flash"    // CEF 87.1.14
	Cef87    = "cef-87"   // CEF 87.1.14
)

const (
	Windows64      = "Windows64"
	Windows32      = "Windows32"
	WindowsARM64   = "WindowsARM64"
	MacOS64        = "MacOS64"
	MacOSARM64     = "MacOSARM64"
	Linux64        = "Linux64"
	Linux64GTK2    = "Linux64GTK2"
	Linux64GTK3    = "Linux64GTK3"
	LinuxARM64     = "LinuxARM64"
	LinuxARM64GTK2 = "LinuxARM64GTK2"
	LinuxARM64GTK3 = "LinuxARM64GTK3"
)

var SupportOSArchList = []string{
	Windows64,
	Windows32,
	WindowsARM64,
	MacOS64,
	MacOSARM64,
	Linux64,
	Linux64GTK2,
	Linux64GTK3,
	LinuxARM64,
	LinuxARM64GTK2,
	LinuxARM64GTK3,
}

const ()
