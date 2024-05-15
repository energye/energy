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

// ICanvas Parent: IFPCustomCanvas
type ICanvas interface {
	IFPCustomCanvas
	TextRect2(aRect *TRect, text string, textFormat TTextFormat)
	Pixels(X, Y int32) TColor                                                                  // property
	SetPixels(X, Y int32, AValue TColor)                                                       // property
	Handle() HDC                                                                               // property
	SetHandle(AValue HDC)                                                                      // property
	TextStyle() (resultTextStyle TTextStyle)                                                   // property
	SetTextStyle(AValue *TTextStyle)                                                           // property
	AntialiasingMode() TAntialiasingMode                                                       // property
	SetAntialiasingMode(AValue TAntialiasingMode)                                              // property
	AutoRedraw() bool                                                                          // property
	SetAutoRedraw(AValue bool)                                                                 // property
	BrushForBrush() IBrush                                                                     // property
	SetBrushForBrush(AValue IBrush)                                                            // property
	CopyMode() TCopyMode                                                                       // property
	SetCopyMode(AValue TCopyMode)                                                              // property
	FontForFont() IFont                                                                        // property
	SetFontForFont(AValue IFont)                                                               // property
	PenForPen() IPen                                                                           // property
	SetPenForPen(AValue IPen)                                                                  // property
	Region() IRegion                                                                           // property
	SetRegion(AValue IRegion)                                                                  // property
	TryLock() bool                                                                             // function
	GetTextMetrics(OutTM *TLCLTextMetric) bool                                                 // function
	TextFitInfo(Text string, MaxWidth int32) int32                                             // function
	HandleAllocated() bool                                                                     // function
	GetUpdatedHandle(ReqState TCanvasState) HDC                                                // function
	Lock()                                                                                     // procedure
	Unlock()                                                                                   // procedure
	Refresh()                                                                                  // procedure
	Changing()                                                                                 // procedure
	Changed()                                                                                  // procedure
	SaveHandleState()                                                                          // procedure
	RestoreHandleState()                                                                       // procedure
	ArcTo(ALeft, ATop, ARight, ABottom, SX, SY, EX, EY int32)                                  // procedure
	AngleArc(X, Y int32, Radius uint32, StartAngle, SweepAngle float32)                        // procedure
	BrushCopy(ADestRect *TRect, ABitmap IBitmap, ASourceRect *TRect, ATransparentColor TColor) // procedure
	Chord(x1, y1, x2, y2, Angle16Deg, Angle16DegLength int32)                                  // procedure
	Chord1(x1, y1, x2, y2, SX, SY, EX, EY int32)                                               // procedure
	CopyRectForRect(Dest *TRect, SrcCanvas ICanvas, Source *TRect)                             // procedure
	DrawForGraphic(X, Y int32, SrcGraphic IGraphic)                                            // procedure
	DrawFocusRect(ARect *TRect)                                                                // procedure
	StretchDrawForRect(DestRect *TRect, SrcGraphic IGraphic)                                   // procedure
	FloodFillForColor(X, Y int32, FillColor TColor, FillStyle TFillStyle)                      // procedure
	Frame3d(ARect *TRect, FrameWidth int32, Style TGraphicsBevelCut)                           // procedure
	Frame3D1(ARect *TRect, TopColor, BottomColor TColor, FrameWidth int32)                     // procedure
	Frame(ARect *TRect)                                                                        // procedure
	Frame1(X1, Y1, X2, Y2 int32)                                                               // procedure
	FrameRect(ARect *TRect)                                                                    // procedure
	FrameRect1(X1, Y1, X2, Y2 int32)                                                           // procedure
	GradientFill(ARect *TRect, AStart, AStop TColor, ADirection TGradientDirection)            // procedure
	Pie(EllipseX1, EllipseY1, EllipseX2, EllipseY2, StartX, StartY, EndX, EndY int32)          // procedure
	Polygon(Points []TPoint, Winding bool)                                                     // procedure
	Polyline(Points []TPoint)                                                                  // procedure
	RoundRect(X1, Y1, X2, Y2 int32, RX, RY int32)                                              // procedure
	RoundRect1(Rect *TRect, RX, RY int32)                                                      // procedure
	TextRect(ARect *TRect, X, Y int32, Text string)                                            // procedure
	TextRect1(ARect *TRect, X, Y int32, Text string, Style *TTextStyle)                        // procedure
	SetOnChange(fn TNotifyEvent)                                                               // property event
	SetOnChanging(fn TNotifyEvent)                                                             // property event
}

// TCanvas Parent: TFPCustomCanvas
type TCanvas struct {
	TFPCustomCanvas
	changePtr   uintptr
	changingPtr uintptr
}

func NewCanvas() ICanvas {
	r1 := LCL().SysCallN(523)
	return AsCanvas(r1)
}

func (m *TCanvas) Pixels(X, Y int32) TColor {
	r1 := LCL().SysCallN(542, 0, m.Instance(), uintptr(X), uintptr(Y))
	return TColor(r1)
}

func (m *TCanvas) SetPixels(X, Y int32, AValue TColor) {
	LCL().SysCallN(542, 1, m.Instance(), uintptr(X), uintptr(Y), uintptr(AValue))
}

func (m *TCanvas) Handle() HDC {
	r1 := LCL().SysCallN(537, 0, m.Instance(), 0)
	return HDC(r1)
}

func (m *TCanvas) SetHandle(AValue HDC) {
	LCL().SysCallN(537, 1, m.Instance(), uintptr(AValue))
}

func (m *TCanvas) TextStyle() (resultTextStyle TTextStyle) {
	LCL().SysCallN(557, 0, m.Instance(), uintptr(unsafePointer(&resultTextStyle)), uintptr(unsafePointer(&resultTextStyle)))
	return
}

func (m *TCanvas) SetTextStyle(AValue *TTextStyle) {
	LCL().SysCallN(557, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCanvas) AntialiasingMode() TAntialiasingMode {
	r1 := LCL().SysCallN(511, 0, m.Instance(), 0)
	return TAntialiasingMode(r1)
}

func (m *TCanvas) SetAntialiasingMode(AValue TAntialiasingMode) {
	LCL().SysCallN(511, 1, m.Instance(), uintptr(AValue))
}

func (m *TCanvas) AutoRedraw() bool {
	r1 := LCL().SysCallN(513, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCanvas) SetAutoRedraw(AValue bool) {
	LCL().SysCallN(513, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCanvas) BrushForBrush() IBrush {
	r1 := LCL().SysCallN(515, 0, m.Instance(), 0)
	return AsBrush(r1)
}

func (m *TCanvas) SetBrushForBrush(AValue IBrush) {
	LCL().SysCallN(515, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCanvas) CopyMode() TCopyMode {
	r1 := LCL().SysCallN(521, 0, m.Instance(), 0)
	return TCopyMode(r1)
}

func (m *TCanvas) SetCopyMode(AValue TCopyMode) {
	LCL().SysCallN(521, 1, m.Instance(), uintptr(AValue))
}

func (m *TCanvas) FontForFont() IFont {
	r1 := LCL().SysCallN(527, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TCanvas) SetFontForFont(AValue IFont) {
	LCL().SysCallN(527, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCanvas) PenForPen() IPen {
	r1 := LCL().SysCallN(540, 0, m.Instance(), 0)
	return AsPen(r1)
}

func (m *TCanvas) SetPenForPen(AValue IPen) {
	LCL().SysCallN(540, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCanvas) Region() IRegion {
	r1 := LCL().SysCallN(546, 0, m.Instance(), 0)
	return AsRegion(r1)
}

func (m *TCanvas) SetRegion(AValue IRegion) {
	LCL().SysCallN(546, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCanvas) TryLock() bool {
	r1 := LCL().SysCallN(558, m.Instance())
	return GoBool(r1)
}

func (m *TCanvas) GetTextMetrics(OutTM *TLCLTextMetric) bool {
	var result0 uintptr
	r1 := LCL().SysCallN(534, m.Instance(), uintptr(unsafePointer(&result0)))
	*OutTM = *(*TLCLTextMetric)(getPointer(result0))
	return GoBool(r1)
}

func (m *TCanvas) TextFitInfo(Text string, MaxWidth int32) int32 {
	r1 := LCL().SysCallN(554, m.Instance(), PascalStr(Text), uintptr(MaxWidth))
	return int32(r1)
}

func (m *TCanvas) HandleAllocated() bool {
	r1 := LCL().SysCallN(538, m.Instance())
	return GoBool(r1)
}

func (m *TCanvas) GetUpdatedHandle(ReqState TCanvasState) HDC {
	r1 := LCL().SysCallN(535, m.Instance(), uintptr(ReqState))
	return HDC(r1)
}

func CanvasClass() TClass {
	ret := LCL().SysCallN(520)
	return TClass(ret)
}

func (m *TCanvas) Lock() {
	LCL().SysCallN(539, m.Instance())
}

func (m *TCanvas) Unlock() {
	LCL().SysCallN(559, m.Instance())
}

func (m *TCanvas) Refresh() {
	LCL().SysCallN(545, m.Instance())
}

func (m *TCanvas) Changing() {
	LCL().SysCallN(517, m.Instance())
}

func (m *TCanvas) Changed() {
	LCL().SysCallN(516, m.Instance())
}

func (m *TCanvas) SaveHandleState() {
	LCL().SysCallN(550, m.Instance())
}

func (m *TCanvas) RestoreHandleState() {
	LCL().SysCallN(547, m.Instance())
}

func (m *TCanvas) ArcTo(ALeft, ATop, ARight, ABottom, SX, SY, EX, EY int32) {
	LCL().SysCallN(512, m.Instance(), uintptr(ALeft), uintptr(ATop), uintptr(ARight), uintptr(ABottom), uintptr(SX), uintptr(SY), uintptr(EX), uintptr(EY))
}

func (m *TCanvas) AngleArc(X, Y int32, Radius uint32, StartAngle, SweepAngle float32) {
	LCL().SysCallN(510, m.Instance(), uintptr(X), uintptr(Y), uintptr(Radius), uintptr(StartAngle), uintptr(SweepAngle))
}

func (m *TCanvas) BrushCopy(ADestRect *TRect, ABitmap IBitmap, ASourceRect *TRect, ATransparentColor TColor) {
	LCL().SysCallN(514, m.Instance(), uintptr(unsafePointer(ADestRect)), GetObjectUintptr(ABitmap), uintptr(unsafePointer(ASourceRect)), uintptr(ATransparentColor))
}

func (m *TCanvas) Chord(x1, y1, x2, y2, Angle16Deg, Angle16DegLength int32) {
	LCL().SysCallN(518, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(Angle16Deg), uintptr(Angle16DegLength))
}

func (m *TCanvas) Chord1(x1, y1, x2, y2, SX, SY, EX, EY int32) {
	LCL().SysCallN(519, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(SX), uintptr(SY), uintptr(EX), uintptr(EY))
}

func (m *TCanvas) CopyRectForRect(Dest *TRect, SrcCanvas ICanvas, Source *TRect) {
	LCL().SysCallN(522, m.Instance(), uintptr(unsafePointer(Dest)), GetObjectUintptr(SrcCanvas), uintptr(unsafePointer(Source)))
}

func (m *TCanvas) DrawForGraphic(X, Y int32, SrcGraphic IGraphic) {
	LCL().SysCallN(525, m.Instance(), uintptr(X), uintptr(Y), GetObjectUintptr(SrcGraphic))
}

func (m *TCanvas) DrawFocusRect(ARect *TRect) {
	LCL().SysCallN(524, m.Instance(), uintptr(unsafePointer(ARect)))
}

func (m *TCanvas) StretchDrawForRect(DestRect *TRect, SrcGraphic IGraphic) {
	LCL().SysCallN(553, m.Instance(), uintptr(unsafePointer(DestRect)), GetObjectUintptr(SrcGraphic))
}

func (m *TCanvas) FloodFillForColor(X, Y int32, FillColor TColor, FillStyle TFillStyle) {
	LCL().SysCallN(526, m.Instance(), uintptr(X), uintptr(Y), uintptr(FillColor), uintptr(FillStyle))
}

func (m *TCanvas) Frame3d(ARect *TRect, FrameWidth int32, Style TGraphicsBevelCut) {
	var result0 uintptr
	LCL().SysCallN(531, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(FrameWidth), uintptr(Style))
	*ARect = *(*TRect)(getPointer(result0))
}

func (m *TCanvas) Frame3D1(ARect *TRect, TopColor, BottomColor TColor, FrameWidth int32) {
	var result0 uintptr
	LCL().SysCallN(530, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(TopColor), uintptr(BottomColor), uintptr(FrameWidth))
	*ARect = *(*TRect)(getPointer(result0))
}

func (m *TCanvas) Frame(ARect *TRect) {
	LCL().SysCallN(528, m.Instance(), uintptr(unsafePointer(ARect)))
}

func (m *TCanvas) Frame1(X1, Y1, X2, Y2 int32) {
	LCL().SysCallN(529, m.Instance(), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
}

func (m *TCanvas) FrameRect(ARect *TRect) {
	LCL().SysCallN(532, m.Instance(), uintptr(unsafePointer(ARect)))
}

func (m *TCanvas) FrameRect1(X1, Y1, X2, Y2 int32) {
	LCL().SysCallN(533, m.Instance(), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
}

func (m *TCanvas) GradientFill(ARect *TRect, AStart, AStop TColor, ADirection TGradientDirection) {
	LCL().SysCallN(536, m.Instance(), uintptr(unsafePointer(ARect)), uintptr(AStart), uintptr(AStop), uintptr(ADirection))
}

func (m *TCanvas) Pie(EllipseX1, EllipseY1, EllipseX2, EllipseY2, StartX, StartY, EndX, EndY int32) {
	LCL().SysCallN(541, m.Instance(), uintptr(EllipseX1), uintptr(EllipseY1), uintptr(EllipseX2), uintptr(EllipseY2), uintptr(StartX), uintptr(StartY), uintptr(EndX), uintptr(EndY))
}

func (m *TCanvas) Polygon(Points []TPoint, Winding bool) {
	sysCallPoint(543, m.Instance(), Points, PascalBool(Winding))
}

func (m *TCanvas) Polyline(Points []TPoint) {
	sysCallPoint(544, m.Instance(), Points)
}

func (m *TCanvas) RoundRect(X1, Y1, X2, Y2 int32, RX, RY int32) {
	LCL().SysCallN(548, m.Instance(), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2), uintptr(RX), uintptr(RY))
}

func (m *TCanvas) RoundRect1(Rect *TRect, RX, RY int32) {
	LCL().SysCallN(549, m.Instance(), uintptr(unsafePointer(Rect)), uintptr(RX), uintptr(RY))
}

func (m *TCanvas) TextRect(ARect *TRect, X, Y int32, Text string) {
	LCL().SysCallN(555, m.Instance(), uintptr(unsafePointer(ARect)), uintptr(X), uintptr(Y), PascalStr(Text))
}

func (m *TCanvas) TextRect1(ARect *TRect, X, Y int32, Text string, Style *TTextStyle) {
	LCL().SysCallN(556, m.Instance(), uintptr(unsafePointer(ARect)), uintptr(X), uintptr(Y), PascalStr(Text), uintptr(unsafePointer(Style)))
}

func (m *TCanvas) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(551, m.Instance(), m.changePtr)
}

func (m *TCanvas) SetOnChanging(fn TNotifyEvent) {
	if m.changingPtr != 0 {
		RemoveEventElement(m.changingPtr)
	}
	m.changingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(552, m.Instance(), m.changingPtr)
}
