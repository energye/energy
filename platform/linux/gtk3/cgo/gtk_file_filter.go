package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// FileFilter is a representation of GTK's GtkFileFilter.
type FileFilter struct {
	*Object
}

func (v *FileFilter) native() *C.GtkFileFilter {
	if v == nil || v.Object == nil {
		return nil
	}
	return C.toGtkFileFilter(unsafe.Pointer(v.GObject))
}

func wrapFileFilter(obj *Object) *FileFilter {
	if obj == nil {
		return nil
	}
	return &FileFilter{obj}
}

// NewFileFilter is a wrapper around gtk_file_filter_new().
func NewFileFilter() IFileFilter {
	c := C.gtk_file_filter_new()
	if c == nil {
		return nil
	}
	return wrapFileFilter(ToGoObject(unsafe.Pointer(c)))
}

// SetName is a wrapper around gtk_file_filter_set_name().
func (v *FileFilter) SetName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_file_filter_set_name(v.native(), (*C.gchar)(cstr))
}

// GetName is a wrapper around gtk_file_filter_get_name().
func (v *FileFilter) GetName() string {
	c := C.gtk_file_filter_get_name(v.native())
	if c == nil {
		return ""
	}
	return C.GoString((*C.char)(c))
}

// AddPattern is a wrapper around gtk_file_filter_add_pattern().
func (v *FileFilter) AddPattern(pattern string) {
	cstr := C.CString(pattern)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_file_filter_add_pattern(v.native(), (*C.gchar)(cstr))
}

// AddMimeType is a wrapper around gtk_file_filter_add_mime_type().
func (v *FileFilter) AddMimeType(mimeType string) {
	cstr := C.CString(mimeType)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_file_filter_add_mime_type(v.native(), (*C.gchar)(cstr))
}
