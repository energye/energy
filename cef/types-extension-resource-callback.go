//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import "github.com/energye/energy/common/imports"

// Instance 实例
func (m *ICefGetExtensionResourceCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefGetExtensionResourceCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefGetExtensionResourceCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefGetExtensionResourceCallback) Cont(stream *ICefStreamReader) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefGetExtensionResourceCallback_Cont).Call(m.Instance(), stream.Instance())
}

func (m *ICefGetExtensionResourceCallback) Cancel() {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefGetExtensionResourceCallback_Cancel).Call(m.Instance())
}
