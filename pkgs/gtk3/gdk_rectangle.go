package gtk3

// #cgo pkg-config: gdk-3.0 glib-2.0 gobject-2.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import "unsafe"

// Rectangle is a representation of GDK's GdkRectangle type.
type Rectangle struct {
	GdkRectangle C.GdkRectangle
}

func WrapRectangle(p uintptr) *Rectangle {
	return wrapRectangle((*C.GdkRectangle)(unsafe.Pointer(p)))
}

func wrapRectangle(obj *C.GdkRectangle) *Rectangle {
	if obj == nil {
		return nil
	}
	return &Rectangle{*obj}
}

// Native() returns a pointer to the underlying GdkRectangle.
func (r *Rectangle) native() *C.GdkRectangle {
	return &r.GdkRectangle
}

// RectangleIntersect is a wrapper around gdk_rectangle_intersect().
func (v *Rectangle) RectangleIntersect(rect *Rectangle) (*Rectangle, bool) {
	r := new(C.GdkRectangle)
	c := C.gdk_rectangle_intersect(v.native(), rect.native(), r)
	return wrapRectangle(r), GoBool(c)
}

// RectangleUnion is a wrapper around gdk_rectangle_union().
func (v *Rectangle) RectangleUnion(rect *Rectangle) *Rectangle {
	r := new(C.GdkRectangle)
	C.gdk_rectangle_union(v.native(), rect.native(), r)
	return wrapRectangle(r)
}

// NewRectangle helper function to create a GdkRectanlge
func NewRectangle(x, y, width, height int) *Rectangle {
	var r C.GdkRectangle
	r.x = C.int(x)
	r.y = C.int(y)
	r.width = C.int(width)
	r.height = C.int(height)
	return &Rectangle{r}
}

// SetRectangleInt helper function to set GdkRectanlge values
func (v *Rectangle) SetRectangleInt(x, y, width, height int) {
	v.native().x = C.int(x)
	v.native().y = C.int(y)
	v.native().width = C.int(width)
	v.native().height = C.int(height)
}

// GetRectangleInt helper function to get GdkRectanlge values
func (v *Rectangle) GetRectangleInt() (x, y, width, height int) {
	return int(v.native().x),
		int(v.native().y),
		int(v.native().width),
		int(v.native().height)
}

// GetX returns x field of the underlying GdkRectangle.
func (r *Rectangle) GetX() int {
	return int(r.native().x)
}

// SetX sets x field of the underlying GdkRectangle.
func (r *Rectangle) SetX(x int) {
	r.native().x = C.int(x)
}

// GetY returns y field of the underlying GdkRectangle.
func (r *Rectangle) GetY() int {
	return int(r.native().y)
}

// SetY sets y field of the underlying GdkRectangle.
func (r *Rectangle) SetY(y int) {
	r.native().y = C.int(y)
}

// GetWidth returns width field of the underlying GdkRectangle.
func (r *Rectangle) GetWidth() int {
	return int(r.native().width)
}

// SetWidth sets width field of the underlying GdkRectangle.
func (r *Rectangle) SetWidth(width int) {
	r.native().width = C.int(width)
}

// GetHeight returns height field of the underlying GdkRectangle.
func (r *Rectangle) GetHeight() int {
	return int(r.native().height)
}

// SetHeight sets height field of the underlying GdkRectangle.
func (r *Rectangle) SetHeight(height int) {
	r.native().height = C.int(height)
}
