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
	"github.com/energye/golcl/lcl/api/dllimports"
)

var (

	//energy 扩展 LCL 导入
	liblclext *DllTable
	//energy CEF 导入
	libenergy *DllTable
)

func init() {
	liblclext = new(DllTable)
	libenergy = new(DllTable)
}

type DllTable struct {
	dll   dllimports.DLL
	table []*dllimports.ImportTable
	ok    bool
}

// LibEnergy libenergy
func LibEnergy() *DllTable {
	return libenergy
}

// LibLCLExt liblcl ext
func LibLCLExt() *DllTable {
	return liblclext
}

// Proc cef energy
func Proc(index int) dllimports.ProcAddr {
	return libenergy.Proc(index)
}

// SysCallN cef energy
func SysCallN(index int, args ...uintptr) uintptr {
	return libenergy.SysCallN(index, args...)
}

// LoadLib 加载动态库，适用自定义，第三方，支持Windows, MacOS, Linux
func LoadLib(libName string) *DllTable {
	if libName == "" {
		return &DllTable{}
	}
	dll, err := dllimports.NewDLL(libName)
	if dll == 0 && err != nil {
		return &DllTable{}
	}
	return &DllTable{
		dll: dll,
		ok:  true,
	}
}

func (m *DllTable) Proc(index int) dllimports.ProcAddr {
	if m.IsOk() {
		return dllimports.ImportDefFunc(m.dll, m.table, index)
	}
	return 0
}

func (m *DllTable) SysCallN(index int, args ...uintptr) (result uintptr) {
	if m.IsOk() {
		proc := dllimports.ImportDefFunc(m.dll, m.table, index)
		if proc > 0 {
			result, _, _ = proc.Call(args...)
		}
	}
	return
}

func (m *DllTable) IsOk() bool {
	return m.ok
}

func (m *DllTable) SetOk(v bool) {
	m.ok = v
}

func (m *DllTable) Dll() dllimports.DLL {
	return m.dll
}

func (m *DllTable) SetDll(dll dllimports.DLL) {
	m.dll = dll
}

func (m *DllTable) ImportTable() []*dllimports.ImportTable {
	return m.table
}

func (m *DllTable) SetImportTable(table []*dllimports.ImportTable) {
	m.table = table
}
