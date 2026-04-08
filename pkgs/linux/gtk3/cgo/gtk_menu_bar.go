package cgo

/*
#cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
#include <gio/gio.h>
#include <gtk/gtk.h>
#include "gtk.go.h"
*/
import "C"
import (
	. "github.com/energye/energy/v3/pkgs/linux/gtk3/types"
	"unsafe"
)

// MenuBar is a representation of GTK's GtkMenuBar.
type MenuBar struct {
	MenuShell
}

func AsMenuBar(p unsafe.Pointer) IMenuBar {
	m := new(MenuBar)
	m.Object = ToGoObject(p)
	return m
}

// native() returns a pointer to the underlying GtkMenuBar.
func (v *MenuBar) native() *C.GtkMenuBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenuBar(p)
}

func wrapMenuBar(obj *Object) *MenuBar {
	if obj == nil {
		return nil
	}

	return &MenuBar{MenuShell{Container{Widget{InitiallyUnowned{obj}}}}}
}

// NewMenuBar is a wrapper around gtk_menu_bar_new().
func NewMenuBar() IMenuBar {
	c := C.gtk_menu_bar_new()
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapMenuBar(obj)
}

// ToggleButton is a representation of GTK's GtkToggleButton.
type ToggleButton struct {
	Button
}

// native returns a pointer to the underlying GtkToggleButton.
func (v *ToggleButton) native() *C.GtkToggleButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkToggleButton(p)
}

// GetActive is a wrapper around gtk_toggle_button_get_active().
func (v *ToggleButton) GetActive() bool {
	c := C.gtk_toggle_button_get_active(v.native())
	return GoBool(c)
}

// SetActive is a wrapper around gtk_toggle_button_set_active().
func (v *ToggleButton) SetActive(isActive bool) {
	C.gtk_toggle_button_set_active(v.native(), CBool(isActive))
}

// GetMode is a wrapper around gtk_toggle_button_get_mode().
func (v *ToggleButton) GetMode() bool {
	c := C.gtk_toggle_button_get_mode(v.native())
	return GoBool(c)
}

// SetMode is a wrapper around gtk_toggle_button_set_mode().
func (v *ToggleButton) SetMode(drawIndicator bool) {
	C.gtk_toggle_button_set_mode(v.native(), CBool(drawIndicator))
}

// Toggled is a wrapper around gtk_toggle_button_toggled().
func (v *ToggleButton) Toggled() {
	C.gtk_toggle_button_toggled(v.native())
}

// GetInconsistent gtk_toggle_button_get_inconsistent().
func (v *ToggleButton) GetInconsistent() bool {
	c := C.gtk_toggle_button_get_inconsistent(v.native())
	return GoBool(c)
}

// SetInconsistent gtk_toggle_button_set_inconsistent().
func (v *ToggleButton) SetInconsistent(setting bool) {
	C.gtk_toggle_button_set_inconsistent(v.native(), CBool(setting))
}

// MenuButton is a representation of GTK's GtkMenuButton.
type MenuButton struct {
	ToggleButton
}

// native returns a pointer to the underlying GtkMenuButton.
func (v *MenuButton) native() *C.GtkMenuButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenuButton(p)
}

// SetPopup is a wrapper around gtk_menu_button_set_popup().
func (v *MenuButton) SetPopup(menu IMenu) {
	C.gtk_menu_button_set_popup(v.native(), menu.toWidget())
}

// GetPopup is a wrapper around gtk_menu_button_get_popup().
func (v *MenuButton) GetPopup() *Menu {
	c := C.gtk_menu_button_get_popup(v.native())
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapMenu(obj)
}

// SetMenuModel is a wrapper around gtk_menu_button_set_menu_model().
func (v *MenuButton) SetMenuModel(menuModel *GMenuModel) {
	C.gtk_menu_button_set_menu_model(v.native(), C.toGMenuModel(unsafe.Pointer(menuModel.Native())))
}

// GetMenuModel is a wrapper around gtk_menu_button_get_menu_model().
func (v *MenuButton) GetMenuModel() *GMenuModel {
	c := C.gtk_menu_button_get_menu_model(v.native())
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return &GMenuModel{obj}
}

// SetDirection is a wrapper around gtk_menu_button_set_direction().
func (v *MenuButton) SetDirection(direction ArrowType) {
	C.gtk_menu_button_set_direction(v.native(), C.GtkArrowType(direction))
}

// GetDirection is a wrapper around gtk_menu_button_get_direction().
func (v *MenuButton) GetDirection() ArrowType {
	c := C.gtk_menu_button_get_direction(v.native())
	return ArrowType(c)
}

// SetAlignWidget is a wrapper around gtk_menu_button_set_align_widget().
func (v *MenuButton) SetAlignWidget(alignWidget IWidget) {
	C.gtk_menu_button_set_align_widget(v.native(), alignWidget.(_IWidget).toWidget())
}

// GetAlignWidget is a wrapper around gtk_menu_button_get_align_widget().
func (v *MenuButton) GetAlignWidget() IWidget {
	c := C.gtk_menu_button_get_align_widget(v.native())
	if c == nil {
		return nil
	}
	return castWidget(c)
}

// MenuItem is a representation of GTK's GtkMenuItem.
type MenuItem struct {
	Bin
}

// IMenuItem is an interface type implemented by all structs
// embedding a MenuItem.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkMenuItem.
type IMenuItem interface {
	toMenuItem() *C.GtkMenuItem
	toWidget() *C.GtkWidget
}

// native returns a pointer to the underlying GtkMenuItem.
func (v *MenuItem) native() *C.GtkMenuItem {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenuItem(p)
}

func (v *MenuItem) toMenuItem() *C.GtkMenuItem {
	if v == nil {
		return nil
	}
	return v.native()
}

func wrapMenuItem(obj *Object) *MenuItem {
	if obj == nil {
		return nil
	}

	return &MenuItem{Bin{Container{Widget{InitiallyUnowned{obj}}}}}
}

// NewMenuItem is a wrapper around gtk_menu_item_new().
func NewMenuItem() (*MenuItem, error) {
	c := C.gtk_menu_item_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapMenuItem(obj), nil
}

// MenuItemNewWithLabel() is a wrapper around gtk_menu_item_new_with_label().
func MenuItemNewWithLabel(label string) (*MenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_menu_item_new_with_label((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapMenuItem(obj), nil
}

// MenuItemNewWithMnemonic() is a wrapper around gtk_menu_item_new_with_mnemonic().
func MenuItemNewWithMnemonic(label string) (*MenuItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_menu_item_new_with_mnemonic((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapMenuItem(obj), nil
}

// SetSubmenu() is a wrapper around gtk_menu_item_set_submenu().
func (v *MenuItem) SetSubmenu(submenu IWidget) {
	C.gtk_menu_item_set_submenu(v.native(), submenu.(_IWidget).toWidget())
}

// GetSubmenu is a wrapper around gtk_menu_item_get_submenu().
func (v *MenuItem) GetSubmenu() (IMenu, error) {
	c := C.gtk_menu_item_get_submenu(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapMenu(obj), nil
}

// SetLabel is a wrapper around gtk_menu_item_set_label().
func (v *MenuItem) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_menu_item_set_label(v.native(), (*C.gchar)(cstr))
}

// GetLabel is a wrapper around gtk_menu_item_get_label().
func (v *MenuItem) GetLabel() string {
	l := C.gtk_menu_item_get_label(v.native())
	return GoString(l)
}

// SetUseUnderline() is a wrapper around gtk_menu_item_set_use_underline()
func (v *MenuItem) SetUseUnderline(settings bool) {
	C.gtk_menu_item_set_use_underline(v.native(), CBool(settings))
}

// GetUseUnderline() is a wrapper around gtk_menu_item_get_use_underline()
func (v *MenuItem) GetUseUnderline() bool {
	c := C.gtk_menu_item_get_use_underline(v.native())
	return GoBool(c)
}

// Select is a wrapper around gtk_menu_item_select()
func (v *MenuItem) Select() {
	C.gtk_menu_item_select(v.native())
}

// Deselect is a wrapper around gtk_menu_item_deselect()
func (v *MenuItem) Deselect() {
	C.gtk_menu_item_deselect(v.native())
}

// Activate is a wrapper around gtk_menu_item_activate()
func (v *MenuItem) Activate() {
	C.gtk_menu_item_activate(v.native())
}

// ToggleSizeRequest is a wrapper around gtk_menu_item_toggle_size_request()
func (v *MenuItem) ToggleSizeRequest(requisition int) int {
	cint := new(C.gint)
	*cint = C.gint(requisition)
	C.gtk_menu_item_toggle_size_request(v.native(), cint)
	return int(*cint)
}

// ToggleSizeAllocate is a wrapper around gtk_menu_item_toggle_size_allocate()
func (v *MenuItem) ToggleSizeAllocate(allocation int) {
	C.gtk_menu_item_toggle_size_allocate(v.native(), C.gint(allocation))
}

// GetReserveIndicator is a wrapper around gtk_menu_item_get_reserve_indicator().
func (v *MenuItem) GetReserveIndicator() bool {
	return GoBool(C.gtk_menu_item_get_reserve_indicator(v.native()))
}

// SetReserveIndicator is a wrapper around gtk_menu_item_set_reserve_indicator().
func (v *MenuItem) SetReserveIndicator(reserve bool) {
	C.gtk_menu_item_set_reserve_indicator(v.native(), CBool(reserve))
}
