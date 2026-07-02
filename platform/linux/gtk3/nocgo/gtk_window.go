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
	"github.com/energye/energy/v3/platform/linux/callback"
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// Window is a representation of GTK's GtkWindow.
type Window struct {
	Bin
}

func AsWindow(ptr unsafe.Pointer) IWindow {
	if ptr == nil {
		return nil
	}
	m := &Window{}
	m.instance = ptr
	return m
}

// NewWindow is a wrapper around gtk_window_new().
func NewWindow(t WindowType) (*Window, error) {
	r := gtk3.SysCall("gtk_window_new", uintptr(t))
	if r == 0 {
		return nil, errNilPtr
	}
	return &Window{Bin{Container{Widget{Object{instance: unsafe.Pointer(r)}}}}}, nil
}

// ToWindow returns the window itself. This is a convenience method for type conversion.
func (m *Window) ToWindow() *Window {
	return m
}

// SetTitlebar is a wrapper around gtk_window_set_titlebar().
func (m *Window) SetTitlebar(titlebar IWidget) {
	gtk3.SysCall("gtk_window_set_titlebar", m.Instance(), titlebar.Instance())
}

// SetTitle is a wrapper around gtk_window_set_title().
func (m *Window) SetTitle(title string) {
	gtk3.SysCall("gtk_window_set_title", m.Instance(), CStr(title))
}

// GetTitle is a wrapper around gtk_window_get_title().
func (m *Window) GetTitle() string {
	r := gtk3.SysCall("gtk_window_get_title", m.Instance())
	return GoStr(r)
}

// SetResizable is a wrapper around gtk_window_set_resizable().
func (m *Window) SetResizable(resizable bool) {
	gtk3.SysCall("gtk_window_set_resizable", m.Instance(), ToCBool(resizable))
}

// GetResizable is a wrapper around gtk_window_get_resizable().
func (m *Window) GetResizable() bool {
	r := gtk3.SysCall("gtk_window_get_resizable", m.Instance())
	return ToGoBool(r)
}

// ActivateFocus is a wrapper around gtk_window_activate_focus().
func (m *Window) ActivateFocus() bool {
	r := gtk3.SysCall("gtk_window_activate_focus", m.Instance())
	return ToGoBool(r)
}

// ActivateDefault is a wrapper around gtk_window_activate_default().
func (m *Window) ActivateDefault() bool {
	r := gtk3.SysCall("gtk_window_activate_default", m.Instance())
	return ToGoBool(r)
}

// SetModal is a wrapper around gtk_window_set_modal().
func (m *Window) SetModal(modal bool) {
	gtk3.SysCall("gtk_window_set_modal", m.Instance(), ToCBool(modal))
}

// GetModal is a wrapper around gtk_window_get_modal().
func (m *Window) GetModal() bool {
	r := gtk3.SysCall("gtk_window_get_modal", m.Instance())
	return ToGoBool(r)
}

// SetDefaultSize is a wrapper around gtk_window_set_default_size().
func (m *Window) SetDefaultSize(width, height int) {
	gtk3.SysCall("gtk_window_set_default_size", m.Instance(), uintptr(width), uintptr(height))
}

// GetDefaultSize is a wrapper around gtk_window_get_default_size().
func (m *Window) GetDefaultSize() (width, height int) {
	gtk3.SysCall("gtk_window_get_default_size", m.Instance(), uintptr(unsafe.Pointer(&width)), uintptr(unsafe.Pointer(&height)))
	return
}

// SetGravity is a wrapper around gtk_window_set_gravity().
func (m *Window) SetGravity(gravity Gravity) {
	gtk3.SysCall("gtk_window_set_gravity", m.Instance(), uintptr(gravity))
}

// GetGravity is a wrapper around gtk_window_get_gravity().
func (m *Window) GetGravity() Gravity {
	r := gtk3.SysCall("gtk_window_get_gravity", m.Instance())
	return Gravity(r)
}

// SetAttachedTo is a wrapper around gtk_window_set_attached_to().
func (m *Window) SetAttachedTo(attachWidget IWidget) {
	var aW uintptr
	if attachWidget != nil {
		aW = attachWidget.Instance()
	}
	gtk3.SysCall("gtk_window_set_attached_to", m.Instance(), aW)
}

// GetAttachedTo is a wrapper around gtk_window_get_attached_to().
func (m *Window) GetAttachedTo() IWidget {
	r := gtk3.SysCall("gtk_window_get_attached_to", m.Instance())
	if r == 0 {
		return nil
	}
	return AsWidget(unsafe.Pointer(r))
}

// SetDestroyWithParent is a wrapper around gtk_window_set_destroy_with_parent().
func (m *Window) SetDestroyWithParent(setting bool) {
	gtk3.SysCall("gtk_window_set_destroy_with_parent", m.Instance(), ToCBool(setting))
}

// GetDestroyWithParent is a wrapper around gtk_window_get_destroy_with_parent().
func (m *Window) GetDestroyWithParent() bool {
	r := gtk3.SysCall("gtk_window_get_destroy_with_parent", m.Instance())
	return ToGoBool(r)
}

// SetHideTitlebarWhenMaximized is a wrapper around gtk_window_set_hide_titlebar_when_maximized().
func (m *Window) SetHideTitlebarWhenMaximized(setting bool) {
	gtk3.SysCall("gtk_window_set_hide_titlebar_when_maximized", m.Instance(), ToCBool(setting))
}

// GetHideTitlebarWhenMaximized is a wrapper around gtk_window_get_hide_titlebar_when_maximized().
func (m *Window) GetHideTitlebarWhenMaximized() bool {
	r := gtk3.SysCall("gtk_window_get_hide_titlebar_when_maximized", m.Instance())
	return ToGoBool(r)
}

// IsActive is a wrapper around gtk_window_is_active().
func (m *Window) IsActive() bool {
	r := gtk3.SysCall("gtk_window_is_active", m.Instance())
	return ToGoBool(r)
}

// HasToplevelFocus is a wrapper around gtk_window_has_toplevel_focus().
func (m *Window) HasToplevelFocus() bool {
	r := gtk3.SysCall("gtk_window_has_toplevel_focus", m.Instance())
	return ToGoBool(r)
}

// GetFocus is a wrapper around gtk_window_get_focus().
func (m *Window) GetFocus() IWidget {
	r := gtk3.SysCall("gtk_window_get_focus", m.Instance())
	if r == 0 {
		return nil
	}
	return AsWidget(unsafe.Pointer(r))
}

// SetFocus is a wrapper around gtk_window_set_focus().
func (m *Window) SetFocus(w IWidget) {
	var wPtr uintptr
	if w != nil {
		wPtr = w.Instance()
	}
	gtk3.SysCall("gtk_window_set_focus", m.Instance(), wPtr)
}

// GetDefaultWidget is a wrapper around gtk_window_get_default_widget().
func (m *Window) GetDefaultWidget() IWidget {
	r := gtk3.SysCall("gtk_window_get_default_widget", m.Instance())
	if r == 0 {
		return nil
	}
	return AsWidget(unsafe.Pointer(r))
}

// SetDefault is a wrapper around gtk_window_set_default().
func (m *Window) SetDefault(widget IWidget) {
	gtk3.SysCall("gtk_window_set_default", m.Instance(), widget.Instance())
}

// Present is a wrapper around gtk_window_present().
func (m *Window) Present() {
	gtk3.SysCall("gtk_window_present", m.Instance())
}

// PresentWithTime is a wrapper around gtk_window_present_with_time().
func (m *Window) PresentWithTime(ts uint32) {
	gtk3.SysCall("gtk_window_present_with_time", m.Instance(), uintptr(ts))
}

// Iconify is a wrapper around gtk_window_iconify().
func (m *Window) Iconify() {
	gtk3.SysCall("gtk_window_iconify", m.Instance())
}

// Deiconify is a wrapper around gtk_window_deiconify().
func (m *Window) Deiconify() {
	gtk3.SysCall("gtk_window_deiconify", m.Instance())
}

// Stick is a wrapper around gtk_window_stick().
func (m *Window) Stick() {
	gtk3.SysCall("gtk_window_stick", m.Instance())
}

// Unstick is a wrapper around gtk_window_unstick().
func (m *Window) Unstick() {
	gtk3.SysCall("gtk_window_unstick", m.Instance())
}

// Maximize is a wrapper around gtk_window_maximize().
func (m *Window) Maximize() {
	gtk3.SysCall("gtk_window_maximize", m.Instance())
}

// Unmaximize is a wrapper around gtk_window_unmaximize().
func (m *Window) Unmaximize() {
	gtk3.SysCall("gtk_window_unmaximize", m.Instance())
}

// Fullscreen is a wrapper around gtk_window_fullscreen().
func (m *Window) Fullscreen() {
	gtk3.SysCall("gtk_window_fullscreen", m.Instance())
}

// Unfullscreen is a wrapper around gtk_window_unfullscreen().
func (m *Window) Unfullscreen() {
	gtk3.SysCall("gtk_window_unfullscreen", m.Instance())
}

// SetKeepAbove is a wrapper around gtk_window_set_keep_above().
func (m *Window) SetKeepAbove(setting bool) {
	gtk3.SysCall("gtk_window_set_keep_above", m.Instance(), ToCBool(setting))
}

// SetKeepBelow is a wrapper around gtk_window_set_keep_below().
func (m *Window) SetKeepBelow(setting bool) {
	gtk3.SysCall("gtk_window_set_keep_below", m.Instance(), ToCBool(setting))
}

// SetDecorated is a wrapper around gtk_window_set_decorated().
func (m *Window) SetDecorated(setting bool) {
	gtk3.SysCall("gtk_window_set_decorated", m.Instance(), ToCBool(setting))
}

// GetDecorated is a wrapper around gtk_window_get_decorated().
func (m *Window) GetDecorated() bool {
	r := gtk3.SysCall("gtk_window_get_decorated", m.Instance())
	return ToGoBool(r)
}

// SetDeletable is a wrapper around gtk_window_set_deletable().
func (m *Window) SetDeletable(setting bool) {
	gtk3.SysCall("gtk_window_set_deletable", m.Instance(), ToCBool(setting))
}

// GetDeletable is a wrapper around gtk_window_get_deletable().
func (m *Window) GetDeletable() bool {
	r := gtk3.SysCall("gtk_window_get_deletable", m.Instance())
	return ToGoBool(r)
}

// SetTypeHint is a wrapper around gtk_window_set_type_hint().
func (m *Window) SetTypeHint(typeHint WindowTypeHint) {
	gtk3.SysCall("gtk_window_set_type_hint", m.Instance(), uintptr(typeHint))
}

// GetTypeHint is a wrapper around gtk_window_get_type_hint().
func (m *Window) GetTypeHint() WindowTypeHint {
	r := gtk3.SysCall("gtk_window_get_type_hint", m.Instance())
	return WindowTypeHint(r)
}

// SetSkipTaskbarHint is a wrapper around gtk_window_set_skip_taskbar_hint().
func (m *Window) SetSkipTaskbarHint(setting bool) {
	gtk3.SysCall("gtk_window_set_skip_taskbar_hint", m.Instance(), ToCBool(setting))
}

// GetSkipTaskbarHint is a wrapper around gtk_window_get_skip_taskbar_hint().
func (m *Window) GetSkipTaskbarHint() bool {
	r := gtk3.SysCall("gtk_window_get_skip_taskbar_hint", m.Instance())
	return ToGoBool(r)
}

// SetSkipPagerHint is a wrapper around gtk_window_set_skip_pager_hint().
func (m *Window) SetSkipPagerHint(setting bool) {
	gtk3.SysCall("gtk_window_set_skip_pager_hint", m.Instance(), ToCBool(setting))
}

// GetSkipPagerHint is a wrapper around gtk_window_get_skip_pager_hint().
func (m *Window) GetSkipPagerHint() bool {
	r := gtk3.SysCall("gtk_window_get_skip_pager_hint", m.Instance())
	return ToGoBool(r)
}

// SetUrgencyHint is a wrapper around gtk_window_set_urgency_hint().
func (m *Window) SetUrgencyHint(setting bool) {
	gtk3.SysCall("gtk_window_set_urgency_hint", m.Instance(), ToCBool(setting))
}

// GetUrgencyHint is a wrapper around gtk_window_get_urgency_hint().
func (m *Window) GetUrgencyHint() bool {
	r := gtk3.SysCall("gtk_window_get_urgency_hint", m.Instance())
	return ToGoBool(r)
}

// SetAcceptFocus is a wrapper around gtk_window_set_accept_focus().
func (m *Window) SetAcceptFocus(setting bool) {
	gtk3.SysCall("gtk_window_set_accept_focus", m.Instance(), ToCBool(setting))
}

// GetAcceptFocus is a wrapper around gtk_window_get_accept_focus().
func (m *Window) GetAcceptFocus() bool {
	r := gtk3.SysCall("gtk_window_get_accept_focus", m.Instance())
	return ToGoBool(r)
}

// SetFocusOnMap is a wrapper around gtk_window_set_focus_on_map().
func (m *Window) SetFocusOnMap(setting bool) {
	gtk3.SysCall("gtk_window_set_focus_on_map", m.Instance(), ToCBool(setting))
}

// GetFocusOnMap is a wrapper around gtk_window_get_focus_on_map().
func (m *Window) GetFocusOnMap() bool {
	r := gtk3.SysCall("gtk_window_get_focus_on_map", m.Instance())
	return ToGoBool(r)
}

// SetStartupID is a wrapper around gtk_window_set_startup_id().
func (m *Window) SetStartupID(sid string) {
	gtk3.SysCall("gtk_window_set_startup_id", m.Instance(), CStr(sid))
}

// SetRole is a wrapper around gtk_window_set_role().
func (m *Window) SetRole(s string) {
	gtk3.SysCall("gtk_window_set_role", m.Instance(), CStr(s))
}

// GetRole is a wrapper around gtk_window_get_role().
func (m *Window) GetRole() (string, error) {
	r := gtk3.SysCall("gtk_window_get_role", m.Instance())
	return GoStr(r), nil
}

// GetPosition is a wrapper around gtk_window_get_position().
func (m *Window) GetPosition() (int, int) {
	var x, y int
	gtk3.SysCall("gtk_window_get_position", m.Instance(), uintptr(unsafe.Pointer(&x)), uintptr(unsafe.Pointer(&y)))
	return x, y
}

// GetSize is a wrapper around gtk_window_get_size().
func (m *Window) GetSize() (width, height int) {
	gtk3.SysCall("gtk_window_get_size", m.Instance(), uintptr(unsafe.Pointer(&width)), uintptr(unsafe.Pointer(&height)))
	return
}

// GetTransientFor is a wrapper around gtk_window_get_transient_for().
func (m *Window) GetTransientFor() (*Window, error) {
	r := gtk3.SysCall("gtk_window_get_transient_for", m.Instance())
	if r == 0 {
		return nil, nil
	}
	return &Window{Bin{Container{Widget{Object{instance: unsafe.Pointer(r)}}}}}, nil
}

// GetIconName is a wrapper around gtk_window_get_icon_name().
func (m *Window) GetIconName() (string, error) {
	r := gtk3.SysCall("gtk_window_get_icon_name", m.Instance())
	return GoStr(r), nil
}

// SetIconName is a wrapper around gtk_window_set_icon_name().
func (m *Window) SetIconName(name string) {
	gtk3.SysCall("gtk_window_set_icon_name", m.Instance(), CStr(name))
}

// SetIconFromFile is a wrapper around gtk_window_set_icon_from_file().
func (m *Window) SetIconFromFile(file string) error {
	gtk3.SysCall("gtk_window_set_icon_from_file", m.Instance(), CStr(file))
	return nil
}

// GetWindowType is a wrapper around gtk_window_get_window_type().
func (m *Window) GetWindowType() WindowType {
	r := gtk3.SysCall("gtk_window_get_window_type", m.Instance())
	return WindowType(r)
}

// HasGroup is a wrapper around gtk_window_has_group().
func (m *Window) HasGroup() bool {
	r := gtk3.SysCall("gtk_window_has_group", m.Instance())
	return ToGoBool(r)
}

// Move is a wrapper around gtk_window_move().
func (m *Window) Move(x, y int) {
	gtk3.SysCall("gtk_window_move", m.Instance(), uintptr(x), uintptr(y))
}

// Resize is a wrapper around gtk_window_resize().
func (m *Window) Resize(width, height int) {
	gtk3.SysCall("gtk_window_resize", m.Instance(), uintptr(width), uintptr(height))
}

// GetMnemonicsVisible is a wrapper around gtk_window_get_mnemonics_visible().
func (m *Window) GetMnemonicsVisible() bool {
	r := gtk3.SysCall("gtk_window_get_mnemonics_visible", m.Instance())
	return ToGoBool(r)
}

// SetMnemonicsVisible is a wrapper around gtk_window_set_mnemonics_visible().
func (m *Window) SetMnemonicsVisible(setting bool) {
	gtk3.SysCall("gtk_window_set_mnemonics_visible", m.Instance(), ToCBool(setting))
}

// GetFocusVisible is a wrapper around gtk_window_get_focus_visible().
func (m *Window) GetFocusVisible() bool {
	r := gtk3.SysCall("gtk_window_get_focus_visible", m.Instance())
	return ToGoBool(r)
}

// SetFocusVisible is a wrapper around gtk_window_set_focus_visible().
func (m *Window) SetFocusVisible(setting bool) {
	gtk3.SysCall("gtk_window_set_focus_visible", m.Instance(), ToCBool(setting))
}

// AddMnemonic is a wrapper around gtk_window_add_mnemonic().
func (m *Window) AddMnemonic(keyval uint, target IWidget) {
	gtk3.SysCall("gtk_window_add_mnemonic", m.Instance(), uintptr(keyval), target.Instance())
}

// RemoveMnemonic is a wrapper around gtk_window_remove_mnemonic().
func (m *Window) RemoveMnemonic(keyval uint, target IWidget) {
	gtk3.SysCall("gtk_window_remove_mnemonic", m.Instance(), uintptr(keyval), target.Instance())
}

// SetMnemonicModifier is a wrapper around gtk_window_set_mnemonic_modifier().
func (m *Window) SetMnemonicModifier(mods uint) {
	gtk3.SysCall("gtk_window_set_mnemonic_modifier", m.Instance(), uintptr(mods))
}

// BeginResizeDrag is a wrapper around gtk_window_begin_resize_drag().
func (m *Window) BeginResizeDrag(edge WindowEdge, button ButtonType, rootX, rootY int, timestamp uint32) {
	gtk3.SysCall("gtk_window_begin_resize_drag", m.Instance(), uintptr(edge), uintptr(button), uintptr(rootX), uintptr(rootY), uintptr(timestamp))
}

// BeginMoveDrag is a wrapper around gtk_window_begin_move_drag().
func (m *Window) BeginMoveDrag(button ButtonType, rootX, rootY int, timestamp uint32) {
	gtk3.SysCall("gtk_window_begin_move_drag", m.Instance(), uintptr(button), uintptr(rootX), uintptr(rootY), uintptr(timestamp))
}

// SetOnConfigure is a callback for the "configure-event" signal.
func (m *Window) SetOnConfigure(fn TConfigureEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnConfigureEvent, callback.C_trampoline_3_gboolean, fn, 0)
}

// SetOnMap is a callback for the "map" signal.
func (m *Window) SetOnMap(fn TMapEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnMapEvent, callback.C_trampoline_2_void, fn, 0)
}

// SetOnDraw is a callback for the "draw" signal.
func (m *Window) SetOnDraw(fn TDrawEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnDrawEvent, callback.C_trampoline_3_gboolean, fn, 0)
}

// SetOnDestroy is a callback for the "destroy" signal.
func (m *Window) SetOnDestroy(fn TNotifyEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnDestroy, callback.C_trampoline_2_void, fn, 0)
}

// SetPosition is a wrapper around gtk_window_set_position().
func (m *Window) SetPosition(position WindowType) {
	gtk3.SysCall("gtk_window_set_position", m.Instance(), uintptr(position))
}

// SetTransientFor is a wrapper around gtk_window_set_transient_for().
func (m *Window) SetTransientFor(parent IWindow) {
	var p uintptr
	if parent != nil {
		p = parent.Instance()
	}
	gtk3.SysCall("gtk_window_set_transient_for", m.Instance(), p)
}

// WindowGetDefaultIconName is a wrapper around gtk_window_get_default_icon_name().
func WindowGetDefaultIconName() (string, error) {
	r := gtk3.SysCall("gtk_window_get_default_icon_name")
	if r == 0 {
		return "", errNilPtr
	}
	return GoStr(r), nil
}

// WindowSetDefaultIconFromFile is a wrapper around gtk_window_set_default_icon_from_file().
func WindowSetDefaultIconFromFile(file string) error {
	cstr := CStr(file)
	var gErr uintptr
	r := gtk3.SysCall("gtk_window_set_default_icon_from_file", cstr, uintptr(unsafe.Pointer(&gErr)))
	if r == 0 {
		if gErr != 0 {
			gError := (*GError)(unsafe.Pointer(gErr))
			msg := GoStr(gError.Message)
			GErrorFree(gErr)
			return errors.New(msg)
		}
		return errNilPtr
	}
	return nil
}

// WindowSetDefaultIconName is a wrapper around gtk_window_set_default_icon_name().
func WindowSetDefaultIconName(s string) {
	cstr := CStr(s)
	gtk3.SysCall("gtk_window_set_default_icon_name", cstr)
}

// WindowSetAutoStartupNotification is a wrapper around gtk_window_set_auto_startup_notification().
func WindowSetAutoStartupNotification(setting bool) {
	gtk3.SysCall("gtk_window_set_auto_startup_notification", ToCBool(setting))
}
