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
	domain             = "https://energy.yanghy.cn"
	DownloadVersionURL = domain + "/api/cmd/edv"
	DownloadInstallURL = domain + "/api/cmd/install"
	DownloadExtractURL = domain + "/api/cmd/extract"
	CheckUpgradeURL    = domain + "/api/cmd/upgrade"
	CheckCreateURL     = domain + "/api/cmd/create"
)
const (
	cefKey         = "cef"
	energyKey      = "energy"
	frameworkCache = "EnergyFrameworkDownloadCache"
	EnergyHomeKey  = "ENERGY_HOME"
)

const (
	isWindows = runtime.GOOS == "windows" //support
	isLinux   = runtime.GOOS == "linux"   //support
	isDarwin  = runtime.GOOS == "darwin"  //support
)
