package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
	"github.com/energye/energy/v3/platform/linux/callback"
)

// MenuShell is a representation of GTK's GtkMenuShell.
type MenuShell struct {
	Container
}

func (v *MenuShell) native() *C.GtkMenuShell {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkMenuShell(unsafe.Pointer(v.GObject))
}

func wrapMenuShell(obj *Object) *MenuShell {
	return &MenuShell{Container{Widget{InitiallyUnowned{obj}}}}
}

// Append is a wrapper around gtk_menu_shell_append().
func (v *MenuShell) Append(child IWidget) {
	C.gtk_menu_shell_append(v.native(), GtkWidget(child))
}

// Prepend is a wrapper around gtk_menu_shell_prepend().
func (v *MenuShell) Prepend(child IWidget) {
	C.gtk_menu_shell_prepend(v.native(), GtkWidget(child))
}

// Insert is a wrapper around gtk_menu_shell_insert().
func (v *MenuShell) Insert(child IWidget, position int) {
	C.gtk_menu_shell_insert(v.native(), GtkWidget(child), C.gint(position))
}

// Deactivate is a wrapper around gtk_menu_shell_deactivate().
func (v *MenuShell) Deactivate() {
	C.gtk_menu_shell_deactivate(v.native())
}

// SelectItem is a wrapper around gtk_menu_shell_select_item().
func (v *MenuShell) SelectItem(child IWidget) {
	C.gtk_menu_shell_select_item(v.native(), GtkWidget(child))
}

// SelectFirst is a wrapper around gtk_menu_shell_select_first().
func (v *MenuShell) SelectFirst(searchSensitive bool) {
	C.gtk_menu_shell_select_first(v.native(), CBool(searchSensitive))
}

// Deselect is a wrapper around gtk_menu_shell_deselect().
func (v *MenuShell) Deselect() {
	C.gtk_menu_shell_deselect(v.native())
}

// ActivateItem is a wrapper around gtk_menu_shell_activate_item().
func (v *MenuShell) ActivateItem(child IWidget, forceDeactivate bool) {
	C.gtk_menu_shell_activate_item(v.native(), GtkWidget(child), CBool(forceDeactivate))
}

// Cancel is a wrapper around gtk_menu_shell_cancel().
func (v *MenuShell) Cancel() {
	C.gtk_menu_shell_cancel(v.native())
}

// SetTakeFocus is a wrapper around gtk_menu_shell_set_take_focus().
func (v *MenuShell) SetTakeFocus(takeFocus bool) {
	C.gtk_menu_shell_set_take_focus(v.native(), CBool(takeFocus))
}

// GetTakeFocus is a wrapper around gtk_menu_shell_get_take_focus().
func (v *MenuShell) GetTakeFocus() bool {
	return GoBool(C.gtk_menu_shell_get_take_focus(v.native()))
}

// GetSelectedItem is a wrapper around gtk_menu_shell_get_selected_item().
func (v *MenuShell) GetSelectedItem() (IWidget, error) {
	c := C.gtk_menu_shell_get_selected_item(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return castWidget(c), nil
}

// GetParentShell is a wrapper around gtk_menu_shell_get_parent_shell().
func (v *MenuShell) GetParentShell() (IMenuShell, error) {
	c := C.gtk_menu_shell_get_parent_shell(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapMenuShell(obj), nil
}

// BindModel is a wrapper around gtk_menu_shell_bind_model().
func (v *MenuShell) BindModel(model *GMenuModel, actionNamespace string, withSeparators bool) {
	cstr := C.CString(actionNamespace)
	defer C.free(unsafe.Pointer(cstr))
	var mptr unsafe.Pointer
	if model != nil {
		mptr = unsafe.Pointer(model.Native())
	}
	C.gtk_menu_shell_bind_model(v.native(), (*C.GMenuModel)(mptr), cstr, CBool(withSeparators))
}

func (m *MenuShell) SetOnDeactivate(fn TNotifyEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnDeactivate, callback.C_trampoline_2_void, fn, 0)
}

func (m *MenuShell) SetOnSelectionDone(fn TNotifyEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnSelectionDone, callback.C_trampoline_2_void, fn, 0)
}
