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

// ICustomImageListResolution Parent: ILCLReferenceComponent
type ICustomImageListResolution interface {
	ILCLReferenceComponent
	ImageList() ICustomImageList                                                                                                                                // property
	Width() int32                                                                                                                                               // property
	Height() int32                                                                                                                                              // property
	Count() int32                                                                                                                                               // property
	AutoCreatedInDesignTime() bool                                                                                                                              // property
	SetAutoCreatedInDesignTime(AValue bool)                                                                                                                     // property
	GetHotSpot() (resultPoint TPoint)                                                                                                                           // function
	GetBitmap(Index int32, Image ICustomBitmap)                                                                                                                 // procedure
	GetBitmap1(Index int32, Image ICustomBitmap, AEffect TGraphicsDrawEffect)                                                                                   // procedure
	GetIcon(Index int32, Image IIcon, AEffect TGraphicsDrawEffect)                                                                                              // procedure
	GetIcon1(Index int32, Image IIcon)                                                                                                                          // procedure
	GetFullBitmap(Image ICustomBitmap, AEffect TGraphicsDrawEffect)                                                                                             // procedure
	GetRawImage(Index int32, OutImage *TRawImage)                                                                                                               // procedure
	Draw(ACanvas ICanvas, AX, AY, AIndex int32, AEnabled bool)                                                                                                  // procedure
	Draw1(ACanvas ICanvas, AX, AY, AIndex int32, ADrawEffect TGraphicsDrawEffect)                                                                               // procedure
	Draw2(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, AEnabled bool)                                             // procedure
	Draw3(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect)                           // procedure
	StretchDraw(Canvas ICanvas, Index int32, ARect *TRect, Enabled bool)                                                                                        // procedure
	DrawOverlay(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, AEnabled bool)                                                                        // procedure
	DrawOverlay1(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawEffect TGraphicsDrawEffect)                                                     // procedure
	DrawOverlay2(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect) // procedure
}

// TCustomImageListResolution Parent: TLCLReferenceComponent
type TCustomImageListResolution struct {
	TLCLReferenceComponent
}

func NewCustomImageListResolution(TheOwner IComponent) ICustomImageListResolution {
	r1 := LCL().SysCallN(1822, GetObjectUintptr(TheOwner))
	return AsCustomImageListResolution(r1)
}

func (m *TCustomImageListResolution) ImageList() ICustomImageList {
	r1 := LCL().SysCallN(1838, m.Instance())
	return AsCustomImageList(r1)
}

func (m *TCustomImageListResolution) Width() int32 {
	r1 := LCL().SysCallN(1840, m.Instance())
	return int32(r1)
}

func (m *TCustomImageListResolution) Height() int32 {
	r1 := LCL().SysCallN(1837, m.Instance())
	return int32(r1)
}

func (m *TCustomImageListResolution) Count() int32 {
	r1 := LCL().SysCallN(1821, m.Instance())
	return int32(r1)
}

func (m *TCustomImageListResolution) AutoCreatedInDesignTime() bool {
	r1 := LCL().SysCallN(1819, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImageListResolution) SetAutoCreatedInDesignTime(AValue bool) {
	LCL().SysCallN(1819, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImageListResolution) GetHotSpot() (resultPoint TPoint) {
	LCL().SysCallN(1833, m.Instance(), uintptr(unsafePointer(&resultPoint)))
	return
}

func CustomImageListResolutionClass() TClass {
	ret := LCL().SysCallN(1820)
	return TClass(ret)
}

func (m *TCustomImageListResolution) GetBitmap(Index int32, Image ICustomBitmap) {
	LCL().SysCallN(1830, m.Instance(), uintptr(Index), GetObjectUintptr(Image))
}

func (m *TCustomImageListResolution) GetBitmap1(Index int32, Image ICustomBitmap, AEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1831, m.Instance(), uintptr(Index), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageListResolution) GetIcon(Index int32, Image IIcon, AEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1834, m.Instance(), uintptr(Index), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageListResolution) GetIcon1(Index int32, Image IIcon) {
	LCL().SysCallN(1835, m.Instance(), uintptr(Index), GetObjectUintptr(Image))
}

func (m *TCustomImageListResolution) GetFullBitmap(Image ICustomBitmap, AEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1832, m.Instance(), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageListResolution) GetRawImage(Index int32, OutImage *TRawImage) {
	var result1 uintptr
	LCL().SysCallN(1836, m.Instance(), uintptr(Index), uintptr(unsafePointer(&result1)))
	*OutImage = *(*TRawImage)(getPointer(result1))
}

func (m *TCustomImageListResolution) Draw(ACanvas ICanvas, AX, AY, AIndex int32, AEnabled bool) {
	LCL().SysCallN(1823, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), PascalBool(AEnabled))
}

func (m *TCustomImageListResolution) Draw1(ACanvas ICanvas, AX, AY, AIndex int32, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1824, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawEffect))
}

func (m *TCustomImageListResolution) Draw2(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, AEnabled bool) {
	LCL().SysCallN(1825, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawingStyle), uintptr(AImageType), PascalBool(AEnabled))
}

func (m *TCustomImageListResolution) Draw3(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1826, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawingStyle), uintptr(AImageType), uintptr(ADrawEffect))
}

func (m *TCustomImageListResolution) StretchDraw(Canvas ICanvas, Index int32, ARect *TRect, Enabled bool) {
	LCL().SysCallN(1839, m.Instance(), GetObjectUintptr(Canvas), uintptr(Index), uintptr(unsafePointer(ARect)), PascalBool(Enabled))
}

func (m *TCustomImageListResolution) DrawOverlay(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, AEnabled bool) {
	LCL().SysCallN(1827, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), PascalBool(AEnabled))
}

func (m *TCustomImageListResolution) DrawOverlay1(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1828, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), uintptr(ADrawEffect))
}

func (m *TCustomImageListResolution) DrawOverlay2(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1829, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), uintptr(ADrawingStyle), uintptr(AImageType), uintptr(ADrawEffect))
}
