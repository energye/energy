package cgo

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

// Predefined attribute names for GMenu
var (
	MENU_ATTRIBUTE_ACTION           string = C.G_MENU_ATTRIBUTE_ACTION
	MENU_ATTRIBUTE_ACTION_NAMESPACE string = C.G_MENU_ATTRIBUTE_ACTION_NAMESPACE
	MENU_ATTRIBUTE_TARGET           string = C.G_MENU_ATTRIBUTE_TARGET
	MENU_ATTRIBUTE_LABEL            string = C.G_MENU_ATTRIBUTE_LABEL
	MENU_ATTRIBUTE_ICON             string = C.G_MENU_ATTRIBUTE_ICON
)

// Predefined link names for GMenu
var (
	MENU_LINK_SECTION string = C.G_MENU_LINK_SECTION
	MENU_LINK_SUBMENU string = C.G_MENU_LINK_SUBMENU
)

// GMenuModel is a representation of GMenuModel.
type GMenuModel struct {
	*Object
}

// native returns a pointer to the underlying GMenuModel.
func (v *GMenuModel) native() *C.GMenuModel {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGMenuModel(unsafe.Pointer(v.GObject))
}

// Native returns a pointer to the underlying GMenuModel.
func (v *GMenuModel) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func wrapMenuModel(obj *Object) *GMenuModel {
	return &GMenuModel{obj}
}

// IsMutable is a wrapper around g_menu_model_is_mutable().
func (v *GMenuModel) IsMutable() bool {
	return GoBool(C.g_menu_model_is_mutable(v.native()))
}

// GetNItems is a wrapper around g_menu_model_get_n_items().
func (v *GMenuModel) GetNItems() int {
	return int(C.g_menu_model_get_n_items(v.native()))
}

// GetItemLink is a wrapper around g_menu_model_get_item_link().
func (v *GMenuModel) GetItemLink(index int, link string) *GMenuModel {
	cstr := (*C.gchar)(C.CString(link))
	defer C.free(unsafe.Pointer(cstr))
	c := C.g_menu_model_get_item_link(v.native(), C.gint(index), cstr)
	if c == nil {
		return nil
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapMenuModel(obj)
}

// ItemsChanged is a wrapper around g_menu_model_items_changed().
func (v *GMenuModel) ItemsChanged(position, removed, added int) {
	C.g_menu_model_items_changed(v.native(), C.gint(position), C.gint(removed), C.gint(added))
}
