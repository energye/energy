//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build amd64 || arm64 || loong64

package nocgo

import "unsafe"

// TreeIter is a representation of GTK's GtkTreeIter.
// 64-bit layout: stamp(4) + pad(4) + user_data(8) + user_data2(8) + user_data3(8) = 32 bytes.
type TreeIter struct {
	data [32]byte
}

func (m *TreeIter) Instance() uintptr {
	return uintptr(unsafe.Pointer(&m.data[0]))
}
