package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"github.com/energye/energy/v3/platform/linux/callback"
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// Dialog is a representation of GTK's GtkDialog.
type Dialog struct {
	Window
}

func (v *Dialog) native() *C.GtkDialog {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkDialog(p)
}

func wrapDialog(obj *Object) *Dialog {
	return &Dialog{Window{Bin{Container{Widget{InitiallyUnowned{obj}}}}}}
}

// NewDialog is a wrapper around gtk_dialog_new().
func NewDialog() *Dialog {
	c := C.gtk_dialog_new()
	if c == nil {
		return nil
	}
	return wrapDialog(ToGoObject(unsafe.Pointer(c)))
}

// Run is a wrapper around gtk_dialog_run().
func (v *Dialog) Run() int {
	return int(C.gtk_dialog_run(v.native()))
}

// Response is a wrapper around gtk_dialog_response().
func (v *Dialog) Response(responseId int) {
	C.gtk_dialog_response(v.native(), C.gint(responseId))
}

// AddButton is a wrapper around gtk_dialog_add_button().
func (v *Dialog) AddButton(buttonText string, responseId int) IButton {
	cstr := C.CString(buttonText)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_dialog_add_button(v.native(), (*C.gchar)(cstr), C.gint(responseId))
	if c == nil {
		return nil
	}
	return wrapButton(ToGoObject(unsafe.Pointer(c)))
}

// SetDefaultResponse is a wrapper around gtk_dialog_set_default_response().
func (v *Dialog) SetDefaultResponse(responseId int) {
	C.gtk_dialog_set_default_response(v.native(), C.gint(responseId))
}

// GetContentArea is a wrapper around gtk_dialog_get_content_area().
func (v *Dialog) GetContentArea() IBox {
	c := C.gtk_dialog_get_content_area(v.native())
	if c == nil {
		return nil
	}
	return wrapBox(ToGoObject(unsafe.Pointer(c)))
}

// SetOnResponse is a callback for the "response" signal.
func (v *Dialog) SetOnResponse(fn TResponseEvent) ISignalHandlerID {
	return callback.Connect(v.Instance(), EsnResponse, callback.C_trampoline_3_void, fn, 0)
}
