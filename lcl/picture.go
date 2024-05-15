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
	r1 := LCL().SysCallN(4530)
	return AsPicture(r1)
}

func (m *TPicture) Bitmap() IBitmap {
	r1 := LCL().SysCallN(4527, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TPicture) SetBitmap(AValue IBitmap) {
	LCL().SysCallN(4527, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) Icon() IIcon {
	r1 := LCL().SysCallN(4533, 0, m.Instance(), 0)
	return AsIcon(r1)
}

func (m *TPicture) SetIcon(AValue IIcon) {
	LCL().SysCallN(4533, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) Jpeg() IJPEGImage {
	r1 := LCL().SysCallN(4534, 0, m.Instance(), 0)
	return AsJPEGImage(r1)
}

func (m *TPicture) SetJpeg(AValue IJPEGImage) {
	LCL().SysCallN(4534, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) Pixmap() IPixmap {
	r1 := LCL().SysCallN(4545, 0, m.Instance(), 0)
	return AsPixmap(r1)
}

func (m *TPicture) SetPixmap(AValue IPixmap) {
	LCL().SysCallN(4545, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) PNG() IPortableNetworkGraphic {
	r1 := LCL().SysCallN(4543, 0, m.Instance(), 0)
	return AsPortableNetworkGraphic(r1)
}

func (m *TPicture) SetPNG(AValue IPortableNetworkGraphic) {
	LCL().SysCallN(4543, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) PNM() IPortableAnyMapGraphic {
	r1 := LCL().SysCallN(4544, 0, m.Instance(), 0)
	return AsPortableAnyMapGraphic(r1)
}

func (m *TPicture) SetPNM(AValue IPortableAnyMapGraphic) {
	LCL().SysCallN(4544, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) Graphic() IGraphic {
	r1 := LCL().SysCallN(4531, 0, m.Instance(), 0)
	return AsGraphic(r1)
}

func (m *TPicture) SetGraphic(AValue IGraphic) {
	LCL().SysCallN(4531, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) Height() int32 {
	r1 := LCL().SysCallN(4532, m.Instance())
	return int32(r1)
}

func (m *TPicture) Width() int32 {
	r1 := LCL().SysCallN(4552, m.Instance())
	return int32(r1)
}

func PictureClass() TClass {
	ret := LCL().SysCallN(4528)
	return TClass(ret)
}

func (m *TPicture) Clear() {
	LCL().SysCallN(4529, m.Instance())
}

func (m *TPicture) LoadFromClipboardFormat(FormatID TClipboardFormat) {
	LCL().SysCallN(4535, m.Instance(), uintptr(FormatID))
}

func (m *TPicture) LoadFromClipboardFormatID(ClipboardType TClipboardType, FormatID TClipboardFormat) {
	LCL().SysCallN(4536, m.Instance(), uintptr(ClipboardType), uintptr(FormatID))
}

func (m *TPicture) LoadFromFile(Filename string) {
	LCL().SysCallN(4537, m.Instance(), PascalStr(Filename))
}

func (m *TPicture) LoadFromResourceName(Instance THandle, ResName string) {
	LCL().SysCallN(4539, m.Instance(), uintptr(Instance), PascalStr(ResName))
}

func (m *TPicture) LoadFromResourceName1(Instance THandle, ResName string, AClass TGraphicClass) {
	LCL().SysCallN(4540, m.Instance(), uintptr(Instance), PascalStr(ResName), uintptr(AClass))
}

func (m *TPicture) LoadFromLazarusResource(AName string) {
	LCL().SysCallN(4538, m.Instance(), PascalStr(AName))
}

func (m *TPicture) LoadFromStream(Stream IStream) {
	LCL().SysCallN(4541, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TPicture) LoadFromStreamWithFileExt(Stream IStream, FileExt string) {
	LCL().SysCallN(4542, m.Instance(), GetObjectUintptr(Stream), PascalStr(FileExt))
}

func (m *TPicture) SaveToClipboardFormat(FormatID TClipboardFormat) {
	LCL().SysCallN(4546, m.Instance(), uintptr(FormatID))
}

func (m *TPicture) SaveToFile(Filename string, FileExt string) {
	LCL().SysCallN(4547, m.Instance(), PascalStr(Filename), PascalStr(FileExt))
}

func (m *TPicture) SaveToStream(Stream IStream) {
	LCL().SysCallN(4548, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TPicture) SaveToStreamWithFileExt(Stream IStream, FileExt string) {
	LCL().SysCallN(4549, m.Instance(), GetObjectUintptr(Stream), PascalStr(FileExt))
}

func (m *TPicture) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4550, m.Instance(), m.changePtr)
}

func (m *TPicture) SetOnProgress(fn TProgressEvent) {
	if m.progressPtr != 0 {
		RemoveEventElement(m.progressPtr)
	}
	m.progressPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4551, m.Instance(), m.progressPtr)
}
