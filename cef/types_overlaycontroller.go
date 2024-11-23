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

import "unsafe"

// ICefOverlayController TODO 未实现
//
//	Controller for an overlay that contains a contents View added via
//	ICefWindow.AddOverlayView. Methods exposed by this controller should be
//	called in preference to functions of the same name exposed by the contents
//	View unless otherwise indicated. Methods must be called on the browser
//	process UI thread unless otherwise indicated.
//	<para><see cref="uCEFTypes|TCefOverlayController">Implements TCefOverlayController</see></para>
//	<para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_overlay_controller_capi.h">CEF source file: /include/capi/views/cef_overlay_controller_capi.h (cef_overlay_controller_t)</see></para>
type ICefOverlayController struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefOverlayController) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefOverlayController) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefOverlayController) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}
