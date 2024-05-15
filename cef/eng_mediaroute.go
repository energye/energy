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

// ICefMediaRoute Parent: ICefBaseRefCounted
//
//	Represents the route between a media source and sink. Instances of this object are created via ICefMediaRouter.CreateRoute and retrieved via ICefMediaObserver.OnRoutes. Contains the status and metadata of a routing operation. The functions of this interface may be called on any browser process thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_route_t))</a>
type ICefMediaRoute interface {
	ICefBaseRefCounted
	// GetId
	//  Returns the ID for this route.
	GetId() string // function
	// GetSource
	//  Returns the source associated with this route.
	GetSource() ICefMediaSource // function
	// GetSink
	//  Returns the sink associated with this route.
	GetSink() ICefMediaSink // function
	// SendRouteMessage
	//  Send a message over this route. |message_| will be copied if necessary.
	SendRouteMessage(message string) // procedure
	// Terminate
	//  Terminate this route. Will result in an asynchronous call to ICefMediaObserver.OnRoutes on all registered observers.
	Terminate() // procedure
}

// TCefMediaRoute Parent: TCefBaseRefCounted
//
//	Represents the route between a media source and sink. Instances of this object are created via ICefMediaRouter.CreateRoute and retrieved via ICefMediaObserver.OnRoutes. Contains the status and metadata of a routing operation. The functions of this interface may be called on any browser process thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_route_t))</a>
type TCefMediaRoute struct {
	TCefBaseRefCounted
}

// MediaRouteRef -> ICefMediaRoute
var MediaRouteRef mediaRoute

// mediaRoute TCefMediaRoute Ref
type mediaRoute uintptr

func (m *mediaRoute) UnWrap(data uintptr) ICefMediaRoute {
	var resultCefMediaRoute uintptr
	CEF().SysCallN(1058, uintptr(data), uintptr(unsafePointer(&resultCefMediaRoute)))
	return AsCefMediaRoute(resultCefMediaRoute)
}

func (m *TCefMediaRoute) GetId() string {
	r1 := CEF().SysCallN(1053, m.Instance())
	return GoStr(r1)
}

func (m *TCefMediaRoute) GetSource() ICefMediaSource {
	var resultCefMediaSource uintptr
	CEF().SysCallN(1055, m.Instance(), uintptr(unsafePointer(&resultCefMediaSource)))
	return AsCefMediaSource(resultCefMediaSource)
}

func (m *TCefMediaRoute) GetSink() ICefMediaSink {
	var resultCefMediaSink uintptr
	CEF().SysCallN(1054, m.Instance(), uintptr(unsafePointer(&resultCefMediaSink)))
	return AsCefMediaSink(resultCefMediaSink)
}

func (m *TCefMediaRoute) SendRouteMessage(message string) {
	CEF().SysCallN(1056, m.Instance(), PascalStr(message))
}

func (m *TCefMediaRoute) Terminate() {
	CEF().SysCallN(1057, m.Instance())
}
