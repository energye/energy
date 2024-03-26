//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

// AsCoreWebView2 Convert a pointer object to an existing class object
func AsCoreWebView2(obj interface{}) ICoreWebView2 {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2 := new(TCoreWebView2)
	SetObjectInstance(coreWebView2, instance)
	return coreWebView2
}

// AsCoreWebView2AcceleratorKeyPressedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2AcceleratorKeyPressedEventArgs(obj interface{}) ICoreWebView2AcceleratorKeyPressedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2AcceleratorKeyPressedEventArgs := new(TCoreWebView2AcceleratorKeyPressedEventArgs)
	SetObjectInstance(coreWebView2AcceleratorKeyPressedEventArgs, instance)
	return coreWebView2AcceleratorKeyPressedEventArgs
}

// AsCoreWebView2BasicAuthenticationRequestedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2BasicAuthenticationRequestedEventArgs(obj interface{}) ICoreWebView2BasicAuthenticationRequestedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2BasicAuthenticationRequestedEventArgs := new(TCoreWebView2BasicAuthenticationRequestedEventArgs)
	SetObjectInstance(coreWebView2BasicAuthenticationRequestedEventArgs, instance)
	return coreWebView2BasicAuthenticationRequestedEventArgs
}

// AsCoreWebView2BasicAuthenticationResponse Convert a pointer object to an existing class object
func AsCoreWebView2BasicAuthenticationResponse(obj interface{}) ICoreWebView2BasicAuthenticationResponse {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2BasicAuthenticationResponse := new(TCoreWebView2BasicAuthenticationResponse)
	SetObjectInstance(coreWebView2BasicAuthenticationResponse, instance)
	return coreWebView2BasicAuthenticationResponse
}

// AsCoreWebView2BrowserExtension Convert a pointer object to an existing class object
func AsCoreWebView2BrowserExtension(obj interface{}) ICoreWebView2BrowserExtension {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2BrowserExtension := new(TCoreWebView2BrowserExtension)
	SetObjectInstance(coreWebView2BrowserExtension, instance)
	return coreWebView2BrowserExtension
}

// AsCoreWebView2BrowserExtensionEnableCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2BrowserExtensionEnableCompletedHandler(obj interface{}) ICoreWebView2BrowserExtensionEnableCompletedHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2BrowserExtensionEnableCompletedHandler := new(TCoreWebView2BrowserExtensionEnableCompletedHandler)
	SetObjectInstance(coreWebView2BrowserExtensionEnableCompletedHandler, instance)
	return coreWebView2BrowserExtensionEnableCompletedHandler
}

// AsCoreWebView2BrowserExtensionList Convert a pointer object to an existing class object
func AsCoreWebView2BrowserExtensionList(obj interface{}) ICoreWebView2BrowserExtensionList {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2BrowserExtensionList := new(TCoreWebView2BrowserExtensionList)
	SetObjectInstance(coreWebView2BrowserExtensionList, instance)
	return coreWebView2BrowserExtensionList
}

// AsCoreWebView2BrowserExtensionRemoveCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2BrowserExtensionRemoveCompletedHandler(obj interface{}) ICoreWebView2BrowserExtensionRemoveCompletedHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2BrowserExtensionRemoveCompletedHandler := new(TCoreWebView2BrowserExtensionRemoveCompletedHandler)
	SetObjectInstance(coreWebView2BrowserExtensionRemoveCompletedHandler, instance)
	return coreWebView2BrowserExtensionRemoveCompletedHandler
}

// AsCoreWebView2BrowserProcessExitedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2BrowserProcessExitedEventArgs(obj interface{}) ICoreWebView2BrowserProcessExitedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2BrowserProcessExitedEventArgs := new(TCoreWebView2BrowserProcessExitedEventArgs)
	SetObjectInstance(coreWebView2BrowserProcessExitedEventArgs, instance)
	return coreWebView2BrowserProcessExitedEventArgs
}

// AsCoreWebView2Certificate Convert a pointer object to an existing class object
func AsCoreWebView2Certificate(obj interface{}) ICoreWebView2Certificate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2Certificate := new(TCoreWebView2Certificate)
	SetObjectInstance(coreWebView2Certificate, instance)
	return coreWebView2Certificate
}

// AsCoreWebView2ClearBrowsingDataCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2ClearBrowsingDataCompletedHandler(obj interface{}) ICoreWebView2ClearBrowsingDataCompletedHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ClearBrowsingDataCompletedHandler := new(TCoreWebView2ClearBrowsingDataCompletedHandler)
	SetObjectInstance(coreWebView2ClearBrowsingDataCompletedHandler, instance)
	return coreWebView2ClearBrowsingDataCompletedHandler
}

// AsCoreWebView2ClientCertificate Convert a pointer object to an existing class object
func AsCoreWebView2ClientCertificate(obj interface{}) ICoreWebView2ClientCertificate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ClientCertificate := new(TCoreWebView2ClientCertificate)
	SetObjectInstance(coreWebView2ClientCertificate, instance)
	return coreWebView2ClientCertificate
}

// AsCoreWebView2ClientCertificateCollection Convert a pointer object to an existing class object
func AsCoreWebView2ClientCertificateCollection(obj interface{}) ICoreWebView2ClientCertificateCollection {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ClientCertificateCollection := new(TCoreWebView2ClientCertificateCollection)
	SetObjectInstance(coreWebView2ClientCertificateCollection, instance)
	return coreWebView2ClientCertificateCollection
}

// AsCoreWebView2ClientCertificateRequestedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2ClientCertificateRequestedEventArgs(obj interface{}) ICoreWebView2ClientCertificateRequestedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ClientCertificateRequestedEventArgs := new(TCoreWebView2ClientCertificateRequestedEventArgs)
	SetObjectInstance(coreWebView2ClientCertificateRequestedEventArgs, instance)
	return coreWebView2ClientCertificateRequestedEventArgs
}

// AsCoreWebView2CompositionController Convert a pointer object to an existing class object
func AsCoreWebView2CompositionController(obj interface{}) ICoreWebView2CompositionController {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2CompositionController := new(TCoreWebView2CompositionController)
	SetObjectInstance(coreWebView2CompositionController, instance)
	return coreWebView2CompositionController
}

// AsCoreWebView2ContentLoadingEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2ContentLoadingEventArgs(obj interface{}) ICoreWebView2ContentLoadingEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ContentLoadingEventArgs := new(TCoreWebView2ContentLoadingEventArgs)
	SetObjectInstance(coreWebView2ContentLoadingEventArgs, instance)
	return coreWebView2ContentLoadingEventArgs
}

// AsCoreWebView2ContextMenuItem Convert a pointer object to an existing class object
func AsCoreWebView2ContextMenuItem(obj interface{}) ICoreWebView2ContextMenuItem {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ContextMenuItem := new(TCoreWebView2ContextMenuItem)
	SetObjectInstance(coreWebView2ContextMenuItem, instance)
	return coreWebView2ContextMenuItem
}

// AsCoreWebView2ContextMenuItemCollection Convert a pointer object to an existing class object
func AsCoreWebView2ContextMenuItemCollection(obj interface{}) ICoreWebView2ContextMenuItemCollection {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ContextMenuItemCollection := new(TCoreWebView2ContextMenuItemCollection)
	SetObjectInstance(coreWebView2ContextMenuItemCollection, instance)
	return coreWebView2ContextMenuItemCollection
}

// AsCoreWebView2ContextMenuRequestedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2ContextMenuRequestedEventArgs(obj interface{}) ICoreWebView2ContextMenuRequestedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ContextMenuRequestedEventArgs := new(TCoreWebView2ContextMenuRequestedEventArgs)
	SetObjectInstance(coreWebView2ContextMenuRequestedEventArgs, instance)
	return coreWebView2ContextMenuRequestedEventArgs
}

// AsCoreWebView2ContextMenuTarget Convert a pointer object to an existing class object
func AsCoreWebView2ContextMenuTarget(obj interface{}) ICoreWebView2ContextMenuTarget {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ContextMenuTarget := new(TCoreWebView2ContextMenuTarget)
	SetObjectInstance(coreWebView2ContextMenuTarget, instance)
	return coreWebView2ContextMenuTarget
}

// AsCoreWebView2Controller Convert a pointer object to an existing class object
func AsCoreWebView2Controller(obj interface{}) ICoreWebView2Controller {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2Controller := new(TCoreWebView2Controller)
	SetObjectInstance(coreWebView2Controller, instance)
	return coreWebView2Controller
}

// AsCoreWebView2ControllerOptions Convert a pointer object to an existing class object
func AsCoreWebView2ControllerOptions(obj interface{}) ICoreWebView2ControllerOptions {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ControllerOptions := new(TCoreWebView2ControllerOptions)
	SetObjectInstance(coreWebView2ControllerOptions, instance)
	return coreWebView2ControllerOptions
}

// AsCoreWebView2Cookie Convert a pointer object to an existing class object
func AsCoreWebView2Cookie(obj interface{}) ICoreWebView2Cookie {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2Cookie := new(TCoreWebView2Cookie)
	SetObjectInstance(coreWebView2Cookie, instance)
	return coreWebView2Cookie
}

// AsCoreWebView2CookieList Convert a pointer object to an existing class object
func AsCoreWebView2CookieList(obj interface{}) ICoreWebView2CookieList {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2CookieList := new(TCoreWebView2CookieList)
	SetObjectInstance(coreWebView2CookieList, instance)
	return coreWebView2CookieList
}

// AsCoreWebView2CookieManager Convert a pointer object to an existing class object
func AsCoreWebView2CookieManager(obj interface{}) ICoreWebView2CookieManager {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2CookieManager := new(TCoreWebView2CookieManager)
	SetObjectInstance(coreWebView2CookieManager, instance)
	return coreWebView2CookieManager
}

// AsCoreWebView2CustomItemSelectedEventHandler Convert a pointer object to an existing class object
func AsCoreWebView2CustomItemSelectedEventHandler(obj interface{}) ICoreWebView2CustomItemSelectedEventHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2CustomItemSelectedEventHandler := new(TCoreWebView2CustomItemSelectedEventHandler)
	SetObjectInstance(coreWebView2CustomItemSelectedEventHandler, instance)
	return coreWebView2CustomItemSelectedEventHandler
}

// AsCoreWebView2DOMContentLoadedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2DOMContentLoadedEventArgs(obj interface{}) ICoreWebView2DOMContentLoadedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2DOMContentLoadedEventArgs := new(TCoreWebView2DOMContentLoadedEventArgs)
	SetObjectInstance(coreWebView2DOMContentLoadedEventArgs, instance)
	return coreWebView2DOMContentLoadedEventArgs
}

// AsCoreWebView2Deferral Convert a pointer object to an existing class object
func AsCoreWebView2Deferral(obj interface{}) ICoreWebView2Deferral {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2Deferral := new(TCoreWebView2Deferral)
	SetObjectInstance(coreWebView2Deferral, instance)
	return coreWebView2Deferral
}

// AsCoreWebView2DevToolsProtocolEventReceivedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2DevToolsProtocolEventReceivedEventArgs(obj interface{}) ICoreWebView2DevToolsProtocolEventReceivedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2DevToolsProtocolEventReceivedEventArgs := new(TCoreWebView2DevToolsProtocolEventReceivedEventArgs)
	SetObjectInstance(coreWebView2DevToolsProtocolEventReceivedEventArgs, instance)
	return coreWebView2DevToolsProtocolEventReceivedEventArgs
}

// AsCoreWebView2DownloadOperation Convert a pointer object to an existing class object
func AsCoreWebView2DownloadOperation(obj interface{}) ICoreWebView2DownloadOperation {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2DownloadOperation := new(TCoreWebView2DownloadOperation)
	SetObjectInstance(coreWebView2DownloadOperation, instance)
	return coreWebView2DownloadOperation
}

// AsCoreWebView2DownloadStartingEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2DownloadStartingEventArgs(obj interface{}) ICoreWebView2DownloadStartingEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2DownloadStartingEventArgs := new(TCoreWebView2DownloadStartingEventArgs)
	SetObjectInstance(coreWebView2DownloadStartingEventArgs, instance)
	return coreWebView2DownloadStartingEventArgs
}

// AsCoreWebView2Environment Convert a pointer object to an existing class object
func AsCoreWebView2Environment(obj interface{}) ICoreWebView2Environment {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2Environment := new(TCoreWebView2Environment)
	SetObjectInstance(coreWebView2Environment, instance)
	return coreWebView2Environment
}

// AsCoreWebView2ExecuteScriptResult Convert a pointer object to an existing class object
func AsCoreWebView2ExecuteScriptResult(obj interface{}) ICoreWebView2ExecuteScriptResult {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ExecuteScriptResult := new(TCoreWebView2ExecuteScriptResult)
	SetObjectInstance(coreWebView2ExecuteScriptResult, instance)
	return coreWebView2ExecuteScriptResult
}

// AsCoreWebView2Frame Convert a pointer object to an existing class object
func AsCoreWebView2Frame(obj interface{}) ICoreWebView2Frame {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2Frame := new(TCoreWebView2Frame)
	SetObjectInstance(coreWebView2Frame, instance)
	return coreWebView2Frame
}

// AsCoreWebView2FrameCreatedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2FrameCreatedEventArgs(obj interface{}) ICoreWebView2FrameCreatedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2FrameCreatedEventArgs := new(TCoreWebView2FrameCreatedEventArgs)
	SetObjectInstance(coreWebView2FrameCreatedEventArgs, instance)
	return coreWebView2FrameCreatedEventArgs
}

// AsCoreWebView2FrameInfo Convert a pointer object to an existing class object
func AsCoreWebView2FrameInfo(obj interface{}) ICoreWebView2FrameInfo {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2FrameInfo := new(TCoreWebView2FrameInfo)
	SetObjectInstance(coreWebView2FrameInfo, instance)
	return coreWebView2FrameInfo
}

// AsCoreWebView2FrameInfoCollection Convert a pointer object to an existing class object
func AsCoreWebView2FrameInfoCollection(obj interface{}) ICoreWebView2FrameInfoCollection {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2FrameInfoCollection := new(TCoreWebView2FrameInfoCollection)
	SetObjectInstance(coreWebView2FrameInfoCollection, instance)
	return coreWebView2FrameInfoCollection
}

// AsCoreWebView2FrameInfoCollectionIterator Convert a pointer object to an existing class object
func AsCoreWebView2FrameInfoCollectionIterator(obj interface{}) ICoreWebView2FrameInfoCollectionIterator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2FrameInfoCollectionIterator := new(TCoreWebView2FrameInfoCollectionIterator)
	SetObjectInstance(coreWebView2FrameInfoCollectionIterator, instance)
	return coreWebView2FrameInfoCollectionIterator
}

// AsCoreWebView2GetCookiesCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2GetCookiesCompletedHandler(obj interface{}) ICoreWebView2GetCookiesCompletedHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2GetCookiesCompletedHandler := new(TCoreWebView2GetCookiesCompletedHandler)
	SetObjectInstance(coreWebView2GetCookiesCompletedHandler, instance)
	return coreWebView2GetCookiesCompletedHandler
}

// AsCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler(obj interface{}) ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2GetNonDefaultPermissionSettingsCompletedHandler := new(TCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler)
	SetObjectInstance(coreWebView2GetNonDefaultPermissionSettingsCompletedHandler, instance)
	return coreWebView2GetNonDefaultPermissionSettingsCompletedHandler
}

// AsCoreWebView2HttpHeadersCollectionIterator Convert a pointer object to an existing class object
func AsCoreWebView2HttpHeadersCollectionIterator(obj interface{}) ICoreWebView2HttpHeadersCollectionIterator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2HttpHeadersCollectionIterator := new(TCoreWebView2HttpHeadersCollectionIterator)
	SetObjectInstance(coreWebView2HttpHeadersCollectionIterator, instance)
	return coreWebView2HttpHeadersCollectionIterator
}

// AsCoreWebView2HttpRequestHeaders Convert a pointer object to an existing class object
func AsCoreWebView2HttpRequestHeaders(obj interface{}) ICoreWebView2HttpRequestHeaders {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2HttpRequestHeaders := new(TCoreWebView2HttpRequestHeaders)
	SetObjectInstance(coreWebView2HttpRequestHeaders, instance)
	return coreWebView2HttpRequestHeaders
}

// AsCoreWebView2HttpResponseHeaders Convert a pointer object to an existing class object
func AsCoreWebView2HttpResponseHeaders(obj interface{}) ICoreWebView2HttpResponseHeaders {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2HttpResponseHeaders := new(TCoreWebView2HttpResponseHeaders)
	SetObjectInstance(coreWebView2HttpResponseHeaders, instance)
	return coreWebView2HttpResponseHeaders
}

// AsCoreWebView2LaunchingExternalUriSchemeEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2LaunchingExternalUriSchemeEventArgs(obj interface{}) ICoreWebView2LaunchingExternalUriSchemeEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2LaunchingExternalUriSchemeEventArgs := new(TCoreWebView2LaunchingExternalUriSchemeEventArgs)
	SetObjectInstance(coreWebView2LaunchingExternalUriSchemeEventArgs, instance)
	return coreWebView2LaunchingExternalUriSchemeEventArgs
}

// AsCoreWebView2MoveFocusRequestedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2MoveFocusRequestedEventArgs(obj interface{}) ICoreWebView2MoveFocusRequestedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2MoveFocusRequestedEventArgs := new(TCoreWebView2MoveFocusRequestedEventArgs)
	SetObjectInstance(coreWebView2MoveFocusRequestedEventArgs, instance)
	return coreWebView2MoveFocusRequestedEventArgs
}

// AsCoreWebView2NavigationCompletedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2NavigationCompletedEventArgs(obj interface{}) ICoreWebView2NavigationCompletedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2NavigationCompletedEventArgs := new(TCoreWebView2NavigationCompletedEventArgs)
	SetObjectInstance(coreWebView2NavigationCompletedEventArgs, instance)
	return coreWebView2NavigationCompletedEventArgs
}

// AsCoreWebView2NavigationStartingEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2NavigationStartingEventArgs(obj interface{}) ICoreWebView2NavigationStartingEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2NavigationStartingEventArgs := new(TCoreWebView2NavigationStartingEventArgs)
	SetObjectInstance(coreWebView2NavigationStartingEventArgs, instance)
	return coreWebView2NavigationStartingEventArgs
}

// AsCoreWebView2NewWindowRequestedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2NewWindowRequestedEventArgs(obj interface{}) ICoreWebView2NewWindowRequestedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2NewWindowRequestedEventArgs := new(TCoreWebView2NewWindowRequestedEventArgs)
	SetObjectInstance(coreWebView2NewWindowRequestedEventArgs, instance)
	return coreWebView2NewWindowRequestedEventArgs
}

// AsCoreWebView2ObjectCollectionView Convert a pointer object to an existing class object
func AsCoreWebView2ObjectCollectionView(obj interface{}) ICoreWebView2ObjectCollectionView {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ObjectCollectionView := new(TCoreWebView2ObjectCollectionView)
	SetObjectInstance(coreWebView2ObjectCollectionView, instance)
	return coreWebView2ObjectCollectionView
}

// AsCoreWebView2PermissionRequestedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2PermissionRequestedEventArgs(obj interface{}) ICoreWebView2PermissionRequestedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2PermissionRequestedEventArgs := new(TCoreWebView2PermissionRequestedEventArgs)
	SetObjectInstance(coreWebView2PermissionRequestedEventArgs, instance)
	return coreWebView2PermissionRequestedEventArgs
}

// AsCoreWebView2PermissionSetting Convert a pointer object to an existing class object
func AsCoreWebView2PermissionSetting(obj interface{}) ICoreWebView2PermissionSetting {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2PermissionSetting := new(TCoreWebView2PermissionSetting)
	SetObjectInstance(coreWebView2PermissionSetting, instance)
	return coreWebView2PermissionSetting
}

// AsCoreWebView2PermissionSettingCollectionView Convert a pointer object to an existing class object
func AsCoreWebView2PermissionSettingCollectionView(obj interface{}) ICoreWebView2PermissionSettingCollectionView {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2PermissionSettingCollectionView := new(TCoreWebView2PermissionSettingCollectionView)
	SetObjectInstance(coreWebView2PermissionSettingCollectionView, instance)
	return coreWebView2PermissionSettingCollectionView
}

// AsCoreWebView2PointerInfo Convert a pointer object to an existing class object
func AsCoreWebView2PointerInfo(obj interface{}) ICoreWebView2PointerInfo {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2PointerInfo := new(TCoreWebView2PointerInfo)
	SetObjectInstance(coreWebView2PointerInfo, instance)
	return coreWebView2PointerInfo
}

// AsCoreWebView2PrintCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2PrintCompletedHandler(obj interface{}) ICoreWebView2PrintCompletedHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2PrintCompletedHandler := new(TCoreWebView2PrintCompletedHandler)
	SetObjectInstance(coreWebView2PrintCompletedHandler, instance)
	return coreWebView2PrintCompletedHandler
}

// AsCoreWebView2PrintSettings Convert a pointer object to an existing class object
func AsCoreWebView2PrintSettings(obj interface{}) ICoreWebView2PrintSettings {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2PrintSettings := new(TCoreWebView2PrintSettings)
	SetObjectInstance(coreWebView2PrintSettings, instance)
	return coreWebView2PrintSettings
}

// AsCoreWebView2PrintToPdfCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2PrintToPdfCompletedHandler(obj interface{}) ICoreWebView2PrintToPdfCompletedHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2PrintToPdfCompletedHandler := new(TCoreWebView2PrintToPdfCompletedHandler)
	SetObjectInstance(coreWebView2PrintToPdfCompletedHandler, instance)
	return coreWebView2PrintToPdfCompletedHandler
}

// AsCoreWebView2PrintToPdfStreamCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2PrintToPdfStreamCompletedHandler(obj interface{}) ICoreWebView2PrintToPdfStreamCompletedHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2PrintToPdfStreamCompletedHandler := new(TCoreWebView2PrintToPdfStreamCompletedHandler)
	SetObjectInstance(coreWebView2PrintToPdfStreamCompletedHandler, instance)
	return coreWebView2PrintToPdfStreamCompletedHandler
}

// AsCoreWebView2ProcessExtendedInfo Convert a pointer object to an existing class object
func AsCoreWebView2ProcessExtendedInfo(obj interface{}) ICoreWebView2ProcessExtendedInfo {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ProcessExtendedInfo := new(TCoreWebView2ProcessExtendedInfo)
	SetObjectInstance(coreWebView2ProcessExtendedInfo, instance)
	return coreWebView2ProcessExtendedInfo
}

// AsCoreWebView2ProcessExtendedInfoCollection Convert a pointer object to an existing class object
func AsCoreWebView2ProcessExtendedInfoCollection(obj interface{}) ICoreWebView2ProcessExtendedInfoCollection {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ProcessExtendedInfoCollection := new(TCoreWebView2ProcessExtendedInfoCollection)
	SetObjectInstance(coreWebView2ProcessExtendedInfoCollection, instance)
	return coreWebView2ProcessExtendedInfoCollection
}

// AsCoreWebView2ProcessFailedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2ProcessFailedEventArgs(obj interface{}) ICoreWebView2ProcessFailedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ProcessFailedEventArgs := new(TCoreWebView2ProcessFailedEventArgs)
	SetObjectInstance(coreWebView2ProcessFailedEventArgs, instance)
	return coreWebView2ProcessFailedEventArgs
}

// AsCoreWebView2ProcessInfo Convert a pointer object to an existing class object
func AsCoreWebView2ProcessInfo(obj interface{}) ICoreWebView2ProcessInfo {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ProcessInfo := new(TCoreWebView2ProcessInfo)
	SetObjectInstance(coreWebView2ProcessInfo, instance)
	return coreWebView2ProcessInfo
}

// AsCoreWebView2ProcessInfoCollection Convert a pointer object to an existing class object
func AsCoreWebView2ProcessInfoCollection(obj interface{}) ICoreWebView2ProcessInfoCollection {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ProcessInfoCollection := new(TCoreWebView2ProcessInfoCollection)
	SetObjectInstance(coreWebView2ProcessInfoCollection, instance)
	return coreWebView2ProcessInfoCollection
}

// AsCoreWebView2Profile Convert a pointer object to an existing class object
func AsCoreWebView2Profile(obj interface{}) ICoreWebView2Profile {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2Profile := new(TCoreWebView2Profile)
	SetObjectInstance(coreWebView2Profile, instance)
	return coreWebView2Profile
}

// AsCoreWebView2ProfileAddBrowserExtensionCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2ProfileAddBrowserExtensionCompletedHandler(obj interface{}) ICoreWebView2ProfileAddBrowserExtensionCompletedHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ProfileAddBrowserExtensionCompletedHandler := new(TCoreWebView2ProfileAddBrowserExtensionCompletedHandler)
	SetObjectInstance(coreWebView2ProfileAddBrowserExtensionCompletedHandler, instance)
	return coreWebView2ProfileAddBrowserExtensionCompletedHandler
}

// AsCoreWebView2ProfileGetBrowserExtensionsCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2ProfileGetBrowserExtensionsCompletedHandler(obj interface{}) ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ProfileGetBrowserExtensionsCompletedHandler := new(TCoreWebView2ProfileGetBrowserExtensionsCompletedHandler)
	SetObjectInstance(coreWebView2ProfileGetBrowserExtensionsCompletedHandler, instance)
	return coreWebView2ProfileGetBrowserExtensionsCompletedHandler
}

// AsCoreWebView2ScriptDialogOpeningEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2ScriptDialogOpeningEventArgs(obj interface{}) ICoreWebView2ScriptDialogOpeningEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ScriptDialogOpeningEventArgs := new(TCoreWebView2ScriptDialogOpeningEventArgs)
	SetObjectInstance(coreWebView2ScriptDialogOpeningEventArgs, instance)
	return coreWebView2ScriptDialogOpeningEventArgs
}

// AsCoreWebView2ScriptException Convert a pointer object to an existing class object
func AsCoreWebView2ScriptException(obj interface{}) ICoreWebView2ScriptException {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ScriptException := new(TCoreWebView2ScriptException)
	SetObjectInstance(coreWebView2ScriptException, instance)
	return coreWebView2ScriptException
}

// AsCoreWebView2ServerCertificateErrorDetectedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2ServerCertificateErrorDetectedEventArgs(obj interface{}) ICoreWebView2ServerCertificateErrorDetectedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2ServerCertificateErrorDetectedEventArgs := new(TCoreWebView2ServerCertificateErrorDetectedEventArgs)
	SetObjectInstance(coreWebView2ServerCertificateErrorDetectedEventArgs, instance)
	return coreWebView2ServerCertificateErrorDetectedEventArgs
}

// AsCoreWebView2SetPermissionStateCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2SetPermissionStateCompletedHandler(obj interface{}) ICoreWebView2SetPermissionStateCompletedHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2SetPermissionStateCompletedHandler := new(TCoreWebView2SetPermissionStateCompletedHandler)
	SetObjectInstance(coreWebView2SetPermissionStateCompletedHandler, instance)
	return coreWebView2SetPermissionStateCompletedHandler
}

// AsCoreWebView2Settings Convert a pointer object to an existing class object
func AsCoreWebView2Settings(obj interface{}) ICoreWebView2Settings {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2Settings := new(TCoreWebView2Settings)
	SetObjectInstance(coreWebView2Settings, instance)
	return coreWebView2Settings
}

// AsCoreWebView2SharedBuffer Convert a pointer object to an existing class object
func AsCoreWebView2SharedBuffer(obj interface{}) ICoreWebView2SharedBuffer {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2SharedBuffer := new(TCoreWebView2SharedBuffer)
	SetObjectInstance(coreWebView2SharedBuffer, instance)
	return coreWebView2SharedBuffer
}

// AsCoreWebView2SourceChangedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2SourceChangedEventArgs(obj interface{}) ICoreWebView2SourceChangedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2SourceChangedEventArgs := new(TCoreWebView2SourceChangedEventArgs)
	SetObjectInstance(coreWebView2SourceChangedEventArgs, instance)
	return coreWebView2SourceChangedEventArgs
}

// AsCoreWebView2StringCollection Convert a pointer object to an existing class object
func AsCoreWebView2StringCollection(obj interface{}) ICoreWebView2StringCollection {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2StringCollection := new(TCoreWebView2StringCollection)
	SetObjectInstance(coreWebView2StringCollection, instance)
	return coreWebView2StringCollection
}

// AsCoreWebView2TrySuspendCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2TrySuspendCompletedHandler(obj interface{}) ICoreWebView2TrySuspendCompletedHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2TrySuspendCompletedHandler := new(TCoreWebView2TrySuspendCompletedHandler)
	SetObjectInstance(coreWebView2TrySuspendCompletedHandler, instance)
	return coreWebView2TrySuspendCompletedHandler
}

// AsCoreWebView2WebMessageReceivedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2WebMessageReceivedEventArgs(obj interface{}) ICoreWebView2WebMessageReceivedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2WebMessageReceivedEventArgs := new(TCoreWebView2WebMessageReceivedEventArgs)
	SetObjectInstance(coreWebView2WebMessageReceivedEventArgs, instance)
	return coreWebView2WebMessageReceivedEventArgs
}

// AsCoreWebView2WebResourceRequestRef Convert a pointer object to an existing class object
func AsCoreWebView2WebResourceRequestRef(obj interface{}) ICoreWebView2WebResourceRequestRef {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2WebResourceRequestRef := new(TCoreWebView2WebResourceRequestRef)
	SetObjectInstance(coreWebView2WebResourceRequestRef, instance)
	return coreWebView2WebResourceRequestRef
}

// AsCoreWebView2WebResourceRequestedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2WebResourceRequestedEventArgs(obj interface{}) ICoreWebView2WebResourceRequestedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2WebResourceRequestedEventArgs := new(TCoreWebView2WebResourceRequestedEventArgs)
	SetObjectInstance(coreWebView2WebResourceRequestedEventArgs, instance)
	return coreWebView2WebResourceRequestedEventArgs
}

// AsCoreWebView2WebResourceResponse Convert a pointer object to an existing class object
func AsCoreWebView2WebResourceResponse(obj interface{}) ICoreWebView2WebResourceResponse {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2WebResourceResponse := new(TCoreWebView2WebResourceResponse)
	SetObjectInstance(coreWebView2WebResourceResponse, instance)
	return coreWebView2WebResourceResponse
}

// AsCoreWebView2WebResourceResponseReceivedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2WebResourceResponseReceivedEventArgs(obj interface{}) ICoreWebView2WebResourceResponseReceivedEventArgs {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2WebResourceResponseReceivedEventArgs := new(TCoreWebView2WebResourceResponseReceivedEventArgs)
	SetObjectInstance(coreWebView2WebResourceResponseReceivedEventArgs, instance)
	return coreWebView2WebResourceResponseReceivedEventArgs
}

// AsCoreWebView2WebResourceResponseView Convert a pointer object to an existing class object
func AsCoreWebView2WebResourceResponseView(obj interface{}) ICoreWebView2WebResourceResponseView {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2WebResourceResponseView := new(TCoreWebView2WebResourceResponseView)
	SetObjectInstance(coreWebView2WebResourceResponseView, instance)
	return coreWebView2WebResourceResponseView
}

// AsCoreWebView2WebResourceResponseViewGetContentCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2WebResourceResponseViewGetContentCompletedHandler(obj interface{}) ICoreWebView2WebResourceResponseViewGetContentCompletedHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2WebResourceResponseViewGetContentCompletedHandler := new(TCoreWebView2WebResourceResponseViewGetContentCompletedHandler)
	SetObjectInstance(coreWebView2WebResourceResponseViewGetContentCompletedHandler, instance)
	return coreWebView2WebResourceResponseViewGetContentCompletedHandler
}

// AsCoreWebView2WindowFeatures Convert a pointer object to an existing class object
func AsCoreWebView2WindowFeatures(obj interface{}) ICoreWebView2WindowFeatures {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coreWebView2WindowFeatures := new(TCoreWebView2WindowFeatures)
	SetObjectInstance(coreWebView2WindowFeatures, instance)
	return coreWebView2WindowFeatures
}

// AsWVBrowser Convert a pointer object to an existing class object
func AsWVBrowser(obj interface{}) IWVBrowser {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wVBrowser := new(TWVBrowser)
	SetObjectInstance(wVBrowser, instance)
	return wVBrowser
}

// AsWVBrowserBase Convert a pointer object to an existing class object
func AsWVBrowserBase(obj interface{}) IWVBrowserBase {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wVBrowserBase := new(TWVBrowserBase)
	SetObjectInstance(wVBrowserBase, instance)
	return wVBrowserBase
}

// AsWVLoader Convert a pointer object to an existing class object
func AsWVLoader(obj interface{}) IWVLoader {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wVLoader := new(TWVLoader)
	SetObjectInstance(wVLoader, instance)
	return wVLoader
}

// AsWVProxySettings Convert a pointer object to an existing class object
func AsWVProxySettings(obj interface{}) IWVProxySettings {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wVProxySettings := new(TWVProxySettings)
	SetObjectInstance(wVProxySettings, instance)
	return wVProxySettings
}

// AsWVWinControl Convert a pointer object to an existing class object
func AsWVWinControl(obj interface{}) IWVWinControl {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wVWinControl := new(TWVWinControl)
	SetObjectInstance(wVWinControl, instance)
	return wVWinControl
}

// AsWVWindowParent Convert a pointer object to an existing class object
func AsWVWindowParent(obj interface{}) IWVWindowParent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wVWindowParent := new(TWVWindowParent)
	SetObjectInstance(wVWindowParent, instance)
	return wVWindowParent
}
