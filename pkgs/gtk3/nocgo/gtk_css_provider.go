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
	. "github.com/energye/energy/v3/pkgs/gtk3/types"
	"unsafe"
)

// CssProvider is a representation of GTK's GtkCssProvider.
type CssProvider struct {
	Object
}

func AsCssProvider(ptr unsafe.Pointer) ICssProvider {
	if ptr == nil {
		return nil
	}
	m := new(CssProvider)
	m.instance = ptr
	return m
}

// CssProviderNew is a wrapper around gtk_css_provider_new().
func NewCssProvider() ICssProvider {
	r := gtk3.SysCall("gtk_css_provider_new")
	if r == 0 {
		return nil
	}
	return AsCssProvider(unsafe.Pointer(r))
}

// LoadFromPath is a wrapper around gtk_css_provider_load_from_path().
func (m *CssProvider) LoadFromPath(path string) error {
	var gErr uintptr
	gtk3.SysCall("gtk_css_provider_load_from_path", m.Instance(), CStr(path), uintptr(unsafe.Pointer(&gErr)))
	if gErr != 0 {
		gError := (*GError)(unsafe.Pointer(gErr))
		msg := GoStr(gError.Message)
		GErrorFree(gErr)
		return errors.New(msg)
	}
	return nil
}

// LoadFromData is a wrapper around gtk_css_provider_load_from_data().
func (m *CssProvider) LoadFromData(data string) error {
	var gErr uintptr
	cData := CStr(data)
	// len = ^uintptr(0) = -1
	gtk3.SysCall("gtk_css_provider_load_from_data", m.Instance(), cData, ^uintptr(0), uintptr(unsafe.Pointer(&gErr)))
	if gErr != 0 {
		gError := (*GError)(unsafe.Pointer(gErr))
		msg := GoStr(gError.Message)
		GErrorFree(gErr)
		return errors.New(msg)
	}
	return nil
}

// ToString is a wrapper around gtk_css_provider_to_string().
func (m *CssProvider) ToString() string {
	r := gtk3.SysCall("gtk_css_provider_to_string", m.Instance())
	return GoStr(r)
}
