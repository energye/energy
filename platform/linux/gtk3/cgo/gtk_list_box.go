package cgo

// #include <gtk/gtk.h>
// #include "gtk_header_bar.go.h"
import "C"
import (
	"github.com/energye/energy/v3/platform/linux/callback"
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

type ListBox struct {
	Container
}

func (v *ListBox) native() *C.GtkListBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	return (*C.GtkListBox)(unsafe.Pointer(v.GObject))
}

func wrapListBox(obj *Object) *ListBox {
	return &ListBox{Container{Widget{InitiallyUnowned{obj}}}}
}

func NewListBox() *ListBox {
	c := C.gtk_list_box_new()
	if c == nil {
		return nil
	}
	return wrapListBox(ToGoObject(unsafe.Pointer(c)))
}

func (v *ListBox) Prepend(child IWidget) {
	C.gtk_list_box_prepend(v.native(), GtkWidget(child))
}

func (v *ListBox) Insert(child IWidget, position int) {
	C.gtk_list_box_insert(v.native(), GtkWidget(child), C.gint(position))
}

func (v *ListBox) SelectRow(row IWidget) {
	C.gtk_list_box_select_row(v.native(), C.toGtkListBoxRow(unsafe.Pointer(row.Instance())))
}

func (v *ListBox) GetSelectedRow() IWidget {
	c := C.gtk_list_box_get_selected_row(v.native())
	if c == nil {
		return nil
	}
	return wrapWidget(ToGoObject(unsafe.Pointer(c)))
}

func (v *ListBox) SetSelectionMode(mode SelectionMode) {
	C.gtk_list_box_set_selection_mode(v.native(), C.GtkSelectionMode(mode))
}

func (v *ListBox) GetSelectionMode() SelectionMode {
	return SelectionMode(C.gtk_list_box_get_selection_mode(v.native()))
}

func (m *ListBox) SetOnRowSelected(fn TNotifyEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnRowSelected, callback.C_trampoline_3_void, fn, 0)
}
