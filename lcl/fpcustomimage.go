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

// IFPCustomImage Is Abstract Class Parent: IPersistent
type IFPCustomImage interface {
	IPersistent
	Height() int32                                               // property
	SetHeight(AValue int32)                                      // property
	Width() int32                                                // property
	SetWidth(AValue int32)                                       // property
	Colors(x, y int32) (resultFPColor TFPColor)                  // property
	SetColors(x, y int32, AValue *TFPColor)                      // property
	UsePalette() bool                                            // property
	SetUsePalette(AValue bool)                                   // property
	Palette() IFPPalette                                         // property
	Pixels(x, y int32) int32                                     // property
	SetPixels(x, y int32, AValue int32)                          // property
	Extra(key string) string                                     // property
	SetExtra(key string, AValue string)                          // property
	ExtraValue(index int32) string                               // property
	SetExtraValue(index int32, AValue string)                    // property
	ExtraKey(index int32) string                                 // property
	SetExtraKey(index int32, AValue string)                      // property
	LoadFromFile(filename string) bool                           // function
	SaveToFile(filename string) bool                             // function
	ExtraCount() int32                                           // function
	LoadFromStream(Str IStream, Handler IFPCustomImageReader)    // procedure
	LoadFromStream1(Str IStream)                                 // procedure
	LoadFromFile1(filename string, Handler IFPCustomImageReader) // procedure
	SaveToStream(Str IStream, Handler IFPCustomImageWriter)      // procedure
	SaveToFile1(filename string, Handler IFPCustomImageWriter)   // procedure
	SetSize(AWidth, AHeight int32)                               // procedure
	RemoveExtra(key string)                                      // procedure
	SetOnProgress(fn TFPImgProgressEvent)                        // property event
}

// TFPCustomImage Is Abstract Class Parent: TPersistent
type TFPCustomImage struct {
	TPersistent
	progressPtr uintptr
}

func (m *TFPCustomImage) Height() int32 {
	r1 := LCL().SysCallN(2924, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomImage) SetHeight(AValue int32) {
	LCL().SysCallN(2924, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomImage) Width() int32 {
	r1 := LCL().SysCallN(2938, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomImage) SetWidth(AValue int32) {
	LCL().SysCallN(2938, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomImage) Colors(x, y int32) (resultFPColor TFPColor) {
	r1 := LCL().SysCallN(2919, 0, m.Instance(), uintptr(x), uintptr(y))
	return *(*TFPColor)(getPointer(r1))
}

func (m *TFPCustomImage) SetColors(x, y int32, AValue *TFPColor) {
	LCL().SysCallN(2919, 1, m.Instance(), uintptr(x), uintptr(y), uintptr(unsafePointer(AValue)))
}

func (m *TFPCustomImage) UsePalette() bool {
	r1 := LCL().SysCallN(2937, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomImage) SetUsePalette(AValue bool) {
	LCL().SysCallN(2937, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomImage) Palette() IFPPalette {
	r1 := LCL().SysCallN(2929, m.Instance())
	return AsFPPalette(r1)
}

func (m *TFPCustomImage) Pixels(x, y int32) int32 {
	r1 := LCL().SysCallN(2930, 0, m.Instance(), uintptr(x), uintptr(y))
	return int32(r1)
}

func (m *TFPCustomImage) SetPixels(x, y int32, AValue int32) {
	LCL().SysCallN(2930, 1, m.Instance(), uintptr(x), uintptr(y), uintptr(AValue))
}

func (m *TFPCustomImage) Extra(key string) string {
	r1 := LCL().SysCallN(2920, 0, m.Instance(), PascalStr(key))
	return GoStr(r1)
}

func (m *TFPCustomImage) SetExtra(key string, AValue string) {
	LCL().SysCallN(2920, 1, m.Instance(), PascalStr(key), PascalStr(AValue))
}

func (m *TFPCustomImage) ExtraValue(index int32) string {
	r1 := LCL().SysCallN(2923, 0, m.Instance(), uintptr(index))
	return GoStr(r1)
}

func (m *TFPCustomImage) SetExtraValue(index int32, AValue string) {
	LCL().SysCallN(2923, 1, m.Instance(), uintptr(index), PascalStr(AValue))
}

func (m *TFPCustomImage) ExtraKey(index int32) string {
	r1 := LCL().SysCallN(2922, 0, m.Instance(), uintptr(index))
	return GoStr(r1)
}

func (m *TFPCustomImage) SetExtraKey(index int32, AValue string) {
	LCL().SysCallN(2922, 1, m.Instance(), uintptr(index), PascalStr(AValue))
}

func (m *TFPCustomImage) LoadFromFile(filename string) bool {
	r1 := LCL().SysCallN(2925, m.Instance(), PascalStr(filename))
	return GoBool(r1)
}

func (m *TFPCustomImage) SaveToFile(filename string) bool {
	r1 := LCL().SysCallN(2932, m.Instance(), PascalStr(filename))
	return GoBool(r1)
}

func (m *TFPCustomImage) ExtraCount() int32 {
	r1 := LCL().SysCallN(2921, m.Instance())
	return int32(r1)
}

func FPCustomImageClass() TClass {
	ret := LCL().SysCallN(2918)
	return TClass(ret)
}

func (m *TFPCustomImage) LoadFromStream(Str IStream, Handler IFPCustomImageReader) {
	LCL().SysCallN(2927, m.Instance(), GetObjectUintptr(Str), GetObjectUintptr(Handler))
}

func (m *TFPCustomImage) LoadFromStream1(Str IStream) {
	LCL().SysCallN(2928, m.Instance(), GetObjectUintptr(Str))
}

func (m *TFPCustomImage) LoadFromFile1(filename string, Handler IFPCustomImageReader) {
	LCL().SysCallN(2926, m.Instance(), PascalStr(filename), GetObjectUintptr(Handler))
}

func (m *TFPCustomImage) SaveToStream(Str IStream, Handler IFPCustomImageWriter) {
	LCL().SysCallN(2934, m.Instance(), GetObjectUintptr(Str), GetObjectUintptr(Handler))
}

func (m *TFPCustomImage) SaveToFile1(filename string, Handler IFPCustomImageWriter) {
	LCL().SysCallN(2933, m.Instance(), PascalStr(filename), GetObjectUintptr(Handler))
}

func (m *TFPCustomImage) SetSize(AWidth, AHeight int32) {
	LCL().SysCallN(2936, m.Instance(), uintptr(AWidth), uintptr(AHeight))
}

func (m *TFPCustomImage) RemoveExtra(key string) {
	LCL().SysCallN(2931, m.Instance(), PascalStr(key))
}

func (m *TFPCustomImage) SetOnProgress(fn TFPImgProgressEvent) {
	if m.progressPtr != 0 {
		RemoveEventElement(m.progressPtr)
	}
	m.progressPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2935, m.Instance(), m.progressPtr)
}
