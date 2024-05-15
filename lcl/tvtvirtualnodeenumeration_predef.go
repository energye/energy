//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

type IVTVirtualNodeEnumeration interface {
	GetEnumerator() IVTVirtualNodeEnumerator
}

type TVTVirtualNodeEnumeration struct {
	instance unsafePointer
}

func (m *TVTVirtualNodeEnumeration) GetEnumerator() IVTVirtualNodeEnumerator {
	//TODO no impl
	return nil
}
