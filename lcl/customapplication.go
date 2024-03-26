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

// ICustomApplication Parent: IComponent
type ICustomApplication interface {
	IComponent
	ExeName() string                                                                                    // property
	HelpFile() string                                                                                   // property
	SetHelpFile(AValue string)                                                                          // property
	Terminated() bool                                                                                   // property
	Title() string                                                                                      // property
	SetTitle(AValue string)                                                                             // property
	Location() string                                                                                   // property
	Params(Index int32) string                                                                          // property
	ParamCount() int32                                                                                  // property
	EnvironmentVariable(envName string) string                                                          // property
	OptionChar() Char                                                                                   // property
	SetOptionChar(AValue Char)                                                                          // property
	CaseSensitiveOptions() bool                                                                         // property
	SetCaseSensitiveOptions(AValue bool)                                                                // property
	StopOnException() bool                                                                              // property
	SetStopOnException(AValue bool)                                                                     // property
	ExceptionExitCode() int32                                                                           // property
	SetExceptionExitCode(AValue int32)                                                                  // property
	EventLogFilter() TEventLogTypes                                                                     // property
	SetEventLogFilter(AValue TEventLogTypes)                                                            // property
	FindOptionIndex(S string, Longopt *bool, StartAt int32) int32                                       // function
	GetOptionValue(S string) string                                                                     // function
	GetOptionValue1(C Char, S string) string                                                            // function
	GetOptionValues(C Char, S string) TStringArray                                                      // function
	HasOption(S string) bool                                                                            // function
	HasOption1(C Char, S string) bool                                                                   // function
	CheckOptions(ShortOptions string, Longopts IStrings, Opts, NonOpts IStrings, AllErrors bool) string // function
	CheckOptions1(ShortOptions string, Longopts IStrings, AllErrors bool) string                        // function
	CheckOptions2(ShortOptions string, LongOpts string, AllErrors bool) string                          // function
	HandleException(Sender IObject)                                                                     // procedure
	Initialize()                                                                                        // procedure
	Run()                                                                                               // procedure
	ShowException(E IException)                                                                         // procedure
	Terminate()                                                                                         // procedure
	Terminate1(AExitCode int32)                                                                         // procedure
	GetEnvironmentList(List IStrings, NamesOnly bool)                                                   // procedure
	GetEnvironmentList1(List IStrings)                                                                  // procedure
	Log(EventType TEventType, Msg string)                                                               // procedure
	SetOnException(fn TExceptionEvent)                                                                  // property event
}

// TCustomApplication Parent: TComponent
type TCustomApplication struct {
	TComponent
	exceptionPtr uintptr
}

func NewCustomApplication(AOwner IComponent) ICustomApplication {
	r1 := LCL().SysCallN(1086, GetObjectUintptr(AOwner))
	return AsCustomApplication(r1)
}

func (m *TCustomApplication) ExeName() string {
	r1 := LCL().SysCallN(1090, m.Instance())
	return GoStr(r1)
}

func (m *TCustomApplication) HelpFile() string {
	r1 := LCL().SysCallN(1100, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomApplication) SetHelpFile(AValue string) {
	LCL().SysCallN(1100, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomApplication) Terminated() bool {
	r1 := LCL().SysCallN(1113, m.Instance())
	return GoBool(r1)
}

func (m *TCustomApplication) Title() string {
	r1 := LCL().SysCallN(1114, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomApplication) SetTitle(AValue string) {
	LCL().SysCallN(1114, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomApplication) Location() string {
	r1 := LCL().SysCallN(1102, m.Instance())
	return GoStr(r1)
}

func (m *TCustomApplication) Params(Index int32) string {
	r1 := LCL().SysCallN(1106, m.Instance(), uintptr(Index))
	return GoStr(r1)
}

func (m *TCustomApplication) ParamCount() int32 {
	r1 := LCL().SysCallN(1105, m.Instance())
	return int32(r1)
}

func (m *TCustomApplication) EnvironmentVariable(envName string) string {
	r1 := LCL().SysCallN(1087, m.Instance(), PascalStr(envName))
	return GoStr(r1)
}

func (m *TCustomApplication) OptionChar() Char {
	r1 := LCL().SysCallN(1104, 0, m.Instance(), 0)
	return Char(r1)
}

func (m *TCustomApplication) SetOptionChar(AValue Char) {
	LCL().SysCallN(1104, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomApplication) CaseSensitiveOptions() bool {
	r1 := LCL().SysCallN(1081, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomApplication) SetCaseSensitiveOptions(AValue bool) {
	LCL().SysCallN(1081, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomApplication) StopOnException() bool {
	r1 := LCL().SysCallN(1110, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomApplication) SetStopOnException(AValue bool) {
	LCL().SysCallN(1110, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomApplication) ExceptionExitCode() int32 {
	r1 := LCL().SysCallN(1089, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomApplication) SetExceptionExitCode(AValue int32) {
	LCL().SysCallN(1089, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomApplication) EventLogFilter() TEventLogTypes {
	r1 := LCL().SysCallN(1088, 0, m.Instance(), 0)
	return TEventLogTypes(r1)
}

func (m *TCustomApplication) SetEventLogFilter(AValue TEventLogTypes) {
	LCL().SysCallN(1088, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomApplication) FindOptionIndex(S string, Longopt *bool, StartAt int32) int32 {
	var result1 uintptr
	r1 := LCL().SysCallN(1091, m.Instance(), PascalStr(S), uintptr(unsafe.Pointer(&result1)), uintptr(StartAt))
	*Longopt = GoBool(result1)
	return int32(r1)
}

func (m *TCustomApplication) GetOptionValue(S string) string {
	r1 := LCL().SysCallN(1094, m.Instance(), PascalStr(S))
	return GoStr(r1)
}

func (m *TCustomApplication) GetOptionValue1(C Char, S string) string {
	r1 := LCL().SysCallN(1095, m.Instance(), uintptr(C), PascalStr(S))
	return GoStr(r1)
}

func (m *TCustomApplication) GetOptionValues(C Char, S string) TStringArray {
	r1 := LCL().SysCallN(1096, m.Instance(), uintptr(C), PascalStr(S))
	return TStringArray(r1)
}

func (m *TCustomApplication) HasOption(S string) bool {
	r1 := LCL().SysCallN(1098, m.Instance(), PascalStr(S))
	return GoBool(r1)
}

func (m *TCustomApplication) HasOption1(C Char, S string) bool {
	r1 := LCL().SysCallN(1099, m.Instance(), uintptr(C), PascalStr(S))
	return GoBool(r1)
}

func (m *TCustomApplication) CheckOptions(ShortOptions string, Longopts IStrings, Opts, NonOpts IStrings, AllErrors bool) string {
	r1 := LCL().SysCallN(1082, m.Instance(), PascalStr(ShortOptions), GetObjectUintptr(Longopts), GetObjectUintptr(Opts), GetObjectUintptr(NonOpts), PascalBool(AllErrors))
	return GoStr(r1)
}

func (m *TCustomApplication) CheckOptions1(ShortOptions string, Longopts IStrings, AllErrors bool) string {
	r1 := LCL().SysCallN(1083, m.Instance(), PascalStr(ShortOptions), GetObjectUintptr(Longopts), PascalBool(AllErrors))
	return GoStr(r1)
}

func (m *TCustomApplication) CheckOptions2(ShortOptions string, LongOpts string, AllErrors bool) string {
	r1 := LCL().SysCallN(1084, m.Instance(), PascalStr(ShortOptions), PascalStr(LongOpts), PascalBool(AllErrors))
	return GoStr(r1)
}

func CustomApplicationClass() TClass {
	ret := LCL().SysCallN(1085)
	return TClass(ret)
}

func (m *TCustomApplication) HandleException(Sender IObject) {
	LCL().SysCallN(1097, m.Instance(), GetObjectUintptr(Sender))
}

func (m *TCustomApplication) Initialize() {
	LCL().SysCallN(1101, m.Instance())
}

func (m *TCustomApplication) Run() {
	LCL().SysCallN(1107, m.Instance())
}

func (m *TCustomApplication) ShowException(E IException) {
	LCL().SysCallN(1109, m.Instance(), GetObjectUintptr(E))
}

func (m *TCustomApplication) Terminate() {
	LCL().SysCallN(1111, m.Instance())
}

func (m *TCustomApplication) Terminate1(AExitCode int32) {
	LCL().SysCallN(1112, m.Instance(), uintptr(AExitCode))
}

func (m *TCustomApplication) GetEnvironmentList(List IStrings, NamesOnly bool) {
	LCL().SysCallN(1092, m.Instance(), GetObjectUintptr(List), PascalBool(NamesOnly))
}

func (m *TCustomApplication) GetEnvironmentList1(List IStrings) {
	LCL().SysCallN(1093, m.Instance(), GetObjectUintptr(List))
}

func (m *TCustomApplication) Log(EventType TEventType, Msg string) {
	LCL().SysCallN(1103, m.Instance(), uintptr(EventType), PascalStr(Msg))
}

func (m *TCustomApplication) SetOnException(fn TExceptionEvent) {
	if m.exceptionPtr != 0 {
		RemoveEventElement(m.exceptionPtr)
	}
	m.exceptionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1108, m.Instance(), m.exceptionPtr)
}
