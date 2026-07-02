//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	"unsafe"
)

// TreeIter is a representation of GTK's GtkTreeIter.
type TreeIter struct {
	Stamp     int32
	userData  int32
	userData2 int32
	userData3 int32
}

func (m *TreeIter) Instance() uintptr {
	return uintptr(unsafe.Pointer(m))
}
