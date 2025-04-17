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

// Callback interface for ICefMediaSink.GetDeviceInfo. The functions of
// this interface will be called on the browser process UI thread.
// <para><see cref="uCEFTypes|TCefMediaSinkDeviceInfoCallback">Implements TCefMediaSinkDeviceInfoCallback</see></para>
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_sink_device_info_callback_t)</see></para>
type ICefMediaSinkDeviceInfoCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
	ct       consts.CefCreateType
}

// MediaSinkDeviceInfoCallbackRef -> ICefMediaSinkDeviceInfoCallback
var MediaSinkDeviceInfoCallbackRef mediaSinkDeviceInfoCallback

type mediaSinkDeviceInfoCallback uintptr

func (*mediaSinkDeviceInfoCallback) New() *ICefMediaSinkDeviceInfoCallback {
	var result uintptr
	imports.Proc(def.MediaSinkDeviceInfoCallbackRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaSinkDeviceInfoCallback{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (*mediaSinkDeviceInfoCallback) NewForChromium(chromium IChromium) *ICefMediaSinkDeviceInfoCallback {
	var result uintptr
	imports.Proc(def.MediaSinkDeviceInfoCallbackRef_CustomCreate).Call(chromium.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMediaSinkDeviceInfoCallback{instance: unsafe.Pointer(result), ct: consts.CtOther}
	}
	return nil
}

func (m *ICefMediaSinkDeviceInfoCallback) IsSelfOwnEvent() bool {
	return m.ct == consts.CtSelfOwn
}

func (m *ICefMediaSinkDeviceInfoCallback) IsOtherEvent() bool {
	return m.ct == consts.CtOther
}

// Instance 实例
func (m *ICefMediaSinkDeviceInfoCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefMediaSinkDeviceInfoCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefMediaSinkDeviceInfoCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

// Method that will be executed asyncronously once device information has  been retrieved.
func (m *ICefMediaRouteCreateCallback) SetOnMediaSinkDeviceInfo(fn onOnMediaSinkDeviceInfo) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.MediaSinkDeviceInfoCallback_OnMediaSinkDeviceInfo).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type onOnMediaSinkDeviceInfo func(ipAddress string, port int32, modelName string)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		switch fn.(type) {
		case onOnMediaSinkDeviceInfo:
			ipAddress := api.GoStr(getVal(0))
			port := int32(getVal(1))
			modelName := api.GoStr(getVal(2))
			fn.(onOnMediaSinkDeviceInfo)(ipAddress, port, modelName)
		default:
			return false
		}
		return true
	})
}
