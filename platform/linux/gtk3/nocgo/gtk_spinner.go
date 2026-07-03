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

type Spinner struct {
	Widget
}

func AsSpinner(ptr unsafe.Pointer) *Spinner {
	if ptr == nil {
		return nil
	}
	m := new(Spinner)
	m.instance = ptr
	return m
}

func NewSpinner() *Spinner {
	r := gtk3.SysCall("gtk_spinner_new")
	if r == 0 {
		return nil
	}
	return AsSpinner(unsafe.Pointer(r))
}

func (m *Spinner) Start() {
	gtk3.SysCall("gtk_spinner_start", m.Instance())
}

func (m *Spinner) Stop() {
	gtk3.SysCall("gtk_spinner_stop", m.Instance())
}
