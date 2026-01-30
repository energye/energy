package gtk3

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import "unsafe"

// Menu is a representation of GTK's GtkMenu.
type Menu struct {
	MenuShell
}

// IMenu is an interface type implemented by all structs embedding
// a Menu.  It is meant to be used as an argument type for wrapper
// functions that wrap around a C GTK function taking a
// GtkMenu.
type IMenu interface {
	toMenu() *C.GtkMenu
	toWidget() *C.GtkWidget
}

// native() returns a pointer to the underlying GtkMenu.
func (v *Menu) native() *C.GtkMenu {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenu(p)
}

func (v *Menu) toMenu() *C.GtkMenu {
	if v == nil {
		return nil
	}
	return v.native()
}

func wrapMenu(obj *Object) *Menu {
	if obj == nil {
		return nil
	}

	return &Menu{MenuShell{Container{Widget{InitiallyUnowned{obj}}}}}
}

// NewMenu is a wrapper around gtk_menu_new().
func NewMenu() (*Menu, error) {
	c := C.gtk_menu_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapMenu(obj), nil
}

// GtkMenuNewFromModel is a wrapper around gtk_menu_new_from_model().
func GtkMenuNewFromModel(model *GMenuModel) (*Menu, error) {
	c := C.gtk_menu_new_from_model(C.toGMenuModel(unsafe.Pointer(model.Native())))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapMenu(obj), nil
}

// SetScreen is a wrapper around gtk_menu_set_screen().
func (v *Menu) SetScreen(screen *Screen) {
	C.gtk_menu_set_screen(v.native(), (*C.GdkScreen)(unsafe.Pointer(screen.Native())))
}

// Attach is a wrapper around gtk_menu_attach().
func (v *Menu) Attach(child IWidget, l, r, t, b uint) {
	C.gtk_menu_attach(
		v.native(),
		child.toWidget(),
		C.guint(l),
		C.guint(r),
		C.guint(t),
		C.guint(b))
}

// SetMonitor is a wrapper around gtk_menu_set_monitor().
func (v *Menu) SetMonitor(monitor_num int) {
	C.gtk_menu_set_monitor(v.native(), C.gint(monitor_num))
}

// GetMonitor is a wrapper around gtk_menu_get_monitor().
func (v *Menu) GetMonitor() int {
	return int(C.gtk_menu_get_monitor(v.native()))
}

// ReorderChild is a wrapper around gtk_menu_reorder_child().
func (v *Menu) ReorderChild(child IWidget, position int) {
	C.gtk_menu_reorder_child(v.native(), child.toWidget(), C.gint(position))
}

// SetReserveToggleSize is a wrapper around gtk_menu_set_reserve_toggle_size().
func (v *Menu) SetReserveToggleSize(reserve bool) {
	C.gtk_menu_set_reserve_toggle_size(v.native(), CBool(reserve))
}

// GetReserveToggleSize is a wrapper around gtk_menu_get_reserve_toggle_size().
func (v *Menu) GetReserveToggleSize() bool {
	return GoBool(C.gtk_menu_get_reserve_toggle_size(v.native()))
}

// Popdown is a wrapper around gtk_menu_popdown().
func (v *Menu) Popdown() {
	C.gtk_menu_popdown(v.native())
}

// GetActive is a wrapper around gtk_menu_get_active().
func (v *Menu) GetActive() (*Menu, error) {
	c := C.gtk_menu_get_active(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapMenu(obj), nil
}

// SetActive is a wrapper around gtk_menu_set_active().
func (v *Menu) SetActive(index uint) {
	C.gtk_menu_set_active(v.native(), C.guint(index))
}

// GetAttachWidget is a wrapper around gtk_menu_get_attach_widget().
func (v *Menu) GetAttachWidget() IWidget {
	c := C.gtk_menu_get_attach_widget(v.native())
	if c == nil {
		return nil
	}
	return castWidget(c)
}
