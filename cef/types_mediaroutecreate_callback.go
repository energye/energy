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

// Callback interface for ICefMediaRouter.CreateRoute. The functions of
// this interface will be called on the browser process UI thread.
// <para><see cref="uCEFTypes|TCefMediaRouteCreateCallback">Implements TCefMediaRouteCreateCallback</see></para>
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_route_create_callback_t)</see></para>
type ICefMediaRouteCreateCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
	ct       consts.CefCreateType
}

// MediaRouteCreateCallbackRef -> ICefMediaRouteCreateCallback
var MediaRouteCreateCallbackRef mediaRouteCreateCallback

type mediaRouteCreateCallback uintptr

func (*mediaRouteCreateCallback) New() *ICefMediaRouteCreateCallback {
	var result uintptr
	imports.Proc(def.MediaRouteCreateCallbackRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaRouteCreateCallback{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (*mediaRouteCreateCallback) NewForChromium(chromium IChromium) *ICefMediaRouteCreateCallback {
	var result uintptr
	imports.Proc(def.MediaRouteCreateCallbackRef_CustomCreate).Call(chromium.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaRouteCreateCallback{instance: unsafe.Pointer(result), ct: consts.CtOther}
	}
	return nil
}

func (m *ICefMediaRouteCreateCallback) IsSelfOwnEvent() bool {
	return m.ct == consts.CtSelfOwn
}

func (m *ICefMediaRouteCreateCallback) IsOtherEvent() bool {
	return m.ct == consts.CtOther
}

// Instance 实例
func (m *ICefMediaRouteCreateCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefMediaRouteCreateCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefMediaRouteCreateCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// Method that will be executed when the route creation has finished.
// |result| will be CEF_MRCR_OK if the route creation succeeded. |error| will
// be a description of the error if the route creation failed. |route| is the
// resulting route, or NULL if the route creation failed.
func (m *ICefMediaRouteCreateCallback) SetOnMediaRouteCreateFinished(fn onMediaRouteCreateFinished) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.MediaRouteCreateCallback_OnMediaRouteCreateFinished).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type onMediaRouteCreateFinished func(result consts.TCefMediaRouterCreateResult, error_ string, route *ICefMediaRoute)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		switch fn.(type) {
		case onMediaRouteCreateFinished:
			result := consts.TCefMediaRouterCreateResult(getVal(0))
			err := api.GoStr(getVal(1))
			route := &ICefMediaRoute{instance: getInstance(getVal(2))}
			fn.(onMediaRouteCreateFinished)(result, err, route)
		default:
			return false
		}
		return true
	})
}
