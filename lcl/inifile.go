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

// IIniFile Parent: ICustomIniFile
type IIniFile interface {
	ICustomIniFile
	Stream() IStream                                 // property
	CacheUpdates() bool                              // property
	SetCacheUpdates(AValue bool)                     // property
	WriteBOM() bool                                  // property
	SetWriteBOM(AValue bool)                         // property
	ReadSectionRaw(Section string, Strings IStrings) // procedure
}

// TIniFile Parent: TCustomIniFile
type TIniFile struct {
	TCustomIniFile
}

func NewIniFile(AFileName string, AOptions TIniFileOptions) IIniFile {
	r1 := LCL().SysCallN(3154, PascalStr(AFileName), uintptr(AOptions))
	return AsIniFile(r1)
}

func NewIniFile1(AStream IStream, AOptions TIniFileOptions) IIniFile {
	r1 := LCL().SysCallN(3155, GetObjectUintptr(AStream), uintptr(AOptions))
	return AsIniFile(r1)
}

func (m *TIniFile) Stream() IStream {
	r1 := LCL().SysCallN(3157, m.Instance())
	return AsStream(r1)
}

func (m *TIniFile) CacheUpdates() bool {
	r1 := LCL().SysCallN(3152, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TIniFile) SetCacheUpdates(AValue bool) {
	LCL().SysCallN(3152, 1, m.Instance(), PascalBool(AValue))
}

func (m *TIniFile) WriteBOM() bool {
	r1 := LCL().SysCallN(3158, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TIniFile) SetWriteBOM(AValue bool) {
	LCL().SysCallN(3158, 1, m.Instance(), PascalBool(AValue))
}

func IniFileClass() TClass {
	ret := LCL().SysCallN(3153)
	return TClass(ret)
}

func (m *TIniFile) ReadSectionRaw(Section string, Strings IStrings) {
	LCL().SysCallN(3156, m.Instance(), PascalStr(Section), GetObjectUintptr(Strings))
}
