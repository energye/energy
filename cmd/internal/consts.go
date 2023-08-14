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
	cefKey                   = "cef"
	energyKey                = "energy"
	downloadVersionConfigUrl = "https://energy.yanghy.cn/autoconfig/edv.json"
	downloadExtractUrl       = "https://energy.yanghy.cn/autoconfig/extract.json"
	frameworkCache           = "EnergyFrameworkDownloadCache"
	EnergyHomeKey            = "ENERGY_HOME"
)

const (
	isWindows = runtime.GOOS == "windows" //support
	isLinux   = runtime.GOOS == "linux"   //support
	isDarwin  = runtime.GOOS == "darwin"  //support
)
