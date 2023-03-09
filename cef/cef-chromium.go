//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF Chromium组件
package cef

import (
	. "github.com/energye/energy/common"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"sync"
	"unsafe"
)

var executeJS = &ExecuteJS{
	msgID:        &MsgID{},
	emitSync:     &ipc.EmitSyncCollection{Mutex: new(sync.Mutex), EmitCollection: sync.Map{}},
	emitCallback: &ipc.EmitCallbackCollection{EmitCollection: sync.Map{}},
}

type ExecuteJS struct {
	msgID        *MsgID                      //ipc 消息ID生成
	emitCallback *ipc.EmitCallbackCollection //回调函数集合
	emitSync     *ipc.EmitSyncCollection     //触发同步事件集合
}

// IChromium 组件接口
type IChromium interface {
	IChromiumProc
	IChromiumEvent
}

// TCEFChromium 组件
type TCEFChromium struct {
	*lcl.TComponent
	instance      unsafe.Pointer
	cfg           *tCefChromiumConfig
	browser       *ICefBrowser
	emitLock      *sync.Mutex
	browserHandle types.HWND
	widgetHandle  types.HWND
	renderHandle  types.HWND
}

// NewChromium 创建一个新的 TCEFChromium
func NewChromium(owner lcl.IComponent, config *tCefChromiumConfig) IChromium {
	m := new(TCEFChromium)
	if config != nil {
		m.cfg = config
	} else {
		m.cfg = NewChromiumConfig()
	}
	m.instance = unsafe.Pointer(_CEFChromium_Create(lcl.CheckPtr(owner), uintptr(unsafe.Pointer(m.cfg))))
	m.emitLock = new(sync.Mutex)
	m.initDefault()
	return m
}

// 默认的初始配置
func (m *TCEFChromium) initDefault() {
	//通过设置这些首选项，可以降低/避免WebRTC的IP泄漏
	m.SetWebRTCIPHandlingPolicy(HpDisableNonProxiedUDP)
	m.SetWebRTCMultipleRoutes(STATE_DISABLED)
	m.SetWebRTCNonproxiedUDP(STATE_DISABLED)
}

// Instance 组件实例指针
func (m *TCEFChromium) Instance() uintptr {
	if m == nil || m.instance == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *TCEFChromium) browseEmitJsOnEvent(browseId int32, frameId int64, name string, argumentList ipc.IArgumentList) ProcessMessageError {
	//data := argumentList.Package()
	//r1 := _CEFFrame_SendProcessMessageByIPC(browseId, frameId, name, PID_RENDER, int32(argumentList.Size()), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
	//return ProcessMessageError(r1)
	return 1 //TODO dev
}

// On 监听事件,事件驱动在Go中监听，JS中触发
func (m *TCEFChromium) On(name string, eventCallback ipc.EventCallback) {
	if eventCallback == nil {
		return
	}
	ipc.IPC.Browser().On(name, eventCallback)
}

// ExecuteJavaScript
// 执行JS代码
//
// code: js代码
//
// scriptURL: js脚本地址 默认about:blank
//
// startLine: js脚本启始执行行号
func (m *TCEFChromium) ExecuteJavaScript(code, scriptURL string, startLine int32) {
	_CEFChromium_ExecuteJavaScript(uintptr(m.instance), code, scriptURL, startLine)
}

// Emit
// 触发JS监听的事件-异步执行
//
// EmitTarget 接收目标, nil:mainBrowser&mainFrame, 可传递browser和指定浏览器窗口，JS监听事件的接收
func (m *TCEFChromium) Emit(eventName string, args ipc.IArgumentList, target ipc.IEmitTarget) ProcessMessageError {
	if eventName == "" {
		return PMErr_NAME_IS_NULL
	}
	m.emitLock.Lock()
	defer m.emitLock.Unlock()
	var (
		browseId int32
		frameId  int64
	)
	if args == nil {
		args = ipc.NewArgumentList()
	}
	if target == nil {
		bsr := m.Browser()
		browseId = bsr.Identifier()
		frameId = bsr.MainFrame().Identifier()
	} else {
		browseId = target.GetBrowserId()
		frameId = target.GetFrameId()
		if m.BrowserById(browseId).GetFrameById(frameId) == nil {
			return PMErr_NOT_FOUND_FRAME
		}
	}
	var idx = args.Size()
	args.SetInt32(idx, int32(Tm_Async))
	args.SetInt32(idx+1, 0)
	args.SetString(idx+2, eventName, true)
	m.browseEmitJsOnEvent(browseId, frameId, ipc.Ln_IPC_GoEmitJS, args)
	return PME_OK
}

// EmitAndCallback
// 触发JS监听的事件-异步执行-带回调
//
// EmitTarget 接收目标, nil = mainBrowser mainFrame
func (m *TCEFChromium) EmitAndCallback(eventName string, args ipc.IArgumentList, target ipc.IEmitTarget, callback ipc.IPCCallback) ProcessMessageError {
	if eventName == "" {
		return PMErr_NAME_IS_NULL
	}
	m.emitLock.Lock()
	defer m.emitLock.Unlock()
	var (
		browseId int32
		frameId  int64
		ipcId    = executeJS.msgID.New()
		idx      = args.Size()
	)
	if args == nil {
		args = ipc.NewArgumentList()
	}
	if target == nil {
		bsr := m.Browser()
		browseId = bsr.Identifier()
		frameId = bsr.MainFrame().Identifier()
	} else {
		browseId = target.GetBrowserId()
		frameId = target.GetFrameId()
		if m.BrowserById(browseId).GetFrameById(frameId) == nil {
			return PMErr_NOT_FOUND_FRAME
		}
	}
	args.SetInt32(idx, int32(Tm_Callback))
	args.SetInt32(idx+1, ipcId)
	args.SetString(idx+2, eventName, true)
	executeJS.emitCallback.EmitCollection.Store(ipcId, callback)
	m.browseEmitJsOnEvent(browseId, frameId, ipc.Ln_IPC_GoEmitJS, args)
	return PME_OK
}

// EmitAndReturn
// 触发JS监听的事件-同步执行-阻塞UI主线程
//
// 使用不当会造成 UI线程 锁死，一搬不在与JS监听中使用，与其它子进程通信时使用
//
// EmitTarget 接收目标, nil = mainBrowser mainFrame
func (m *TCEFChromium) EmitAndReturn(eventName string, args ipc.IArgumentList, target ipc.IEmitTarget) (ipc.IIPCContext, ProcessMessageError) {
	if eventName == "" {
		return nil, PMErr_NAME_IS_NULL
	}
	m.emitLock.Lock()
	defer m.emitLock.Unlock()
	var (
		browseId int32
		frameId  int64
		ipcId    = executeJS.msgID.New()
		idx      = args.Size()
	)
	if args == nil {
		args = ipc.NewArgumentList()
	}
	if target == nil {
		bsr := m.Browser()
		browseId = bsr.Identifier()
		frameId = bsr.MainFrame().Identifier()
	} else {
		browseId = target.GetBrowserId()
		frameId = target.GetFrameId()
		if m.BrowserById(browseId).GetFrameById(frameId) == nil {
			return nil, PMErr_NOT_FOUND_FRAME
		}
	}
	args.SetInt32(idx, int32(Tm_Sync))
	args.SetInt32(idx+1, ipcId)
	args.SetString(idx+2, eventName, true)
	var callback = func(emitAsync *ipc.EmitSyncCollection, ipcId int32) ipc.IIPCContext {
		emitAsync.Mutex.Lock()
		defer emitAsync.Mutex.Unlock()
		var chn = make(chan ipc.IIPCContext)
		var ret ipc.IIPCContext
		emitAsync.EmitCollection.Store(ipcId, chn)
		ret = <-chn //锁住当前线程
		executeJS.emitSync.EmitCollection.Delete(ipcId)
		return ret
	}
	m.browseEmitJsOnEvent(browseId, frameId, ipc.Ln_IPC_GoEmitJS, args)
	return callback(executeJS.emitSync, ipcId), PME_OK
}
