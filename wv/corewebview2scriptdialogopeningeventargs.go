//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// ICoreWebView2ScriptDialogOpeningEventArgs Parent: IObject
//
//	Event args for the ScriptDialogOpening event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs">See the ICoreWebView2ScriptDialogOpeningEventArgs article.</a>
type ICoreWebView2ScriptDialogOpeningEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ScriptDialogOpeningEventArgs // property
	// URI
	//  The URI of the page that requested the dialog box.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs#get_uri">See the ICoreWebView2ScriptDialogOpeningEventArgs article.</a>
	URI() string // property
	// Kind
	//  The kind of JavaScript dialog box. `alert`, `confirm`, `prompt`, or
	//  `beforeunload`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs#get_kind">See the ICoreWebView2ScriptDialogOpeningEventArgs article.</a>
	Kind() TWVScriptDialogKind // property
	// Message
	//  The message of the dialog box. From JavaScript this is the first
	//  parameter passed to `alert`, `confirm`, and `prompt` and is empty for
	//  `beforeunload`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs#get_message">See the ICoreWebView2ScriptDialogOpeningEventArgs article.</a>
	Message() string // property
	// DefaultText
	//  The second parameter passed to the JavaScript prompt dialog.
	//  The result of the prompt JavaScript function uses this value as the
	//  default value.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs#get_defaulttext">See the ICoreWebView2ScriptDialogOpeningEventArgs article.</a>
	DefaultText() string // property
	// ResultText
	//  The return value from the JavaScript prompt function if `Accept` is run.
	//  This value is ignored for dialog kinds other than prompt. If `Accept`
	//  is not run, this value is ignored and `FALSE` is returned from prompt.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs#get_resulttext">See the ICoreWebView2ScriptDialogOpeningEventArgs article.</a>
	ResultText() string // property
	// SetResultText Set ResultText
	SetResultText(AValue string) // property
	// Deferral
	//  Returns an `ICoreWebView2Deferral` object. Use this operation to
	//  complete the event at a later time.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs#getdeferral">See the ICoreWebView2ScriptDialogOpeningEventArgs article.</a>
	Deferral() ICoreWebView2Deferral // property
	// Accept
	//  The host may run this to respond with **OK** to `confirm`, `prompt`, and
	//  `beforeunload` dialogs. Do not run this method to indicate cancel.
	//  From JavaScript, this means that the `confirm` and `beforeunload` function
	//  returns `TRUE` if `Accept` is run. And for the prompt function it returns
	//  the value of `ResultText` if `Accept` is run and otherwise returns
	//  `FALSE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs#accept">See the ICoreWebView2ScriptDialogOpeningEventArgs article.</a>
	Accept() bool // function
}

// TCoreWebView2ScriptDialogOpeningEventArgs Parent: TObject
//
//	Event args for the ScriptDialogOpening event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2scriptdialogopeningeventargs">See the ICoreWebView2ScriptDialogOpeningEventArgs article.</a>
type TCoreWebView2ScriptDialogOpeningEventArgs struct {
	TObject
}

func NewCoreWebView2ScriptDialogOpeningEventArgs(aArgs ICoreWebView2ScriptDialogOpeningEventArgs) ICoreWebView2ScriptDialogOpeningEventArgs {
	r1 := WV().SysCallN(585, GetObjectUintptr(aArgs))
	return AsCoreWebView2ScriptDialogOpeningEventArgs(r1)
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) Initialized() bool {
	r1 := WV().SysCallN(588, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) BaseIntf() ICoreWebView2ScriptDialogOpeningEventArgs {
	var resultCoreWebView2ScriptDialogOpeningEventArgs uintptr
	WV().SysCallN(583, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ScriptDialogOpeningEventArgs)))
	return AsCoreWebView2ScriptDialogOpeningEventArgs(resultCoreWebView2ScriptDialogOpeningEventArgs)
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) URI() string {
	r1 := WV().SysCallN(592, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) Kind() TWVScriptDialogKind {
	r1 := WV().SysCallN(589, m.Instance())
	return TWVScriptDialogKind(r1)
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) Message() string {
	r1 := WV().SysCallN(590, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) DefaultText() string {
	r1 := WV().SysCallN(586, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) ResultText() string {
	r1 := WV().SysCallN(591, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) SetResultText(AValue string) {
	WV().SysCallN(591, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) Deferral() ICoreWebView2Deferral {
	var resultCoreWebView2Deferral uintptr
	WV().SysCallN(587, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Deferral)))
	return AsCoreWebView2Deferral(resultCoreWebView2Deferral)
}

func (m *TCoreWebView2ScriptDialogOpeningEventArgs) Accept() bool {
	r1 := WV().SysCallN(582, m.Instance())
	return GoBool(r1)
}

func CoreWebView2ScriptDialogOpeningEventArgsClass() TClass {
	ret := WV().SysCallN(584)
	return TClass(ret)
}
