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

// ICefExtension Parent: ICefBaseRefCounted
//
//	Object representing an extension. Methods may be called on any thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_extension_capi.h">CEF source file: /include/capi/cef_extension_capi.h (cef_extension_t))</a>
type ICefExtension interface {
	ICefBaseRefCounted
	// GetIdentifier
	//  Returns the unique extension identifier. This is calculated based on the extension public key, if available, or on the extension path. See https://developer.chrome.com/extensions/manifest/key for details.
	GetIdentifier() string // function
	// GetPath
	//  Returns the absolute path to the extension directory on disk. This value will be prefixed with PK_DIR_RESOURCES if a relative path was passed to ICefRequestContext.LoadExtension.
	GetPath() string // function
	// GetManifest
	//  Returns the extension manifest contents as a ICefDictionaryValue object. See https://developer.chrome.com/extensions/manifest for details.
	GetManifest() ICefDictionaryValue // function
	// IsSame
	//  Returns true (1) if this object is the same extension as |that| object. Extensions are considered the same if identifier, path and loader context match.
	IsSame(that ICefExtension) bool // function
	// GetHandler
	//  Returns the handler for this extension. Will return NULL for internal extensions or if no handler was passed to ICefRequestContext.LoadExtension.
	GetHandler() ICefExtensionHandler // function
	// GetLoaderContext
	//  Returns the request context that loaded this extension. Will return NULL for internal extensions or if the extension has been unloaded. See the ICefRequestContext.LoadExtension documentation for more information about loader contexts. Must be called on the browser process UI thread.
	GetLoaderContext() ICefRequestContext // function
	// IsLoaded
	//  Returns true (1) if this extension is currently loaded. Must be called on the browser process UI thread.
	IsLoaded() bool                // function
	GetBrowserActionPopup() string // function
	GetBrowserActionIcon() string  // function
	GetPageActionPopup() string    // function
	GetPageActionIcon() string     // function
	GetOptionsPage() string        // function
	GetOptionsUIPage() string      // function
	GetBackgroundPage() string     // function
	GetURL() string                // function
	// Unload
	//  Unload this extension if it is not an internal extension and is currently loaded. Will result in a call to ICefExtensionHandler.OnExtensionUnloaded on success.
	Unload() // procedure
}

// TCefExtension Parent: TCefBaseRefCounted
//
//	Object representing an extension. Methods may be called on any thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_extension_capi.h">CEF source file: /include/capi/cef_extension_capi.h (cef_extension_t))</a>
type TCefExtension struct {
	TCefBaseRefCounted
}

// ExtensionRef -> ICefExtension
var ExtensionRef extension

// extension TCefExtension Ref
type extension uintptr

func (m *extension) UnWrap(data uintptr) ICefExtension {
	var resultCefExtension uintptr
	CEF().SysCallN(949, uintptr(data), uintptr(unsafePointer(&resultCefExtension)))
	return AsCefExtension(resultCefExtension)
}

func (m *TCefExtension) GetIdentifier() string {
	r1 := CEF().SysCallN(938, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetPath() string {
	r1 := CEF().SysCallN(945, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetManifest() ICefDictionaryValue {
	var resultCefDictionaryValue uintptr
	CEF().SysCallN(940, m.Instance(), uintptr(unsafePointer(&resultCefDictionaryValue)))
	return AsCefDictionaryValue(resultCefDictionaryValue)
}

func (m *TCefExtension) IsSame(that ICefExtension) bool {
	r1 := CEF().SysCallN(948, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefExtension) GetHandler() ICefExtensionHandler {
	var resultCefExtensionHandler uintptr
	CEF().SysCallN(937, m.Instance(), uintptr(unsafePointer(&resultCefExtensionHandler)))
	return AsCefExtensionHandler(resultCefExtensionHandler)
}

func (m *TCefExtension) GetLoaderContext() ICefRequestContext {
	var resultCefRequestContext uintptr
	CEF().SysCallN(939, m.Instance(), uintptr(unsafePointer(&resultCefRequestContext)))
	return AsCefRequestContext(resultCefRequestContext)
}

func (m *TCefExtension) IsLoaded() bool {
	r1 := CEF().SysCallN(947, m.Instance())
	return GoBool(r1)
}

func (m *TCefExtension) GetBrowserActionPopup() string {
	r1 := CEF().SysCallN(936, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetBrowserActionIcon() string {
	r1 := CEF().SysCallN(935, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetPageActionPopup() string {
	r1 := CEF().SysCallN(944, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetPageActionIcon() string {
	r1 := CEF().SysCallN(943, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetOptionsPage() string {
	r1 := CEF().SysCallN(941, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetOptionsUIPage() string {
	r1 := CEF().SysCallN(942, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetBackgroundPage() string {
	r1 := CEF().SysCallN(934, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) GetURL() string {
	r1 := CEF().SysCallN(946, m.Instance())
	return GoStr(r1)
}

func (m *TCefExtension) Unload() {
	CEF().SysCallN(950, m.Instance())
}
