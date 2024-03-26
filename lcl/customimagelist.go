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

// ICustomImageList Parent: ILCLComponent
type ICustomImageList interface {
	ILCLComponent
	HasOverlays() bool                                                                                                                                          // property
	AllocBy() int32                                                                                                                                             // property
	SetAllocBy(AValue int32)                                                                                                                                    // property
	BlendColor() TColor                                                                                                                                         // property
	SetBlendColor(AValue TColor)                                                                                                                                // property
	BkColor() TColor                                                                                                                                            // property
	SetBkColor(AValue TColor)                                                                                                                                   // property
	Count() int32                                                                                                                                               // property
	DrawingStyle() TDrawingStyle                                                                                                                                // property
	SetDrawingStyle(AValue TDrawingStyle)                                                                                                                       // property
	Height() int32                                                                                                                                              // property
	SetHeight(AValue int32)                                                                                                                                     // property
	HeightForPPI(AImageWidth, APPI int32) int32                                                                                                                 // property
	HeightForWidth(AWidth int32) int32                                                                                                                          // property
	Width() int32                                                                                                                                               // property
	SetWidth(AValue int32)                                                                                                                                      // property
	WidthForPPI(AImageWidth, APPI int32) int32                                                                                                                  // property
	SizeForPPI(AImageWidth, APPI int32) (resultSize TSize)                                                                                                      // property
	Masked() bool                                                                                                                                               // property
	SetMasked(AValue bool)                                                                                                                                      // property
	Resolution(AImageWidth int32) ICustomImageListResolution                                                                                                    // property
	ResolutionByIndex(AIndex int32) ICustomImageListResolution                                                                                                  // property
	ResolutionForPPI(AImageWidth, APPI int32, ACanvasScaleFactor float64) (resultScaledImageListResolution TScaledImageListResolution)                          // property
	ResolutionCount() int32                                                                                                                                     // property
	Scaled() bool                                                                                                                                               // property
	SetScaled(AValue bool)                                                                                                                                      // property
	ShareImages() bool                                                                                                                                          // property
	SetShareImages(AValue bool)                                                                                                                                 // property
	ImageType() TImageType                                                                                                                                      // property
	SetImageType(AValue TImageType)                                                                                                                             // property
	Add(Image, Mask ICustomBitmap) int32                                                                                                                        // function
	AddSliced(Image ICustomBitmap, AHorizontalCount, AVerticalCount int32) int32                                                                                // function
	AddSlice(Image ICustomBitmap, AImageRect *TRect) int32                                                                                                      // function
	AddSliceCentered(Image ICustomBitmap) int32                                                                                                                 // function
	AddIcon(Image ICustomIcon) int32                                                                                                                            // function
	AddMasked(Image IBitmap, MaskColor TColor) int32                                                                                                            // function
	AddLazarusResource(ResourceName string, MaskColor TColor) int32                                                                                             // function
	AddResourceName(Instance THandle, ResourceName string, MaskColor TColor) int32                                                                              // function
	GetHotSpot() (resultPoint TPoint)                                                                                                                           // function
	FindResolution(AImageWidth int32, OutResolution *ICustomImageListResolution) bool                                                                           // function
	Resolutions() ICustomImageListResolutionEnumerator                                                                                                          // function
	ResolutionsDesc() ICustomImageListResolutionEnumerator                                                                                                      // function
	AssignTo(Dest IPersistent)                                                                                                                                  // procedure
	WriteData(AStream IStream)                                                                                                                                  // procedure
	ReadData(AStream IStream)                                                                                                                                   // procedure
	WriteAdvData(AStream IStream)                                                                                                                               // procedure
	ReadAdvData(AStream IStream)                                                                                                                                // procedure
	BeginUpdate()                                                                                                                                               // procedure
	EndUpdate()                                                                                                                                                 // procedure
	AddImages(AValue ICustomImageList)                                                                                                                          // procedure
	Change()                                                                                                                                                    // procedure
	Clear()                                                                                                                                                     // procedure
	Delete(AIndex int32)                                                                                                                                        // procedure
	Draw(ACanvas ICanvas, AX, AY, AIndex int32, AEnabled bool)                                                                                                  // procedure
	Draw1(ACanvas ICanvas, AX, AY, AIndex int32, ADrawEffect TGraphicsDrawEffect)                                                                               // procedure
	Draw2(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, AEnabled bool)                                             // procedure
	Draw3(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect)                           // procedure
	DrawForPPI(ACanvas ICanvas, AX, AY, AIndex int32, AImageWidthAt96PPI, ATargetPPI int32, ACanvasFactor float64, AEnabled bool)                               // procedure
	DrawForPPI1(ACanvas ICanvas, AX, AY, AIndex int32, AImageWidthAt96PPI, ATargetPPI int32, ACanvasFactor float64, ADrawEffect TGraphicsDrawEffect)            // procedure
	DrawOverlay(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, AEnabled bool)                                                                        // procedure
	DrawOverlay1(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawEffect TGraphicsDrawEffect)                                                     // procedure
	DrawOverlay2(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect) // procedure
	GetBitmap(Index int32, Image ICustomBitmap)                                                                                                                 // procedure
	GetBitmap1(Index int32, Image ICustomBitmap, AEffect TGraphicsDrawEffect)                                                                                   // procedure
	GetFullBitmap(Image ICustomBitmap, AEffect TGraphicsDrawEffect)                                                                                             // procedure
	GetFullRawImage(OutImage *TRawImage)                                                                                                                        // procedure
	GetIcon(Index int32, Image IIcon, AEffect TGraphicsDrawEffect)                                                                                              // procedure
	GetIcon1(Index int32, Image IIcon)                                                                                                                          // procedure
	GetRawImage(Index int32, OutImage *TRawImage)                                                                                                               // procedure
	Insert(AIndex int32, AImage, AMask ICustomBitmap)                                                                                                           // procedure
	InsertIcon(AIndex int32, AIcon ICustomIcon)                                                                                                                 // procedure
	InsertMasked(Index int32, AImage ICustomBitmap, MaskColor TColor)                                                                                           // procedure
	Move(ACurIndex, ANewIndex int32)                                                                                                                            // procedure
	Overlay(AIndex int32, Overlay TOverlay)                                                                                                                     // procedure
	Replace(AIndex int32, AImage, AMask ICustomBitmap, AllResolutions bool)                                                                                     // procedure
	ReplaceSlice(AIndex int32, Image ICustomBitmap, AImageRect *TRect, AllResolutions bool)                                                                     // procedure
	ReplaceSliceCentered(AIndex, AImageWidth int32, Image ICustomBitmap, AllResolutions bool)                                                                   // procedure
	ReplaceIcon(AIndex int32, AIcon ICustomIcon)                                                                                                                // procedure
	ReplaceMasked(Index int32, NewImage ICustomBitmap, MaskColor TColor, AllResolutions bool)                                                                   // procedure
	RegisterChanges(Value IChangeLink)                                                                                                                          // procedure
	StretchDraw(Canvas ICanvas, Index int32, ARect *TRect, Enabled bool)                                                                                        // procedure
	UnRegisterChanges(Value IChangeLink)                                                                                                                        // procedure
	DeleteResolution(AWidth int32)                                                                                                                              // procedure
	SetOnChange(fn TNotifyEvent)                                                                                                                                // property event
	SetOnGetWidthForPPI(fn TCustomImageListGetWidthForPPI)                                                                                                      // property event
}

// TCustomImageList Parent: TLCLComponent
type TCustomImageList struct {
	TLCLComponent
	changePtr         uintptr
	getWidthForPPIPtr uintptr
}

func NewCustomImageList(AOwner IComponent) ICustomImageList {
	r1 := LCL().SysCallN(1669, GetObjectUintptr(AOwner))
	return AsCustomImageList(r1)
}

func NewCustomImageListSize(AWidth, AHeight int32) ICustomImageList {
	r1 := LCL().SysCallN(1670, uintptr(AWidth), uintptr(AHeight))
	return AsCustomImageList(r1)
}

func (m *TCustomImageList) HasOverlays() bool {
	r1 := LCL().SysCallN(1693, m.Instance())
	return GoBool(r1)
}

func (m *TCustomImageList) AllocBy() int32 {
	r1 := LCL().SysCallN(1660, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomImageList) SetAllocBy(AValue int32) {
	LCL().SysCallN(1660, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) BlendColor() TColor {
	r1 := LCL().SysCallN(1664, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomImageList) SetBlendColor(AValue TColor) {
	LCL().SysCallN(1664, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) BkColor() TColor {
	r1 := LCL().SysCallN(1663, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomImageList) SetBkColor(AValue TColor) {
	LCL().SysCallN(1663, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) Count() int32 {
	r1 := LCL().SysCallN(1668, m.Instance())
	return int32(r1)
}

func (m *TCustomImageList) DrawingStyle() TDrawingStyle {
	r1 := LCL().SysCallN(1682, 0, m.Instance(), 0)
	return TDrawingStyle(r1)
}

func (m *TCustomImageList) SetDrawingStyle(AValue TDrawingStyle) {
	LCL().SysCallN(1682, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) Height() int32 {
	r1 := LCL().SysCallN(1694, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomImageList) SetHeight(AValue int32) {
	LCL().SysCallN(1694, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) HeightForPPI(AImageWidth, APPI int32) int32 {
	r1 := LCL().SysCallN(1695, m.Instance(), uintptr(AImageWidth), uintptr(APPI))
	return int32(r1)
}

func (m *TCustomImageList) HeightForWidth(AWidth int32) int32 {
	r1 := LCL().SysCallN(1696, m.Instance(), uintptr(AWidth))
	return int32(r1)
}

func (m *TCustomImageList) Width() int32 {
	r1 := LCL().SysCallN(1725, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomImageList) SetWidth(AValue int32) {
	LCL().SysCallN(1725, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) WidthForPPI(AImageWidth, APPI int32) int32 {
	r1 := LCL().SysCallN(1726, m.Instance(), uintptr(AImageWidth), uintptr(APPI))
	return int32(r1)
}

func (m *TCustomImageList) SizeForPPI(AImageWidth, APPI int32) (resultSize TSize) {
	LCL().SysCallN(1722, m.Instance(), uintptr(AImageWidth), uintptr(APPI), uintptr(unsafe.Pointer(&resultSize)))
	return
}

func (m *TCustomImageList) Masked() bool {
	r1 := LCL().SysCallN(1701, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImageList) SetMasked(AValue bool) {
	LCL().SysCallN(1701, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImageList) Resolution(AImageWidth int32) ICustomImageListResolution {
	r1 := LCL().SysCallN(1712, m.Instance(), uintptr(AImageWidth))
	return AsCustomImageListResolution(r1)
}

func (m *TCustomImageList) ResolutionByIndex(AIndex int32) ICustomImageListResolution {
	r1 := LCL().SysCallN(1713, m.Instance(), uintptr(AIndex))
	return AsCustomImageListResolution(r1)
}

func (m *TCustomImageList) ResolutionForPPI(AImageWidth, APPI int32, ACanvasScaleFactor float64) (resultScaledImageListResolution TScaledImageListResolution) {
	r1 := LCL().SysCallN(1715, m.Instance(), uintptr(AImageWidth), uintptr(APPI), uintptr(unsafe.Pointer(&ACanvasScaleFactor)))
	return *(*TScaledImageListResolution)(getPointer(r1))
}

func (m *TCustomImageList) ResolutionCount() int32 {
	r1 := LCL().SysCallN(1714, m.Instance())
	return int32(r1)
}

func (m *TCustomImageList) Scaled() bool {
	r1 := LCL().SysCallN(1718, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImageList) SetScaled(AValue bool) {
	LCL().SysCallN(1718, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImageList) ShareImages() bool {
	r1 := LCL().SysCallN(1721, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImageList) SetShareImages(AValue bool) {
	LCL().SysCallN(1721, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImageList) ImageType() TImageType {
	r1 := LCL().SysCallN(1697, 0, m.Instance(), 0)
	return TImageType(r1)
}

func (m *TCustomImageList) SetImageType(AValue TImageType) {
	LCL().SysCallN(1697, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) Add(Image, Mask ICustomBitmap) int32 {
	r1 := LCL().SysCallN(1651, m.Instance(), GetObjectUintptr(Image), GetObjectUintptr(Mask))
	return int32(r1)
}

func (m *TCustomImageList) AddSliced(Image ICustomBitmap, AHorizontalCount, AVerticalCount int32) int32 {
	r1 := LCL().SysCallN(1659, m.Instance(), GetObjectUintptr(Image), uintptr(AHorizontalCount), uintptr(AVerticalCount))
	return int32(r1)
}

func (m *TCustomImageList) AddSlice(Image ICustomBitmap, AImageRect *TRect) int32 {
	r1 := LCL().SysCallN(1657, m.Instance(), GetObjectUintptr(Image), uintptr(unsafe.Pointer(AImageRect)))
	return int32(r1)
}

func (m *TCustomImageList) AddSliceCentered(Image ICustomBitmap) int32 {
	r1 := LCL().SysCallN(1658, m.Instance(), GetObjectUintptr(Image))
	return int32(r1)
}

func (m *TCustomImageList) AddIcon(Image ICustomIcon) int32 {
	r1 := LCL().SysCallN(1652, m.Instance(), GetObjectUintptr(Image))
	return int32(r1)
}

func (m *TCustomImageList) AddMasked(Image IBitmap, MaskColor TColor) int32 {
	r1 := LCL().SysCallN(1655, m.Instance(), GetObjectUintptr(Image), uintptr(MaskColor))
	return int32(r1)
}

func (m *TCustomImageList) AddLazarusResource(ResourceName string, MaskColor TColor) int32 {
	r1 := LCL().SysCallN(1654, m.Instance(), PascalStr(ResourceName), uintptr(MaskColor))
	return int32(r1)
}

func (m *TCustomImageList) AddResourceName(Instance THandle, ResourceName string, MaskColor TColor) int32 {
	r1 := LCL().SysCallN(1656, m.Instance(), uintptr(Instance), PascalStr(ResourceName), uintptr(MaskColor))
	return int32(r1)
}

func (m *TCustomImageList) GetHotSpot() (resultPoint TPoint) {
	LCL().SysCallN(1689, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TCustomImageList) FindResolution(AImageWidth int32, OutResolution *ICustomImageListResolution) bool {
	var result1 uintptr
	r1 := LCL().SysCallN(1684, m.Instance(), uintptr(AImageWidth), uintptr(unsafe.Pointer(&result1)))
	*OutResolution = AsCustomImageListResolution(result1)
	return GoBool(r1)
}

func (m *TCustomImageList) Resolutions() ICustomImageListResolutionEnumerator {
	r1 := LCL().SysCallN(1716, m.Instance())
	return AsCustomImageListResolutionEnumerator(r1)
}

func (m *TCustomImageList) ResolutionsDesc() ICustomImageListResolutionEnumerator {
	r1 := LCL().SysCallN(1717, m.Instance())
	return AsCustomImageListResolutionEnumerator(r1)
}

func CustomImageListClass() TClass {
	ret := LCL().SysCallN(1666)
	return TClass(ret)
}

func (m *TCustomImageList) AssignTo(Dest IPersistent) {
	LCL().SysCallN(1661, m.Instance(), GetObjectUintptr(Dest))
}

func (m *TCustomImageList) WriteData(AStream IStream) {
	LCL().SysCallN(1728, m.Instance(), GetObjectUintptr(AStream))
}

func (m *TCustomImageList) ReadData(AStream IStream) {
	LCL().SysCallN(1705, m.Instance(), GetObjectUintptr(AStream))
}

func (m *TCustomImageList) WriteAdvData(AStream IStream) {
	LCL().SysCallN(1727, m.Instance(), GetObjectUintptr(AStream))
}

func (m *TCustomImageList) ReadAdvData(AStream IStream) {
	LCL().SysCallN(1704, m.Instance(), GetObjectUintptr(AStream))
}

func (m *TCustomImageList) BeginUpdate() {
	LCL().SysCallN(1662, m.Instance())
}

func (m *TCustomImageList) EndUpdate() {
	LCL().SysCallN(1683, m.Instance())
}

func (m *TCustomImageList) AddImages(AValue ICustomImageList) {
	LCL().SysCallN(1653, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomImageList) Change() {
	LCL().SysCallN(1665, m.Instance())
}

func (m *TCustomImageList) Clear() {
	LCL().SysCallN(1667, m.Instance())
}

func (m *TCustomImageList) Delete(AIndex int32) {
	LCL().SysCallN(1671, m.Instance(), uintptr(AIndex))
}

func (m *TCustomImageList) Draw(ACanvas ICanvas, AX, AY, AIndex int32, AEnabled bool) {
	LCL().SysCallN(1673, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), PascalBool(AEnabled))
}

func (m *TCustomImageList) Draw1(ACanvas ICanvas, AX, AY, AIndex int32, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1674, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawEffect))
}

func (m *TCustomImageList) Draw2(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, AEnabled bool) {
	LCL().SysCallN(1675, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawingStyle), uintptr(AImageType), PascalBool(AEnabled))
}

func (m *TCustomImageList) Draw3(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1676, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawingStyle), uintptr(AImageType), uintptr(ADrawEffect))
}

func (m *TCustomImageList) DrawForPPI(ACanvas ICanvas, AX, AY, AIndex int32, AImageWidthAt96PPI, ATargetPPI int32, ACanvasFactor float64, AEnabled bool) {
	LCL().SysCallN(1677, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AImageWidthAt96PPI), uintptr(ATargetPPI), uintptr(unsafe.Pointer(&ACanvasFactor)), PascalBool(AEnabled))
}

func (m *TCustomImageList) DrawForPPI1(ACanvas ICanvas, AX, AY, AIndex int32, AImageWidthAt96PPI, ATargetPPI int32, ACanvasFactor float64, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1678, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AImageWidthAt96PPI), uintptr(ATargetPPI), uintptr(unsafe.Pointer(&ACanvasFactor)), uintptr(ADrawEffect))
}

func (m *TCustomImageList) DrawOverlay(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, AEnabled bool) {
	LCL().SysCallN(1679, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), PascalBool(AEnabled))
}

func (m *TCustomImageList) DrawOverlay1(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1680, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), uintptr(ADrawEffect))
}

func (m *TCustomImageList) DrawOverlay2(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1681, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), uintptr(ADrawingStyle), uintptr(AImageType), uintptr(ADrawEffect))
}

func (m *TCustomImageList) GetBitmap(Index int32, Image ICustomBitmap) {
	LCL().SysCallN(1685, m.Instance(), uintptr(Index), GetObjectUintptr(Image))
}

func (m *TCustomImageList) GetBitmap1(Index int32, Image ICustomBitmap, AEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1686, m.Instance(), uintptr(Index), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageList) GetFullBitmap(Image ICustomBitmap, AEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1687, m.Instance(), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageList) GetFullRawImage(OutImage *TRawImage) {
	var result0 uintptr
	LCL().SysCallN(1688, m.Instance(), uintptr(unsafe.Pointer(&result0)))
	*OutImage = *(*TRawImage)(getPointer(result0))
}

func (m *TCustomImageList) GetIcon(Index int32, Image IIcon, AEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1690, m.Instance(), uintptr(Index), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageList) GetIcon1(Index int32, Image IIcon) {
	LCL().SysCallN(1691, m.Instance(), uintptr(Index), GetObjectUintptr(Image))
}

func (m *TCustomImageList) GetRawImage(Index int32, OutImage *TRawImage) {
	var result1 uintptr
	LCL().SysCallN(1692, m.Instance(), uintptr(Index), uintptr(unsafe.Pointer(&result1)))
	*OutImage = *(*TRawImage)(getPointer(result1))
}

func (m *TCustomImageList) Insert(AIndex int32, AImage, AMask ICustomBitmap) {
	LCL().SysCallN(1698, m.Instance(), uintptr(AIndex), GetObjectUintptr(AImage), GetObjectUintptr(AMask))
}

func (m *TCustomImageList) InsertIcon(AIndex int32, AIcon ICustomIcon) {
	LCL().SysCallN(1699, m.Instance(), uintptr(AIndex), GetObjectUintptr(AIcon))
}

func (m *TCustomImageList) InsertMasked(Index int32, AImage ICustomBitmap, MaskColor TColor) {
	LCL().SysCallN(1700, m.Instance(), uintptr(Index), GetObjectUintptr(AImage), uintptr(MaskColor))
}

func (m *TCustomImageList) Move(ACurIndex, ANewIndex int32) {
	LCL().SysCallN(1702, m.Instance(), uintptr(ACurIndex), uintptr(ANewIndex))
}

func (m *TCustomImageList) Overlay(AIndex int32, Overlay TOverlay) {
	LCL().SysCallN(1703, m.Instance(), uintptr(AIndex), uintptr(Overlay))
}

func (m *TCustomImageList) Replace(AIndex int32, AImage, AMask ICustomBitmap, AllResolutions bool) {
	LCL().SysCallN(1707, m.Instance(), uintptr(AIndex), GetObjectUintptr(AImage), GetObjectUintptr(AMask), PascalBool(AllResolutions))
}

func (m *TCustomImageList) ReplaceSlice(AIndex int32, Image ICustomBitmap, AImageRect *TRect, AllResolutions bool) {
	LCL().SysCallN(1710, m.Instance(), uintptr(AIndex), GetObjectUintptr(Image), uintptr(unsafe.Pointer(AImageRect)), PascalBool(AllResolutions))
}

func (m *TCustomImageList) ReplaceSliceCentered(AIndex, AImageWidth int32, Image ICustomBitmap, AllResolutions bool) {
	LCL().SysCallN(1711, m.Instance(), uintptr(AIndex), uintptr(AImageWidth), GetObjectUintptr(Image), PascalBool(AllResolutions))
}

func (m *TCustomImageList) ReplaceIcon(AIndex int32, AIcon ICustomIcon) {
	LCL().SysCallN(1708, m.Instance(), uintptr(AIndex), GetObjectUintptr(AIcon))
}

func (m *TCustomImageList) ReplaceMasked(Index int32, NewImage ICustomBitmap, MaskColor TColor, AllResolutions bool) {
	LCL().SysCallN(1709, m.Instance(), uintptr(Index), GetObjectUintptr(NewImage), uintptr(MaskColor), PascalBool(AllResolutions))
}

func (m *TCustomImageList) RegisterChanges(Value IChangeLink) {
	LCL().SysCallN(1706, m.Instance(), GetObjectUintptr(Value))
}

func (m *TCustomImageList) StretchDraw(Canvas ICanvas, Index int32, ARect *TRect, Enabled bool) {
	LCL().SysCallN(1723, m.Instance(), GetObjectUintptr(Canvas), uintptr(Index), uintptr(unsafe.Pointer(ARect)), PascalBool(Enabled))
}

func (m *TCustomImageList) UnRegisterChanges(Value IChangeLink) {
	LCL().SysCallN(1724, m.Instance(), GetObjectUintptr(Value))
}

func (m *TCustomImageList) DeleteResolution(AWidth int32) {
	LCL().SysCallN(1672, m.Instance(), uintptr(AWidth))
}

func (m *TCustomImageList) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1719, m.Instance(), m.changePtr)
}

func (m *TCustomImageList) SetOnGetWidthForPPI(fn TCustomImageListGetWidthForPPI) {
	if m.getWidthForPPIPtr != 0 {
		RemoveEventElement(m.getWidthForPPIPtr)
	}
	m.getWidthForPPIPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1720, m.Instance(), m.getWidthForPPIPtr)
}
