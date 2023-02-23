//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 应用程序
package cef

import (
	"fmt"
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/imports"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"strings"
	"unsafe"
)

var application *TCEFApplication

// CEF应用对象
type TCEFApplication struct {
	instance unsafe.Pointer
	cfg      *tCefApplicationConfig
}

// 创建CEF应用程序
func NewCEFApplication(cfg *tCefApplicationConfig) *TCEFApplication {
	if application != nil {
		return application
	}
	if cfg == nil {
		cfg = NewApplicationConfig()
	}
	cfg.framework()
	application = new(TCEFApplication)
	r1, _, _ := imports.Proc(internale_CEFApplication_Create).Call(uintptr(unsafe.Pointer(cfg)))
	application.instance = unsafe.Pointer(r1)
	application.cfg = cfg

	return application
}

// 创建应用程序
//
// 带有默认的应用事件
func NewApplication(cfg *tCefApplicationConfig) *TCEFApplication {
	cefApp := NewCEFApplication(cfg)
	cefApp.registerDefaultEvent()
	return cefApp
}

func (m *TCEFApplication) registerDefaultEvent() {
	//register default function
	m.defaultSetOnContextCreated()
	m.defaultSetOnProcessMessageReceived()
	m.defaultSetOnBeforeChildProcessLaunch()
}

func (m *TCEFApplication) Instance() uintptr {
	return uintptr(m.instance)
}

// 启动主进程
func (m *TCEFApplication) StartMainProcess() bool {
	if m.instance != nullptr {
		logger.Debug("application single exe,", common.Args.ProcessType(), "process start")
		r1, _, _ := imports.Proc(internale_CEFStartMainProcess).Call(m.Instance())
		return api.GoBool(r1)
	}
	return false
}

// 启动子进程, 如果指定了子进程执行程序, 将执行指定的子进程程序
func (m *TCEFApplication) StartSubProcess() (result bool) {
	if m.instance != nullptr {
		logger.Debug("application multiple exe,", common.Args.ProcessType(), "process start")
		r1, _, _ := imports.Proc(internale_CEFStartSubProcess).Call(m.Instance())
		result = api.GoBool(r1)
	}
	return false
}

func (m *TCEFApplication) RunMessageLoop() {
	defer func() {
		logger.Debug("application run message loop end")
		api.EnergyLibRelease()
	}()
	logger.Debug("application run message loop start")
	imports.Proc(internale_CEFApplication_RunMessageLoop).Call()
}

func (m *TCEFApplication) QuitMessageLoop() {
	logger.Debug("application quit message loop")
	imports.Proc(internale_CEFApplication_QuitMessageLoop).Call()
}

func (m *TCEFApplication) StopScheduler() {
	imports.Proc(internale_CEFApplication_StopScheduler).Call()
}

func (m *TCEFApplication) Destroy() {
	imports.Proc(internale_CEFApplication_Destroy).Call()
}

func (m *TCEFApplication) Free() {
	if m.instance != nullptr {
		imports.Proc(internale_CEFApplication_Free).Call()
		m.instance = nullptr
	}
}

func (m *TCEFApplication) ExecuteJS(browserId int32, code string) {
	imports.Proc(internale_CEFApplication_ExecuteJS).Call()
}

// 上下文件创建回调
//
// 返回值 false 将会创建 render进程的IPC和GO绑定变量
//
// 对于一些不想GO绑定变量的URL地址，实现该函数，通过 frame.Url
func (m *TCEFApplication) SetOnContextCreated(fn GlobalCEFAppEventOnContextCreated) {
	imports.Proc(internale_CEFGlobalApp_SetOnContextCreated).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) defaultSetOnContextCreated() {
	m.SetOnContextCreated(func(browse *ICefBrowser, frame *ICefFrame, context *ICefV8Context) bool {
		return false
	})
}

func (m *TCEFApplication) SetOnContextInitialized(fn GlobalCEFAppEventOnContextInitialized) {
	imports.Proc(internale_CEFGlobalApp_SetOnContextInitialized).Call(api.MakeEventDataPtr(fn))
}

// 初始化设置全局回调
func (m *TCEFApplication) SetOnWebKitInitialized(fn GlobalCEFAppEventOnWebKitInitialized) {
	imports.Proc(internale_CEFGlobalApp_SetOnWebKitInitialized).Call(api.MakeEventDataPtr(fn))
}

// 进程间通信处理消息接收
func (m *TCEFApplication) SetOnProcessMessageReceived(fn RenderProcessMessageReceived) {
	imports.Proc(internale_CEFGlobalApp_SetOnProcessMessageReceived).Call(api.MakeEventDataPtr(fn))
}
func (m *TCEFApplication) defaultSetOnProcessMessageReceived() {
	m.SetOnProcessMessageReceived(func(browse *ICefBrowser, frame *ICefFrame, sourceProcess CefProcessId, processMessage *ipc.ICefProcessMessage) bool {
		return false
	})
}

func (m *TCEFApplication) AddCustomCommandLine(commandLine, value string) {
	imports.Proc(internale_AddCustomCommandLine).Call(api.PascalStr(commandLine), api.PascalStr(value))
}

// 启动子进程之前自定义命令行参数设置
func (m *TCEFApplication) SetOnBeforeChildProcessLaunch(fn GlobalCEFAppEventOnBeforeChildProcessLaunch) {
	imports.Proc(internale_CEFGlobalApp_SetOnBeforeChildProcessLaunch).Call(api.MakeEventDataPtr(fn))
}
func (m *TCEFApplication) defaultSetOnBeforeChildProcessLaunch() {
	m.SetOnBeforeChildProcessLaunch(func(commandLine *TCefCommandLine) {})
}

func (m *TCEFApplication) SetOnBrowserDestroyed(fn GlobalCEFAppEventOnBrowserDestroyed) {
	imports.Proc(internale_CEFGlobalApp_SetOnBrowserDestroyed).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnLoadStart(fn GlobalCEFAppEventOnRenderLoadStart) {
	imports.Proc(internale_CEFGlobalApp_SetOnLoadStart).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnLoadEnd(fn GlobalCEFAppEventOnRenderLoadEnd) {
	imports.Proc(internale_CEFGlobalApp_SetOnLoadEnd).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnLoadError(fn GlobalCEFAppEventOnRenderLoadError) {
	imports.Proc(internale_CEFGlobalApp_SetOnLoadError).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnLoadingStateChange(fn GlobalCEFAppEventOnRenderLoadingStateChange) {
	imports.Proc(internale_CEFGlobalApp_SetOnLoadingStateChange).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnGetDefaultClient(fn GlobalCEFAppEventOnGetDefaultClient) {
	imports.Proc(internale_CEFGlobalApp_SetOnGetDefaultClient).Call(api.MakeEventDataPtr(fn))
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
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
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      common.StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			fn.(GlobalCEFAppEventOnRenderLoadStart)(browser, frame, TCefTransitionType(getVal(2)))
		case GlobalCEFAppEventOnRenderLoadEnd:
			browser := &ICefBrowser{browseId: int32(getVal(0))}
			tempFrame := (*cefFrame)(getPtr(1))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      common.StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			fn.(GlobalCEFAppEventOnRenderLoadEnd)(browser, frame, int32(getVal(2)))
		case GlobalCEFAppEventOnRenderLoadError:
			browser := &ICefBrowser{browseId: int32(getVal(0))}
			tempFrame := (*cefFrame)(getPtr(1))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      common.StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			fn.(GlobalCEFAppEventOnRenderLoadError)(browser, frame, TCefErrorCode(getVal(2)), api.GoStr(getVal(3)), api.GoStr(getVal(4)))
		case GlobalCEFAppEventOnRenderLoadingStateChange:
			browser := &ICefBrowser{browseId: int32(getVal(0))}
			tempFrame := (*cefFrame)(getPtr(1))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      common.StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			fn.(GlobalCEFAppEventOnRenderLoadingStateChange)(browser, frame, api.GoBool(getVal(2)), api.GoBool(getVal(3)), api.GoBool(getVal(4)))
		case RenderProcessMessageReceived:
			browser := &ICefBrowser{browseId: int32(getVal(0))}
			tempFrame := (*cefFrame)(getPtr(1))
			frame := &ICefFrame{
				Browser: browser,
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      common.StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			cefProcMsg := (*ipc.CefProcessMessagePtr)(getPtr(3))
			args := ipc.NewArgumentList()
			args.UnPackageBytePtr(cefProcMsg.Data, int32(cefProcMsg.DataLen))
			processMessage := &ipc.ICefProcessMessage{
				Name:         api.GoStr(cefProcMsg.Name),
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
				Name:    api.GoStr(tempFrame.Name),
				Url:     api.GoStr(tempFrame.Url),
				Id:      common.StrToInt64(api.GoStr(tempFrame.Identifier)),
			}
			if strings.Index(frame.Url, "devtools://") == 0 {
				processName = common.PT_DEVTOOLS
				return true
			} else {
				processName = common.Args.ProcessType()
			}
			v8ctx := (*iCefV8ContextPtr)(getPtr(2))
			ctx := &ICefV8Context{
				instance: unsafe.Pointer(v8ctx.V8Context),
				Browser:  browser,
				Frame:    frame,
				Global:   &ICefV8Value{instance: unsafe.Pointer(v8ctx.Global)},
			}
			fmt.Println("iCefV8ContextPtr", v8ctx, "Global.IsValid:", ctx.Global.IsValid(), ctx.Global.IsUndefined(), ctx.Global.GetDateValue())
			fmt.Println("iCefV8ContextPtr GetStringValuer", ctx.Global.GetStringValue())
			fmt.Println("iCefV8ContextPtr GetValueByIndex", ctx.Global.GetValueByIndex(0).IsValid())
			fmt.Println("iCefV8ContextPtr GetValueByIndex", ctx.Global.GetValueByIndex(1).IsValid())
			fmt.Println("iCefV8ContextPtr GetValueByKey", ctx.Global.GetValueByKey("name").IsValid())
			fmt.Println("iCefV8ContextPtr SetValueByAccessor", ctx.Global.SetValueByAccessor("nametest", V8_ACCESS_CONTROL_DEFAULT, V8_PROPERTY_ATTRIBUTE_NONE))
			fmt.Println("iCefV8ContextPtr GetExternallyAllocatedMemory", ctx.Global.GetExternallyAllocatedMemory())
			fmt.Println("iCefV8ContextPtr AdjustExternallyAllocatedMemory", ctx.Global.AdjustExternallyAllocatedMemory(0))
			fmt.Println("iCefV8ContextPtr GetExternallyAllocatedMemory.GetValueByKey", ctx.Global.GetValueByKey("name").GetExternallyAllocatedMemory())
			fmt.Println("iCefV8ContextPtr GetFunctionName", ctx.Global.GetFunctionName())
			fmt.Println("V8ValueRef IsValid", V8ValueRef.NewUndefined().IsValid())
			handler := CreateCefV8Handler()
			accessor := CreateCefV8Accessor()
			fmt.Println("handler-accessor:", handler, accessor)
			accessor.Get(func(name string, object *ICefV8Value, retVal *ResultV8Value, exception *Exception) bool {
				retVal.SetResult(V8ValueRef.NewString("这能返回？"))
				return true
			})
			accessor.Set(func(name string, object *ICefV8Value, value *ICefV8Value, exception *Exception) bool {
				fmt.Println("name", name, "object.IsValid", object.IsValid(), object.IsObject(), object.IsString(), "value.IsValid", value.IsValid(), value.IsString(), value.IsObject())
				return true
			})
			handler.Execute(func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *Exception) bool {
				fmt.Println(arguments.Size())
				fmt.Println(arguments.Get(0).IsValid(), arguments.Get(0).GetStringValue())
				fmt.Println(arguments.Get(1).IsValid(), arguments.Get(1).GetIntValue())
				retVal.SetResult(V8ValueRef.NewString("函数返回值？"))
				return true
			})
			object := V8ValueRef.NewObject(accessor, nil)
			fmt.Println("V8ValueRef NewObject", object, object.IsValid())
			object.SetValueByAccessor("testcca", V8_ACCESS_CONTROL_DEFAULT, V8_PROPERTY_ATTRIBUTE_NONE)
			object.SetValueByKey("testcca", V8ValueRef.NewObject(accessor, nil), V8_PROPERTY_ATTRIBUTE_NONE)
			object.SetValueByKey("testcca", V8ValueRef.NewObject(accessor, nil), V8_PROPERTY_ATTRIBUTE_NONE)
			object.SetValueByKey("testfn", V8ValueRef.NewFunction("testfn", handler), V8_PROPERTY_ATTRIBUTE_NONE)
			fmt.Println("Global.SetValueByKey", ctx.Global.SetValueByKey("testset", object, V8_PROPERTY_ATTRIBUTE_READONLY))

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
		case GlobalCEFAppEventOnContextInitialized:
			fn.(GlobalCEFAppEventOnContextInitialized)()
		case GlobalCEFAppEventOnBeforeChildProcessLaunch:
			commands := (*uintptr)(getPtr(0))
			commandLine := &TCefCommandLine{commandLines: make(map[string]string)}
			ipc.IPC.SetPort()
			commandLine.AppendSwitch(MAINARGS_NETIPCPORT, fmt.Sprintf("%d", ipc.IPC.Port()))
			fn.(GlobalCEFAppEventOnBeforeChildProcessLaunch)(commandLine)
			*commands = api.PascalStr(commandLine.toString())
		case GlobalCEFAppEventOnGetDefaultClient:
			client := (*uintptr)(getPtr(0))
			getClient := &ICefClient{instance: unsafe.Pointer(client)}
			fn.(GlobalCEFAppEventOnGetDefaultClient)(getClient)
			*client = uintptr(getClient.instance)
		default:
			return false
		}
		return true
	})
}
