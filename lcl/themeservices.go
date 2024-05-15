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

// IThemeServices Parent: IObject
// TThemeServices is a small foot print class to provide the user with pure
// Windows XP theme related abilities like painting elements and text or
// retrieving certain info.
type IThemeServices interface {
	IObject
	DottedBrush() HBRUSH                                                                                                                                       // property
	ThemesAvailable() bool                                                                                                                                     // property
	ThemesEnabled() bool                                                                                                                                       // property
	IsDisabled(Details *TThemedElementDetails) bool                                                                                                            // function
	IsPushed(Details *TThemedElementDetails) bool                                                                                                              // function
	IsHot(Details *TThemedElementDetails) bool                                                                                                                 // function
	IsChecked(Details *TThemedElementDetails) bool                                                                                                             // function
	IsMixed(Details *TThemedElementDetails) bool                                                                                                               // function
	GetElementDetails(Detail TThemedButton) (resultThemedElementDetails TThemedElementDetails)                                                                 // function
	GetElementDetails1(Detail TThemedClock) (resultThemedElementDetails TThemedElementDetails)                                                                 // function
	GetElementDetails2(Detail TThemedComboBox) (resultThemedElementDetails TThemedElementDetails)                                                              // function
	GetElementDetails3(Detail TThemedEdit) (resultThemedElementDetails TThemedElementDetails)                                                                  // function
	GetElementDetails4(Detail TThemedExplorerBar) (resultThemedElementDetails TThemedElementDetails)                                                           // function
	GetElementDetails5(Detail TThemedHeader) (resultThemedElementDetails TThemedElementDetails)                                                                // function
	GetElementDetails6(Detail TThemedListView) (resultThemedElementDetails TThemedElementDetails)                                                              // function
	GetElementDetails7(Detail TThemedMenu) (resultThemedElementDetails TThemedElementDetails)                                                                  // function
	GetElementDetails8(Detail TThemedPage) (resultThemedElementDetails TThemedElementDetails)                                                                  // function
	GetElementDetails9(Detail TThemedProgress) (resultThemedElementDetails TThemedElementDetails)                                                              // function
	GetElementDetails10(Detail TThemedRebar) (resultThemedElementDetails TThemedElementDetails)                                                                // function
	GetElementDetails11(Detail TThemedScrollBar) (resultThemedElementDetails TThemedElementDetails)                                                            // function
	GetElementDetails12(Detail TThemedSpin) (resultThemedElementDetails TThemedElementDetails)                                                                 // function
	GetElementDetails13(Detail TThemedStartPanel) (resultThemedElementDetails TThemedElementDetails)                                                           // function
	GetElementDetails14(Detail TThemedStatus) (resultThemedElementDetails TThemedElementDetails)                                                               // function
	GetElementDetails15(Detail TThemedTab) (resultThemedElementDetails TThemedElementDetails)                                                                  // function
	GetElementDetails16(Detail TThemedTaskBand) (resultThemedElementDetails TThemedElementDetails)                                                             // function
	GetElementDetails17(Detail TThemedTaskBar) (resultThemedElementDetails TThemedElementDetails)                                                              // function
	GetElementDetails18(Detail TThemedToolBar) (resultThemedElementDetails TThemedElementDetails)                                                              // function
	GetElementDetails19(Detail TThemedToolTip) (resultThemedElementDetails TThemedElementDetails)                                                              // function
	GetElementDetails20(Detail TThemedTrackBar) (resultThemedElementDetails TThemedElementDetails)                                                             // function
	GetElementDetails21(Detail TThemedTrayNotify) (resultThemedElementDetails TThemedElementDetails)                                                           // function
	GetElementDetails22(Detail TThemedTreeview) (resultThemedElementDetails TThemedElementDetails)                                                             // function
	GetElementDetails23(Detail TThemedWindow) (resultThemedElementDetails TThemedElementDetails)                                                               // function
	GetDetailSizeForWindow(Details *TThemedElementDetails, AWindow HWND) (resultSize TSize)                                                                    // function
	GetDetailSizeForPPI(Details *TThemedElementDetails, PPI int32) (resultSize TSize)                                                                          // function
	GetDetailRegion(DC HDC, Details *TThemedElementDetails, R *TRect) HRGN                                                                                     // function
	GetStockImage(StockID int32, OutImage, OutMask *HBITMAP) bool                                                                                              // function
	GetStockImage1(StockID int32, AWidth, AHeight int32, OutImage, OutMask *HBITMAP) bool                                                                      // function
	GetOption(AOption TThemeOption) int32                                                                                                                      // function
	GetTextExtent(DC HDC, Details *TThemedElementDetails, S string, Flags uint32, BoundingRect *TRect) (resultRect TRect)                                      // function
	ColorToRGB(Color int32, Details *PThemedElementDetails) COLORREF                                                                                           // function
	ContentRect(DC HDC, Details *TThemedElementDetails, BoundingRect *TRect) (resultRect TRect)                                                                // function
	HasTransparentParts(Details *TThemedElementDetails) bool                                                                                                   // function
	IntfDoOnThemeChange()                                                                                                                                      // procedure
	DrawEdge(DC HDC, Details *TThemedElementDetails, R *TRect, Edge, Flags uint32, AContentRect *TRect)                                                        // procedure
	DrawElement(DC HDC, Details *TThemedElementDetails, R *TRect, ClipRect *TRect)                                                                             // procedure
	DrawIcon(DC HDC, Details *TThemedElementDetails, R *TRect, himl HIMAGELIST, Index int32)                                                                   // procedure
	DrawIcon1(ACanvas IPersistent, Details *TThemedElementDetails, P *TPoint, AImageList IPersistent, Index int32, AImageWidth int32, ARefControl IPersistent) // procedure
	DrawParentBackground(Window HWND, Target HDC, Details *PThemedElementDetails, OnlyIfTransparent bool, Bounds *TRect)                                       // procedure
	DrawText(DC HDC, Details *TThemedElementDetails, S string, R *TRect, Flags, Flags2 uint32)                                                                 // procedure
	DrawText1(ACanvas IPersistent, Details *TThemedElementDetails, S string, R *TRect, Flags, Flags2 uint32)                                                   // procedure
	PaintBorder(Control IObject, EraseLRCorner bool)                                                                                                           // procedure
	UpdateThemes()                                                                                                                                             // procedure
	SetOnThemeChange(fn TNotifyEvent)                                                                                                                          // property event
}

// TThemeServices Parent: TObject
// TThemeServices is a small foot print class to provide the user with pure
// Windows XP theme related abilities like painting elements and text or
// retrieving certain info.
type TThemeServices struct {
	TObject
	themeChangePtr uintptr
}

func NewThemeServices() IThemeServices {
	r1 := LCL().SysCallN(5379)
	return AsThemeServices(r1)
}

func (m *TThemeServices) DottedBrush() HBRUSH {
	r1 := LCL().SysCallN(5380, m.Instance())
	return HBRUSH(r1)
}

func (m *TThemeServices) ThemesAvailable() bool {
	r1 := LCL().SysCallN(5428, m.Instance())
	return GoBool(r1)
}

func (m *TThemeServices) ThemesEnabled() bool {
	r1 := LCL().SysCallN(5429, m.Instance())
	return GoBool(r1)
}

func (m *TThemeServices) IsDisabled(Details *TThemedElementDetails) bool {
	r1 := LCL().SysCallN(5422, m.Instance(), uintptr(unsafePointer(Details)))
	return GoBool(r1)
}

func (m *TThemeServices) IsPushed(Details *TThemedElementDetails) bool {
	r1 := LCL().SysCallN(5425, m.Instance(), uintptr(unsafePointer(Details)))
	return GoBool(r1)
}

func (m *TThemeServices) IsHot(Details *TThemedElementDetails) bool {
	r1 := LCL().SysCallN(5423, m.Instance(), uintptr(unsafePointer(Details)))
	return GoBool(r1)
}

func (m *TThemeServices) IsChecked(Details *TThemedElementDetails) bool {
	r1 := LCL().SysCallN(5421, m.Instance(), uintptr(unsafePointer(Details)))
	return GoBool(r1)
}

func (m *TThemeServices) IsMixed(Details *TThemedElementDetails) bool {
	r1 := LCL().SysCallN(5424, m.Instance(), uintptr(unsafePointer(Details)))
	return GoBool(r1)
}

func (m *TThemeServices) GetElementDetails(Detail TThemedButton) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5391, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails1(Detail TThemedClock) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5392, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails2(Detail TThemedComboBox) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5403, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails3(Detail TThemedEdit) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5408, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails4(Detail TThemedExplorerBar) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5409, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails5(Detail TThemedHeader) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5410, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails6(Detail TThemedListView) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5411, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails7(Detail TThemedMenu) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5412, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails8(Detail TThemedPage) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5413, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails9(Detail TThemedProgress) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5414, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails10(Detail TThemedRebar) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5393, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails11(Detail TThemedScrollBar) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5394, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails12(Detail TThemedSpin) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5395, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails13(Detail TThemedStartPanel) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5396, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails14(Detail TThemedStatus) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5397, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails15(Detail TThemedTab) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5398, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails16(Detail TThemedTaskBand) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5399, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails17(Detail TThemedTaskBar) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5400, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails18(Detail TThemedToolBar) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5401, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails19(Detail TThemedToolTip) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5402, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails20(Detail TThemedTrackBar) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5404, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails21(Detail TThemedTrayNotify) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5405, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails22(Detail TThemedTreeview) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5406, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails23(Detail TThemedWindow) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(5407, m.Instance(), uintptr(Detail), uintptr(unsafePointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetDetailSizeForWindow(Details *TThemedElementDetails, AWindow HWND) (resultSize TSize) {
	LCL().SysCallN(5390, m.Instance(), uintptr(unsafePointer(Details)), uintptr(AWindow), uintptr(unsafePointer(&resultSize)))
	return
}

func (m *TThemeServices) GetDetailSizeForPPI(Details *TThemedElementDetails, PPI int32) (resultSize TSize) {
	LCL().SysCallN(5389, m.Instance(), uintptr(unsafePointer(Details)), uintptr(PPI), uintptr(unsafePointer(&resultSize)))
	return
}

func (m *TThemeServices) GetDetailRegion(DC HDC, Details *TThemedElementDetails, R *TRect) HRGN {
	r1 := LCL().SysCallN(5388, m.Instance(), uintptr(DC), uintptr(unsafePointer(Details)), uintptr(unsafePointer(R)))
	return HRGN(r1)
}

func (m *TThemeServices) GetStockImage(StockID int32, OutImage, OutMask *HBITMAP) bool {
	var result1 uintptr
	var result2 uintptr
	r1 := LCL().SysCallN(5416, m.Instance(), uintptr(StockID), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*OutImage = HBITMAP(result1)
	*OutMask = HBITMAP(result2)
	return GoBool(r1)
}

func (m *TThemeServices) GetStockImage1(StockID int32, AWidth, AHeight int32, OutImage, OutMask *HBITMAP) bool {
	var result2 uintptr
	var result3 uintptr
	r1 := LCL().SysCallN(5417, m.Instance(), uintptr(StockID), uintptr(AWidth), uintptr(AHeight), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)))
	*OutImage = HBITMAP(result2)
	*OutMask = HBITMAP(result3)
	return GoBool(r1)
}

func (m *TThemeServices) GetOption(AOption TThemeOption) int32 {
	r1 := LCL().SysCallN(5415, m.Instance(), uintptr(AOption))
	return int32(r1)
}

func (m *TThemeServices) GetTextExtent(DC HDC, Details *TThemedElementDetails, S string, Flags uint32, BoundingRect *TRect) (resultRect TRect) {
	LCL().SysCallN(5418, m.Instance(), uintptr(DC), uintptr(unsafePointer(Details)), PascalStr(S), uintptr(Flags), uintptr(unsafePointer(BoundingRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TThemeServices) ColorToRGB(Color int32, Details *PThemedElementDetails) COLORREF {
	r1 := LCL().SysCallN(5377, m.Instance(), uintptr(Color), uintptr(unsafePointer(Details)))
	return COLORREF(r1)
}

func (m *TThemeServices) ContentRect(DC HDC, Details *TThemedElementDetails, BoundingRect *TRect) (resultRect TRect) {
	LCL().SysCallN(5378, m.Instance(), uintptr(DC), uintptr(unsafePointer(Details)), uintptr(unsafePointer(BoundingRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TThemeServices) HasTransparentParts(Details *TThemedElementDetails) bool {
	r1 := LCL().SysCallN(5419, m.Instance(), uintptr(unsafePointer(Details)))
	return GoBool(r1)
}

func ThemeServicesClass() TClass {
	ret := LCL().SysCallN(5376)
	return TClass(ret)
}

func (m *TThemeServices) IntfDoOnThemeChange() {
	LCL().SysCallN(5420, m.Instance())
}

func (m *TThemeServices) DrawEdge(DC HDC, Details *TThemedElementDetails, R *TRect, Edge, Flags uint32, AContentRect *TRect) {
	LCL().SysCallN(5381, m.Instance(), uintptr(DC), uintptr(unsafePointer(Details)), uintptr(unsafePointer(R)), uintptr(Edge), uintptr(Flags), uintptr(unsafePointer(AContentRect)))
}

func (m *TThemeServices) DrawElement(DC HDC, Details *TThemedElementDetails, R *TRect, ClipRect *TRect) {
	LCL().SysCallN(5382, m.Instance(), uintptr(DC), uintptr(unsafePointer(Details)), uintptr(unsafePointer(R)), uintptr(unsafePointer(ClipRect)))
}

func (m *TThemeServices) DrawIcon(DC HDC, Details *TThemedElementDetails, R *TRect, himl HIMAGELIST, Index int32) {
	LCL().SysCallN(5383, m.Instance(), uintptr(DC), uintptr(unsafePointer(Details)), uintptr(unsafePointer(R)), uintptr(himl), uintptr(Index))
}

func (m *TThemeServices) DrawIcon1(ACanvas IPersistent, Details *TThemedElementDetails, P *TPoint, AImageList IPersistent, Index int32, AImageWidth int32, ARefControl IPersistent) {
	LCL().SysCallN(5384, m.Instance(), GetObjectUintptr(ACanvas), uintptr(unsafePointer(Details)), uintptr(unsafePointer(P)), GetObjectUintptr(AImageList), uintptr(Index), uintptr(AImageWidth), GetObjectUintptr(ARefControl))
}

func (m *TThemeServices) DrawParentBackground(Window HWND, Target HDC, Details *PThemedElementDetails, OnlyIfTransparent bool, Bounds *TRect) {
	LCL().SysCallN(5385, m.Instance(), uintptr(Window), uintptr(Target), uintptr(unsafePointer(Details)), PascalBool(OnlyIfTransparent), uintptr(unsafePointer(Bounds)))
}

func (m *TThemeServices) DrawText(DC HDC, Details *TThemedElementDetails, S string, R *TRect, Flags, Flags2 uint32) {
	LCL().SysCallN(5386, m.Instance(), uintptr(DC), uintptr(unsafePointer(Details)), PascalStr(S), uintptr(unsafePointer(R)), uintptr(Flags), uintptr(Flags2))
}

func (m *TThemeServices) DrawText1(ACanvas IPersistent, Details *TThemedElementDetails, S string, R *TRect, Flags, Flags2 uint32) {
	LCL().SysCallN(5387, m.Instance(), GetObjectUintptr(ACanvas), uintptr(unsafePointer(Details)), PascalStr(S), uintptr(unsafePointer(R)), uintptr(Flags), uintptr(Flags2))
}

func (m *TThemeServices) PaintBorder(Control IObject, EraseLRCorner bool) {
	LCL().SysCallN(5426, m.Instance(), GetObjectUintptr(Control), PascalBool(EraseLRCorner))
}

func (m *TThemeServices) UpdateThemes() {
	LCL().SysCallN(5430, m.Instance())
}

func (m *TThemeServices) SetOnThemeChange(fn TNotifyEvent) {
	if m.themeChangePtr != 0 {
		RemoveEventElement(m.themeChangePtr)
	}
	m.themeChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5427, m.Instance(), m.themeChangePtr)
}
