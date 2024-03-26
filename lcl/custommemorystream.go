//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// ICustomMemoryStream Parent: IStream
type ICustomMemoryStream interface {
	IStream
	Memory() uintptr             // property
	SaveToStream(Stream IStream) // procedure
	SaveToFile(FileName string)  // procedure
}

// TCustomMemoryStream Parent: TStream
type TCustomMemoryStream struct {
	TStream
}

func NewCustomMemoryStream() ICustomMemoryStream {
	r1 := LCL().SysCallN(1913)
	return AsCustomMemoryStream(r1)
}

func (m *TCustomMemoryStream) Memory() uintptr {
	r1 := LCL().SysCallN(1914, m.Instance())
	return uintptr(r1)
}

func CustomMemoryStreamClass() TClass {
	ret := LCL().SysCallN(1912)
	return TClass(ret)
}

func (m *TCustomMemoryStream) SaveToStream(Stream IStream) {
	LCL().SysCallN(1916, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TCustomMemoryStream) SaveToFile(FileName string) {
	LCL().SysCallN(1915, m.Instance(), PascalStr(FileName))
}
