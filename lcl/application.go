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

// IApplication Parent: ICustomApplication
type IApplication interface {
	ICustomApplication
	CreateForm(fields ...interface{}) IForm
	Initialize()
	SetRunLoopReceived(proc uintptr)
	SetIconResId(id int)
	Active() bool                                                              // property
	ApplicationType() TApplicationType                                         // property
	SetApplicationType(AValue TApplicationType)                                // property
	BidiMode() TBiDiMode                                                       // property
	SetBidiMode(AValue TBiDiMode)                                              // property
	CaptureExceptions() bool                                                   // property
	SetCaptureExceptions(AValue bool)                                          // property
	DoubleBuffered() TApplicationDoubleBuffered                                // property
	SetDoubleBuffered(AValue TApplicationDoubleBuffered)                       // property
	ExtendedKeysSupport() bool                                                 // property
	SetExtendedKeysSupport(AValue bool)                                        // property
	ExceptionDialog() TApplicationExceptionDlg                                 // property
	SetExceptionDialog(AValue TApplicationExceptionDlg)                        // property
	FindGlobalComponentEnabled() bool                                          // property
	SetFindGlobalComponentEnabled(AValue bool)                                 // property
	Flags() TApplicationFlags                                                  // property
	SetFlags(AValue TApplicationFlags)                                         // property
	Handle() THandle                                                           // property
	SetHandle(AValue THandle)                                                  // property
	Hint() string                                                              // property
	SetHint(AValue string)                                                     // property
	HintColor() TColor                                                         // property
	SetHintColor(AValue TColor)                                                // property
	HintHidePause() int32                                                      // property
	SetHintHidePause(AValue int32)                                             // property
	HintHidePausePerChar() int32                                               // property
	SetHintHidePausePerChar(AValue int32)                                      // property
	HintPause() int32                                                          // property
	SetHintPause(AValue int32)                                                 // property
	HintShortCuts() bool                                                       // property
	SetHintShortCuts(AValue bool)                                              // property
	HintShortPause() int32                                                     // property
	SetHintShortPause(AValue int32)                                            // property
	Icon() IIcon                                                               // property
	SetIcon(AValue IIcon)                                                      // property
	LayoutAdjustmentPolicy() TLayoutAdjustmentPolicy                           // property
	SetLayoutAdjustmentPolicy(AValue TLayoutAdjustmentPolicy)                  // property
	Navigation() TApplicationNavigationOptions                                 // property
	SetNavigation(AValue TApplicationNavigationOptions)                        // property
	MainForm() IForm                                                           // property
	MainFormHandle() HWND                                                      // property
	MainFormOnTaskBar() bool                                                   // property
	SetMainFormOnTaskBar(AValue bool)                                          // property
	ModalLevel() int32                                                         // property
	MoveFormFocusToChildren() bool                                             // property
	SetMoveFormFocusToChildren(AValue bool)                                    // property
	MouseControl() IControl                                                    // property
	TaskBarBehavior() TTaskBarBehavior                                         // property
	SetTaskBarBehavior(AValue TTaskBarBehavior)                                // property
	UpdateFormatSettings() bool                                                // property
	SetUpdateFormatSettings(AValue bool)                                       // property
	ShowButtonGlyphs() TApplicationShowGlyphs                                  // property
	SetShowButtonGlyphs(AValue TApplicationShowGlyphs)                         // property
	ShowMenuGlyphs() TApplicationShowGlyphs                                    // property
	SetShowMenuGlyphs(AValue TApplicationShowGlyphs)                           // property
	ShowHint() bool                                                            // property
	SetShowHint(AValue bool)                                                   // property
	ShowMainForm() bool                                                        // property
	SetShowMainForm(AValue bool)                                               // property
	Scaled() bool                                                              // property
	SetScaled(AValue bool)                                                     // property
	GetControlAtMouse() IControl                                               // function
	GetControlAtPos(P *TPoint) IControl                                        // function
	BigIconHandle() HICON                                                      // function
	SmallIconHandle() HICON                                                    // function
	HelpCommand(Command Word, Data uint32) bool                                // function
	HelpContext(Context THelpContext) bool                                     // function
	HelpKeyword(Keyword string) bool                                           // function
	HelpShowTableOfContents() bool                                             // function
	IsWaiting() bool                                                           // function
	MessageBox(Text, Caption string, Flags int32) int32                        // function
	IsShortcut(Message *TLMKey) bool                                           // function
	IsRightToLeft() bool                                                       // function
	IsRTLLang(ALang string) bool                                               // function
	Direction(ALang string) TBiDiMode                                          // function
	ActivateHint(CursorPos *TPoint, CheckHintControlChange bool)               // procedure
	ControlDestroyed(AControl IControl)                                        // procedure
	BringToFront()                                                             // procedure
	UpdateMainForm(AForm IForm)                                                // procedure
	RemoveAsyncCalls(AnObject IObject)                                         // procedure
	ReleaseComponent(AComponent IComponent)                                    // procedure
	HandleMessage()                                                            // procedure
	ShowHelpForObject(Sender IObject)                                          // procedure
	RemoveStayOnTop(ASystemTopAlso bool)                                       // procedure
	RestoreStayOnTop(ASystemTopAlso bool)                                      // procedure
	CancelHint()                                                               // procedure
	HideHint()                                                                 // procedure
	HintMouseMessage(Control IControl, AMessage *TLMessage)                    // procedure
	Minimize()                                                                 // procedure
	ModalStarted()                                                             // procedure
	ModalFinished()                                                            // procedure
	Restore()                                                                  // procedure
	Notification(AComponent IComponent, Operation TOperation)                  // procedure
	ProcessMessages()                                                          // procedure
	Idle(Wait bool)                                                            // procedure
	DisableIdleHandler()                                                       // procedure
	EnableIdleHandler()                                                        // procedure
	NotifyUserInputHandler(Sender IObject, Msg uint32)                         // procedure
	NotifyKeyDownBeforeHandler(Sender IObject, Key *Word, Shift TShiftState)   // procedure
	NotifyKeyDownHandler(Sender IObject, Key *Word, Shift TShiftState)         // procedure
	ControlKeyDown(Sender IObject, Key *Word, Shift TShiftState)               // procedure
	ControlKeyUp(Sender IObject, Key *Word, Shift TShiftState)                 // procedure
	RemoveAllHandlersOfObject(AnObject IObject)                                // procedure
	DoBeforeMouseMessage(CurMouseControl IControl)                             // procedure
	IntfEndSession()                                                           // procedure
	IntfAppActivate(Async bool)                                                // procedure
	IntfAppDeactivate(Async bool)                                              // procedure
	IntfAppMinimize()                                                          // procedure
	IntfAppRestore()                                                           // procedure
	IntfSettingsChange()                                                       // procedure
	IntfThemeOptionChange(AThemeServices IThemeServices, AOption TThemeOption) // procedure
	DoArrowKey(AControl IWinControl, Key *Word, Shift TShiftState)             // procedure
	DoTabKey(AControl IWinControl, Key *Word, Shift TShiftState)               // procedure
	DoEscapeKey(AControl IWinControl, Key *Word, Shift TShiftState)            // procedure
	DoReturnKey(AControl IWinControl, Key *Word, Shift TShiftState)            // procedure
	SetOnActionExecute(fn TActionEvent)                                        // property event
	SetOnActionUpdate(fn TActionEvent)                                         // property event
	SetOnActivate(fn TNotifyEvent)                                             // property event
	SetOnDeactivate(fn TNotifyEvent)                                           // property event
	SetOnGetMainFormHandle(fn TGetHandleEvent)                                 // property event
	SetOnIdleEnd(fn TNotifyEvent)                                              // property event
	SetOnEndSession(fn TNotifyEvent)                                           // property event
	SetOnMinimize(fn TNotifyEvent)                                             // property event
	SetOnModalBegin(fn TNotifyEvent)                                           // property event
	SetOnModalEnd(fn TNotifyEvent)                                             // property event
	SetOnRestore(fn TNotifyEvent)                                              // property event
	SetOnDropFiles(fn TDropFilesEvent)                                         // property event
	SetOnHelp(fn THelpEvent)                                                   // property event
	SetOnHint(fn TNotifyEvent)                                                 // property event
	SetOnShortcut(fn TShortCutEvent)                                           // property event
	SetOnShowHint(fn TShowHintEvent)                                           // property event
	SetOnDestroy(fn TNotifyEvent)                                              // property event
	SetOnCircularException(fn TExceptionEvent)                                 // property event
}

// TApplication Parent: TCustomApplication
type TApplication struct {
	TCustomApplication
	actionExecutePtr     uintptr
	actionUpdatePtr      uintptr
	activatePtr          uintptr
	deactivatePtr        uintptr
	getMainFormHandlePtr uintptr
	idleEndPtr           uintptr
	endSessionPtr        uintptr
	minimizePtr          uintptr
	modalBeginPtr        uintptr
	modalEndPtr          uintptr
	restorePtr           uintptr
	dropFilesPtr         uintptr
	helpPtr              uintptr
	hintPtr              uintptr
	shortcutPtr          uintptr
	showHintPtr          uintptr
	destroyPtr           uintptr
	circularExceptionPtr uintptr
}

func NewApplication(AOwner IComponent) IApplication {
	r1 := LCL().SysCallN(109, GetObjectUintptr(AOwner))
	return AsApplication(r1)
}

func (m *TApplication) Active() bool {
	r1 := LCL().SysCallN(98, m.Instance())
	return GoBool(r1)
}

func (m *TApplication) ApplicationType() TApplicationType {
	r1 := LCL().SysCallN(99, 0, m.Instance(), 0)
	return TApplicationType(r1)
}

func (m *TApplication) SetApplicationType(AValue TApplicationType) {
	LCL().SysCallN(99, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) BidiMode() TBiDiMode {
	r1 := LCL().SysCallN(100, 0, m.Instance(), 0)
	return TBiDiMode(r1)
}

func (m *TApplication) SetBidiMode(AValue TBiDiMode) {
	LCL().SysCallN(100, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) CaptureExceptions() bool {
	r1 := LCL().SysCallN(104, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetCaptureExceptions(AValue bool) {
	LCL().SysCallN(104, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) DoubleBuffered() TApplicationDoubleBuffered {
	r1 := LCL().SysCallN(117, 0, m.Instance(), 0)
	return TApplicationDoubleBuffered(r1)
}

func (m *TApplication) SetDoubleBuffered(AValue TApplicationDoubleBuffered) {
	LCL().SysCallN(117, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) ExtendedKeysSupport() bool {
	r1 := LCL().SysCallN(120, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetExtendedKeysSupport(AValue bool) {
	LCL().SysCallN(120, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) ExceptionDialog() TApplicationExceptionDlg {
	r1 := LCL().SysCallN(119, 0, m.Instance(), 0)
	return TApplicationExceptionDlg(r1)
}

func (m *TApplication) SetExceptionDialog(AValue TApplicationExceptionDlg) {
	LCL().SysCallN(119, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) FindGlobalComponentEnabled() bool {
	r1 := LCL().SysCallN(121, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetFindGlobalComponentEnabled(AValue bool) {
	LCL().SysCallN(121, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) Flags() TApplicationFlags {
	r1 := LCL().SysCallN(122, 0, m.Instance(), 0)
	return TApplicationFlags(r1)
}

func (m *TApplication) SetFlags(AValue TApplicationFlags) {
	LCL().SysCallN(122, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) Handle() THandle {
	r1 := LCL().SysCallN(125, 0, m.Instance(), 0)
	return THandle(r1)
}

func (m *TApplication) SetHandle(AValue THandle) {
	LCL().SysCallN(125, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) Hint() string {
	r1 := LCL().SysCallN(132, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TApplication) SetHint(AValue string) {
	LCL().SysCallN(132, 1, m.Instance(), PascalStr(AValue))
}

func (m *TApplication) HintColor() TColor {
	r1 := LCL().SysCallN(133, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TApplication) SetHintColor(AValue TColor) {
	LCL().SysCallN(133, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) HintHidePause() int32 {
	r1 := LCL().SysCallN(134, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TApplication) SetHintHidePause(AValue int32) {
	LCL().SysCallN(134, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) HintHidePausePerChar() int32 {
	r1 := LCL().SysCallN(135, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TApplication) SetHintHidePausePerChar(AValue int32) {
	LCL().SysCallN(135, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) HintPause() int32 {
	r1 := LCL().SysCallN(137, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TApplication) SetHintPause(AValue int32) {
	LCL().SysCallN(137, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) HintShortCuts() bool {
	r1 := LCL().SysCallN(138, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetHintShortCuts(AValue bool) {
	LCL().SysCallN(138, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) HintShortPause() int32 {
	r1 := LCL().SysCallN(139, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TApplication) SetHintShortPause(AValue int32) {
	LCL().SysCallN(139, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) Icon() IIcon {
	r1 := LCL().SysCallN(140, 0, m.Instance(), 0)
	return AsIcon(r1)
}

func (m *TApplication) SetIcon(AValue IIcon) {
	LCL().SysCallN(140, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TApplication) LayoutAdjustmentPolicy() TLayoutAdjustmentPolicy {
	r1 := LCL().SysCallN(153, 0, m.Instance(), 0)
	return TLayoutAdjustmentPolicy(r1)
}

func (m *TApplication) SetLayoutAdjustmentPolicy(AValue TLayoutAdjustmentPolicy) {
	LCL().SysCallN(153, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) Navigation() TApplicationNavigationOptions {
	r1 := LCL().SysCallN(164, 0, m.Instance(), 0)
	return TApplicationNavigationOptions(r1)
}

func (m *TApplication) SetNavigation(AValue TApplicationNavigationOptions) {
	LCL().SysCallN(164, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) MainForm() IForm {
	r1 := LCL().SysCallN(154, m.Instance())
	return AsForm(r1)
}

func (m *TApplication) MainFormHandle() HWND {
	r1 := LCL().SysCallN(155, m.Instance())
	return HWND(r1)
}

func (m *TApplication) MainFormOnTaskBar() bool {
	r1 := LCL().SysCallN(156, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetMainFormOnTaskBar(AValue bool) {
	LCL().SysCallN(156, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) ModalLevel() int32 {
	r1 := LCL().SysCallN(160, m.Instance())
	return int32(r1)
}

func (m *TApplication) MoveFormFocusToChildren() bool {
	r1 := LCL().SysCallN(163, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetMoveFormFocusToChildren(AValue bool) {
	LCL().SysCallN(163, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) MouseControl() IControl {
	r1 := LCL().SysCallN(162, m.Instance())
	return AsControl(r1)
}

func (m *TApplication) TaskBarBehavior() TTaskBarBehavior {
	r1 := LCL().SysCallN(201, 0, m.Instance(), 0)
	return TTaskBarBehavior(r1)
}

func (m *TApplication) SetTaskBarBehavior(AValue TTaskBarBehavior) {
	LCL().SysCallN(201, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) UpdateFormatSettings() bool {
	r1 := LCL().SysCallN(202, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetUpdateFormatSettings(AValue bool) {
	LCL().SysCallN(202, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) ShowButtonGlyphs() TApplicationShowGlyphs {
	r1 := LCL().SysCallN(195, 0, m.Instance(), 0)
	return TApplicationShowGlyphs(r1)
}

func (m *TApplication) SetShowButtonGlyphs(AValue TApplicationShowGlyphs) {
	LCL().SysCallN(195, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) ShowMenuGlyphs() TApplicationShowGlyphs {
	r1 := LCL().SysCallN(199, 0, m.Instance(), 0)
	return TApplicationShowGlyphs(r1)
}

func (m *TApplication) SetShowMenuGlyphs(AValue TApplicationShowGlyphs) {
	LCL().SysCallN(199, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) ShowHint() bool {
	r1 := LCL().SysCallN(197, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetShowHint(AValue bool) {
	LCL().SysCallN(197, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) ShowMainForm() bool {
	r1 := LCL().SysCallN(198, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetShowMainForm(AValue bool) {
	LCL().SysCallN(198, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) Scaled() bool {
	r1 := LCL().SysCallN(176, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetScaled(AValue bool) {
	LCL().SysCallN(176, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) GetControlAtMouse() IControl {
	r1 := LCL().SysCallN(123, m.Instance())
	return AsControl(r1)
}

func (m *TApplication) GetControlAtPos(P *TPoint) IControl {
	r1 := LCL().SysCallN(124, m.Instance(), uintptr(unsafe.Pointer(P)))
	return AsControl(r1)
}

func (m *TApplication) BigIconHandle() HICON {
	r1 := LCL().SysCallN(101, m.Instance())
	return HICON(r1)
}

func (m *TApplication) SmallIconHandle() HICON {
	r1 := LCL().SysCallN(200, m.Instance())
	return HICON(r1)
}

func (m *TApplication) HelpCommand(Command Word, Data uint32) bool {
	r1 := LCL().SysCallN(127, m.Instance(), uintptr(Command), uintptr(Data))
	return GoBool(r1)
}

func (m *TApplication) HelpContext(Context THelpContext) bool {
	r1 := LCL().SysCallN(128, m.Instance(), uintptr(Context))
	return GoBool(r1)
}

func (m *TApplication) HelpKeyword(Keyword string) bool {
	r1 := LCL().SysCallN(129, m.Instance(), PascalStr(Keyword))
	return GoBool(r1)
}

func (m *TApplication) HelpShowTableOfContents() bool {
	r1 := LCL().SysCallN(130, m.Instance())
	return GoBool(r1)
}

func (m *TApplication) IsWaiting() bool {
	r1 := LCL().SysCallN(152, m.Instance())
	return GoBool(r1)
}

func (m *TApplication) MessageBox(Text, Caption string, Flags int32) int32 {
	r1 := LCL().SysCallN(157, m.Instance(), PascalStr(Text), PascalStr(Caption), uintptr(Flags))
	return int32(r1)
}

func (m *TApplication) IsShortcut(Message *TLMKey) bool {
	var result0 uintptr
	r1 := LCL().SysCallN(151, m.Instance(), uintptr(unsafe.Pointer(&result0)))
	*Message = *(*TLMKey)(getPointer(result0))
	return GoBool(r1)
}

func (m *TApplication) IsRightToLeft() bool {
	r1 := LCL().SysCallN(150, m.Instance())
	return GoBool(r1)
}

func (m *TApplication) IsRTLLang(ALang string) bool {
	r1 := LCL().SysCallN(149, m.Instance(), PascalStr(ALang))
	return GoBool(r1)
}

func (m *TApplication) Direction(ALang string) TBiDiMode {
	r1 := LCL().SysCallN(110, m.Instance(), PascalStr(ALang))
	return TBiDiMode(r1)
}

func ApplicationClass() TClass {
	ret := LCL().SysCallN(105)
	return TClass(ret)
}

func (m *TApplication) ActivateHint(CursorPos *TPoint, CheckHintControlChange bool) {
	LCL().SysCallN(97, m.Instance(), uintptr(unsafe.Pointer(CursorPos)), PascalBool(CheckHintControlChange))
}

func (m *TApplication) ControlDestroyed(AControl IControl) {
	LCL().SysCallN(106, m.Instance(), GetObjectUintptr(AControl))
}

func (m *TApplication) BringToFront() {
	LCL().SysCallN(102, m.Instance())
}

func (m *TApplication) UpdateMainForm(AForm IForm) {
	LCL().SysCallN(203, m.Instance(), GetObjectUintptr(AForm))
}

func (m *TApplication) RemoveAsyncCalls(AnObject IObject) {
	LCL().SysCallN(172, m.Instance(), GetObjectUintptr(AnObject))
}

func (m *TApplication) ReleaseComponent(AComponent IComponent) {
	LCL().SysCallN(170, m.Instance(), GetObjectUintptr(AComponent))
}

func (m *TApplication) HandleMessage() {
	LCL().SysCallN(126, m.Instance())
}

func (m *TApplication) ShowHelpForObject(Sender IObject) {
	LCL().SysCallN(196, m.Instance(), GetObjectUintptr(Sender))
}

func (m *TApplication) RemoveStayOnTop(ASystemTopAlso bool) {
	LCL().SysCallN(173, m.Instance(), PascalBool(ASystemTopAlso))
}

func (m *TApplication) RestoreStayOnTop(ASystemTopAlso bool) {
	LCL().SysCallN(175, m.Instance(), PascalBool(ASystemTopAlso))
}

func (m *TApplication) CancelHint() {
	LCL().SysCallN(103, m.Instance())
}

func (m *TApplication) HideHint() {
	LCL().SysCallN(131, m.Instance())
}

func (m *TApplication) HintMouseMessage(Control IControl, AMessage *TLMessage) {
	var result1 uintptr
	LCL().SysCallN(136, m.Instance(), GetObjectUintptr(Control), uintptr(unsafe.Pointer(&result1)))
	*AMessage = *(*TLMessage)(getPointer(result1))
}

func (m *TApplication) Minimize() {
	LCL().SysCallN(158, m.Instance())
}

func (m *TApplication) ModalStarted() {
	LCL().SysCallN(161, m.Instance())
}

func (m *TApplication) ModalFinished() {
	LCL().SysCallN(159, m.Instance())
}

func (m *TApplication) Restore() {
	LCL().SysCallN(174, m.Instance())
}

func (m *TApplication) Notification(AComponent IComponent, Operation TOperation) {
	LCL().SysCallN(165, m.Instance(), GetObjectUintptr(AComponent), uintptr(Operation))
}

func (m *TApplication) ProcessMessages() {
	LCL().SysCallN(169, m.Instance())
}

func (m *TApplication) Idle(Wait bool) {
	LCL().SysCallN(141, m.Instance(), PascalBool(Wait))
}

func (m *TApplication) DisableIdleHandler() {
	LCL().SysCallN(111, m.Instance())
}

func (m *TApplication) EnableIdleHandler() {
	LCL().SysCallN(118, m.Instance())
}

func (m *TApplication) NotifyUserInputHandler(Sender IObject, Msg uint32) {
	LCL().SysCallN(168, m.Instance(), GetObjectUintptr(Sender), uintptr(Msg))
}

func (m *TApplication) NotifyKeyDownBeforeHandler(Sender IObject, Key *Word, Shift TShiftState) {
	var result1 uintptr
	LCL().SysCallN(166, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafe.Pointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TApplication) NotifyKeyDownHandler(Sender IObject, Key *Word, Shift TShiftState) {
	var result1 uintptr
	LCL().SysCallN(167, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafe.Pointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TApplication) ControlKeyDown(Sender IObject, Key *Word, Shift TShiftState) {
	var result1 uintptr
	LCL().SysCallN(107, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafe.Pointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TApplication) ControlKeyUp(Sender IObject, Key *Word, Shift TShiftState) {
	var result1 uintptr
	LCL().SysCallN(108, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafe.Pointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TApplication) RemoveAllHandlersOfObject(AnObject IObject) {
	LCL().SysCallN(171, m.Instance(), GetObjectUintptr(AnObject))
}

func (m *TApplication) DoBeforeMouseMessage(CurMouseControl IControl) {
	LCL().SysCallN(113, m.Instance(), GetObjectUintptr(CurMouseControl))
}

func (m *TApplication) IntfEndSession() {
	LCL().SysCallN(146, m.Instance())
}

func (m *TApplication) IntfAppActivate(Async bool) {
	LCL().SysCallN(142, m.Instance(), PascalBool(Async))
}

func (m *TApplication) IntfAppDeactivate(Async bool) {
	LCL().SysCallN(143, m.Instance(), PascalBool(Async))
}

func (m *TApplication) IntfAppMinimize() {
	LCL().SysCallN(144, m.Instance())
}

func (m *TApplication) IntfAppRestore() {
	LCL().SysCallN(145, m.Instance())
}

func (m *TApplication) IntfSettingsChange() {
	LCL().SysCallN(147, m.Instance())
}

func (m *TApplication) IntfThemeOptionChange(AThemeServices IThemeServices, AOption TThemeOption) {
	LCL().SysCallN(148, m.Instance(), GetObjectUintptr(AThemeServices), uintptr(AOption))
}

func (m *TApplication) DoArrowKey(AControl IWinControl, Key *Word, Shift TShiftState) {
	var result1 uintptr
	LCL().SysCallN(112, m.Instance(), GetObjectUintptr(AControl), uintptr(unsafe.Pointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TApplication) DoTabKey(AControl IWinControl, Key *Word, Shift TShiftState) {
	var result1 uintptr
	LCL().SysCallN(116, m.Instance(), GetObjectUintptr(AControl), uintptr(unsafe.Pointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TApplication) DoEscapeKey(AControl IWinControl, Key *Word, Shift TShiftState) {
	var result1 uintptr
	LCL().SysCallN(114, m.Instance(), GetObjectUintptr(AControl), uintptr(unsafe.Pointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TApplication) DoReturnKey(AControl IWinControl, Key *Word, Shift TShiftState) {
	var result1 uintptr
	LCL().SysCallN(115, m.Instance(), GetObjectUintptr(AControl), uintptr(unsafe.Pointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TApplication) SetOnActionExecute(fn TActionEvent) {
	if m.actionExecutePtr != 0 {
		RemoveEventElement(m.actionExecutePtr)
	}
	m.actionExecutePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(177, m.Instance(), m.actionExecutePtr)
}

func (m *TApplication) SetOnActionUpdate(fn TActionEvent) {
	if m.actionUpdatePtr != 0 {
		RemoveEventElement(m.actionUpdatePtr)
	}
	m.actionUpdatePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(178, m.Instance(), m.actionUpdatePtr)
}

func (m *TApplication) SetOnActivate(fn TNotifyEvent) {
	if m.activatePtr != 0 {
		RemoveEventElement(m.activatePtr)
	}
	m.activatePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(179, m.Instance(), m.activatePtr)
}

func (m *TApplication) SetOnDeactivate(fn TNotifyEvent) {
	if m.deactivatePtr != 0 {
		RemoveEventElement(m.deactivatePtr)
	}
	m.deactivatePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(181, m.Instance(), m.deactivatePtr)
}

func (m *TApplication) SetOnGetMainFormHandle(fn TGetHandleEvent) {
	if m.getMainFormHandlePtr != 0 {
		RemoveEventElement(m.getMainFormHandlePtr)
	}
	m.getMainFormHandlePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(185, m.Instance(), m.getMainFormHandlePtr)
}

func (m *TApplication) SetOnIdleEnd(fn TNotifyEvent) {
	if m.idleEndPtr != 0 {
		RemoveEventElement(m.idleEndPtr)
	}
	m.idleEndPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(188, m.Instance(), m.idleEndPtr)
}

func (m *TApplication) SetOnEndSession(fn TNotifyEvent) {
	if m.endSessionPtr != 0 {
		RemoveEventElement(m.endSessionPtr)
	}
	m.endSessionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(184, m.Instance(), m.endSessionPtr)
}

func (m *TApplication) SetOnMinimize(fn TNotifyEvent) {
	if m.minimizePtr != 0 {
		RemoveEventElement(m.minimizePtr)
	}
	m.minimizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(189, m.Instance(), m.minimizePtr)
}

func (m *TApplication) SetOnModalBegin(fn TNotifyEvent) {
	if m.modalBeginPtr != 0 {
		RemoveEventElement(m.modalBeginPtr)
	}
	m.modalBeginPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(190, m.Instance(), m.modalBeginPtr)
}

func (m *TApplication) SetOnModalEnd(fn TNotifyEvent) {
	if m.modalEndPtr != 0 {
		RemoveEventElement(m.modalEndPtr)
	}
	m.modalEndPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(191, m.Instance(), m.modalEndPtr)
}

func (m *TApplication) SetOnRestore(fn TNotifyEvent) {
	if m.restorePtr != 0 {
		RemoveEventElement(m.restorePtr)
	}
	m.restorePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(192, m.Instance(), m.restorePtr)
}

func (m *TApplication) SetOnDropFiles(fn TDropFilesEvent) {
	if m.dropFilesPtr != 0 {
		RemoveEventElement(m.dropFilesPtr)
	}
	m.dropFilesPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(183, m.Instance(), m.dropFilesPtr)
}

func (m *TApplication) SetOnHelp(fn THelpEvent) {
	if m.helpPtr != 0 {
		RemoveEventElement(m.helpPtr)
	}
	m.helpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(186, m.Instance(), m.helpPtr)
}

func (m *TApplication) SetOnHint(fn TNotifyEvent) {
	if m.hintPtr != 0 {
		RemoveEventElement(m.hintPtr)
	}
	m.hintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(187, m.Instance(), m.hintPtr)
}

func (m *TApplication) SetOnShortcut(fn TShortCutEvent) {
	if m.shortcutPtr != 0 {
		RemoveEventElement(m.shortcutPtr)
	}
	m.shortcutPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(193, m.Instance(), m.shortcutPtr)
}

func (m *TApplication) SetOnShowHint(fn TShowHintEvent) {
	if m.showHintPtr != 0 {
		RemoveEventElement(m.showHintPtr)
	}
	m.showHintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(194, m.Instance(), m.showHintPtr)
}

func (m *TApplication) SetOnDestroy(fn TNotifyEvent) {
	if m.destroyPtr != 0 {
		RemoveEventElement(m.destroyPtr)
	}
	m.destroyPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(182, m.Instance(), m.destroyPtr)
}

func (m *TApplication) SetOnCircularException(fn TExceptionEvent) {
	if m.circularExceptionPtr != 0 {
		RemoveEventElement(m.circularExceptionPtr)
	}
	m.circularExceptionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(180, m.Instance(), m.circularExceptionPtr)
}
