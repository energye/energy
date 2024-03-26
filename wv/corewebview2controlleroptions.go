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

// ICoreWebView2ControllerOptions Parent: IObject
//
//	This class is a ICoreWebView2ControllerOptions wrapper.
//	ICoreWebView2ControllerOptions is used to manage profile options that created by 'CreateCoreWebView2ControllerOptions'.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions">See the ICoreWebView2ControllerOptions article.</a>
type ICoreWebView2ControllerOptions interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ControllerOptions // property
	// SetBaseIntf Set BaseIntf
	SetBaseIntf(AValue ICoreWebView2ControllerOptions) // property
	// ProfileName
	//  `ProfileName` property is to specify a profile name, which is only allowed to contain
	//  the following ASCII characters. It has a maximum length of 64 characters excluding the null-terminator.
	//  It is ASCII case insensitive.
	//  * alphabet characters: a-z and A-Z
	//  * digit characters: 0-9
	//  * and '#', '@', '$', '(', ')', '+', '-', '_', '~', '.', ' '(space).
	//  Note: the text must not end with a period '.' or ' '(space). And, although upper-case letters are
	//  allowed, they're treated just as lower-case counterparts because the profile name will be mapped to
	//  the real profile directory path on disk and Windows file system handles path names in a case-insensitive way.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions#get_profilename">See the ICoreWebView2ControllerOptions article.</a>
	ProfileName() string // property
	// SetProfileName Set ProfileName
	SetProfileName(AValue string) // property
	// IsInPrivateModeEnabled
	//  `IsInPrivateModeEnabled` property is to enable/disable InPrivate mode.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions#get_isinprivatemodeenabled">See the ICoreWebView2ControllerOptions article.</a>
	IsInPrivateModeEnabled() bool // property
	// SetIsInPrivateModeEnabled Set IsInPrivateModeEnabled
	SetIsInPrivateModeEnabled(AValue bool) // property
	// ScriptLocale
	//  The default locale for the WebView2. It sets the default locale for all
	//  Intl JavaScript APIs and other JavaScript APIs that depend on it, namely
	//  `Intl.DateTimeFormat()` which affects string formatting like
	//  in the time/date formats. Example: `Intl.DateTimeFormat().format(new Date())`
	//  The intended locale value is in the format of
	//  BCP 47 Language Tags. More information can be found from
	//  [IETF BCP47](https://www.ietf.org/rfc/bcp/bcp47.html).
	//  This property sets the locale for a CoreWebView2Environment used to create the
	//  WebView2ControllerOptions object, which is passed as a parameter in
	//  `CreateCoreWebView2ControllerWithOptions`.
	//  Changes to the ScriptLocale property apply to renderer processes created after
	//  the change. Any existing renderer processes will continue to use the previous
	//  ScriptLocale value. To ensure changes are applied to all renderer process,
	//  close and restart the CoreWebView2Environment and all associated WebView2 objects.
	//  The default value for ScriptLocale will depend on the WebView2 language
	//  and OS region. If the language portions of the WebView2 language and OS region
	//  match, then it will use the OS region. Otherwise, it will use the WebView2
	//  language.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions2#get_scriptlocale">See the ICoreWebView2ControllerOptions2 article.</a>
	ScriptLocale() string // property
	// SetScriptLocale Set ScriptLocale
	SetScriptLocale(AValue string) // property
}

// TCoreWebView2ControllerOptions Parent: TObject
//
//	This class is a ICoreWebView2ControllerOptions wrapper.
//	ICoreWebView2ControllerOptions is used to manage profile options that created by 'CreateCoreWebView2ControllerOptions'.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions">See the ICoreWebView2ControllerOptions article.</a>
type TCoreWebView2ControllerOptions struct {
	TObject
}

func NewCoreWebView2ControllerOptions(aBaseIntf ICoreWebView2ControllerOptions) ICoreWebView2ControllerOptions {
	r1 := WV().SysCallN(177, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ControllerOptions(r1)
}

func (m *TCoreWebView2ControllerOptions) Initialized() bool {
	r1 := WV().SysCallN(178, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ControllerOptions) BaseIntf() ICoreWebView2ControllerOptions {
	var resultCoreWebView2ControllerOptions uintptr
	WV().SysCallN(175, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2ControllerOptions)))
	return AsCoreWebView2ControllerOptions(resultCoreWebView2ControllerOptions)
}

func (m *TCoreWebView2ControllerOptions) SetBaseIntf(AValue ICoreWebView2ControllerOptions) {
	WV().SysCallN(175, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2ControllerOptions) ProfileName() string {
	r1 := WV().SysCallN(180, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2ControllerOptions) SetProfileName(AValue string) {
	WV().SysCallN(180, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoreWebView2ControllerOptions) IsInPrivateModeEnabled() bool {
	r1 := WV().SysCallN(179, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2ControllerOptions) SetIsInPrivateModeEnabled(AValue bool) {
	WV().SysCallN(179, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2ControllerOptions) ScriptLocale() string {
	r1 := WV().SysCallN(181, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2ControllerOptions) SetScriptLocale(AValue string) {
	WV().SysCallN(181, 1, m.Instance(), PascalStr(AValue))
}

func CoreWebView2ControllerOptionsClass() TClass {
	ret := WV().SysCallN(176)
	return TClass(ret)
}
