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

import "unsafe"

type NSView struct {
	NSResponder
}

func AsNSView(ptr unsafe.Pointer) *NSView {
	if ptr == nil {
		return nil
	}
	m := new(NSView)
	m.SetInstance(ptr)
	return m
}
