//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"fmt"
	. "github.com/energye/energy/commons"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"strings"
	"unsafe"
)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		defer func() {
			if err := recover(); err != nil {
				logger.Logger.Error("v8event Error:", err)
			}
		}()
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case GlobalCEFAppEventOnBrowserDestroyed:
			fn.(GlobalCEFAppEventOnBrowserDestroyed)(&ICefBrowser{browseId: int32(getVal(0))})
		case GlobalCEFAppEventOnRenderLoadStart:
			browser := &ICefBrowser{browseId: int32(getVal(0))}
			tempFrame := (*cefFrame)(getPtr(1))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			fn.(GlobalCEFAppEventOnRenderLoadStart)(browser, frame, TCefTransitionType(getVal(2)))
		case GlobalCEFAppEventOnRenderLoadEnd:
			browser := &ICefBrowser{browseId: int32(getVal(0))}
			tempFrame := (*cefFrame)(getPtr(1))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			fn.(GlobalCEFAppEventOnRenderLoadEnd)(browser, frame, int32(getVal(2)))
		case GlobalCEFAppEventOnRenderLoadError:
			browser := &ICefBrowser{browseId: int32(getVal(0))}
			tempFrame := (*cefFrame)(getPtr(1))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			fn.(GlobalCEFAppEventOnRenderLoadError)(browser, frame, TCefErrorCode(getVal(2)), api.DStrToGoStr(getVal(3)), api.DStrToGoStr(getVal(4)))
		case GlobalCEFAppEventOnRenderLoadingStateChange:
			browser := &ICefBrowser{browseId: int32(getVal(0))}
			tempFrame := (*cefFrame)(getPtr(1))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			fn.(GlobalCEFAppEventOnRenderLoadingStateChange)(browser, frame, api.DBoolToGoBool(getVal(2)), api.DBoolToGoBool(getVal(3)), api.DBoolToGoBool(getVal(4)))
		case RenderProcessMessageReceived:
			browser := &ICefBrowser{browseId: int32(getVal(0))}
			tempFrame := (*cefFrame)(getPtr(1))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			cefProcMsg := (*ipc.CefProcessMessagePtr)(getPtr(3))
			args := ipc.NewArgumentList()
			args.UnPackageBytePtr(cefProcMsg.Data, int32(cefProcMsg.DataLen))
			processMessage := &ipc.ICefProcessMessage{
				Name:         api.DStrToGoStr(cefProcMsg.Name),
				ArgumentList: args,
			}
			var result = (*bool)(getPtr(4))
			*result = fn.(RenderProcessMessageReceived)(browser, frame, CefProcessId(getVal(2)), processMessage)
			args.Clear()
			cefProcMsg.Data = 0
			cefProcMsg.DataLen = 0
			cefProcMsg.Name = 0
			cefProcMsg = nil
			args = nil
		case GlobalCEFAppEventOnContextCreated:
			browser := &ICefBrowser{browseId: int32(getVal(0))}
			tempFrame := (*cefFrame)(getPtr(1))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.DStrToGoStr(tempFrame.Name),
				Url:     api.DStrToGoStr(tempFrame.Url),
				Id:      StrToInt64(api.DStrToGoStr(tempFrame.Identifier)),
			}
			if strings.Index(frame.Url, "devtools://") == 0 {
				processName = PT_DEVTOOLS
				return true
			} else {
				processName = Args.ProcessType()
			}
			v8ctx := (*cefV8Context)(getPtr(2))
			ctx := &ICefV8Context{
				Browser: browser,
				Frame:   frame,
				Global:  &ICEFv8Value{instance: v8ctx.Global, ptr: unsafe.Pointer(v8ctx.Global)},
			}
			var status = (*bool)(getPtr(3))
			//用户定义返回 false 创建 render IPC 及 变量绑定
			var result = fn.(GlobalCEFAppEventOnContextCreated)(browser, frame, ctx)
			if !result {
				cefAppContextCreated(browser, frame)
				*status = true
			} else {
				*status = false
			}
		case GlobalCEFAppEventOnWebKitInitialized:
			fn.(GlobalCEFAppEventOnWebKitInitialized)()
		case GlobalCEFAppEventOnBeforeChildProcessLaunch:
			commands := (*uintptr)(getPtr(0))
			commandLine := &TCefCommandLine{commandLines: make(map[string]string)}
			ipc.IPC.SetPort()
			commandLine.AppendSwitch(MAINARGS_NETIPCPORT, fmt.Sprintf("%d", ipc.IPC.Port()))
			fn.(GlobalCEFAppEventOnBeforeChildProcessLaunch)(commandLine)
			*commands = api.GoStrToDStr(commandLine.toString())
		default:
			return false
		}
		return true
	})
}

func _SetCEFCallbackEvent(fnName CEF_ON_EVENTS, fn interface{}) {
	var eventId = api.GetAddEventToMapFn()(CommonPtr.Instance(), fn)
	//Logger.Debug("CEFApplication Event name:", fnName, "eventId:", eventId, "commonInstance.instance:", commonInstance.instance)
	Proc("SetCEFCallbackEvent").Call(api.GoStrToDStr(string(fnName)), eventId)
}
