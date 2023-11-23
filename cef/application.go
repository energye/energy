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
	"github.com/energye/energy/v2/cef/internal/cef"
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/cef/process"
	"github.com/energye/energy/v2/common/imports"
	. "github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/logger"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

var application *TCEFApplication

// TCEFApplication CEF应用对象
type TCEFApplication struct {
	instance                 unsafe.Pointer
	ui                       UITool
	onContextCreated         GlobalCEFAppEventOnContextCreated
	onProcessMessageReceived RenderProcessMessageReceived
	onWebKitInitialized      GlobalCEFAppEventOnWebKitInitialized
	onRegCustomSchemes       GlobalCEFAppEventOnRegCustomSchemes
	onRenderLoadStart        GlobalCEFAppEventOnRenderLoadStart
}

// NewApplication 创建CEF应用
//
// 参数: disableRegisDefaultEvent = true 时不会注册默认事件
func NewApplication(disableRegisDefaultEvent ...bool) *TCEFApplication {
	if application == nil {
		application = CreateApplication()
		if len(disableRegisDefaultEvent) == 0 || !disableRegisDefaultEvent[0] {
			application.registerDefaultEvent()
		}
		application.initDefaultSettings()
		// 将应用设置到内部实现
		cef.SetApplication(application)
	}
	return application
}

// CreateApplication
//
//	创建CEF Application
//	初始化CEF时必须创建，多进程模式每个application配置都应该相同
func CreateApplication() *TCEFApplication {
	var result uintptr
	imports.Proc(def.CEFApplication_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &TCEFApplication{instance: unsafe.Pointer(result)}
}

// AddCrDelegate MacOSX Delegate
func (m *TCEFApplication) AddCrDelegate() {
	imports.Proc(def.CEF_AddCrDelegate).Call()
}

// registerDefaultEvent 注册默认事件
func (m *TCEFApplication) registerDefaultEvent() {
	m.defaultSetOnContextCreated()
	m.defaultSetOnProcessMessageReceived()
	m.defaultSetOnWebKitInitialized()
	m.defaultSetOnRegCustomSchemes()
	m.defaultSetOnRenderLoadStart()
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
		r1, _, _ := imports.Proc(def.CEFStartMainProcess).Call(m.Instance())
		return api.GoBool(r1)
	}
	return false
}

// StartSubProcess 启动子进程, 如果指定了子进程执行程序, 将执行指定的子进程程序
func (m *TCEFApplication) StartSubProcess() (result bool) {
	if m.instance != nil {
		v8init()
		logger.Debug("application multiple exe,", process.Args.ProcessType(), "process start")
		r1, _, _ := imports.Proc(def.CEFStartSubProcess).Call(m.Instance())
		result = api.GoBool(r1)
	}
	return false
}

// DoMessageLoopWork
func (m *TCEFApplication) DoMessageLoopWork() {
	imports.Proc(def.CEFApplication_DoMessageLoopWork).Call()
}

// RunMessageLoop
func (m *TCEFApplication) RunMessageLoop() {
	defer func() {
		logger.Debug("application run value loop end")
		api.EnergyLibRelease()
	}()
	logger.Debug("application run value loop start")
	imports.Proc(def.CEFApplication_RunMessageLoop).Call()
}

// QuitMessageLoop 退出消息轮询
func (m *TCEFApplication) QuitMessageLoop() {
	logger.Debug("application quit value loop")
	imports.Proc(def.CEFApplication_QuitMessageLoop).Call()
}

func (m *TCEFApplication) StopScheduler() {
	GlobalWorkScheduler.StopScheduler()
}

func (m *TCEFApplication) Destroy() {
	imports.Proc(def.CEFApplication_Destroy).Call()
}

func (m *TCEFApplication) Free() {
	if m.instance != nil {
		imports.Proc(def.CEFApplication_Free).Call()
		m.instance = nil
	}
}

// AddCustomCommandLine
//
//	添加自定义进程启动时添加的命令行参数
func (m *TCEFApplication) AddCustomCommandLine(commandLine, value string) {
	imports.Proc(def.AddCustomCommandLine).Call(api.PascalStr(commandLine), api.PascalStr(value))
}

// SetOnRegCustomSchemes
//
//	自定义协议注册回调函数
func (m *TCEFApplication) SetOnRegCustomSchemes(fn GlobalCEFAppEventOnRegCustomSchemes) {
	m.onRegCustomSchemes = fn
}

func (m *TCEFApplication) setOnRegCustomSchemes(fn GlobalCEFAppEventOnRegCustomSchemes) {
	imports.Proc(def.CEFGlobalApp_SetOnRegCustomSchemes).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) defaultSetOnRegCustomSchemes() {
	m.setOnRegCustomSchemes(func(registrar *TCefSchemeRegistrarRef) {
		regCustomSchemes(registrar)
		if m.onRegCustomSchemes != nil {
			m.onRegCustomSchemes(registrar)
		}
	})
}

// SetOnRegisterCustomPreferences
//
//	TODO 该函数还未完全实现
func (m *TCEFApplication) SetOnRegisterCustomPreferences(fn GlobalCEFAppEventOnRegisterCustomPreferences) {
	imports.Proc(def.CEFGlobalApp_SetOnRegisterCustomPreferences).Call(api.MakeEventDataPtr(fn))
}

// SetOnContextInitialized
//
//	上下文初始化
func (m *TCEFApplication) SetOnContextInitialized(fn GlobalCEFAppEventOnContextInitialized) {
	imports.Proc(def.CEFGlobalApp_SetOnContextInitialized).Call(api.MakeEventDataPtr(fn))
}

// SetOnBeforeChildProcessLaunch
//
//	启动子进程之前自定义命令行参数设置
func (m *TCEFApplication) SetOnBeforeChildProcessLaunch(fn GlobalCEFAppEventOnBeforeChildProcessLaunch) {
	imports.Proc(def.CEFGlobalApp_SetOnBeforeChildProcessLaunch).Call(api.MakeEventDataPtr(fn))
}

// SetOnGetDefaultClient
//
//	获取并返回CefClient, 我们自己创建并返回到 *ICefClient = myCefClient
func (m *TCEFApplication) SetOnGetDefaultClient(fn GlobalCEFAppEventOnGetDefaultClient) {
	imports.Proc(def.CEFGlobalApp_SetOnGetDefaultClient).Call(api.MakeEventDataPtr(fn))
}

// SetOnGetLocalizedString
//
//	获取并返回本地化
func (m *TCEFApplication) SetOnGetLocalizedString(fn GlobalCEFAppEventOnGetLocalizedString) {
	imports.Proc(def.CEFGlobalApp_SetOnGetLocalizedString).Call(api.MakeEventDataPtr(fn))
}

// SetOnGetDataResource
//
//	获取并返回本地资源
func (m *TCEFApplication) SetOnGetDataResource(fn GlobalCEFAppEventOnGetDataResource) {
	imports.Proc(def.CEFGlobalApp_SetOnGetDataResource).Call(api.MakeEventDataPtr(fn))
}

// SetOnGetDataResourceForScale
//
//	获取并返回本地资源大小
func (m *TCEFApplication) SetOnGetDataResourceForScale(fn GlobalCEFAppEventOnGetDataResourceForScale) {
	imports.Proc(def.CEFGlobalApp_SetOnGetDataResourceForScale).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnWebKitInitialized(fn GlobalCEFAppEventOnWebKitInitialized) {
	m.onWebKitInitialized = fn
}

func (m *TCEFApplication) setOnWebKitInitialized(fn GlobalCEFAppEventOnWebKitInitialized) {
	imports.Proc(def.CEFGlobalApp_SetOnWebKitInitialized).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) defaultSetOnWebKitInitialized() {
	m.setOnWebKitInitialized(func() {
		appWebKitInitialized()
		if m.onWebKitInitialized != nil {
			m.onWebKitInitialized()
		}
	})
}

func (m *TCEFApplication) SetOnBrowserCreated(fn GlobalCEFAppEventOnBrowserCreated) {
	imports.Proc(def.CEFGlobalApp_SetOnBrowserCreated).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnBrowserDestroyed(fn GlobalCEFAppEventOnBrowserDestroyed) {
	imports.Proc(def.CEFGlobalApp_SetOnBrowserDestroyed).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnContextCreated(fn GlobalCEFAppEventOnContextCreated) {
	m.onContextCreated = fn
}

func (m *TCEFApplication) setOnContextCreated(fn GlobalCEFAppEventOnContextCreated) {
	imports.Proc(def.CEFGlobalApp_SetOnContextCreated).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) defaultSetOnContextCreated() {
	m.setOnContextCreated(func(browse *ICefBrowser, frame *ICefFrame, context *ICefV8Context) bool {
		var flag bool
		if m.onContextCreated != nil {
			flag = m.onContextCreated(browse, frame, context)
		}
		if !flag {
			appOnContextCreated(browse, frame, context)
		}
		return false
	})
}

func (m *TCEFApplication) SetOnContextReleased(fn GlobalCEFAppEventOnContextReleased) {
	imports.Proc(def.CEFGlobalApp_SetOnContextReleased).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnUncaughtException(fn GlobalCEFAppEventOnUncaughtException) {
	imports.Proc(def.CEFGlobalApp_SetOnUncaughtException).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnFocusedNodeChanged(fn GlobalCEFAppEventOnFocusedNodeChanged) {
	imports.Proc(def.CEFGlobalApp_SetOnFocusedNodeChanged).Call(api.MakeEventDataPtr(fn))
}

// SetOnProcessMessageReceived
//
//	进程间通信处理消息接收回调函数
func (m *TCEFApplication) SetOnProcessMessageReceived(fn RenderProcessMessageReceived) {
	m.onProcessMessageReceived = fn
}

func (m *TCEFApplication) setOnProcessMessageReceived(fn RenderProcessMessageReceived) {
	imports.Proc(def.CEFGlobalApp_SetOnProcessMessageReceived).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) defaultSetOnProcessMessageReceived() {
	m.setOnProcessMessageReceived(func(browse *ICefBrowser, frame *ICefFrame, sourceProcess CefProcessId, processMessage *ICefProcessMessage) bool {
		var result = renderProcessMessageReceived(browse, frame, sourceProcess, processMessage)
		if m.onProcessMessageReceived != nil && !result {
			result = m.onProcessMessageReceived(browse, frame, sourceProcess, processMessage)
		}
		return result
	})
}

func (m *TCEFApplication) SetOnRenderLoadingStateChange(fn GlobalCEFAppEventOnRenderLoadingStateChange) {
	imports.Proc(def.CEFGlobalApp_SetOnRenderLoadingStateChange).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnRenderLoadStart(fn GlobalCEFAppEventOnRenderLoadStart) {
	m.onRenderLoadStart = fn
}

func (m *TCEFApplication) setOnRenderLoadStart(fn GlobalCEFAppEventOnRenderLoadStart) {
	imports.Proc(def.CEFGlobalApp_SetOnRenderLoadStart).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) defaultSetOnRenderLoadStart() {
	m.setOnRenderLoadStart(func(browser *ICefBrowser, frame *ICefFrame, transitionType TCefTransitionType) {
		if ipcRender != nil {
			ipcRender.clear()
		}
		if m.onRenderLoadStart != nil {
			m.onRenderLoadStart(browser, frame, transitionType)
		}
	})
}

func (m *TCEFApplication) SetOnRenderLoadEnd(fn GlobalCEFAppEventOnRenderLoadEnd) {
	imports.Proc(def.CEFGlobalApp_SetOnRenderLoadEnd).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnRenderLoadError(fn GlobalCEFAppEventOnRenderLoadError) {
	imports.Proc(def.CEFGlobalApp_SetOnRenderLoadError).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnScheduleMessagePumpWork(fn GlobalCEFAppEventOnScheduleMessagePumpWork) {
	var callback uintptr
	if fn != nil {
		callback = api.MakeEventDataPtr(fn)
	}
	imports.Proc(def.CEFGlobalApp_SetOnScheduleMessagePumpWork).Call(callback)
}
