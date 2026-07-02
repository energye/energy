package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import "unsafe"

// TreeIter is a representation of GTK's GtkTreeIter.
type TreeIter struct {
	GtkTreeIter C.GtkTreeIter
}

func (v *TreeIter) native() *C.GtkTreeIter {
	if v == nil {
		return nil
	}
	return &v.GtkTreeIter
}

func (m *TreeIter) Instance() uintptr {
	return uintptr(unsafe.Pointer(m.native()))
}
