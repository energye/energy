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

// IPicture Parent: IPersistent
type IPicture interface {
	IPersistent
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
	Bitmap() IBitmap                                                                   // property
	SetBitmap(AValue IBitmap)                                                          // property
	Icon() IIcon                                                                       // property
	SetIcon(AValue IIcon)                                                              // property
	Jpeg() IJPEGImage                                                                  // property
	SetJpeg(AValue IJPEGImage)                                                         // property
	Pixmap() IPixmap                                                                   // property
	SetPixmap(AValue IPixmap)                                                          // property
	PNG() IPortableNetworkGraphic                                                      // property
	SetPNG(AValue IPortableNetworkGraphic)                                             // property
	PNM() IPortableAnyMapGraphic                                                       // property
	SetPNM(AValue IPortableAnyMapGraphic)                                              // property
	Graphic() IGraphic                                                                 // property
	SetGraphic(AValue IGraphic)                                                        // property
	Height() int32                                                                     // property
	Width() int32                                                                      // property
	Clear()                                                                            // procedure
	LoadFromClipboardFormat(FormatID TClipboardFormat)                                 // procedure
	LoadFromClipboardFormatID(ClipboardType TClipboardType, FormatID TClipboardFormat) // procedure
	LoadFromFile(Filename string)                                                      // procedure
	LoadFromResourceName(Instance THandle, ResName string)                             // procedure
	LoadFromResourceName1(Instance THandle, ResName string, AClass TGraphicClass)      // procedure
	LoadFromLazarusResource(AName string)                                              // procedure
	LoadFromStream(Stream IStream)                                                     // procedure
	LoadFromStreamWithFileExt(Stream IStream, FileExt string)                          // procedure
	SaveToClipboardFormat(FormatID TClipboardFormat)                                   // procedure
	SaveToFile(Filename string, FileExt string)                                        // procedure
	SaveToStream(Stream IStream)                                                       // procedure
	SaveToStreamWithFileExt(Stream IStream, FileExt string)                            // procedure
	SetOnChange(fn TNotifyEvent)                                                       // property event
	SetOnProgress(fn TProgressEvent)                                                   // property event
}

// TPicture Parent: TPersistent
type TPicture struct {
	TPersistent
	changePtr   uintptr
	progressPtr uintptr
}

func NewPicture() IPicture {
	r1 := LCL().SysCallN(3888)
	return AsPicture(r1)
}

func (m *TPicture) Bitmap() IBitmap {
	r1 := LCL().SysCallN(3885, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TPicture) SetBitmap(AValue IBitmap) {
	LCL().SysCallN(3885, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) Icon() IIcon {
	r1 := LCL().SysCallN(3891, 0, m.Instance(), 0)
	return AsIcon(r1)
}

func (m *TPicture) SetIcon(AValue IIcon) {
	LCL().SysCallN(3891, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) Jpeg() IJPEGImage {
	r1 := LCL().SysCallN(3892, 0, m.Instance(), 0)
	return AsJPEGImage(r1)
}

func (m *TPicture) SetJpeg(AValue IJPEGImage) {
	LCL().SysCallN(3892, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) Pixmap() IPixmap {
	r1 := LCL().SysCallN(3903, 0, m.Instance(), 0)
	return AsPixmap(r1)
}

func (m *TPicture) SetPixmap(AValue IPixmap) {
	LCL().SysCallN(3903, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) PNG() IPortableNetworkGraphic {
	r1 := LCL().SysCallN(3901, 0, m.Instance(), 0)
	return AsPortableNetworkGraphic(r1)
}

func (m *TPicture) SetPNG(AValue IPortableNetworkGraphic) {
	LCL().SysCallN(3901, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) PNM() IPortableAnyMapGraphic {
	r1 := LCL().SysCallN(3902, 0, m.Instance(), 0)
	return AsPortableAnyMapGraphic(r1)
}

func (m *TPicture) SetPNM(AValue IPortableAnyMapGraphic) {
	LCL().SysCallN(3902, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) Graphic() IGraphic {
	r1 := LCL().SysCallN(3889, 0, m.Instance(), 0)
	return AsGraphic(r1)
}

func (m *TPicture) SetGraphic(AValue IGraphic) {
	LCL().SysCallN(3889, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) Height() int32 {
	r1 := LCL().SysCallN(3890, m.Instance())
	return int32(r1)
}

func (m *TPicture) Width() int32 {
	r1 := LCL().SysCallN(3910, m.Instance())
	return int32(r1)
}

func PictureClass() TClass {
	ret := LCL().SysCallN(3886)
	return TClass(ret)
}

func (m *TPicture) Clear() {
	LCL().SysCallN(3887, m.Instance())
}

func (m *TPicture) LoadFromClipboardFormat(FormatID TClipboardFormat) {
	LCL().SysCallN(3893, m.Instance(), uintptr(FormatID))
}

func (m *TPicture) LoadFromClipboardFormatID(ClipboardType TClipboardType, FormatID TClipboardFormat) {
	LCL().SysCallN(3894, m.Instance(), uintptr(ClipboardType), uintptr(FormatID))
}

func (m *TPicture) LoadFromFile(Filename string) {
	LCL().SysCallN(3895, m.Instance(), PascalStr(Filename))
}

func (m *TPicture) LoadFromResourceName(Instance THandle, ResName string) {
	LCL().SysCallN(3897, m.Instance(), uintptr(Instance), PascalStr(ResName))
}

func (m *TPicture) LoadFromResourceName1(Instance THandle, ResName string, AClass TGraphicClass) {
	LCL().SysCallN(3898, m.Instance(), uintptr(Instance), PascalStr(ResName), uintptr(AClass))
}

func (m *TPicture) LoadFromLazarusResource(AName string) {
	LCL().SysCallN(3896, m.Instance(), PascalStr(AName))
}

func (m *TPicture) LoadFromStream(Stream IStream) {
	LCL().SysCallN(3899, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TPicture) LoadFromStreamWithFileExt(Stream IStream, FileExt string) {
	LCL().SysCallN(3900, m.Instance(), GetObjectUintptr(Stream), PascalStr(FileExt))
}

func (m *TPicture) SaveToClipboardFormat(FormatID TClipboardFormat) {
	LCL().SysCallN(3904, m.Instance(), uintptr(FormatID))
}

func (m *TPicture) SaveToFile(Filename string, FileExt string) {
	LCL().SysCallN(3905, m.Instance(), PascalStr(Filename), PascalStr(FileExt))
}

func (m *TPicture) SaveToStream(Stream IStream) {
	LCL().SysCallN(3906, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TPicture) SaveToStreamWithFileExt(Stream IStream, FileExt string) {
	LCL().SysCallN(3907, m.Instance(), GetObjectUintptr(Stream), PascalStr(FileExt))
}

func (m *TPicture) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3908, m.Instance(), m.changePtr)
}

func (m *TPicture) SetOnProgress(fn TProgressEvent) {
	if m.progressPtr != 0 {
		RemoveEventElement(m.progressPtr)
	}
	m.progressPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3909, m.Instance(), m.progressPtr)
}
