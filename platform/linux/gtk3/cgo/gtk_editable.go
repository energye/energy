package cgo

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gio/gio.h>
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"
)

// Editable is a representation of GTK's GtkEditable GInterface.
type Editable struct {
	*Object
}

// IEditable is an interface type implemented by all structs
// embedding an Editable.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkEditable.
type IEditable interface {
	toEditable() *C.GtkEditable
}

// native() returns a pointer to the underlying GObject as a GtkEditable.
func (v *Editable) native() *C.GtkEditable {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkEditable(p)
}

func wrapEditable(obj *Object) *Editable {
	if obj == nil {
		return nil
	}

	return &Editable{obj}
}

func (v *Editable) toEditable() *C.GtkEditable {
	if v == nil {
		return nil
	}
	return v.native()
}

// SelectRegion is a wrapper around gtk_editable_select_region().
func (v *Editable) SelectRegion(startPos, endPos int) {
	C.gtk_editable_select_region(v.native(), C.gint(startPos),
		C.gint(endPos))
}

// GetSelectionBounds is a wrapper around gtk_editable_get_selection_bounds().
func (v *Editable) GetSelectionBounds() (start, end int, nonEmpty bool) {
	var cstart, cend C.gint
	c := C.gtk_editable_get_selection_bounds(v.native(), &cstart, &cend)
	return int(cstart), int(cend), GoBool(c)
}

// InsertText is a wrapper around gtk_editable_insert_text(). The returned
// int is the position after the inserted text.
func (v *Editable) InsertText(newText string, position int) int {
	cstr := C.CString(newText)
	defer C.free(unsafe.Pointer(cstr))
	pos := new(C.gint)
	*pos = C.gint(position)
	C.gtk_editable_insert_text(v.native(), (*C.gchar)(cstr),
		C.gint(len(newText)), pos)
	return int(*pos)
}

// DeleteText is a wrapper around gtk_editable_delete_text().
func (v *Editable) DeleteText(startPos, endPos int) {
	C.gtk_editable_delete_text(v.native(), C.gint(startPos), C.gint(endPos))
}

// GetChars is a wrapper around gtk_editable_get_chars().
func (v *Editable) GetChars(startPos, endPos int) string {
	c := C.gtk_editable_get_chars(v.native(), C.gint(startPos),
		C.gint(endPos))
	defer C.free(unsafe.Pointer(c))
	return GoString(c)
}

// CutClipboard is a wrapper around gtk_editable_cut_clipboard().
func (v *Editable) CutClipboard() {
	C.gtk_editable_cut_clipboard(v.native())
}

// CopyClipboard is a wrapper around gtk_editable_copy_clipboard().
func (v *Editable) CopyClipboard() {
	C.gtk_editable_copy_clipboard(v.native())
}

// PasteClipboard is a wrapper around gtk_editable_paste_clipboard().
func (v *Editable) PasteClipboard() {
	C.gtk_editable_paste_clipboard(v.native())
}

// DeleteSelection is a wrapper around gtk_editable_delete_selection().
func (v *Editable) DeleteSelection() {
	C.gtk_editable_delete_selection(v.native())
}

// SetPosition is a wrapper around gtk_editable_set_position().
func (v *Editable) SetPosition(position int) {
	C.gtk_editable_set_position(v.native(), C.gint(position))
}

// GetPosition is a wrapper around gtk_editable_get_position().
func (v *Editable) GetPosition() int {
	c := C.gtk_editable_get_position(v.native())
	return int(c)
}

// SetEditable is a wrapper around gtk_editable_set_editable().
func (v *Editable) SetEditable(isEditable bool) {
	C.gtk_editable_set_editable(v.native(), CBool(isEditable))
}

// GetEditable is a wrapper around gtk_editable_get_editable().
func (v *Editable) GetEditable() bool {
	c := C.gtk_editable_get_editable(v.native())
	return GoBool(c)
}

// CellEditable is a representation of GTK's GtkCellEditable GInterface.
type CellEditable struct {
	InitiallyUnowned
}

// ICellEditable is an interface type implemented by all structs
// embedding an CellEditable. It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkCellEditable.
type ICellEditable interface {
	toCellEditable() *C.GtkCellEditable
	ToEntry() *Entry
}

// native() returns a pointer to the underlying GObject as a GtkCellEditable.
func (v *CellEditable) native() *C.GtkCellEditable {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkCellEditable(p)
}

func wrapCellEditable(obj *Object) *CellEditable {
	if obj == nil {
		return nil
	}
	return &CellEditable{InitiallyUnowned{obj}}
}

func (v *CellEditable) toCellEditable() *C.GtkCellEditable {
	if v == nil {
		return nil
	}
	return v.native()
}

// ToEntry is a helper tool, e.g: it returns *gtk.CellEditable as a *gtk.Entry
// that embedding this CellEditable instance, then it can be used with
// CellRendererText to adding EntryCompletion tools or intercepting EntryBuffer,
// (to bypass "canceled" signal for example) then record entry, and much more.
func (v *CellEditable) ToEntry() *Entry {
	return &Entry{Widget{InitiallyUnowned{v.Object}},
		Editable{v.Object},
		*v}
}

// StartEditing is a wrapper around gtk_cell_editable_start_editing().
func (v *CellEditable) StartEditing(event *Event) {
	C.gtk_cell_editable_start_editing(v.native(),
		(*C.GdkEvent)(unsafe.Pointer(event.Native())))
}

// EditingDone is a wrapper around gtk_cell_editable_editing_done().
func (v *CellEditable) EditingDone() {
	C.gtk_cell_editable_editing_done(v.native())
}

// RemoveWidget is a wrapper around gtk_cell_editable_remove_widget().
func (v *CellEditable) RemoveWidget() {
	C.gtk_cell_editable_remove_widget(v.native())
}
