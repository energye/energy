package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// cellRenderer is a cgo-internal interface for getting the C GtkCellRenderer pointer.
type cellRenderer interface {
	toCellRenderer() *C.GtkCellRenderer
}

// CellRendererText is a representation of GTK's GtkCellRendererText.
type CellRendererText struct {
	*Object
}

func (v *CellRendererText) native() *C.GtkCellRendererText {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkCellRendererText(unsafe.Pointer(v.GObject))
}

func (v *CellRendererText) toCellRenderer() *C.GtkCellRenderer {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkCellRenderer(unsafe.Pointer(v.GObject))
}

func wrapCellRendererText(obj *Object) *CellRendererText {
	return &CellRendererText{obj}
}

// NewCellRendererText is a wrapper around gtk_cell_renderer_text_new().
func NewCellRendererText() *CellRendererText {
	c := C.gtk_cell_renderer_text_new()
	if c == nil {
		return nil
	}
	return wrapCellRendererText(ToGoObject(unsafe.Pointer(c)))
}

// SetText sets the "text" property.
func (v *CellRendererText) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C._cell_renderer_text_set_text(v.native(), cstr)
}

// SetMarkup sets the "markup" property.
func (v *CellRendererText) SetMarkup(markup string) {
	cstr := C.CString(markup)
	defer C.free(unsafe.Pointer(cstr))
	C._cell_renderer_text_set_markup(v.native(), cstr)
}

// SetFont sets the "font" property (e.g. "Sans 12").
func (v *CellRendererText) SetFont(font string) {
	cstr := C.CString(font)
	defer C.free(unsafe.Pointer(cstr))
	C._cell_renderer_text_set_font(v.native(), cstr)
}

// SetForeground sets the "foreground" color property (e.g. "red", "#FF0000").
func (v *CellRendererText) SetForeground(color string) {
	cstr := C.CString(color)
	defer C.free(unsafe.Pointer(cstr))
	C._cell_renderer_text_set_foreground(v.native(), cstr)
}

// SetBackground sets the "background" color property.
func (v *CellRendererText) SetBackground(color string) {
	cstr := C.CString(color)
	defer C.free(unsafe.Pointer(cstr))
	C._cell_renderer_text_set_background(v.native(), cstr)
}

// SetAlignment sets the "alignment" property (0.0=left, 0.5=center, 1.0=right).
func (v *CellRendererText) SetAlignment(align float32) {
	C._cell_renderer_text_set_alignment(v.native(), C.float(align))
}

// SetEllipsize sets the "ellipsize" property.
func (v *CellRendererText) SetEllipsize(mode EllipsizeMode) {
	C._cell_renderer_text_set_ellipsize(v.native(), C.int(mode))
}

// SetWidthChars sets the "width-chars" property.
func (v *CellRendererText) SetWidthChars(n int) {
	C._cell_renderer_text_set_width_chars(v.native(), C.int(n))
}
