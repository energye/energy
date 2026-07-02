//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build cgo

package types

/*
#cgo pkg-config: glib-2.0 gobject-2.0 gdk-3.0
#include <glib.h>
#include <glib-object.h>
#include <gdk/gdk.h>
#include <stdlib.h>

// G_TYPE_IS_VALUE is a macro, wrap it as a function.
static inline gboolean _g_type_is_value(GType t) {
    return G_TYPE_IS_VALUE(t);
}

// G_TYPE_FROM_INSTANCE is a macro, wrap it as a function.
static inline GType _g_type_from_instance(gpointer inst) {
    return G_TYPE_FROM_INSTANCE(inst);
}

static GdkAtom toGdkAtom(void *p) { return (GdkAtom)p; }
*/
import "C"
import "unsafe"

// IsValue checks whether the passed in type can be used for g_value_init().
func (t Type) IsValue() bool {
	return C._g_type_is_value(C.GType(t)) != 0
}

// Name is a wrapper around g_type_name().
func (t Type) Name() string {
	return C.GoString((*C.char)(C.g_type_name(C.GType(t))))
}

// Depth is a wrapper around g_type_depth().
func (t Type) Depth() uint {
	return uint(C.g_type_depth(C.GType(t)))
}

// Parent is a wrapper around g_type_parent().
func (t Type) Parent() Type {
	return Type(C.g_type_parent(C.GType(t)))
}

// IsA is a wrapper around g_type_is_a().
func (t Type) IsA(isAType Type) bool {
	return C.g_type_is_a(C.GType(t), C.GType(isAType)) != 0
}

// TypeFromName is a wrapper around g_type_from_name().
func TypeFromName(typeName string) Type {
	cstr := (*C.gchar)(C.CString(typeName))
	defer C.free(unsafe.Pointer(cstr))
	return Type(C.g_type_from_name(cstr))
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
