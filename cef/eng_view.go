//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICefView Parent: ICefBaseRefCounted
//
//	A View is a rectangle within the views View hierarchy. It is the base
//	interface for all Views. All size and position values are in density
//	independent pixels (DIP) unless otherwise indicated. Methods must be called
//	on the browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_view_capi.h">CEF source file: /include/capi/views/cef_view_capi.h (cef_view_t)</a>
type ICefView interface {
	ICefBaseRefCounted
	// AsBrowserView
	//  Returns this View as a BrowserView or NULL if this is not a BrowserView.
	AsBrowserView() ICefBrowserView // function
	// AsButton
	//  Returns this View as a Button or NULL if this is not a Button.
	AsButton() ICefButton // function
	// AsPanel
	//  Returns this View as a Panel or NULL if this is not a Panel.
	AsPanel() ICefPanel // function
	// AsScrollView
	//  Returns this View as a ScrollView or NULL if this is not a ScrollView.
	AsScrollView() ICefScrollView // function
	// AsTextfield
	//  Returns this View as a Textfield or NULL if this is not a Textfield.
	AsTextfield() ICefTextfield // function
	// GetTypeString
	//  Returns the type of this View as a string. Used primarily for testing
	//  purposes.
	GetTypeString() string // function
	// ToStringEx
	//  Returns a string representation of this View which includes the type and
	//  various type-specific identifying attributes. If |include_children| is
	//  true(1) any child Views will also be included. Used primarily for testing
	//  purposes.
	ToStringEx(includechildren bool) string // function
	// IsValid
	//  Returns true(1) if this View is valid.
	IsValid() bool // function
	// IsAttached
	//  Returns true(1) if this View is currently attached to another View. A
	//  View can only be attached to one View at a time.
	IsAttached() bool // function
	// IsSame
	//  Returns true(1) if this View is the same as |that| View.
	IsSame(that ICefView) bool // function
	// GetDelegate
	//  Returns the delegate associated with this View, if any.
	GetDelegate() ICefViewDelegate // function
	// GetWindow
	//  Returns the top-level Window hosting this View, if any.
	GetWindow() ICefWindow // function
	// GetID
	//  Returns the ID for this View.
	GetID() int32 // function
	// GetGroupID
	//  Returns the group id of this View, or -1 if not set.
	GetGroupID() int32 // function
	// GetParentView
	//  Returns the View that contains this View, if any.
	GetParentView() ICefView // function
	// GetViewForID
	//  Recursively descends the view tree starting at this View, and returns the
	//  first child that it encounters with the given ID. Returns NULL if no
	//  matching child view is found.
	GetViewForID(id int32) ICefView // function
	// GetBounds
	//  Returns the bounds(size and position) of this View in parent coordinates,
	//  or DIP screen coordinates if there is no parent.
	GetBounds() (resultCefRect TCefRect) // function
	// GetBoundsInScreen
	//  Returns the bounds(size and position) of this View in DIP screen
	//  coordinates.
	GetBoundsInScreen() (resultCefRect TCefRect) // function
	// GetSize
	//  Returns the size of this View in parent coordinates, or DIP screen
	//  coordinates if there is no parent.
	GetSize() (resultCefSize TCefSize) // function
	// GetPosition
	//  Returns the position of this View. Position is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	GetPosition() (resultCefPoint TCefPoint) // function
	// GetInsets
	//  Returns the insets for this View in parent coordinates, or DIP screen
	//  coordinates if there is no parent.
	GetInsets() (resultCefInsets TCefInsets) // function
	// GetPreferredSize
	//  Returns the size this View would like to be if enough space is available.
	//  Size is in parent coordinates, or DIP screen coordinates if there is no
	//  parent.
	GetPreferredSize() (resultCefSize TCefSize) // function
	// GetMinimumSize
	//  Returns the minimum size for this View. Size is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	GetMinimumSize() (resultCefSize TCefSize) // function
	// GetMaximumSize
	//  Returns the maximum size for this View. Size is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	GetMaximumSize() (resultCefSize TCefSize) // function
	// GetHeightForWidth
	//  Returns the height necessary to display this View with the provided width.
	GetHeightForWidth(width int32) int32 // function
	// IsVisible
	//  Returns whether this View is visible. A view may be visible but still not
	//  drawn in a Window if any parent views are hidden. If this View is a Window
	//  then a return value of true(1) indicates that this Window is currently
	//  visible to the user on-screen. If this View is not a Window then call
	//  is_drawn() to determine whether this View and all parent views are visible
	//  and will be drawn.
	IsVisible() bool // function
	// IsDrawn
	//  Returns whether this View is visible and drawn in a Window. A view is
	//  drawn if it and all parent views are visible. If this View is a Window
	//  then calling this function is equivalent to calling is_visible().
	//  Otherwise, to determine if the containing Window is visible to the user
	//  on-screen call is_visible() on the Window.
	IsDrawn() bool // function
	// IsEnabled
	//  Returns whether this View is enabled.
	IsEnabled() bool // function
	// IsFocusable
	//  Returns true(1) if this View is focusable, enabled and drawn.
	IsFocusable() bool // function
	// IsAccessibilityFocusable
	//  Return whether this View is focusable when the user requires full keyboard
	//  access, even though it may not be normally focusable.
	IsAccessibilityFocusable() bool // function
	// GetBackgroundColor
	//  Returns the background color for this View.
	GetBackgroundColor() TCefColor // function
	// ConvertPointToScreen
	//  Convert |point| from this View's coordinate system to DIP screen
	//  coordinates. This View must belong to a Window when calling this function.
	//  Returns true(1) if the conversion is successful or false(0) otherwise.
	//  Use ICefDisplay.ConvertPointToPixels() after calling this function
	//  if further conversion to display-specific pixel coordinates is desired.
	ConvertPointToScreen(point *TCefPoint) bool // function
	// ConvertPointFromScreen
	//  Convert |point| to this View's coordinate system from DIP screen
	//  coordinates. This View must belong to a Window when calling this function.
	//  Returns true(1) if the conversion is successful or false(0) otherwise.
	//  Use ICefDisplay.ConvertPointFromPixels() before calling this
	//  function if conversion from display-specific pixel coordinates is
	//  necessary.
	ConvertPointFromScreen(point *TCefPoint) bool // function
	// ConvertPointToWindow
	//  Convert |point| from this View's coordinate system to that of the Window.
	//  This View must belong to a Window when calling this function. Returns true
	// (1) if the conversion is successful or false(0) otherwise.
	ConvertPointToWindow(point *TCefPoint) bool // function
	// ConvertPointFromWindow
	//  Convert |point| to this View's coordinate system from that of the Window.
	//  This View must belong to a Window when calling this function. Returns true
	// (1) if the conversion is successful or false(0) otherwise.
	ConvertPointFromWindow(point *TCefPoint) bool // function
	// ConvertPointToView
	//  Convert |point| from this View's coordinate system to that of |view|.
	//  |view| needs to be in the same Window but not necessarily the same view
	//  hierarchy. Returns true(1) if the conversion is successful or false(0)
	//  otherwise.
	ConvertPointToView(view ICefView, point *TCefPoint) bool // function
	// ConvertPointFromView
	//  Convert |point| to this View's coordinate system from that |view|. |view|
	//  needs to be in the same Window but not necessarily the same view
	//  hierarchy. Returns true(1) if the conversion is successful or false(0)
	//  otherwise.
	ConvertPointFromView(view ICefView, point *TCefPoint) bool // function
	// SetID
	//  Sets the ID for this View. ID should be unique within the subtree that you
	//  intend to search for it. 0 is the default ID for views.
	SetID(id int32) // procedure
	// SetGroupID
	//  A group id is used to tag Views which are part of the same logical group.
	//  Focus can be moved between views with the same group using the arrow keys.
	//  The group id is immutable once it's set.
	SetGroupID(groupid int32) // procedure
	// SetBounds
	//  Sets the bounds(size and position) of this View. |bounds| is in parent
	//  coordinates, or DIP screen coordinates if there is no parent.
	SetBounds(bounds *TCefRect) // procedure
	// SetSize
	//  Sets the size of this View without changing the position. |size| in parent
	//  coordinates, or DIP screen coordinates if there is no parent.
	SetSize(size *TCefSize) // procedure
	// SetPosition
	//  Sets the position of this View without changing the size. |position| is in
	//  parent coordinates, or DIP screen coordinates if there is no parent.
	SetPosition(position *TCefPoint) // procedure
	// SetInsets
	//  Sets the insets for this View. |insets| is in parent coordinates, or DIP
	//  screen coordinates if there is no parent.
	SetInsets(insets *TCefInsets) // procedure
	// SizeToPreferredSize
	//  Size this View to its preferred size. Size is in parent coordinates, or
	//  DIP screen coordinates if there is no parent.
	SizeToPreferredSize() // procedure
	// InvalidateLayout
	//  Indicate that this View and all parent Views require a re-layout. This
	//  ensures the next call to layout() will propagate to this View even if the
	//  bounds of parent Views do not change.
	InvalidateLayout() // procedure
	// SetVisible
	//  Sets whether this View is visible. Windows are hidden by default and other
	//  views are visible by default. This View and any parent views must be set
	//  as visible for this View to be drawn in a Window. If this View is set as
	//  hidden then it and any child views will not be drawn and, if any of those
	//  views currently have focus, then focus will also be cleared. Painting is
	//  scheduled as needed. If this View is a Window then calling this function
	//  is equivalent to calling the Window show() and hide() functions.
	SetVisible(visible bool) // procedure
	// SetEnabled
	//  Set whether this View is enabled. A disabled View does not receive
	//  keyboard or mouse inputs. If |enabled| differs from the current value the
	//  View will be repainted. Also, clears focus if the focused View is
	//  disabled.
	SetEnabled(enabled bool) // procedure
	// SetFocusable
	//  Sets whether this View is capable of taking focus. It will clear focus if
	//  the focused View is set to be non-focusable. This is false(0) by default
	//  so that a View used as a container does not get the focus.
	SetFocusable(focusable bool) // procedure
	// RequestFocus
	//  Request keyboard focus. If this View is focusable it will become the
	//  focused View.
	RequestFocus() // procedure
	// SetBackgroundColor
	//  Sets the background color for this View.
	SetBackgroundColor(color TCefColor) // procedure
}

// TCefView Parent: TCefBaseRefCounted
//
//	A View is a rectangle within the views View hierarchy. It is the base
//	interface for all Views. All size and position values are in density
//	independent pixels (DIP) unless otherwise indicated. Methods must be called
//	on the browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_view_capi.h">CEF source file: /include/capi/views/cef_view_capi.h (cef_view_t)</a>
type TCefView struct {
	TCefBaseRefCounted
}

// ViewRef -> ICefView
var ViewRef view

// view TCefView Ref
type view uintptr

// UnWrap
//
//	Returns a ICefView instance using a PCefView data pointer.
func (m *view) UnWrap(data uintptr) ICefView {
	var resultCefView uintptr
	CEF().SysCallN(1571, uintptr(data), uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCefView) AsBrowserView() ICefBrowserView {
	var resultCefBrowserView uintptr
	CEF().SysCallN(1521, m.Instance(), uintptr(unsafePointer(&resultCefBrowserView)))
	return AsCefBrowserView(resultCefBrowserView)
}

func (m *TCefView) AsButton() ICefButton {
	var resultCefButton uintptr
	CEF().SysCallN(1522, m.Instance(), uintptr(unsafePointer(&resultCefButton)))
	return AsCefButton(resultCefButton)
}

func (m *TCefView) AsPanel() ICefPanel {
	var resultCefPanel uintptr
	CEF().SysCallN(1523, m.Instance(), uintptr(unsafePointer(&resultCefPanel)))
	return AsCefPanel(resultCefPanel)
}

func (m *TCefView) AsScrollView() ICefScrollView {
	var resultCefScrollView uintptr
	CEF().SysCallN(1524, m.Instance(), uintptr(unsafePointer(&resultCefScrollView)))
	return AsCefScrollView(resultCefScrollView)
}

func (m *TCefView) AsTextfield() ICefTextfield {
	var resultCefTextfield uintptr
	CEF().SysCallN(1525, m.Instance(), uintptr(unsafePointer(&resultCefTextfield)))
	return AsCefTextfield(resultCefTextfield)
}

func (m *TCefView) GetTypeString() string {
	r1 := CEF().SysCallN(1546, m.Instance())
	return GoStr(r1)
}

func (m *TCefView) ToStringEx(includechildren bool) string {
	r1 := CEF().SysCallN(1570, m.Instance(), PascalBool(includechildren))
	return GoStr(r1)
}

func (m *TCefView) IsValid() bool {
	r1 := CEF().SysCallN(1556, m.Instance())
	return GoBool(r1)
}

func (m *TCefView) IsAttached() bool {
	r1 := CEF().SysCallN(1551, m.Instance())
	return GoBool(r1)
}

func (m *TCefView) IsSame(that ICefView) bool {
	r1 := CEF().SysCallN(1555, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefView) GetDelegate() ICefViewDelegate {
	var resultCefViewDelegate uintptr
	CEF().SysCallN(1535, m.Instance(), uintptr(unsafePointer(&resultCefViewDelegate)))
	return AsCefViewDelegate(resultCefViewDelegate)
}

func (m *TCefView) GetWindow() ICefWindow {
	var resultCefWindow uintptr
	CEF().SysCallN(1548, m.Instance(), uintptr(unsafePointer(&resultCefWindow)))
	return AsCefWindow(resultCefWindow)
}

func (m *TCefView) GetID() int32 {
	r1 := CEF().SysCallN(1538, m.Instance())
	return int32(r1)
}

func (m *TCefView) GetGroupID() int32 {
	r1 := CEF().SysCallN(1536, m.Instance())
	return int32(r1)
}

func (m *TCefView) GetParentView() ICefView {
	var resultCefView uintptr
	CEF().SysCallN(1542, m.Instance(), uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCefView) GetViewForID(id int32) ICefView {
	var resultCefView uintptr
	CEF().SysCallN(1547, m.Instance(), uintptr(id), uintptr(unsafePointer(&resultCefView)))
	return AsCefView(resultCefView)
}

func (m *TCefView) GetBounds() (resultCefRect TCefRect) {
	CEF().SysCallN(1533, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCefView) GetBoundsInScreen() (resultCefRect TCefRect) {
	CEF().SysCallN(1534, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCefView) GetSize() (resultCefSize TCefSize) {
	CEF().SysCallN(1545, m.Instance(), uintptr(unsafePointer(&resultCefSize)))
	return
}

func (m *TCefView) GetPosition() (resultCefPoint TCefPoint) {
	CEF().SysCallN(1543, m.Instance(), uintptr(unsafePointer(&resultCefPoint)))
	return
}

func (m *TCefView) GetInsets() (resultCefInsets TCefInsets) {
	CEF().SysCallN(1539, m.Instance(), uintptr(unsafePointer(&resultCefInsets)))
	return
}

func (m *TCefView) GetPreferredSize() (resultCefSize TCefSize) {
	CEF().SysCallN(1544, m.Instance(), uintptr(unsafePointer(&resultCefSize)))
	return
}

func (m *TCefView) GetMinimumSize() (resultCefSize TCefSize) {
	CEF().SysCallN(1541, m.Instance(), uintptr(unsafePointer(&resultCefSize)))
	return
}

func (m *TCefView) GetMaximumSize() (resultCefSize TCefSize) {
	CEF().SysCallN(1540, m.Instance(), uintptr(unsafePointer(&resultCefSize)))
	return
}

func (m *TCefView) GetHeightForWidth(width int32) int32 {
	r1 := CEF().SysCallN(1537, m.Instance(), uintptr(width))
	return int32(r1)
}

func (m *TCefView) IsVisible() bool {
	r1 := CEF().SysCallN(1557, m.Instance())
	return GoBool(r1)
}

func (m *TCefView) IsDrawn() bool {
	r1 := CEF().SysCallN(1552, m.Instance())
	return GoBool(r1)
}

func (m *TCefView) IsEnabled() bool {
	r1 := CEF().SysCallN(1553, m.Instance())
	return GoBool(r1)
}

func (m *TCefView) IsFocusable() bool {
	r1 := CEF().SysCallN(1554, m.Instance())
	return GoBool(r1)
}

func (m *TCefView) IsAccessibilityFocusable() bool {
	r1 := CEF().SysCallN(1550, m.Instance())
	return GoBool(r1)
}

func (m *TCefView) GetBackgroundColor() TCefColor {
	r1 := CEF().SysCallN(1532, m.Instance())
	return TCefColor(r1)
}

func (m *TCefView) ConvertPointToScreen(point *TCefPoint) bool {
	var result0 uintptr
	r1 := CEF().SysCallN(1529, m.Instance(), uintptr(unsafePointer(&result0)))
	*point = *(*TCefPoint)(unsafePointer(result0))
	return GoBool(r1)
}

func (m *TCefView) ConvertPointFromScreen(point *TCefPoint) bool {
	var result0 uintptr
	r1 := CEF().SysCallN(1526, m.Instance(), uintptr(unsafePointer(&result0)))
	*point = *(*TCefPoint)(unsafePointer(result0))
	return GoBool(r1)
}

func (m *TCefView) ConvertPointToWindow(point *TCefPoint) bool {
	var result0 uintptr
	r1 := CEF().SysCallN(1531, m.Instance(), uintptr(unsafePointer(&result0)))
	*point = *(*TCefPoint)(unsafePointer(result0))
	return GoBool(r1)
}

func (m *TCefView) ConvertPointFromWindow(point *TCefPoint) bool {
	var result0 uintptr
	r1 := CEF().SysCallN(1528, m.Instance(), uintptr(unsafePointer(&result0)))
	*point = *(*TCefPoint)(unsafePointer(result0))
	return GoBool(r1)
}

func (m *TCefView) ConvertPointToView(view ICefView, point *TCefPoint) bool {
	var result1 uintptr
	r1 := CEF().SysCallN(1530, m.Instance(), GetObjectUintptr(view), uintptr(unsafePointer(&result1)))
	*point = *(*TCefPoint)(unsafePointer(result1))
	return GoBool(r1)
}

func (m *TCefView) ConvertPointFromView(view ICefView, point *TCefPoint) bool {
	var result1 uintptr
	r1 := CEF().SysCallN(1527, m.Instance(), GetObjectUintptr(view), uintptr(unsafePointer(&result1)))
	*point = *(*TCefPoint)(unsafePointer(result1))
	return GoBool(r1)
}

func (m *TCefView) SetID(id int32) {
	CEF().SysCallN(1564, m.Instance(), uintptr(id))
}

func (m *TCefView) SetGroupID(groupid int32) {
	CEF().SysCallN(1563, m.Instance(), uintptr(groupid))
}

func (m *TCefView) SetBounds(bounds *TCefRect) {
	CEF().SysCallN(1560, m.Instance(), uintptr(unsafePointer(bounds)))
}

func (m *TCefView) SetSize(size *TCefSize) {
	CEF().SysCallN(1567, m.Instance(), uintptr(unsafePointer(size)))
}

func (m *TCefView) SetPosition(position *TCefPoint) {
	CEF().SysCallN(1566, m.Instance(), uintptr(unsafePointer(position)))
}

func (m *TCefView) SetInsets(insets *TCefInsets) {
	CEF().SysCallN(1565, m.Instance(), uintptr(unsafePointer(insets)))
}

func (m *TCefView) SizeToPreferredSize() {
	CEF().SysCallN(1569, m.Instance())
}

func (m *TCefView) InvalidateLayout() {
	CEF().SysCallN(1549, m.Instance())
}

func (m *TCefView) SetVisible(visible bool) {
	CEF().SysCallN(1568, m.Instance(), PascalBool(visible))
}

func (m *TCefView) SetEnabled(enabled bool) {
	CEF().SysCallN(1561, m.Instance(), PascalBool(enabled))
}

func (m *TCefView) SetFocusable(focusable bool) {
	CEF().SysCallN(1562, m.Instance(), PascalBool(focusable))
}

func (m *TCefView) RequestFocus() {
	CEF().SysCallN(1558, m.Instance())
}

func (m *TCefView) SetBackgroundColor(color TCefColor) {
	CEF().SysCallN(1559, m.Instance(), uintptr(color))
}
