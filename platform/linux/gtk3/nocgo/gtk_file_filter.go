//----------------------------------------
// Copyright © yanghy. All Rights Reserved.
// Licensed under Apache License Version 2.0, January 2004
// https://www.apache.org/licenses/LICENSE-2.0
//----------------------------------------

package nocgo

import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// FileFilter is a representation of GTK's GtkFileFilter.
type FileFilter struct {
	Object
}

func AsFileFilter(ptr unsafe.Pointer) *FileFilter {
	if ptr == nil {
		return nil
	}
	m := new(FileFilter)
	m.instance = ptr
	return m
}

// NewFileFilter is a wrapper around gtk_file_filter_new().
func NewFileFilter() IFileFilter {
	r := gtk3.SysCall("gtk_file_filter_new")
	if r == 0 {
		return nil
	}
	return AsFileFilter(unsafe.Pointer(r))
}

// SetName is a wrapper around gtk_file_filter_set_name().
func (m *FileFilter) SetName(name string) {
	gtk3.SysCall("gtk_file_filter_set_name", m.Instance(), CStr(name))
}

// GetName is a wrapper around gtk_file_filter_get_name().
func (m *FileFilter) GetName() string {
	r := gtk3.SysCall("gtk_file_filter_get_name", m.Instance())
	if r == 0 {
		return ""
	}
	return GoStr(r)
}

// AddPattern is a wrapper around gtk_file_filter_add_pattern().
func (m *FileFilter) AddPattern(pattern string) {
	gtk3.SysCall("gtk_file_filter_add_pattern", m.Instance(), CStr(pattern))
}

// AddMimeType is a wrapper around gtk_file_filter_add_mime_type().
func (m *FileFilter) AddMimeType(mimeType string) {
	gtk3.SysCall("gtk_file_filter_add_mime_type", m.Instance(), CStr(mimeType))
}
