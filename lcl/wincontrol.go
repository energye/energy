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
	r1 := LCL().SysCallN(6029, GetObjectUintptr(TheOwner))
	return AsWinControl(r1)
}

func NewWinControlParented(AParentWindow HWND) IWinControl {
	r1 := LCL().SysCallN(6030, uintptr(AParentWindow))
	return AsWinControl(r1)
}

func (m *TWinControl) BorderWidth() TBorderWidth {
	r1 := LCL().SysCallN(6012, 0, m.Instance(), 0)
	return TBorderWidth(r1)
}

func (m *TWinControl) SetBorderWidth(AValue TBorderWidth) {
	LCL().SysCallN(6012, 1, m.Instance(), uintptr(AValue))
}

func (m *TWinControl) BoundsLockCount() int32 {
	r1 := LCL().SysCallN(6013, m.Instance())
	return int32(r1)
}

func (m *TWinControl) Brush() IBrush {
	r1 := LCL().SysCallN(6014, m.Instance())
	return AsBrush(r1)
}

func (m *TWinControl) CachedClientHeight() int32 {
	r1 := LCL().SysCallN(6016, m.Instance())
	return int32(r1)
}

func (m *TWinControl) CachedClientWidth() int32 {
	r1 := LCL().SysCallN(6017, m.Instance())
	return int32(r1)
}

func (m *TWinControl) ChildSizing() IControlChildSizing {
	r1 := LCL().SysCallN(6020, 0, m.Instance(), 0)
	return AsControlChildSizing(r1)
}

func (m *TWinControl) SetChildSizing(AValue IControlChildSizing) {
	LCL().SysCallN(6020, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TWinControl) ControlCount() int32 {
	r1 := LCL().SysCallN(6027, m.Instance())
	return int32(r1)
}

func (m *TWinControl) Controls(Index int32) IControl {
	r1 := LCL().SysCallN(6028, m.Instance(), uintptr(Index))
	return AsControl(r1)
}

func (m *TWinControl) DockClientCount() int32 {
	r1 := LCL().SysCallN(6034, m.Instance())
	return int32(r1)
}

func (m *TWinControl) DockClients(Index int32) IControl {
	r1 := LCL().SysCallN(6035, m.Instance(), uintptr(Index))
	return AsControl(r1)
}

func (m *TWinControl) DockManager() IDockManager {
	r1 := LCL().SysCallN(6037, 0, m.Instance(), 0)
	return AsDockManager(r1)
}

func (m *TWinControl) SetDockManager(AValue IDockManager) {
	LCL().SysCallN(6037, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TWinControl) DockSite() bool {
	r1 := LCL().SysCallN(6038, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWinControl) SetDockSite(AValue bool) {
	LCL().SysCallN(6038, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) DoubleBuffered() bool {
	r1 := LCL().SysCallN(6039, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWinControl) SetDoubleBuffered(AValue bool) {
	LCL().SysCallN(6039, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) Handle() HWND {
	r1 := LCL().SysCallN(6051, 0, m.Instance(), 0)
	return HWND(r1)
}

func (m *TWinControl) SetHandle(AValue HWND) {
	LCL().SysCallN(6051, 1, m.Instance(), uintptr(AValue))
}

func (m *TWinControl) IsFlipped() bool {
	r1 := LCL().SysCallN(6059, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) IsResizing() bool {
	r1 := LCL().SysCallN(6060, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) TabOrder() TTabOrder {
	r1 := LCL().SysCallN(6091, 0, m.Instance(), 0)
	return TTabOrder(r1)
}

func (m *TWinControl) SetTabOrder(AValue TTabOrder) {
	LCL().SysCallN(6091, 1, m.Instance(), uintptr(AValue))
}

func (m *TWinControl) TabStop() bool {
	r1 := LCL().SysCallN(6092, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWinControl) SetTabStop(AValue bool) {
	LCL().SysCallN(6092, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) ParentDoubleBuffered() bool {
	r1 := LCL().SysCallN(6066, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWinControl) SetParentDoubleBuffered(AValue bool) {
	LCL().SysCallN(6066, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) ParentWindow() HWND {
	r1 := LCL().SysCallN(6067, 0, m.Instance(), 0)
	return HWND(r1)
}

func (m *TWinControl) SetParentWindow(AValue HWND) {
	LCL().SysCallN(6067, 1, m.Instance(), uintptr(AValue))
}

func (m *TWinControl) Showing() bool {
	r1 := LCL().SysCallN(6090, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) UseDockManager() bool {
	r1 := LCL().SysCallN(6095, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWinControl) SetUseDockManager(AValue bool) {
	LCL().SysCallN(6095, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) SetDesignerDeleting(AValue bool) {
	LCL().SysCallN(6031, m.Instance(), PascalBool(AValue))
}

func (m *TWinControl) IsSpecialSubControl() bool {
	r1 := LCL().SysCallN(6061, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) VisibleDockClientCount() int32 {
	r1 := LCL().SysCallN(6096, m.Instance())
	return int32(r1)
}

func (m *TWinControl) ControlAtPos(Pos *TPoint, AllowDisabled bool) IControl {
	r1 := LCL().SysCallN(6024, m.Instance(), uintptr(unsafePointer(Pos)), PascalBool(AllowDisabled))
	return AsControl(r1)
}

func (m *TWinControl) ControlAtPos1(Pos *TPoint, AllowDisabled, AllowWinControls bool) IControl {
	r1 := LCL().SysCallN(6025, m.Instance(), uintptr(unsafePointer(Pos)), PascalBool(AllowDisabled), PascalBool(AllowWinControls))
	return AsControl(r1)
}

func (m *TWinControl) ControlAtPos2(Pos *TPoint, Flags TControlAtPosFlags) IControl {
	r1 := LCL().SysCallN(6026, m.Instance(), uintptr(unsafePointer(Pos)), uintptr(Flags))
	return AsControl(r1)
}

func (m *TWinControl) ContainsControl(Control IControl) bool {
	r1 := LCL().SysCallN(6023, m.Instance(), GetObjectUintptr(Control))
	return GoBool(r1)
}

func (m *TWinControl) ClientRectNeedsInterfaceUpdate() bool {
	r1 := LCL().SysCallN(6022, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) CanFocus() bool {
	r1 := LCL().SysCallN(6018, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) CanSetFocus() bool {
	r1 := LCL().SysCallN(6019, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) GetControlIndex(AControl IControl) int32 {
	r1 := LCL().SysCallN(6047, m.Instance(), GetObjectUintptr(AControl))
	return int32(r1)
}

func (m *TWinControl) Focused() bool {
	r1 := LCL().SysCallN(6046, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) PerformTab(ForwardTab bool) bool {
	r1 := LCL().SysCallN(6068, m.Instance(), PascalBool(ForwardTab))
	return GoBool(r1)
}

func (m *TWinControl) FindChildControl(ControlName string) IControl {
	r1 := LCL().SysCallN(6043, m.Instance(), PascalStr(ControlName))
	return AsControl(r1)
}

func (m *TWinControl) GetEnumeratorControls() IWinControlEnumerator {
	r1 := LCL().SysCallN(6049, m.Instance())
	return AsWinControlEnumerator(r1)
}

func (m *TWinControl) GetEnumeratorControlsReverse() IWinControlEnumerator {
	r1 := LCL().SysCallN(6050, m.Instance())
	return AsWinControlEnumerator(r1)
}

func (m *TWinControl) GetDockCaption(AControl IControl) string {
	r1 := LCL().SysCallN(6048, m.Instance(), GetObjectUintptr(AControl))
	return GoStr(r1)
}

func (m *TWinControl) HandleAllocated() bool {
	r1 := LCL().SysCallN(6052, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) BrushCreated() bool {
	r1 := LCL().SysCallN(6015, m.Instance())
	return GoBool(r1)
}

func (m *TWinControl) IntfUTF8KeyPress(UTF8Key *TUTF8Char, RepeatCount int32, SystemKey bool) bool {
	var result0 uintptr
	r1 := LCL().SysCallN(6057, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(RepeatCount), PascalBool(SystemKey))
	*UTF8Key = *(*TUTF8Char)(getPointer(result0))
	return GoBool(r1)
}

func (m *TWinControl) IntfGetDropFilesTarget() IWinControl {
	r1 := LCL().SysCallN(6056, m.Instance())
	return AsWinControl(r1)
}

func WinControlClass() TClass {
	ret := LCL().SysCallN(6021)
	return TClass(ret)
}

func (m *TWinControl) BeginUpdateBounds() {
	LCL().SysCallN(6011, m.Instance())
}

func (m *TWinControl) EndUpdateBounds() {
	LCL().SysCallN(6041, m.Instance())
}

func (m *TWinControl) LockRealizeBounds() {
	LCL().SysCallN(6062, m.Instance())
}

func (m *TWinControl) UnlockRealizeBounds() {
	LCL().SysCallN(6093, m.Instance())
}

func (m *TWinControl) DoAdjustClientRectChange(InvalidateRect bool) {
	LCL().SysCallN(6033, m.Instance(), PascalBool(InvalidateRect))
}

func (m *TWinControl) InvalidateClientRectCache(WithChildControls bool) {
	LCL().SysCallN(6058, m.Instance(), PascalBool(WithChildControls))
}

func (m *TWinControl) DisableAlign() {
	LCL().SysCallN(6032, m.Instance())
}

func (m *TWinControl) EnableAlign() {
	LCL().SysCallN(6040, m.Instance())
}

func (m *TWinControl) ReAlign() {
	LCL().SysCallN(6069, m.Instance())
}

func (m *TWinControl) ScrollByWS(DeltaX, DeltaY int32) {
	LCL().SysCallN(6073, m.Instance(), uintptr(DeltaX), uintptr(DeltaY))
}

func (m *TWinControl) ScrollBy(DeltaX, DeltaY int32) {
	LCL().SysCallN(6072, m.Instance(), uintptr(DeltaX), uintptr(DeltaY))
}

func (m *TWinControl) FixDesignFontsPPIWithChildren(ADesignTimePPI int32) {
	LCL().SysCallN(6044, m.Instance(), uintptr(ADesignTimePPI))
}

func (m *TWinControl) DockDrop(DragDockObject IDragDockObject, X, Y int32) {
	LCL().SysCallN(6036, m.Instance(), GetObjectUintptr(DragDockObject), uintptr(X), uintptr(Y))
}

func (m *TWinControl) SetControlIndex(AControl IControl, NewIndex int32) {
	LCL().SysCallN(6075, m.Instance(), GetObjectUintptr(AControl), uintptr(NewIndex))
}

func (m *TWinControl) SelectNext(CurControl IWinControl, GoForward, CheckTabStop bool) {
	LCL().SysCallN(6074, m.Instance(), GetObjectUintptr(CurControl), PascalBool(GoForward), PascalBool(CheckTabStop))
}

func (m *TWinControl) NotifyControls(Msg Word) {
	LCL().SysCallN(6063, m.Instance(), uintptr(Msg))
}

func (m *TWinControl) AddControl() {
	LCL().SysCallN(6010, m.Instance())
}

func (m *TWinControl) InsertControl(AControl IControl) {
	LCL().SysCallN(6054, m.Instance(), GetObjectUintptr(AControl))
}

func (m *TWinControl) InsertControl1(AControl IControl, Index int32) {
	LCL().SysCallN(6055, m.Instance(), GetObjectUintptr(AControl), uintptr(Index))
}

func (m *TWinControl) RemoveControl(AControl IControl) {
	LCL().SysCallN(6070, m.Instance(), GetObjectUintptr(AControl))
}

func (m *TWinControl) SetFocus() {
	LCL().SysCallN(6076, m.Instance())
}

func (m *TWinControl) FlipChildren(AllLevels bool) {
	LCL().SysCallN(6045, m.Instance(), PascalBool(AllLevels))
}

func (m *TWinControl) ScaleBy(Multiplier, Divider int32) {
	LCL().SysCallN(6071, m.Instance(), uintptr(Multiplier), uintptr(Divider))
}

func (m *TWinControl) UpdateDockCaption(Exclude IControl) {
	LCL().SysCallN(6094, m.Instance(), GetObjectUintptr(Exclude))
}

func (m *TWinControl) HandleNeeded() {
	LCL().SysCallN(6053, m.Instance())
}

func (m *TWinControl) EraseBackground(DC HDC) {
	LCL().SysCallN(6042, m.Instance(), uintptr(DC))
}

func (m *TWinControl) PaintTo(DC HDC, X, Y int32) {
	LCL().SysCallN(6064, m.Instance(), uintptr(DC), uintptr(X), uintptr(Y))
}

func (m *TWinControl) PaintTo1(ACanvas ICanvas, X, Y int32) {
	LCL().SysCallN(6065, m.Instance(), GetObjectUintptr(ACanvas), uintptr(X), uintptr(Y))
}

func (m *TWinControl) SetShape(AShape IBitmap) {
	LCL().SysCallN(6088, m.Instance(), GetObjectUintptr(AShape))
}

func (m *TWinControl) SetShape1(AShape IRegion) {
	LCL().SysCallN(6089, m.Instance(), GetObjectUintptr(AShape))
}

func (m *TWinControl) SetOnAlignInsertBefore(fn TAlignInsertBeforeEvent) {
	if m.alignInsertBeforePtr != 0 {
		RemoveEventElement(m.alignInsertBeforePtr)
	}
	m.alignInsertBeforePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6077, m.Instance(), m.alignInsertBeforePtr)
}

func (m *TWinControl) SetOnAlignPosition(fn TAlignPositionEvent) {
	if m.alignPositionPtr != 0 {
		RemoveEventElement(m.alignPositionPtr)
	}
	m.alignPositionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6078, m.Instance(), m.alignPositionPtr)
}

func (m *TWinControl) SetOnDockDrop(fn TDockDropEvent) {
	if m.dockDropPtr != 0 {
		RemoveEventElement(m.dockDropPtr)
	}
	m.dockDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6079, m.Instance(), m.dockDropPtr)
}

func (m *TWinControl) SetOnDockOver(fn TDockOverEvent) {
	if m.dockOverPtr != 0 {
		RemoveEventElement(m.dockOverPtr)
	}
	m.dockOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6080, m.Instance(), m.dockOverPtr)
}

func (m *TWinControl) SetOnEnter(fn TNotifyEvent) {
	if m.enterPtr != 0 {
		RemoveEventElement(m.enterPtr)
	}
	m.enterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6081, m.Instance(), m.enterPtr)
}

func (m *TWinControl) SetOnExit(fn TNotifyEvent) {
	if m.exitPtr != 0 {
		RemoveEventElement(m.exitPtr)
	}
	m.exitPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6082, m.Instance(), m.exitPtr)
}

func (m *TWinControl) SetOnKeyDown(fn TKeyEvent) {
	if m.keyDownPtr != 0 {
		RemoveEventElement(m.keyDownPtr)
	}
	m.keyDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6083, m.Instance(), m.keyDownPtr)
}

func (m *TWinControl) SetOnKeyPress(fn TKeyPressEvent) {
	if m.keyPressPtr != 0 {
		RemoveEventElement(m.keyPressPtr)
	}
	m.keyPressPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6084, m.Instance(), m.keyPressPtr)
}

func (m *TWinControl) SetOnKeyUp(fn TKeyEvent) {
	if m.keyUpPtr != 0 {
		RemoveEventElement(m.keyUpPtr)
	}
	m.keyUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6085, m.Instance(), m.keyUpPtr)
}

func (m *TWinControl) SetOnUnDock(fn TUnDockEvent) {
	if m.unDockPtr != 0 {
		RemoveEventElement(m.unDockPtr)
	}
	m.unDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6087, m.Instance(), m.unDockPtr)
}

func (m *TWinControl) SetOnUTF8KeyPress(fn TUTF8KeyPressEvent) {
	if m.uTF8KeyPressPtr != 0 {
		RemoveEventElement(m.uTF8KeyPressPtr)
	}
	m.uTF8KeyPressPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6086, m.Instance(), m.uTF8KeyPressPtr)
}
