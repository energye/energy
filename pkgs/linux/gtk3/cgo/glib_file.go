package cgo

// #include <gio/gio.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import "unsafe"

// File is a representation of GIO's GFile.
type File struct {
	*Object
}

// native returns a pointer to the underlying GFile.
func (v *File) native() *C.GFile {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGFile(p)
}

// NativePrivate: to be used inside Gotk3 only.
func (v *File) NativePrivate() *C.GFile {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGFile(p)
}

// Native returns a pointer to the underlying GFile.
func (v *File) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func wrapFile(obj *Object) *File {
	return &File{obj}
}

// NewFile is a wrapper around g_file_new_for_path().
// To avoid breaking previous implementation of GFile ...
func NewFile(path string) *File {
	return NewFileForPath(path)
}

// NewFileForPath is a wrapper around g_file_new_for_path().
func NewFileForPath(path string) *File {
	cstr := (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(cstr))

	c := C.g_file_new_for_path(cstr)
	if c == nil {
		return nil
	}
	return wrapFile(ToGoObject(unsafe.Pointer(c)))
}
