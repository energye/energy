//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF + Energy 应用程序

package cef

import (
	"github.com/energye/energy/v2/cef/process"
	"github.com/energye/energy/v2/common/imports"
	. "github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/logger"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

var application *TCEFApplication

// TCEFApplication CEF应用对象
type TCEFApplication struct {
	instance unsafe.Pointer
}

// NewApplication 创建CEF应用
//
// 参数: disableRegisDefaultEvent = true 时不会注册默认事件
func NewApplication(disableRegisDefaultEvent ...bool) *TCEFApplication {
	if application == nil {
		var result uintptr
		imports.Proc(internale_CEFApplication_Create).Call(uintptr(unsafe.Pointer(&result)))
		application = &TCEFApplication{instance: unsafe.Pointer(result)}
		if len(disableRegisDefaultEvent) == 0 || !disableRegisDefaultEvent[0] {
			application.registerDefaultEvent()
		}
		application.initDefaultSettings()
	}
	return application
}

// AddCrDelegate MacOSX Delegate
func (m *TCEFApplication) AddCrDelegate() {
	imports.Proc(internale_CEF_AddCrDelegate).Call()
}

//registerDefaultEvent 注册默认事件
func (m *TCEFApplication) registerDefaultEvent() {
	m.defaultSetOnContextCreated()
	m.defaultSetOnProcessMessageReceived()
	m.defaultSetOnRenderLoadStart()
	//m.defaultSetOnBeforeChildProcessLaunch()
	m.defaultSetOnWebKitInitialized()
}

// Instance 实例
func (m *TCEFApplication) Instance() uintptr {
	return uintptr(m.instance)
}

// StartMainProcess 启动主进程
func (m *TCEFApplication) StartMainProcess() bool {
	if m.instance != nil {
		v8init()
		logger.Debug("application single exe,", process.Args.ProcessType(), "process start")
		r1, _, _ := imports.Proc(internale_CEFStartMainProcess).Call(m.Instance())
		return api.GoBool(r1)
	}
	return false
}

// StartSubProcess 启动子进程, 如果指定了子进程执行程序, 将执行指定的子进程程序
func (m *TCEFApplication) StartSubProcess() (result bool) {
	if m.instance != nil {
		v8init()
		logger.Debug("application multiple exe,", process.Args.ProcessType(), "process start")
		r1, _, _ := imports.Proc(internale_CEFStartSubProcess).Call(m.Instance())
		result = api.GoBool(r1)
	}
	return false
}

// RunMessageLoop
func (m *TCEFApplication) RunMessageLoop() {
	defer func() {
		logger.Debug("application run value loop end")
		api.EnergyLibRelease()
	}()
	logger.Debug("application run value loop start")
	imports.Proc(internale_CEFApplication_RunMessageLoop).Call()
}

// QuitMessageLoop 退出消息轮询
func (m *TCEFApplication) QuitMessageLoop() {
	logger.Debug("application quit value loop")
	imports.Proc(internale_CEFApplication_QuitMessageLoop).Call()
}

func (m *TCEFApplication) StopScheduler() {
	imports.Proc(internale_CEFApplication_StopScheduler).Call()
}

func (m *TCEFApplication) Destroy() {
	imports.Proc(internale_CEFApplication_Destroy).Call()
}

func (m *TCEFApplication) Free() {
	if m.instance != nil {
		imports.Proc(internale_CEFApplication_Free).Call()
		m.instance = nil
	}
}

func (m *TCEFApplication) AddCustomCommandLine(commandLine, value string) {
	imports.Proc(internale_AddCustomCommandLine).Call(api.PascalStr(commandLine), api.PascalStr(value))
}

func (m *TCEFApplication) SetOnRegCustomSchemes(fn GlobalCEFAppEventOnRegCustomSchemes) {
	imports.Proc(internale_CEFGlobalApp_SetOnRegCustomSchemes).Call(api.MakeEventDataPtr(fn))
}

// TODO TCefPreferenceRegistrarRef
func (m *TCEFApplication) SetOnRegisterCustomPreferences(fn GlobalCEFAppEventOnRegisterCustomPreferences) {
	imports.Proc(internale_CEFGlobalApp_SetOnRegisterCustomPreferences).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnContextInitialized(fn GlobalCEFAppEventOnContextInitialized) {
	imports.Proc(internale_CEFGlobalApp_SetOnContextInitialized).Call(api.MakeEventDataPtr(fn))
}

// 启动子进程之前自定义命令行参数设置
func (m *TCEFApplication) SetOnBeforeChildProcessLaunch(fn GlobalCEFAppEventOnBeforeChildProcessLaunch) {
	imports.Proc(internale_CEFGlobalApp_SetOnBeforeChildProcessLaunch).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnGetDefaultClient(fn GlobalCEFAppEventOnGetDefaultClient) {
	imports.Proc(internale_CEFGlobalApp_SetOnGetDefaultClient).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnGetLocalizedString(fn GlobalCEFAppEventOnGetLocalizedString) {
	imports.Proc(internale_CEFGlobalApp_SetOnGetLocalizedString).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnGetDataResource(fn GlobalCEFAppEventOnGetDataResource) {
	imports.Proc(internale_CEFGlobalApp_SetOnGetDataResource).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnGetDataResourceForScale(fn GlobalCEFAppEventOnGetDataResourceForScale) {
	imports.Proc(internale_CEFGlobalApp_SetOnGetDataResourceForScale).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnWebKitInitialized(fn GlobalCEFAppEventOnWebKitInitialized) {
	imports.Proc(internale_CEFGlobalApp_SetOnWebKitInitialized).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) defaultSetOnWebKitInitialized() {
	m.SetOnWebKitInitialized(func() {})
}

func (m *TCEFApplication) SetOnBrowserCreated(fn GlobalCEFAppEventOnBrowserCreated) {
	imports.Proc(internale_CEFGlobalApp_SetOnBrowserCreated).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnBrowserDestroyed(fn GlobalCEFAppEventOnBrowserDestroyed) {
	imports.Proc(internale_CEFGlobalApp_SetOnBrowserDestroyed).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnContextCreated(fn GlobalCEFAppEventOnContextCreated) {
	imports.Proc(internale_CEFGlobalApp_SetOnContextCreated).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) defaultSetOnContextCreated() {
	m.SetOnContextCreated(func(browse *ICefBrowser, frame *ICefFrame, context *ICefV8Context) bool {
		return false
	})
}

func (m *TCEFApplication) SetOnContextReleased(fn GlobalCEFAppEventOnContextReleased) {
	imports.Proc(internale_CEFGlobalApp_SetOnContextReleased).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnUncaughtException(fn GlobalCEFAppEventOnUncaughtException) {
	imports.Proc(internale_CEFGlobalApp_SetOnUncaughtException).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnFocusedNodeChanged(fn GlobalCEFAppEventOnFocusedNodeChanged) {
	imports.Proc(internale_CEFGlobalApp_SetOnFocusedNodeChanged).Call(api.MakeEventDataPtr(fn))
}

// 进程间通信处理消息接收
func (m *TCEFApplication) SetOnProcessMessageReceived(fn RenderProcessMessageReceived) {
	imports.Proc(internale_CEFGlobalApp_SetOnProcessMessageReceived).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) defaultSetOnProcessMessageReceived() {
	m.SetOnProcessMessageReceived(func(browse *ICefBrowser, frame *ICefFrame, sourceProcess CefProcessId, processMessage *ICefProcessMessage) bool {
		return false
	})
}

//func (m *TCEFApplication) defaultSetOnBeforeChildProcessLaunch() {
//	m.SetOnBeforeChildProcessLaunch(func(commandLine *TCefCommandLine) {})
//}

//func (m *TCEFApplication) SetOnScheduleMessagePumpWork(fn ) {
//	imports.Proc(internale_CEFGlobalApp_SetOnScheduleMessagePumpWork).Call(api.MakeEventDataPtr(fn))
//}

func (m *TCEFApplication) SetOnRenderLoadingStateChange(fn GlobalCEFAppEventOnRenderLoadingStateChange) {
	imports.Proc(internale_CEFGlobalApp_SetOnRenderLoadingStateChange).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnRenderLoadStart(fn GlobalCEFAppEventOnRenderLoadStart) {
	imports.Proc(internale_CEFGlobalApp_SetOnRenderLoadStart).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) defaultSetOnRenderLoadStart() {
	m.SetOnRenderLoadStart(func(browser *ICefBrowser, frame *ICefFrame, transitionType TCefTransitionType) {

	})
}

func (m *TCEFApplication) SetOnRenderLoadEnd(fn GlobalCEFAppEventOnRenderLoadEnd) {
	imports.Proc(internale_CEFGlobalApp_SetOnRenderLoadEnd).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnRenderLoadError(fn GlobalCEFAppEventOnRenderLoadError) {
	imports.Proc(internale_CEFGlobalApp_SetOnRenderLoadError).Call(api.MakeEventDataPtr(fn))
}

func init() {
	//var renderLock sync.Mutex
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case GlobalCEFAppEventOnRegCustomSchemes:
			fn.(GlobalCEFAppEventOnRegCustomSchemes)(&TCefSchemeRegistrarRef{instance: getPtr(0)})
		case GlobalCEFAppEventOnRegisterCustomPreferences:
			fn.(GlobalCEFAppEventOnRegisterCustomPreferences)(TCefPreferencesType(getVal(0)), &TCefPreferenceRegistrarRef{instance: getPtr(1)})
		case GlobalCEFAppEventOnContextInitialized:
			fn.(GlobalCEFAppEventOnContextInitialized)()
		case GlobalCEFAppEventOnBeforeChildProcessLaunch:
			//commands := (*uintptr)(getPtr(0))
			//commandLine := &TCefCommandLine{commandLines: make(map[string]string)}
			fn.(GlobalCEFAppEventOnBeforeChildProcessLaunch)(&ICefCommandLine{instance: getPtr(0)})
			//*commands = api.PascalStr(commandLine.toString())
		case GlobalCEFAppEventOnGetDefaultClient:
			client := (*uintptr)(getPtr(0))
			getClient := &ICefClient{instance: unsafe.Pointer(client)}
			fn.(GlobalCEFAppEventOnGetDefaultClient)(getClient)
			*client = uintptr(getClient.instance)
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
			fn.(GlobalCEFAppEventOnGetDataResourceForScale)(int32(getVal(0)), TCefScaleFactor(getVal(1)), resultBytes, resultBool)
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
			appWebKitInitialized()
		case GlobalCEFAppEventOnBrowserCreated:
			fn.(GlobalCEFAppEventOnBrowserCreated)(&ICefBrowser{instance: getPtr(0)}, &ICefDictionaryValue{instance: getPtr(1)})
		case GlobalCEFAppEventOnBrowserDestroyed:
			fn.(GlobalCEFAppEventOnBrowserDestroyed)(&ICefBrowser{instance: getPtr(0)})
		case GlobalCEFAppEventOnContextCreated:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			ctx := &ICefV8Context{instance: getPtr(2)}
			var result = fn.(GlobalCEFAppEventOnContextCreated)(browse, frame, ctx)
			if !result {
				appOnContextCreated(browse, frame, ctx)
			}
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
			if ipcRender != nil {
				ipcRender.clear()
			}
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			fn.(GlobalCEFAppEventOnRenderLoadStart)(browse, frame, TCefTransitionType(getVal(2)))
		case GlobalCEFAppEventOnRenderLoadEnd:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			fn.(GlobalCEFAppEventOnRenderLoadEnd)(browse, frame, int32(getVal(2)))
		case GlobalCEFAppEventOnRenderLoadError:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			fn.(GlobalCEFAppEventOnRenderLoadError)(browse, frame, TCefErrorCode(getVal(2)), api.GoStr(getVal(3)), api.GoStr(getVal(4)))
		case RenderProcessMessageReceived:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			processId := CefProcessId(getVal(2))
			message := &ICefProcessMessage{instance: getPtr(3)}
			var result = (*bool)(getPtr(4))
			*result = renderProcessMessageReceived(browse, frame, processId, message)
			if !*result {
				*result = fn.(RenderProcessMessageReceived)(browse, frame, processId, message)
			}
			frame.Free()
			browse.Free()
			message.Free()
		default:
			return false
		}
		return true
	})
}
