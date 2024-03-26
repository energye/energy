//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// 导入表实例

package imports

import "errors"

// CallImport 调用导入表接口
type CallImport interface {
	Proc(index int) ProcAddr
	SysCallN(index int, args ...uintptr) (result uintptr)
}

// Imports 导入表实例
type Imports struct {
	dll   DLL
	table []*Table
	ok    bool
}

// LoadLib 加载动态库，适用自定义导入动态链接库，支持Windows, MacOS, Linux
func LoadLib(libName string) (*Imports, error) {
	if libName == "" {
		return &Imports{}, errors.New("the dynamic link library is empty")
	}
	dll, err := NewDLL(libName)
	if dll == 0 && err != nil {
		return nil, err
	}
	return &Imports{
		dll: dll,
		ok:  true,
	}, nil
}

// Proc 返回api实例
//  参数
//	index: 导入表索引
func (m *Imports) Proc(index int) ProcAddr {
	if m.IsOk() {
		return internalGetImportFunc(m.dll, m.table, index)
	}
	return 0
}

// SysCallN 调用api
//  参数
//	index: 导入表索引
//	args: 调用api参数, 指针数组
func (m *Imports) SysCallN(index int, args ...uintptr) (result uintptr) {
	if m.IsOk() {
		proc := internalGetImportFunc(m.dll, m.table, index)
		if proc > 0 {
			result, _, _ = proc.Call(args...)
		}
	}
	return
}

// IsOk 是否成功导入, 如果为false不能调用api
func (m *Imports) IsOk() bool {
	return m.ok
}

// SetOk 设置是否成功导入
func (m *Imports) SetOk(v bool) {
	m.ok = v
}

// Dll 返回导入DLL实例
func (m *Imports) Dll() DLL {
	return m.dll
}

// SetDll 设置导入DLL实例
func (m *Imports) SetDll(dll DLL) {
	m.dll = dll
}

// ImportTable 返回导入表
func (m *Imports) ImportTable() []*Table {
	return m.table
}

// SetImportTable 设置导入表
func (m *Imports) SetImportTable(table []*Table) {
	m.table = table
}
