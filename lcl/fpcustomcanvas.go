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

// IFPCustomCanvas Is Abstract Class Parent: IPersistent
type IFPCustomCanvas interface {
	IPersistent
	LockCount() int32                                                     // property
	Font() IFPCustomFont                                                  // property
	SetFont(AValue IFPCustomFont)                                         // property
	Pen() IFPCustomPen                                                    // property
	SetPen(AValue IFPCustomPen)                                           // property
	Brush() IFPCustomBrush                                                // property
	SetBrush(AValue IFPCustomBrush)                                       // property
	Colors(x, y int32) (resultFPColor TFPColor)                           // property
	SetColors(x, y int32, AValue *TFPColor)                               // property
	ClipRect() (resultRect TRect)                                         // property
	SetClipRect(AValue *TRect)                                            // property
	ClipRegion() IFPCustomRegion                                          // property
	SetClipRegion(AValue IFPCustomRegion)                                 // property
	Clipping() bool                                                       // property
	SetClipping(AValue bool)                                              // property
	PenPos() (resultPoint TPoint)                                         // property
	SetPenPos(AValue *TPoint)                                             // property
	Height() int32                                                        // property
	SetHeight(AValue int32)                                               // property
	Width() int32                                                         // property
	SetWidth(AValue int32)                                                // property
	ManageResources() bool                                                // property
	SetManageResources(AValue bool)                                       // property
	DrawingMode() TFPDrawingMode                                          // property
	SetDrawingMode(AValue TFPDrawingMode)                                 // property
	Locked() bool                                                         // function
	CreateFont() IFPCustomFont                                            // function
	CreatePen() IFPCustomPen                                              // function
	CreateBrush() IFPCustomBrush                                          // function
	GetTextHeight(text string) int32                                      // function
	GetTextWidth(text string) int32                                       // function
	TextExtent(Text string) (resultSize TSize)                            // function
	TextHeight(Text string) int32                                         // function
	TextWidth(Text string) int32                                          // function
	GetTextHeight1(text string) int32                                     // function
	GetTextWidth1(text string) int32                                      // function
	TextExtent1(Text string) (resultSize TSize)                           // function
	TextHeight1(Text string) int32                                        // function
	TextWidth1(Text string) int32                                         // function
	LockCanvas()                                                          // procedure
	UnlockCanvas()                                                        // procedure
	TextOut(x, y int32, text string)                                      // procedure
	GetTextSize(text string, w, h *int32)                                 // procedure
	TextOut1(x, y int32, text string)                                     // procedure
	GetTextSize1(text string, w, h *int32)                                // procedure
	Arc(ALeft, ATop, ARight, ABottom, Angle16Deg, Angle16DegLength int32) // procedure
	Arc1(ALeft, ATop, ARight, ABottom, SX, SY, EX, EY int32)              // procedure
	Ellipse(Bounds *TRect)                                                // procedure
	Ellipse1(left, top, right, bottom int32)                              // procedure
	EllipseC(x, y int32, rx, ry uint32)                                   // procedure
	RadialPie(x1, y1, x2, y2, StartAngle16Deg, Angle16DegLength int32)    // procedure
	PolyBezier(Points []TPoint, Filled bool, Continuous bool)             // procedure
	Rectangle(Bounds *TRect)                                              // procedure
	Rectangle1(left, top, right, bottom int32)                            // procedure
	FillRect(ARect *TRect)                                                // procedure
	FillRect1(X1, Y1, X2, Y2 int32)                                       // procedure
	FloodFill(x, y int32)                                                 // procedure
	Clear()                                                               // procedure
	MoveTo(x, y int32)                                                    // procedure
	MoveTo1(p *TPoint)                                                    // procedure
	LineTo(x, y int32)                                                    // procedure
	LineTo1(p *TPoint)                                                    // procedure
	Line(x1, y1, x2, y2 int32)                                            // procedure
	Line1(p1, p2 *TPoint)                                                 // procedure
	Line2(points *TRect)                                                  // procedure
	CopyRect(x, y int32, canvas IFPCustomCanvas, SourceRect *TRect)       // procedure
	Draw(x, y int32, image IFPCustomImage)                                // procedure
	StretchDraw(x, y, w, h int32, source IFPCustomImage)                  // procedure
	Erase()                                                               // procedure
	DrawPixel(x, y int32, newcolor *TFPColor)                             // procedure
	SetOnCombineColors(fn TFPCanvasCombineColors)                         // property event
}

// TFPCustomCanvas Is Abstract Class Parent: TPersistent
type TFPCustomCanvas struct {
	TPersistent
	combineColorsPtr uintptr
}

func (m *TFPCustomCanvas) LockCount() int32 {
	r1 := LCL().SysCallN(2873, m.Instance())
	return int32(r1)
}

func (m *TFPCustomCanvas) Font() IFPCustomFont {
	r1 := LCL().SysCallN(2859, 0, m.Instance(), 0)
	return AsFPCustomFont(r1)
}

func (m *TFPCustomCanvas) SetFont(AValue IFPCustomFont) {
	LCL().SysCallN(2859, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFPCustomCanvas) Pen() IFPCustomPen {
	r1 := LCL().SysCallN(2878, 0, m.Instance(), 0)
	return AsFPCustomPen(r1)
}

func (m *TFPCustomCanvas) SetPen(AValue IFPCustomPen) {
	LCL().SysCallN(2878, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFPCustomCanvas) Brush() IFPCustomBrush {
	r1 := LCL().SysCallN(2838, 0, m.Instance(), 0)
	return AsFPCustomBrush(r1)
}

func (m *TFPCustomCanvas) SetBrush(AValue IFPCustomBrush) {
	LCL().SysCallN(2838, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFPCustomCanvas) Colors(x, y int32) (resultFPColor TFPColor) {
	r1 := LCL().SysCallN(2844, 0, m.Instance(), uintptr(x), uintptr(y))
	return *(*TFPColor)(getPointer(r1))
}

func (m *TFPCustomCanvas) SetColors(x, y int32, AValue *TFPColor) {
	LCL().SysCallN(2844, 1, m.Instance(), uintptr(x), uintptr(y), uintptr(unsafePointer(AValue)))
}

func (m *TFPCustomCanvas) ClipRect() (resultRect TRect) {
	LCL().SysCallN(2841, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TFPCustomCanvas) SetClipRect(AValue *TRect) {
	LCL().SysCallN(2841, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TFPCustomCanvas) ClipRegion() IFPCustomRegion {
	r1 := LCL().SysCallN(2842, 0, m.Instance(), 0)
	return AsFPCustomRegion(r1)
}

func (m *TFPCustomCanvas) SetClipRegion(AValue IFPCustomRegion) {
	LCL().SysCallN(2842, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFPCustomCanvas) Clipping() bool {
	r1 := LCL().SysCallN(2843, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomCanvas) SetClipping(AValue bool) {
	LCL().SysCallN(2843, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomCanvas) PenPos() (resultPoint TPoint) {
	LCL().SysCallN(2879, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TFPCustomCanvas) SetPenPos(AValue *TPoint) {
	LCL().SysCallN(2879, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TFPCustomCanvas) Height() int32 {
	r1 := LCL().SysCallN(2866, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomCanvas) SetHeight(AValue int32) {
	LCL().SysCallN(2866, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomCanvas) Width() int32 {
	r1 := LCL().SysCallN(2895, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomCanvas) SetWidth(AValue int32) {
	LCL().SysCallN(2895, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomCanvas) ManageResources() bool {
	r1 := LCL().SysCallN(2875, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomCanvas) SetManageResources(AValue bool) {
	LCL().SysCallN(2875, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomCanvas) DrawingMode() TFPDrawingMode {
	r1 := LCL().SysCallN(2851, 0, m.Instance(), 0)
	return TFPDrawingMode(r1)
}

func (m *TFPCustomCanvas) SetDrawingMode(AValue TFPDrawingMode) {
	LCL().SysCallN(2851, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomCanvas) Locked() bool {
	r1 := LCL().SysCallN(2874, m.Instance())
	return GoBool(r1)
}

func (m *TFPCustomCanvas) CreateFont() IFPCustomFont {
	r1 := LCL().SysCallN(2847, m.Instance())
	return AsFPCustomFont(r1)
}

func (m *TFPCustomCanvas) CreatePen() IFPCustomPen {
	r1 := LCL().SysCallN(2848, m.Instance())
	return AsFPCustomPen(r1)
}

func (m *TFPCustomCanvas) CreateBrush() IFPCustomBrush {
	r1 := LCL().SysCallN(2846, m.Instance())
	return AsFPCustomBrush(r1)
}

func (m *TFPCustomCanvas) GetTextHeight(text string) int32 {
	r1 := LCL().SysCallN(2860, m.Instance(), PascalStr(text))
	return int32(r1)
}

func (m *TFPCustomCanvas) GetTextWidth(text string) int32 {
	r1 := LCL().SysCallN(2864, m.Instance(), PascalStr(text))
	return int32(r1)
}

func (m *TFPCustomCanvas) TextExtent(Text string) (resultSize TSize) {
	LCL().SysCallN(2886, m.Instance(), PascalStr(Text), uintptr(unsafePointer(&resultSize)))
	return
}

func (m *TFPCustomCanvas) TextHeight(Text string) int32 {
	r1 := LCL().SysCallN(2888, m.Instance(), PascalStr(Text))
	return int32(r1)
}

func (m *TFPCustomCanvas) TextWidth(Text string) int32 {
	r1 := LCL().SysCallN(2892, m.Instance(), PascalStr(Text))
	return int32(r1)
}

func (m *TFPCustomCanvas) GetTextHeight1(text string) int32 {
	r1 := LCL().SysCallN(2861, m.Instance(), PascalStr(text))
	return int32(r1)
}

func (m *TFPCustomCanvas) GetTextWidth1(text string) int32 {
	r1 := LCL().SysCallN(2865, m.Instance(), PascalStr(text))
	return int32(r1)
}

func (m *TFPCustomCanvas) TextExtent1(Text string) (resultSize TSize) {
	LCL().SysCallN(2887, m.Instance(), PascalStr(Text), uintptr(unsafePointer(&resultSize)))
	return
}

func (m *TFPCustomCanvas) TextHeight1(Text string) int32 {
	r1 := LCL().SysCallN(2889, m.Instance(), PascalStr(Text))
	return int32(r1)
}

func (m *TFPCustomCanvas) TextWidth1(Text string) int32 {
	r1 := LCL().SysCallN(2893, m.Instance(), PascalStr(Text))
	return int32(r1)
}

func FPCustomCanvasClass() TClass {
	ret := LCL().SysCallN(2839)
	return TClass(ret)
}

func (m *TFPCustomCanvas) LockCanvas() {
	LCL().SysCallN(2872, m.Instance())
}

func (m *TFPCustomCanvas) UnlockCanvas() {
	LCL().SysCallN(2894, m.Instance())
}

func (m *TFPCustomCanvas) TextOut(x, y int32, text string) {
	LCL().SysCallN(2890, m.Instance(), uintptr(x), uintptr(y), PascalStr(text))
}

func (m *TFPCustomCanvas) GetTextSize(text string, w, h *int32) {
	var result1 uintptr
	var result2 uintptr
	LCL().SysCallN(2862, m.Instance(), PascalStr(text), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*w = int32(result1)
	*h = int32(result2)
}

func (m *TFPCustomCanvas) TextOut1(x, y int32, text string) {
	LCL().SysCallN(2891, m.Instance(), uintptr(x), uintptr(y), PascalStr(text))
}

func (m *TFPCustomCanvas) GetTextSize1(text string, w, h *int32) {
	var result1 uintptr
	var result2 uintptr
	LCL().SysCallN(2863, m.Instance(), PascalStr(text), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*w = int32(result1)
	*h = int32(result2)
}

func (m *TFPCustomCanvas) Arc(ALeft, ATop, ARight, ABottom, Angle16Deg, Angle16DegLength int32) {
	LCL().SysCallN(2836, m.Instance(), uintptr(ALeft), uintptr(ATop), uintptr(ARight), uintptr(ABottom), uintptr(Angle16Deg), uintptr(Angle16DegLength))
}

func (m *TFPCustomCanvas) Arc1(ALeft, ATop, ARight, ABottom, SX, SY, EX, EY int32) {
	LCL().SysCallN(2837, m.Instance(), uintptr(ALeft), uintptr(ATop), uintptr(ARight), uintptr(ABottom), uintptr(SX), uintptr(SY), uintptr(EX), uintptr(EY))
}

func (m *TFPCustomCanvas) Ellipse(Bounds *TRect) {
	LCL().SysCallN(2852, m.Instance(), uintptr(unsafePointer(Bounds)))
}

func (m *TFPCustomCanvas) Ellipse1(left, top, right, bottom int32) {
	LCL().SysCallN(2853, m.Instance(), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
}

func (m *TFPCustomCanvas) EllipseC(x, y int32, rx, ry uint32) {
	LCL().SysCallN(2854, m.Instance(), uintptr(x), uintptr(y), uintptr(rx), uintptr(ry))
}

func (m *TFPCustomCanvas) RadialPie(x1, y1, x2, y2, StartAngle16Deg, Angle16DegLength int32) {
	LCL().SysCallN(2881, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(StartAngle16Deg), uintptr(Angle16DegLength))
}

func (m *TFPCustomCanvas) PolyBezier(Points []TPoint, Filled bool, Continuous bool) {
	sysCallPoint(2880, m.Instance(), Points, PascalBool(Filled), PascalBool(Continuous))
}

func (m *TFPCustomCanvas) Rectangle(Bounds *TRect) {
	LCL().SysCallN(2882, m.Instance(), uintptr(unsafePointer(Bounds)))
}

func (m *TFPCustomCanvas) Rectangle1(left, top, right, bottom int32) {
	LCL().SysCallN(2883, m.Instance(), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
}

func (m *TFPCustomCanvas) FillRect(ARect *TRect) {
	LCL().SysCallN(2856, m.Instance(), uintptr(unsafePointer(ARect)))
}

func (m *TFPCustomCanvas) FillRect1(X1, Y1, X2, Y2 int32) {
	LCL().SysCallN(2857, m.Instance(), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
}

func (m *TFPCustomCanvas) FloodFill(x, y int32) {
	LCL().SysCallN(2858, m.Instance(), uintptr(x), uintptr(y))
}

func (m *TFPCustomCanvas) Clear() {
	LCL().SysCallN(2840, m.Instance())
}

func (m *TFPCustomCanvas) MoveTo(x, y int32) {
	LCL().SysCallN(2876, m.Instance(), uintptr(x), uintptr(y))
}

func (m *TFPCustomCanvas) MoveTo1(p *TPoint) {
	LCL().SysCallN(2877, m.Instance(), uintptr(unsafePointer(p)))
}

func (m *TFPCustomCanvas) LineTo(x, y int32) {
	LCL().SysCallN(2870, m.Instance(), uintptr(x), uintptr(y))
}

func (m *TFPCustomCanvas) LineTo1(p *TPoint) {
	LCL().SysCallN(2871, m.Instance(), uintptr(unsafePointer(p)))
}

func (m *TFPCustomCanvas) Line(x1, y1, x2, y2 int32) {
	LCL().SysCallN(2867, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
}

func (m *TFPCustomCanvas) Line1(p1, p2 *TPoint) {
	LCL().SysCallN(2868, m.Instance(), uintptr(unsafePointer(p1)), uintptr(unsafePointer(p2)))
}

func (m *TFPCustomCanvas) Line2(points *TRect) {
	LCL().SysCallN(2869, m.Instance(), uintptr(unsafePointer(points)))
}

func (m *TFPCustomCanvas) CopyRect(x, y int32, canvas IFPCustomCanvas, SourceRect *TRect) {
	LCL().SysCallN(2845, m.Instance(), uintptr(x), uintptr(y), GetObjectUintptr(canvas), uintptr(unsafePointer(SourceRect)))
}

func (m *TFPCustomCanvas) Draw(x, y int32, image IFPCustomImage) {
	LCL().SysCallN(2849, m.Instance(), uintptr(x), uintptr(y), GetObjectUintptr(image))
}

func (m *TFPCustomCanvas) StretchDraw(x, y, w, h int32, source IFPCustomImage) {
	LCL().SysCallN(2885, m.Instance(), uintptr(x), uintptr(y), uintptr(w), uintptr(h), GetObjectUintptr(source))
}

func (m *TFPCustomCanvas) Erase() {
	LCL().SysCallN(2855, m.Instance())
}

func (m *TFPCustomCanvas) DrawPixel(x, y int32, newcolor *TFPColor) {
	LCL().SysCallN(2850, m.Instance(), uintptr(x), uintptr(y), uintptr(unsafePointer(newcolor)))
}

func (m *TFPCustomCanvas) SetOnCombineColors(fn TFPCanvasCombineColors) {
	if m.combineColorsPtr != 0 {
		RemoveEventElement(m.combineColorsPtr)
	}
	m.combineColorsPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2884, m.Instance(), m.combineColorsPtr)
}
