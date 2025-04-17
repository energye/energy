//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package cef All CEF implementations of Energy in Go
package cef

import (
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

func init() {
	//var renderLock sync.Mutex
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case GlobalCEFAppEventOnRegCustomSchemes:
			registrar := &TCefSchemeRegistrarRef{instance: getPtr(0)}
			fn.(GlobalCEFAppEventOnRegCustomSchemes)(registrar)
			registrar.Free()
		case GlobalCEFAppEventOnRegisterCustomPreferences:
			fn.(GlobalCEFAppEventOnRegisterCustomPreferences)(consts.TCefPreferencesType(getVal(0)), &TCefPreferenceRegistrarRef{instance: getPtr(1)})
		case GlobalCEFAppEventOnContextInitialized:
			fn.(GlobalCEFAppEventOnContextInitialized)()
		case GlobalCEFAppEventOnBeforeChildProcessLaunch:
			fn.(GlobalCEFAppEventOnBeforeChildProcessLaunch)(&ICefCommandLine{instance: getPtr(0)})
		case GlobalCEFAppEventOnAlreadyRunningAppRelaunchEvent:
			result := (*bool)(getPtr(2))
			*result = fn.(GlobalCEFAppEventOnAlreadyRunningAppRelaunchEvent)(&ICefCommandLine{instance: getPtr(0)}, api.GoStr(getVal(1)))
		case GlobalCEFAppEventOnGetDefaultClient:
			client := (*uintptr)(getPtr(0))
			getClient := &ICefClient{instance: unsafe.Pointer(client)}
			fn.(GlobalCEFAppEventOnGetDefaultClient)(getClient)
			if client != nil {
				*client = uintptr(getClient.instance)
			}
		case GlobalCEFAppEventOnGetLocalizedString:
			stringVal := (*uintptr)(getPtr(1))
			result := (*bool)(getPtr(2))
			resultStringVal := &ResultString{}
			resultBool := &ResultBool{}
			fn.(GlobalCEFAppEventOnGetLocalizedString)(int32(getVal(0)), resultStringVal, resultBool)
			if resultStringVal.Value() != "" {
				*stringVal = api.PascalStr(resultStringVal.Value())
			} else {
				*stringVal = 0
			}
			*result = resultBool.Value()
		case GlobalCEFAppEventOnGetDataResource:
			resultBytes := &ResultBytes{}
			resultData := (*uintptr)(getPtr(1))
			resultDataSize := (*uint32)(getPtr(2))
			result := (*bool)(getPtr(3))
			resultBool := &ResultBool{}
			fn.(GlobalCEFAppEventOnGetDataResource)(int32(getVal(0)), resultBytes, resultBool)
			*result = resultBool.Value()
			if resultBytes.Value() != nil {
				*resultData = uintptr(unsafe.Pointer(&resultBytes.Value()[0]))
				*resultDataSize = uint32(len(resultBytes.Value()))
			} else {
				*resultData = 0
				*resultDataSize = 0
			}
		case GlobalCEFAppEventOnGetDataResourceForScale:
			resultBytes := &ResultBytes{}
			resultData := (*uintptr)(getPtr(2))
			resultDataSize := (*uint32)(getPtr(3))
			result := (*bool)(getPtr(4))
			resultBool := &ResultBool{}
			fn.(GlobalCEFAppEventOnGetDataResourceForScale)(int32(getVal(0)), consts.TCefScaleFactor(getVal(1)), resultBytes, resultBool)
			*result = resultBool.Value()
			if resultBytes.Value() != nil {
				*resultData = uintptr(unsafe.Pointer(&resultBytes.Value()[0]))
				*resultDataSize = uint32(len(resultBytes.Value()))
			} else {
				*resultData = 0
				*resultDataSize = 0
			}
		case GlobalCEFAppEventOnWebKitInitialized:
			fn.(GlobalCEFAppEventOnWebKitInitialized)()
		case GlobalCEFAppEventOnBrowserCreated:
			fn.(GlobalCEFAppEventOnBrowserCreated)(&ICefBrowser{instance: getPtr(0)}, &ICefDictionaryValue{instance: getPtr(1)})
		case GlobalCEFAppEventOnBrowserDestroyed:
			fn.(GlobalCEFAppEventOnBrowserDestroyed)(&ICefBrowser{instance: getPtr(0)})
		case GlobalCEFAppEventOnContextCreated:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			ctx := &ICefV8Context{instance: getPtr(2)}
			fn.(GlobalCEFAppEventOnContextCreated)(browse, frame, ctx)
		case GlobalCEFAppEventOnContextReleased:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			ctx := &ICefV8Context{instance: getPtr(2)}
			fn.(GlobalCEFAppEventOnContextReleased)(browse, frame, ctx)
		case GlobalCEFAppEventOnUncaughtException:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			ctx := &ICefV8Context{instance: getPtr(2)}
			v8Exception := &ICefV8Exception{instance: getPtr(3)}
			v8StackTrace := &ICefV8StackTrace{instance: getPtr(3)}
			fn.(GlobalCEFAppEventOnUncaughtException)(browse, frame, ctx, v8Exception, v8StackTrace)
		case GlobalCEFAppEventOnFocusedNodeChanged:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			node := &ICefDomNode{instance: getPtr(2)}
			fn.(GlobalCEFAppEventOnFocusedNodeChanged)(browse, frame, node)
		case GlobalCEFAppEventOnRenderLoadingStateChange:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			fn.(GlobalCEFAppEventOnRenderLoadingStateChange)(browse, frame, api.GoBool(getVal(2)), api.GoBool(getVal(3)), api.GoBool(getVal(4)))
		case GlobalCEFAppEventOnRenderLoadStart:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			fn.(GlobalCEFAppEventOnRenderLoadStart)(browse, frame, consts.TCefTransitionType(getVal(2)))
		case GlobalCEFAppEventOnRenderLoadEnd:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			fn.(GlobalCEFAppEventOnRenderLoadEnd)(browse, frame, int32(getVal(2)))
		case GlobalCEFAppEventOnRenderLoadError:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			fn.(GlobalCEFAppEventOnRenderLoadError)(browse, frame, consts.TCefErrorCode(getVal(2)), api.GoStr(getVal(3)), api.GoStr(getVal(4)))
		case RenderProcessMessageReceived:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			processId := consts.CefProcessId(getVal(2))
			message := &ICefProcessMessage{instance: getPtr(3)}
			var result = (*bool)(getPtr(4))
			*result = fn.(RenderProcessMessageReceived)(browse, frame, processId, message)
			frame.Free()
			browse.Free()
			message.Free()
		case GlobalCEFAppEventOnScheduleMessagePumpWork:
			fn.(GlobalCEFAppEventOnScheduleMessagePumpWork)(*(*int64)(getPtr(0)))
		default:
			return false
		}
		return true
	})
}
