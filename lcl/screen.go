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

// IScreen Parent: ILCLComponent
type IScreen interface {
	ILCLComponent
	ActiveControl() IWinControl                                                  // property
	ActiveCustomForm() ICustomForm                                               // property
	ActiveForm() IForm                                                           // property
	Cursor() TCursor                                                             // property
	SetCursor(AValue TCursor)                                                    // property
	RealCursor() TCursor                                                         // property
	Cursors(Index int32) HCURSOR                                                 // property
	SetCursors(Index int32, AValue HCURSOR)                                      // property
	CustomFormCount() int32                                                      // property
	CustomForms(Index int32) ICustomForm                                         // property
	CustomFormZOrderCount() int32                                                // property
	CustomFormsZOrdered(Index int32) ICustomForm                                 // property
	DesktopLeft() int32                                                          // property
	DesktopTop() int32                                                           // property
	DesktopHeight() int32                                                        // property
	DesktopWidth() int32                                                         // property
	DesktopRect() (resultRect TRect)                                             // property
	FocusedForm() ICustomForm                                                    // property
	FormCount() int32                                                            // property
	Forms(Index int32) IForm                                                     // property
	DataModuleCount() int32                                                      // property
	DataModules(Index int32) IDataModule                                         // property
	HintFont() IFont                                                             // property
	SetHintFont(AValue IFont)                                                    // property
	IconFont() IFont                                                             // property
	SetIconFont(AValue IFont)                                                    // property
	MenuFont() IFont                                                             // property
	SetMenuFont(AValue IFont)                                                    // property
	SystemFont() IFont                                                           // property
	SetSystemFont(AValue IFont)                                                  // property
	Fonts() IStrings                                                             // property
	Height() int32                                                               // property
	MonitorCount() int32                                                         // property
	Monitors(Index int32) IMonitor                                               // property
	PixelsPerInch() int32                                                        // property
	PrimaryMonitor() IMonitor                                                    // property
	Width() int32                                                                // property
	WorkAreaRect() (resultRect TRect)                                            // property
	WorkAreaHeight() int32                                                       // property
	WorkAreaLeft() int32                                                         // property
	WorkAreaTop() int32                                                          // property
	WorkAreaWidth() int32                                                        // property
	CustomFormIndex(AForm ICustomForm) int32                                     // function
	FormIndex(AForm IForm) int32                                                 // function
	CustomFormZIndex(AForm ICustomForm) int32                                    // function
	GetCurrentModalForm() ICustomForm                                            // function
	GetCurrentModalFormZIndex() int32                                            // function
	CustomFormBelongsToActiveGroup(AForm ICustomForm) bool                       // function
	FindNonDesignerForm(FormName string) ICustomForm                             // function
	FindForm(FormName string) ICustomForm                                        // function
	FindNonDesignerDataModule(DataModuleName string) IDataModule                 // function
	FindDataModule(DataModuleName string) IDataModule                            // function
	DisableForms(SkipForm ICustomForm, DisabledList IList) IList                 // function
	MonitorFromPoint(Point *TPoint, MonitorDefault TMonitorDefaultTo) IMonitor   // function
	MonitorFromRect(Rect *TRect, MonitorDefault TMonitorDefaultTo) IMonitor      // function
	MonitorFromWindow(Handle THandle, MonitorDefault TMonitorDefaultTo) IMonitor // function
	MoveFormToFocusFront(ACustomForm ICustomForm)                                // procedure
	MoveFormToZFront(ACustomForm ICustomForm)                                    // procedure
	NewFormWasCreated(AForm ICustomForm)                                         // procedure
	UpdateMonitors()                                                             // procedure
	UpdateScreen()                                                               // procedure
	EnableForms(AFormList *IList)                                                // procedure
	BeginTempCursor(aCursor TCursor)                                             // procedure
	EndTempCursor(aCursor TCursor)                                               // procedure
	BeginWaitCursor()                                                            // procedure
	EndWaitCursor()                                                              // procedure
	BeginScreenCursor()                                                          // procedure
	EndScreenCursor()                                                            // procedure
	SetOnActiveControlChange(fn TNotifyEvent)                                    // property event
	SetOnActiveFormChange(fn TNotifyEvent)                                       // property event
}

// TScreen Parent: TLCLComponent
type TScreen struct {
	TLCLComponent
	activeControlChangePtr uintptr
	activeFormChangePtr    uintptr
}

func NewScreen(AOwner IComponent) IScreen {
	r1 := LCL().SysCallN(4214, GetObjectUintptr(AOwner))
	return AsScreen(r1)
}

func (m *TScreen) ActiveControl() IWinControl {
	r1 := LCL().SysCallN(4207, m.Instance())
	return AsWinControl(r1)
}

func (m *TScreen) ActiveCustomForm() ICustomForm {
	r1 := LCL().SysCallN(4208, m.Instance())
	return AsCustomForm(r1)
}

func (m *TScreen) ActiveForm() IForm {
	r1 := LCL().SysCallN(4209, m.Instance())
	return AsForm(r1)
}

func (m *TScreen) Cursor() TCursor {
	r1 := LCL().SysCallN(4215, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TScreen) SetCursor(AValue TCursor) {
	LCL().SysCallN(4215, 1, m.Instance(), uintptr(AValue))
}

func (m *TScreen) RealCursor() TCursor {
	r1 := LCL().SysCallN(4261, m.Instance())
	return TCursor(r1)
}

func (m *TScreen) Cursors(Index int32) HCURSOR {
	r1 := LCL().SysCallN(4216, 0, m.Instance(), uintptr(Index))
	return HCURSOR(r1)
}

func (m *TScreen) SetCursors(Index int32, AValue HCURSOR) {
	LCL().SysCallN(4216, 1, m.Instance(), uintptr(Index), uintptr(AValue))
}

func (m *TScreen) CustomFormCount() int32 {
	r1 := LCL().SysCallN(4218, m.Instance())
	return int32(r1)
}

func (m *TScreen) CustomForms(Index int32) ICustomForm {
	r1 := LCL().SysCallN(4222, m.Instance(), uintptr(Index))
	return AsCustomForm(r1)
}

func (m *TScreen) CustomFormZOrderCount() int32 {
	r1 := LCL().SysCallN(4221, m.Instance())
	return int32(r1)
}

func (m *TScreen) CustomFormsZOrdered(Index int32) ICustomForm {
	r1 := LCL().SysCallN(4223, m.Instance(), uintptr(Index))
	return AsCustomForm(r1)
}

func (m *TScreen) DesktopLeft() int32 {
	r1 := LCL().SysCallN(4227, m.Instance())
	return int32(r1)
}

func (m *TScreen) DesktopTop() int32 {
	r1 := LCL().SysCallN(4229, m.Instance())
	return int32(r1)
}

func (m *TScreen) DesktopHeight() int32 {
	r1 := LCL().SysCallN(4226, m.Instance())
	return int32(r1)
}

func (m *TScreen) DesktopWidth() int32 {
	r1 := LCL().SysCallN(4230, m.Instance())
	return int32(r1)
}

func (m *TScreen) DesktopRect() (resultRect TRect) {
	LCL().SysCallN(4228, m.Instance(), uintptr(unsafe.Pointer(&resultRect)))
	return
}

func (m *TScreen) FocusedForm() ICustomForm {
	r1 := LCL().SysCallN(4240, m.Instance())
	return AsCustomForm(r1)
}

func (m *TScreen) FormCount() int32 {
	r1 := LCL().SysCallN(4242, m.Instance())
	return int32(r1)
}

func (m *TScreen) Forms(Index int32) IForm {
	r1 := LCL().SysCallN(4244, m.Instance(), uintptr(Index))
	return AsForm(r1)
}

func (m *TScreen) DataModuleCount() int32 {
	r1 := LCL().SysCallN(4224, m.Instance())
	return int32(r1)
}

func (m *TScreen) DataModules(Index int32) IDataModule {
	r1 := LCL().SysCallN(4225, m.Instance(), uintptr(Index))
	return AsDataModule(r1)
}

func (m *TScreen) HintFont() IFont {
	r1 := LCL().SysCallN(4248, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TScreen) SetHintFont(AValue IFont) {
	LCL().SysCallN(4248, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TScreen) IconFont() IFont {
	r1 := LCL().SysCallN(4249, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TScreen) SetIconFont(AValue IFont) {
	LCL().SysCallN(4249, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TScreen) MenuFont() IFont {
	r1 := LCL().SysCallN(4250, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TScreen) SetMenuFont(AValue IFont) {
	LCL().SysCallN(4250, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TScreen) SystemFont() IFont {
	r1 := LCL().SysCallN(4264, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TScreen) SetSystemFont(AValue IFont) {
	LCL().SysCallN(4264, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TScreen) Fonts() IStrings {
	r1 := LCL().SysCallN(4241, m.Instance())
	return AsStrings(r1)
}

func (m *TScreen) Height() int32 {
	r1 := LCL().SysCallN(4247, m.Instance())
	return int32(r1)
}

func (m *TScreen) MonitorCount() int32 {
	r1 := LCL().SysCallN(4251, m.Instance())
	return int32(r1)
}

func (m *TScreen) Monitors(Index int32) IMonitor {
	r1 := LCL().SysCallN(4255, m.Instance(), uintptr(Index))
	return AsMonitor(r1)
}

func (m *TScreen) PixelsPerInch() int32 {
	r1 := LCL().SysCallN(4259, m.Instance())
	return int32(r1)
}

func (m *TScreen) PrimaryMonitor() IMonitor {
	r1 := LCL().SysCallN(4260, m.Instance())
	return AsMonitor(r1)
}

func (m *TScreen) Width() int32 {
	r1 := LCL().SysCallN(4267, m.Instance())
	return int32(r1)
}

func (m *TScreen) WorkAreaRect() (resultRect TRect) {
	LCL().SysCallN(4270, m.Instance(), uintptr(unsafe.Pointer(&resultRect)))
	return
}

func (m *TScreen) WorkAreaHeight() int32 {
	r1 := LCL().SysCallN(4268, m.Instance())
	return int32(r1)
}

func (m *TScreen) WorkAreaLeft() int32 {
	r1 := LCL().SysCallN(4269, m.Instance())
	return int32(r1)
}

func (m *TScreen) WorkAreaTop() int32 {
	r1 := LCL().SysCallN(4271, m.Instance())
	return int32(r1)
}

func (m *TScreen) WorkAreaWidth() int32 {
	r1 := LCL().SysCallN(4272, m.Instance())
	return int32(r1)
}

func (m *TScreen) CustomFormIndex(AForm ICustomForm) int32 {
	r1 := LCL().SysCallN(4219, m.Instance(), GetObjectUintptr(AForm))
	return int32(r1)
}

func (m *TScreen) FormIndex(AForm IForm) int32 {
	r1 := LCL().SysCallN(4243, m.Instance(), GetObjectUintptr(AForm))
	return int32(r1)
}

func (m *TScreen) CustomFormZIndex(AForm ICustomForm) int32 {
	r1 := LCL().SysCallN(4220, m.Instance(), GetObjectUintptr(AForm))
	return int32(r1)
}

func (m *TScreen) GetCurrentModalForm() ICustomForm {
	r1 := LCL().SysCallN(4245, m.Instance())
	return AsCustomForm(r1)
}

func (m *TScreen) GetCurrentModalFormZIndex() int32 {
	r1 := LCL().SysCallN(4246, m.Instance())
	return int32(r1)
}

func (m *TScreen) CustomFormBelongsToActiveGroup(AForm ICustomForm) bool {
	r1 := LCL().SysCallN(4217, m.Instance(), GetObjectUintptr(AForm))
	return GoBool(r1)
}

func (m *TScreen) FindNonDesignerForm(FormName string) ICustomForm {
	r1 := LCL().SysCallN(4239, m.Instance(), PascalStr(FormName))
	return AsCustomForm(r1)
}

func (m *TScreen) FindForm(FormName string) ICustomForm {
	r1 := LCL().SysCallN(4237, m.Instance(), PascalStr(FormName))
	return AsCustomForm(r1)
}

func (m *TScreen) FindNonDesignerDataModule(DataModuleName string) IDataModule {
	r1 := LCL().SysCallN(4238, m.Instance(), PascalStr(DataModuleName))
	return AsDataModule(r1)
}

func (m *TScreen) FindDataModule(DataModuleName string) IDataModule {
	r1 := LCL().SysCallN(4236, m.Instance(), PascalStr(DataModuleName))
	return AsDataModule(r1)
}

func (m *TScreen) DisableForms(SkipForm ICustomForm, DisabledList IList) IList {
	r1 := LCL().SysCallN(4231, m.Instance(), GetObjectUintptr(SkipForm), GetObjectUintptr(DisabledList))
	return AsList(r1)
}

func (m *TScreen) MonitorFromPoint(Point *TPoint, MonitorDefault TMonitorDefaultTo) IMonitor {
	r1 := LCL().SysCallN(4252, m.Instance(), uintptr(unsafe.Pointer(Point)), uintptr(MonitorDefault))
	return AsMonitor(r1)
}

func (m *TScreen) MonitorFromRect(Rect *TRect, MonitorDefault TMonitorDefaultTo) IMonitor {
	r1 := LCL().SysCallN(4253, m.Instance(), uintptr(unsafe.Pointer(Rect)), uintptr(MonitorDefault))
	return AsMonitor(r1)
}

func (m *TScreen) MonitorFromWindow(Handle THandle, MonitorDefault TMonitorDefaultTo) IMonitor {
	r1 := LCL().SysCallN(4254, m.Instance(), uintptr(Handle), uintptr(MonitorDefault))
	return AsMonitor(r1)
}

func ScreenClass() TClass {
	ret := LCL().SysCallN(4213)
	return TClass(ret)
}

func (m *TScreen) MoveFormToFocusFront(ACustomForm ICustomForm) {
	LCL().SysCallN(4256, m.Instance(), GetObjectUintptr(ACustomForm))
}

func (m *TScreen) MoveFormToZFront(ACustomForm ICustomForm) {
	LCL().SysCallN(4257, m.Instance(), GetObjectUintptr(ACustomForm))
}

func (m *TScreen) NewFormWasCreated(AForm ICustomForm) {
	LCL().SysCallN(4258, m.Instance(), GetObjectUintptr(AForm))
}

func (m *TScreen) UpdateMonitors() {
	LCL().SysCallN(4265, m.Instance())
}

func (m *TScreen) UpdateScreen() {
	LCL().SysCallN(4266, m.Instance())
}

func (m *TScreen) EnableForms(AFormList *IList) {
	var result0 uintptr
	LCL().SysCallN(4232, m.Instance(), uintptr(unsafe.Pointer(&result0)))
	*AFormList = AsList(result0)
}

func (m *TScreen) BeginTempCursor(aCursor TCursor) {
	LCL().SysCallN(4211, m.Instance(), uintptr(aCursor))
}

func (m *TScreen) EndTempCursor(aCursor TCursor) {
	LCL().SysCallN(4234, m.Instance(), uintptr(aCursor))
}

func (m *TScreen) BeginWaitCursor() {
	LCL().SysCallN(4212, m.Instance())
}

func (m *TScreen) EndWaitCursor() {
	LCL().SysCallN(4235, m.Instance())
}

func (m *TScreen) BeginScreenCursor() {
	LCL().SysCallN(4210, m.Instance())
}

func (m *TScreen) EndScreenCursor() {
	LCL().SysCallN(4233, m.Instance())
}

func (m *TScreen) SetOnActiveControlChange(fn TNotifyEvent) {
	if m.activeControlChangePtr != 0 {
		RemoveEventElement(m.activeControlChangePtr)
	}
	m.activeControlChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4262, m.Instance(), m.activeControlChangePtr)
}

func (m *TScreen) SetOnActiveFormChange(fn TNotifyEvent) {
	if m.activeFormChangePtr != 0 {
		RemoveEventElement(m.activeFormChangePtr)
	}
	m.activeFormChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4263, m.Instance(), m.activeFormChangePtr)
}
