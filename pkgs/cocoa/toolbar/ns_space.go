//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin

package toolbar

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa
#include "config.h"

*/
import "C"

func (m *NSToolBar) AddFlexibleSpace() {
	m.items.Add(GetStringConstValue(C.NSToolbarFlexibleSpaceItemIdentifier), nil)
	C.AddToolbarFlexibleSpace(m.toolbar)
}

func (m *NSToolBar) AddSpace() {
	m.items.Add(GetStringConstValue(C.NSToolbarSpaceItemIdentifier), nil)
	C.AddToolbarSpace(m.toolbar)
}
