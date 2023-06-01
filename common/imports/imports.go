//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package imports Dynamic Library Import Based on LCL
// You can also import and add custom dynamic libraries here
package imports

import (
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/api/dllimports"
)

var (
	//energy CEF 导入
	energyImportDefs []*dllimports.ImportTable
	//energy 扩展 LCL 导入
	lclExtImportDefs []*dllimports.ImportTable
)

func SetEnergyImportDefs(importDefs []*dllimports.ImportTable) {
	energyImportDefs = importDefs
}

func SetLClExtImportDefs(importDefs []*dllimports.ImportTable) {
	lclExtImportDefs = importDefs
}

func Proc(index int) dllimports.ProcAddr {
	return api.ImportDefFunc(energyImportDefs, index)
}

func ExtProc(index int) dllimports.ProcAddr {
	return api.ImportDefFunc(lclExtImportDefs, index)
}
