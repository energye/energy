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
	r1 := LCL().SysCallN(4721)
	return AsThemeServices(r1)
}

func (m *TThemeServices) DottedBrush() HBRUSH {
	r1 := LCL().SysCallN(4722, m.Instance())
	return HBRUSH(r1)
}

func (m *TThemeServices) ThemesAvailable() bool {
	r1 := LCL().SysCallN(4770, m.Instance())
	return GoBool(r1)
}

func (m *TThemeServices) ThemesEnabled() bool {
	r1 := LCL().SysCallN(4771, m.Instance())
	return GoBool(r1)
}

func (m *TThemeServices) IsDisabled(Details *TThemedElementDetails) bool {
	r1 := LCL().SysCallN(4764, m.Instance(), uintptr(unsafe.Pointer(Details)))
	return GoBool(r1)
}

func (m *TThemeServices) IsPushed(Details *TThemedElementDetails) bool {
	r1 := LCL().SysCallN(4767, m.Instance(), uintptr(unsafe.Pointer(Details)))
	return GoBool(r1)
}

func (m *TThemeServices) IsHot(Details *TThemedElementDetails) bool {
	r1 := LCL().SysCallN(4765, m.Instance(), uintptr(unsafe.Pointer(Details)))
	return GoBool(r1)
}

func (m *TThemeServices) IsChecked(Details *TThemedElementDetails) bool {
	r1 := LCL().SysCallN(4763, m.Instance(), uintptr(unsafe.Pointer(Details)))
	return GoBool(r1)
}

func (m *TThemeServices) IsMixed(Details *TThemedElementDetails) bool {
	r1 := LCL().SysCallN(4766, m.Instance(), uintptr(unsafe.Pointer(Details)))
	return GoBool(r1)
}

func (m *TThemeServices) GetElementDetails(Detail TThemedButton) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4733, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails1(Detail TThemedClock) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4734, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails2(Detail TThemedComboBox) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4745, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails3(Detail TThemedEdit) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4750, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails4(Detail TThemedExplorerBar) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4751, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails5(Detail TThemedHeader) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4752, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails6(Detail TThemedListView) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4753, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails7(Detail TThemedMenu) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4754, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails8(Detail TThemedPage) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4755, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails9(Detail TThemedProgress) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4756, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails10(Detail TThemedRebar) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4735, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails11(Detail TThemedScrollBar) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4736, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails12(Detail TThemedSpin) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4737, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails13(Detail TThemedStartPanel) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4738, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails14(Detail TThemedStatus) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4739, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails15(Detail TThemedTab) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4740, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails16(Detail TThemedTaskBand) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4741, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails17(Detail TThemedTaskBar) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4742, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails18(Detail TThemedToolBar) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4743, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails19(Detail TThemedToolTip) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4744, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails20(Detail TThemedTrackBar) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4746, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails21(Detail TThemedTrayNotify) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4747, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails22(Detail TThemedTreeview) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4748, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetElementDetails23(Detail TThemedWindow) (resultThemedElementDetails TThemedElementDetails) {
	LCL().SysCallN(4749, m.Instance(), uintptr(Detail), uintptr(unsafe.Pointer(&resultThemedElementDetails)))
	return
}

func (m *TThemeServices) GetDetailSizeForWindow(Details *TThemedElementDetails, AWindow HWND) (resultSize TSize) {
	LCL().SysCallN(4732, m.Instance(), uintptr(unsafe.Pointer(Details)), uintptr(AWindow), uintptr(unsafe.Pointer(&resultSize)))
	return
}

func (m *TThemeServices) GetDetailSizeForPPI(Details *TThemedElementDetails, PPI int32) (resultSize TSize) {
	LCL().SysCallN(4731, m.Instance(), uintptr(unsafe.Pointer(Details)), uintptr(PPI), uintptr(unsafe.Pointer(&resultSize)))
	return
}

func (m *TThemeServices) GetDetailRegion(DC HDC, Details *TThemedElementDetails, R *TRect) HRGN {
	r1 := LCL().SysCallN(4730, m.Instance(), uintptr(DC), uintptr(unsafe.Pointer(Details)), uintptr(unsafe.Pointer(R)))
	return HRGN(r1)
}

func (m *TThemeServices) GetStockImage(StockID int32, OutImage, OutMask *HBITMAP) bool {
	var result1 uintptr
	var result2 uintptr
	r1 := LCL().SysCallN(4758, m.Instance(), uintptr(StockID), uintptr(unsafe.Pointer(&result1)), uintptr(unsafe.Pointer(&result2)))
	*OutImage = HBITMAP(result1)
	*OutMask = HBITMAP(result2)
	return GoBool(r1)
}

func (m *TThemeServices) GetStockImage1(StockID int32, AWidth, AHeight int32, OutImage, OutMask *HBITMAP) bool {
	var result2 uintptr
	var result3 uintptr
	r1 := LCL().SysCallN(4759, m.Instance(), uintptr(StockID), uintptr(AWidth), uintptr(AHeight), uintptr(unsafe.Pointer(&result2)), uintptr(unsafe.Pointer(&result3)))
	*OutImage = HBITMAP(result2)
	*OutMask = HBITMAP(result3)
	return GoBool(r1)
}

func (m *TThemeServices) GetOption(AOption TThemeOption) int32 {
	r1 := LCL().SysCallN(4757, m.Instance(), uintptr(AOption))
	return int32(r1)
}

func (m *TThemeServices) GetTextExtent(DC HDC, Details *TThemedElementDetails, S string, Flags uint32, BoundingRect *TRect) (resultRect TRect) {
	LCL().SysCallN(4760, m.Instance(), uintptr(DC), uintptr(unsafe.Pointer(Details)), PascalStr(S), uintptr(Flags), uintptr(unsafe.Pointer(BoundingRect)), uintptr(unsafe.Pointer(&resultRect)))
	return
}

func (m *TThemeServices) ColorToRGB(Color int32, Details *PThemedElementDetails) COLORREF {
	r1 := LCL().SysCallN(4719, m.Instance(), uintptr(Color), uintptr(unsafe.Pointer(Details)))
	return COLORREF(r1)
}

func (m *TThemeServices) ContentRect(DC HDC, Details *TThemedElementDetails, BoundingRect *TRect) (resultRect TRect) {
	LCL().SysCallN(4720, m.Instance(), uintptr(DC), uintptr(unsafe.Pointer(Details)), uintptr(unsafe.Pointer(BoundingRect)), uintptr(unsafe.Pointer(&resultRect)))
	return
}

func (m *TThemeServices) HasTransparentParts(Details *TThemedElementDetails) bool {
	r1 := LCL().SysCallN(4761, m.Instance(), uintptr(unsafe.Pointer(Details)))
	return GoBool(r1)
}

func ThemeServicesClass() TClass {
	ret := LCL().SysCallN(4718)
	return TClass(ret)
}

func (m *TThemeServices) IntfDoOnThemeChange() {
	LCL().SysCallN(4762, m.Instance())
}

func (m *TThemeServices) DrawEdge(DC HDC, Details *TThemedElementDetails, R *TRect, Edge, Flags uint32, AContentRect *TRect) {
	LCL().SysCallN(4723, m.Instance(), uintptr(DC), uintptr(unsafe.Pointer(Details)), uintptr(unsafe.Pointer(R)), uintptr(Edge), uintptr(Flags), uintptr(unsafe.Pointer(AContentRect)))
}

func (m *TThemeServices) DrawElement(DC HDC, Details *TThemedElementDetails, R *TRect, ClipRect *TRect) {
	LCL().SysCallN(4724, m.Instance(), uintptr(DC), uintptr(unsafe.Pointer(Details)), uintptr(unsafe.Pointer(R)), uintptr(unsafe.Pointer(ClipRect)))
}

func (m *TThemeServices) DrawIcon(DC HDC, Details *TThemedElementDetails, R *TRect, himl HIMAGELIST, Index int32) {
	LCL().SysCallN(4725, m.Instance(), uintptr(DC), uintptr(unsafe.Pointer(Details)), uintptr(unsafe.Pointer(R)), uintptr(himl), uintptr(Index))
}

func (m *TThemeServices) DrawIcon1(ACanvas IPersistent, Details *TThemedElementDetails, P *TPoint, AImageList IPersistent, Index int32, AImageWidth int32, ARefControl IPersistent) {
	LCL().SysCallN(4726, m.Instance(), GetObjectUintptr(ACanvas), uintptr(unsafe.Pointer(Details)), uintptr(unsafe.Pointer(P)), GetObjectUintptr(AImageList), uintptr(Index), uintptr(AImageWidth), GetObjectUintptr(ARefControl))
}

func (m *TThemeServices) DrawParentBackground(Window HWND, Target HDC, Details *PThemedElementDetails, OnlyIfTransparent bool, Bounds *TRect) {
	LCL().SysCallN(4727, m.Instance(), uintptr(Window), uintptr(Target), uintptr(unsafe.Pointer(Details)), PascalBool(OnlyIfTransparent), uintptr(unsafe.Pointer(Bounds)))
}

func (m *TThemeServices) DrawText(DC HDC, Details *TThemedElementDetails, S string, R *TRect, Flags, Flags2 uint32) {
	LCL().SysCallN(4728, m.Instance(), uintptr(DC), uintptr(unsafe.Pointer(Details)), PascalStr(S), uintptr(unsafe.Pointer(R)), uintptr(Flags), uintptr(Flags2))
}

func (m *TThemeServices) DrawText1(ACanvas IPersistent, Details *TThemedElementDetails, S string, R *TRect, Flags, Flags2 uint32) {
	LCL().SysCallN(4729, m.Instance(), GetObjectUintptr(ACanvas), uintptr(unsafe.Pointer(Details)), PascalStr(S), uintptr(unsafe.Pointer(R)), uintptr(Flags), uintptr(Flags2))
}

func (m *TThemeServices) PaintBorder(Control IObject, EraseLRCorner bool) {
	LCL().SysCallN(4768, m.Instance(), GetObjectUintptr(Control), PascalBool(EraseLRCorner))
}

func (m *TThemeServices) UpdateThemes() {
	LCL().SysCallN(4772, m.Instance())
}

func (m *TThemeServices) SetOnThemeChange(fn TNotifyEvent) {
	if m.themeChangePtr != 0 {
		RemoveEventElement(m.themeChangePtr)
	}
	m.themeChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4769, m.Instance(), m.themeChangePtr)
}
