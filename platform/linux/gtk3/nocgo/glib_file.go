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

import "unsafe"

// File is a representation of GIO's GFile.
type File struct {
	Object
}

func AsFile(ptr unsafe.Pointer) *File {
	if ptr == nil {
		return nil
	}
	m := new(File)
	m.instance = ptr
	return m
}

// NativePrivate returns the underlying GFile pointer.
func (m *File) NativePrivate() uintptr {
	if m == nil {
		return 0
	}
	return m.Instance()
}

// Native returns the underlying GFile pointer as uintptr.
func (m *File) Native() uintptr {
	if m == nil {
		return 0
	}
	return m.Instance()
}

// NewFile is a wrapper around g_file_new_for_path().
func NewFile(path string) *File {
	return NewFileForPath(path)
}

// NewFileForPath is a wrapper around g_file_new_for_path().
func NewFileForPath(path string) *File {
	cstr := CStr(path)
	r := glib2_0.SysCall("g_file_new_for_path", cstr)
	if r == 0 {
		return nil
	}
	return AsFile(unsafe.Pointer(r))
}
