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
	r1 := LCL().SysCallN(1859, GetObjectUintptr(AOwner))
	return AsCustomImageList(r1)
}

func NewCustomImageListSize(AWidth, AHeight int32) ICustomImageList {
	r1 := LCL().SysCallN(1860, uintptr(AWidth), uintptr(AHeight))
	return AsCustomImageList(r1)
}

func (m *TCustomImageList) HasOverlays() bool {
	r1 := LCL().SysCallN(1883, m.Instance())
	return GoBool(r1)
}

func (m *TCustomImageList) AllocBy() int32 {
	r1 := LCL().SysCallN(1850, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomImageList) SetAllocBy(AValue int32) {
	LCL().SysCallN(1850, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) BlendColor() TColor {
	r1 := LCL().SysCallN(1854, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomImageList) SetBlendColor(AValue TColor) {
	LCL().SysCallN(1854, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) BkColor() TColor {
	r1 := LCL().SysCallN(1853, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomImageList) SetBkColor(AValue TColor) {
	LCL().SysCallN(1853, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) Count() int32 {
	r1 := LCL().SysCallN(1858, m.Instance())
	return int32(r1)
}

func (m *TCustomImageList) DrawingStyle() TDrawingStyle {
	r1 := LCL().SysCallN(1872, 0, m.Instance(), 0)
	return TDrawingStyle(r1)
}

func (m *TCustomImageList) SetDrawingStyle(AValue TDrawingStyle) {
	LCL().SysCallN(1872, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) Height() int32 {
	r1 := LCL().SysCallN(1884, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomImageList) SetHeight(AValue int32) {
	LCL().SysCallN(1884, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) HeightForPPI(AImageWidth, APPI int32) int32 {
	r1 := LCL().SysCallN(1885, m.Instance(), uintptr(AImageWidth), uintptr(APPI))
	return int32(r1)
}

func (m *TCustomImageList) HeightForWidth(AWidth int32) int32 {
	r1 := LCL().SysCallN(1886, m.Instance(), uintptr(AWidth))
	return int32(r1)
}

func (m *TCustomImageList) Width() int32 {
	r1 := LCL().SysCallN(1915, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomImageList) SetWidth(AValue int32) {
	LCL().SysCallN(1915, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) WidthForPPI(AImageWidth, APPI int32) int32 {
	r1 := LCL().SysCallN(1916, m.Instance(), uintptr(AImageWidth), uintptr(APPI))
	return int32(r1)
}

func (m *TCustomImageList) SizeForPPI(AImageWidth, APPI int32) (resultSize TSize) {
	LCL().SysCallN(1912, m.Instance(), uintptr(AImageWidth), uintptr(APPI), uintptr(unsafePointer(&resultSize)))
	return
}

func (m *TCustomImageList) Masked() bool {
	r1 := LCL().SysCallN(1891, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImageList) SetMasked(AValue bool) {
	LCL().SysCallN(1891, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImageList) Resolution(AImageWidth int32) ICustomImageListResolution {
	r1 := LCL().SysCallN(1902, m.Instance(), uintptr(AImageWidth))
	return AsCustomImageListResolution(r1)
}

func (m *TCustomImageList) ResolutionByIndex(AIndex int32) ICustomImageListResolution {
	r1 := LCL().SysCallN(1903, m.Instance(), uintptr(AIndex))
	return AsCustomImageListResolution(r1)
}

func (m *TCustomImageList) ResolutionForPPI(AImageWidth, APPI int32, ACanvasScaleFactor float64) (resultScaledImageListResolution TScaledImageListResolution) {
	r1 := LCL().SysCallN(1905, m.Instance(), uintptr(AImageWidth), uintptr(APPI), uintptr(unsafePointer(&ACanvasScaleFactor)))
	return *(*TScaledImageListResolution)(getPointer(r1))
}

func (m *TCustomImageList) ResolutionCount() int32 {
	r1 := LCL().SysCallN(1904, m.Instance())
	return int32(r1)
}

func (m *TCustomImageList) Scaled() bool {
	r1 := LCL().SysCallN(1908, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImageList) SetScaled(AValue bool) {
	LCL().SysCallN(1908, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImageList) ShareImages() bool {
	r1 := LCL().SysCallN(1911, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImageList) SetShareImages(AValue bool) {
	LCL().SysCallN(1911, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImageList) ImageType() TImageType {
	r1 := LCL().SysCallN(1887, 0, m.Instance(), 0)
	return TImageType(r1)
}

func (m *TCustomImageList) SetImageType(AValue TImageType) {
	LCL().SysCallN(1887, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) Add(Image, Mask ICustomBitmap) int32 {
	r1 := LCL().SysCallN(1841, m.Instance(), GetObjectUintptr(Image), GetObjectUintptr(Mask))
	return int32(r1)
}

func (m *TCustomImageList) AddSliced(Image ICustomBitmap, AHorizontalCount, AVerticalCount int32) int32 {
	r1 := LCL().SysCallN(1849, m.Instance(), GetObjectUintptr(Image), uintptr(AHorizontalCount), uintptr(AVerticalCount))
	return int32(r1)
}

func (m *TCustomImageList) AddSlice(Image ICustomBitmap, AImageRect *TRect) int32 {
	r1 := LCL().SysCallN(1847, m.Instance(), GetObjectUintptr(Image), uintptr(unsafePointer(AImageRect)))
	return int32(r1)
}

func (m *TCustomImageList) AddSliceCentered(Image ICustomBitmap) int32 {
	r1 := LCL().SysCallN(1848, m.Instance(), GetObjectUintptr(Image))
	return int32(r1)
}

func (m *TCustomImageList) AddIcon(Image ICustomIcon) int32 {
	r1 := LCL().SysCallN(1842, m.Instance(), GetObjectUintptr(Image))
	return int32(r1)
}

func (m *TCustomImageList) AddMasked(Image IBitmap, MaskColor TColor) int32 {
	r1 := LCL().SysCallN(1845, m.Instance(), GetObjectUintptr(Image), uintptr(MaskColor))
	return int32(r1)
}

func (m *TCustomImageList) AddLazarusResource(ResourceName string, MaskColor TColor) int32 {
	r1 := LCL().SysCallN(1844, m.Instance(), PascalStr(ResourceName), uintptr(MaskColor))
	return int32(r1)
}

func (m *TCustomImageList) AddResourceName(Instance THandle, ResourceName string, MaskColor TColor) int32 {
	r1 := LCL().SysCallN(1846, m.Instance(), uintptr(Instance), PascalStr(ResourceName), uintptr(MaskColor))
	return int32(r1)
}

func (m *TCustomImageList) GetHotSpot() (resultPoint TPoint) {
	LCL().SysCallN(1879, m.Instance(), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCustomImageList) FindResolution(AImageWidth int32, OutResolution *ICustomImageListResolution) bool {
	var result1 uintptr
	r1 := LCL().SysCallN(1874, m.Instance(), uintptr(AImageWidth), uintptr(unsafePointer(&result1)))
	*OutResolution = AsCustomImageListResolution(result1)
	return GoBool(r1)
}

func (m *TCustomImageList) Resolutions() ICustomImageListResolutionEnumerator {
	r1 := LCL().SysCallN(1906, m.Instance())
	return AsCustomImageListResolutionEnumerator(r1)
}

func (m *TCustomImageList) ResolutionsDesc() ICustomImageListResolutionEnumerator {
	r1 := LCL().SysCallN(1907, m.Instance())
	return AsCustomImageListResolutionEnumerator(r1)
}

func CustomImageListClass() TClass {
	ret := LCL().SysCallN(1856)
	return TClass(ret)
}

func (m *TCustomImageList) AssignTo(Dest IPersistent) {
	LCL().SysCallN(1851, m.Instance(), GetObjectUintptr(Dest))
}

func (m *TCustomImageList) WriteData(AStream IStream) {
	LCL().SysCallN(1918, m.Instance(), GetObjectUintptr(AStream))
}

func (m *TCustomImageList) ReadData(AStream IStream) {
	LCL().SysCallN(1895, m.Instance(), GetObjectUintptr(AStream))
}

func (m *TCustomImageList) WriteAdvData(AStream IStream) {
	LCL().SysCallN(1917, m.Instance(), GetObjectUintptr(AStream))
}

func (m *TCustomImageList) ReadAdvData(AStream IStream) {
	LCL().SysCallN(1894, m.Instance(), GetObjectUintptr(AStream))
}

func (m *TCustomImageList) BeginUpdate() {
	LCL().SysCallN(1852, m.Instance())
}

func (m *TCustomImageList) EndUpdate() {
	LCL().SysCallN(1873, m.Instance())
}

func (m *TCustomImageList) AddImages(AValue ICustomImageList) {
	LCL().SysCallN(1843, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomImageList) Change() {
	LCL().SysCallN(1855, m.Instance())
}

func (m *TCustomImageList) Clear() {
	LCL().SysCallN(1857, m.Instance())
}

func (m *TCustomImageList) Delete(AIndex int32) {
	LCL().SysCallN(1861, m.Instance(), uintptr(AIndex))
}

func (m *TCustomImageList) Draw(ACanvas ICanvas, AX, AY, AIndex int32, AEnabled bool) {
	LCL().SysCallN(1863, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), PascalBool(AEnabled))
}

func (m *TCustomImageList) Draw1(ACanvas ICanvas, AX, AY, AIndex int32, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1864, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawEffect))
}

func (m *TCustomImageList) Draw2(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, AEnabled bool) {
	LCL().SysCallN(1865, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawingStyle), uintptr(AImageType), PascalBool(AEnabled))
}

func (m *TCustomImageList) Draw3(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1866, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawingStyle), uintptr(AImageType), uintptr(ADrawEffect))
}

func (m *TCustomImageList) DrawForPPI(ACanvas ICanvas, AX, AY, AIndex int32, AImageWidthAt96PPI, ATargetPPI int32, ACanvasFactor float64, AEnabled bool) {
	LCL().SysCallN(1867, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AImageWidthAt96PPI), uintptr(ATargetPPI), uintptr(unsafePointer(&ACanvasFactor)), PascalBool(AEnabled))
}

func (m *TCustomImageList) DrawForPPI1(ACanvas ICanvas, AX, AY, AIndex int32, AImageWidthAt96PPI, ATargetPPI int32, ACanvasFactor float64, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1868, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AImageWidthAt96PPI), uintptr(ATargetPPI), uintptr(unsafePointer(&ACanvasFactor)), uintptr(ADrawEffect))
}

func (m *TCustomImageList) DrawOverlay(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, AEnabled bool) {
	LCL().SysCallN(1869, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), PascalBool(AEnabled))
}

func (m *TCustomImageList) DrawOverlay1(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1870, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), uintptr(ADrawEffect))
}

func (m *TCustomImageList) DrawOverlay2(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1871, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), uintptr(ADrawingStyle), uintptr(AImageType), uintptr(ADrawEffect))
}

func (m *TCustomImageList) GetBitmap(Index int32, Image ICustomBitmap) {
	LCL().SysCallN(1875, m.Instance(), uintptr(Index), GetObjectUintptr(Image))
}

func (m *TCustomImageList) GetBitmap1(Index int32, Image ICustomBitmap, AEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1876, m.Instance(), uintptr(Index), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageList) GetFullBitmap(Image ICustomBitmap, AEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1877, m.Instance(), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageList) GetFullRawImage(OutImage *TRawImage) {
	var result0 uintptr
	LCL().SysCallN(1878, m.Instance(), uintptr(unsafePointer(&result0)))
	*OutImage = *(*TRawImage)(getPointer(result0))
}

func (m *TCustomImageList) GetIcon(Index int32, Image IIcon, AEffect TGraphicsDrawEffect) {
	LCL().SysCallN(1880, m.Instance(), uintptr(Index), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageList) GetIcon1(Index int32, Image IIcon) {
	LCL().SysCallN(1881, m.Instance(), uintptr(Index), GetObjectUintptr(Image))
}

func (m *TCustomImageList) GetRawImage(Index int32, OutImage *TRawImage) {
	var result1 uintptr
	LCL().SysCallN(1882, m.Instance(), uintptr(Index), uintptr(unsafePointer(&result1)))
	*OutImage = *(*TRawImage)(getPointer(result1))
}

func (m *TCustomImageList) Insert(AIndex int32, AImage, AMask ICustomBitmap) {
	LCL().SysCallN(1888, m.Instance(), uintptr(AIndex), GetObjectUintptr(AImage), GetObjectUintptr(AMask))
}

func (m *TCustomImageList) InsertIcon(AIndex int32, AIcon ICustomIcon) {
	LCL().SysCallN(1889, m.Instance(), uintptr(AIndex), GetObjectUintptr(AIcon))
}

func (m *TCustomImageList) InsertMasked(Index int32, AImage ICustomBitmap, MaskColor TColor) {
	LCL().SysCallN(1890, m.Instance(), uintptr(Index), GetObjectUintptr(AImage), uintptr(MaskColor))
}

func (m *TCustomImageList) Move(ACurIndex, ANewIndex int32) {
	LCL().SysCallN(1892, m.Instance(), uintptr(ACurIndex), uintptr(ANewIndex))
}

func (m *TCustomImageList) Overlay(AIndex int32, Overlay TOverlay) {
	LCL().SysCallN(1893, m.Instance(), uintptr(AIndex), uintptr(Overlay))
}

func (m *TCustomImageList) Replace(AIndex int32, AImage, AMask ICustomBitmap, AllResolutions bool) {
	LCL().SysCallN(1897, m.Instance(), uintptr(AIndex), GetObjectUintptr(AImage), GetObjectUintptr(AMask), PascalBool(AllResolutions))
}

func (m *TCustomImageList) ReplaceSlice(AIndex int32, Image ICustomBitmap, AImageRect *TRect, AllResolutions bool) {
	LCL().SysCallN(1900, m.Instance(), uintptr(AIndex), GetObjectUintptr(Image), uintptr(unsafePointer(AImageRect)), PascalBool(AllResolutions))
}

func (m *TCustomImageList) ReplaceSliceCentered(AIndex, AImageWidth int32, Image ICustomBitmap, AllResolutions bool) {
	LCL().SysCallN(1901, m.Instance(), uintptr(AIndex), uintptr(AImageWidth), GetObjectUintptr(Image), PascalBool(AllResolutions))
}

func (m *TCustomImageList) ReplaceIcon(AIndex int32, AIcon ICustomIcon) {
	LCL().SysCallN(1898, m.Instance(), uintptr(AIndex), GetObjectUintptr(AIcon))
}

func (m *TCustomImageList) ReplaceMasked(Index int32, NewImage ICustomBitmap, MaskColor TColor, AllResolutions bool) {
	LCL().SysCallN(1899, m.Instance(), uintptr(Index), GetObjectUintptr(NewImage), uintptr(MaskColor), PascalBool(AllResolutions))
}

func (m *TCustomImageList) RegisterChanges(Value IChangeLink) {
	LCL().SysCallN(1896, m.Instance(), GetObjectUintptr(Value))
}

func (m *TCustomImageList) StretchDraw(Canvas ICanvas, Index int32, ARect *TRect, Enabled bool) {
	LCL().SysCallN(1913, m.Instance(), GetObjectUintptr(Canvas), uintptr(Index), uintptr(unsafePointer(ARect)), PascalBool(Enabled))
}

func (m *TCustomImageList) UnRegisterChanges(Value IChangeLink) {
	LCL().SysCallN(1914, m.Instance(), GetObjectUintptr(Value))
}

func (m *TCustomImageList) DeleteResolution(AWidth int32) {
	LCL().SysCallN(1862, m.Instance(), uintptr(AWidth))
}

func (m *TCustomImageList) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1909, m.Instance(), m.changePtr)
}

func (m *TCustomImageList) SetOnGetWidthForPPI(fn TCustomImageListGetWidthForPPI) {
	if m.getWidthForPPIPtr != 0 {
		RemoveEventElement(m.getWidthForPPIPtr)
	}
	m.getWidthForPPIPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1910, m.Instance(), m.getWidthForPPIPtr)
}
