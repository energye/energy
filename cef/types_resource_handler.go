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
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefResourceHandler
//
//	/include/capi/cef_resource_handler_capi.h (cef_resource_handler_t)
type ICefResourceHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// ResourceHandlerRef -> ICefResourceHandler
var ResourceHandlerRef resourceHandler

type resourceHandler uintptr

func (*resourceHandler) New(browser *ICefBrowser, frame *ICefFrame, schemeName string, request *ICefRequest) *ICefResourceHandler {
	var result uintptr
	imports.Proc(def.ResourceHandlerRef_Create).Call(browser.Instance(), frame.Instance(), api.PascalStr(schemeName), request.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefResourceHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefResourceHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefResourceHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefResourceHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefResourceHandler) Open(fn resourceHandlerOpen) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ResourceHandler_Open).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefResourceHandler) GetResponseHeaders(fn resourceHandlerGetResponseHeaders) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ResourceHandler_GetResponseHeaders).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefResourceHandler) Skip(fn resourceHandlerSkip) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ResourceHandler_Skip).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefResourceHandler) Read(fn resourceHandlerRead) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ResourceHandler_Read).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ProcessRequest
//
// Deprecated: Use Open instead.
func (m *ICefResourceHandler) ProcessRequest(fn resourceHandlerProcessRequest) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ResourceHandler_ProcessRequest).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ReadResponse
//
// Deprecated: Use Skip and Read instead.
func (m *ICefResourceHandler) ReadResponse(fn resourceHandlerReadResponse) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ResourceHandler_ReadResponse).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefResourceHandler) Cancel(fn resourceHandlerCancel) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ResourceHandler_Cancel).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type resourceHandlerOpen func(request *ICefRequest, callback *ICefCallback) (handleRequest, result bool)                             // Open
type resourceHandlerGetResponseHeaders func(response *ICefResponse) (responseLength int64, redirectUrl string)                       // GetResponseHeaders
type resourceHandlerSkip func(bytesToSkip int64, callback *ICefResourceSkipCallback) (bytesSkipped int64, result bool)               // Skip
type resourceHandlerRead func(dataOut uintptr, bytesToRead int32, callback *ICefResourceReadCallback) (bytesRead int32, result bool) // Read
type resourceHandlerProcessRequest func(request *ICefRequest, callback *ICefCallback) bool                                           // ProcessRequest deprecated
type resourceHandlerReadResponse func(dataOut uintptr, bytesToRead int32, callback *ICefCallback) (bytesRead int32, result bool)     // ReadResponse deprecated
type resourceHandlerCancel func()                                                                                                    // Cancel

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case resourceHandlerOpen:
			request := &ICefRequest{instance: getPtr(0)}
			handleRequestPtr := (*bool)(getPtr(1))
			callback := &ICefCallback{instance: getPtr(2)}
			resultPtr := (*bool)(getPtr(3))
			*handleRequestPtr, *resultPtr = fn.(resourceHandlerOpen)(request, callback)
		case resourceHandlerGetResponseHeaders:
			response := &ICefResponse{instance: getPtr(0)}
			responseLengthPtr := (*int64)(getPtr(1))
			redirectUrlPtr := (*uintptr)(getPtr(2))
			responseLength, redirectUrl := fn.(resourceHandlerGetResponseHeaders)(response)
			*responseLengthPtr = responseLength
			*redirectUrlPtr = api.PascalStr(redirectUrl)
		case resourceHandlerSkip:
			bytesToSkip := *(*int64)(getPtr(0))
			bytesSkippedPtr := (*int64)(getPtr(1))
			callback := &ICefResourceSkipCallback{instance: getPtr(2)}
			resultPtr := (*bool)(getPtr(3))
			bytesSkipped, result := fn.(resourceHandlerSkip)(bytesToSkip, callback)
			*bytesSkippedPtr = bytesSkipped
			*resultPtr = result
		case resourceHandlerRead:
			dataOut := getVal(0)
			bytesToRead := int32(getVal(1))
			bytesReadPtr := (*int32)(getPtr(2))
			callback := &ICefResourceReadCallback{instance: getPtr(3)}
			resultPtr := (*bool)(getPtr(4))
			bytesRead, result := fn.(resourceHandlerRead)(dataOut, bytesToRead, callback)
			*bytesReadPtr = bytesRead
			*resultPtr = result
		case resourceHandlerProcessRequest:
			request := &ICefRequest{instance: getPtr(0)}
			callback := &ICefCallback{instance: getPtr(1)}
			resultPtr := (*bool)(getPtr(2))
			result := fn.(resourceHandlerProcessRequest)(request, callback)
			*resultPtr = result
		case resourceHandlerReadResponse:
			dataOut := getVal(0)
			bytesToRead := int32(getVal(1))
			bytesReadPtr := (*int32)(getPtr(2))
			callback := &ICefCallback{instance: getPtr(3)}
			resultPtr := (*bool)(getPtr(4))
			bytesRead, result := fn.(resourceHandlerReadResponse)(dataOut, bytesToRead, callback)
			*bytesReadPtr = bytesRead
			*resultPtr = result
		case resourceHandlerCancel:
			fn.(resourceHandlerCancel)()
		default:
			return false
		}
		return true
	})
}
