//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICefPreferenceManager Parent: ICefBaseRefCounted
//
//	Manage access to preferences. Many built-in preferences are registered by Chromium. Custom preferences can be registered in ICefBrowserProcessHandler.OnRegisterCustomPreferences.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_preference_capi.h">CEF source file: /include/capi/cef_preference_capi.h (cef_preference_manager_t)) <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_preference_manager_capi.h">CEF source file: /include/capi/cef_preference_manager_capi.h (cef_preference_manager_t))
type ICefPreferenceManager interface {
	ICefBaseRefCounted
	// HasPreference
	//  Returns true (1) if a preference with the specified |name| exists. This function must be called on the browser process UI thread.
	HasPreference(name string) bool // function
	// GetPreference
	//  Returns the value for the preference with the specified |name|. Returns NULL if the preference does not exist. The returned object contains a copy of the underlying preference value and modifications to the returned object will not modify the underlying preference value. This function must be called on the browser process UI thread.
	GetPreference(name string) ICefValue // function
	// GetAllPreferences
	//  Returns all preferences as a dictionary. If |include_defaults| is true (1) then preferences currently at their default value will be included. The returned object contains a copy of the underlying preference values and modifications to the returned object will not modify the underlying preference values. This function must be called on the browser process UI thread.
	GetAllPreferences(includeDefaults bool) ICefDictionaryValue // function
	// CanSetPreference
	//  Returns true (1) if the preference with the specified |name| can be modified using SetPreference. As one example preferences set via the command-line usually cannot be modified. This function must be called on the browser process UI thread.
	CanSetPreference(name string) bool // function
	// SetPreference
	//  Set the |value| associated with preference |name|. Returns true (1) if the value is set successfully and false (0) otherwise. If |value| is NULL the preference will be restored to its default value. If setting the preference fails then |error| will be populated with a detailed description of the problem. This function must be called on the browser process UI thread.
	SetPreference(name string, value ICefValue, outError *string) bool // function
}

// TCefPreferenceManager Parent: TCefBaseRefCounted
//
//	Manage access to preferences. Many built-in preferences are registered by Chromium. Custom preferences can be registered in ICefBrowserProcessHandler.OnRegisterCustomPreferences.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_preference_capi.h">CEF source file: /include/capi/cef_preference_capi.h (cef_preference_manager_t)) <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_preference_manager_capi.h">CEF source file: /include/capi/cef_preference_manager_capi.h (cef_preference_manager_t))
type TCefPreferenceManager struct {
	TCefBaseRefCounted
}

// PreferenceManagerRef -> ICefPreferenceManager
var PreferenceManagerRef preferenceManager

// preferenceManager TCefPreferenceManager Ref
type preferenceManager uintptr

func (m *preferenceManager) UnWrap(data uintptr) ICefPreferenceManager {
	var resultCefPreferenceManager uintptr
	CEF().SysCallN(1212, uintptr(data), uintptr(unsafePointer(&resultCefPreferenceManager)))
	return AsCefPreferenceManager(resultCefPreferenceManager)
}

func (m *preferenceManager) Global() ICefPreferenceManager {
	var resultCefPreferenceManager uintptr
	CEF().SysCallN(1209, uintptr(unsafePointer(&resultCefPreferenceManager)))
	return AsCefPreferenceManager(resultCefPreferenceManager)
}

func (m *TCefPreferenceManager) HasPreference(name string) bool {
	r1 := CEF().SysCallN(1210, m.Instance(), PascalStr(name))
	return GoBool(r1)
}

func (m *TCefPreferenceManager) GetPreference(name string) ICefValue {
	var resultCefValue uintptr
	CEF().SysCallN(1208, m.Instance(), PascalStr(name), uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *TCefPreferenceManager) GetAllPreferences(includeDefaults bool) ICefDictionaryValue {
	var resultCefDictionaryValue uintptr
	CEF().SysCallN(1207, m.Instance(), PascalBool(includeDefaults), uintptr(unsafePointer(&resultCefDictionaryValue)))
	return AsCefDictionaryValue(resultCefDictionaryValue)
}

func (m *TCefPreferenceManager) CanSetPreference(name string) bool {
	r1 := CEF().SysCallN(1206, m.Instance(), PascalStr(name))
	return GoBool(r1)
}

func (m *TCefPreferenceManager) SetPreference(name string, value ICefValue, outError *string) bool {
	var result2 uintptr
	r1 := CEF().SysCallN(1211, m.Instance(), PascalStr(name), GetObjectUintptr(value), uintptr(unsafePointer(&result2)))
	*outError = GoStr(result2)
	return GoBool(r1)
}
