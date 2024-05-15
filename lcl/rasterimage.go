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

// IRasterImage Is Abstract Class Parent: IGraphic
type IRasterImage interface {
	IGraphic
	Canvas() ICanvas                                            // property
	BitmapHandle() HBITMAP                                      // property
	SetBitmapHandle(AValue HBITMAP)                             // property
	Masked() bool                                               // property
	SetMasked(AValue bool)                                      // property
	MaskHandle() HBITMAP                                        // property
	SetMaskHandle(AValue HBITMAP)                               // property
	PixelFormat() TPixelFormat                                  // property
	SetPixelFormat(AValue TPixelFormat)                         // property
	RawImage() (resultRawImage TRawImage)                       // property
	ScanLine(Row int32) uintptr                                 // property
	TransparentColor() TColor                                   // property
	SetTransparentColor(AValue TColor)                          // property
	TransparentMode() TTransparentMode                          // property
	SetTransparentMode(AValue TTransparentMode)                 // property
	BitmapHandleAllocated() bool                                // function Is Abstract
	MaskHandleAllocated() bool                                  // function Is Abstract
	PaletteAllocated() bool                                     // function Is Abstract
	ReleaseBitmapHandle() HBITMAP                               // function
	ReleaseMaskHandle() HBITMAP                                 // function
	ReleasePalette() HPALETTE                                   // function
	HandleAllocated() bool                                      // function
	BeginUpdate(ACanvasOnly bool)                               // procedure
	EndUpdate(AStreamIsValid bool)                              // procedure
	FreeImage()                                                 // procedure
	LoadFromBitmapHandles(ABitmap, AMask HBITMAP, ARect *TRect) // procedure
	LoadFromDevice(DC HDC)                                      // procedure
	LoadFromStreamForStream(AStream IStream, ASize uint32)      // procedure
	LoadFromRawImage(AIMage *TRawImage, ADataOwner bool)        // procedure
	GetSize(OutWidth, OutHeight *int32)                         // procedure
	Mask(ATransparentColor TColor)                              // procedure
	SetHandles(ABitmap, AMask HBITMAP)                          // procedure Is Abstract
}

// TRasterImage Is Abstract Class Parent: TGraphic
type TRasterImage struct {
	TGraphic
}

func (m *TRasterImage) Canvas() ICanvas {
	r1 := LCL().SysCallN(4713, m.Instance())
	return AsCanvas(r1)
}

func (m *TRasterImage) BitmapHandle() HBITMAP {
	r1 := LCL().SysCallN(4711, 0, m.Instance(), 0)
	return HBITMAP(r1)
}

func (m *TRasterImage) SetBitmapHandle(AValue HBITMAP) {
	LCL().SysCallN(4711, 1, m.Instance(), uintptr(AValue))
}

func (m *TRasterImage) Masked() bool {
	r1 := LCL().SysCallN(4726, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRasterImage) SetMasked(AValue bool) {
	LCL().SysCallN(4726, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRasterImage) MaskHandle() HBITMAP {
	r1 := LCL().SysCallN(4724, 0, m.Instance(), 0)
	return HBITMAP(r1)
}

func (m *TRasterImage) SetMaskHandle(AValue HBITMAP) {
	LCL().SysCallN(4724, 1, m.Instance(), uintptr(AValue))
}

func (m *TRasterImage) PixelFormat() TPixelFormat {
	r1 := LCL().SysCallN(4728, 0, m.Instance(), 0)
	return TPixelFormat(r1)
}

func (m *TRasterImage) SetPixelFormat(AValue TPixelFormat) {
	LCL().SysCallN(4728, 1, m.Instance(), uintptr(AValue))
}

func (m *TRasterImage) RawImage() (resultRawImage TRawImage) {
	r1 := LCL().SysCallN(4729, m.Instance())
	return *(*TRawImage)(getPointer(r1))
}

func (m *TRasterImage) ScanLine(Row int32) uintptr {
	r1 := LCL().SysCallN(4733, m.Instance(), uintptr(Row))
	return uintptr(r1)
}

func (m *TRasterImage) TransparentColor() TColor {
	r1 := LCL().SysCallN(4735, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TRasterImage) SetTransparentColor(AValue TColor) {
	LCL().SysCallN(4735, 1, m.Instance(), uintptr(AValue))
}

func (m *TRasterImage) TransparentMode() TTransparentMode {
	r1 := LCL().SysCallN(4736, 0, m.Instance(), 0)
	return TTransparentMode(r1)
}

func (m *TRasterImage) SetTransparentMode(AValue TTransparentMode) {
	LCL().SysCallN(4736, 1, m.Instance(), uintptr(AValue))
}

func (m *TRasterImage) BitmapHandleAllocated() bool {
	r1 := LCL().SysCallN(4712, m.Instance())
	return GoBool(r1)
}

func (m *TRasterImage) MaskHandleAllocated() bool {
	r1 := LCL().SysCallN(4725, m.Instance())
	return GoBool(r1)
}

func (m *TRasterImage) PaletteAllocated() bool {
	r1 := LCL().SysCallN(4727, m.Instance())
	return GoBool(r1)
}

func (m *TRasterImage) ReleaseBitmapHandle() HBITMAP {
	r1 := LCL().SysCallN(4730, m.Instance())
	return HBITMAP(r1)
}

func (m *TRasterImage) ReleaseMaskHandle() HBITMAP {
	r1 := LCL().SysCallN(4731, m.Instance())
	return HBITMAP(r1)
}

func (m *TRasterImage) ReleasePalette() HPALETTE {
	r1 := LCL().SysCallN(4732, m.Instance())
	return HPALETTE(r1)
}

func (m *TRasterImage) HandleAllocated() bool {
	r1 := LCL().SysCallN(4718, m.Instance())
	return GoBool(r1)
}

func RasterImageClass() TClass {
	ret := LCL().SysCallN(4714)
	return TClass(ret)
}

func (m *TRasterImage) BeginUpdate(ACanvasOnly bool) {
	LCL().SysCallN(4710, m.Instance(), PascalBool(ACanvasOnly))
}

func (m *TRasterImage) EndUpdate(AStreamIsValid bool) {
	LCL().SysCallN(4715, m.Instance(), PascalBool(AStreamIsValid))
}

func (m *TRasterImage) FreeImage() {
	LCL().SysCallN(4716, m.Instance())
}

func (m *TRasterImage) LoadFromBitmapHandles(ABitmap, AMask HBITMAP, ARect *TRect) {
	LCL().SysCallN(4719, m.Instance(), uintptr(ABitmap), uintptr(AMask), uintptr(unsafePointer(ARect)))
}

func (m *TRasterImage) LoadFromDevice(DC HDC) {
	LCL().SysCallN(4720, m.Instance(), uintptr(DC))
}

func (m *TRasterImage) LoadFromStreamForStream(AStream IStream, ASize uint32) {
	LCL().SysCallN(4722, m.Instance(), GetObjectUintptr(AStream), uintptr(ASize))
}

func (m *TRasterImage) LoadFromRawImage(AIMage *TRawImage, ADataOwner bool) {
	LCL().SysCallN(4721, m.Instance(), uintptr(unsafePointer(AIMage)), PascalBool(ADataOwner))
}

func (m *TRasterImage) GetSize(OutWidth, OutHeight *int32) {
	var result0 uintptr
	var result1 uintptr
	LCL().SysCallN(4717, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)))
	*OutWidth = int32(result0)
	*OutHeight = int32(result1)
}

func (m *TRasterImage) Mask(ATransparentColor TColor) {
	LCL().SysCallN(4723, m.Instance(), uintptr(ATransparentColor))
}

func (m *TRasterImage) SetHandles(ABitmap, AMask HBITMAP) {
	LCL().SysCallN(4734, m.Instance(), uintptr(ABitmap), uintptr(AMask))
}
