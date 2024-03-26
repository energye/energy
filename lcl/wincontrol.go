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

// IWinControl Parent: IControl
type IWinControl interface {
	IControl
	BorderWidth() TBorderWidth                                                   // property
	SetBorderWidth(AValue TBorderWidth)                                          // property
	BoundsLockCount() int32                                                      // property
	Brush() IBrush                                                               // property
	CachedClientHeight() int32                                                   // property
	CachedClientWidth() int32                                                    // property
	ChildSizing() IControlChildSizing                                            // property
	SetChildSizing(AValue IControlChildSizing)                                   // property
	ControlCount() int32                                                         // property
	Controls(Index int32) IControl                                               // property
	DockClientCount() int32                                                      // property
	DockClients(Index int32) IControl                                            // property
	DockManager() IDockManager                                                   // property
	SetDockManager(AValue IDockManager)                                          // property
	DockSite() bool                                                              // property
	SetDockSite(AValue bool)                                                     // property
	DoubleBuffered() bool                                                        // property
	SetDoubleBuffered(AValue bool)                                               // property
	Handle() HWND                                                                // property
	SetHandle(AValue HWND)                                                       // property
	IsFlipped() bool                                                             // property
	IsResizing() bool                                                            // property
	TabOrder() TTabOrder                                                         // property
	SetTabOrder(AValue TTabOrder)                                                // property
	TabStop() bool                                                               // property
	SetTabStop(AValue bool)                                                      // property
	ParentDoubleBuffered() bool                                                  // property
	SetParentDoubleBuffered(AValue bool)                                         // property
	ParentWindow() HWND                                                          // property
	SetParentWindow(AValue HWND)                                                 // property
	Showing() bool                                                               // property
	UseDockManager() bool                                                        // property
	SetUseDockManager(AValue bool)                                               // property
	SetDesignerDeleting(AValue bool)                                             // property
	IsSpecialSubControl() bool                                                   // property
	VisibleDockClientCount() int32                                               // property
	ControlAtPos(Pos *TPoint, AllowDisabled bool) IControl                       // function
	ControlAtPos1(Pos *TPoint, AllowDisabled, AllowWinControls bool) IControl    // function
	ControlAtPos2(Pos *TPoint, Flags TControlAtPosFlags) IControl                // function
	ContainsControl(Control IControl) bool                                       // function
	ClientRectNeedsInterfaceUpdate() bool                                        // function
	CanFocus() bool                                                              // function
	CanSetFocus() bool                                                           // function
	GetControlIndex(AControl IControl) int32                                     // function
	Focused() bool                                                               // function
	PerformTab(ForwardTab bool) bool                                             // function
	FindChildControl(ControlName string) IControl                                // function
	GetEnumeratorControls() IWinControlEnumerator                                // function
	GetEnumeratorControlsReverse() IWinControlEnumerator                         // function
	GetDockCaption(AControl IControl) string                                     // function
	HandleAllocated() bool                                                       // function
	BrushCreated() bool                                                          // function
	IntfUTF8KeyPress(UTF8Key *TUTF8Char, RepeatCount int32, SystemKey bool) bool // function
	IntfGetDropFilesTarget() IWinControl                                         // function
	BeginUpdateBounds()                                                          // procedure
	EndUpdateBounds()                                                            // procedure
	LockRealizeBounds()                                                          // procedure
	UnlockRealizeBounds()                                                        // procedure
	DoAdjustClientRectChange(InvalidateRect bool)                                // procedure
	InvalidateClientRectCache(WithChildControls bool)                            // procedure
	DisableAlign()                                                               // procedure
	EnableAlign()                                                                // procedure
	ReAlign()                                                                    // procedure
	ScrollByWS(DeltaX, DeltaY int32)                                             // procedure
	ScrollBy(DeltaX, DeltaY int32)                                               // procedure
	FixDesignFontsPPIWithChildren(ADesignTimePPI int32)                          // procedure
	DockDrop(DragDockObject IDragDockObject, X, Y int32)                         // procedure
	SetControlIndex(AControl IControl, NewIndex int32)                           // procedure
	SelectNext(CurControl IWinControl, GoForward, CheckTabStop bool)             // procedure
	NotifyControls(Msg Word)                                                     // procedure
	AddControl()                                                                 // procedure
	InsertControl(AControl IControl)                                             // procedure
	InsertControl1(AControl IControl, Index int32)                               // procedure
	RemoveControl(AControl IControl)                                             // procedure
	SetFocus()                                                                   // procedure
	FlipChildren(AllLevels bool)                                                 // procedure
	ScaleBy(Multiplier, Divider int32)                                           // procedure
	UpdateDockCaption(Exclude IControl)                                          // procedure
	HandleNeeded()                                                               // procedure
	EraseBackground(DC HDC)                                                      // procedure
	PaintTo(DC HDC, X, Y int32)                                                  // procedure
	PaintTo1(ACanvas ICanvas, X, Y int32)                                        // procedure
	SetShape(AShape IBitmap)                                                     // procedure
	SetShape1(AShape IRegion)                                                    // procedure
	SetOnAlignInsertBefore(fn TAlignInsertBeforeEvent)                           // property event
	SetOnAlignPosition(fn TAlignPositionEvent)                                   // property event
	SetOnDockDrop(fn TDockDropEvent)                                             // property event
	SetOnDockOver(fn TDockOverEvent)                                             // property event
	SetOnEnter(fn TNotifyEvent)                                                  // property event
	SetOnExit(fn TNotifyEvent)                                                   // property event
	SetOnKeyDown(fn TKeyEvent)                                                   // property event
	SetOnKeyPress(fn TKeyPressEvent)                                             // property event
	SetOnKeyUp(fn TKeyEvent)                                                     // property event
	SetOnUnDock(fn TUnDockEvent)                                                 // property event
	SetOnUTF8KeyPress(fn TUTF8KeyPressEvent)                                     // property event
}

// TWinControl Parent: TControl
type TWinControl struct {
	TControl
	alignInsertBeforePtr uintptr
	alignPositionPtr     uintptr
	dockDropPtr          uintptr
	dockOverPtr          uintptr
	enterPtr             uintptr
	exitPtr              uintptr
	keyDownPtr           uintptr
	keyPressPtr          uintptr
	keyUpPtr             uintptr
	unDockPtr            uintptr
	uTF8KeyPressPtr      uintptr
}

func NewWinControl(TheOwner IComponent) IWinControl {
	r1 := LCL().SysCallN(5220, GetObjectUintptr(TheOwner))
	return AsWinControl(r1)
}

func NewWinControlParented(AParentWindow HWND) IWinControl {
	r1 := LCL().SysCallN(5221, uintptr(AParentWindow))
	return AsWinControl(r1)
}

func (m *TWinControl) BorderWidth() TBorderWidth {
	r1 := LCL().SysCallN(5203, 0, m.Instance(), 0)
	return TBorderWidth(r1)
}

func (m *TWinControl) SetBorderWidth(AValue TBorderWidth) {
	LCL().SysCallN(5203, 1, m.Instance(), uintptr(AValue))
}

func (m *TWinControl) BoundsLockCount() int32 {
	r1 := LCL().SysCallN(5204, m.Instance())
	return int32(r1)
}

func (m *TWinControl) Brush() IBrush {
	r1 := LCL().SysCallN(5205, m.Instance())
	return AsBrush(r1)
}

func (m *TWinControl) CachedClientHeight() int32 {
	r1 := LCL().SysCallN(5207, m.Instance())
	return int32(r1)
}

func (m *TWinControl) CachedClientWidth() int32 {
	r1 := LCL().SysCallN(5208, m.Instance())
	return int32(r1)
}

func (m *TWinControl) ChildSizing() IControlChildSizing {
	r1 := LCL().SysCallN(5211, 0, m.Instance(), 0)
	return AsControlChildSizing(r1)
}

func (m *TWinControl) SetChildSizing(AValue IControlChildSizing) {
	LCL().SysCallN(5211, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TWinControl) ControlCount() int32 {
	r1 := LCL().SysCallN(5218, m.Instance())
	return int32(r1)
}

func (m *TWinControl) Controls(Index int32) IControl {
	r1 := LCL().SysCallN(5219, m.Instance(), uintptr(Index))
	return AsControl(r1)
}

func (m *TWinControl) DockClientCount() int32 {
	r1 := LCL().SysCallN(5225, m.Instance())
	return int32(r1)
}

func (m *TWinControl) DockClients(Index int32) IControl {
	r1 := LCL().SysCallN(5226, m.Instance(), uintptr(Index))
	return AsControl(r1)
}

func (m *TWinControl) DockManager() IDockManager {
	r1 := LCL().SysCallN(5228, 0, m.Instance(), 0)
	return AsDockManager(r1)
}

func (m *TWinControl) SetDockManager(AValue IDockManager) {
	LCL().SysCallN(5228, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TWinControl) DockSite() bool {
	r1 := LCL().SysCallN(5229, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWinControl) SetDockSite(AValue bool) {
	LCL().SysCallN(5229, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) DoubleBuffered() bool {
	r1 := LCL().SysCallN(5230, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWinControl) SetDoubleBuffered(AValue bool) {
	LCL().SysCallN(5230, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) Handle() HWND {
	r1 := LCL().SysCallN(5242, 0, m.Instance(), 0)
	return HWND(r1)
}

func (m *TWinControl) SetHandle(AValue HWND) {
	LCL().SysCallN(5242, 1, m.Instance(), uintptr(AValue))
}

func (m *TWinControl) IsFlipped() bool {
	r1 := LCL().SysCallN(5250, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) IsResizing() bool {
	r1 := LCL().SysCallN(5251, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) TabOrder() TTabOrder {
	r1 := LCL().SysCallN(5282, 0, m.Instance(), 0)
	return TTabOrder(r1)
}

func (m *TWinControl) SetTabOrder(AValue TTabOrder) {
	LCL().SysCallN(5282, 1, m.Instance(), uintptr(AValue))
}

func (m *TWinControl) TabStop() bool {
	r1 := LCL().SysCallN(5283, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWinControl) SetTabStop(AValue bool) {
	LCL().SysCallN(5283, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) ParentDoubleBuffered() bool {
	r1 := LCL().SysCallN(5257, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWinControl) SetParentDoubleBuffered(AValue bool) {
	LCL().SysCallN(5257, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) ParentWindow() HWND {
	r1 := LCL().SysCallN(5258, 0, m.Instance(), 0)
	return HWND(r1)
}

func (m *TWinControl) SetParentWindow(AValue HWND) {
	LCL().SysCallN(5258, 1, m.Instance(), uintptr(AValue))
}

func (m *TWinControl) Showing() bool {
	r1 := LCL().SysCallN(5281, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) UseDockManager() bool {
	r1 := LCL().SysCallN(5286, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWinControl) SetUseDockManager(AValue bool) {
	LCL().SysCallN(5286, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) SetDesignerDeleting(AValue bool) {
	LCL().SysCallN(5222, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) IsSpecialSubControl() bool {
	r1 := LCL().SysCallN(5252, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) VisibleDockClientCount() int32 {
	r1 := LCL().SysCallN(5287, m.Instance())
	return int32(r1)
}

func (m *TWinControl) ControlAtPos(Pos *TPoint, AllowDisabled bool) IControl {
	r1 := LCL().SysCallN(5215, m.Instance(), uintptr(unsafe.Pointer(Pos)), PascalBool(AllowDisabled))
	return AsControl(r1)
}

func (m *TWinControl) ControlAtPos1(Pos *TPoint, AllowDisabled, AllowWinControls bool) IControl {
	r1 := LCL().SysCallN(5216, m.Instance(), uintptr(unsafe.Pointer(Pos)), PascalBool(AllowDisabled), PascalBool(AllowWinControls))
	return AsControl(r1)
}

func (m *TWinControl) ControlAtPos2(Pos *TPoint, Flags TControlAtPosFlags) IControl {
	r1 := LCL().SysCallN(5217, m.Instance(), uintptr(unsafe.Pointer(Pos)), uintptr(Flags))
	return AsControl(r1)
}

func (m *TWinControl) ContainsControl(Control IControl) bool {
	r1 := LCL().SysCallN(5214, m.Instance(), GetObjectUintptr(Control))
	return GoBool(r1)
}

func (m *TWinControl) ClientRectNeedsInterfaceUpdate() bool {
	r1 := LCL().SysCallN(5213, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) CanFocus() bool {
	r1 := LCL().SysCallN(5209, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) CanSetFocus() bool {
	r1 := LCL().SysCallN(5210, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) GetControlIndex(AControl IControl) int32 {
	r1 := LCL().SysCallN(5238, m.Instance(), GetObjectUintptr(AControl))
	return int32(r1)
}

func (m *TWinControl) Focused() bool {
	r1 := LCL().SysCallN(5237, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) PerformTab(ForwardTab bool) bool {
	r1 := LCL().SysCallN(5259, m.Instance(), PascalBool(ForwardTab))
	return GoBool(r1)
}

func (m *TWinControl) FindChildControl(ControlName string) IControl {
	r1 := LCL().SysCallN(5234, m.Instance(), PascalStr(ControlName))
	return AsControl(r1)
}

func (m *TWinControl) GetEnumeratorControls() IWinControlEnumerator {
	r1 := LCL().SysCallN(5240, m.Instance())
	return AsWinControlEnumerator(r1)
}

func (m *TWinControl) GetEnumeratorControlsReverse() IWinControlEnumerator {
	r1 := LCL().SysCallN(5241, m.Instance())
	return AsWinControlEnumerator(r1)
}

func (m *TWinControl) GetDockCaption(AControl IControl) string {
	r1 := LCL().SysCallN(5239, m.Instance(), GetObjectUintptr(AControl))
	return GoStr(r1)
}

func (m *TWinControl) HandleAllocated() bool {
	r1 := LCL().SysCallN(5243, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) BrushCreated() bool {
	r1 := LCL().SysCallN(5206, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) IntfUTF8KeyPress(UTF8Key *TUTF8Char, RepeatCount int32, SystemKey bool) bool {
	var result0 uintptr
	r1 := LCL().SysCallN(5248, m.Instance(), uintptr(unsafe.Pointer(&result0)), uintptr(RepeatCount), PascalBool(SystemKey))
	*UTF8Key = *(*TUTF8Char)(getPointer(result0))
	return GoBool(r1)
}

func (m *TWinControl) IntfGetDropFilesTarget() IWinControl {
	r1 := LCL().SysCallN(5247, m.Instance())
	return AsWinControl(r1)
}

func WinControlClass() TClass {
	ret := LCL().SysCallN(5212)
	return TClass(ret)
}

func (m *TWinControl) BeginUpdateBounds() {
	LCL().SysCallN(5202, m.Instance())
}

func (m *TWinControl) EndUpdateBounds() {
	LCL().SysCallN(5232, m.Instance())
}

func (m *TWinControl) LockRealizeBounds() {
	LCL().SysCallN(5253, m.Instance())
}

func (m *TWinControl) UnlockRealizeBounds() {
	LCL().SysCallN(5284, m.Instance())
}

func (m *TWinControl) DoAdjustClientRectChange(InvalidateRect bool) {
	LCL().SysCallN(5224, m.Instance(), PascalBool(InvalidateRect))
}

func (m *TWinControl) InvalidateClientRectCache(WithChildControls bool) {
	LCL().SysCallN(5249, m.Instance(), PascalBool(WithChildControls))
}

func (m *TWinControl) DisableAlign() {
	LCL().SysCallN(5223, m.Instance())
}

func (m *TWinControl) EnableAlign() {
	LCL().SysCallN(5231, m.Instance())
}

func (m *TWinControl) ReAlign() {
	LCL().SysCallN(5260, m.Instance())
}

func (m *TWinControl) ScrollByWS(DeltaX, DeltaY int32) {
	LCL().SysCallN(5264, m.Instance(), uintptr(DeltaX), uintptr(DeltaY))
}

func (m *TWinControl) ScrollBy(DeltaX, DeltaY int32) {
	LCL().SysCallN(5263, m.Instance(), uintptr(DeltaX), uintptr(DeltaY))
}

func (m *TWinControl) FixDesignFontsPPIWithChildren(ADesignTimePPI int32) {
	LCL().SysCallN(5235, m.Instance(), uintptr(ADesignTimePPI))
}

func (m *TWinControl) DockDrop(DragDockObject IDragDockObject, X, Y int32) {
	LCL().SysCallN(5227, m.Instance(), GetObjectUintptr(DragDockObject), uintptr(X), uintptr(Y))
}

func (m *TWinControl) SetControlIndex(AControl IControl, NewIndex int32) {
	LCL().SysCallN(5266, m.Instance(), GetObjectUintptr(AControl), uintptr(NewIndex))
}

func (m *TWinControl) SelectNext(CurControl IWinControl, GoForward, CheckTabStop bool) {
	LCL().SysCallN(5265, m.Instance(), GetObjectUintptr(CurControl), PascalBool(GoForward), PascalBool(CheckTabStop))
}

func (m *TWinControl) NotifyControls(Msg Word) {
	LCL().SysCallN(5254, m.Instance(), uintptr(Msg))
}

func (m *TWinControl) AddControl() {
	LCL().SysCallN(5201, m.Instance())
}

func (m *TWinControl) InsertControl(AControl IControl) {
	LCL().SysCallN(5245, m.Instance(), GetObjectUintptr(AControl))
}

func (m *TWinControl) InsertControl1(AControl IControl, Index int32) {
	LCL().SysCallN(5246, m.Instance(), GetObjectUintptr(AControl), uintptr(Index))
}

func (m *TWinControl) RemoveControl(AControl IControl) {
	LCL().SysCallN(5261, m.Instance(), GetObjectUintptr(AControl))
}

func (m *TWinControl) SetFocus() {
	LCL().SysCallN(5267, m.Instance())
}

func (m *TWinControl) FlipChildren(AllLevels bool) {
	LCL().SysCallN(5236, m.Instance(), PascalBool(AllLevels))
}

func (m *TWinControl) ScaleBy(Multiplier, Divider int32) {
	LCL().SysCallN(5262, m.Instance(), uintptr(Multiplier), uintptr(Divider))
}

func (m *TWinControl) UpdateDockCaption(Exclude IControl) {
	LCL().SysCallN(5285, m.Instance(), GetObjectUintptr(Exclude))
}

func (m *TWinControl) HandleNeeded() {
	LCL().SysCallN(5244, m.Instance())
}

func (m *TWinControl) EraseBackground(DC HDC) {
	LCL().SysCallN(5233, m.Instance(), uintptr(DC))
}

func (m *TWinControl) PaintTo(DC HDC, X, Y int32) {
	LCL().SysCallN(5255, m.Instance(), uintptr(DC), uintptr(X), uintptr(Y))
}

func (m *TWinControl) PaintTo1(ACanvas ICanvas, X, Y int32) {
	LCL().SysCallN(5256, m.Instance(), GetObjectUintptr(ACanvas), uintptr(X), uintptr(Y))
}

func (m *TWinControl) SetShape(AShape IBitmap) {
	LCL().SysCallN(5279, m.Instance(), GetObjectUintptr(AShape))
}

func (m *TWinControl) SetShape1(AShape IRegion) {
	LCL().SysCallN(5280, m.Instance(), GetObjectUintptr(AShape))
}

func (m *TWinControl) SetOnAlignInsertBefore(fn TAlignInsertBeforeEvent) {
	if m.alignInsertBeforePtr != 0 {
		RemoveEventElement(m.alignInsertBeforePtr)
	}
	m.alignInsertBeforePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5268, m.Instance(), m.alignInsertBeforePtr)
}

func (m *TWinControl) SetOnAlignPosition(fn TAlignPositionEvent) {
	if m.alignPositionPtr != 0 {
		RemoveEventElement(m.alignPositionPtr)
	}
	m.alignPositionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5269, m.Instance(), m.alignPositionPtr)
}

func (m *TWinControl) SetOnDockDrop(fn TDockDropEvent) {
	if m.dockDropPtr != 0 {
		RemoveEventElement(m.dockDropPtr)
	}
	m.dockDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5270, m.Instance(), m.dockDropPtr)
}

func (m *TWinControl) SetOnDockOver(fn TDockOverEvent) {
	if m.dockOverPtr != 0 {
		RemoveEventElement(m.dockOverPtr)
	}
	m.dockOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5271, m.Instance(), m.dockOverPtr)
}

func (m *TWinControl) SetOnEnter(fn TNotifyEvent) {
	if m.enterPtr != 0 {
		RemoveEventElement(m.enterPtr)
	}
	m.enterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5272, m.Instance(), m.enterPtr)
}

func (m *TWinControl) SetOnExit(fn TNotifyEvent) {
	if m.exitPtr != 0 {
		RemoveEventElement(m.exitPtr)
	}
	m.exitPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5273, m.Instance(), m.exitPtr)
}

func (m *TWinControl) SetOnKeyDown(fn TKeyEvent) {
	if m.keyDownPtr != 0 {
		RemoveEventElement(m.keyDownPtr)
	}
	m.keyDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5274, m.Instance(), m.keyDownPtr)
}

func (m *TWinControl) SetOnKeyPress(fn TKeyPressEvent) {
	if m.keyPressPtr != 0 {
		RemoveEventElement(m.keyPressPtr)
	}
	m.keyPressPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5275, m.Instance(), m.keyPressPtr)
}

func (m *TWinControl) SetOnKeyUp(fn TKeyEvent) {
	if m.keyUpPtr != 0 {
		RemoveEventElement(m.keyUpPtr)
	}
	m.keyUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5276, m.Instance(), m.keyUpPtr)
}

func (m *TWinControl) SetOnUnDock(fn TUnDockEvent) {
	if m.unDockPtr != 0 {
		RemoveEventElement(m.unDockPtr)
	}
	m.unDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5278, m.Instance(), m.unDockPtr)
}

func (m *TWinControl) SetOnUTF8KeyPress(fn TUTF8KeyPressEvent) {
	if m.uTF8KeyPressPtr != 0 {
		RemoveEventElement(m.uTF8KeyPressPtr)
	}
	m.uTF8KeyPressPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5277, m.Instance(), m.uTF8KeyPressPtr)
}
