//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

package term

import (
	"github.com/energye/golcl/lcl/rtl/version"
	"github.com/pterm/pterm"
)

func init() {
	// < windows 10 禁用颜色
	version.VersionInit()
	ov := version.OSVersion
	if ov.Major < 10 {
		pterm.DisableColor()
	}
}
