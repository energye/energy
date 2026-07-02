//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build 386 || arm

package nocgo

import "unsafe"

// TreeIter is a representation of GTK's GtkTreeIter.
// 32-bit layout: stamp(4) + user_data(4) + user_data2(4) + user_data3(4) = 16 bytes.
type TreeIter struct {
	data [16]byte
}

func (m *TreeIter) Instance() uintptr {
	return uintptr(unsafe.Pointer(&m.data[0]))
}
