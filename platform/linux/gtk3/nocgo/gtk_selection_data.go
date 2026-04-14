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
	. "github.com/energye/energy/v3/platform/linux/types"
	"reflect"
	"unsafe"
)

type SelectionData struct {
	instance unsafe.Pointer
}

func AsSelectionData(ptr unsafe.Pointer) ISelectionData {
	if ptr == nil {
		return nil
	}
	m := new(SelectionData)
	m.instance = ptr
	return m
}

func (m *SelectionData) Instance() uintptr {
	return uintptr(m.instance)
}

// GetLength is a wrapper around gtk_selection_data_get_length().
func (m *SelectionData) GetLength() int {
	r := gtk3.SysCall("gtk_selection_data_get_length", m.Instance())
	return int(r)
}

// GetData is a wrapper around gtk_selection_data_get_data_with_length().
// It returns a slice of the correct size with the copy of the selection's data.
func (m *SelectionData) GetData() []byte {
	var length uintptr
	r := gtk3.SysCall("gtk_selection_data_get_data_with_length", m.Instance(), uintptr(unsafe.Pointer(&length)))
	// Only set if length is valid.
	if int(length) < 1 {
		return nil
	}
	var data []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sliceHeader.Len = int(length)
	sliceHeader.Cap = int(length)
	sliceHeader.Data = r
	return data
}

// SetData is a wrapper around gtk_selection_data_set_data_with_length().
func (m *SelectionData) SetData(atom TAtom, data []byte) {
	gtk3.SysCall("gtk_selection_data_set", m.Instance(), uintptr(atom), uintptr(8),
		uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
}

// GetText is a wrapper around gtk_selection_data_get_text(). It returns a copy
// of the string from SelectionData and frees the C reference.
func (m *SelectionData) GetText() string {
	charPtr := gtk3.SysCall("gtk_selection_data_get_text", m.Instance())
	if charPtr == 0 {
		return ""
	}
	defer GFree(charPtr)
	return ucharString(charPtr)

}

// SetText is a wrapper around gtk_selection_data_set_text().
func (m *SelectionData) SetText(text string) bool {
	textPtr := *(*[]byte)(unsafe.Pointer(&text))
	r := gtk3.SysCall("gtk_selection_data_set_text", m.Instance(),
		uintptr(unsafe.Pointer(&textPtr[0])), uintptr(len(text)))
	return ToGoBool(r)
}

// SetURIs is a wrapper around gtk_selection_data_set_uris().
func (m *SelectionData) SetURIs(uris []string) bool {
	if m.Instance() == 0 || len(uris) == 0 {
		return false
	}
	data, items := makeCStringArray(uris)
	defer freeCStringArray(data, items)
	r := gtk3.SysCall("gtk_selection_data_set_uris", m.Instance(), data)
	return ToGoBool(r)
}

// GetURIs is a wrapper around gtk_selection_data_get_uris().
func (m *SelectionData) GetURIs() []string {
	uriPtrs := gtk3.SysCall("gtk_selection_data_get_uris", m.Instance())
	return toGoStringArray(uriPtrs)
}

func (m *SelectionData) Free() {
	gtk3.SysCall("gtk_selection_data_free", m.Instance())
}
