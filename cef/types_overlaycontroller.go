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
// include/capi/views/cef_overlay_controller_capi.h (cef_overlay_controller_t)
type ICefOverlayController struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}
