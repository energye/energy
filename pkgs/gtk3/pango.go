package gtk3

// #include <pango/pango.h>
import "C"

// EllipsizeMode is a representation of Pango's PangoEllipsizeMode.
type EllipsizeMode int

const (
	ELLIPSIZE_NONE   EllipsizeMode = C.PANGO_ELLIPSIZE_NONE
	ELLIPSIZE_START  EllipsizeMode = C.PANGO_ELLIPSIZE_START
	ELLIPSIZE_MIDDLE EllipsizeMode = C.PANGO_ELLIPSIZE_MIDDLE
	ELLIPSIZE_END    EllipsizeMode = C.PANGO_ELLIPSIZE_END
)
