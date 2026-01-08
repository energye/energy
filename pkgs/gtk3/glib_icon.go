package gtk3

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import (
	"errors"
	"unsafe"
)

// Icon is a representation of GIO's GIcon.
// Interface for icons
type Icon struct {
	*Object
}

// native returns a pointer to the underlying GIcon.
func (v *Icon) native() *C.GIcon {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGIcon(p)
}

// NativePrivate: to be used inside Gotk3 only.
func (v *Icon) NativePrivate() *C.GIcon {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGIcon(p)
}

// Native returns a pointer to the underlying GIcon.
func (v *Icon) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalIcon(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapIcon(obj), nil
}

func wrapIcon(obj *Object) *Icon {
	return &Icon{obj}
}

// Equal is a wrapper around g_icon_equal().
func (v *Icon) Equal(icon *Icon) bool {
	return GoBool(C.g_icon_equal(v.native(), icon.native()))
}

// ToString is a wrapper around g_icon_to_string().
func (v *Icon) ToString() string {
	var s string
	if c := C.g_icon_to_string(v.native()); c != nil {
		s = GoString(c)
		defer C.g_free((C.gpointer)(c))
	}

	return s
}

// NewIconForString is a wrapper around g_icon_new_for_string().
func NewIconForString(str string) (*Icon, error) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError
	c := C.g_icon_new_for_string((*C.gchar)(cstr), &err)
	if c == nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}
	obj := &Object{ToCObject(unsafe.Pointer(c))}
	i := &Icon{obj}
	return i, nil
}

// FileIcon is a representation of GIO's GFileIcon.
type FileIcon struct {
	*Object
}

// native returns a pointer to the underlying GFileIcon.
func (v *FileIcon) native() *C.GFileIcon {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGFileIcon(p)
}

// NativePrivate to be used inside Gotk3 only.
func (v *FileIcon) NativePrivate() *C.GFileIcon {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGFileIcon(p)
}

// Native returns a pointer to the underlying GFileIcon.
func (v *FileIcon) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalFileIcon(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := ToGoObject(unsafe.Pointer(c))
	return wrapFileIcon(obj), nil
}

func wrapFileIcon(obj *Object) *FileIcon {
	return &FileIcon{obj}
}

// NewFileIconN is a wrapper around g_file_icon_new().
// This version respect Gtk3 documentation.
func NewFileIconN(file *File) *Icon {
	c := C.g_file_icon_new(file.native())
	if c == nil {
		return nil
	}
	return wrapIcon(ToGoObject(unsafe.Pointer(c)))
}

// NewFileIcon is a wrapper around g_file_icon_new().
// To not break previous implementation of GFileIcon ...
func NewFileIcon(path string) *Icon {
	file := NewFile(path)

	c := C.g_file_icon_new(file.native())
	if c == nil {
		return nil
	}
	return wrapIcon(ToGoObject(unsafe.Pointer(c)))
}
