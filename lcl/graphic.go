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

// IGraphic Is Abstract Class Parent: IPersistent
type IGraphic interface {
	IPersistent
	LoadFromBytes(data []byte)
	Empty() bool                                                                       // property
	Height() int32                                                                     // property
	SetHeight(AValue int32)                                                            // property
	Modified() bool                                                                    // property
	SetModified(AValue bool)                                                           // property
	MimeType() string                                                                  // property
	Palette() HPALETTE                                                                 // property
	SetPalette(AValue HPALETTE)                                                        // property
	PaletteModified() bool                                                             // property
	SetPaletteModified(AValue bool)                                                    // property
	Transparent() bool                                                                 // property
	SetTransparent(AValue bool)                                                        // property
	Width() int32                                                                      // property
	SetWidth(AValue int32)                                                             // property
	LazarusResourceTypeValid(AResourceType string) bool                                // function
	GetResourceType() TResourceType                                                    // function
	Clear()                                                                            // procedure
	LoadFromFile(Filename string)                                                      // procedure
	LoadFromStream(Stream IStream)                                                     // procedure Is Abstract
	LoadFromMimeStream(AStream IStream, AMimeType string)                              // procedure
	LoadFromLazarusResource(ResName string)                                            // procedure
	LoadFromResourceName(Instance THandle, ResName string)                             // procedure
	LoadFromResourceID(Instance THandle, ResID uint32)                                 // procedure
	LoadFromClipboardFormat(FormatID TClipboardFormat)                                 // procedure
	LoadFromClipboardFormatID(ClipboardType TClipboardType, FormatID TClipboardFormat) // procedure
	SaveToFile(Filename string)                                                        // procedure
	SaveToStream(Stream IStream)                                                       // procedure Is Abstract
	SaveToClipboardFormat(FormatID TClipboardFormat)                                   // procedure
	SaveToClipboardFormatID(ClipboardType TClipboardType, FormatID TClipboardFormat)   // procedure
	GetSupportedSourceMimeTypes(List IStrings)                                         // procedure
	SetOnChange(fn TNotifyEvent)                                                       // property event
	SetOnProgress(fn TProgressEvent)                                                   // property event
}

// TGraphic Is Abstract Class Parent: TPersistent
type TGraphic struct {
	TPersistent
	changePtr   uintptr
	progressPtr uintptr
}

func (m *TGraphic) Empty() bool {
	r1 := LCL().SysCallN(2934, m.Instance())
	return GoBool(r1)
}

func (m *TGraphic) Height() int32 {
	r1 := LCL().SysCallN(2937, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGraphic) SetHeight(AValue int32) {
	LCL().SysCallN(2937, 1, m.Instance(), uintptr(AValue))
}

func (m *TGraphic) Modified() bool {
	r1 := LCL().SysCallN(2948, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGraphic) SetModified(AValue bool) {
	LCL().SysCallN(2948, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGraphic) MimeType() string {
	r1 := LCL().SysCallN(2947, m.Instance())
	return GoStr(r1)
}

func (m *TGraphic) Palette() HPALETTE {
	r1 := LCL().SysCallN(2949, 0, m.Instance(), 0)
	return HPALETTE(r1)
}

func (m *TGraphic) SetPalette(AValue HPALETTE) {
	LCL().SysCallN(2949, 1, m.Instance(), uintptr(AValue))
}

func (m *TGraphic) PaletteModified() bool {
	r1 := LCL().SysCallN(2950, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGraphic) SetPaletteModified(AValue bool) {
	LCL().SysCallN(2950, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGraphic) Transparent() bool {
	r1 := LCL().SysCallN(2957, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGraphic) SetTransparent(AValue bool) {
	LCL().SysCallN(2957, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGraphic) Width() int32 {
	r1 := LCL().SysCallN(2958, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGraphic) SetWidth(AValue int32) {
	LCL().SysCallN(2958, 1, m.Instance(), uintptr(AValue))
}

func (m *TGraphic) LazarusResourceTypeValid(AResourceType string) bool {
	r1 := LCL().SysCallN(2938, m.Instance(), PascalStr(AResourceType))
	return GoBool(r1)
}

func (m *TGraphic) GetResourceType() TResourceType {
	r1 := LCL().SysCallN(2935, m.Instance())
	return TResourceType(r1)
}

func GraphicClass() TClass {
	ret := LCL().SysCallN(2932)
	return TClass(ret)
}

func (m *TGraphic) Clear() {
	LCL().SysCallN(2933, m.Instance())
}

func (m *TGraphic) LoadFromFile(Filename string) {
	LCL().SysCallN(2941, m.Instance(), PascalStr(Filename))
}

func (m *TGraphic) LoadFromStream(Stream IStream) {
	LCL().SysCallN(2946, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TGraphic) LoadFromMimeStream(AStream IStream, AMimeType string) {
	LCL().SysCallN(2943, m.Instance(), GetObjectUintptr(AStream), PascalStr(AMimeType))
}

func (m *TGraphic) LoadFromLazarusResource(ResName string) {
	LCL().SysCallN(2942, m.Instance(), PascalStr(ResName))
}

func (m *TGraphic) LoadFromResourceName(Instance THandle, ResName string) {
	LCL().SysCallN(2945, m.Instance(), uintptr(Instance), PascalStr(ResName))
}

func (m *TGraphic) LoadFromResourceID(Instance THandle, ResID uint32) {
	LCL().SysCallN(2944, m.Instance(), uintptr(Instance), uintptr(ResID))
}

func (m *TGraphic) LoadFromClipboardFormat(FormatID TClipboardFormat) {
	LCL().SysCallN(2939, m.Instance(), uintptr(FormatID))
}

func (m *TGraphic) LoadFromClipboardFormatID(ClipboardType TClipboardType, FormatID TClipboardFormat) {
	LCL().SysCallN(2940, m.Instance(), uintptr(ClipboardType), uintptr(FormatID))
}

func (m *TGraphic) SaveToFile(Filename string) {
	LCL().SysCallN(2953, m.Instance(), PascalStr(Filename))
}

func (m *TGraphic) SaveToStream(Stream IStream) {
	LCL().SysCallN(2954, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TGraphic) SaveToClipboardFormat(FormatID TClipboardFormat) {
	LCL().SysCallN(2951, m.Instance(), uintptr(FormatID))
}

func (m *TGraphic) SaveToClipboardFormatID(ClipboardType TClipboardType, FormatID TClipboardFormat) {
	LCL().SysCallN(2952, m.Instance(), uintptr(ClipboardType), uintptr(FormatID))
}

func (m *TGraphic) GetSupportedSourceMimeTypes(List IStrings) {
	LCL().SysCallN(2936, m.Instance(), GetObjectUintptr(List))
}

func (m *TGraphic) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2955, m.Instance(), m.changePtr)
}

func (m *TGraphic) SetOnProgress(fn TProgressEvent) {
	if m.progressPtr != 0 {
		RemoveEventElement(m.progressPtr)
	}
	m.progressPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2956, m.Instance(), m.progressPtr)
}
