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
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefMediaRouter
//
//	Supports discovery of and communication with media devices on the local
//	network via the Cast and DIAL protocols. The func (m *ICefMediaRouter)s of this interface may
//	be called on any browser process thread unless otherwise indicated.
//	<para><see cref="uCEFTypes|TCefMediaRouter">Implements TCefMediaRouter</see></para>
//	<para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_router_t)</see></para>
type ICefMediaRouter struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// MediaRouterRef -> ICefPanel
var MediaRouterRef mediaRouter

type mediaRouter uintptr

func (*mediaRouter) UnWrap(data *ICefMediaRouter) *ICefMediaRouter {
	var result uintptr
	imports.Proc(def.MediaRouterRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaRouter{instance: getInstance(result)}
	}
	return nil
}

func (*mediaRouter) Global() *ICefMediaRouter {
	var result uintptr
	imports.Proc(def.MediaRouterRef_Global).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaRouter{instance: getInstance(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefMediaRouter) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefMediaRouter) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefMediaRouter) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// Add an observer for MediaRouter events. The observer will remain
// registered until the returned Registration object is destroyed.
func (m *ICefMediaRouter) AddObserver(observer *ICefMediaObserver) *ICefRegistration {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.MediaRouter_AddObserver).Call(m.Instance(), observer.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRegistration{instance: getInstance(result)}
	}
	return nil
}

// Returns a MediaSource object for the specified media source URN. Supported
// URN schemes include "cast:" and "dial:", and will be already known by the
// client application (e.g. "cast:<appId>?clientId=<clientId>").
func (m *ICefMediaRouter) GetSource(urn string) *ICefMediaSource {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.MediaRouter_GetSource).Call(m.Instance(), api.PascalStr(urn), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaSource{instance: getInstance(result)}
	}
	return nil
}

// Trigger an asynchronous call to ICefMediaObserver.OnSinks on all
// registered observers.
func (m *ICefMediaRouter) NotifyCurrentSinks() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MediaRouter_NotifyCurrentSinks).Call(m.Instance())
}

// Create a new route between |source| and |sink|. Source and sink must be
// valid, compatible (as reported by ICefMediaSink.IsCompatibleWith), and
// a route between them must not already exist. |callback| will be executed
// on success or failure. If route creation succeeds it will also trigger an
// asynchronous call to ICefMediaObserver.OnRoutes on all registered
// observers.
func (m *ICefMediaRouter) CreateRoute(source *ICefMediaSource, sink *ICefMediaSink, callback *ICefMediaRouteCreateCallback) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MediaRouter_CreateRoute).Call(m.Instance(), source.Instance(), sink.Instance(), callback.Instance())
}

// Trigger an asynchronous call to ICefMediaObserver.OnRoutes on all
// registered observers.
func (m *ICefMediaRouter) NotifyCurrentRoutes() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MediaRouter_NotifyCurrentRoutes).Call(m.Instance())
}
