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

// IControl Parent: ILCLComponent
type IControl interface {
	ILCLComponent
	AnchoredControls(Index int32) IControl                                                                         // property
	BaseBounds() (resultRect TRect)                                                                                // property
	ReadBounds() (resultRect TRect)                                                                                // property
	BaseParentClientSize() (resultSize TSize)                                                                      // property
	AccessibleName() string                                                                                        // property
	SetAccessibleName(AValue string)                                                                               // property
	AccessibleDescription() string                                                                                 // property
	SetAccessibleDescription(AValue string)                                                                        // property
	AccessibleValue() string                                                                                       // property
	SetAccessibleValue(AValue string)                                                                              // property
	AccessibleRole() TLazAccessibilityRole                                                                         // property
	SetAccessibleRole(AValue TLazAccessibilityRole)                                                                // property
	Action() IBasicAction                                                                                          // property
	SetAction(AValue IBasicAction)                                                                                 // property
	Align() TAlign                                                                                                 // property
	SetAlign(AValue TAlign)                                                                                        // property
	Anchors() TAnchors                                                                                             // property
	SetAnchors(AValue TAnchors)                                                                                    // property
	AnchorSide(Kind TAnchorKind) IAnchorSide                                                                       // property
	AutoSize() bool                                                                                                // property
	SetAutoSize(AValue bool)                                                                                       // property
	BorderSpacing() IControlBorderSpacing                                                                          // property
	SetBorderSpacing(AValue IControlBorderSpacing)                                                                 // property
	BoundsRect() (resultRect TRect)                                                                                // property
	SetBoundsRect(AValue *TRect)                                                                                   // property
	BoundsRectForNewParent() (resultRect TRect)                                                                    // property
	SetBoundsRectForNewParent(AValue *TRect)                                                                       // property
	Caption() string                                                                                               // property
	SetCaption(AValue string)                                                                                      // property
	CaptureMouseButtons() TCaptureMouseButtons                                                                     // property
	SetCaptureMouseButtons(AValue TCaptureMouseButtons)                                                            // property
	ClientHeight() int32                                                                                           // property
	SetClientHeight(AValue int32)                                                                                  // property
	ClientOrigin() (resultPoint TPoint)                                                                            // property
	ClientRect() (resultRect TRect)                                                                                // property
	ClientWidth() int32                                                                                            // property
	SetClientWidth(AValue int32)                                                                                   // property
	Color() TColor                                                                                                 // property
	SetColor(AValue TColor)                                                                                        // property
	Constraints() ISizeConstraints                                                                                 // property
	SetConstraints(AValue ISizeConstraints)                                                                        // property
	ControlOrigin() (resultPoint TPoint)                                                                           // property
	ControlState() TControlState                                                                                   // property
	SetControlState(AValue TControlState)                                                                          // property
	ControlStyle() TControlStyle                                                                                   // property
	SetControlStyle(AValue TControlStyle)                                                                          // property
	Enabled() bool                                                                                                 // property
	SetEnabled(AValue bool)                                                                                        // property
	Font() IFont                                                                                                   // property
	SetFont(AValue IFont)                                                                                          // property
	IsControl() bool                                                                                               // property
	SetIsControl(AValue bool)                                                                                      // property
	MouseInClient() bool                                                                                           // property
	Parent() IWinControl                                                                                           // property
	SetParent(AValue IWinControl)                                                                                  // property
	PopupMenu() IPopupMenu                                                                                         // property
	SetPopupMenu(AValue IPopupMenu)                                                                                // property
	ShowHint() bool                                                                                                // property
	SetShowHint(AValue bool)                                                                                       // property
	Visible() bool                                                                                                 // property
	SetVisible(AValue bool)                                                                                        // property
	DockOrientation() TDockOrientation                                                                             // property
	SetDockOrientation(AValue TDockOrientation)                                                                    // property
	Floating() bool                                                                                                // property
	FloatingDockSiteClass() TWinControlClass                                                                       // property
	SetFloatingDockSiteClass(AValue TWinControlClass)                                                              // property
	HostDockSite() IWinControl                                                                                     // property
	SetHostDockSite(AValue IWinControl)                                                                            // property
	LRDockWidth() int32                                                                                            // property
	SetLRDockWidth(AValue int32)                                                                                   // property
	TBDockHeight() int32                                                                                           // property
	SetTBDockHeight(AValue int32)                                                                                  // property
	UndockHeight() int32                                                                                           // property
	SetUndockHeight(AValue int32)                                                                                  // property
	UndockWidth() int32                                                                                            // property
	SetUndockWidth(AValue int32)                                                                                   // property
	BiDiMode() TBiDiMode                                                                                           // property
	SetBiDiMode(AValue TBiDiMode)                                                                                  // property
	ParentBiDiMode() bool                                                                                          // property
	SetParentBiDiMode(AValue bool)                                                                                 // property
	AnchorSideLeft() IAnchorSide                                                                                   // property
	SetAnchorSideLeft(AValue IAnchorSide)                                                                          // property
	AnchorSideTop() IAnchorSide                                                                                    // property
	SetAnchorSideTop(AValue IAnchorSide)                                                                           // property
	AnchorSideRight() IAnchorSide                                                                                  // property
	SetAnchorSideRight(AValue IAnchorSide)                                                                         // property
	AnchorSideBottom() IAnchorSide                                                                                 // property
	SetAnchorSideBottom(AValue IAnchorSide)                                                                        // property
	Cursor() TCursor                                                                                               // property
	SetCursor(AValue TCursor)                                                                                      // property
	Left() int32                                                                                                   // property
	SetLeft(AValue int32)                                                                                          // property
	Height() int32                                                                                                 // property
	SetHeight(AValue int32)                                                                                        // property
	Hint() string                                                                                                  // property
	SetHint(AValue string)                                                                                         // property
	Top() int32                                                                                                    // property
	SetTop(AValue int32)                                                                                           // property
	Width() int32                                                                                                  // property
	SetWidth(AValue int32)                                                                                         // property
	HelpType() THelpType                                                                                           // property
	SetHelpType(AValue THelpType)                                                                                  // property
	HelpKeyword() string                                                                                           // property
	SetHelpKeyword(AValue string)                                                                                  // property
	HelpContext() THelpContext                                                                                     // property
	SetHelpContext(AValue THelpContext)                                                                            // property
	ManualDock(NewDockSite IWinControl, DropControl IControl, ControlSide TAlign, KeepDockSiteSize bool) bool      // function
	ManualFloat(TheScreenRect *TRect, KeepDockSiteSize bool) bool                                                  // function
	ReplaceDockedControl(Control IControl, NewDockSite IWinControl, DropControl IControl, ControlSide TAlign) bool // function
	Docked() bool                                                                                                  // function
	Dragging() bool                                                                                                // function
	GetAccessibleObject() ILazAccessibleObject                                                                     // function
	CreateAccessibleObject() ILazAccessibleObject                                                                  // function
	GetSelectedChildAccessibleObject() ILazAccessibleObject                                                        // function
	GetChildAccessibleObjectAtPos(APos *TPoint) ILazAccessibleObject                                               // function
	ScaleDesignToForm(ASize int32) int32                                                                           // function
	ScaleFormToDesign(ASize int32) int32                                                                           // function
	Scale96ToForm(ASize int32) int32                                                                               // function
	ScaleFormTo96(ASize int32) int32                                                                               // function
	Scale96ToFont(ASize int32) int32                                                                               // function
	ScaleFontTo96(ASize int32) int32                                                                               // function
	ScaleScreenToFont(ASize int32) int32                                                                           // function
	ScaleFontToScreen(ASize int32) int32                                                                           // function
	Scale96ToScreen(ASize int32) int32                                                                             // function
	ScaleScreenTo96(ASize int32) int32                                                                             // function
	AutoSizePhases() TControlAutoSizePhases                                                                        // function
	AutoSizeDelayed() bool                                                                                         // function
	AutoSizeDelayedReport() string                                                                                 // function
	AutoSizeDelayedHandle() bool                                                                                   // function
	AnchoredControlCount() int32                                                                                   // function
	GetCanvasScaleFactor() (resultDouble float64)                                                                  // function
	GetDefaultWidth() int32                                                                                        // function
	GetDefaultHeight() int32                                                                                       // function
	GetDefaultColor(DefaultColorType TDefaultColorType) TColor                                                     // function
	GetColorResolvingParent() TColor                                                                               // function
	GetRGBColorResolvingParent() TColor                                                                            // function
	GetSidePosition(Side TAnchorKind) int32                                                                        // function
	GetAnchorsDependingOnParent(WithNormalAnchors bool) TAnchors                                                   // function
	IsParentOf(AControl IControl) bool                                                                             // function
	GetTopParent() IControl                                                                                        // function
	FindSubComponent(AName string) IComponent                                                                      // function
	IsVisible() bool                                                                                               // function
	IsControlVisible() bool                                                                                        // function
	IsEnabled() bool                                                                                               // function
	IsParentColor() bool                                                                                           // function
	IsParentFont() bool                                                                                            // function
	FormIsUpdating() bool                                                                                          // function
	IsProcessingPaintMsg() bool                                                                                    // function
	CheckChildClassAllowed(ChildClass TClass, ExceptionOnInvalid bool) bool                                        // function
	GetTextBuf(Buffer *string, BufSize int32) int32                                                                // function
	GetTextLen() int32                                                                                             // function
	Perform(Msg uint32, WParam WPARAM, LParam LPARAM) LRESULT                                                      // function
	ScreenToClient(APoint *TPoint) (resultPoint TPoint)                                                            // function
	ClientToScreen(APoint *TPoint) (resultPoint TPoint)                                                            // function
	ClientToScreen1(ARect *TRect) (resultRect TRect)                                                               // function
	ScreenToControl(APoint *TPoint) (resultPoint TPoint)                                                           // function
	ControlToScreen(APoint *TPoint) (resultPoint TPoint)                                                           // function
	ClientToParent(Point *TPoint, AParent IWinControl) (resultPoint TPoint)                                        // function
	ParentToClient(Point *TPoint, AParent IWinControl) (resultPoint TPoint)                                        // function
	GetChildrenRect(Scrolled bool) (resultRect TRect)                                                              // function
	HandleObjectShouldBeVisible() bool                                                                             // function
	ParentDestroyingHandle() bool                                                                                  // function
	ParentHandlesAllocated() bool                                                                                  // function
	HasHelp() bool                                                                                                 // function
	UseRightToLeftAlignment() bool                                                                                 // function
	UseRightToLeftReading() bool                                                                                   // function
	UseRightToLeftScrollBar() bool                                                                                 // function
	IsRightToLeft() bool                                                                                           // function
	DragDrop(Source IObject, X, Y int32)                                                                           // procedure
	Dock(NewDockSite IWinControl, ARect *TRect)                                                                    // procedure
	AdjustSize()                                                                                                   // procedure
	AnchorToNeighbour(Side TAnchorKind, Space TSpacingSize, Sibling IControl)                                      // procedure
	AnchorParallel(Side TAnchorKind, Space TSpacingSize, Sibling IControl)                                         // procedure
	AnchorHorizontalCenterTo(Sibling IControl)                                                                     // procedure
	AnchorVerticalCenterTo(Sibling IControl)                                                                       // procedure
	AnchorToCompanion(Side TAnchorKind, Space TSpacingSize, Sibling IControl, FreeCompositeSide bool)              // procedure
	AnchorSame(Side TAnchorKind, Sibling IControl)                                                                 // procedure
	AnchorAsAlign(TheAlign TAlign, Space TSpacingSize)                                                             // procedure
	AnchorClient(Space TSpacingSize)                                                                               // procedure
	SetBounds(aLeft, aTop, aWidth, aHeight int32)                                                                  // procedure
	SetInitialBounds(aLeft, aTop, aWidth, aHeight int32)                                                           // procedure
	SetBoundsKeepBase(aLeft, aTop, aWidth, aHeight int32)                                                          // procedure
	GetPreferredSize(PreferredWidth, PreferredHeight *int32, Raw bool, WithThemeSpace bool)                        // procedure
	CNPreferredSizeChanged()                                                                                       // procedure
	InvalidatePreferredSize()                                                                                      // procedure
	DisableAutoSizing()                                                                                            // procedure
	EnableAutoSizing()                                                                                             // procedure
	UpdateBaseBounds(StoreBounds, StoreParentClientSize, UseLoadedValues bool)                                     // procedure
	WriteLayoutDebugReport(Prefix string)                                                                          // procedure
	AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth int32)          // procedure
	ShouldAutoAdjust(AWidth, AHeight *bool)                                                                        // procedure
	FixDesignFontsPPI(ADesignTimePPI int32)                                                                        // procedure
	ScaleFontsPPI(AToPPI int32, AProportion float64)                                                               // procedure
	EditingDone()                                                                                                  // procedure
	ExecuteDefaultAction()                                                                                         // procedure
	ExecuteCancelAction()                                                                                          // procedure
	BeginDrag(Immediate bool, Threshold int32)                                                                     // procedure
	EndDrag(Drop bool)                                                                                             // procedure
	BringToFront()                                                                                                 // procedure
	Hide()                                                                                                         // procedure
	Refresh()                                                                                                      // procedure
	Repaint()                                                                                                      // procedure
	Invalidate()                                                                                                   // procedure
	CheckNewParent(AParent IWinControl)                                                                            // procedure
	SendToBack()                                                                                                   // procedure
	SetTempCursor(Value TCursor)                                                                                   // procedure
	UpdateRolesForForm()                                                                                           // procedure
	ActiveDefaultControlChanged(NewControl IControl)                                                               // procedure
	SetTextBuf(Buffer string)                                                                                      // procedure
	Show()                                                                                                         // procedure
	Update()                                                                                                       // procedure
	InitiateAction()                                                                                               // procedure
	ShowHelp()                                                                                                     // procedure
	SetOnChangeBounds(fn TNotifyEvent)                                                                             // property event
	SetOnClick(fn TNotifyEvent)                                                                                    // property event
	SetOnResize(fn TNotifyEvent)                                                                                   // property event
	SetOnShowHint(fn TControlShowHintEvent)                                                                        // property event
}

// TControl Parent: TLCLComponent
type TControl struct {
	TLCLComponent
	changeBoundsPtr uintptr
	clickPtr        uintptr
	resizePtr       uintptr
	showHintPtr     uintptr
}

func NewControl(TheOwner IComponent) IControl {
	r1 := LCL().SysCallN(1031, GetObjectUintptr(TheOwner))
	return AsControl(r1)
}

func (m *TControl) AnchoredControls(Index int32) IControl {
	r1 := LCL().SysCallN(996, m.Instance(), uintptr(Index))
	return AsControl(r1)
}

func (m *TControl) BaseBounds() (resultRect TRect) {
	LCL().SysCallN(1004, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TControl) ReadBounds() (resultRect TRect) {
	LCL().SysCallN(1101, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TControl) BaseParentClientSize() (resultSize TSize) {
	LCL().SysCallN(1005, m.Instance(), uintptr(unsafePointer(&resultSize)))
	return
}

func (m *TControl) AccessibleName() string {
	r1 := LCL().SysCallN(975, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TControl) SetAccessibleName(AValue string) {
	LCL().SysCallN(975, 1, m.Instance(), PascalStr(AValue))
}

func (m *TControl) AccessibleDescription() string {
	r1 := LCL().SysCallN(974, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TControl) SetAccessibleDescription(AValue string) {
	LCL().SysCallN(974, 1, m.Instance(), PascalStr(AValue))
}

func (m *TControl) AccessibleValue() string {
	r1 := LCL().SysCallN(977, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TControl) SetAccessibleValue(AValue string) {
	LCL().SysCallN(977, 1, m.Instance(), PascalStr(AValue))
}

func (m *TControl) AccessibleRole() TLazAccessibilityRole {
	r1 := LCL().SysCallN(976, 0, m.Instance(), 0)
	return TLazAccessibilityRole(r1)
}

func (m *TControl) SetAccessibleRole(AValue TLazAccessibilityRole) {
	LCL().SysCallN(976, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Action() IBasicAction {
	r1 := LCL().SysCallN(978, 0, m.Instance(), 0)
	return AsBasicAction(r1)
}

func (m *TControl) SetAction(AValue IBasicAction) {
	LCL().SysCallN(978, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) Align() TAlign {
	r1 := LCL().SysCallN(981, 0, m.Instance(), 0)
	return TAlign(r1)
}

func (m *TControl) SetAlign(AValue TAlign) {
	LCL().SysCallN(981, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Anchors() TAnchors {
	r1 := LCL().SysCallN(997, 0, m.Instance(), 0)
	return TAnchors(r1)
}

func (m *TControl) SetAnchors(AValue TAnchors) {
	LCL().SysCallN(997, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) AnchorSide(Kind TAnchorKind) IAnchorSide {
	r1 := LCL().SysCallN(987, m.Instance(), uintptr(Kind))
	return AsAnchorSide(r1)
}

func (m *TControl) AutoSize() bool {
	r1 := LCL().SysCallN(999, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControl) SetAutoSize(AValue bool) {
	LCL().SysCallN(999, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControl) BorderSpacing() IControlBorderSpacing {
	r1 := LCL().SysCallN(1008, 0, m.Instance(), 0)
	return AsControlBorderSpacing(r1)
}

func (m *TControl) SetBorderSpacing(AValue IControlBorderSpacing) {
	LCL().SysCallN(1008, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) BoundsRect() (resultRect TRect) {
	LCL().SysCallN(1009, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TControl) SetBoundsRect(AValue *TRect) {
	LCL().SysCallN(1009, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TControl) BoundsRectForNewParent() (resultRect TRect) {
	LCL().SysCallN(1010, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TControl) SetBoundsRectForNewParent(AValue *TRect) {
	LCL().SysCallN(1010, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TControl) Caption() string {
	r1 := LCL().SysCallN(1013, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TControl) SetCaption(AValue string) {
	LCL().SysCallN(1013, 1, m.Instance(), PascalStr(AValue))
}

func (m *TControl) CaptureMouseButtons() TCaptureMouseButtons {
	r1 := LCL().SysCallN(1014, 0, m.Instance(), 0)
	return TCaptureMouseButtons(r1)
}

func (m *TControl) SetCaptureMouseButtons(AValue TCaptureMouseButtons) {
	LCL().SysCallN(1014, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) ClientHeight() int32 {
	r1 := LCL().SysCallN(1018, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetClientHeight(AValue int32) {
	LCL().SysCallN(1018, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) ClientOrigin() (resultPoint TPoint) {
	LCL().SysCallN(1019, m.Instance(), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TControl) ClientRect() (resultRect TRect) {
	LCL().SysCallN(1020, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TControl) ClientWidth() int32 {
	r1 := LCL().SysCallN(1024, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetClientWidth(AValue int32) {
	LCL().SysCallN(1024, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Color() TColor {
	r1 := LCL().SysCallN(1025, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TControl) SetColor(AValue TColor) {
	LCL().SysCallN(1025, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Constraints() ISizeConstraints {
	r1 := LCL().SysCallN(1026, 0, m.Instance(), 0)
	return AsSizeConstraints(r1)
}

func (m *TControl) SetConstraints(AValue ISizeConstraints) {
	LCL().SysCallN(1026, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) ControlOrigin() (resultPoint TPoint) {
	LCL().SysCallN(1027, m.Instance(), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TControl) ControlState() TControlState {
	r1 := LCL().SysCallN(1028, 0, m.Instance(), 0)
	return TControlState(r1)
}

func (m *TControl) SetControlState(AValue TControlState) {
	LCL().SysCallN(1028, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) ControlStyle() TControlStyle {
	r1 := LCL().SysCallN(1029, 0, m.Instance(), 0)
	return TControlStyle(r1)
}

func (m *TControl) SetControlStyle(AValue TControlStyle) {
	LCL().SysCallN(1029, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Enabled() bool {
	r1 := LCL().SysCallN(1042, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControl) SetEnabled(AValue bool) {
	LCL().SysCallN(1042, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControl) Font() IFont {
	r1 := LCL().SysCallN(1050, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TControl) SetFont(AValue IFont) {
	LCL().SysCallN(1050, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) IsControl() bool {
	r1 := LCL().SysCallN(1080, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControl) SetIsControl(AValue bool) {
	LCL().SysCallN(1080, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControl) MouseInClient() bool {
	r1 := LCL().SysCallN(1093, m.Instance())
	return GoBool(r1)
}

func (m *TControl) Parent() IWinControl {
	r1 := LCL().SysCallN(1094, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TControl) SetParent(AValue IWinControl) {
	LCL().SysCallN(1094, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) PopupMenu() IPopupMenu {
	r1 := LCL().SysCallN(1100, 0, m.Instance(), 0)
	return AsPopupMenu(r1)
}

func (m *TControl) SetPopupMenu(AValue IPopupMenu) {
	LCL().SysCallN(1100, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) ShowHint() bool {
	r1 := LCL().SysCallN(1131, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControl) SetShowHint(AValue bool) {
	LCL().SysCallN(1131, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControl) Visible() bool {
	r1 := LCL().SysCallN(1142, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControl) SetVisible(AValue bool) {
	LCL().SysCallN(1142, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControl) DockOrientation() TDockOrientation {
	r1 := LCL().SysCallN(1036, 0, m.Instance(), 0)
	return TDockOrientation(r1)
}

func (m *TControl) SetDockOrientation(AValue TDockOrientation) {
	LCL().SysCallN(1036, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Floating() bool {
	r1 := LCL().SysCallN(1048, m.Instance())
	return GoBool(r1)
}

func (m *TControl) FloatingDockSiteClass() TWinControlClass {
	r1 := LCL().SysCallN(1049, 0, m.Instance(), 0)
	return TWinControlClass(r1)
}

func (m *TControl) SetFloatingDockSiteClass(AValue TWinControlClass) {
	LCL().SysCallN(1049, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) HostDockSite() IWinControl {
	r1 := LCL().SysCallN(1076, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TControl) SetHostDockSite(AValue IWinControl) {
	LCL().SysCallN(1076, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) LRDockWidth() int32 {
	r1 := LCL().SysCallN(1089, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetLRDockWidth(AValue int32) {
	LCL().SysCallN(1089, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) TBDockHeight() int32 {
	r1 := LCL().SysCallN(1132, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetTBDockHeight(AValue int32) {
	LCL().SysCallN(1132, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) UndockHeight() int32 {
	r1 := LCL().SysCallN(1134, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetUndockHeight(AValue int32) {
	LCL().SysCallN(1134, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) UndockWidth() int32 {
	r1 := LCL().SysCallN(1135, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetUndockWidth(AValue int32) {
	LCL().SysCallN(1135, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) BiDiMode() TBiDiMode {
	r1 := LCL().SysCallN(1007, 0, m.Instance(), 0)
	return TBiDiMode(r1)
}

func (m *TControl) SetBiDiMode(AValue TBiDiMode) {
	LCL().SysCallN(1007, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) ParentBiDiMode() bool {
	r1 := LCL().SysCallN(1095, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControl) SetParentBiDiMode(AValue bool) {
	LCL().SysCallN(1095, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControl) AnchorSideLeft() IAnchorSide {
	r1 := LCL().SysCallN(989, 0, m.Instance(), 0)
	return AsAnchorSide(r1)
}

func (m *TControl) SetAnchorSideLeft(AValue IAnchorSide) {
	LCL().SysCallN(989, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) AnchorSideTop() IAnchorSide {
	r1 := LCL().SysCallN(991, 0, m.Instance(), 0)
	return AsAnchorSide(r1)
}

func (m *TControl) SetAnchorSideTop(AValue IAnchorSide) {
	LCL().SysCallN(991, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) AnchorSideRight() IAnchorSide {
	r1 := LCL().SysCallN(990, 0, m.Instance(), 0)
	return AsAnchorSide(r1)
}

func (m *TControl) SetAnchorSideRight(AValue IAnchorSide) {
	LCL().SysCallN(990, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) AnchorSideBottom() IAnchorSide {
	r1 := LCL().SysCallN(988, 0, m.Instance(), 0)
	return AsAnchorSide(r1)
}

func (m *TControl) SetAnchorSideBottom(AValue IAnchorSide) {
	LCL().SysCallN(988, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TControl) Cursor() TCursor {
	r1 := LCL().SysCallN(1033, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TControl) SetCursor(AValue TCursor) {
	LCL().SysCallN(1033, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Left() int32 {
	r1 := LCL().SysCallN(1090, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetLeft(AValue int32) {
	LCL().SysCallN(1090, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Height() int32 {
	r1 := LCL().SysCallN(1070, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetHeight(AValue int32) {
	LCL().SysCallN(1070, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Hint() string {
	r1 := LCL().SysCallN(1075, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TControl) SetHint(AValue string) {
	LCL().SysCallN(1075, 1, m.Instance(), PascalStr(AValue))
}

func (m *TControl) Top() int32 {
	r1 := LCL().SysCallN(1133, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetTop(AValue int32) {
	LCL().SysCallN(1133, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) Width() int32 {
	r1 := LCL().SysCallN(1143, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControl) SetWidth(AValue int32) {
	LCL().SysCallN(1143, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) HelpType() THelpType {
	r1 := LCL().SysCallN(1073, 0, m.Instance(), 0)
	return THelpType(r1)
}

func (m *TControl) SetHelpType(AValue THelpType) {
	LCL().SysCallN(1073, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) HelpKeyword() string {
	r1 := LCL().SysCallN(1072, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TControl) SetHelpKeyword(AValue string) {
	LCL().SysCallN(1072, 1, m.Instance(), PascalStr(AValue))
}

func (m *TControl) HelpContext() THelpContext {
	r1 := LCL().SysCallN(1071, 0, m.Instance(), 0)
	return THelpContext(r1)
}

func (m *TControl) SetHelpContext(AValue THelpContext) {
	LCL().SysCallN(1071, 1, m.Instance(), uintptr(AValue))
}

func (m *TControl) ManualDock(NewDockSite IWinControl, DropControl IControl, ControlSide TAlign, KeepDockSiteSize bool) bool {
	r1 := LCL().SysCallN(1091, m.Instance(), GetObjectUintptr(NewDockSite), GetObjectUintptr(DropControl), uintptr(ControlSide), PascalBool(KeepDockSiteSize))
	return GoBool(r1)
}

func (m *TControl) ManualFloat(TheScreenRect *TRect, KeepDockSiteSize bool) bool {
	r1 := LCL().SysCallN(1092, m.Instance(), uintptr(unsafePointer(TheScreenRect)), PascalBool(KeepDockSiteSize))
	return GoBool(r1)
}

func (m *TControl) ReplaceDockedControl(Control IControl, NewDockSite IWinControl, DropControl IControl, ControlSide TAlign) bool {
	r1 := LCL().SysCallN(1104, m.Instance(), GetObjectUintptr(Control), GetObjectUintptr(NewDockSite), GetObjectUintptr(DropControl), uintptr(ControlSide))
	return GoBool(r1)
}

func (m *TControl) Docked() bool {
	r1 := LCL().SysCallN(1037, m.Instance())
	return GoBool(r1)
}

func (m *TControl) Dragging() bool {
	r1 := LCL().SysCallN(1039, m.Instance())
	return GoBool(r1)
}

func (m *TControl) GetAccessibleObject() ILazAccessibleObject {
	r1 := LCL().SysCallN(1052, m.Instance())
	return AsLazAccessibleObject(r1)
}

func (m *TControl) CreateAccessibleObject() ILazAccessibleObject {
	r1 := LCL().SysCallN(1032, m.Instance())
	return AsLazAccessibleObject(r1)
}

func (m *TControl) GetSelectedChildAccessibleObject() ILazAccessibleObject {
	r1 := LCL().SysCallN(1063, m.Instance())
	return AsLazAccessibleObject(r1)
}

func (m *TControl) GetChildAccessibleObjectAtPos(APos *TPoint) ILazAccessibleObject {
	r1 := LCL().SysCallN(1055, m.Instance(), uintptr(unsafePointer(APos)))
	return AsLazAccessibleObject(r1)
}

func (m *TControl) ScaleDesignToForm(ASize int32) int32 {
	r1 := LCL().SysCallN(1108, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) ScaleFormToDesign(ASize int32) int32 {
	r1 := LCL().SysCallN(1113, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) Scale96ToForm(ASize int32) int32 {
	r1 := LCL().SysCallN(1106, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) ScaleFormTo96(ASize int32) int32 {
	r1 := LCL().SysCallN(1112, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) Scale96ToFont(ASize int32) int32 {
	r1 := LCL().SysCallN(1105, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) ScaleFontTo96(ASize int32) int32 {
	r1 := LCL().SysCallN(1109, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) ScaleScreenToFont(ASize int32) int32 {
	r1 := LCL().SysCallN(1115, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) ScaleFontToScreen(ASize int32) int32 {
	r1 := LCL().SysCallN(1110, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) Scale96ToScreen(ASize int32) int32 {
	r1 := LCL().SysCallN(1107, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) ScaleScreenTo96(ASize int32) int32 {
	r1 := LCL().SysCallN(1114, m.Instance(), uintptr(ASize))
	return int32(r1)
}

func (m *TControl) AutoSizePhases() TControlAutoSizePhases {
	r1 := LCL().SysCallN(1003, m.Instance())
	return TControlAutoSizePhases(r1)
}

func (m *TControl) AutoSizeDelayed() bool {
	r1 := LCL().SysCallN(1000, m.Instance())
	return GoBool(r1)
}

func (m *TControl) AutoSizeDelayedReport() string {
	r1 := LCL().SysCallN(1002, m.Instance())
	return GoStr(r1)
}

func (m *TControl) AutoSizeDelayedHandle() bool {
	r1 := LCL().SysCallN(1001, m.Instance())
	return GoBool(r1)
}

func (m *TControl) AnchoredControlCount() int32 {
	r1 := LCL().SysCallN(995, m.Instance())
	return int32(r1)
}

func (m *TControl) GetCanvasScaleFactor() (resultDouble float64) {
	LCL().SysCallN(1054, m.Instance(), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TControl) GetDefaultWidth() int32 {
	r1 := LCL().SysCallN(1060, m.Instance())
	return int32(r1)
}

func (m *TControl) GetDefaultHeight() int32 {
	r1 := LCL().SysCallN(1059, m.Instance())
	return int32(r1)
}

func (m *TControl) GetDefaultColor(DefaultColorType TDefaultColorType) TColor {
	r1 := LCL().SysCallN(1058, m.Instance(), uintptr(DefaultColorType))
	return TColor(r1)
}

func (m *TControl) GetColorResolvingParent() TColor {
	r1 := LCL().SysCallN(1057, m.Instance())
	return TColor(r1)
}

func (m *TControl) GetRGBColorResolvingParent() TColor {
	r1 := LCL().SysCallN(1062, m.Instance())
	return TColor(r1)
}

func (m *TControl) GetSidePosition(Side TAnchorKind) int32 {
	r1 := LCL().SysCallN(1064, m.Instance(), uintptr(Side))
	return int32(r1)
}

func (m *TControl) GetAnchorsDependingOnParent(WithNormalAnchors bool) TAnchors {
	r1 := LCL().SysCallN(1053, m.Instance(), PascalBool(WithNormalAnchors))
	return TAnchors(r1)
}

func (m *TControl) IsParentOf(AControl IControl) bool {
	r1 := LCL().SysCallN(1085, m.Instance(), GetObjectUintptr(AControl))
	return GoBool(r1)
}

func (m *TControl) GetTopParent() IControl {
	r1 := LCL().SysCallN(1067, m.Instance())
	return AsControl(r1)
}

func (m *TControl) FindSubComponent(AName string) IComponent {
	r1 := LCL().SysCallN(1046, m.Instance(), PascalStr(AName))
	return AsComponent(r1)
}

func (m *TControl) IsVisible() bool {
	r1 := LCL().SysCallN(1088, m.Instance())
	return GoBool(r1)
}

func (m *TControl) IsControlVisible() bool {
	r1 := LCL().SysCallN(1081, m.Instance())
	return GoBool(r1)
}

func (m *TControl) IsEnabled() bool {
	r1 := LCL().SysCallN(1082, m.Instance())
	return GoBool(r1)
}

func (m *TControl) IsParentColor() bool {
	r1 := LCL().SysCallN(1083, m.Instance())
	return GoBool(r1)
}

func (m *TControl) IsParentFont() bool {
	r1 := LCL().SysCallN(1084, m.Instance())
	return GoBool(r1)
}

func (m *TControl) FormIsUpdating() bool {
	r1 := LCL().SysCallN(1051, m.Instance())
	return GoBool(r1)
}

func (m *TControl) IsProcessingPaintMsg() bool {
	r1 := LCL().SysCallN(1086, m.Instance())
	return GoBool(r1)
}

func (m *TControl) CheckChildClassAllowed(ChildClass TClass, ExceptionOnInvalid bool) bool {
	r1 := LCL().SysCallN(1015, m.Instance(), uintptr(ChildClass), PascalBool(ExceptionOnInvalid))
	return GoBool(r1)
}

func (m *TControl) GetTextBuf(Buffer *string, BufSize int32) int32 {
	r1 := sysCallGetTextBuf(1065, m.Instance(), Buffer, BufSize)
	return int32(r1)
}

func (m *TControl) GetTextLen() int32 {
	r1 := LCL().SysCallN(1066, m.Instance())
	return int32(r1)
}

func (m *TControl) Perform(Msg uint32, WParam WPARAM, LParam LPARAM) LRESULT {
	r1 := LCL().SysCallN(1099, m.Instance(), uintptr(Msg), uintptr(WParam), uintptr(LParam))
	return LRESULT(r1)
}

func (m *TControl) ScreenToClient(APoint *TPoint) (resultPoint TPoint) {
	LCL().SysCallN(1116, m.Instance(), uintptr(unsafePointer(APoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TControl) ClientToScreen(APoint *TPoint) (resultPoint TPoint) {
	LCL().SysCallN(1022, m.Instance(), uintptr(unsafePointer(APoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TControl) ClientToScreen1(ARect *TRect) (resultRect TRect) {
	LCL().SysCallN(1023, m.Instance(), uintptr(unsafePointer(ARect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TControl) ScreenToControl(APoint *TPoint) (resultPoint TPoint) {
	LCL().SysCallN(1117, m.Instance(), uintptr(unsafePointer(APoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TControl) ControlToScreen(APoint *TPoint) (resultPoint TPoint) {
	LCL().SysCallN(1030, m.Instance(), uintptr(unsafePointer(APoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TControl) ClientToParent(Point *TPoint, AParent IWinControl) (resultPoint TPoint) {
	LCL().SysCallN(1021, m.Instance(), uintptr(unsafePointer(Point)), GetObjectUintptr(AParent), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TControl) ParentToClient(Point *TPoint, AParent IWinControl) (resultPoint TPoint) {
	LCL().SysCallN(1098, m.Instance(), uintptr(unsafePointer(Point)), GetObjectUintptr(AParent), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TControl) GetChildrenRect(Scrolled bool) (resultRect TRect) {
	LCL().SysCallN(1056, m.Instance(), PascalBool(Scrolled), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TControl) HandleObjectShouldBeVisible() bool {
	r1 := LCL().SysCallN(1068, m.Instance())
	return GoBool(r1)
}

func (m *TControl) ParentDestroyingHandle() bool {
	r1 := LCL().SysCallN(1096, m.Instance())
	return GoBool(r1)
}

func (m *TControl) ParentHandlesAllocated() bool {
	r1 := LCL().SysCallN(1097, m.Instance())
	return GoBool(r1)
}

func (m *TControl) HasHelp() bool {
	r1 := LCL().SysCallN(1069, m.Instance())
	return GoBool(r1)
}

func (m *TControl) UseRightToLeftAlignment() bool {
	r1 := LCL().SysCallN(1139, m.Instance())
	return GoBool(r1)
}

func (m *TControl) UseRightToLeftReading() bool {
	r1 := LCL().SysCallN(1140, m.Instance())
	return GoBool(r1)
}

func (m *TControl) UseRightToLeftScrollBar() bool {
	r1 := LCL().SysCallN(1141, m.Instance())
	return GoBool(r1)
}

func (m *TControl) IsRightToLeft() bool {
	r1 := LCL().SysCallN(1087, m.Instance())
	return GoBool(r1)
}

func ControlClass() TClass {
	ret := LCL().SysCallN(1017)
	return TClass(ret)
}

func (m *TControl) DragDrop(Source IObject, X, Y int32) {
	LCL().SysCallN(1038, m.Instance(), GetObjectUintptr(Source), uintptr(X), uintptr(Y))
}

func (m *TControl) Dock(NewDockSite IWinControl, ARect *TRect) {
	LCL().SysCallN(1035, m.Instance(), GetObjectUintptr(NewDockSite), uintptr(unsafePointer(ARect)))
}

func (m *TControl) AdjustSize() {
	LCL().SysCallN(980, m.Instance())
}

func (m *TControl) AnchorToNeighbour(Side TAnchorKind, Space TSpacingSize, Sibling IControl) {
	LCL().SysCallN(993, m.Instance(), uintptr(Side), uintptr(Space), GetObjectUintptr(Sibling))
}

func (m *TControl) AnchorParallel(Side TAnchorKind, Space TSpacingSize, Sibling IControl) {
	LCL().SysCallN(985, m.Instance(), uintptr(Side), uintptr(Space), GetObjectUintptr(Sibling))
}

func (m *TControl) AnchorHorizontalCenterTo(Sibling IControl) {
	LCL().SysCallN(984, m.Instance(), GetObjectUintptr(Sibling))
}

func (m *TControl) AnchorVerticalCenterTo(Sibling IControl) {
	LCL().SysCallN(994, m.Instance(), GetObjectUintptr(Sibling))
}

func (m *TControl) AnchorToCompanion(Side TAnchorKind, Space TSpacingSize, Sibling IControl, FreeCompositeSide bool) {
	LCL().SysCallN(992, m.Instance(), uintptr(Side), uintptr(Space), GetObjectUintptr(Sibling), PascalBool(FreeCompositeSide))
}

func (m *TControl) AnchorSame(Side TAnchorKind, Sibling IControl) {
	LCL().SysCallN(986, m.Instance(), uintptr(Side), GetObjectUintptr(Sibling))
}

func (m *TControl) AnchorAsAlign(TheAlign TAlign, Space TSpacingSize) {
	LCL().SysCallN(982, m.Instance(), uintptr(TheAlign), uintptr(Space))
}

func (m *TControl) AnchorClient(Space TSpacingSize) {
	LCL().SysCallN(983, m.Instance(), uintptr(Space))
}

func (m *TControl) SetBounds(aLeft, aTop, aWidth, aHeight int32) {
	LCL().SysCallN(1119, m.Instance(), uintptr(aLeft), uintptr(aTop), uintptr(aWidth), uintptr(aHeight))
}

func (m *TControl) SetInitialBounds(aLeft, aTop, aWidth, aHeight int32) {
	LCL().SysCallN(1121, m.Instance(), uintptr(aLeft), uintptr(aTop), uintptr(aWidth), uintptr(aHeight))
}

func (m *TControl) SetBoundsKeepBase(aLeft, aTop, aWidth, aHeight int32) {
	LCL().SysCallN(1120, m.Instance(), uintptr(aLeft), uintptr(aTop), uintptr(aWidth), uintptr(aHeight))
}

func (m *TControl) GetPreferredSize(PreferredWidth, PreferredHeight *int32, Raw bool, WithThemeSpace bool) {
	var result0 uintptr
	var result1 uintptr
	LCL().SysCallN(1061, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)), PascalBool(Raw), PascalBool(WithThemeSpace))
	*PreferredWidth = int32(result0)
	*PreferredHeight = int32(result1)
}

func (m *TControl) CNPreferredSizeChanged() {
	LCL().SysCallN(1012, m.Instance())
}

func (m *TControl) InvalidatePreferredSize() {
	LCL().SysCallN(1079, m.Instance())
}

func (m *TControl) DisableAutoSizing() {
	LCL().SysCallN(1034, m.Instance())
}

func (m *TControl) EnableAutoSizing() {
	LCL().SysCallN(1041, m.Instance())
}

func (m *TControl) UpdateBaseBounds(StoreBounds, StoreParentClientSize, UseLoadedValues bool) {
	LCL().SysCallN(1137, m.Instance(), PascalBool(StoreBounds), PascalBool(StoreParentClientSize), PascalBool(UseLoadedValues))
}

func (m *TControl) WriteLayoutDebugReport(Prefix string) {
	LCL().SysCallN(1144, m.Instance(), PascalStr(Prefix))
}

func (m *TControl) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI, AToPPI, AOldFormWidth, ANewFormWidth int32) {
	LCL().SysCallN(998, m.Instance(), uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func (m *TControl) ShouldAutoAdjust(AWidth, AHeight *bool) {
	var result0 uintptr
	var result1 uintptr
	LCL().SysCallN(1128, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)))
	*AWidth = GoBool(result0)
	*AHeight = GoBool(result1)
}

func (m *TControl) FixDesignFontsPPI(ADesignTimePPI int32) {
	LCL().SysCallN(1047, m.Instance(), uintptr(ADesignTimePPI))
}

func (m *TControl) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	LCL().SysCallN(1111, m.Instance(), uintptr(AToPPI), uintptr(unsafePointer(&AProportion)))
}

func (m *TControl) EditingDone() {
	LCL().SysCallN(1040, m.Instance())
}

func (m *TControl) ExecuteDefaultAction() {
	LCL().SysCallN(1045, m.Instance())
}

func (m *TControl) ExecuteCancelAction() {
	LCL().SysCallN(1044, m.Instance())
}

func (m *TControl) BeginDrag(Immediate bool, Threshold int32) {
	LCL().SysCallN(1006, m.Instance(), PascalBool(Immediate), uintptr(Threshold))
}

func (m *TControl) EndDrag(Drop bool) {
	LCL().SysCallN(1043, m.Instance(), PascalBool(Drop))
}

func (m *TControl) BringToFront() {
	LCL().SysCallN(1011, m.Instance())
}

func (m *TControl) Hide() {
	LCL().SysCallN(1074, m.Instance())
}

func (m *TControl) Refresh() {
	LCL().SysCallN(1102, m.Instance())
}

func (m *TControl) Repaint() {
	LCL().SysCallN(1103, m.Instance())
}

func (m *TControl) Invalidate() {
	LCL().SysCallN(1078, m.Instance())
}

func (m *TControl) CheckNewParent(AParent IWinControl) {
	LCL().SysCallN(1016, m.Instance(), GetObjectUintptr(AParent))
}

func (m *TControl) SendToBack() {
	LCL().SysCallN(1118, m.Instance())
}

func (m *TControl) SetTempCursor(Value TCursor) {
	LCL().SysCallN(1126, m.Instance(), uintptr(Value))
}

func (m *TControl) UpdateRolesForForm() {
	LCL().SysCallN(1138, m.Instance())
}

func (m *TControl) ActiveDefaultControlChanged(NewControl IControl) {
	LCL().SysCallN(979, m.Instance(), GetObjectUintptr(NewControl))
}

func (m *TControl) SetTextBuf(Buffer string) {
	LCL().SysCallN(1127, m.Instance(), PascalStr(Buffer))
}

func (m *TControl) Show() {
	LCL().SysCallN(1129, m.Instance())
}

func (m *TControl) Update() {
	LCL().SysCallN(1136, m.Instance())
}

func (m *TControl) InitiateAction() {
	LCL().SysCallN(1077, m.Instance())
}

func (m *TControl) ShowHelp() {
	LCL().SysCallN(1130, m.Instance())
}

func (m *TControl) SetOnChangeBounds(fn TNotifyEvent) {
	if m.changeBoundsPtr != 0 {
		RemoveEventElement(m.changeBoundsPtr)
	}
	m.changeBoundsPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1122, m.Instance(), m.changeBoundsPtr)
}

func (m *TControl) SetOnClick(fn TNotifyEvent) {
	if m.clickPtr != 0 {
		RemoveEventElement(m.clickPtr)
	}
	m.clickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1123, m.Instance(), m.clickPtr)
}

func (m *TControl) SetOnResize(fn TNotifyEvent) {
	if m.resizePtr != 0 {
		RemoveEventElement(m.resizePtr)
	}
	m.resizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1124, m.Instance(), m.resizePtr)
}

func (m *TControl) SetOnShowHint(fn TControlShowHintEvent) {
	if m.showHintPtr != 0 {
		RemoveEventElement(m.showHintPtr)
	}
	m.showHintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1125, m.Instance(), m.showHintPtr)
}
