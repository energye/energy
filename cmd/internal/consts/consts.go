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
	DomainYangHY         = "https://energy.yanghy.cn"
	DomainGithub         = "https://energye.github.io"
	BASE_CONFIG_URL      = "/data/model-base-config.json"
	LATEST_VERSION_URL   = "/data/latest-version.json"
	CMD_VERSION_URL      = "/data/command-line-tools.json"
	VERSIONS_UPGRADE_URL = "/data/versions-upgrade.json"
	MODEL_CEF_URL        = "/data/model-cef.json"
	MODEL_LCL_URL        = "/data/model-liblcl.json"
)

const (
	CefKey         = "cef"
	LiblclKey      = "liblcl"
	FrameworkCache = "EnergyFrameworkDownloadCache"
)

var (
	// 生成项目多平台配置文件
	EnergyProjectConfig = []string{"energy_darwin.json", "energy_linux.json", "energy_windows.json"}
	// 项目配置模板
	ConfigFile = "energy_%s.json"
)

const (
	IsWindows = runtime.GOOS == "windows"
	IsLinux   = runtime.GOOS == "linux"
	IsDarwin  = runtime.GOOS == "darwin"
	Is386     = runtime.GOARCH == "386"
	IsAMD64   = runtime.GOARCH == "amd64"
	IsARM     = runtime.GOARCH == "arm"
	IsARM64   = runtime.GOARCH == "arm64"
	IsLoong64 = runtime.GOARCH == "loong64"
)

const (
	CEFLatestEmpty = ""
	CEFLatest      = "latest"
	CEF109         = "109"      // CEF 109
	CEF109118      = "109.1.18" // CEF 109.1.18
	CEF101         = "101"      // CEF 101
	CEF101018      = "101.0.18" // CEF 101.0.18
	CEF106         = "106"      // CEF 106
	CEF10611       = "106.1.1"  // CEF 106.1.1
	CEF87          = "87"       // CEF 87
	CEF89          = "89"       // CEF 87
	CEF87114       = "87.1.14"  // CEF 87.1.14
)

const (
	Windows32    = "Windows32"
	Windows64    = "Windows64"
	WindowsARM64 = "WindowsARM64"
	MacOSX64     = "MacOSX64"
	MacOSARM64   = "MacOSARM64"
	Linux32      = "Linux32"
	Linux64      = "Linux64"
	LinuxARM     = "LinuxARM"
	LinuxARM64   = "LinuxARM64"
	LinuxLoong64 = "LinuxLoong64"
)
