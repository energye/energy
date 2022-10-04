//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/golcl/lcl"
	"sync"
	"unsafe"
)

var executeJS = &ExecuteJS{
	msgID:        &msgID{},
	emitSync:     &emitSyncCollection{mutex: new(sync.Mutex), emitCollection: sync.Map{}},
	emitCallback: &emitCallbackCollection{emitCollection: sync.Map{}},
}

type ExecuteJS struct {
	msgID        *msgID                  //ipc 消息ID生成
	emitCallback *emitCallbackCollection //回调函数集合
	emitSync     *emitSyncCollection     //触发同步事件集合
}

type IChromium interface {
	IChromiumProc
	IChromiumEvent
	//启用独立事件 默认禁用, 启用后所有默认事件行为将不在主窗口chromium event执行
	//
	//启用后注册的事件才生效
	//
	//只对当前chromium对象有效
	EnableIndependentEvent()

	//禁用独立事件 默认禁用, 禁用后所有默认事件行为在主窗口chromium event执行
	//
	//禁用后注册的事件才生效
	//
	//只对当前chromium对象有效
	DisableIndependentEvent()
}

type TCEFChromium struct {
	BaseComponent
	cfg              *tCefChromiumConfig
	independentEvent bool
	emitLock         *sync.Mutex
}

func NewChromium(owner lcl.IComponent, config *tCefChromiumConfig) *TCEFChromium {
	m := new(TCEFChromium)
	m.procName = "CEFChromium"
	if config != nil {
		m.cfg = config
	} else {
		m.cfg = NewChromiumConfig()
	}
	m.instance = _CEFChromium_Create(lcl.CheckPtr(owner), uintptr(unsafe.Pointer(m.cfg)))
	m.ptr = unsafe.Pointer(m.instance)
	m.emitLock = new(sync.Mutex)
	return m
}

func (m *TCEFChromium) GetBrowserById(browserId int32) *ICefBrowser {
	return &ICefBrowser{
		browseId: browserId,
		chromium: m.instance,
	}
}

//启用独立事件 默认 false
func (m *TCEFChromium) EnableIndependentEvent() {
	m.independentEvent = true
}

//禁用独立事件 默认 false
func (m *TCEFChromium) DisableIndependentEvent() {
	m.independentEvent = false
}

func (m *TCEFChromium) browseEmitJsOnEvent(browseId int32, frameId int64, name string, argumentList IArgumentList) ProcessMessageError {
	data := argumentList.Package()
	r1 := _CEFFrame_SendProcessMessage(browseId, frameId, name, PID_RENDERER, int32(argumentList.Size()), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
	return ProcessMessageError(r1)
}

func (m *TCEFChromium) On(name string, eventCallback EventCallback) {
	if eventCallback == nil {
		return
	}
	IPC.browser.On(name, eventCallback)
}

// 执行JS代码
//
// code: js代码
//
// scriptURL: js脚本地址默认about:blank
//
// startLine: js脚本启始执行行号
func (m *TCEFChromium) ExecuteJavaScript(code, scriptURL string, startLine int32) {
	_CEFChromium_ExecuteJavaScript(m.instance, code, scriptURL, startLine)
}

// 触发JS监听的事件-异步执行
//
// GoEmitTarget 接收目标, nil = mainBrowser mainFrame
func (m *TCEFChromium) Emit(eventName string, args IArgumentList, target *GoEmitTarget) ProcessMessageError {
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
		args = NewArgumentList()
	}
	if target == nil {
		browser := m.Browser()
		browseId = browser.Identifier()
		frameId = browser.MainFrame().Id
	} else {
		browseId = target.BrowseId
		frameId = target.FrameId
		if m.GetBrowserById(browseId).GetFrameById(frameId) == nil {
			return PMErr_NOT_FOUND_FRAME
		}
	}
	var idx = args.Size()
	args.SetInt32(idx, int32(tm_Async))
	args.SetInt32(idx+1, 0)
	args.SetString(idx+2, eventName, true)
	m.browseEmitJsOnEvent(browseId, frameId, ln_IPC_GoEmitJS, args)
	return PME_OK
}

// 触发JS监听的事件-异步执行-带回调
//
// GoEmitTarget 接收目标, nil = mainBrowser mainFrame
func (m *TCEFChromium) EmitAndCallback(eventName string, args IArgumentList, target *GoEmitTarget, callback ipcCallback) ProcessMessageError {
	if eventName == "" {
		return PMErr_NAME_IS_NULL
	}
	m.emitLock.Lock()
	defer m.emitLock.Unlock()
	var (
		browseId int32
		frameId  int64
		ipcId    = executeJS.msgID.new()
		idx      = args.Size()
	)
	if args == nil {
		args = NewArgumentList()
	}
	if target == nil {
		browser := m.Browser()
		browseId = browser.Identifier()
		frameId = browser.MainFrame().Id
	} else {
		browseId = target.BrowseId
		frameId = target.FrameId
		if m.GetBrowserById(browseId).GetFrameById(frameId) == nil {
			return PMErr_NOT_FOUND_FRAME
		}
	}
	args.SetInt32(idx, int32(tm_Callback))
	args.SetInt32(idx+1, ipcId)
	args.SetString(idx+2, eventName, true)
	executeJS.emitCallback.emitCollection.Store(ipcId, callback)
	m.browseEmitJsOnEvent(browseId, frameId, ln_IPC_GoEmitJS, args)
	return PME_OK
}

// 触发JS监听的事件-同步执行-阻塞UI主线程
//
// 使用不当会造成 UI线程 锁死
//
// GoEmitTarget 接收目标, nil = mainBrowser mainFrame
func (m *TCEFChromium) EmitAndReturn(eventName string, args IArgumentList, target *GoEmitTarget) (IIPCContext, ProcessMessageError) {
	if eventName == "" {
		return nil, PMErr_NAME_IS_NULL
	}
	m.emitLock.Lock()
	defer m.emitLock.Unlock()
	var (
		browseId int32
		frameId  int64
		ipcId    = executeJS.msgID.new()
		idx      = args.Size()
	)
	if args == nil {
		args = NewArgumentList()
	}
	if target == nil {
		browser := m.Browser()
		browseId = browser.Identifier()
		frameId = browser.MainFrame().Id
	} else {
		browseId = target.BrowseId
		frameId = target.FrameId
		if m.GetBrowserById(browseId).GetFrameById(frameId) == nil {
			return nil, PMErr_NOT_FOUND_FRAME
		}
	}
	args.SetInt32(idx, int32(tm_Sync))
	args.SetInt32(idx+1, ipcId)
	args.SetString(idx+2, eventName, true)
	var callback = func(emitAsync *emitSyncCollection, ipcId int32) IIPCContext {
		emitAsync.mutex.Lock()
		defer emitAsync.mutex.Unlock()
		var chn = make(chan IIPCContext)
		var ret IIPCContext
		emitAsync.emitCollection.Store(ipcId, chn)
		ret = <-chn //锁住当前线程
		executeJS.emitSync.emitCollection.Delete(ipcId)
		return ret
	}
	m.browseEmitJsOnEvent(browseId, frameId, ln_IPC_GoEmitJS, args)
	return callback(executeJS.emitSync, ipcId), PME_OK
}
