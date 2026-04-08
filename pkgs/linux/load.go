//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package linux

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

const (
	Libgtk3                    = "libgtk-3.so.0"
	Libgdk3                    = "libgdk-3.so.0"
	Libgobject2_0              = "libgobject-2.0.so.0"
	Libglib2_0                 = "libglib-2.0.so.0"
	Libgio2_0                  = "libgio-2.0.so.0"
	Libcairo                   = "libcairo.so.2"
	Libpango1_0                = "libpango-1.0.so.0"
	Libwebkit2gtk4_0_37        = "libwebkit2gtk-4.0.so.37"
	Libjavascriptcoregtk4_0_18 = "libjavascriptcoregtk-4.0.so.18"
)

type DnyLibrary struct {
	*imports.Imports
	index map[string]int
}

func LibLoad(libName string) *DnyLibrary {
	dll, err := imports.NewDLL(libName)
	if err == nil {
		m := &DnyLibrary{Imports: &imports.Imports{Dll: dll}}
		m.NextType()
		return m
	}
	return nil
}

func (m *DnyLibrary) SetLibClose() {
	if m == nil {
		return
	}
	api.SetOnReleaseCallback(func() {
		m.Dll.Release()
	})
}

func (m *DnyLibrary) MapperIndex() {
	if m.index == nil {
		m.index = make(map[string]int, len(m.Table))
	}
	for i, table := range m.Table {
		m.index[table.Name()] = i
	}
}

func (m *DnyLibrary) Proc(index int) imports.ProcAddr {
	if m == nil {
		return 0
	}
	return m.Imports.Proc(index)
}

func (m *DnyLibrary) SysCallN(index int, args ...uintptr) uintptr {
	if m == nil || m.Imports == nil {
		return 0
	}
	return m.Imports.SysCallN(index, args...)
}

func (m *DnyLibrary) SysCall(name string, args ...uintptr) uintptr {
	if m == nil || m.Imports == nil {
		return 0
	}
	if m.index != nil {
		if index, ok := m.index[name]; ok {
			return m.SysCallN(index, args...)
		}
	}
	return 0
}
