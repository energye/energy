package cgo

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"reflect"
	"runtime"
	"unsafe"
)

type SelectionData struct {
	GtkSelectionData *C.GtkSelectionData
}

func ToSelectionData(pointer unsafe.Pointer) *SelectionData {
	p := (*C.GtkSelectionData)(pointer)
	return &SelectionData{GtkSelectionData: p}
}

// native returns a pointer to the underlying GtkSelectionData.
func (v *SelectionData) native() *C.GtkSelectionData {
	if v == nil {
		return nil
	}
	return v.GtkSelectionData
}

// GetLength is a wrapper around gtk_selection_data_get_length().
func (v *SelectionData) GetLength() int {
	return int(C.gtk_selection_data_get_length(v.native()))
}

// GetData is a wrapper around gtk_selection_data_get_data_with_length().
// It returns a slice of the correct size with the copy of the selection's data.
func (v *SelectionData) GetData() []byte {
	var length C.gint
	c := C.gtk_selection_data_get_data_with_length(v.native(), &length)

	// Only set if length is valid.
	if int(length) < 1 {
		return nil
	}

	var data []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sliceHeader.Len = int(length)
	sliceHeader.Cap = int(length)
	sliceHeader.Data = uintptr(unsafe.Pointer(c))

	// Keep the SelectionData alive for as long as the byte slice is.
	runtime.SetFinalizer(&data, func(*[]byte) {
		runtime.KeepAlive(v)
	})

	return data
}

// SetData is a wrapper around gtk_selection_data_set_data_with_length().
func (v *SelectionData) SetData(atom Atom, data []byte) {
	C.gtk_selection_data_set(
		v.native(),
		C.GdkAtom(unsafe.Pointer(atom)),
		C.gint(8), (*C.guchar)(&data[0]), C.gint(len(data)),
	)
}

// GetText is a wrapper around gtk_selection_data_get_text(). It returns a copy
// of the string from SelectionData and frees the C reference.
func (v *SelectionData) GetText() string {
	charptr := C.gtk_selection_data_get_text(v.native())
	if charptr == nil {
		return ""
	}
	defer C.g_free(C.gpointer(charptr))
	return ucharString(charptr)
}

// SetText is a wrapper around gtk_selection_data_set_text().
func (v *SelectionData) SetText(text string) bool {
	textPtr := *(*[]byte)(unsafe.Pointer(&text))
	return GoBool(C.gtk_selection_data_set_text(
		v.native(),
		// https://play.golang.org/p/PmGaLDhRuEU
		// This is probably safe since we expect Gdk to copy the string anyway.
		(*C.gchar)(unsafe.Pointer(&textPtr[0])), C.int(len(text)),
	))
}

// SetURIs is a wrapper around gtk_selection_data_set_uris().
func (v *SelectionData) SetURIs(uris []string) bool {
	var clist = C.make_strings(C.int(len(uris)))
	for i := range uris {
		cstring := C.CString(uris[i])
		// This defer will only run once the function exits, not once the loop
		// exits, so it's perfectly fine.
		defer C.free(unsafe.Pointer(cstring))
		C.set_string(clist, C.int(i), (*C.gchar)(cstring))
	}
	return GoBool(C.gtk_selection_data_set_uris(v.native(), clist))
}

// GetURIs is a wrapper around gtk_selection_data_get_uris().
func (v *SelectionData) GetURIs() []string {
	uriPtrs := C.gtk_selection_data_get_uris(v.native())
	return toGoStringArray(uriPtrs)
}

func (v *SelectionData) Free() {
	C.gtk_selection_data_free(v.native())
}
