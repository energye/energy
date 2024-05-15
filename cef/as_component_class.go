//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

// AsAccessibilityHandler Convert a pointer object to an existing class object
func AsAccessibilityHandler(obj interface{}) IAccessibilityHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	accessibilityHandler := new(TAccessibilityHandler)
	SetObjectInstance(accessibilityHandler, instance)
	return accessibilityHandler
}

// AsBufferPanel Convert a pointer object to an existing class object
func AsBufferPanel(obj interface{}) IBufferPanel {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	bufferPanel := new(TBufferPanel)
	SetObjectInstance(bufferPanel, instance)
	return bufferPanel
}

// AsCEFAccessibilityHandler Convert a pointer object to an existing class object
func AsCEFAccessibilityHandler(obj interface{}) ICEFAccessibilityHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFAccessibilityHandler := new(TCEFAccessibilityHandler)
	SetObjectInstance(cEFAccessibilityHandler, instance)
	return cEFAccessibilityHandler
}

// AsCEFBaseScopedWrapperRef Convert a pointer object to an existing class object
func AsCEFBaseScopedWrapperRef(obj interface{}) ICEFBaseScopedWrapperRef {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFBaseScopedWrapperRef := new(TCEFBaseScopedWrapperRef)
	SetObjectInstance(cEFBaseScopedWrapperRef, instance)
	return cEFBaseScopedWrapperRef
}

// AsCEFBitmapBitBuffer Convert a pointer object to an existing class object
func AsCEFBitmapBitBuffer(obj interface{}) ICEFBitmapBitBuffer {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFBitmapBitBuffer := new(TCEFBitmapBitBuffer)
	SetObjectInstance(cEFBitmapBitBuffer, instance)
	return cEFBitmapBitBuffer
}

// AsCEFBrowserViewComponent Convert a pointer object to an existing class object
func AsCEFBrowserViewComponent(obj interface{}) ICEFBrowserViewComponent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFBrowserViewComponent := new(TCEFBrowserViewComponent)
	SetObjectInstance(cEFBrowserViewComponent, instance)
	return cEFBrowserViewComponent
}

// AsCEFButtonComponent Convert a pointer object to an existing class object
func AsCEFButtonComponent(obj interface{}) ICEFButtonComponent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFButtonComponent := new(TCEFButtonComponent)
	SetObjectInstance(cEFButtonComponent, instance)
	return cEFButtonComponent
}

// AsCefDevToolsMessageObserver Convert a pointer object to an existing class object
func AsCefDevToolsMessageObserver(obj interface{}) ICefDevToolsMessageObserver {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefDevToolsMessageObserver := new(TCefDevToolsMessageObserver)
	SetObjectInstance(cefDevToolsMessageObserver, instance)
	return cefDevToolsMessageObserver
}

// AsCEFFileDialogInfo Convert a pointer object to an existing class object
func AsCEFFileDialogInfo(obj interface{}) ICEFFileDialogInfo {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFFileDialogInfo := new(TCEFFileDialogInfo)
	SetObjectInstance(cEFFileDialogInfo, instance)
	return cEFFileDialogInfo
}

// AsCEFJson Convert a pointer object to an existing class object
func AsCEFJson(obj interface{}) ICEFJson {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFJson := new(TCEFJson)
	SetObjectInstance(cEFJson, instance)
	return cEFJson
}

// AsCEFLabelButtonComponent Convert a pointer object to an existing class object
func AsCEFLabelButtonComponent(obj interface{}) ICEFLabelButtonComponent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFLabelButtonComponent := new(TCEFLabelButtonComponent)
	SetObjectInstance(cEFLabelButtonComponent, instance)
	return cEFLabelButtonComponent
}

// AsCEFLinkedWinControlBase Convert a pointer object to an existing class object
func AsCEFLinkedWinControlBase(obj interface{}) ICEFLinkedWinControlBase {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFLinkedWinControlBase := new(TCEFLinkedWinControlBase)
	SetObjectInstance(cEFLinkedWinControlBase, instance)
	return cEFLinkedWinControlBase
}

// AsCEFLinkedWindowParent Convert a pointer object to an existing class object
func AsCEFLinkedWindowParent(obj interface{}) ICEFLinkedWindowParent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFLinkedWindowParent := new(TCEFLinkedWindowParent)
	SetObjectInstance(cEFLinkedWindowParent, instance)
	return cEFLinkedWindowParent
}

// AsCEFMenuButtonComponent Convert a pointer object to an existing class object
func AsCEFMenuButtonComponent(obj interface{}) ICEFMenuButtonComponent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFMenuButtonComponent := new(TCEFMenuButtonComponent)
	SetObjectInstance(cEFMenuButtonComponent, instance)
	return cEFMenuButtonComponent
}

// AsCEFPanelComponent Convert a pointer object to an existing class object
func AsCEFPanelComponent(obj interface{}) ICEFPanelComponent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFPanelComponent := new(TCEFPanelComponent)
	SetObjectInstance(cEFPanelComponent, instance)
	return cEFPanelComponent
}

// AsCEFScrollViewComponent Convert a pointer object to an existing class object
func AsCEFScrollViewComponent(obj interface{}) ICEFScrollViewComponent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFScrollViewComponent := new(TCEFScrollViewComponent)
	SetObjectInstance(cEFScrollViewComponent, instance)
	return cEFScrollViewComponent
}

// AsCEFSentinel Convert a pointer object to an existing class object
func AsCEFSentinel(obj interface{}) ICEFSentinel {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFSentinel := new(TCEFSentinel)
	SetObjectInstance(cEFSentinel, instance)
	return cEFSentinel
}

// AsCEFServerComponent Convert a pointer object to an existing class object
func AsCEFServerComponent(obj interface{}) ICEFServerComponent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFServerComponent := new(TCEFServerComponent)
	SetObjectInstance(cEFServerComponent, instance)
	return cEFServerComponent
}

// AsCEFServerHandler Convert a pointer object to an existing class object
func AsCEFServerHandler(obj interface{}) ICEFServerHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFServerHandler := new(TCEFServerHandler)
	SetObjectInstance(cEFServerHandler, instance)
	return cEFServerHandler
}

// AsCEFServer Convert a pointer object to an existing class object
func AsCEFServer(obj interface{}) ICEFServer {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFServer := new(TCEFServer)
	SetObjectInstance(cEFServer, instance)
	return cEFServer
}

// AsCEFTextfieldComponent Convert a pointer object to an existing class object
func AsCEFTextfieldComponent(obj interface{}) ICEFTextfieldComponent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFTextfieldComponent := new(TCEFTextfieldComponent)
	SetObjectInstance(cEFTextfieldComponent, instance)
	return cEFTextfieldComponent
}

// AsCEFUrlRequestClientComponent Convert a pointer object to an existing class object
func AsCEFUrlRequestClientComponent(obj interface{}) ICEFUrlRequestClientComponent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFUrlRequestClientComponent := new(TCEFUrlRequestClientComponent)
	SetObjectInstance(cEFUrlRequestClientComponent, instance)
	return cEFUrlRequestClientComponent
}

// AsCEFViewComponent Convert a pointer object to an existing class object
func AsCEFViewComponent(obj interface{}) ICEFViewComponent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFViewComponent := new(TCEFViewComponent)
	SetObjectInstance(cEFViewComponent, instance)
	return cEFViewComponent
}

// AsCEFWinControl Convert a pointer object to an existing class object
func AsCEFWinControl(obj interface{}) ICEFWinControl {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFWinControl := new(TCEFWinControl)
	SetObjectInstance(cEFWinControl, instance)
	return cEFWinControl
}

// AsCEFWindowComponent Convert a pointer object to an existing class object
func AsCEFWindowComponent(obj interface{}) ICEFWindowComponent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFWindowComponent := new(TCEFWindowComponent)
	SetObjectInstance(cEFWindowComponent, instance)
	return cEFWindowComponent
}

// AsCEFWindowParent Convert a pointer object to an existing class object
func AsCEFWindowParent(obj interface{}) ICEFWindowParent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFWindowParent := new(TCEFWindowParent)
	SetObjectInstance(cEFWindowParent, instance)
	return cEFWindowParent
}

// AsCEFWorkScheduler Convert a pointer object to an existing class object
func AsCEFWorkScheduler(obj interface{}) ICEFWorkScheduler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cEFWorkScheduler := new(TCEFWorkScheduler)
	SetObjectInstance(cEFWorkScheduler, instance)
	return cEFWorkScheduler
}

// AsCefX509Certificate Convert a pointer object to an existing class object
func AsCefX509Certificate(obj interface{}) ICefX509Certificate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefX509Certificate := new(TCefX509Certificate)
	SetObjectInstance(cefX509Certificate, instance)
	return cefX509Certificate
}

// AsCefApplication Convert a pointer object to an existing class object
func AsCefApplication(obj interface{}) ICefApplication {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefApplication := new(TCefApplication)
	SetObjectInstance(cefApplication, instance)
	return cefApplication
}

// AsCefApplicationCore Convert a pointer object to an existing class object
func AsCefApplicationCore(obj interface{}) ICefApplicationCore {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefApplicationCore := new(TCefApplicationCore)
	SetObjectInstance(cefApplicationCore, instance)
	return cefApplicationCore
}

// AsCefAudioHandler Convert a pointer object to an existing class object
func AsCefAudioHandler(obj interface{}) ICefAudioHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefAudioHandler := new(TCefAudioHandler)
	SetObjectInstance(cefAudioHandler, instance)
	return cefAudioHandler
}

// AsCefAuthCallback Convert a pointer object to an existing class object
func AsCefAuthCallback(obj interface{}) ICefAuthCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefAuthCallback := new(TCefAuthCallback)
	SetObjectInstance(cefAuthCallback, instance)
	return cefAuthCallback
}

// AsCefBaseRefCountedOwn Convert a pointer object to an existing class object
func AsCefBaseRefCountedOwn(obj interface{}) ICefBaseRefCountedOwn {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefBaseRefCountedOwn := new(TCefBaseRefCountedOwn)
	SetObjectInstance(cefBaseRefCountedOwn, instance)
	return cefBaseRefCountedOwn
}

// AsCefBaseRefCounted Convert a pointer object to an existing class object
func AsCefBaseRefCounted(obj interface{}) ICefBaseRefCounted {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefBaseRefCounted := new(TCefBaseRefCounted)
	SetObjectInstance(cefBaseRefCounted, instance)
	return cefBaseRefCounted
}

// AsCefBeforeDownloadCallback Convert a pointer object to an existing class object
func AsCefBeforeDownloadCallback(obj interface{}) ICefBeforeDownloadCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefBeforeDownloadCallback := new(TCefBeforeDownloadCallback)
	SetObjectInstance(cefBeforeDownloadCallback, instance)
	return cefBeforeDownloadCallback
}

// AsCefBinaryValue Convert a pointer object to an existing class object
func AsCefBinaryValue(obj interface{}) ICefBinaryValue {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefBinaryValue := new(TCefBinaryValue)
	SetObjectInstance(cefBinaryValue, instance)
	return cefBinaryValue
}

// AsCefBoxLayout Convert a pointer object to an existing class object
func AsCefBoxLayout(obj interface{}) ICefBoxLayout {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefBoxLayout := new(TCefBoxLayout)
	SetObjectInstance(cefBoxLayout, instance)
	return cefBoxLayout
}

// AsCefBrowserHost Convert a pointer object to an existing class object
func AsCefBrowserHost(obj interface{}) ICefBrowserHost {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefBrowserHost := new(TCefBrowserHost)
	SetObjectInstance(cefBrowserHost, instance)
	return cefBrowserHost
}

// AsCefBrowserNavigationTask Convert a pointer object to an existing class object
func AsCefBrowserNavigationTask(obj interface{}) ICefBrowserNavigationTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefBrowserNavigationTask := new(TCefBrowserNavigationTask)
	SetObjectInstance(cefBrowserNavigationTask, instance)
	return cefBrowserNavigationTask
}

// AsCefBrowser Convert a pointer object to an existing class object
func AsCefBrowser(obj interface{}) ICefBrowser {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefBrowser := new(TCefBrowser)
	SetObjectInstance(cefBrowser, instance)
	return cefBrowser
}

// AsCefBrowserViewDelegate Convert a pointer object to an existing class object
func AsCefBrowserViewDelegate(obj interface{}) ICefBrowserViewDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefBrowserViewDelegate := new(TCefBrowserViewDelegate)
	SetObjectInstance(cefBrowserViewDelegate, instance)
	return cefBrowserViewDelegate
}

// AsCefBrowserView Convert a pointer object to an existing class object
func AsCefBrowserView(obj interface{}) ICefBrowserView {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefBrowserView := new(TCefBrowserView)
	SetObjectInstance(cefBrowserView, instance)
	return cefBrowserView
}

// AsCefButtonDelegate Convert a pointer object to an existing class object
func AsCefButtonDelegate(obj interface{}) ICefButtonDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefButtonDelegate := new(TCefButtonDelegate)
	SetObjectInstance(cefButtonDelegate, instance)
	return cefButtonDelegate
}

// AsCefButton Convert a pointer object to an existing class object
func AsCefButton(obj interface{}) ICefButton {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefButton := new(TCefButton)
	SetObjectInstance(cefButton, instance)
	return cefButton
}

// AsCefBytesWriteHandler Convert a pointer object to an existing class object
func AsCefBytesWriteHandler(obj interface{}) ICefBytesWriteHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefBytesWriteHandler := new(TCefBytesWriteHandler)
	SetObjectInstance(cefBytesWriteHandler, instance)
	return cefBytesWriteHandler
}

// AsCefCallback Convert a pointer object to an existing class object
func AsCefCallback(obj interface{}) ICefCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCallback := new(TCefCallback)
	SetObjectInstance(cefCallback, instance)
	return cefCallback
}

// AsCefClient Convert a pointer object to an existing class object
func AsCefClient(obj interface{}) ICefClient {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefClient := new(TCefClient)
	SetObjectInstance(cefClient, instance)
	return cefClient
}

// AsCefCommandHandler Convert a pointer object to an existing class object
func AsCefCommandHandler(obj interface{}) ICefCommandHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCommandHandler := new(TCefCommandHandler)
	SetObjectInstance(cefCommandHandler, instance)
	return cefCommandHandler
}

// AsCefCommandLine Convert a pointer object to an existing class object
func AsCefCommandLine(obj interface{}) ICefCommandLine {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCommandLine := new(TCefCommandLine)
	SetObjectInstance(cefCommandLine, instance)
	return cefCommandLine
}

// AsCefCompletionCallback Convert a pointer object to an existing class object
func AsCefCompletionCallback(obj interface{}) ICefCompletionCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCompletionCallback := new(TCefCompletionCallback)
	SetObjectInstance(cefCompletionCallback, instance)
	return cefCompletionCallback
}

// AsCefContextMenuHandler Convert a pointer object to an existing class object
func AsCefContextMenuHandler(obj interface{}) ICefContextMenuHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefContextMenuHandler := new(TCefContextMenuHandler)
	SetObjectInstance(cefContextMenuHandler, instance)
	return cefContextMenuHandler
}

// AsCefContextMenuParams Convert a pointer object to an existing class object
func AsCefContextMenuParams(obj interface{}) ICefContextMenuParams {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefContextMenuParams := new(TCefContextMenuParams)
	SetObjectInstance(cefContextMenuParams, instance)
	return cefContextMenuParams
}

// AsCefCookieManager Convert a pointer object to an existing class object
func AsCefCookieManager(obj interface{}) ICefCookieManager {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCookieManager := new(TCefCookieManager)
	SetObjectInstance(cefCookieManager, instance)
	return cefCookieManager
}

// AsCefCookieVisitor Convert a pointer object to an existing class object
func AsCefCookieVisitor(obj interface{}) ICefCookieVisitor {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCookieVisitor := new(TCefCookieVisitor)
	SetObjectInstance(cefCookieVisitor, instance)
	return cefCookieVisitor
}

// AsCefCreateCustomViewTask Convert a pointer object to an existing class object
func AsCefCreateCustomViewTask(obj interface{}) ICefCreateCustomViewTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCreateCustomViewTask := new(TCefCreateCustomViewTask)
	SetObjectInstance(cefCreateCustomViewTask, instance)
	return cefCreateCustomViewTask
}

// AsCefCustomCompletionCallback Convert a pointer object to an existing class object
func AsCefCustomCompletionCallback(obj interface{}) ICefCustomCompletionCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCustomCompletionCallback := new(TCefCustomCompletionCallback)
	SetObjectInstance(cefCustomCompletionCallback, instance)
	return cefCustomCompletionCallback
}

// AsCefCustomCookieVisitor Convert a pointer object to an existing class object
func AsCefCustomCookieVisitor(obj interface{}) ICefCustomCookieVisitor {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCustomCookieVisitor := new(TCefCustomCookieVisitor)
	SetObjectInstance(cefCustomCookieVisitor, instance)
	return cefCustomCookieVisitor
}

// AsCefCustomDeleteCookiesCallback Convert a pointer object to an existing class object
func AsCefCustomDeleteCookiesCallback(obj interface{}) ICefCustomDeleteCookiesCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCustomDeleteCookiesCallback := new(TCefCustomDeleteCookiesCallback)
	SetObjectInstance(cefCustomDeleteCookiesCallback, instance)
	return cefCustomDeleteCookiesCallback
}

// AsCefCustomDownloadImageCallback Convert a pointer object to an existing class object
func AsCefCustomDownloadImageCallback(obj interface{}) ICefCustomDownloadImageCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCustomDownloadImageCallback := new(TCefCustomDownloadImageCallback)
	SetObjectInstance(cefCustomDownloadImageCallback, instance)
	return cefCustomDownloadImageCallback
}

// AsCefCustomMediaRouteCreateCallback Convert a pointer object to an existing class object
func AsCefCustomMediaRouteCreateCallback(obj interface{}) ICefCustomMediaRouteCreateCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCustomMediaRouteCreateCallback := new(TCefCustomMediaRouteCreateCallback)
	SetObjectInstance(cefCustomMediaRouteCreateCallback, instance)
	return cefCustomMediaRouteCreateCallback
}

// AsCefCustomMediaSinkDeviceInfoCallback Convert a pointer object to an existing class object
func AsCefCustomMediaSinkDeviceInfoCallback(obj interface{}) ICefCustomMediaSinkDeviceInfoCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCustomMediaSinkDeviceInfoCallback := new(TCefCustomMediaSinkDeviceInfoCallback)
	SetObjectInstance(cefCustomMediaSinkDeviceInfoCallback, instance)
	return cefCustomMediaSinkDeviceInfoCallback
}

// AsCefCustomPDFPrintCallBack Convert a pointer object to an existing class object
func AsCefCustomPDFPrintCallBack(obj interface{}) ICefCustomPDFPrintCallBack {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCustomPDFPrintCallBack := new(TCefCustomPDFPrintCallBack)
	SetObjectInstance(cefCustomPDFPrintCallBack, instance)
	return cefCustomPDFPrintCallBack
}

// AsCefCustomRenderProcessHandler Convert a pointer object to an existing class object
func AsCefCustomRenderProcessHandler(obj interface{}) ICefCustomRenderProcessHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCustomRenderProcessHandler := new(TCefCustomRenderProcessHandler)
	SetObjectInstance(cefCustomRenderProcessHandler, instance)
	return cefCustomRenderProcessHandler
}

// AsCefCustomResolveCallback Convert a pointer object to an existing class object
func AsCefCustomResolveCallback(obj interface{}) ICefCustomResolveCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCustomResolveCallback := new(TCefCustomResolveCallback)
	SetObjectInstance(cefCustomResolveCallback, instance)
	return cefCustomResolveCallback
}

// AsCefCustomSetCookieCallback Convert a pointer object to an existing class object
func AsCefCustomSetCookieCallback(obj interface{}) ICefCustomSetCookieCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCustomSetCookieCallback := new(TCefCustomSetCookieCallback)
	SetObjectInstance(cefCustomSetCookieCallback, instance)
	return cefCustomSetCookieCallback
}

// AsCefCustomStreamReader Convert a pointer object to an existing class object
func AsCefCustomStreamReader(obj interface{}) ICefCustomStreamReader {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCustomStreamReader := new(TCefCustomStreamReader)
	SetObjectInstance(cefCustomStreamReader, instance)
	return cefCustomStreamReader
}

// AsCefCustomStringList Convert a pointer object to an existing class object
func AsCefCustomStringList(obj interface{}) ICefCustomStringList {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefCustomStringList := new(TCefCustomStringList)
	SetObjectInstance(cefCustomStringList, instance)
	return cefCustomStringList
}

// AsCefStringMap Convert a pointer object to an existing class object
func AsCefStringMap(obj interface{}) ICefStringMap {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefStringMap := new(TCefStringMap)
	SetObjectInstance(cefStringMap, instance)
	return cefStringMap
}

// AsCefStringMultimap Convert a pointer object to an existing class object
func AsCefStringMultimap(obj interface{}) ICefStringMultimap {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefStringMultimap := new(TCefStringMultimap)
	SetObjectInstance(cefStringMultimap, instance)
	return cefStringMultimap
}

// AsCefDeleteCookiesCallback Convert a pointer object to an existing class object
func AsCefDeleteCookiesCallback(obj interface{}) ICefDeleteCookiesCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefDeleteCookiesCallback := new(TCefDeleteCookiesCallback)
	SetObjectInstance(cefDeleteCookiesCallback, instance)
	return cefDeleteCookiesCallback
}

// AsCefDialogHandler Convert a pointer object to an existing class object
func AsCefDialogHandler(obj interface{}) ICefDialogHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefDialogHandler := new(TCefDialogHandler)
	SetObjectInstance(cefDialogHandler, instance)
	return cefDialogHandler
}

// AsCefDictionaryValue Convert a pointer object to an existing class object
func AsCefDictionaryValue(obj interface{}) ICefDictionaryValue {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefDictionaryValue := new(TCefDictionaryValue)
	SetObjectInstance(cefDictionaryValue, instance)
	return cefDictionaryValue
}

// AsCefDisplayHandler Convert a pointer object to an existing class object
func AsCefDisplayHandler(obj interface{}) ICefDisplayHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefDisplayHandler := new(TCefDisplayHandler)
	SetObjectInstance(cefDisplayHandler, instance)
	return cefDisplayHandler
}

// AsCefDisplay Convert a pointer object to an existing class object
func AsCefDisplay(obj interface{}) ICefDisplay {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefDisplay := new(TCefDisplay)
	SetObjectInstance(cefDisplay, instance)
	return cefDisplay
}

// AsCefDomDocument Convert a pointer object to an existing class object
func AsCefDomDocument(obj interface{}) ICefDomDocument {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefDomDocument := new(TCefDomDocument)
	SetObjectInstance(cefDomDocument, instance)
	return cefDomDocument
}

// AsCefDomNode Convert a pointer object to an existing class object
func AsCefDomNode(obj interface{}) ICefDomNode {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefDomNode := new(TCefDomNode)
	SetObjectInstance(cefDomNode, instance)
	return cefDomNode
}

// AsCefDomVisitor Convert a pointer object to an existing class object
func AsCefDomVisitor(obj interface{}) ICefDomVisitor {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefDomVisitor := new(TCefDomVisitor)
	SetObjectInstance(cefDomVisitor, instance)
	return cefDomVisitor
}

// AsCefDownloadHandler Convert a pointer object to an existing class object
func AsCefDownloadHandler(obj interface{}) ICefDownloadHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefDownloadHandler := new(TCefDownloadHandler)
	SetObjectInstance(cefDownloadHandler, instance)
	return cefDownloadHandler
}

// AsCefDownloadImageCallback Convert a pointer object to an existing class object
func AsCefDownloadImageCallback(obj interface{}) ICefDownloadImageCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefDownloadImageCallback := new(TCefDownloadImageCallback)
	SetObjectInstance(cefDownloadImageCallback, instance)
	return cefDownloadImageCallback
}

// AsCefDownloadItemCallback Convert a pointer object to an existing class object
func AsCefDownloadItemCallback(obj interface{}) ICefDownloadItemCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefDownloadItemCallback := new(TCefDownloadItemCallback)
	SetObjectInstance(cefDownloadItemCallback, instance)
	return cefDownloadItemCallback
}

// AsCefDownloadItem Convert a pointer object to an existing class object
func AsCefDownloadItem(obj interface{}) ICefDownloadItem {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefDownloadItem := new(TCefDownloadItem)
	SetObjectInstance(cefDownloadItem, instance)
	return cefDownloadItem
}

// AsCefDragData Convert a pointer object to an existing class object
func AsCefDragData(obj interface{}) ICefDragData {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefDragData := new(TCefDragData)
	SetObjectInstance(cefDragData, instance)
	return cefDragData
}

// AsCefDragHandler Convert a pointer object to an existing class object
func AsCefDragHandler(obj interface{}) ICefDragHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefDragHandler := new(TCefDragHandler)
	SetObjectInstance(cefDragHandler, instance)
	return cefDragHandler
}

// AsCefEnableFocusTask Convert a pointer object to an existing class object
func AsCefEnableFocusTask(obj interface{}) ICefEnableFocusTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefEnableFocusTask := new(TCefEnableFocusTask)
	SetObjectInstance(cefEnableFocusTask, instance)
	return cefEnableFocusTask
}

// AsCefExtensionHandler Convert a pointer object to an existing class object
func AsCefExtensionHandler(obj interface{}) ICefExtensionHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefExtensionHandler := new(TCefExtensionHandler)
	SetObjectInstance(cefExtensionHandler, instance)
	return cefExtensionHandler
}

// AsCefExtension Convert a pointer object to an existing class object
func AsCefExtension(obj interface{}) ICefExtension {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefExtension := new(TCefExtension)
	SetObjectInstance(cefExtension, instance)
	return cefExtension
}

// AsCefFileDialogCallback Convert a pointer object to an existing class object
func AsCefFileDialogCallback(obj interface{}) ICefFileDialogCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefFileDialogCallback := new(TCefFileDialogCallback)
	SetObjectInstance(cefFileDialogCallback, instance)
	return cefFileDialogCallback
}

// AsCefFillLayout Convert a pointer object to an existing class object
func AsCefFillLayout(obj interface{}) ICefFillLayout {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefFillLayout := new(TCefFillLayout)
	SetObjectInstance(cefFillLayout, instance)
	return cefFillLayout
}

// AsCefFindHandler Convert a pointer object to an existing class object
func AsCefFindHandler(obj interface{}) ICefFindHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefFindHandler := new(TCefFindHandler)
	SetObjectInstance(cefFindHandler, instance)
	return cefFindHandler
}

// AsCefFocusHandler Convert a pointer object to an existing class object
func AsCefFocusHandler(obj interface{}) ICefFocusHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefFocusHandler := new(TCefFocusHandler)
	SetObjectInstance(cefFocusHandler, instance)
	return cefFocusHandler
}

// AsCefFrameHandler Convert a pointer object to an existing class object
func AsCefFrameHandler(obj interface{}) ICefFrameHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefFrameHandler := new(TCefFrameHandler)
	SetObjectInstance(cefFrameHandler, instance)
	return cefFrameHandler
}

// AsCefFrame Convert a pointer object to an existing class object
func AsCefFrame(obj interface{}) ICefFrame {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefFrame := new(TCefFrame)
	SetObjectInstance(cefFrame, instance)
	return cefFrame
}

// AsCefGenericTask Convert a pointer object to an existing class object
func AsCefGenericTask(obj interface{}) ICefGenericTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefGenericTask := new(TCefGenericTask)
	SetObjectInstance(cefGenericTask, instance)
	return cefGenericTask
}

// AsCefGetExtensionResourceCallback Convert a pointer object to an existing class object
func AsCefGetExtensionResourceCallback(obj interface{}) ICefGetExtensionResourceCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefGetExtensionResourceCallback := new(TCefGetExtensionResourceCallback)
	SetObjectInstance(cefGetExtensionResourceCallback, instance)
	return cefGetExtensionResourceCallback
}

// AsCefImage Convert a pointer object to an existing class object
func AsCefImage(obj interface{}) ICefImage {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefImage := new(TCefImage)
	SetObjectInstance(cefImage, instance)
	return cefImage
}

// AsCefJsDialogCallback Convert a pointer object to an existing class object
func AsCefJsDialogCallback(obj interface{}) ICefJsDialogCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefJsDialogCallback := new(TCefJsDialogCallback)
	SetObjectInstance(cefJsDialogCallback, instance)
	return cefJsDialogCallback
}

// AsCefJsDialogHandler Convert a pointer object to an existing class object
func AsCefJsDialogHandler(obj interface{}) ICefJsDialogHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefJsDialogHandler := new(TCefJsDialogHandler)
	SetObjectInstance(cefJsDialogHandler, instance)
	return cefJsDialogHandler
}

// AsCefKeyboardHandler Convert a pointer object to an existing class object
func AsCefKeyboardHandler(obj interface{}) ICefKeyboardHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefKeyboardHandler := new(TCefKeyboardHandler)
	SetObjectInstance(cefKeyboardHandler, instance)
	return cefKeyboardHandler
}

// AsCefLabelButton Convert a pointer object to an existing class object
func AsCefLabelButton(obj interface{}) ICefLabelButton {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefLabelButton := new(TCefLabelButton)
	SetObjectInstance(cefLabelButton, instance)
	return cefLabelButton
}

// AsCefLayout Convert a pointer object to an existing class object
func AsCefLayout(obj interface{}) ICefLayout {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefLayout := new(TCefLayout)
	SetObjectInstance(cefLayout, instance)
	return cefLayout
}

// AsCefLifeSpanHandler Convert a pointer object to an existing class object
func AsCefLifeSpanHandler(obj interface{}) ICefLifeSpanHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefLifeSpanHandler := new(TCefLifeSpanHandler)
	SetObjectInstance(cefLifeSpanHandler, instance)
	return cefLifeSpanHandler
}

// AsCefListValue Convert a pointer object to an existing class object
func AsCefListValue(obj interface{}) ICefListValue {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefListValue := new(TCefListValue)
	SetObjectInstance(cefListValue, instance)
	return cefListValue
}

// AsCefLoadHandler Convert a pointer object to an existing class object
func AsCefLoadHandler(obj interface{}) ICefLoadHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefLoadHandler := new(TCefLoadHandler)
	SetObjectInstance(cefLoadHandler, instance)
	return cefLoadHandler
}

// AsCefMediaAccessCallback Convert a pointer object to an existing class object
func AsCefMediaAccessCallback(obj interface{}) ICefMediaAccessCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefMediaAccessCallback := new(TCefMediaAccessCallback)
	SetObjectInstance(cefMediaAccessCallback, instance)
	return cefMediaAccessCallback
}

// AsCefMediaObserver Convert a pointer object to an existing class object
func AsCefMediaObserver(obj interface{}) ICefMediaObserver {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefMediaObserver := new(TCefMediaObserver)
	SetObjectInstance(cefMediaObserver, instance)
	return cefMediaObserver
}

// AsCefMediaRouteCreateCallback Convert a pointer object to an existing class object
func AsCefMediaRouteCreateCallback(obj interface{}) ICefMediaRouteCreateCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefMediaRouteCreateCallback := new(TCefMediaRouteCreateCallback)
	SetObjectInstance(cefMediaRouteCreateCallback, instance)
	return cefMediaRouteCreateCallback
}

// AsCefMediaRoute Convert a pointer object to an existing class object
func AsCefMediaRoute(obj interface{}) ICefMediaRoute {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefMediaRoute := new(TCefMediaRoute)
	SetObjectInstance(cefMediaRoute, instance)
	return cefMediaRoute
}

// AsCefMediaRouter Convert a pointer object to an existing class object
func AsCefMediaRouter(obj interface{}) ICefMediaRouter {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefMediaRouter := new(TCefMediaRouter)
	SetObjectInstance(cefMediaRouter, instance)
	return cefMediaRouter
}

// AsCefMediaSinkDeviceInfoCallback Convert a pointer object to an existing class object
func AsCefMediaSinkDeviceInfoCallback(obj interface{}) ICefMediaSinkDeviceInfoCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefMediaSinkDeviceInfoCallback := new(TCefMediaSinkDeviceInfoCallback)
	SetObjectInstance(cefMediaSinkDeviceInfoCallback, instance)
	return cefMediaSinkDeviceInfoCallback
}

// AsCefMediaSink Convert a pointer object to an existing class object
func AsCefMediaSink(obj interface{}) ICefMediaSink {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefMediaSink := new(TCefMediaSink)
	SetObjectInstance(cefMediaSink, instance)
	return cefMediaSink
}

// AsCefMediaSource Convert a pointer object to an existing class object
func AsCefMediaSource(obj interface{}) ICefMediaSource {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefMediaSource := new(TCefMediaSource)
	SetObjectInstance(cefMediaSource, instance)
	return cefMediaSource
}

// AsCefMenuButtonDelegate Convert a pointer object to an existing class object
func AsCefMenuButtonDelegate(obj interface{}) ICefMenuButtonDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefMenuButtonDelegate := new(TCefMenuButtonDelegate)
	SetObjectInstance(cefMenuButtonDelegate, instance)
	return cefMenuButtonDelegate
}

// AsCefMenuButtonPressedLock Convert a pointer object to an existing class object
func AsCefMenuButtonPressedLock(obj interface{}) ICefMenuButtonPressedLock {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefMenuButtonPressedLock := new(TCefMenuButtonPressedLock)
	SetObjectInstance(cefMenuButtonPressedLock, instance)
	return cefMenuButtonPressedLock
}

// AsCefMenuButton Convert a pointer object to an existing class object
func AsCefMenuButton(obj interface{}) ICefMenuButton {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefMenuButton := new(TCefMenuButton)
	SetObjectInstance(cefMenuButton, instance)
	return cefMenuButton
}

// AsCefMenuModelDelegate Convert a pointer object to an existing class object
func AsCefMenuModelDelegate(obj interface{}) ICefMenuModelDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefMenuModelDelegate := new(TCefMenuModelDelegate)
	SetObjectInstance(cefMenuModelDelegate, instance)
	return cefMenuModelDelegate
}

// AsCefMenuModel Convert a pointer object to an existing class object
func AsCefMenuModel(obj interface{}) ICefMenuModel {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefMenuModel := new(TCefMenuModel)
	SetObjectInstance(cefMenuModel, instance)
	return cefMenuModel
}

// AsCefNavigationEntry Convert a pointer object to an existing class object
func AsCefNavigationEntry(obj interface{}) ICefNavigationEntry {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefNavigationEntry := new(TCefNavigationEntry)
	SetObjectInstance(cefNavigationEntry, instance)
	return cefNavigationEntry
}

// AsCefNavigationEntryVisitor Convert a pointer object to an existing class object
func AsCefNavigationEntryVisitor(obj interface{}) ICefNavigationEntryVisitor {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefNavigationEntryVisitor := new(TCefNavigationEntryVisitor)
	SetObjectInstance(cefNavigationEntryVisitor, instance)
	return cefNavigationEntryVisitor
}

// AsCefOverlayController Convert a pointer object to an existing class object
func AsCefOverlayController(obj interface{}) ICefOverlayController {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefOverlayController := new(TCefOverlayController)
	SetObjectInstance(cefOverlayController, instance)
	return cefOverlayController
}

// AsCefPanelDelegate Convert a pointer object to an existing class object
func AsCefPanelDelegate(obj interface{}) ICefPanelDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefPanelDelegate := new(TCefPanelDelegate)
	SetObjectInstance(cefPanelDelegate, instance)
	return cefPanelDelegate
}

// AsCefPanel Convert a pointer object to an existing class object
func AsCefPanel(obj interface{}) ICefPanel {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefPanel := new(TCefPanel)
	SetObjectInstance(cefPanel, instance)
	return cefPanel
}

// AsCefPdfPrintCallback Convert a pointer object to an existing class object
func AsCefPdfPrintCallback(obj interface{}) ICefPdfPrintCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefPdfPrintCallback := new(TCefPdfPrintCallback)
	SetObjectInstance(cefPdfPrintCallback, instance)
	return cefPdfPrintCallback
}

// AsCefPermissionHandler Convert a pointer object to an existing class object
func AsCefPermissionHandler(obj interface{}) ICefPermissionHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefPermissionHandler := new(TCefPermissionHandler)
	SetObjectInstance(cefPermissionHandler, instance)
	return cefPermissionHandler
}

// AsCefPermissionPromptCallback Convert a pointer object to an existing class object
func AsCefPermissionPromptCallback(obj interface{}) ICefPermissionPromptCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefPermissionPromptCallback := new(TCefPermissionPromptCallback)
	SetObjectInstance(cefPermissionPromptCallback, instance)
	return cefPermissionPromptCallback
}

// AsCefPostDataElement Convert a pointer object to an existing class object
func AsCefPostDataElement(obj interface{}) ICefPostDataElement {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefPostDataElement := new(TCefPostDataElement)
	SetObjectInstance(cefPostDataElement, instance)
	return cefPostDataElement
}

// AsCefPostData Convert a pointer object to an existing class object
func AsCefPostData(obj interface{}) ICefPostData {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefPostData := new(TCefPostData)
	SetObjectInstance(cefPostData, instance)
	return cefPostData
}

// AsCefPreferenceManager Convert a pointer object to an existing class object
func AsCefPreferenceManager(obj interface{}) ICefPreferenceManager {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefPreferenceManager := new(TCefPreferenceManager)
	SetObjectInstance(cefPreferenceManager, instance)
	return cefPreferenceManager
}

// AsCefPreferenceRegistrarRef Convert a pointer object to an existing class object
func AsCefPreferenceRegistrarRef(obj interface{}) ICefPreferenceRegistrarRef {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefPreferenceRegistrarRef := new(TCefPreferenceRegistrarRef)
	SetObjectInstance(cefPreferenceRegistrarRef, instance)
	return cefPreferenceRegistrarRef
}

// AsCefPrintDialogCallback Convert a pointer object to an existing class object
func AsCefPrintDialogCallback(obj interface{}) ICefPrintDialogCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefPrintDialogCallback := new(TCefPrintDialogCallback)
	SetObjectInstance(cefPrintDialogCallback, instance)
	return cefPrintDialogCallback
}

// AsCefPrintHandler Convert a pointer object to an existing class object
func AsCefPrintHandler(obj interface{}) ICefPrintHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefPrintHandler := new(TCefPrintHandler)
	SetObjectInstance(cefPrintHandler, instance)
	return cefPrintHandler
}

// AsCefPrintJobCallback Convert a pointer object to an existing class object
func AsCefPrintJobCallback(obj interface{}) ICefPrintJobCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefPrintJobCallback := new(TCefPrintJobCallback)
	SetObjectInstance(cefPrintJobCallback, instance)
	return cefPrintJobCallback
}

// AsCefPrintSettings Convert a pointer object to an existing class object
func AsCefPrintSettings(obj interface{}) ICefPrintSettings {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefPrintSettings := new(TCefPrintSettings)
	SetObjectInstance(cefPrintSettings, instance)
	return cefPrintSettings
}

// AsCefProcessMessage Convert a pointer object to an existing class object
func AsCefProcessMessage(obj interface{}) ICefProcessMessage {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefProcessMessage := new(TCefProcessMessage)
	SetObjectInstance(cefProcessMessage, instance)
	return cefProcessMessage
}

// AsCefReadZoomTask Convert a pointer object to an existing class object
func AsCefReadZoomTask(obj interface{}) ICefReadZoomTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefReadZoomTask := new(TCefReadZoomTask)
	SetObjectInstance(cefReadZoomTask, instance)
	return cefReadZoomTask
}

// AsCefRegistration Convert a pointer object to an existing class object
func AsCefRegistration(obj interface{}) ICefRegistration {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefRegistration := new(TCefRegistration)
	SetObjectInstance(cefRegistration, instance)
	return cefRegistration
}

// AsCefRenderHandler Convert a pointer object to an existing class object
func AsCefRenderHandler(obj interface{}) ICefRenderHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefRenderHandler := new(TCefRenderHandler)
	SetObjectInstance(cefRenderHandler, instance)
	return cefRenderHandler
}

// AsCefRenderProcessHandler Convert a pointer object to an existing class object
func AsCefRenderProcessHandler(obj interface{}) ICefRenderProcessHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefRenderProcessHandler := new(TCefRenderProcessHandler)
	SetObjectInstance(cefRenderProcessHandler, instance)
	return cefRenderProcessHandler
}

// AsCefRequestContextHandler Convert a pointer object to an existing class object
func AsCefRequestContextHandler(obj interface{}) ICefRequestContextHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefRequestContextHandler := new(TCefRequestContextHandler)
	SetObjectInstance(cefRequestContextHandler, instance)
	return cefRequestContextHandler
}

// AsCefRequestContext Convert a pointer object to an existing class object
func AsCefRequestContext(obj interface{}) ICefRequestContext {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefRequestContext := new(TCefRequestContext)
	SetObjectInstance(cefRequestContext, instance)
	return cefRequestContext
}

// AsCefRequestHandler Convert a pointer object to an existing class object
func AsCefRequestHandler(obj interface{}) ICefRequestHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefRequestHandler := new(TCefRequestHandler)
	SetObjectInstance(cefRequestHandler, instance)
	return cefRequestHandler
}

// AsCefRequest Convert a pointer object to an existing class object
func AsCefRequest(obj interface{}) ICefRequest {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefRequest := new(TCefRequest)
	SetObjectInstance(cefRequest, instance)
	return cefRequest
}

// AsCefResolveCallback Convert a pointer object to an existing class object
func AsCefResolveCallback(obj interface{}) ICefResolveCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefResolveCallback := new(TCefResolveCallback)
	SetObjectInstance(cefResolveCallback, instance)
	return cefResolveCallback
}

// AsCefResourceHandler Convert a pointer object to an existing class object
func AsCefResourceHandler(obj interface{}) ICefResourceHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefResourceHandler := new(TCefResourceHandler)
	SetObjectInstance(cefResourceHandler, instance)
	return cefResourceHandler
}

// AsCefResourceReadCallback Convert a pointer object to an existing class object
func AsCefResourceReadCallback(obj interface{}) ICefResourceReadCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefResourceReadCallback := new(TCefResourceReadCallback)
	SetObjectInstance(cefResourceReadCallback, instance)
	return cefResourceReadCallback
}

// AsCefResourceRequestHandler Convert a pointer object to an existing class object
func AsCefResourceRequestHandler(obj interface{}) ICefResourceRequestHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefResourceRequestHandler := new(TCefResourceRequestHandler)
	SetObjectInstance(cefResourceRequestHandler, instance)
	return cefResourceRequestHandler
}

// AsCefResourceSkipCallback Convert a pointer object to an existing class object
func AsCefResourceSkipCallback(obj interface{}) ICefResourceSkipCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefResourceSkipCallback := new(TCefResourceSkipCallback)
	SetObjectInstance(cefResourceSkipCallback, instance)
	return cefResourceSkipCallback
}

// AsCefResponseFilter Convert a pointer object to an existing class object
func AsCefResponseFilter(obj interface{}) ICefResponseFilter {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefResponseFilter := new(TCefResponseFilter)
	SetObjectInstance(cefResponseFilter, instance)
	return cefResponseFilter
}

// AsCefResponse Convert a pointer object to an existing class object
func AsCefResponse(obj interface{}) ICefResponse {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefResponse := new(TCefResponse)
	SetObjectInstance(cefResponse, instance)
	return cefResponse
}

// AsCefRunContextMenuCallback Convert a pointer object to an existing class object
func AsCefRunContextMenuCallback(obj interface{}) ICefRunContextMenuCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefRunContextMenuCallback := new(TCefRunContextMenuCallback)
	SetObjectInstance(cefRunContextMenuCallback, instance)
	return cefRunContextMenuCallback
}

// AsCefRunFileDialogCallback Convert a pointer object to an existing class object
func AsCefRunFileDialogCallback(obj interface{}) ICefRunFileDialogCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefRunFileDialogCallback := new(TCefRunFileDialogCallback)
	SetObjectInstance(cefRunFileDialogCallback, instance)
	return cefRunFileDialogCallback
}

// AsCefRunQuickMenuCallback Convert a pointer object to an existing class object
func AsCefRunQuickMenuCallback(obj interface{}) ICefRunQuickMenuCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefRunQuickMenuCallback := new(TCefRunQuickMenuCallback)
	SetObjectInstance(cefRunQuickMenuCallback, instance)
	return cefRunQuickMenuCallback
}

// AsCefSSLStatus Convert a pointer object to an existing class object
func AsCefSSLStatus(obj interface{}) ICefSSLStatus {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefSSLStatus := new(TCefSSLStatus)
	SetObjectInstance(cefSSLStatus, instance)
	return cefSSLStatus
}

// AsCefSavePrefsTask Convert a pointer object to an existing class object
func AsCefSavePrefsTask(obj interface{}) ICefSavePrefsTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefSavePrefsTask := new(TCefSavePrefsTask)
	SetObjectInstance(cefSavePrefsTask, instance)
	return cefSavePrefsTask
}

// AsCefSchemeHandlerFactory Convert a pointer object to an existing class object
func AsCefSchemeHandlerFactory(obj interface{}) ICefSchemeHandlerFactory {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefSchemeHandlerFactory := new(TCefSchemeHandlerFactory)
	SetObjectInstance(cefSchemeHandlerFactory, instance)
	return cefSchemeHandlerFactory
}

// AsCefSchemeRegistrarRef Convert a pointer object to an existing class object
func AsCefSchemeRegistrarRef(obj interface{}) ICefSchemeRegistrarRef {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefSchemeRegistrarRef := new(TCefSchemeRegistrarRef)
	SetObjectInstance(cefSchemeRegistrarRef, instance)
	return cefSchemeRegistrarRef
}

// AsCefScrollView Convert a pointer object to an existing class object
func AsCefScrollView(obj interface{}) ICefScrollView {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefScrollView := new(TCefScrollView)
	SetObjectInstance(cefScrollView, instance)
	return cefScrollView
}

// AsCefSelectClientCertificateCallback Convert a pointer object to an existing class object
func AsCefSelectClientCertificateCallback(obj interface{}) ICefSelectClientCertificateCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefSelectClientCertificateCallback := new(TCefSelectClientCertificateCallback)
	SetObjectInstance(cefSelectClientCertificateCallback, instance)
	return cefSelectClientCertificateCallback
}

// AsCefSetAudioMutedTask Convert a pointer object to an existing class object
func AsCefSetAudioMutedTask(obj interface{}) ICefSetAudioMutedTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefSetAudioMutedTask := new(TCefSetAudioMutedTask)
	SetObjectInstance(cefSetAudioMutedTask, instance)
	return cefSetAudioMutedTask
}

// AsCefSetCookieCallback Convert a pointer object to an existing class object
func AsCefSetCookieCallback(obj interface{}) ICefSetCookieCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefSetCookieCallback := new(TCefSetCookieCallback)
	SetObjectInstance(cefSetCookieCallback, instance)
	return cefSetCookieCallback
}

// AsCefSetZoomLevelTask Convert a pointer object to an existing class object
func AsCefSetZoomLevelTask(obj interface{}) ICefSetZoomLevelTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefSetZoomLevelTask := new(TCefSetZoomLevelTask)
	SetObjectInstance(cefSetZoomLevelTask, instance)
	return cefSetZoomLevelTask
}

// AsCefSetZoomPctTask Convert a pointer object to an existing class object
func AsCefSetZoomPctTask(obj interface{}) ICefSetZoomPctTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefSetZoomPctTask := new(TCefSetZoomPctTask)
	SetObjectInstance(cefSetZoomPctTask, instance)
	return cefSetZoomPctTask
}

// AsCefSetZoomStepTask Convert a pointer object to an existing class object
func AsCefSetZoomStepTask(obj interface{}) ICefSetZoomStepTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefSetZoomStepTask := new(TCefSetZoomStepTask)
	SetObjectInstance(cefSetZoomStepTask, instance)
	return cefSetZoomStepTask
}

// AsCefSharedMemoryRegion Convert a pointer object to an existing class object
func AsCefSharedMemoryRegion(obj interface{}) ICefSharedMemoryRegion {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefSharedMemoryRegion := new(TCefSharedMemoryRegion)
	SetObjectInstance(cefSharedMemoryRegion, instance)
	return cefSharedMemoryRegion
}

// AsCefSharedProcessMessageBuilder Convert a pointer object to an existing class object
func AsCefSharedProcessMessageBuilder(obj interface{}) ICefSharedProcessMessageBuilder {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefSharedProcessMessageBuilder := new(TCefSharedProcessMessageBuilder)
	SetObjectInstance(cefSharedProcessMessageBuilder, instance)
	return cefSharedProcessMessageBuilder
}

// AsCefSslInfo Convert a pointer object to an existing class object
func AsCefSslInfo(obj interface{}) ICefSslInfo {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefSslInfo := new(TCefSslInfo)
	SetObjectInstance(cefSslInfo, instance)
	return cefSslInfo
}

// AsCefStreamReader Convert a pointer object to an existing class object
func AsCefStreamReader(obj interface{}) ICefStreamReader {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefStreamReader := new(TCefStreamReader)
	SetObjectInstance(cefStreamReader, instance)
	return cefStreamReader
}

// AsCefStreamWriter Convert a pointer object to an existing class object
func AsCefStreamWriter(obj interface{}) ICefStreamWriter {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefStreamWriter := new(TCefStreamWriter)
	SetObjectInstance(cefStreamWriter, instance)
	return cefStreamWriter
}

// AsCefStringListRef Convert a pointer object to an existing class object
func AsCefStringListRef(obj interface{}) ICefStringListRef {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefStringListRef := new(TCefStringListRef)
	SetObjectInstance(cefStringListRef, instance)
	return cefStringListRef
}

// AsCefStringMapRef Convert a pointer object to an existing class object
func AsCefStringMapRef(obj interface{}) ICefStringMapRef {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefStringMapRef := new(TCefStringMapRef)
	SetObjectInstance(cefStringMapRef, instance)
	return cefStringMapRef
}

// AsCefStringMultimapRef Convert a pointer object to an existing class object
func AsCefStringMultimapRef(obj interface{}) ICefStringMultimapRef {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefStringMultimapRef := new(TCefStringMultimapRef)
	SetObjectInstance(cefStringMultimapRef, instance)
	return cefStringMultimapRef
}

// AsCefStringVisitor Convert a pointer object to an existing class object
func AsCefStringVisitor(obj interface{}) ICefStringVisitor {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefStringVisitor := new(TCefStringVisitor)
	SetObjectInstance(cefStringVisitor, instance)
	return cefStringVisitor
}

// AsCefTask Convert a pointer object to an existing class object
func AsCefTask(obj interface{}) ICefTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefTask := new(TCefTask)
	SetObjectInstance(cefTask, instance)
	return cefTask
}

// AsCefTaskRunner Convert a pointer object to an existing class object
func AsCefTaskRunner(obj interface{}) ICefTaskRunner {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefTaskRunner := new(TCefTaskRunner)
	SetObjectInstance(cefTaskRunner, instance)
	return cefTaskRunner
}

// AsCefTextfieldDelegate Convert a pointer object to an existing class object
func AsCefTextfieldDelegate(obj interface{}) ICefTextfieldDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefTextfieldDelegate := new(TCefTextfieldDelegate)
	SetObjectInstance(cefTextfieldDelegate, instance)
	return cefTextfieldDelegate
}

// AsCefTextfield Convert a pointer object to an existing class object
func AsCefTextfield(obj interface{}) ICefTextfield {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefTextfield := new(TCefTextfield)
	SetObjectInstance(cefTextfield, instance)
	return cefTextfield
}

// AsCefToggleAudioMutedTask Convert a pointer object to an existing class object
func AsCefToggleAudioMutedTask(obj interface{}) ICefToggleAudioMutedTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefToggleAudioMutedTask := new(TCefToggleAudioMutedTask)
	SetObjectInstance(cefToggleAudioMutedTask, instance)
	return cefToggleAudioMutedTask
}

// AsCefURLRequestTask Convert a pointer object to an existing class object
func AsCefURLRequestTask(obj interface{}) ICefURLRequestTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefURLRequestTask := new(TCefURLRequestTask)
	SetObjectInstance(cefURLRequestTask, instance)
	return cefURLRequestTask
}

// AsCefUpdatePrefsTask Convert a pointer object to an existing class object
func AsCefUpdatePrefsTask(obj interface{}) ICefUpdatePrefsTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefUpdatePrefsTask := new(TCefUpdatePrefsTask)
	SetObjectInstance(cefUpdatePrefsTask, instance)
	return cefUpdatePrefsTask
}

// AsCefUpdateZoomPctTask Convert a pointer object to an existing class object
func AsCefUpdateZoomPctTask(obj interface{}) ICefUpdateZoomPctTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefUpdateZoomPctTask := new(TCefUpdateZoomPctTask)
	SetObjectInstance(cefUpdateZoomPctTask, instance)
	return cefUpdateZoomPctTask
}

// AsCefUpdateZoomStepTask Convert a pointer object to an existing class object
func AsCefUpdateZoomStepTask(obj interface{}) ICefUpdateZoomStepTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefUpdateZoomStepTask := new(TCefUpdateZoomStepTask)
	SetObjectInstance(cefUpdateZoomStepTask, instance)
	return cefUpdateZoomStepTask
}

// AsCefUrlRequest Convert a pointer object to an existing class object
func AsCefUrlRequest(obj interface{}) ICefUrlRequest {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefUrlRequest := new(TCefUrlRequest)
	SetObjectInstance(cefUrlRequest, instance)
	return cefUrlRequest
}

// AsCefUrlRequestClient Convert a pointer object to an existing class object
func AsCefUrlRequestClient(obj interface{}) ICefUrlRequestClient {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefUrlRequestClient := new(TCefUrlRequestClient)
	SetObjectInstance(cefUrlRequestClient, instance)
	return cefUrlRequestClient
}

// AsCefV8Accessor Convert a pointer object to an existing class object
func AsCefV8Accessor(obj interface{}) ICefV8Accessor {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefV8Accessor := new(TCefV8Accessor)
	SetObjectInstance(cefV8Accessor, instance)
	return cefV8Accessor
}

// AsCefV8Exception Convert a pointer object to an existing class object
func AsCefV8Exception(obj interface{}) ICefV8Exception {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefV8Exception := new(TCefV8Exception)
	SetObjectInstance(cefV8Exception, instance)
	return cefV8Exception
}

// AsCefV8Interceptor Convert a pointer object to an existing class object
func AsCefV8Interceptor(obj interface{}) ICefV8Interceptor {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefV8Interceptor := new(TCefV8Interceptor)
	SetObjectInstance(cefV8Interceptor, instance)
	return cefV8Interceptor
}

// AsCefV8StackFrame Convert a pointer object to an existing class object
func AsCefV8StackFrame(obj interface{}) ICefV8StackFrame {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefV8StackFrame := new(TCefV8StackFrame)
	SetObjectInstance(cefV8StackFrame, instance)
	return cefV8StackFrame
}

// AsCefV8StackTrace Convert a pointer object to an existing class object
func AsCefV8StackTrace(obj interface{}) ICefV8StackTrace {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefV8StackTrace := new(TCefV8StackTrace)
	SetObjectInstance(cefV8StackTrace, instance)
	return cefV8StackTrace
}

// AsCefValue Convert a pointer object to an existing class object
func AsCefValue(obj interface{}) ICefValue {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefValue := new(TCefValue)
	SetObjectInstance(cefValue, instance)
	return cefValue
}

// AsCefViewDelegate Convert a pointer object to an existing class object
func AsCefViewDelegate(obj interface{}) ICefViewDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefViewDelegate := new(TCefViewDelegate)
	SetObjectInstance(cefViewDelegate, instance)
	return cefViewDelegate
}

// AsCefView Convert a pointer object to an existing class object
func AsCefView(obj interface{}) ICefView {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefView := new(TCefView)
	SetObjectInstance(cefView, instance)
	return cefView
}

// AsCefWindowDelegate Convert a pointer object to an existing class object
func AsCefWindowDelegate(obj interface{}) ICefWindowDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefWindowDelegate := new(TCefWindowDelegate)
	SetObjectInstance(cefWindowDelegate, instance)
	return cefWindowDelegate
}

// AsCefWindow Convert a pointer object to an existing class object
func AsCefWindow(obj interface{}) ICefWindow {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefWindow := new(TCefWindow)
	SetObjectInstance(cefWindow, instance)
	return cefWindow
}

// AsCefWriteHandler Convert a pointer object to an existing class object
func AsCefWriteHandler(obj interface{}) ICefWriteHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefWriteHandler := new(TCefWriteHandler)
	SetObjectInstance(cefWriteHandler, instance)
	return cefWriteHandler
}

// AsCefX509CertPrincipal Convert a pointer object to an existing class object
func AsCefX509CertPrincipal(obj interface{}) ICefX509CertPrincipal {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefX509CertPrincipal := new(TCefX509CertPrincipal)
	SetObjectInstance(cefX509CertPrincipal, instance)
	return cefX509CertPrincipal
}

// AsCefv8ArrayBufferReleaseCallback Convert a pointer object to an existing class object
func AsCefv8ArrayBufferReleaseCallback(obj interface{}) ICefv8ArrayBufferReleaseCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefv8ArrayBufferReleaseCallback := new(TCefv8ArrayBufferReleaseCallback)
	SetObjectInstance(cefv8ArrayBufferReleaseCallback, instance)
	return cefv8ArrayBufferReleaseCallback
}

// AsCefv8Context Convert a pointer object to an existing class object
func AsCefv8Context(obj interface{}) ICefv8Context {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefv8Context := new(TCefv8Context)
	SetObjectInstance(cefv8Context, instance)
	return cefv8Context
}

// AsCefv8Handler Convert a pointer object to an existing class object
func AsCefv8Handler(obj interface{}) ICefv8Handler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefv8Handler := new(TCefv8Handler)
	SetObjectInstance(cefv8Handler, instance)
	return cefv8Handler
}

// AsCefv8Value Convert a pointer object to an existing class object
func AsCefv8Value(obj interface{}) ICefv8Value {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	cefv8Value := new(TCefv8Value)
	SetObjectInstance(cefv8Value, instance)
	return cefv8Value
}

// AsChromium Convert a pointer object to an existing class object
func AsChromium(obj interface{}) IChromium {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	chromium := new(TChromium)
	SetObjectInstance(chromium, instance)
	return chromium
}

// AsChromiumCore Convert a pointer object to an existing class object
func AsChromiumCore(obj interface{}) IChromiumCore {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	chromiumCore := new(TChromiumCore)
	SetObjectInstance(chromiumCore, instance)
	return chromiumCore
}

// AsChromiumFontOptions Convert a pointer object to an existing class object
func AsChromiumFontOptions(obj interface{}) IChromiumFontOptions {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	chromiumFontOptions := new(TChromiumFontOptions)
	SetObjectInstance(chromiumFontOptions, instance)
	return chromiumFontOptions
}

// AsChromiumOptions Convert a pointer object to an existing class object
func AsChromiumOptions(obj interface{}) IChromiumOptions {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	chromiumOptions := new(TChromiumOptions)
	SetObjectInstance(chromiumOptions, instance)
	return chromiumOptions
}

// AsCustomAudioHandler Convert a pointer object to an existing class object
func AsCustomAudioHandler(obj interface{}) ICustomAudioHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customAudioHandler := new(TCustomAudioHandler)
	SetObjectInstance(customAudioHandler, instance)
	return customAudioHandler
}

// AsCustomBrowserViewDelegate Convert a pointer object to an existing class object
func AsCustomBrowserViewDelegate(obj interface{}) ICustomBrowserViewDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customBrowserViewDelegate := new(TCustomBrowserViewDelegate)
	SetObjectInstance(customBrowserViewDelegate, instance)
	return customBrowserViewDelegate
}

// AsCustomButtonDelegate Convert a pointer object to an existing class object
func AsCustomButtonDelegate(obj interface{}) ICustomButtonDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customButtonDelegate := new(TCustomButtonDelegate)
	SetObjectInstance(customButtonDelegate, instance)
	return customButtonDelegate
}

// AsCustomCefNavigationEntryVisitor Convert a pointer object to an existing class object
func AsCustomCefNavigationEntryVisitor(obj interface{}) ICustomCefNavigationEntryVisitor {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customCefNavigationEntryVisitor := new(TCustomCefNavigationEntryVisitor)
	SetObjectInstance(customCefNavigationEntryVisitor, instance)
	return customCefNavigationEntryVisitor
}

// AsCustomCefStringVisitor Convert a pointer object to an existing class object
func AsCustomCefStringVisitor(obj interface{}) ICustomCefStringVisitor {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customCefStringVisitor := new(TCustomCefStringVisitor)
	SetObjectInstance(customCefStringVisitor, instance)
	return customCefStringVisitor
}

// AsCustomCefUrlrequestClient Convert a pointer object to an existing class object
func AsCustomCefUrlrequestClient(obj interface{}) ICustomCefUrlrequestClient {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customCefUrlrequestClient := new(TCustomCefUrlrequestClient)
	SetObjectInstance(customCefUrlrequestClient, instance)
	return customCefUrlrequestClient
}

// AsCustomClientHandler Convert a pointer object to an existing class object
func AsCustomClientHandler(obj interface{}) ICustomClientHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customClientHandler := new(TCustomClientHandler)
	SetObjectInstance(customClientHandler, instance)
	return customClientHandler
}

// AsCustomCommandHandler Convert a pointer object to an existing class object
func AsCustomCommandHandler(obj interface{}) ICustomCommandHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customCommandHandler := new(TCustomCommandHandler)
	SetObjectInstance(customCommandHandler, instance)
	return customCommandHandler
}

// AsCustomContextMenuHandler Convert a pointer object to an existing class object
func AsCustomContextMenuHandler(obj interface{}) ICustomContextMenuHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customContextMenuHandler := new(TCustomContextMenuHandler)
	SetObjectInstance(customContextMenuHandler, instance)
	return customContextMenuHandler
}

// AsCustomDevToolsMessageObserver Convert a pointer object to an existing class object
func AsCustomDevToolsMessageObserver(obj interface{}) ICustomDevToolsMessageObserver {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customDevToolsMessageObserver := new(TCustomDevToolsMessageObserver)
	SetObjectInstance(customDevToolsMessageObserver, instance)
	return customDevToolsMessageObserver
}

// AsCustomDialogHandler Convert a pointer object to an existing class object
func AsCustomDialogHandler(obj interface{}) ICustomDialogHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customDialogHandler := new(TCustomDialogHandler)
	SetObjectInstance(customDialogHandler, instance)
	return customDialogHandler
}

// AsCustomDisplayHandler Convert a pointer object to an existing class object
func AsCustomDisplayHandler(obj interface{}) ICustomDisplayHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customDisplayHandler := new(TCustomDisplayHandler)
	SetObjectInstance(customDisplayHandler, instance)
	return customDisplayHandler
}

// AsCustomDownloadHandler Convert a pointer object to an existing class object
func AsCustomDownloadHandler(obj interface{}) ICustomDownloadHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customDownloadHandler := new(TCustomDownloadHandler)
	SetObjectInstance(customDownloadHandler, instance)
	return customDownloadHandler
}

// AsCustomDragHandler Convert a pointer object to an existing class object
func AsCustomDragHandler(obj interface{}) ICustomDragHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customDragHandler := new(TCustomDragHandler)
	SetObjectInstance(customDragHandler, instance)
	return customDragHandler
}

// AsCustomExtensionHandler Convert a pointer object to an existing class object
func AsCustomExtensionHandler(obj interface{}) ICustomExtensionHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customExtensionHandler := new(TCustomExtensionHandler)
	SetObjectInstance(customExtensionHandler, instance)
	return customExtensionHandler
}

// AsCustomFindHandler Convert a pointer object to an existing class object
func AsCustomFindHandler(obj interface{}) ICustomFindHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customFindHandler := new(TCustomFindHandler)
	SetObjectInstance(customFindHandler, instance)
	return customFindHandler
}

// AsCustomFocusHandler Convert a pointer object to an existing class object
func AsCustomFocusHandler(obj interface{}) ICustomFocusHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customFocusHandler := new(TCustomFocusHandler)
	SetObjectInstance(customFocusHandler, instance)
	return customFocusHandler
}

// AsCustomFrameHandler Convert a pointer object to an existing class object
func AsCustomFrameHandler(obj interface{}) ICustomFrameHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customFrameHandler := new(TCustomFrameHandler)
	SetObjectInstance(customFrameHandler, instance)
	return customFrameHandler
}

// AsCustomJsDialogHandler Convert a pointer object to an existing class object
func AsCustomJsDialogHandler(obj interface{}) ICustomJsDialogHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customJsDialogHandler := new(TCustomJsDialogHandler)
	SetObjectInstance(customJsDialogHandler, instance)
	return customJsDialogHandler
}

// AsCustomKeyboardHandler Convert a pointer object to an existing class object
func AsCustomKeyboardHandler(obj interface{}) ICustomKeyboardHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customKeyboardHandler := new(TCustomKeyboardHandler)
	SetObjectInstance(customKeyboardHandler, instance)
	return customKeyboardHandler
}

// AsCustomLifeSpanHandler Convert a pointer object to an existing class object
func AsCustomLifeSpanHandler(obj interface{}) ICustomLifeSpanHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customLifeSpanHandler := new(TCustomLifeSpanHandler)
	SetObjectInstance(customLifeSpanHandler, instance)
	return customLifeSpanHandler
}

// AsCustomLoadHandler Convert a pointer object to an existing class object
func AsCustomLoadHandler(obj interface{}) ICustomLoadHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customLoadHandler := new(TCustomLoadHandler)
	SetObjectInstance(customLoadHandler, instance)
	return customLoadHandler
}

// AsCustomMediaObserver Convert a pointer object to an existing class object
func AsCustomMediaObserver(obj interface{}) ICustomMediaObserver {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customMediaObserver := new(TCustomMediaObserver)
	SetObjectInstance(customMediaObserver, instance)
	return customMediaObserver
}

// AsCustomMenuButtonDelegate Convert a pointer object to an existing class object
func AsCustomMenuButtonDelegate(obj interface{}) ICustomMenuButtonDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customMenuButtonDelegate := new(TCustomMenuButtonDelegate)
	SetObjectInstance(customMenuButtonDelegate, instance)
	return customMenuButtonDelegate
}

// AsCustomPanelDelegate Convert a pointer object to an existing class object
func AsCustomPanelDelegate(obj interface{}) ICustomPanelDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customPanelDelegate := new(TCustomPanelDelegate)
	SetObjectInstance(customPanelDelegate, instance)
	return customPanelDelegate
}

// AsCustomPermissionHandler Convert a pointer object to an existing class object
func AsCustomPermissionHandler(obj interface{}) ICustomPermissionHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customPermissionHandler := new(TCustomPermissionHandler)
	SetObjectInstance(customPermissionHandler, instance)
	return customPermissionHandler
}

// AsCustomPrintHandler Convert a pointer object to an existing class object
func AsCustomPrintHandler(obj interface{}) ICustomPrintHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customPrintHandler := new(TCustomPrintHandler)
	SetObjectInstance(customPrintHandler, instance)
	return customPrintHandler
}

// AsCustomRenderHandler Convert a pointer object to an existing class object
func AsCustomRenderHandler(obj interface{}) ICustomRenderHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customRenderHandler := new(TCustomRenderHandler)
	SetObjectInstance(customRenderHandler, instance)
	return customRenderHandler
}

// AsCustomRenderLoadHandler Convert a pointer object to an existing class object
func AsCustomRenderLoadHandler(obj interface{}) ICustomRenderLoadHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customRenderLoadHandler := new(TCustomRenderLoadHandler)
	SetObjectInstance(customRenderLoadHandler, instance)
	return customRenderLoadHandler
}

// AsCustomRequestContextHandler Convert a pointer object to an existing class object
func AsCustomRequestContextHandler(obj interface{}) ICustomRequestContextHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customRequestContextHandler := new(TCustomRequestContextHandler)
	SetObjectInstance(customRequestContextHandler, instance)
	return customRequestContextHandler
}

// AsCustomRequestHandler Convert a pointer object to an existing class object
func AsCustomRequestHandler(obj interface{}) ICustomRequestHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customRequestHandler := new(TCustomRequestHandler)
	SetObjectInstance(customRequestHandler, instance)
	return customRequestHandler
}

// AsCustomResourceRequestHandler Convert a pointer object to an existing class object
func AsCustomResourceRequestHandler(obj interface{}) ICustomResourceRequestHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customResourceRequestHandler := new(TCustomResourceRequestHandler)
	SetObjectInstance(customResourceRequestHandler, instance)
	return customResourceRequestHandler
}

// AsCustomResponseFilter Convert a pointer object to an existing class object
func AsCustomResponseFilter(obj interface{}) ICustomResponseFilter {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customResponseFilter := new(TCustomResponseFilter)
	SetObjectInstance(customResponseFilter, instance)
	return customResponseFilter
}

// AsCustomServerHandler Convert a pointer object to an existing class object
func AsCustomServerHandler(obj interface{}) ICustomServerHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customServerHandler := new(TCustomServerHandler)
	SetObjectInstance(customServerHandler, instance)
	return customServerHandler
}

// AsCustomTextfieldDelegate Convert a pointer object to an existing class object
func AsCustomTextfieldDelegate(obj interface{}) ICustomTextfieldDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customTextfieldDelegate := new(TCustomTextfieldDelegate)
	SetObjectInstance(customTextfieldDelegate, instance)
	return customTextfieldDelegate
}

// AsCustomViewDelegate Convert a pointer object to an existing class object
func AsCustomViewDelegate(obj interface{}) ICustomViewDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customViewDelegate := new(TCustomViewDelegate)
	SetObjectInstance(customViewDelegate, instance)
	return customViewDelegate
}

// AsCustomWindowDelegate Convert a pointer object to an existing class object
func AsCustomWindowDelegate(obj interface{}) ICustomWindowDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customWindowDelegate := new(TCustomWindowDelegate)
	SetObjectInstance(customWindowDelegate, instance)
	return customWindowDelegate
}

// AsDomVisitor Convert a pointer object to an existing class object
func AsDomVisitor(obj interface{}) IDomVisitor {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	domVisitor := new(TDomVisitor)
	SetObjectInstance(domVisitor, instance)
	return domVisitor
}

// AsEnergyStringList Convert a pointer object to an existing class object
func AsEnergyStringList(obj interface{}) IEnergyStringList {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	energyStringList := new(TEnergyStringList)
	SetObjectInstance(energyStringList, instance)
	return energyStringList
}

// AsEnergyStringMap Convert a pointer object to an existing class object
func AsEnergyStringMap(obj interface{}) IEnergyStringMap {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	energyStringMap := new(TEnergyStringMap)
	SetObjectInstance(energyStringMap, instance)
	return energyStringMap
}

// AsEnergyStringMultiMap Convert a pointer object to an existing class object
func AsEnergyStringMultiMap(obj interface{}) IEnergyStringMultiMap {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	energyStringMultiMap := new(TEnergyStringMultiMap)
	SetObjectInstance(energyStringMultiMap, instance)
	return energyStringMultiMap
}

// AsMenuModelDelegate Convert a pointer object to an existing class object
func AsMenuModelDelegate(obj interface{}) IMenuModelDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	menuModelDelegate := new(TMenuModelDelegate)
	SetObjectInstance(menuModelDelegate, instance)
	return menuModelDelegate
}

// AsPDFPrintOptions Convert a pointer object to an existing class object
func AsPDFPrintOptions(obj interface{}) IPDFPrintOptions {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	pDFPrintOptions := new(TPDFPrintOptions)
	SetObjectInstance(pDFPrintOptions, instance)
	return pDFPrintOptions
}

// AsResourceHandler Convert a pointer object to an existing class object
func AsResourceHandler(obj interface{}) IResourceHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	resourceHandler := new(TResourceHandler)
	SetObjectInstance(resourceHandler, instance)
	return resourceHandler
}

// AsRunFileDialogCallback Convert a pointer object to an existing class object
func AsRunFileDialogCallback(obj interface{}) IRunFileDialogCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	runFileDialogCallback := new(TRunFileDialogCallback)
	SetObjectInstance(runFileDialogCallback, instance)
	return runFileDialogCallback
}

// AsSchemeHandlerFactory Convert a pointer object to an existing class object
func AsSchemeHandlerFactory(obj interface{}) ISchemeHandlerFactory {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	schemeHandlerFactory := new(TSchemeHandlerFactory)
	SetObjectInstance(schemeHandlerFactory, instance)
	return schemeHandlerFactory
}

// AsTask Convert a pointer object to an existing class object
func AsTask(obj interface{}) ITask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	task := new(TTask)
	SetObjectInstance(task, instance)
	return task
}

// AsV8Accessor Convert a pointer object to an existing class object
func AsV8Accessor(obj interface{}) IV8Accessor {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	v8Accessor := new(TV8Accessor)
	SetObjectInstance(v8Accessor, instance)
	return v8Accessor
}

// AsV8ArrayBufferReleaseCallback Convert a pointer object to an existing class object
func AsV8ArrayBufferReleaseCallback(obj interface{}) IV8ArrayBufferReleaseCallback {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	v8ArrayBufferReleaseCallback := new(TV8ArrayBufferReleaseCallback)
	SetObjectInstance(v8ArrayBufferReleaseCallback, instance)
	return v8ArrayBufferReleaseCallback
}

// AsV8Handler Convert a pointer object to an existing class object
func AsV8Handler(obj interface{}) IV8Handler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	v8Handler := new(TV8Handler)
	SetObjectInstance(v8Handler, instance)
	return v8Handler
}

// AsV8Interceptor Convert a pointer object to an existing class object
func AsV8Interceptor(obj interface{}) IV8Interceptor {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	v8Interceptor := new(TV8Interceptor)
	SetObjectInstance(v8Interceptor, instance)
	return v8Interceptor
}
