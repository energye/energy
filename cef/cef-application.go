//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/commons"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

type TCEFApplication struct {
	instance uintptr
	ptr      unsafe.Pointer
	procName string
}

//创建应用程序
func NewApplication(cfg *tCefApplicationConfig) *TCEFApplication {
	if cfg == nil {
		cfg = NewApplicationConfig()
	}
	m := new(TCEFApplication)
	m.procName = "CEFApplication"
	m.instance = _CEFApplication_Create(uintptr(unsafe.Pointer(cfg)))
	m.ptr = unsafe.Pointer(m.instance)
	//注册默认的函数
	m.defaultSetOnContextCreated()
	m.defaultSetOnProcessMessageReceived()
	m.defaultSetOnBeforeChildProcessLaunch()
	return m
}

func (m *TCEFApplication) StartMainProcess() bool {
	if m.instance != 0 {
		b := api.DBoolToGoBool(_CEFStartMainProcess(m.instance))
		if b {
			internalBrowserIPCOnEventInit()
			ipc.IPC.StartBrowserIPC()
			bindGoToJS(nil, nil)
		}
		return b
	}
	return false
}

func (m *TCEFApplication) StartSubProcess() bool {
	if m.instance != 0 {
		b := api.DBoolToGoBool(_CEFStartSubProcess(m.instance))
		return b
	}
	return false
}

func (m *TCEFApplication) Free() {
	if m.instance != 0 {
		_CEFApplication_Free()
		m.instance, m.ptr = 0, nullptr
	}
}

//上下文件创建回调
//
//返回false 创建 render IPC 和 变量绑定
//
//对于一些不想绑定的URL地址，实现该函数，通过 frame.Url
func (m *TCEFApplication) SetOnContextCreated(fn GlobalCEFAppEventOnContextCreated) {
	_SetCEFCallbackEvent(OnContextCreated, fn)
}
func (m *TCEFApplication) defaultSetOnContextCreated() {
	m.SetOnContextCreated(func(browse *ICefBrowser, frame *ICefFrame, context *ICefV8Context) bool {
		return false
	})
}

//初始化设置全局回调
func (m *TCEFApplication) SetOnWebKitInitialized(fn GlobalCEFAppEventOnWebKitInitialized) {
	_SetCEFCallbackEvent(OnWebKitInitialized, fn)
}

//进程间通信处理消息接收
func (m *TCEFApplication) SetOnProcessMessageReceived(fn RenderProcessMessageReceived) {
	_SetCEFCallbackEvent(OnProcessMessageReceived, fn)
}
func (m *TCEFApplication) defaultSetOnProcessMessageReceived() {
	m.SetOnProcessMessageReceived(func(browse *ICefBrowser, frame *ICefFrame, sourceProcess CefProcessId, processMessage *ipc.ICefProcessMessage) bool {
		return false
	})
}

func (m *TCEFApplication) AddCustomCommandLine(commandLine, value string) {
	_AddCustomCommandLine(api.GoStrToDStr(commandLine), api.GoStrToDStr(value))
}

//启动子进程之前自定义命令行参数设置
func (m *TCEFApplication) SetOnBeforeChildProcessLaunch(fn GlobalCEFAppEventOnBeforeChildProcessLaunch) {
	if commons.Args.IsMain() {
		_SetCEFCallbackEvent(OnBeforeChildProcessLaunch, fn)
	}
}
func (m *TCEFApplication) defaultSetOnBeforeChildProcessLaunch() {
	m.SetOnBeforeChildProcessLaunch(func(commandLine *TCefCommandLine) {})
}

func (m *TCEFApplication) SetOnBrowserDestroyed(fn GlobalCEFAppEventOnBrowserDestroyed) {
	_SetCEFCallbackEvent(OnBrowserDestroyed, fn)
}

func (m *TCEFApplication) SetOnLoadStart(fn GlobalCEFAppEventOnRenderLoadStart) {
	_SetCEFCallbackEvent(OnRenderLoadStart, fn)
}

func (m *TCEFApplication) SetOnLoadEnd(fn GlobalCEFAppEventOnRenderLoadEnd) {
	_SetCEFCallbackEvent(OnRenderLoadEnd, fn)
}

func (m *TCEFApplication) SetOnLoadError(fn GlobalCEFAppEventOnRenderLoadError) {
	_SetCEFCallbackEvent(OnRenderLoadError, fn)
}

func (m *TCEFApplication) SetOnLoadingStateChange(fn GlobalCEFAppEventOnRenderLoadingStateChange) {
	_SetCEFCallbackEvent(OnRenderLoadingStateChange, fn)
}
