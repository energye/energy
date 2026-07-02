package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"

import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// MessageDialog is a representation of GTK's GtkMessageDialog.
type MessageDialog struct {
	Dialog
}

func (v *MessageDialog) native() *C.GtkMessageDialog {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkMessageDialog(unsafe.Pointer(v.GObject))
}

func wrapMessageDialog(obj *Object) *MessageDialog {
	return &MessageDialog{Dialog{Window{Bin{Container{Widget{InitiallyUnowned{obj}}}}}}}
}

// MessageDialogNew is a wrapper around gtk_message_dialog_new().
func MessageDialogNew(parent IWindow, flags DialogFlags, mType MessageType, buttons ButtonsType, message string) *MessageDialog {
	cstr := C.CString(message)
	defer C.free(unsafe.Pointer(cstr))
	var w *C.GtkWindow
	if parent != nil {
		w = (*C.GtkWindow)(unsafe.Pointer(parent.Instance()))
	}
	c := C._gtk_message_dialog_new(w,
		C.GtkDialogFlags(flags), C.GtkMessageType(mType),
		C.GtkButtonsType(buttons), (*C.char)(cstr))
	if c == nil {
		return nil
	}
	return wrapMessageDialog(ToGoObject(unsafe.Pointer(c)))
}

// FormatSecondaryText is a wrapper around gtk_message_dialog_format_secondary_text().
func (v *MessageDialog) FormatSecondaryText(message string) {
	cstr := C.CString(message)
	defer C.free(unsafe.Pointer(cstr))
	C._gtk_message_dialog_format_secondary_text(v.native(), (*C.gchar)(cstr))
}
