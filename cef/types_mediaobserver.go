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

// / Implemented by the client to observe MediaRouter events and registered via
// / ICefMediaRouter.AddObserver. The functions of this interface will be
// / called on the browser process UI thread.
// / <para><see cref="uCEFTypes|TCefMediaObserver">Implements TCefMediaObserver</see></para>
// / <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_observer_t)</see></para>
type ICefMediaObserver struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
	ct       consts.CefCreateType
}

// MediaObserverRef -> ICefMediaObserver
var MediaObserverRef mediaObserver

type mediaObserver uintptr

func (*mediaObserver) New() *ICefMediaObserver {
	var result uintptr
	imports.Proc(def.MediaObserverRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaObserver{instance: getInstance(result)}
	}
	return nil
}
func (*mediaObserver) NewForChromium(chromium IChromium) *ICefMediaObserver {
	var result uintptr
	imports.Proc(def.MediaObserverRef_CustomCreate).Call(chromium.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaObserver{instance: getInstance(result), ct: consts.CtOther}
	}
	return nil
}

// Instance 实例
func (m *ICefMediaObserver) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefMediaObserver) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefMediaObserver) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefMediaObserver) IsSelfOwnEvent() bool {
	return m.ct == consts.CtSelfOwn
}

func (m *ICefMediaObserver) IsOtherEvent() bool {
	return m.ct == consts.CtOther
}

func (m *ICefMediaObserver) SetOnSinks(fn onSinks) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.MediaObserver_OnSinks).Call(m.Instance(), api.MakeEventDataPtr(fn))
}
func (m *ICefMediaObserver) SetOnRoutes(fn onRoutes) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.MediaObserver_OnRoutes).Call(m.Instance(), api.MakeEventDataPtr(fn))
}
func (m *ICefMediaObserver) SetOnRouteStateChanged(fn onRouteStateChanged) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.MediaObserver_OnRouteStateChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}
func (m *ICefMediaObserver) SetOnRouteMessageReceived(fn onRouteMessageReceived) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.MediaObserver_OnRouteMessageReceived).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type onSinks func(sinks *TCefMediaSinkArray)

type onRoutes func(routes *TCefMediaRouteArray)

type onRouteStateChanged func(route *ICefMediaRoute, state consts.TCefMediaRouteConnectionState)

type onRouteMessageReceived func(route *ICefMediaRoute, message string)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		switch fn.(type) {
		case onSinks:
			sinks := &TCefMediaSinkArray{instance: getInstance(getVal(0)), count: uint32(getVal(1))}
			fn.(onSinks)(sinks)
		case onRoutes:
			routes := &TCefMediaRouteArray{instance: getInstance(getVal(0)), count: uint32(getVal(1))}
			fn.(onRoutes)(routes)
		case onRouteStateChanged:
			route := &ICefMediaRoute{instance: getInstance(getVal(0))}
			fn.(onRouteStateChanged)(route, consts.TCefMediaRouteConnectionState(getVal(1)))
		case onRouteMessageReceived:
			route := &ICefMediaRoute{instance: getInstance(getVal(0))}
			fn.(onRouteMessageReceived)(route, api.GoStr(getVal(1)))
		default:
			return false
		}
		return true
	})
}
