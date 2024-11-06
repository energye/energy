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
	ENERGY               = "energy"
	BASE_CONFIG_URL      = "https://energye.github.io/data/model-base-config.json"
	LATEST_VERSION_URL   = "https://energye.github.io/data/latest-version.json"
	CMD_VERSION_URL      = "https://energye.github.io/data/command-line-tools.json"
	VERSIONS_UPGRADE_URL = "https://energye.github.io/data/versions-upgrade.json"
	MODEL_CEF_URL        = "https://energye.github.io/data/model-cef.json"
	MODEL_LCL_URL        = "https://energye.github.io/data/model-liblcl.json"
)

const (
	CefKey         = "cef"
	LiblclKey      = "liblcl"
	GolanHomeKey   = "GOROOT"
	NSISHomeKey    = "NSIS_HOME"
	Z7ZHomeKey     = "Z7Z_HOME"
	UPXHomeKey     = "UPX_HOME"
	EnergyHomeKey  = "ENERGY_HOME"
	FrameworkCache = "EnergyFrameworkDownloadCache"
)

var (
	EnergyProjectConfig = []string{"energy_darwin.json", "energy_linux.json", "energy_windows.json"}
	ConfigFile          = "energy_%s.json"
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
