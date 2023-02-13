//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl/api"
)

var (
	cef_version       string
	lib_build_version string
)

// cef version
func CEFVersion() string {
	if cef_version == "" {
		r1, _, _ := imports.Proc(internale_CEFVersion).Call()
		cef_version = api.GoStr(r1)
	}
	return cef_version
}

// lib build version
func LibBuildVersion() string {
	if lib_build_version == "" {
		r1, _, _ := imports.Proc(internale_LibBuildVersion).Call()
		lib_build_version = api.GoStr(r1)
	}
	return lib_build_version
}
