//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefRequestHandler
type ICefRequestHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// RequestHandlerRef -> ICefRequestHandler
var RequestHandlerRef requestHandler

type requestHandler uintptr

func (*requestHandler) New() *ICefRequestHandler {
	var result uintptr
	imports.Proc(def.CefRequestHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRequestHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefRequestHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefRequestHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefRequestHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefRequestHandler) SetOnBeforeBrowse(fn requestHandlerOnBeforeBrowse) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequestHandler_OnBeforeBrowse).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRequestHandler) SetOnOpenUrlFromTab(fn requestHandlerOnOpenUrlFromTab) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequestHandler_OnOpenUrlFromTab).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRequestHandler) SetGetResourceRequestHandler(fn requestHandlerGetResourceRequestHandler) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequestHandler_GetResourceRequestHandler).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRequestHandler) SetGetAuthCredentials(fn requestHandlerGetAuthCredentials) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequestHandler_GetAuthCredentials).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRequestHandler) SetOnCertificateError(fn requestHandlerOnCertificateError) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequestHandler_OnCertificateError).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRequestHandler) SetOnSelectClientCertificate(fn requestHandlerOnSelectClientCertificate) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequestHandler_OnSelectClientCertificate).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRequestHandler) SetOnRenderViewReady(fn requestHandlerOnRenderViewReady) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequestHandler_OnRenderViewReady).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRequestHandler) SetOnRenderProcessTerminated(fn requestHandlerOnRenderProcessTerminated) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequestHandler_OnRenderProcessTerminated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRequestHandler) SetOnDocumentAvailableInMainFrame(fn requestHandlerOnDocumentAvailableInMainFrame) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefRequestHandler_OnDocumentAvailableInMainFrame).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type requestHandlerOnBeforeBrowse func(browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, userGesture, isRedirect bool) bool
type requestHandlerOnOpenUrlFromTab func(browser *ICefBrowser, frame *ICefFrame, targetUrl string, targetDisposition consts.TCefWindowOpenDisposition, userGesture bool) bool
type requestHandlerGetResourceRequestHandler func(browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, isNavigation, isDownload bool, requestInitiator string) (disableDefaultHandling bool, resourceRequestHandler *ICefResourceRequestHandler)
type requestHandlerGetAuthCredentials func(browser *ICefBrowser, originUrl string, isProxy bool, host string, port int32, realm, scheme string, callback *ICefAuthCallback) bool
type requestHandlerOnCertificateError func(browser *ICefBrowser, certError consts.TCefErrorCode, requestUrl string, sslInfo *ICefSslInfo, callback *ICefCallback) bool
type requestHandlerOnSelectClientCertificate func(browser *ICefBrowser, isProxy bool, host string, port int32, certificates *TCefX509CertificateArray, callback *ICefSelectClientCertificateCallback) bool
type requestHandlerOnRenderViewReady func(browser *ICefBrowser)
type requestHandlerOnRenderProcessTerminated func(browser *ICefBrowser, status consts.TCefTerminationStatus, errorCode int32, error_ string)
type requestHandlerOnDocumentAvailableInMainFrame func(browser *ICefBrowser)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case requestHandlerOnBeforeBrowse:
			browser := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			request := &ICefRequest{instance: getPtr(2)}
			userGesture, isRedirect := api.GoBool(getVal(3)), api.GoBool(getVal(4))
			resultPtr := (*bool)(getPtr(5))
			*resultPtr = fn.(requestHandlerOnBeforeBrowse)(browser, frame, request, userGesture, isRedirect)
		case requestHandlerOnOpenUrlFromTab:
			browser := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			targetUrl := api.GoStr(getVal(2))
			targetDisposition := consts.TCefWindowOpenDisposition(getVal(3))
			userGesture := api.GoBool(getVal(4))
			resultPtr := (*bool)(getPtr(5))
			*resultPtr = fn.(requestHandlerOnOpenUrlFromTab)(browser, frame, targetUrl, targetDisposition, userGesture)
		case requestHandlerGetResourceRequestHandler:
			browser := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			request := &ICefRequest{instance: getPtr(2)}
			isNavigation, isDownload := api.GoBool(getVal(3)), api.GoBool(getVal(4))
			requestInitiator := api.GoStr(getVal(5))
			disableDefaultHandlingPtr := (*bool)(getPtr(6))
			resourceRequestHandlerPtr := (*uintptr)(getPtr(7))
			disableDefaultHandling, resourceRequestHandler := fn.(requestHandlerGetResourceRequestHandler)(browser, frame, request, isNavigation, isDownload, requestInitiator)
			*disableDefaultHandlingPtr = disableDefaultHandling
			if resourceRequestHandler != nil && resourceRequestHandler.IsValid() {
				*resourceRequestHandlerPtr = resourceRequestHandler.Instance()
			}
		case requestHandlerGetAuthCredentials:
			browser := &ICefBrowser{instance: getPtr(0)}
			originUrl := api.GoStr(getVal(1))
			isProxy := api.GoBool(getVal(2))
			host := api.GoStr(getVal(3))
			port := int32(getVal(4))
			realm, scheme := api.GoStr(getVal(5)), api.GoStr(getVal(6))
			callback := &ICefAuthCallback{instance: getPtr(7)}
			resultPtr := (*bool)(getPtr(8))
			*resultPtr = fn.(requestHandlerGetAuthCredentials)(browser, originUrl, isProxy, host, port, realm, scheme, callback)
		case requestHandlerOnCertificateError:
			browser := &ICefBrowser{instance: getPtr(0)}
			certError := consts.TCefErrorCode(getVal(1))
			requestUrl := api.GoStr(getVal(2))
			sslInfo := &ICefSslInfo{instance: getPtr(3)}
			callback := &ICefCallback{instance: getPtr(4)}
			resultPtr := (*bool)(getPtr(5))
			*resultPtr = fn.(requestHandlerOnCertificateError)(browser, certError, requestUrl, sslInfo, callback)
		case requestHandlerOnSelectClientCertificate:
			browser := &ICefBrowser{instance: getPtr(0)}
			isProxy := api.GoBool(getVal(1))
			host := api.GoStr(getVal(2))
			port := int32(getVal(3))
			certificates := &TCefX509CertificateArray{count: uint32(getVal(4)), instance: getPtr(5)}
			callback := &ICefSelectClientCertificateCallback{instance: getPtr(6)}
			resultPtr := (*bool)(getPtr(7))
			*resultPtr = fn.(requestHandlerOnSelectClientCertificate)(browser, isProxy, host, port, certificates, callback)
		case requestHandlerOnRenderViewReady:
			browser := &ICefBrowser{instance: getPtr(0)}
			fn.(requestHandlerOnRenderViewReady)(browser)
		case requestHandlerOnRenderProcessTerminated:
			browser := &ICefBrowser{instance: getPtr(0)}
			status := consts.TCefTerminationStatus(getVal(1))
			errorCode := int32(getVal(2))
			error_ := api.GoStr(getVal(3))
			fn.(requestHandlerOnRenderProcessTerminated)(browser, status, errorCode, error_)
		case requestHandlerOnDocumentAvailableInMainFrame:
			browser := &ICefBrowser{instance: getPtr(0)}
			fn.(requestHandlerOnDocumentAvailableInMainFrame)(browser)
		default:
			return false
		}
		return true
	})
}
