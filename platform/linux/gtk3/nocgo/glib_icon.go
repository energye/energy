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
	"errors"
	"unsafe"
)

// Icon is a representation of GIO's GIcon.
type Icon struct {
	Object
}

func AsIcon(ptr unsafe.Pointer) *Icon {
	if ptr == nil {
		return nil
	}
	m := new(Icon)
	m.instance = ptr
	return m
}

// NativePrivate returns the underlying GIcon pointer.
func (m *Icon) NativePrivate() uintptr {
	if m == nil {
		return 0
	}
	return m.Instance()
}

// Native returns the underlying GIcon pointer.
func (m *Icon) Native() uintptr {
	if m == nil {
		return 0
	}
	return m.Instance()
}

// Equal is a wrapper around g_icon_equal().
func (m *Icon) Equal(icon *Icon) bool {
	r := glib2_0.SysCall("g_icon_equal", m.Instance(), icon.Instance())
	return ToGoBool(r)
}

// ToString is a wrapper around g_icon_to_string().
func (m *Icon) ToString() string {
	r := glib2_0.SysCall("g_icon_to_string", m.Instance())
	if r == 0 {
		return ""
	}
	s := GoStr(r)
	GFree(r)
	return s
}

// NewIconForString is a wrapper around g_icon_new_for_string().
func NewIconForString(str string) (*Icon, error) {
	cstr := CStr(str)
	var gErr uintptr
	r := glib2_0.SysCall("g_icon_new_for_string", cstr, uintptr(unsafe.Pointer(&gErr)))
	if gErr != 0 {
		gError := (*GError)(unsafe.Pointer(gErr))
		msg := GoStr(gError.Message)
		GErrorFree(gErr)
		return nil, errors.New(msg)
	}
	if r == 0 {
		return nil, errNilPtr
	}
	return AsIcon(unsafe.Pointer(r)), nil
}

// FileIcon is a representation of GIO's GFileIcon.
type FileIcon struct {
	Object
}

func AsFileIcon(ptr unsafe.Pointer) *FileIcon {
	if ptr == nil {
		return nil
	}
	m := new(FileIcon)
	m.instance = ptr
	return m
}

// NativePrivate returns the underlying GFileIcon pointer.
func (m *FileIcon) NativePrivate() uintptr {
	if m == nil {
		return 0
	}
	return m.Instance()
}

// Native returns the underlying GFileIcon pointer.
func (m *FileIcon) Native() uintptr {
	if m == nil {
		return 0
	}
	return m.Instance()
}

// NewFileIconN is a wrapper around g_file_icon_new() with a File argument.
func NewFileIconN(file *File) *Icon {
	r := glib2_0.SysCall("g_file_icon_new", file.Instance())
	if r == 0 {
		return nil
	}
	return AsIcon(unsafe.Pointer(r))
}

// NewFileIcon creates a new FileIcon from a file path.
func NewFileIcon(path string) *Icon {
	file := NewFile(path)
	if file == nil {
		return nil
	}
	return NewFileIconN(file)
}
