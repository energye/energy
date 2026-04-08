//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"unsafe"
)

var ptrSize = unsafe.Sizeof(uintptr(0))

const (
	libgtk3                    = "libgtk-3.so.0"
	libgdk3                    = "libgdk-3.so.0"
	libgobject2_0              = "libgobject-2.0.so.0"
	libglib2_0                 = "libglib-2.0.so.0"
	libgio2_0                  = "libgio-2.0.so.0"
	libcairo                   = "libcairo.so.2"
	libpango1_0                = "libpango-1.0.so.0"
	libwebkit2gtk4_0_37        = "libwebkit2gtk-4.0.so.37"
	libjavascriptcoregtk4_0_18 = "libjavascriptcoregtk-4.0.so.18"
)

type dnyLibrary struct {
	*imports.Imports
	index map[string]int
}

func libLoad(libName string) *dnyLibrary {
	dll, err := imports.NewDLL(libName)
	if err == nil {
		m := &dnyLibrary{Imports: &imports.Imports{Dll: dll}}
		m.NextType()
		return m
	}
	return nil
}

func setLibClose(lib *dnyLibrary) {
	api.SetOnReleaseCallback(func() {
		if lib != nil {
			lib.Dll.Release()
		}
	})
}

func (m *dnyLibrary) mapperIndex() {
	if m.index == nil {
		m.index = make(map[string]int, len(m.Table))
	}
	for i, table := range m.Table {
		m.index[table.Name()] = i
	}
}

func (m *dnyLibrary) Proc(index int) imports.ProcAddr {
	if m == nil {
		return 0
	}
	return m.Imports.Proc(index)
}

func (m *dnyLibrary) SysCallN(index int, args ...uintptr) uintptr {
	if m == nil || m.Imports == nil {
		return 0
	}
	return m.Imports.SysCallN(index, args...)
}

func (m *dnyLibrary) SysCall(name string, args ...uintptr) uintptr {
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
