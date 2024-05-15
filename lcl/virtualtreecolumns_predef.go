//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/energy/v2/api"
)

// TColumnsArray array of TVirtualTreeColumn;
type TColumnsArray struct {
	instance unsafePointer
	count    int
}

func (m *TVirtualTreeColumns) GetVisibleColumns() *TColumnsArray {
	var cols uintptr
	r1 := api.LCLPreDef().SysCallN(api.VirtualTreeColumnsGetVisibleColumns(), uintptr(unsafePointer(&cols)))
	return &TColumnsArray{
		count:    int(r1),
		instance: unsafePointer(cols),
	}
}

func (m *TColumnsArray) Count() int {
	return m.count
}

func (m *TColumnsArray) Get(index int) IVirtualTreeColumn {
	if m.instance == nil {
		return nil
	}
	if index < m.count {
		return AsVirtualTreeColumn(getParamOf(index, m.Instance()))
	}
	return nil
}

func (m *TColumnsArray) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *TColumnsArray) Free() {
	m.count = 0
	m.instance = nil
}
