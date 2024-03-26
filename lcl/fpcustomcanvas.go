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
	r1 := LCL().SysCallN(2630, m.Instance())
	return int32(r1)
}

func (m *TFPCustomCanvas) Font() IFPCustomFont {
	r1 := LCL().SysCallN(2616, 0, m.Instance(), 0)
	return AsFPCustomFont(r1)
}

func (m *TFPCustomCanvas) SetFont(AValue IFPCustomFont) {
	LCL().SysCallN(2616, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFPCustomCanvas) Pen() IFPCustomPen {
	r1 := LCL().SysCallN(2635, 0, m.Instance(), 0)
	return AsFPCustomPen(r1)
}

func (m *TFPCustomCanvas) SetPen(AValue IFPCustomPen) {
	LCL().SysCallN(2635, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFPCustomCanvas) Brush() IFPCustomBrush {
	r1 := LCL().SysCallN(2595, 0, m.Instance(), 0)
	return AsFPCustomBrush(r1)
}

func (m *TFPCustomCanvas) SetBrush(AValue IFPCustomBrush) {
	LCL().SysCallN(2595, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFPCustomCanvas) Colors(x, y int32) (resultFPColor TFPColor) {
	r1 := LCL().SysCallN(2601, 0, m.Instance(), uintptr(x), uintptr(y))
	return *(*TFPColor)(getPointer(r1))
}

func (m *TFPCustomCanvas) SetColors(x, y int32, AValue *TFPColor) {
	LCL().SysCallN(2601, 1, m.Instance(), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(AValue)))
}

func (m *TFPCustomCanvas) ClipRect() (resultRect TRect) {
	LCL().SysCallN(2598, 0, m.Instance(), uintptr(unsafe.Pointer(&resultRect)), uintptr(unsafe.Pointer(&resultRect)))
	return
}

func (m *TFPCustomCanvas) SetClipRect(AValue *TRect) {
	LCL().SysCallN(2598, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TFPCustomCanvas) ClipRegion() IFPCustomRegion {
	r1 := LCL().SysCallN(2599, 0, m.Instance(), 0)
	return AsFPCustomRegion(r1)
}

func (m *TFPCustomCanvas) SetClipRegion(AValue IFPCustomRegion) {
	LCL().SysCallN(2599, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFPCustomCanvas) Clipping() bool {
	r1 := LCL().SysCallN(2600, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomCanvas) SetClipping(AValue bool) {
	LCL().SysCallN(2600, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomCanvas) PenPos() (resultPoint TPoint) {
	LCL().SysCallN(2636, 0, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TFPCustomCanvas) SetPenPos(AValue *TPoint) {
	LCL().SysCallN(2636, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TFPCustomCanvas) Height() int32 {
	r1 := LCL().SysCallN(2623, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomCanvas) SetHeight(AValue int32) {
	LCL().SysCallN(2623, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomCanvas) Width() int32 {
	r1 := LCL().SysCallN(2652, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomCanvas) SetWidth(AValue int32) {
	LCL().SysCallN(2652, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomCanvas) ManageResources() bool {
	r1 := LCL().SysCallN(2632, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomCanvas) SetManageResources(AValue bool) {
	LCL().SysCallN(2632, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomCanvas) DrawingMode() TFPDrawingMode {
	r1 := LCL().SysCallN(2608, 0, m.Instance(), 0)
	return TFPDrawingMode(r1)
}

func (m *TFPCustomCanvas) SetDrawingMode(AValue TFPDrawingMode) {
	LCL().SysCallN(2608, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomCanvas) Locked() bool {
	r1 := LCL().SysCallN(2631, m.Instance())
	return GoBool(r1)
}

func (m *TFPCustomCanvas) CreateFont() IFPCustomFont {
	r1 := LCL().SysCallN(2604, m.Instance())
	return AsFPCustomFont(r1)
}

func (m *TFPCustomCanvas) CreatePen() IFPCustomPen {
	r1 := LCL().SysCallN(2605, m.Instance())
	return AsFPCustomPen(r1)
}

func (m *TFPCustomCanvas) CreateBrush() IFPCustomBrush {
	r1 := LCL().SysCallN(2603, m.Instance())
	return AsFPCustomBrush(r1)
}

func (m *TFPCustomCanvas) GetTextHeight(text string) int32 {
	r1 := LCL().SysCallN(2617, m.Instance(), PascalStr(text))
	return int32(r1)
}

func (m *TFPCustomCanvas) GetTextWidth(text string) int32 {
	r1 := LCL().SysCallN(2621, m.Instance(), PascalStr(text))
	return int32(r1)
}

func (m *TFPCustomCanvas) TextExtent(Text string) (resultSize TSize) {
	LCL().SysCallN(2643, m.Instance(), PascalStr(Text), uintptr(unsafe.Pointer(&resultSize)))
	return
}

func (m *TFPCustomCanvas) TextHeight(Text string) int32 {
	r1 := LCL().SysCallN(2645, m.Instance(), PascalStr(Text))
	return int32(r1)
}

func (m *TFPCustomCanvas) TextWidth(Text string) int32 {
	r1 := LCL().SysCallN(2649, m.Instance(), PascalStr(Text))
	return int32(r1)
}

func (m *TFPCustomCanvas) GetTextHeight1(text string) int32 {
	r1 := LCL().SysCallN(2618, m.Instance(), PascalStr(text))
	return int32(r1)
}

func (m *TFPCustomCanvas) GetTextWidth1(text string) int32 {
	r1 := LCL().SysCallN(2622, m.Instance(), PascalStr(text))
	return int32(r1)
}

func (m *TFPCustomCanvas) TextExtent1(Text string) (resultSize TSize) {
	LCL().SysCallN(2644, m.Instance(), PascalStr(Text), uintptr(unsafe.Pointer(&resultSize)))
	return
}

func (m *TFPCustomCanvas) TextHeight1(Text string) int32 {
	r1 := LCL().SysCallN(2646, m.Instance(), PascalStr(Text))
	return int32(r1)
}

func (m *TFPCustomCanvas) TextWidth1(Text string) int32 {
	r1 := LCL().SysCallN(2650, m.Instance(), PascalStr(Text))
	return int32(r1)
}

func FPCustomCanvasClass() TClass {
	ret := LCL().SysCallN(2596)
	return TClass(ret)
}

func (m *TFPCustomCanvas) LockCanvas() {
	LCL().SysCallN(2629, m.Instance())
}

func (m *TFPCustomCanvas) UnlockCanvas() {
	LCL().SysCallN(2651, m.Instance())
}

func (m *TFPCustomCanvas) TextOut(x, y int32, text string) {
	LCL().SysCallN(2647, m.Instance(), uintptr(x), uintptr(y), PascalStr(text))
}

func (m *TFPCustomCanvas) GetTextSize(text string, w, h *int32) {
	var result1 uintptr
	var result2 uintptr
	LCL().SysCallN(2619, m.Instance(), PascalStr(text), uintptr(unsafe.Pointer(&result1)), uintptr(unsafe.Pointer(&result2)))
	*w = int32(result1)
	*h = int32(result2)
}

func (m *TFPCustomCanvas) TextOut1(x, y int32, text string) {
	LCL().SysCallN(2648, m.Instance(), uintptr(x), uintptr(y), PascalStr(text))
}

func (m *TFPCustomCanvas) GetTextSize1(text string, w, h *int32) {
	var result1 uintptr
	var result2 uintptr
	LCL().SysCallN(2620, m.Instance(), PascalStr(text), uintptr(unsafe.Pointer(&result1)), uintptr(unsafe.Pointer(&result2)))
	*w = int32(result1)
	*h = int32(result2)
}

func (m *TFPCustomCanvas) Arc(ALeft, ATop, ARight, ABottom, Angle16Deg, Angle16DegLength int32) {
	LCL().SysCallN(2593, m.Instance(), uintptr(ALeft), uintptr(ATop), uintptr(ARight), uintptr(ABottom), uintptr(Angle16Deg), uintptr(Angle16DegLength))
}

func (m *TFPCustomCanvas) Arc1(ALeft, ATop, ARight, ABottom, SX, SY, EX, EY int32) {
	LCL().SysCallN(2594, m.Instance(), uintptr(ALeft), uintptr(ATop), uintptr(ARight), uintptr(ABottom), uintptr(SX), uintptr(SY), uintptr(EX), uintptr(EY))
}

func (m *TFPCustomCanvas) Ellipse(Bounds *TRect) {
	LCL().SysCallN(2609, m.Instance(), uintptr(unsafe.Pointer(Bounds)))
}

func (m *TFPCustomCanvas) Ellipse1(left, top, right, bottom int32) {
	LCL().SysCallN(2610, m.Instance(), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
}

func (m *TFPCustomCanvas) EllipseC(x, y int32, rx, ry uint32) {
	LCL().SysCallN(2611, m.Instance(), uintptr(x), uintptr(y), uintptr(rx), uintptr(ry))
}

func (m *TFPCustomCanvas) RadialPie(x1, y1, x2, y2, StartAngle16Deg, Angle16DegLength int32) {
	LCL().SysCallN(2638, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(StartAngle16Deg), uintptr(Angle16DegLength))
}

func (m *TFPCustomCanvas) PolyBezier(Points []TPoint, Filled bool, Continuous bool) {
	sysCallPoint(2637, m.Instance(), Points, PascalBool(Filled), PascalBool(Continuous))
}

func (m *TFPCustomCanvas) Rectangle(Bounds *TRect) {
	LCL().SysCallN(2639, m.Instance(), uintptr(unsafe.Pointer(Bounds)))
}

func (m *TFPCustomCanvas) Rectangle1(left, top, right, bottom int32) {
	LCL().SysCallN(2640, m.Instance(), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
}

func (m *TFPCustomCanvas) FillRect(ARect *TRect) {
	LCL().SysCallN(2613, m.Instance(), uintptr(unsafe.Pointer(ARect)))
}

func (m *TFPCustomCanvas) FillRect1(X1, Y1, X2, Y2 int32) {
	LCL().SysCallN(2614, m.Instance(), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
}

func (m *TFPCustomCanvas) FloodFill(x, y int32) {
	LCL().SysCallN(2615, m.Instance(), uintptr(x), uintptr(y))
}

func (m *TFPCustomCanvas) Clear() {
	LCL().SysCallN(2597, m.Instance())
}

func (m *TFPCustomCanvas) MoveTo(x, y int32) {
	LCL().SysCallN(2633, m.Instance(), uintptr(x), uintptr(y))
}

func (m *TFPCustomCanvas) MoveTo1(p *TPoint) {
	LCL().SysCallN(2634, m.Instance(), uintptr(unsafe.Pointer(p)))
}

func (m *TFPCustomCanvas) LineTo(x, y int32) {
	LCL().SysCallN(2627, m.Instance(), uintptr(x), uintptr(y))
}

func (m *TFPCustomCanvas) LineTo1(p *TPoint) {
	LCL().SysCallN(2628, m.Instance(), uintptr(unsafe.Pointer(p)))
}

func (m *TFPCustomCanvas) Line(x1, y1, x2, y2 int32) {
	LCL().SysCallN(2624, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
}

func (m *TFPCustomCanvas) Line1(p1, p2 *TPoint) {
	LCL().SysCallN(2625, m.Instance(), uintptr(unsafe.Pointer(p1)), uintptr(unsafe.Pointer(p2)))
}

func (m *TFPCustomCanvas) Line2(points *TRect) {
	LCL().SysCallN(2626, m.Instance(), uintptr(unsafe.Pointer(points)))
}

func (m *TFPCustomCanvas) CopyRect(x, y int32, canvas IFPCustomCanvas, SourceRect *TRect) {
	LCL().SysCallN(2602, m.Instance(), uintptr(x), uintptr(y), GetObjectUintptr(canvas), uintptr(unsafe.Pointer(SourceRect)))
}

func (m *TFPCustomCanvas) Draw(x, y int32, image IFPCustomImage) {
	LCL().SysCallN(2606, m.Instance(), uintptr(x), uintptr(y), GetObjectUintptr(image))
}

func (m *TFPCustomCanvas) StretchDraw(x, y, w, h int32, source IFPCustomImage) {
	LCL().SysCallN(2642, m.Instance(), uintptr(x), uintptr(y), uintptr(w), uintptr(h), GetObjectUintptr(source))
}

func (m *TFPCustomCanvas) Erase() {
	LCL().SysCallN(2612, m.Instance())
}

func (m *TFPCustomCanvas) DrawPixel(x, y int32, newcolor *TFPColor) {
	LCL().SysCallN(2607, m.Instance(), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(newcolor)))
}

func (m *TFPCustomCanvas) SetOnCombineColors(fn TFPCanvasCombineColors) {
	if m.combineColorsPtr != 0 {
		RemoveEventElement(m.combineColorsPtr)
	}
	m.combineColorsPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2641, m.Instance(), m.combineColorsPtr)
}
