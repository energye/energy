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

// IMemoryStream Parent: ICustomMemoryStream
type IMemoryStream interface {
	ICustomMemoryStream
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
	Clear()                        // procedure
	LoadFromStream(Stream IStream) // procedure
	LoadFromFile(FileName string)  // procedure
}

// TMemoryStream Parent: TCustomMemoryStream
type TMemoryStream struct {
	TCustomMemoryStream
}

func NewMemoryStream() IMemoryStream {
	r1 := LCL().SysCallN(3591)
	return AsMemoryStream(r1)
}

func MemoryStreamClass() TClass {
	ret := LCL().SysCallN(3589)
	return TClass(ret)
}

func (m *TMemoryStream) Clear() {
	LCL().SysCallN(3590, m.Instance())
}

func (m *TMemoryStream) LoadFromStream(Stream IStream) {
	LCL().SysCallN(3593, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TMemoryStream) LoadFromFile(FileName string) {
	LCL().SysCallN(3592, m.Instance(), PascalStr(FileName))
}
