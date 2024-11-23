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
	"github.com/energye/energy/v2/consts"
	"unsafe"
)

// / Callback interface for ICefMediaRouter.CreateRoute. The functions of
// / this interface will be called on the browser process UI thread.
// / <para><see cref="uCEFTypes|TCefMediaRouteCreateCallback">Implements TCefMediaRouteCreateCallback</see></para>
// / <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_route_create_callback_t)</see></para>
type ICefMediaRouteCreateCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
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

// / Method that will be executed when the route creation has finished.
// / |result| will be CEF_MRCR_OK if the route creation succeeded. |error| will
// / be a description of the error if the route creation failed. |route| is the
// / resulting route, or NULL if the route creation failed.
func (m *ICefMediaRouteCreateCallback) SetOnMediaRouteCreateFinished(fn onMediaRouteCreateFinished) {

}

type onMediaRouteCreateFinished func(result consts.TCefMediaRouterCreateResult, error string, route ICefMediaRoute)
