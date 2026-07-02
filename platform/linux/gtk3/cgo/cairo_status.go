package cgo

// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

var key_Status = map[Status]string{

	STATUS_SUCCESS:                   "CAIRO_STATUS_SUCCESS",
	STATUS_NO_MEMORY:                 "CAIRO_STATUS_NO_MEMORY",
	STATUS_INVALID_RESTORE:           "CAIRO_STATUS_INVALID_RESTORE",
	STATUS_INVALID_POP_GROUP:         "CAIRO_STATUS_INVALID_POP_GROUP",
	STATUS_NO_CURRENT_POINT:          "CAIRO_STATUS_NO_CURRENT_POINT",
	STATUS_INVALID_MATRIX:            "CAIRO_STATUS_INVALID_MATRIX",
	STATUS_INVALID_STATUS:            "CAIRO_STATUS_INVALID_STATUS",
	STATUS_NULL_POINTER:              "CAIRO_STATUS_NULL_POINTER",
	STATUS_INVALID_STRING:            "CAIRO_STATUS_INVALID_STRING",
	STATUS_INVALID_PATH_DATA:         "CAIRO_STATUS_INVALID_PATH_DATA",
	STATUS_READ_ERROR:                "CAIRO_STATUS_READ_ERROR",
	STATUS_WRITE_ERROR:               "CAIRO_STATUS_WRITE_ERROR",
	STATUS_SURFACE_FINISHED:          "CAIRO_STATUS_SURFACE_FINISHED",
	STATUS_SURFACE_TYPE_MISMATCH:     "CAIRO_STATUS_SURFACE_TYPE_MISMATCH",
	STATUS_PATTERN_TYPE_MISMATCH:     "CAIRO_STATUS_PATTERN_TYPE_MISMATCH",
	STATUS_INVALID_CONTENT:           "CAIRO_STATUS_INVALID_CONTENT",
	STATUS_INVALID_FORMAT:            "CAIRO_STATUS_INVALID_FORMAT",
	STATUS_INVALID_VISUAL:            "CAIRO_STATUS_INVALID_VISUAL",
	STATUS_FILE_NOT_FOUND:            "CAIRO_STATUS_FILE_NOT_FOUND",
	STATUS_INVALID_DASH:              "CAIRO_STATUS_INVALID_DASH",
	STATUS_INVALID_DSC_COMMENT:       "CAIRO_STATUS_INVALID_DSC_COMMENT",
	STATUS_INVALID_INDEX:             "CAIRO_STATUS_INVALID_INDEX",
	STATUS_CLIP_NOT_REPRESENTABLE:    "CAIRO_STATUS_CLIP_NOT_REPRESENTABLE",
	STATUS_TEMP_FILE_ERROR:           "CAIRO_STATUS_TEMP_FILE_ERROR",
	STATUS_INVALID_STRIDE:            "CAIRO_STATUS_INVALID_STRIDE",
	STATUS_FONT_TYPE_MISMATCH:        "CAIRO_STATUS_FONT_TYPE_MISMATCH",
	STATUS_USER_FONT_IMMUTABLE:       "CAIRO_STATUS_USER_FONT_IMMUTABLE",
	STATUS_USER_FONT_ERROR:           "CAIRO_STATUS_USER_FONT_ERROR",
	STATUS_NEGATIVE_COUNT:            "CAIRO_STATUS_NEGATIVE_COUNT",
	STATUS_INVALID_CLUSTERS:          "CAIRO_STATUS_INVALID_CLUSTERS",
	STATUS_INVALID_SLANT:             "CAIRO_STATUS_INVALID_SLANT",
	STATUS_INVALID_WEIGHT:            "CAIRO_STATUS_INVALID_WEIGHT",
	STATUS_INVALID_SIZE:              "CAIRO_STATUS_INVALID_SIZE",
	STATUS_USER_FONT_NOT_IMPLEMENTED: "CAIRO_STATUS_USER_FONT_NOT_IMPLEMENTED",
	STATUS_DEVICE_TYPE_MISMATCH:      "CAIRO_STATUS_DEVICE_TYPE_MISMATCH",
	STATUS_DEVICE_ERROR:              "CAIRO_STATUS_DEVICE_ERROR",
}

//func StatusToString(status Status) string {
//	s, ok := key_Status[status]
//	if !ok {
//		s = "CAIRO_STATUS_UNDEFINED"
//	}
//	return s
//}// String returns a readable status messsage usable in texts.
//func (s Status) String() string {
//	str := StatusToString(s)
//	str = strings.Replace(str, "CAIRO_STATUS_", "", 1)
//	str = strings.Replace(str, "_", " ", 0)
//	return strings.ToLower(str)
//}
//
//// ToError returns the error for the status. Returns nil if success.
//func (s Status) ToError() error {
//	if s == STATUS_SUCCESS {
//		return nil
//	}
//	return errors.New(s.String())
//}

func marshalStatus(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Status(c), nil
}

func marshalFillRule(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return FillRule(c), nil
}

func marshalLineCap(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return LineCap(c), nil
}

func marshalLineJoin(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return LineJoin(c), nil
}

func marshalOperator(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Operator(c), nil
}
