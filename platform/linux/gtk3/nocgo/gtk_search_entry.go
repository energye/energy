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
	"unsafe"
)

type SearchEntry struct {
	Entry
}

func AsSearchEntry(ptr unsafe.Pointer) *SearchEntry {
	if ptr == nil {
		return nil
	}
	m := new(SearchEntry)
	m.instance = ptr
	return m
}

func NewSearchEntry() *SearchEntry {
	r := gtk3.SysCall("gtk_search_entry_new")
	if r == 0 {
		return nil
	}
	return AsSearchEntry(unsafe.Pointer(r))
}
