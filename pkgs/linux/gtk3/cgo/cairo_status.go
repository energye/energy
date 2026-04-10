package cgo

// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"
import (
	. "github.com/energye/energy/v3/pkgs/linux/types"
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

// FillRule is a representation of Cairo's cairo_fill_rule_t.
type FillRule int

const (
	FILL_RULE_WINDING  FillRule = C.CAIRO_FILL_RULE_WINDING
	FILL_RULE_EVEN_ODD FillRule = C.CAIRO_FILL_RULE_EVEN_ODD
)

func marshalFillRule(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return FillRule(c), nil
}

// LineCap is a representation of Cairo's cairo_line_cap_t.
type LineCap int

const (
	LINE_CAP_BUTT   LineCap = C.CAIRO_LINE_CAP_BUTT
	LINE_CAP_ROUND  LineCap = C.CAIRO_LINE_CAP_ROUND
	LINE_CAP_SQUARE LineCap = C.CAIRO_LINE_CAP_SQUARE
)

func marshalLineCap(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return LineCap(c), nil
}

// LineJoin is a representation of Cairo's cairo_line_join_t.
type LineJoin int

const (
	LINE_JOIN_MITER LineJoin = C.CAIRO_LINE_JOIN_MITER
	LINE_JOIN_ROUND LineJoin = C.CAIRO_LINE_JOIN_ROUND
	LINE_JOIN_BEVEL LineJoin = C.CAIRO_LINE_JOIN_BEVEL
)

func marshalLineJoin(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return LineJoin(c), nil
}

// Operator is a representation of Cairo's cairo_operator_t.
type Operator int

const (
	OPERATOR_CLEAR          Operator = C.CAIRO_OPERATOR_CLEAR
	OPERATOR_SOURCE         Operator = C.CAIRO_OPERATOR_SOURCE
	OPERATOR_OVER           Operator = C.CAIRO_OPERATOR_OVER
	OPERATOR_IN             Operator = C.CAIRO_OPERATOR_IN
	OPERATOR_OUT            Operator = C.CAIRO_OPERATOR_OUT
	OPERATOR_ATOP           Operator = C.CAIRO_OPERATOR_ATOP
	OPERATOR_DEST           Operator = C.CAIRO_OPERATOR_DEST
	OPERATOR_DEST_OVER      Operator = C.CAIRO_OPERATOR_DEST_OVER
	OPERATOR_DEST_IN        Operator = C.CAIRO_OPERATOR_DEST_IN
	OPERATOR_DEST_OUT       Operator = C.CAIRO_OPERATOR_DEST_OUT
	OPERATOR_DEST_ATOP      Operator = C.CAIRO_OPERATOR_DEST_ATOP
	OPERATOR_XOR            Operator = C.CAIRO_OPERATOR_XOR
	OPERATOR_ADD            Operator = C.CAIRO_OPERATOR_ADD
	OPERATOR_SATURATE       Operator = C.CAIRO_OPERATOR_SATURATE
	OPERATOR_MULTIPLY       Operator = C.CAIRO_OPERATOR_MULTIPLY
	OPERATOR_SCREEN         Operator = C.CAIRO_OPERATOR_SCREEN
	OPERATOR_OVERLAY        Operator = C.CAIRO_OPERATOR_OVERLAY
	OPERATOR_DARKEN         Operator = C.CAIRO_OPERATOR_DARKEN
	OPERATOR_LIGHTEN        Operator = C.CAIRO_OPERATOR_LIGHTEN
	OPERATOR_COLOR_DODGE    Operator = C.CAIRO_OPERATOR_COLOR_DODGE
	OPERATOR_COLOR_BURN     Operator = C.CAIRO_OPERATOR_COLOR_BURN
	OPERATOR_HARD_LIGHT     Operator = C.CAIRO_OPERATOR_HARD_LIGHT
	OPERATOR_SOFT_LIGHT     Operator = C.CAIRO_OPERATOR_SOFT_LIGHT
	OPERATOR_DIFFERENCE     Operator = C.CAIRO_OPERATOR_DIFFERENCE
	OPERATOR_EXCLUSION      Operator = C.CAIRO_OPERATOR_EXCLUSION
	OPERATOR_HSL_HUE        Operator = C.CAIRO_OPERATOR_HSL_HUE
	OPERATOR_HSL_SATURATION Operator = C.CAIRO_OPERATOR_HSL_SATURATION
	OPERATOR_HSL_COLOR      Operator = C.CAIRO_OPERATOR_HSL_COLOR
	OPERATOR_HSL_LUMINOSITY Operator = C.CAIRO_OPERATOR_HSL_LUMINOSITY
)

func marshalOperator(p uintptr) (any, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return Operator(c), nil
}
