package cgo

// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

type SearchEntry struct {
	Entry
}

func (v *SearchEntry) native() *C.GtkSearchEntry {
	if v == nil || v.GObject == nil {
		return nil
	}
	return (*C.GtkSearchEntry)(unsafe.Pointer(v.GObject))
}

func wrapSearchEntry(obj *Object) *SearchEntry {
	if obj == nil {
		return nil
	}
	e := wrapEditable(obj)
	ce := wrapCellEditable(obj)
	return &SearchEntry{Entry{Widget{InitiallyUnowned{obj}}, *e, *ce}}
}

func NewSearchEntry() *SearchEntry {
	c := C.gtk_search_entry_new()
	if c == nil {
		return nil
	}
	return wrapSearchEntry(ToGoObject(unsafe.Pointer(c)))
}
