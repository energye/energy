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

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

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

// OverlayControllerRef -> ICefOverlayController
var OverlayControllerRef overlayController

type overlayController uintptr

func (*overlayController) UnWrap(data *ICefOverlayController) *ICefOverlayController {
	var result uintptr
	imports.Proc(def.OverlayControllerRefRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefOverlayController{instance: getInstance(result)}
	}
	return nil
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
	r1, _, _ := imports.Proc(def.OverlayController_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

// / Returns true (1) if this object is the same as |that| object.
func (m *ICefOverlayController) IsSame(that *ICefOverlayController) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.OverlayController_IsSame).Call(m.Instance(), that.Instance())
	return api.GoBool(r1)
}

// / Returns the contents View for this overlay.
func (m *ICefOverlayController) GetContentsView() *ICefView {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.OverlayController_GetContentsView).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefView{instance: getInstance(result)}
	}
	return nil
}

// / Returns the top-level Window hosting this overlay. Use this function
// / instead of calling get_window() on the contents View.
func (m *ICefOverlayController) GetWindow() *ICefWindow {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.OverlayController_GetWindow).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefWindow{&ICefPanel{&ICefView{instance: getInstance(result)}}}
	}
	return nil
}

// / Returns the docking mode for this overlay.
func (m *ICefOverlayController) GetDockingMode() consts.TCefDockingMode {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.OverlayController_GetDockingMode).Call(m.Instance())
	return consts.TCefDockingMode(r1)
}

// / Destroy this overlay.
func (m *ICefOverlayController) DestroyOverlay() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.OverlayController_DestroyOverlay).Call(m.Instance())
}

// / Sets the bounds (size and position) of this overlay. This will set the
// / bounds of the contents View to match and trigger a re-layout if necessary.
// / |bounds| is in parent coordinates and any insets configured on this
// / overlay will be ignored. Use this function only for overlays created with
// / a docking mode value of CEF_DOCKING_MODE_CUSTOM. With other docking modes
// / modify the insets of this overlay and/or layout of the contents View and
// / call size_to_preferred_size() instead to calculate the new size and re-
// / position the overlay if necessary.
func (m *ICefOverlayController) SetBounds(bounds TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.OverlayController_SetBounds).Call(m.Instance(), uintptr(unsafe.Pointer(&bounds)))
}

// / Returns the bounds (size and position) of this overlay in parent  coordinates.
func (m *ICefOverlayController) GetBounds() (bounds TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.OverlayController_GetBounds).Call(m.Instance(), uintptr(unsafe.Pointer(&bounds)))
	return
}

// / Returns the bounds (size and position) of this overlay in DIP screen coordinates.
func (m *ICefOverlayController) GetBoundsInScreen() (bounds TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.OverlayController_GetBoundsInScreen).Call(m.Instance(), uintptr(unsafe.Pointer(&bounds)))
	return
}

// / Sets the size of this overlay without changing the position. This will set
// / the size of the contents View to match and trigger a re-layout if
// / necessary. |size| is in parent coordinates and any insets configured on
// / this overlay will be ignored. Use this function only for overlays created
// / with a docking mode value of CEF_DOCKING_MODE_CUSTOM. With other docking
// / modes modify the insets of this overlay and/or layout of the contents View
// / and call size_to_preferred_size() instead to calculate the new size and
// / re-position the overlay if necessary.
func (m *ICefOverlayController) SetSize(size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.OverlayController_SetSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
}

// / Returns the size of this overlay in parent coordinates.
func (m *ICefOverlayController) GetSize() (size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.OverlayController_GetSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
	return
}

// / Sets the position of this overlay without changing the size. |position| is
// / in parent coordinates and any insets configured on this overlay will be
// / ignored. Use this function only for overlays created with a docking mode
// / value of CEF_DOCKING_MODE_CUSTOM. With other docking modes modify the
// / insets of this overlay and/or layout of the contents View and call
// / size_to_preferred_size() instead to calculate the new size and re-position
// / the overlay if necessary.
func (m *ICefOverlayController) SetPosition(position TCefPoint) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.OverlayController_SetPosition).Call(m.Instance(), uintptr(unsafe.Pointer(&position)))
}

// / Returns the position of this overlay in parent coordinates.
func (m *ICefOverlayController) GetPosition() (point TCefPoint) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.OverlayController_GetPosition).Call(m.Instance(), uintptr(unsafe.Pointer(&point)))
	return
}

// / Sets the insets for this overlay. |insets| is in parent coordinates. Use
// / this function only for overlays created with a docking mode value other
// / than CEF_DOCKING_MODE_CUSTOM.
func (m *ICefOverlayController) SetInsets(insets TCefInsets) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.OverlayController_SetInsets).Call(m.Instance(), uintptr(unsafe.Pointer(&insets)))
}

// / Returns the insets for this overlay in parent coordinates.
func (m *ICefOverlayController) GetInsets() (insets TCefInsets) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.OverlayController_GetInsets).Call(m.Instance(), uintptr(unsafe.Pointer(&insets)))
	return
}

// / Size this overlay to its preferred size and trigger a re-layout if
// / necessary. The position of overlays created with a docking mode value of
// / CEF_DOCKING_MODE_CUSTOM will not be modified by calling this function.
// / With other docking modes this function may re-position the overlay if
// / necessary to accommodate the new size and any insets configured on the
// / contents View.
func (m *ICefOverlayController) SizeToPreferredSize() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.OverlayController_SizeToPreferredSize).Call(m.Instance())
}

// / Sets whether this overlay is visible. Overlays are hidden by default. If
// / this overlay is hidden then it and any child Views will not be drawn and,
// / if any of those Views currently have focus, then focus will also be
// / cleared. Painting is scheduled as needed.
func (m *ICefOverlayController) SetVisible(visible bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.OverlayController_SetVisible).Call(m.Instance(), api.PascalBool(visible))
}

// / Returns whether this overlay is visible. A View may be visible but still
// / not drawn in a Window if any parent Views are hidden. Call is_drawn() to
// / determine whether this overlay and all parent Views are visible and will
// / be drawn.
func (m *ICefOverlayController) IsVisible() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.OverlayController_IsVisible).Call(m.Instance())
	return api.GoBool(r1)
}

// / Returns whether this overlay is visible and drawn in a Window. A View is
// / drawn if it and all parent Views are visible. To determine if the
// / containing Window is visible to the user on-screen call is_visible() on
// / the Window.
func (m *ICefOverlayController) IsDrawn() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.OverlayController_IsDrawn).Call(m.Instance())
	return api.GoBool(r1)
}
