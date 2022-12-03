//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

//CEF应用对象
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
	cfg.framework()
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

//启动主进程
func (m *TCEFApplication) StartMainProcess() bool {
	if m.instance != 0 {
		b := api.GoBool(_CEFStartMainProcess(m.instance))
		if b {
			internalBrowserIPCOnEventInit()
			ipc.IPC.StartBrowserIPC()
			bindGoToJS(nil, nil)
		}
		return b
	}
	return false
}

//启动子进程, 如果指定了子进程执行程序将执行指定的子进程程序
func (m *TCEFApplication) StartSubProcess() bool {
	if m.instance != 0 {
		b := api.GoBool(_CEFStartSubProcess(m.instance))
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
//返回值 false 将会创建 render进程的IPC和GO绑定变量
//
//对于一些不想GO绑定变量的URL地址，实现该函数，通过 frame.Url
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
	_AddCustomCommandLine(api.PascalStr(commandLine), api.PascalStr(value))
}

//启动子进程之前自定义命令行参数设置
func (m *TCEFApplication) SetOnBeforeChildProcessLaunch(fn GlobalCEFAppEventOnBeforeChildProcessLaunch) {
	if common.Args.IsMain() {
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
