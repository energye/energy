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

// ICustomForm Parent: ICustomDesignControl
type ICustomForm interface {
	ICustomDesignControl
	Active() bool                                                                // property
	ActiveControl() IWinControl                                                  // property
	SetActiveControl(AValue IWinControl)                                         // property
	ActiveDefaultControl() IControl                                              // property
	SetActiveDefaultControl(AValue IControl)                                     // property
	AllowDropFiles() bool                                                        // property
	SetAllowDropFiles(AValue bool)                                               // property
	AlphaBlend() bool                                                            // property
	SetAlphaBlend(AValue bool)                                                   // property
	AlphaBlendValue() Byte                                                       // property
	SetAlphaBlendValue(AValue Byte)                                              // property
	AutoScroll() bool                                                            // property
	SetAutoScroll(AValue bool)                                                   // property
	BorderIcons() TBorderIcons                                                   // property
	SetBorderIcons(AValue TBorderIcons)                                          // property
	BorderStyleForFormBorderStyle() TFormBorderStyle                             // property
	SetBorderStyleForFormBorderStyle(AValue TFormBorderStyle)                    // property
	CancelControl() IControl                                                     // property
	SetCancelControl(AValue IControl)                                            // property
	DefaultControl() IControl                                                    // property
	SetDefaultControl(AValue IControl)                                           // property
	DefaultMonitor() TDefaultMonitor                                             // property
	SetDefaultMonitor(AValue TDefaultMonitor)                                    // property
	Designer() IIDesigner                                                        // property
	SetDesigner(AValue IIDesigner)                                               // property
	EffectiveShowInTaskBar() TShowInTaskBar                                      // property
	FormState() TFormState                                                       // property
	FormStyle() TFormStyle                                                       // property
	SetFormStyle(AValue TFormStyle)                                              // property
	HelpFile() string                                                            // property
	SetHelpFile(AValue string)                                                   // property
	Icon() IIcon                                                                 // property
	SetIcon(AValue IIcon)                                                        // property
	KeyPreview() bool                                                            // property
	SetKeyPreview(AValue bool)                                                   // property
	MDIChildren(I int32) ICustomForm                                             // property
	Menu() IMainMenu                                                             // property
	SetMenu(AValue IMainMenu)                                                    // property
	ModalResult() TModalResult                                                   // property
	SetModalResult(AValue TModalResult)                                          // property
	Monitor() IMonitor                                                           // property
	LastActiveControl() IWinControl                                              // property
	PopupMode() TPopupMode                                                       // property
	SetPopupMode(AValue TPopupMode)                                              // property
	PopupParent() ICustomForm                                                    // property
	SetPopupParent(AValue ICustomForm)                                           // property
	SnapOptions() IWindowMagnetOptions                                           // property
	SetSnapOptions(AValue IWindowMagnetOptions)                                  // property
	ScreenSnap() bool                                                            // property
	SetScreenSnap(AValue bool)                                                   // property
	SnapBuffer() int32                                                           // property
	SetSnapBuffer(AValue int32)                                                  // property
	ParentFont() bool                                                            // property
	SetParentFont(AValue bool)                                                   // property
	Position() TPosition                                                         // property
	SetPosition(AValue TPosition)                                                // property
	RestoredLeft() int32                                                         // property
	RestoredTop() int32                                                          // property
	RestoredWidth() int32                                                        // property
	RestoredHeight() int32                                                       // property
	ShowInTaskBar() TShowInTaskBar                                               // property
	SetShowInTaskBar(AValue TShowInTaskBar)                                      // property
	WindowState() TWindowState                                                   // property
	SetWindowState(AValue TWindowState)                                          // property
	BigIconHandle() HICON                                                        // function
	CloseQuery() bool                                                            // function
	GetFormImage() IBitmap                                                       // function
	GetRolesForControl(AControl IControl) TControlRolesForForm                   // function
	GetRealPopupParent() ICustomForm                                             // function
	IsShortcut(Message *TLMKey) bool                                             // function
	SetFocusedControl(Control IWinControl) bool                                  // function
	ShowModal() int32                                                            // function
	SmallIconHandle() HICON                                                      // function
	WantChildKey(Child IControl, Message *TLMessage) bool                        // function
	ActiveMDIChild() ICustomForm                                                 // function
	GetMDIChildren(AIndex int32) ICustomForm                                     // function
	MDIChildCount() int32                                                        // function
	Close()                                                                      // procedure
	DefocusControl(Control IWinControl, Removing bool)                           // procedure
	DestroyWnd()                                                                 // procedure
	EnsureVisible(AMoveToTop bool)                                               // procedure
	FocusControl(WinControl IWinControl)                                         // procedure
	IntfHelp(AComponent IComponent)                                              // procedure
	MakeFullyVisible(AMonitor IMonitor, UseWorkarea bool)                        // procedure
	Release()                                                                    // procedure
	SetRestoredBounds(ALeft, ATop, AWidth, AHeight int32, ADefaultPosition bool) // procedure
	ShowOnTop()                                                                  // procedure
	AutoScale()                                                                  // procedure
	SetOnActivate(fn TNotifyEvent)                                               // property event
	SetOnClose(fn TCloseEvent)                                                   // property event
	SetOnCloseQuery(fn TCloseQueryEvent)                                         // property event
	SetOnCreate(fn TNotifyEvent)                                                 // property event
	SetOnDeactivate(fn TNotifyEvent)                                             // property event
	SetOnDestroy(fn TNotifyEvent)                                                // property event
	SetOnDropFiles(fn TDropFilesEvent)                                           // property event
	SetOnHelp(fn THelpEvent)                                                     // property event
	SetOnHide(fn TNotifyEvent)                                                   // property event
	SetOnShortcut(fn TShortCutEvent)                                             // property event
	SetOnShow(fn TNotifyEvent)                                                   // property event
	SetOnShowModalFinished(fn TModalDialogFinished)                              // property event
	SetOnWindowStateChange(fn TNotifyEvent)                                      // property event
}

// TCustomForm Parent: TCustomDesignControl
type TCustomForm struct {
	TCustomDesignControl
	activatePtr          uintptr
	closePtr             uintptr
	closeQueryPtr        uintptr
	createPtr            uintptr
	deactivatePtr        uintptr
	destroyPtr           uintptr
	dropFilesPtr         uintptr
	helpPtr              uintptr
	hidePtr              uintptr
	shortcutPtr          uintptr
	showPtr              uintptr
	showModalFinishedPtr uintptr
	windowStateChangePtr uintptr
}

func NewCustomForm(AOwner IComponent) ICustomForm {
	r1 := LCL().SysCallN(1673, GetObjectUintptr(AOwner))
	return AsCustomForm(r1)
}

func NewCustomFormNew(AOwner IComponent, Num int32) ICustomForm {
	r1 := LCL().SysCallN(1674, GetObjectUintptr(AOwner), uintptr(Num))
	return AsCustomForm(r1)
}

func (m *TCustomForm) Active() bool {
	r1 := LCL().SysCallN(1657, m.Instance())
	return GoBool(r1)
}

func (m *TCustomForm) ActiveControl() IWinControl {
	r1 := LCL().SysCallN(1658, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TCustomForm) SetActiveControl(AValue IWinControl) {
	LCL().SysCallN(1658, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) ActiveDefaultControl() IControl {
	r1 := LCL().SysCallN(1659, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TCustomForm) SetActiveDefaultControl(AValue IControl) {
	LCL().SysCallN(1659, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) AllowDropFiles() bool {
	r1 := LCL().SysCallN(1661, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomForm) SetAllowDropFiles(AValue bool) {
	LCL().SysCallN(1661, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomForm) AlphaBlend() bool {
	r1 := LCL().SysCallN(1662, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomForm) SetAlphaBlend(AValue bool) {
	LCL().SysCallN(1662, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomForm) AlphaBlendValue() Byte {
	r1 := LCL().SysCallN(1663, 0, m.Instance(), 0)
	return Byte(r1)
}

func (m *TCustomForm) SetAlphaBlendValue(AValue Byte) {
	LCL().SysCallN(1663, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) AutoScroll() bool {
	r1 := LCL().SysCallN(1665, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomForm) SetAutoScroll(AValue bool) {
	LCL().SysCallN(1665, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomForm) BorderIcons() TBorderIcons {
	r1 := LCL().SysCallN(1667, 0, m.Instance(), 0)
	return TBorderIcons(r1)
}

func (m *TCustomForm) SetBorderIcons(AValue TBorderIcons) {
	LCL().SysCallN(1667, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) BorderStyleForFormBorderStyle() TFormBorderStyle {
	r1 := LCL().SysCallN(1668, 0, m.Instance(), 0)
	return TFormBorderStyle(r1)
}

func (m *TCustomForm) SetBorderStyleForFormBorderStyle(AValue TFormBorderStyle) {
	LCL().SysCallN(1668, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) CancelControl() IControl {
	r1 := LCL().SysCallN(1669, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TCustomForm) SetCancelControl(AValue IControl) {
	LCL().SysCallN(1669, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) DefaultControl() IControl {
	r1 := LCL().SysCallN(1675, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TCustomForm) SetDefaultControl(AValue IControl) {
	LCL().SysCallN(1675, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) DefaultMonitor() TDefaultMonitor {
	r1 := LCL().SysCallN(1676, 0, m.Instance(), 0)
	return TDefaultMonitor(r1)
}

func (m *TCustomForm) SetDefaultMonitor(AValue TDefaultMonitor) {
	LCL().SysCallN(1676, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) Designer() IIDesigner {
	r1 := LCL().SysCallN(1678, 0, m.Instance(), 0)
	return AsIDesigner(r1)
}

func (m *TCustomForm) SetDesigner(AValue IIDesigner) {
	LCL().SysCallN(1678, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) EffectiveShowInTaskBar() TShowInTaskBar {
	r1 := LCL().SysCallN(1680, m.Instance())
	return TShowInTaskBar(r1)
}

func (m *TCustomForm) FormState() TFormState {
	r1 := LCL().SysCallN(1683, m.Instance())
	return TFormState(r1)
}

func (m *TCustomForm) FormStyle() TFormStyle {
	r1 := LCL().SysCallN(1684, 0, m.Instance(), 0)
	return TFormStyle(r1)
}

func (m *TCustomForm) SetFormStyle(AValue TFormStyle) {
	LCL().SysCallN(1684, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) HelpFile() string {
	r1 := LCL().SysCallN(1689, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomForm) SetHelpFile(AValue string) {
	LCL().SysCallN(1689, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomForm) Icon() IIcon {
	r1 := LCL().SysCallN(1690, 0, m.Instance(), 0)
	return AsIcon(r1)
}

func (m *TCustomForm) SetIcon(AValue IIcon) {
	LCL().SysCallN(1690, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) KeyPreview() bool {
	r1 := LCL().SysCallN(1693, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomForm) SetKeyPreview(AValue bool) {
	LCL().SysCallN(1693, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomForm) MDIChildren(I int32) ICustomForm {
	r1 := LCL().SysCallN(1696, m.Instance(), uintptr(I))
	return AsCustomForm(r1)
}

func (m *TCustomForm) Menu() IMainMenu {
	r1 := LCL().SysCallN(1698, 0, m.Instance(), 0)
	return AsMainMenu(r1)
}

func (m *TCustomForm) SetMenu(AValue IMainMenu) {
	LCL().SysCallN(1698, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) ModalResult() TModalResult {
	r1 := LCL().SysCallN(1699, 0, m.Instance(), 0)
	return TModalResult(r1)
}

func (m *TCustomForm) SetModalResult(AValue TModalResult) {
	LCL().SysCallN(1699, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) Monitor() IMonitor {
	r1 := LCL().SysCallN(1700, m.Instance())
	return AsMonitor(r1)
}

func (m *TCustomForm) LastActiveControl() IWinControl {
	r1 := LCL().SysCallN(1694, m.Instance())
	return AsWinControl(r1)
}

func (m *TCustomForm) PopupMode() TPopupMode {
	r1 := LCL().SysCallN(1702, 0, m.Instance(), 0)
	return TPopupMode(r1)
}

func (m *TCustomForm) SetPopupMode(AValue TPopupMode) {
	LCL().SysCallN(1702, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) PopupParent() ICustomForm {
	r1 := LCL().SysCallN(1703, 0, m.Instance(), 0)
	return AsCustomForm(r1)
}

func (m *TCustomForm) SetPopupParent(AValue ICustomForm) {
	LCL().SysCallN(1703, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) SnapOptions() IWindowMagnetOptions {
	r1 := LCL().SysCallN(1731, 0, m.Instance(), 0)
	return AsWindowMagnetOptions(r1)
}

func (m *TCustomForm) SetSnapOptions(AValue IWindowMagnetOptions) {
	LCL().SysCallN(1731, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) ScreenSnap() bool {
	r1 := LCL().SysCallN(1710, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomForm) SetScreenSnap(AValue bool) {
	LCL().SysCallN(1710, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomForm) SnapBuffer() int32 {
	r1 := LCL().SysCallN(1730, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomForm) SetSnapBuffer(AValue int32) {
	LCL().SysCallN(1730, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) ParentFont() bool {
	r1 := LCL().SysCallN(1701, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomForm) SetParentFont(AValue bool) {
	LCL().SysCallN(1701, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomForm) Position() TPosition {
	r1 := LCL().SysCallN(1704, 0, m.Instance(), 0)
	return TPosition(r1)
}

func (m *TCustomForm) SetPosition(AValue TPosition) {
	LCL().SysCallN(1704, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) RestoredLeft() int32 {
	r1 := LCL().SysCallN(1707, m.Instance())
	return int32(r1)
}

func (m *TCustomForm) RestoredTop() int32 {
	r1 := LCL().SysCallN(1708, m.Instance())
	return int32(r1)
}

func (m *TCustomForm) RestoredWidth() int32 {
	r1 := LCL().SysCallN(1709, m.Instance())
	return int32(r1)
}

func (m *TCustomForm) RestoredHeight() int32 {
	r1 := LCL().SysCallN(1706, m.Instance())
	return int32(r1)
}

func (m *TCustomForm) ShowInTaskBar() TShowInTaskBar {
	r1 := LCL().SysCallN(1726, 0, m.Instance(), 0)
	return TShowInTaskBar(r1)
}

func (m *TCustomForm) SetShowInTaskBar(AValue TShowInTaskBar) {
	LCL().SysCallN(1726, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) WindowState() TWindowState {
	r1 := LCL().SysCallN(1733, 0, m.Instance(), 0)
	return TWindowState(r1)
}

func (m *TCustomForm) SetWindowState(AValue TWindowState) {
	LCL().SysCallN(1733, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) BigIconHandle() HICON {
	r1 := LCL().SysCallN(1666, m.Instance())
	return HICON(r1)
}

func (m *TCustomForm) CloseQuery() bool {
	r1 := LCL().SysCallN(1672, m.Instance())
	return GoBool(r1)
}

func (m *TCustomForm) GetFormImage() IBitmap {
	r1 := LCL().SysCallN(1685, m.Instance())
	return AsBitmap(r1)
}

func (m *TCustomForm) GetRolesForControl(AControl IControl) TControlRolesForForm {
	r1 := LCL().SysCallN(1688, m.Instance(), GetObjectUintptr(AControl))
	return TControlRolesForForm(r1)
}

func (m *TCustomForm) GetRealPopupParent() ICustomForm {
	r1 := LCL().SysCallN(1687, m.Instance())
	return AsCustomForm(r1)
}

func (m *TCustomForm) IsShortcut(Message *TLMKey) bool {
	var result0 uintptr
	r1 := LCL().SysCallN(1692, m.Instance(), uintptr(unsafePointer(&result0)))
	*Message = *(*TLMKey)(getPointer(result0))
	return GoBool(r1)
}

func (m *TCustomForm) SetFocusedControl(Control IWinControl) bool {
	r1 := LCL().SysCallN(1711, m.Instance(), GetObjectUintptr(Control))
	return GoBool(r1)
}

func (m *TCustomForm) ShowModal() int32 {
	r1 := LCL().SysCallN(1727, m.Instance())
	return int32(r1)
}

func (m *TCustomForm) SmallIconHandle() HICON {
	r1 := LCL().SysCallN(1729, m.Instance())
	return HICON(r1)
}

func (m *TCustomForm) WantChildKey(Child IControl, Message *TLMessage) bool {
	var result1 uintptr
	r1 := LCL().SysCallN(1732, m.Instance(), GetObjectUintptr(Child), uintptr(unsafePointer(&result1)))
	*Message = *(*TLMessage)(getPointer(result1))
	return GoBool(r1)
}

func (m *TCustomForm) ActiveMDIChild() ICustomForm {
	r1 := LCL().SysCallN(1660, m.Instance())
	return AsCustomForm(r1)
}

func (m *TCustomForm) GetMDIChildren(AIndex int32) ICustomForm {
	r1 := LCL().SysCallN(1686, m.Instance(), uintptr(AIndex))
	return AsCustomForm(r1)
}

func (m *TCustomForm) MDIChildCount() int32 {
	r1 := LCL().SysCallN(1695, m.Instance())
	return int32(r1)
}

func CustomFormClass() TClass {
	ret := LCL().SysCallN(1670)
	return TClass(ret)
}

func (m *TCustomForm) Close() {
	LCL().SysCallN(1671, m.Instance())
}

func (m *TCustomForm) DefocusControl(Control IWinControl, Removing bool) {
	LCL().SysCallN(1677, m.Instance(), GetObjectUintptr(Control), PascalBool(Removing))
}

func (m *TCustomForm) DestroyWnd() {
	LCL().SysCallN(1679, m.Instance())
}

func (m *TCustomForm) EnsureVisible(AMoveToTop bool) {
	LCL().SysCallN(1681, m.Instance(), PascalBool(AMoveToTop))
}

func (m *TCustomForm) FocusControl(WinControl IWinControl) {
	LCL().SysCallN(1682, m.Instance(), GetObjectUintptr(WinControl))
}

func (m *TCustomForm) IntfHelp(AComponent IComponent) {
	LCL().SysCallN(1691, m.Instance(), GetObjectUintptr(AComponent))
}

func (m *TCustomForm) MakeFullyVisible(AMonitor IMonitor, UseWorkarea bool) {
	LCL().SysCallN(1697, m.Instance(), GetObjectUintptr(AMonitor), PascalBool(UseWorkarea))
}

func (m *TCustomForm) Release() {
	LCL().SysCallN(1705, m.Instance())
}

func (m *TCustomForm) SetRestoredBounds(ALeft, ATop, AWidth, AHeight int32, ADefaultPosition bool) {
	LCL().SysCallN(1725, m.Instance(), uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight), PascalBool(ADefaultPosition))
}

func (m *TCustomForm) ShowOnTop() {
	LCL().SysCallN(1728, m.Instance())
}

func (m *TCustomForm) AutoScale() {
	LCL().SysCallN(1664, m.Instance())
}

func (m *TCustomForm) SetOnActivate(fn TNotifyEvent) {
	if m.activatePtr != 0 {
		RemoveEventElement(m.activatePtr)
	}
	m.activatePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1712, m.Instance(), m.activatePtr)
}

func (m *TCustomForm) SetOnClose(fn TCloseEvent) {
	if m.closePtr != 0 {
		RemoveEventElement(m.closePtr)
	}
	m.closePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1713, m.Instance(), m.closePtr)
}

func (m *TCustomForm) SetOnCloseQuery(fn TCloseQueryEvent) {
	if m.closeQueryPtr != 0 {
		RemoveEventElement(m.closeQueryPtr)
	}
	m.closeQueryPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1714, m.Instance(), m.closeQueryPtr)
}

func (m *TCustomForm) SetOnCreate(fn TNotifyEvent) {
	if m.createPtr != 0 {
		RemoveEventElement(m.createPtr)
	}
	m.createPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1715, m.Instance(), m.createPtr)
}

func (m *TCustomForm) SetOnDeactivate(fn TNotifyEvent) {
	if m.deactivatePtr != 0 {
		RemoveEventElement(m.deactivatePtr)
	}
	m.deactivatePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1716, m.Instance(), m.deactivatePtr)
}

func (m *TCustomForm) SetOnDestroy(fn TNotifyEvent) {
	if m.destroyPtr != 0 {
		RemoveEventElement(m.destroyPtr)
	}
	m.destroyPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1717, m.Instance(), m.destroyPtr)
}

func (m *TCustomForm) SetOnDropFiles(fn TDropFilesEvent) {
	if m.dropFilesPtr != 0 {
		RemoveEventElement(m.dropFilesPtr)
	}
	m.dropFilesPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1718, m.Instance(), m.dropFilesPtr)
}

func (m *TCustomForm) SetOnHelp(fn THelpEvent) {
	if m.helpPtr != 0 {
		RemoveEventElement(m.helpPtr)
	}
	m.helpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1719, m.Instance(), m.helpPtr)
}

func (m *TCustomForm) SetOnHide(fn TNotifyEvent) {
	if m.hidePtr != 0 {
		RemoveEventElement(m.hidePtr)
	}
	m.hidePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1720, m.Instance(), m.hidePtr)
}

func (m *TCustomForm) SetOnShortcut(fn TShortCutEvent) {
	if m.shortcutPtr != 0 {
		RemoveEventElement(m.shortcutPtr)
	}
	m.shortcutPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1721, m.Instance(), m.shortcutPtr)
}

func (m *TCustomForm) SetOnShow(fn TNotifyEvent) {
	if m.showPtr != 0 {
		RemoveEventElement(m.showPtr)
	}
	m.showPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1722, m.Instance(), m.showPtr)
}

func (m *TCustomForm) SetOnShowModalFinished(fn TModalDialogFinished) {
	if m.showModalFinishedPtr != 0 {
		RemoveEventElement(m.showModalFinishedPtr)
	}
	m.showModalFinishedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1723, m.Instance(), m.showModalFinishedPtr)
}

func (m *TCustomForm) SetOnWindowStateChange(fn TNotifyEvent) {
	if m.windowStateChangePtr != 0 {
		RemoveEventElement(m.windowStateChangePtr)
	}
	m.windowStateChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1724, m.Instance(), m.windowStateChangePtr)
}
