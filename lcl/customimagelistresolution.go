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
	"unsafe"
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
	r1 := LCL().SysCallN(1632, GetObjectUintptr(TheOwner))
	return AsCustomImageListResolution(r1)
}

func (m *TCustomImageListResolution) ImageList() ICustomImageList {
	r1 := LCL().SysCallN(1648, m.Instance())
	return AsCustomImageList(r1)
}

func (m *TCustomImageListResolution) Width() int32 {
	r1 := LCL().SysCallN(1650, m.Instance())
	return int32(r1)
}

func (m *TCustomImageListResolution) Height() int32 {
	r1 := LCL().SysCallN(1647, m.Instance())
	return int32(r1)
}

func (m *TCustomImageListResolution) Count() int32 {
	r1 := LCL().SysCallN(1631, m.Instance())
	return int32(r1)
}

func (m *TCustomImageListResolution) AutoCreatedInDesignTime() bool {
	r1 := LCL().SysCallN(1629, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImageListResolution) SetAutoCreatedInDesignTime(AValue bool) {
	LCL().SysCallN(1629, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImageListResolution) GetHotSpot() (resultPoint TPoint) {
	LCL().SysCallN(1643, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func CustomImageListResolutionClass() TClass {
	ret := LCL().SysCallN(1630)
	return TClass(ret)
}

func (m *TCustomImageListResolution) GetBitmap(Index int32, Image ICustomBitmap) {
	LCL().SysCallN(1640, m.Instance(), uintptr(Index), GetObjectUintptr(Image))
}

func (m *TCustomImageListResolution) GetBitmap1(Index int32, Image ICustomBitmap, AEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1641, m.Instance(), uintptr(Index), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageListResolution) GetIcon(Index int32, Image IIcon, AEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1644, m.Instance(), uintptr(Index), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageListResolution) GetIcon1(Index int32, Image IIcon) {
	LCL().SysCallN(1645, m.Instance(), uintptr(Index), GetObjectUintptr(Image))
}

func (m *TCustomImageListResolution) GetFullBitmap(Image ICustomBitmap, AEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1642, m.Instance(), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageListResolution) GetRawImage(Index int32, OutImage *TRawImage) {
	var result1 uintptr
	LCL().SysCallN(1646, m.Instance(), uintptr(Index), uintptr(unsafe.Pointer(&result1)))
	*OutImage = *(*TRawImage)(getPointer(result1))
}

func (m *TCustomImageListResolution) Draw(ACanvas ICanvas, AX, AY, AIndex int32, AEnabled bool) {
	LCL().SysCallN(1633, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), PascalBool(AEnabled))
}

func (m *TCustomImageListResolution) Draw1(ACanvas ICanvas, AX, AY, AIndex int32, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1634, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawEffect))
}

func (m *TCustomImageListResolution) Draw2(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, AEnabled bool) {
	LCL().SysCallN(1635, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawingStyle), uintptr(AImageType), PascalBool(AEnabled))
}

func (m *TCustomImageListResolution) Draw3(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1636, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawingStyle), uintptr(AImageType), uintptr(ADrawEffect))
}

func (m *TCustomImageListResolution) StretchDraw(Canvas ICanvas, Index int32, ARect *TRect, Enabled bool) {
	LCL().SysCallN(1649, m.Instance(), GetObjectUintptr(Canvas), uintptr(Index), uintptr(unsafe.Pointer(ARect)), PascalBool(Enabled))
}

func (m *TCustomImageListResolution) DrawOverlay(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, AEnabled bool) {
	LCL().SysCallN(1637, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), PascalBool(AEnabled))
}

func (m *TCustomImageListResolution) DrawOverlay1(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1638, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), uintptr(ADrawEffect))
}

func (m *TCustomImageListResolution) DrawOverlay2(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1639, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), uintptr(ADrawingStyle), uintptr(AImageType), uintptr(ADrawEffect))
}
