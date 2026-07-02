package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// ColorChooserDialog is a representation of GTK's GtkColorChooserDialog.
type ColorChooserDialog struct {
	Dialog
}

func (v *ColorChooserDialog) native() *C.GtkColorChooserDialog {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkColorChooserDialog(unsafe.Pointer(v.GObject))
}

func wrapColorChooserDialog(obj *Object) *ColorChooserDialog {
	return &ColorChooserDialog{Dialog{Window{Bin{Container{Widget{InitiallyUnowned{obj}}}}}}}
}

// NewColorChooserDialog is a wrapper around gtk_color_chooser_dialog_new().
func NewColorChooserDialog(title string, parent IWindow) *ColorChooserDialog {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	var w *C.GtkWindow
	if parent != nil {
		w = (*C.GtkWindow)(unsafe.Pointer(parent.Instance()))
	}
	c := C.gtk_color_chooser_dialog_new((*C.gchar)(cstr), w)
	if c == nil {
		return nil
	}
	return wrapColorChooserDialog(ToGoObject(unsafe.Pointer(c)))
}

// GetUseAlpha is a wrapper around gtk_color_chooser_get_use_alpha().
func (v *ColorChooserDialog) GetUseAlpha() bool {
	return GoBool(C.gtk_color_chooser_get_use_alpha(C.toGtkColorChooser(unsafe.Pointer(v.GObject))))
}

// SetUseAlpha is a wrapper around gtk_color_chooser_set_use_alpha().
func (v *ColorChooserDialog) SetUseAlpha(useAlpha bool) {
	C.gtk_color_chooser_set_use_alpha(C.toGtkColorChooser(unsafe.Pointer(v.GObject)), CBool(useAlpha))
}
