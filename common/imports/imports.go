//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

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
