//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https//www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// TCEFViewComponent
type TCEFViewComponent struct {
	lcl.IComponent
	instance unsafe.Pointer
}

// ViewComponentRef -> TCEFViewComponent
var ViewComponentRef viewComponent

type viewComponent uintptr

func (*viewComponent) New(AOwner lcl.IComponent) *TCEFViewComponent {
	var result uintptr
	imports.Proc(def.ViewComponent_Create).Call(AOwner.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &TCEFViewComponent{instance: getInstance(result)}
	}
	return nil
}

// Instance 实例
func (m *TCEFViewComponent) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *TCEFViewComponent) Free() {
	if m.instance != nil {
		imports.Proc(def.ViewComponent_Free).Call(m.Instance())
		m.instance = nil
	}
}

func (m *TCEFViewComponent) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(def.ViewComponent_GetValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFViewComponent) ToStringEx(includeChildren bool) string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.ViewComponent_ToStringEx).Call(m.Instance(), api.PascalBool(includeChildren))
	return api.GoStr(r1)
}

func (m *TCEFViewComponent) IsSame(that *ICefView) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ViewComponent_IsSame).Call(m.Instance(), that.Instance())
	return api.GoBool(r1)
}

func (m *TCEFViewComponent) SizeToPreferredSize() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_SizeToPreferredSize).Call(m.Instance())
}

func (m *TCEFViewComponent) InvalidateLayout() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_InvalidateLayout).Call(m.Instance())
}

func (m *TCEFViewComponent) RequestFocus() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_RequestFocus).Call(m.Instance())
}

func (m *TCEFViewComponent) ConvertPointToScreen(point TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ViewComponent_ConvertPointToScreen).Call(m.Instance(), uintptr(unsafe.Pointer(&point)))
	return api.GoBool(r1)
}

func (m *TCEFViewComponent) ConvertPointFromScreen(point TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ViewComponent_ConvertPointFromScreen).Call(m.Instance(), uintptr(unsafe.Pointer(&point)))
	return api.GoBool(r1)
}

func (m *TCEFViewComponent) ConvertPointToWindow(point TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ViewComponent_ConvertPointToWindow).Call(m.Instance(), uintptr(unsafe.Pointer(&point)))
	return api.GoBool(r1)
}

func (m *TCEFViewComponent) ConvertPointFromWindow(point TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ViewComponent_ConvertPointFromWindow).Call(m.Instance(), uintptr(unsafe.Pointer(&point)))
	return api.GoBool(r1)
}

func (m *TCEFViewComponent) ConvertPointToView(view *ICefView, point TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ViewComponent_ConvertPointToView).Call(m.Instance(), view.Instance(), uintptr(unsafe.Pointer(&point)))
	return api.GoBool(r1)
}

func (m *TCEFViewComponent) ConvertPointFromView(view *ICefView, point TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ViewComponent_ConvertPointFromView).Call(m.Instance(), view.Instance(), uintptr(unsafe.Pointer(&point)))
	return api.GoBool(r1)
}

func (m *TCEFViewComponent) Initialized() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ViewComponent_GetInitialized).Call(m.Instance())
	return api.GoBool(r1)
}

// Returns this control as a View.
func (m *TCEFViewComponent) AsView() *ICefView {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ViewComponent_GetAsView).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefView{instance: getInstance(result)}
	}
	return nil
}

// Returns this View as a BrowserView or NULL if this is not a BrowserView.
func (m *TCEFViewComponent) AsBrowserView() *ICefBrowserView {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ViewComponent_GetAsBrowserView).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefBrowserView{&ICefView{instance: getInstance(result)}}
	}
	return nil
}

// Returns this View as a Button or NULL if this is not a Button.
func (m *TCEFViewComponent) AsButton() *ICefButton {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ViewComponent_GetAsButton).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefButton{&ICefView{instance: getInstance(result)}}
	}
	return nil
}

// Returns this View as a Panel or NULL if this is not a Panel.
func (m *TCEFViewComponent) AsPanel() *ICefPanel {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ViewComponent_GetAsPanel).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefPanel{&ICefView{instance: getInstance(result)}}
	}
	return nil
}

// Returns this View as a ScrollView or NULL if this is not a ScrollView.
func (m *TCEFViewComponent) AsScrollView() *ICefScrollView {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ViewComponent_GetAsScrollView).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefScrollView{&ICefView{instance: getInstance(result)}}
	}
	return nil
}

// Returns this View as a Textfield or NULL if this is not a Textfield.
func (m *TCEFViewComponent) AsTextfield() *ICefTextfield {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ViewComponent_GetAsTextfield).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefTextfield{&ICefView{instance: getInstance(result)}}
	}
	return nil
}

// Recursively descends the view tree starting at this View, and returns the
// first child that it encounters with the given ID. Returns NULL if no
// matching child view is found.
func (m *TCEFViewComponent) GetViewForID(id int32) *ICefView {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ViewComponent_GetViewForID).Call(m.Instance(), uintptr(id), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefView{instance: getInstance(result)}
	}
	return nil
}

// Returns the delegate associated with this View, if any.
func (m *TCEFViewComponent) IsAttached() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ViewComponent_GetAttached).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFViewComponent) GetDelegate() *ICefViewDelegate {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ViewComponent_GetDelegate).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefViewDelegate{instance: getInstance(result)}
	}
	return nil
}

// Returns the top-level Window hosting this View, if any.
func (m *TCEFViewComponent) GetWindow() *ICefWindow {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ViewComponent_GetWindow).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefWindow{&ICefPanel{&ICefView{instance: getInstance(result)}}}
	}
	return nil
}

// Returns the View that contains this View, if any.
func (m *TCEFViewComponent) GetParentView() *ICefView {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ViewComponent_GetParentView).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefView{instance: getInstance(result)}
	}
	return nil
}

// Returns the bounds (size and position) of this View in DIP screen
// coordinates.
func (m *TCEFViewComponent) BoundsInScreen() (bounds TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_GetBoundsInScreen).Call(m.Instance(), uintptr(unsafe.Pointer(&bounds)))
	return
}

// Returns the size this View would like to be if enough space is available.
// Size is in parent coordinates, or DIP screen coordinates if there is no
// parent.
func (m *TCEFViewComponent) GetPreferredSize() (size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_GetPreferredSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
	return
}

// Returns the minimum size for this View. Size is in parent coordinates, or
// DIP screen coordinates if there is no parent.
func (m *TCEFViewComponent) MinimumSize() (size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_GetMinimumSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
	return
}

// Returns the maximum size for this View. Size is in parent coordinates, or
// DIP screen coordinates if there is no parent.
func (m *TCEFViewComponent) MaximumSize() (size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_GetMaximumSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
	return
}

// Returns whether this View is visible. A view may be visible but still not
// drawn in a Window if any parent views are hidden. If this View is a Window
// then a return value of true (1) indicates that this Window is currently
// visible to the user on-screen. If this View is not a Window then call
// is_drawn() to determine whether this View and all parent views are visible
// and will be drawn.
func (m *TCEFViewComponent) IsVisible() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ViewComponent_GetVisible).Call(m.Instance())
	return api.GoBool(r1)
}

// Sets whether this View is visible. Windows are hidden by default and other
// views are visible by default. This View and any parent views must be set
// as visible for this View to be drawn in a Window. If this View is set as
// hidden then it and any child views will not be drawn and, if any of those
// views currently have focus, then focus will also be cleared. Painting is
// scheduled as needed. If this View is a Window then calling this function
// is equivalent to calling the Window show() and hide() functions.
func (m *TCEFViewComponent) SetVisible(visible bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_SetVisible).Call(m.Instance(), api.PascalBool(visible))
}

// Returns whether this View is visible and drawn in a Window. A view is
// drawn if it and all parent views are visible. If this View is a Window
// then calling this function is equivalent to calling is_visible().
// Otherwise, to determine if the containing Window is visible to the user
// on-screen call is_visible() on the Window.
func (m *TCEFViewComponent) IsDrawn() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ViewComponent_GetDrawn).Call(m.Instance())
	return api.GoBool(r1)
}

// Set whether this View is enabled. A disabled View does not receive
// keyboard or mouse inputs. If |enabled| differs from the current value the
// View will be repainted. Also, clears focus if the focused View is
// disabled.
func (m *TCEFViewComponent) IsEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ViewComponent_GetEnabled).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFViewComponent) SetEnabled(enabled bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_SetEnabled).Call(m.Instance(), api.PascalBool(enabled))
}

// Returns true (1) if this View is focusable, enabled and drawn.
func (m *TCEFViewComponent) IsFocusable() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ViewComponent_GetFocusable).Call(m.Instance())
	return api.GoBool(r1)
}

// Sets whether this View is capable of taking focus. It will clear focus if
// the focused View is set to be non-focusable. This is false (0) by default
// so that a View used as a container does not get the focus.
func (m *TCEFViewComponent) SetFocusable(focusable bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_SetFocusable).Call(m.Instance(), api.PascalBool(focusable))
}

// Return whether this View is focusable when the user requires full keyboard
// access, even though it may not be normally focusable.
func (m *TCEFViewComponent) IsAccessibilityFocusable() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ViewComponent_GetAccessibilityFocusable).Call(m.Instance())
	return api.GoBool(r1)
}

// Returns the background color for this View. If the background color is
// unset then the current `GetThemeColor(CEF_ColorPrimaryBackground)` value
// will be returned. If this View belongs to an overlay (created with
// ICefWindow.AddOverlayView), and the background color is unset, then a
// value of transparent (0) will be returned.
func (m *TCEFViewComponent) GetBackgroundColor() (color types.TCefColor) {
	if !m.IsValid() {
		return 0
	}
	imports.Proc(def.ViewComponent_GetBackgroundColor).Call(m.Instance(), uintptr(unsafe.Pointer(&color)))
	return
}

// Sets the background color for this View. The background color will be
// automatically reset when ICefViewDelegate.OnThemeChanged is called.
func (m *TCEFViewComponent) SetBackgroundColor(color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_SetBackgroundColor).Call(m.Instance(), uintptr(color))
}

// Returns the ID for this View.
func (m *TCEFViewComponent) GetID() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.ViewComponent_GetID).Call(m.Instance())
	return int32(r1)
}

// Sets the ID for this View. ID should be unique within the subtree that you
// intend to search for it. 0 is the default ID for views.
func (m *TCEFViewComponent) SetID(id int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_SetID).Call(m.Instance(), uintptr(id))
}

// Returns the group id of this View, or -1 if not set.
func (m *TCEFViewComponent) GetGroupID() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.ViewComponent_GetGroupID).Call(m.Instance())
	return int32(r1)
}

// A group id is used to tag Views which are part of the same logical group.
// Focus can be moved between views with the same group using the arrow keys.
// The group id is immutable once it's set.
func (m *TCEFViewComponent) SetGroupID(groupId int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_SetGroupID).Call(m.Instance(), uintptr(groupId))
}

// Returns the bounds (size and position) of this View in parent coordinates,
// or DIP screen coordinates if there is no parent.
func (m *TCEFViewComponent) GetBounds() (bounds TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_GetBounds).Call(m.Instance(), uintptr(unsafe.Pointer(&bounds)))
	return
}

// Sets the bounds (size and position) of this View. |bounds| is in parent
// coordinates, or DIP screen coordinates if there is no parent.
func (m *TCEFViewComponent) SetBounds(bounds TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_SetBounds).Call(m.Instance(), uintptr(unsafe.Pointer(&bounds)))
}

// Returns the size of this View in parent coordinates, or DIP screen
// coordinates if there is no parent.
func (m *TCEFViewComponent) GetSize() (size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_GetSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
	return
}

// Sets the size of this View without changing the position. |size| in parent
// coordinates, or DIP screen coordinates if there is no parent.
func (m *TCEFViewComponent) SetSize(size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_SetSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
}

// Returns the position of this View. Position is in parent coordinates, or
// DIP screen coordinates if there is no parent.
func (m *TCEFViewComponent) GetPosition() (point TCefPoint) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_GetPosition).Call(m.Instance(), uintptr(unsafe.Pointer(&point)))
	return
}

// Sets the position of this View without changing the size. |position| is in
// parent coordinates, or DIP screen coordinates if there is no parent.
func (m *TCEFViewComponent) SetPosition(position TCefPoint) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ViewComponent_SetPosition).Call(m.Instance(), uintptr(unsafe.Pointer(&position)))
}

// Returns the type of this View as a string. Used primarily for testing
// purposes.
func (m *TCEFViewComponent) GetTypeString() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.ViewComponent_GetTypeString).Call(m.Instance())
	return api.GoStr(r1)
}

// Returns the height necessary to display this View with the provided width.
func (m *TCEFViewComponent) GetHeightForWidth(width int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.ViewComponent_GetHeightForWidth).Call(m.Instance(), uintptr(width))
	return int32(r1)
}

// Return the preferred size for |view|. The Layout will use this information
// to determine the display size.
func (m *TCEFViewComponent) SetOnGetPreferredSize(fn viewOnGetPreferredSize) {
	imports.Proc(def.ViewComponent_SetOnGetPreferredSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Return the minimum size for |view|.
func (m *TCEFViewComponent) SetOnGetMinimumSize(fn viewOnGetMinimumSize) {
	imports.Proc(def.ViewComponent_SetOnGetMinimumSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Return the maximum size for |view|.
func (m *TCEFViewComponent) SetOnGetMaximumSize(fn viewOnGetMaximumSize) {
	imports.Proc(def.ViewComponent_SetOnGetMaximumSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Return the height necessary to display |view| with the provided |width|.
// If not specified the result of get_preferred_size().height will be used by
// default. Override if |view|'s preferred height depends upon the width (for
// example, with Labels).
func (m *TCEFViewComponent) SetOnGetHeightForWidth(fn viewOnGetHeightForWidth) {
	imports.Proc(def.ViewComponent_SetOnGetHeightForWidth).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Called when the parent of |view| has changed. If |view| is being added to
// |parent| then |added| will be true (1). If |view| is being removed from
// |parent| then |added| will be false (0). If |view| is being reparented the
// remove notification will be sent before the add notification. Do not
// modify the view hierarchy in this callback.
func (m *TCEFViewComponent) SetOnParentViewChanged(fn viewOnParentViewChanged) {
	imports.Proc(def.ViewComponent_SetOnParentViewChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Called when a child of |view| has changed. If |child| is being added to
// |view| then |added| will be true (1). If |child| is being removed from
// |view| then |added| will be false (0). If |child| is being reparented the
// remove notification will be sent to the old parent before the add
// notification is sent to the new parent. Do not modify the view hierarchy
// in this callback.
func (m *TCEFViewComponent) SetOnChildViewChanged(fn viewOnChildViewChanged) {
	imports.Proc(def.ViewComponent_SetOnChildViewChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Called when |view| is added or removed from the ICefWindow.
func (m *TCEFViewComponent) SetOnWindowChanged(fn viewOnWindowChanged) {
	imports.Proc(def.ViewComponent_SetOnWindowChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Called when the layout of |view| has changed.
func (m *TCEFViewComponent) SetOnLayoutChanged(fn viewOnLayoutChanged) {
	imports.Proc(def.ViewComponent_SetOnLayoutChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Called when |view| gains focus.
func (m *TCEFViewComponent) SetOnFocus(fn viewOnFocus) {
	imports.Proc(def.ViewComponent_SetOnFocus).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Called when |view| loses focus.
func (m *TCEFViewComponent) SetOnBlur(fn viewOnBlur) {
	imports.Proc(def.ViewComponent_SetOnBlur).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// <para>Called when the theme for |view| has changed, after the new theme colors
// have already been applied. Views are notified via the component hierarchy
// in depth-first reverse order (children before parents).</para>
// <para>This will be called in the following cases:</para>
// <code>
//  1. When |view|, or a parent of |view|, is added to a Window.
//  2. When the native/OS or Chrome theme changes for the Window that contains
//     |view|. See ICefWindowDelegate.OnThemeColorsChanged documentation.
//  3. When the client explicitly calls ICefWindow.ThemeChanged on the
//     Window that contains |view|.
//
// </code>
// <para>Optionally use this callback to override the new per-View theme colors by
// calling ICefView.SetBackgroundColor or the appropriate component-
// specific function. See ICefWindow.SetThemeColor documentation for how
// to customize additional Window theme colors.</para>
func (m *TCEFViewComponent) SetOnThemeChanged(fn viewOnThemeChanged) {
	imports.Proc(def.ViewComponent_SetOnThemeChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}
