//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

//CEF应用对象
type TCEFApplication struct {
	instance unsafe.Pointer
}

//创建应用程序
func NewApplication(cfg *tCefApplicationConfig) *TCEFApplication {
	if cfg == nil {
		cfg = NewApplicationConfig()
	}
	cfg.framework()
	result := new(TCEFApplication)
	r1, _, _ := Proc(internale_CEFApplication_Create).Call(uintptr(unsafe.Pointer(cfg)))
	result.instance = unsafe.Pointer(r1)
	//register default function
	result.defaultSetOnContextCreated()
	result.defaultSetOnProcessMessageReceived()
	result.defaultSetOnBeforeChildProcessLaunch()
	return result
}

func (m *TCEFApplication) Instance() uintptr {
	return uintptr(m.instance)
}

//启动主进程
func (m *TCEFApplication) StartMainProcess() bool {
	if m.instance != nullptr {
		r1, _, _ := Proc(internale_CEFStartMainProcess).Call(m.Instance())
		var b = api.GoBool(r1)
		if b {
			internalBrowserIPCOnEventInit()
			ipc.IPC.StartBrowserIPC()
			bindGoToJS(nil, nil)
		}
		return b
	}
	return false
}

//启动子进程, 如果指定了子进程执行程序, 将执行指定的子进程程序
func (m *TCEFApplication) StartSubProcess() (result bool) {
	if m.instance != nullptr {
		r1, _, _ := Proc(internale_CEFStartSubProcess).Call(m.Instance())
		result = api.GoBool(r1)
	}
	return false
}

func (m *TCEFApplication) StopScheduler() {
	Proc(internale_CEFApplication_StopScheduler).Call()
}

func (m *TCEFApplication) Destroy() {
	Proc(internale_CEFApplication_Destroy).Call()
}

func (m *TCEFApplication) Free() {
	if m.instance != nullptr {
		Proc(internale_CEFApplication_Free).Call()
		m.instance = nullptr
	}
}

func (m *TCEFApplication) ExecuteJS(browserId int32, code string) {
	Proc(internale_CEFApplication_ExecuteJS).Call()
}

//上下文件创建回调
//
//返回值 false 将会创建 render进程的IPC和GO绑定变量
//
//对于一些不想GO绑定变量的URL地址，实现该函数，通过 frame.Url
func (m *TCEFApplication) SetOnContextCreated(fn GlobalCEFAppEventOnContextCreated) {
	Proc(internale_CEFGlobalApp_SetOnContextCreated).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) defaultSetOnContextCreated() {
	m.SetOnContextCreated(func(browse *ICefBrowser, frame *ICefFrame, context *ICefV8Context) bool {
		return false
	})
}

//初始化设置全局回调
func (m *TCEFApplication) SetOnWebKitInitialized(fn GlobalCEFAppEventOnWebKitInitialized) {
	Proc(internale_CEFGlobalApp_SetOnWebKitInitialized).Call(api.MakeEventDataPtr(fn))
}

//进程间通信处理消息接收
func (m *TCEFApplication) SetOnProcessMessageReceived(fn RenderProcessMessageReceived) {
	Proc(internale_CEFGlobalApp_SetOnProcessMessageReceived).Call(api.MakeEventDataPtr(fn))
}
func (m *TCEFApplication) defaultSetOnProcessMessageReceived() {
	m.SetOnProcessMessageReceived(func(browse *ICefBrowser, frame *ICefFrame, sourceProcess CefProcessId, processMessage *ipc.ICefProcessMessage) bool {
		return false
	})
}

func (m *TCEFApplication) AddCustomCommandLine(commandLine, value string) {
	Proc(internale_AddCustomCommandLine).Call(api.PascalStr(commandLine), api.PascalStr(value))
}

//启动子进程之前自定义命令行参数设置
func (m *TCEFApplication) SetOnBeforeChildProcessLaunch(fn GlobalCEFAppEventOnBeforeChildProcessLaunch) {
	Proc(internale_CEFGlobalApp_SetOnBeforeChildProcessLaunch).Call(api.MakeEventDataPtr(fn))
}
func (m *TCEFApplication) defaultSetOnBeforeChildProcessLaunch() {
	m.SetOnBeforeChildProcessLaunch(func(commandLine *TCefCommandLine) {})
}

func (m *TCEFApplication) SetOnBrowserDestroyed(fn GlobalCEFAppEventOnBrowserDestroyed) {
	Proc(internale_CEFGlobalApp_SetOnBrowserDestroyed).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnLoadStart(fn GlobalCEFAppEventOnRenderLoadStart) {
	Proc(internale_CEFGlobalApp_SetOnLoadStart).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnLoadEnd(fn GlobalCEFAppEventOnRenderLoadEnd) {
	Proc(internale_CEFGlobalApp_SetOnLoadEnd).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnLoadError(fn GlobalCEFAppEventOnRenderLoadError) {
	Proc(internale_CEFGlobalApp_SetOnLoadError).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnLoadingStateChange(fn GlobalCEFAppEventOnRenderLoadingStateChange) {
	Proc(internale_CEFGlobalApp_SetOnLoadingStateChange).Call(api.MakeEventDataPtr(fn))
}
