//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
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
