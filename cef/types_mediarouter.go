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

import "unsafe"

// TODO no impl

// ICefMediaRouter
type ICefMediaRouter struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// ICefMediaRoute
//
//	/include/capi/cef_media_router_capi.h (cef_media_observer_t)
type ICefMediaRoute struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// TCefMediaSinkArray
//
//	[]ICefMediaSink
type TCefMediaSinkArray struct {
	instance  unsafe.Pointer
	mediaSink []*ICefMediaSink
	count     uint32
}

// ICefMediaSink
//
//	/include/capi/cef_media_router_capi.h (cef_media_sink_t)
type ICefMediaSink struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// TCefMediaRouteArray
//
//	[]ICefMediaRoute
type TCefMediaRouteArray struct {
	instance   unsafe.Pointer
	mediaRoute []*ICefMediaRoute
	count      uint32
}
