package gtk3

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gio/gio.h>
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"
)

// NewImage is a wrapper around gtk_image_new().
func NewImage() *Image {
	c := C.gtk_image_new()
	if c == nil {
		return nil
	}
	return wrapImage(ToGoObject(unsafe.Pointer(c)))
}

// NewImageFromFile is a wrapper around gtk_image_new_from_file().
func NewImageFromFile(filename string) *Image {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_image_new_from_file((*C.gchar)(cstr))
	if c == nil {
		return nil
	}
	return wrapImage(ToGoObject(unsafe.Pointer(c)))
}

// NewImageFromResource is a wrapper around gtk_image_new_from_resource().
func NewImageFromResource(resourcePath string) *Image {
	cstr := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_image_new_from_resource((*C.gchar)(cstr))
	if c == nil {
		return nil
	}
	return wrapImage(ToGoObject(unsafe.Pointer(c)))
}

// NewImageFromIconName is a wrapper around gtk_image_new_from_icon_name().
func NewImageFromIconName(iconName string, size IconSize) *Image {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_image_new_from_icon_name((*C.gchar)(cstr), C.GtkIconSize(size))
	if c == nil {
		return nil
	}
	return wrapImage(ToGoObject(unsafe.Pointer(c)))
}

// Image is a representation of GTK's GtkImage.
type Image struct {
	Widget
}

// native returns a pointer to the underlying GtkImage.
func (v *Image) native() *C.GtkImage {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkImage(p)
}

func marshalImage(p uintptr) any {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapImage(ToGoObject(unsafe.Pointer(c)))
}

func wrapImage(obj *Object) *Image {
	if obj == nil {
		return nil
	}
	return &Image{Widget{InitiallyUnowned{obj}}}
}

// Clear is a wrapper around gtk_image_clear().
func (v *Image) Clear() {
	C.gtk_image_clear(v.native())
}

// SetFromFile is a wrapper around gtk_image_set_from_file().
func (v *Image) SetFromFile(filename string) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_image_set_from_file(v.native(), (*C.gchar)(cstr))
}

// SetFromResource is a wrapper around gtk_image_set_from_resource().
func (v *Image) SetFromResource(resourcePath string) {
	cstr := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_image_set_from_resource(v.native(), (*C.gchar)(cstr))
}

// SetFromIconName is a wrapper around gtk_image_set_from_icon_name().
func (v *Image) SetFromIconName(iconName string, size IconSize) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_image_set_from_icon_name(v.native(), (*C.gchar)(cstr),
		C.GtkIconSize(size))
}

// SetFromGIcon is a wrapper around gtk_image_set_from_gicon()
func (v *Image) SetFromGIcon(icon *Icon, size IconSize) {
	C.gtk_image_set_from_gicon(
		v.native(),
		(*C.GIcon)(icon.NativePrivate()),
		C.GtkIconSize(size))
}

// SetPixelSize is a wrapper around gtk_image_set_pixel_size().
func (v *Image) SetPixelSize(pixelSize int) {
	C.gtk_image_set_pixel_size(v.native(), C.gint(pixelSize))
}

// GetStorageType is a wrapper around gtk_image_get_storage_type().
func (v *Image) GetStorageType() ImageType {
	c := C.gtk_image_get_storage_type(v.native())
	return ImageType(c)
}

// GetIconName is a wrapper around gtk_image_get_icon_name().
func (v *Image) GetIconName() (string, IconSize) {
	var iconName *C.gchar
	var size C.GtkIconSize
	C.gtk_image_get_icon_name(v.native(), &iconName, &size)
	return GoString(iconName), IconSize(size)
}

// GetGIcon is a wrapper around gtk_image_get_gicon()
func (v *Image) GetGIcon() (*Icon, IconSize, error) {
	var gicon *C.GIcon
	var size *C.GtkIconSize
	C.gtk_image_get_gicon(v.native(), &gicon, size)

	if gicon == nil {
		return nil, ICON_SIZE_INVALID, nilPtrErr
	}

	obj := &Object{ToCObject(unsafe.Pointer(gicon))}
	i := &Icon{obj}
	return i, IconSize(*size), nil
}

// GetPixelSize is a wrapper around gtk_image_get_pixel_size().
func (v *Image) GetPixelSize() int {
	c := C.gtk_image_get_pixel_size(v.native())
	return int(c)
}
