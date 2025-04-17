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

// ICefMediaRoute
//
//	Represents the route between a media source and sink. Instances of this
//	object are created via ICefMediaRouter.CreateRoute and retrieved via
//	ICefMediaObserver.OnRoutes. Contains the status and metadata of a
//	routing operation. The func (m *ICefMediaRouter)s of this interface may be called on any
//	browser process thread unless otherwise indicated.
//	<para><see cref="uCEFTypes|TCefMediaRoute">Implements TCefMediaRoute</see></para>
//	<para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_route_t)</see></para>
type ICefMediaRoute struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// MediaRouteRef -> ICefMediaRoute
var MediaRouteRef mediaRoute

type mediaRoute uintptr

func (*mediaRoute) UnWrap(delegate *ICefMediaRoute) *ICefMediaRoute {
	var result uintptr
	imports.Proc(def.MediaRouteRef_UnWrap).Call(delegate.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaRoute{instance: getInstance(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefMediaRoute) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefMediaRoute) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefMediaRoute) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// Returns the ID for this route.
func (m *ICefMediaRoute) GetId() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.MediaRoute_GetId).Call(m.Instance())
	return api.GoStr(r1)
}

// Returns the source associated with this route.
func (m *ICefMediaRoute) GetSource() *ICefMediaSource {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.MediaRoute_GetSource).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaSource{instance: getInstance(result)}
	}
	return nil
}

// Returns the sour
// Returns the sink associated with this route.
func (m *ICefMediaRoute) GetSink() *ICefMediaSink {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.MediaRoute_GetSink).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaSink{instance: getInstance(result)}
	}
	return nil
}

// Returns the sour
// Send a message over this route. |message_| will be copied if necessary.
func (m *ICefMediaRoute) SendRouteMessage(message string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MediaRoute_SendRouteMessage).Call(m.Instance(), api.PascalStr(message))
}

// Returns the sour
// Terminate this route. Will result in an asynchronous call to
// ICefMediaObserver.OnRoutes on all registered observers.
func (m *ICefMediaRoute) Terminate() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MediaRoute_Terminate).Call(m.Instance())
}

/// Returns the sour

// TCefMediaRouteArray
//
//	of []ICefMediaRoute
type TCefMediaRouteArray struct {
	instance   unsafe.Pointer
	mediaRoute []*ICefMediaRoute
	count      uint32
}

// Instance 实例
func (m *TCefMediaRouteArray) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *TCefMediaRouteArray) Free() {
	if m.instance != nil {
		m.instance = nil
		m.count = 0
		m.mediaRoute = nil
	}
}

func (m *TCefMediaRouteArray) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *TCefMediaRouteArray) Get(index int) *ICefMediaRoute {
	if !m.IsValid() && index >= 0 && index < int(m.count) {
		return nil
	}
	var result uintptr
	imports.Proc(def.MediaRouteArray_Get).Call(m.Instance(), uintptr(index), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaRoute{instance: getInstance(result)}
	}
	return nil
}
