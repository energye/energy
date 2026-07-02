package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// FontChooserDialog is a representation of GTK's GtkFontChooserDialog.
type FontChooserDialog struct {
	Dialog
}

func (v *FontChooserDialog) native() *C.GtkFontChooserDialog {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkFontChooserDialog(unsafe.Pointer(v.GObject))
}

func wrapFontChooserDialog(obj *Object) *FontChooserDialog {
	return &FontChooserDialog{Dialog{Window{Bin{Container{Widget{InitiallyUnowned{obj}}}}}}}
}

// NewFontChooserDialog is a wrapper around gtk_font_chooser_dialog_new().
func NewFontChooserDialog(title string, parent IWindow) *FontChooserDialog {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	var w *C.GtkWindow
	if parent != nil {
		w = (*C.GtkWindow)(unsafe.Pointer(parent.Instance()))
	}
	c := C.gtk_font_chooser_dialog_new((*C.gchar)(cstr), w)
	if c == nil {
		return nil
	}
	return wrapFontChooserDialog(ToGoObject(unsafe.Pointer(c)))
}

// GetFont is a wrapper around gtk_font_chooser_get_font().
func (v *FontChooserDialog) GetFont() string {
	c := C.gtk_font_chooser_get_font(C.toGtkFontChooser(unsafe.Pointer(v.GObject)))
	if c == nil {
		return ""
	}
	s := C.GoString((*C.char)(c))
	C.g_free(C.gpointer(c))
	return s
}

// SetFont is a wrapper around gtk_font_chooser_set_font().
func (v *FontChooserDialog) SetFont(font string) {
	cstr := C.CString(font)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_font_chooser_set_font(C.toGtkFontChooser(unsafe.Pointer(v.GObject)), (*C.gchar)(cstr))
}

// GetPreviewText is a wrapper around gtk_font_chooser_get_preview_text().
func (v *FontChooserDialog) GetPreviewText() string {
	c := C.gtk_font_chooser_get_preview_text(C.toGtkFontChooser(unsafe.Pointer(v.GObject)))
	if c == nil {
		return ""
	}
	s := C.GoString((*C.char)(c))
	C.g_free(C.gpointer(c))
	return s
}

// SetPreviewText is a wrapper around gtk_font_chooser_set_preview_text().
func (v *FontChooserDialog) SetPreviewText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_font_chooser_set_preview_text(C.toGtkFontChooser(unsafe.Pointer(v.GObject)), (*C.gchar)(cstr))
}
