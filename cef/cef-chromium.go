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

type IChromium interface {
	IChromiumProc
	IChromiumEvent
}

type TCEFChromium struct {
	*lcl.TComponent
	instance      unsafe.Pointer
	cfg           *tCefChromiumConfig
	emitLock      *sync.Mutex
	browserHandle types.HWND
	widgetHandle  types.HWND
	renderHandle  types.HWND
}

func NewChromium(owner lcl.IComponent, config *tCefChromiumConfig) IChromium {
	m := new(TCEFChromium)
	if config != nil {
		m.cfg = config
	} else {
		m.cfg = NewChromiumConfig()
	}
	m.instance = unsafe.Pointer(_CEFChromium_Create(lcl.CheckPtr(owner), uintptr(unsafe.Pointer(m.cfg))))
	m.emitLock = new(sync.Mutex)
	return m
}

func (m *TCEFChromium) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *TCEFChromium) GetBrowserById(browserId int32) *ICefBrowser {
	return &ICefBrowser{
		browseId: browserId,
		chromium: m.instance,
	}
}

func (m *TCEFChromium) browseEmitJsOnEvent(browseId int32, frameId int64, name string, argumentList ipc.IArgumentList) ProcessMessageError {
	data := argumentList.Package()
	r1 := _CEFFrame_SendProcessMessage(browseId, frameId, name, PID_RENDER, int32(argumentList.Size()), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
	return ProcessMessageError(r1)
}

func (m *TCEFChromium) On(name string, eventCallback ipc.EventCallback) {
	if eventCallback == nil {
		return
	}
	ipc.IPC.Browser().On(name, eventCallback)
}

// 执行JS代码
//
// code: js代码
//
// scriptURL: js脚本地址默认about:blank
//
// startLine: js脚本启始执行行号
func (m *TCEFChromium) ExecuteJavaScript(code, scriptURL string, startLine int32) {
	_CEFChromium_ExecuteJavaScript(uintptr(m.instance), code, scriptURL, startLine)
}

// 触发JS监听的事件-异步执行
//
// EmitTarget 接收目标, nil = mainBrowser mainFrame
func (m *TCEFChromium) Emit(eventName string, args ipc.IArgumentList, target IEmitTarget) ProcessMessageError {
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
		frameId = bsr.MainFrame().Id
	} else {
		browseId = target.GetBrowserId()
		frameId = target.GetFrameId()
		if m.GetBrowserById(browseId).GetFrameById(frameId) == nil {
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

// 触发JS监听的事件-异步执行-带回调
//
// EmitTarget 接收目标, nil = mainBrowser mainFrame
func (m *TCEFChromium) EmitAndCallback(eventName string, args ipc.IArgumentList, target IEmitTarget, callback ipc.IPCCallback) ProcessMessageError {
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
		frameId = bsr.MainFrame().Id
	} else {
		browseId = target.GetBrowserId()
		frameId = target.GetFrameId()
		if m.GetBrowserById(browseId).GetFrameById(frameId) == nil {
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

// 触发JS监听的事件-同步执行-阻塞UI主线程
//
// 使用不当会造成 UI线程 锁死
//
// EmitTarget 接收目标, nil = mainBrowser mainFrame
func (m *TCEFChromium) EmitAndReturn(eventName string, args ipc.IArgumentList, target IEmitTarget) (ipc.IIPCContext, ProcessMessageError) {
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
		frameId = bsr.MainFrame().Id
	} else {
		browseId = target.GetBrowserId()
		frameId = target.GetFrameId()
		if m.GetBrowserById(browseId).GetFrameById(frameId) == nil {
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
