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

// ICefRequestContext Parent: ICefPreferenceManager
//
//	A request context provides request handling for a set of related browser or
//	URL request objects. A request context can be specified when creating a new
//	browser via the cef_browser_host_t static factory functions or when creating
//	a new URL request via the cef_urlrequest_t static factory functions. Browser
//	objects with different request contexts will never be hosted in the same
//	render process. Browser objects with the same request context may or may not
//	be hosted in the same render process depending on the process model. Browser
//	objects created indirectly via the JavaScript window.open function or
//	targeted links will share the same render process and the same request
//	context as the source browser. When running in single-process mode there is
//	only a single render process (the main process) and so all browsers created
//	in single-process mode will share the same request context. This will be the
//	first request context passed into a cef_browser_host_t static factory
//	function and all other request context objects will be ignored.
type ICefRequestContext interface {
	ICefPreferenceManager
	// IsSame
	//  Returns true(1) if this object is pointing to the same context as |that|
	//  object.
	IsSame(other ICefRequestContext) bool // function
	// IsSharingWith
	//  Returns true(1) if this object is sharing the same storage as |that|
	//  object.
	IsSharingWith(other ICefRequestContext) bool // function
	// IsGlobal
	//  Returns true(1) if this object is the global context. The global context
	//  is used by default when creating a browser or URL request with a NULL
	//  context argument.
	IsGlobal() bool // function
	// GetHandler
	//  Returns the handler for this context if any.
	GetHandler() ICefRequestContextHandler // function
	// GetCachePath
	//  Returns the cache path for this object. If NULL an "incognito mode" in-
	//  memory cache is being used.
	GetCachePath() string // function
	// GetCookieManager
	//  Returns the cookie manager for this object. If |callback| is non-NULL it
	//  will be executed asnychronously on the UI thread after the manager's
	//  storage has been initialized.
	GetCookieManager(callback ICefCompletionCallback) ICefCookieManager // function
	// RegisterSchemeHandlerFactory
	//  Register a scheme handler factory for the specified |scheme_name| and
	//  optional |domain_name|. An NULL |domain_name| value for a standard scheme
	//  will cause the factory to match all domain names. The |domain_name| value
	//  will be ignored for non-standard schemes. If |scheme_name| is a built-in
	//  scheme and no handler is returned by |factory| then the built-in scheme
	//  handler factory will be called. If |scheme_name| is a custom scheme then
	//  you must also implement the cef_app_t::on_register_custom_schemes()
	//  function in all processes. This function may be called multiple times to
	//  change or remove the factory that matches the specified |scheme_name| and
	//  optional |domain_name|. Returns false(0) if an error occurs. This
	//  function may be called on any thread in the browser process.
	RegisterSchemeHandlerFactory(schemeName, domainName string, factory ICefSchemeHandlerFactory) bool // function
	// ClearSchemeHandlerFactories
	//  Clear all registered scheme handler factories. Returns false(0) on error.
	//  This function may be called on any thread in the browser process.
	ClearSchemeHandlerFactories() bool // function
	// DidLoadExtension
	//  Returns true(1) if this context was used to load the extension identified
	//  by |extension_id|. Other contexts sharing the same storage will also have
	//  access to the extension(see HasExtension). This function must be called
	//  on the browser process UI thread.
	DidLoadExtension(extensionid string) bool // function
	// HasExtension
	//  Returns true(1) if this context has access to the extension identified by
	//  |extension_id|. This may not be the context that was used to load the
	//  extension(see DidLoadExtension). This function must be called on the
	//  browser process UI thread.
	HasExtension(extensionid string) bool // function
	// GetExtensions
	//  Retrieve the list of all extensions that this context has access to(see
	//  HasExtension). |extension_ids| will be populated with the list of
	//  extension ID values. Returns true(1) on success. This function must be
	//  called on the browser process UI thread.
	GetExtensions(extensionids IStringList) bool // function
	// GetExtension
	//  Returns the extension matching |extension_id| or NULL if no matching
	//  extension is accessible in this context(see HasExtension). This function
	//  must be called on the browser process UI thread.
	GetExtension(extensionid string) ICefExtension // function
	// GetMediaRouter
	//  Returns the MediaRouter object associated with this context. If
	//  |callback| is non-NULL it will be executed asnychronously on the UI thread
	//  after the manager's context has been initialized.
	GetMediaRouter(callback ICefCompletionCallback) ICefMediaRouter // function
	// GetWebsiteSetting
	//  Returns the current value for |content_type| that applies for the
	//  specified URLs. If both URLs are NULL the default value will be returned.
	//  Returns nullptr if no value is configured. Must be called on the browser
	//  process UI thread.
	GetWebsiteSetting(requestingurl, toplevelurl string, contenttype TCefContentSettingTypes) ICefValue // function
	// GetContentSetting
	//  Returns the current value for |content_type| that applies for the
	//  specified URLs. If both URLs are NULL the default value will be returned.
	//  Returns CEF_CONTENT_SETTING_VALUE_DEFAULT if no value is configured. Must
	//  be called on the browser process UI thread.
	GetContentSetting(requestingurl, toplevelurl string, contenttype TCefContentSettingTypes) TCefContentSettingValues // function
	// ClearCertificateExceptions
	//  Clears all certificate exceptions that were added as part of handling
	//  cef_request_handler_t::on_certificate_error(). If you call this it is
	//  recommended that you also call close_all_connections() or you risk not
	//  being prompted again for server certificates if you reconnect quickly. If
	//  |callback| is non-NULL it will be executed on the UI thread after
	//  completion.
	ClearCertificateExceptions(callback ICefCompletionCallback) // procedure
	// ClearHttpAuthCredentials
	//  Clears all HTTP authentication credentials that were added as part of
	//  handling GetAuthCredentials. If |callback| is non-NULL it will be executed
	//  on the UI thread after completion.
	ClearHttpAuthCredentials(callback ICefCompletionCallback) // procedure
	// CloseAllConnections
	//  Clears all active and idle connections that Chromium currently has. This
	//  is only recommended if you have released all other CEF objects but don't
	//  yet want to call cef_shutdown(). If |callback| is non-NULL it will be
	//  executed on the UI thread after completion.
	CloseAllConnections(callback ICefCompletionCallback) // procedure
	// ResolveHost
	//  Attempts to resolve |origin| to a list of associated IP addresses.
	//  |callback| will be executed on the UI thread after completion.
	ResolveHost(origin string, callback ICefResolveCallback) // procedure
	// LoadExtension
	//  Load an extension.
	//  If extension resources will be read from disk using the default load
	//  implementation then |root_directory| should be the absolute path to the
	//  extension resources directory and |manifest| should be NULL. If extension
	//  resources will be provided by the client(e.g. via cef_request_handler_t
	//  and/or cef_extension_handler_t) then |root_directory| should be a path
	//  component unique to the extension(if not absolute this will be internally
	//  prefixed with the PK_DIR_RESOURCES path) and |manifest| should contain the
	//  contents that would otherwise be read from the "manifest.json" file on
	//  disk.
	//  The loaded extension will be accessible in all contexts sharing the same
	//  storage(HasExtension returns true(1)). However, only the context on
	//  which this function was called is considered the loader(DidLoadExtension
	//  returns true(1)) and only the loader will receive
	//  cef_request_context_handler_t callbacks for the extension.
	//  cef_extension_handler_t::OnExtensionLoaded will be called on load success
	//  or cef_extension_handler_t::OnExtensionLoadFailed will be called on load
	//  failure.
	//  If the extension specifies a background script via the "background"
	//  manifest key then cef_extension_handler_t::OnBeforeBackgroundBrowser will
	//  be called to create the background browser. See that function for
	//  additional information about background scripts.
	//  For visible extension views the client application should evaluate the
	//  manifest to determine the correct extension URL to load and then pass that
	//  URL to the cef_browser_host_t::CreateBrowser* function after the extension
	//  has loaded. For example, the client can look for the "browser_action"
	//  manifest key as documented at
	//  https://developer.chrome.com/extensions/browserAction. Extension URLs take
	//  the form "chrome-extension://<extension_id>/<path>".
	//  Browsers that host extensions differ from normal browsers as follows:
	//  - Can access chrome.* JavaScript APIs if allowed by the manifest. Visit
	//  chrome://extensions-support for the list of extension APIs currently
	//  supported by CEF.
	//  - Main frame navigation to non-extension content is blocked.
	//  - Pinch-zooming is disabled.
	//  - CefBrowserHost::GetExtension returns the hosted extension.
	//  - CefBrowserHost::IsBackgroundHost returns true for background hosts.
	//  See https://developer.chrome.com/extensions for extension implementation
	//  and usage documentation.
	LoadExtension(rootdirectory string, manifest ICefDictionaryValue, handler ICefExtensionHandler) // procedure
	// SetWebsiteSetting
	//  Sets the current value for |content_type| for the specified URLs in the
	//  default scope. If both URLs are NULL, and the context is not incognito,
	//  the default value will be set. Pass nullptr for |value| to remove the
	//  default value for this content type.
	//  WARNING: Incorrect usage of this function may cause instability or
	//  security issues in Chromium. Make sure that you first understand the
	//  potential impact of any changes to |content_type| by reviewing the related
	//  source code in Chromium. For example, if you plan to modify
	//  CEF_CONTENT_SETTING_TYPE_POPUPS, first review and understand the usage of
	//  ContentSettingsType::POPUPS in Chromium:
	//  https://source.chromium.org/search?q=ContentSettingsType::POPUPS
	SetWebsiteSetting(requestingurl, toplevelurl string, contenttype TCefContentSettingTypes, value ICefValue) // procedure
	// SetContentSetting
	//  Sets the current value for |content_type| for the specified URLs in the
	//  default scope. If both URLs are NULL, and the context is not incognito,
	//  the default value will be set. Pass CEF_CONTENT_SETTING_VALUE_DEFAULT for
	//  |value| to use the default value for this content type.
	//  WARNING: Incorrect usage of this function may cause instability or
	//  security issues in Chromium. Make sure that you first understand the
	//  potential impact of any changes to |content_type| by reviewing the related
	//  source code in Chromium. For example, if you plan to modify
	//  CEF_CONTENT_SETTING_TYPE_POPUPS, first review and understand the usage of
	//  ContentSettingsType::POPUPS in Chromium:
	//  https://source.chromium.org/search?q=ContentSettingsType::POPUPS
	SetContentSetting(requestingurl, toplevelurl string, contenttype TCefContentSettingTypes, value TCefContentSettingValues) // procedure
}

// TCefRequestContext Parent: TCefPreferenceManager
//
//	A request context provides request handling for a set of related browser or
//	URL request objects. A request context can be specified when creating a new
//	browser via the cef_browser_host_t static factory functions or when creating
//	a new URL request via the cef_urlrequest_t static factory functions. Browser
//	objects with different request contexts will never be hosted in the same
//	render process. Browser objects with the same request context may or may not
//	be hosted in the same render process depending on the process model. Browser
//	objects created indirectly via the JavaScript window.open function or
//	targeted links will share the same render process and the same request
//	context as the source browser. When running in single-process mode there is
//	only a single render process (the main process) and so all browsers created
//	in single-process mode will share the same request context. This will be the
//	first request context passed into a cef_browser_host_t static factory
//	function and all other request context objects will be ignored.
type TCefRequestContext struct {
	TCefPreferenceManager
}

// RequestContextRef -> ICefRequestContext
var RequestContextRef requestContext

// requestContext TCefRequestContext Ref
type requestContext uintptr

func (m *requestContext) UnWrap(data uintptr) ICefRequestContext {
	var resultCefRequestContext uintptr
	CEF().SysCallN(1282, uintptr(data), uintptr(unsafePointer(&resultCefRequestContext)))
	return AsCefRequestContext(resultCefRequestContext)
}

// Global
//
//	Returns the global context object.
func (m *requestContext) Global() ICefRequestContext {
	var resultCefRequestContext uintptr
	CEF().SysCallN(1269, uintptr(unsafePointer(&resultCefRequestContext)))
	return AsCefRequestContext(resultCefRequestContext)
}

// New
//
//	Creates a new context object with the specified |settings| and optional
//	|handler|.
func (m *requestContext) New(settings *TCefRequestContextSettings, handler ICefRequestContextHandler) ICefRequestContext {
	inArgs0 := settings.Pointer()
	var resultCefRequestContext uintptr
	CEF().SysCallN(1275, uintptr(unsafePointer(inArgs0)), GetObjectUintptr(handler), uintptr(unsafePointer(&resultCefRequestContext)))
	return AsCefRequestContext(resultCefRequestContext)
}

func (m *requestContext) New1(aCache, aAcceptLanguageList, aCookieableSchemesList string, aCookieableSchemesExcludeDefaults, aPersistSessionCookies, aPersistUserPreferences bool, handler ICefRequestContextHandler) ICefRequestContext {
	var resultCefRequestContext uintptr
	CEF().SysCallN(1276, PascalStr(aCache), PascalStr(aAcceptLanguageList), PascalStr(aCookieableSchemesList), PascalBool(aCookieableSchemesExcludeDefaults), PascalBool(aPersistSessionCookies), PascalBool(aPersistUserPreferences), GetObjectUintptr(handler), uintptr(unsafePointer(&resultCefRequestContext)))
	return AsCefRequestContext(resultCefRequestContext)
}

// Shared
//
//	Creates a new context object that shares storage with |other| and uses an
//	optional |handler|.
func (m *requestContext) Shared(other ICefRequestContext, handler ICefRequestContextHandler) ICefRequestContext {
	var resultCefRequestContext uintptr
	CEF().SysCallN(1281, GetObjectUintptr(other), GetObjectUintptr(handler), uintptr(unsafePointer(&resultCefRequestContext)))
	return AsCefRequestContext(resultCefRequestContext)
}

func (m *TCefRequestContext) IsSame(other ICefRequestContext) bool {
	r1 := CEF().SysCallN(1272, m.Instance(), GetObjectUintptr(other))
	return GoBool(r1)
}

func (m *TCefRequestContext) IsSharingWith(other ICefRequestContext) bool {
	r1 := CEF().SysCallN(1273, m.Instance(), GetObjectUintptr(other))
	return GoBool(r1)
}

func (m *TCefRequestContext) IsGlobal() bool {
	r1 := CEF().SysCallN(1271, m.Instance())
	return GoBool(r1)
}

func (m *TCefRequestContext) GetHandler() ICefRequestContextHandler {
	var resultCefRequestContextHandler uintptr
	CEF().SysCallN(1266, m.Instance(), uintptr(unsafePointer(&resultCefRequestContextHandler)))
	return AsCefRequestContextHandler(resultCefRequestContextHandler)
}

func (m *TCefRequestContext) GetCachePath() string {
	r1 := CEF().SysCallN(1261, m.Instance())
	return GoStr(r1)
}

func (m *TCefRequestContext) GetCookieManager(callback ICefCompletionCallback) ICefCookieManager {
	var resultCefCookieManager uintptr
	CEF().SysCallN(1263, m.Instance(), GetObjectUintptr(callback), uintptr(unsafePointer(&resultCefCookieManager)))
	return AsCefCookieManager(resultCefCookieManager)
}

func (m *TCefRequestContext) RegisterSchemeHandlerFactory(schemeName, domainName string, factory ICefSchemeHandlerFactory) bool {
	r1 := CEF().SysCallN(1277, m.Instance(), PascalStr(schemeName), PascalStr(domainName), GetObjectUintptr(factory))
	return GoBool(r1)
}

func (m *TCefRequestContext) ClearSchemeHandlerFactories() bool {
	r1 := CEF().SysCallN(1258, m.Instance())
	return GoBool(r1)
}

func (m *TCefRequestContext) DidLoadExtension(extensionid string) bool {
	r1 := CEF().SysCallN(1260, m.Instance(), PascalStr(extensionid))
	return GoBool(r1)
}

func (m *TCefRequestContext) HasExtension(extensionid string) bool {
	r1 := CEF().SysCallN(1270, m.Instance(), PascalStr(extensionid))
	return GoBool(r1)
}

func (m *TCefRequestContext) GetExtensions(extensionids IStringList) bool {
	r1 := CEF().SysCallN(1265, m.Instance(), GetObjectUintptr(extensionids))
	return GoBool(r1)
}

func (m *TCefRequestContext) GetExtension(extensionid string) ICefExtension {
	var resultCefExtension uintptr
	CEF().SysCallN(1264, m.Instance(), PascalStr(extensionid), uintptr(unsafePointer(&resultCefExtension)))
	return AsCefExtension(resultCefExtension)
}

func (m *TCefRequestContext) GetMediaRouter(callback ICefCompletionCallback) ICefMediaRouter {
	var resultCefMediaRouter uintptr
	CEF().SysCallN(1267, m.Instance(), GetObjectUintptr(callback), uintptr(unsafePointer(&resultCefMediaRouter)))
	return AsCefMediaRouter(resultCefMediaRouter)
}

func (m *TCefRequestContext) GetWebsiteSetting(requestingurl, toplevelurl string, contenttype TCefContentSettingTypes) ICefValue {
	var resultCefValue uintptr
	CEF().SysCallN(1268, m.Instance(), PascalStr(requestingurl), PascalStr(toplevelurl), uintptr(contenttype), uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *TCefRequestContext) GetContentSetting(requestingurl, toplevelurl string, contenttype TCefContentSettingTypes) TCefContentSettingValues {
	r1 := CEF().SysCallN(1262, m.Instance(), PascalStr(requestingurl), PascalStr(toplevelurl), uintptr(contenttype))
	return TCefContentSettingValues(r1)
}

func (m *TCefRequestContext) ClearCertificateExceptions(callback ICefCompletionCallback) {
	CEF().SysCallN(1256, m.Instance(), GetObjectUintptr(callback))
}

func (m *TCefRequestContext) ClearHttpAuthCredentials(callback ICefCompletionCallback) {
	CEF().SysCallN(1257, m.Instance(), GetObjectUintptr(callback))
}

func (m *TCefRequestContext) CloseAllConnections(callback ICefCompletionCallback) {
	CEF().SysCallN(1259, m.Instance(), GetObjectUintptr(callback))
}

func (m *TCefRequestContext) ResolveHost(origin string, callback ICefResolveCallback) {
	CEF().SysCallN(1278, m.Instance(), PascalStr(origin), GetObjectUintptr(callback))
}

func (m *TCefRequestContext) LoadExtension(rootdirectory string, manifest ICefDictionaryValue, handler ICefExtensionHandler) {
	CEF().SysCallN(1274, m.Instance(), PascalStr(rootdirectory), GetObjectUintptr(manifest), GetObjectUintptr(handler))
}

func (m *TCefRequestContext) SetWebsiteSetting(requestingurl, toplevelurl string, contenttype TCefContentSettingTypes, value ICefValue) {
	CEF().SysCallN(1280, m.Instance(), PascalStr(requestingurl), PascalStr(toplevelurl), uintptr(contenttype), GetObjectUintptr(value))
}

func (m *TCefRequestContext) SetContentSetting(requestingurl, toplevelurl string, contenttype TCefContentSettingTypes, value TCefContentSettingValues) {
	CEF().SysCallN(1279, m.Instance(), PascalStr(requestingurl), PascalStr(toplevelurl), uintptr(contenttype), uintptr(value))
}
