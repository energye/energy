//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package ipc

import (
	"github.com/energye/energy/common"
	"sync"
)

var (
	browser *browserIPC
)

// browserIPC 主进程 IPC
type browserIPC struct {
	onEvent map[string]EmitCallback
	lock    sync.Mutex
}

type EmitCallback func(context IContext)

func init() {
	if common.Args.IsMain() {
		browser = &browserIPC{onEvent: make(map[string]EmitCallback)}
	}
}

type IReplay interface {
	Result() []interface{}
	SetResult(data []interface{})
	Clear()
}

type Replay struct {
	data []interface{}
}

// IContext 进程间IPC通信回调上下文
type IContext interface {
	ArgumentList() IArrayValue  //参数列表
	BrowserId() int32           //事件所属 browser id
	FrameId() int64             //事件所属 frame id
	Replay() IReplay            //回复
	Result(data ...interface{}) //返回结果
}

// Context IPC 事件上下文
type Context struct {
	browserId int32
	frameId   int64
	argument  IArrayValue
	replay    IReplay
}

func NewContext(browserId int32, frameId int64, argument IArrayValue, isCallback bool) IContext {
	ctx := &Context{
		browserId: browserId,
		frameId:   frameId,
		argument:  argument,
	}
	if isCallback {
		ctx.replay = &Replay{}
	}
	return ctx
}

//On
// IPC GO 监听事件
func On(name string, fn EmitCallback) {
	if browser != nil {
		browser.addOnEvent(name, fn)
	}
}

//RemoveOn
// IPC GO 移除监听事件
func RemoveOn(name string) {
	if browser == nil || name == "" {
		return
	}
	browser.lock.Lock()
	defer browser.lock.Unlock()
	delete(browser.onEvent, name)
}

//Emit
// IPC GO 中触发 JS 监听的事件
func Emit(name string) {

}

//CheckOnEvent
// IPC 检查 GO 中监听的事件是存在, 并返回回调函数
func CheckOnEvent(name string) EmitCallback {
	if browser == nil || name == "" {
		return nil
	}
	browser.lock.Lock()
	defer browser.lock.Unlock()
	if callback, ok := browser.onEvent[name]; ok {
		return callback
	}
	return nil
}

// addOnEvent 添加监听事件
func (m *browserIPC) addOnEvent(name string, fn EmitCallback) {
	if m == nil || name == "" || fn == nil {
		return
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	m.onEvent[name] = fn
}

// emitOnEvent 触发监听事件
func (m *browserIPC) emitOnEvent(name string, argumentList IArrayValue) {
	if m == nil || name == "" || argumentList == nil {
		return
	}
	m.lock.Lock()
	defer m.lock.Unlock()

}

// ArgumentList 参数列表
func (m *Context) ArgumentList() IArrayValue {
	return m.argument
}

func (m *Context) BrowserId() int32 {
	return m.browserId
}

func (m *Context) FrameId() int64 {
	return m.frameId
}

func (m *Context) Replay() IReplay {
	return m.replay
}

func (m *Context) Result(data ...interface{}) {
	if m.replay != nil {
		m.replay.SetResult(data)
	}
}

func (m *Replay) Result() []interface{} {
	if m == nil {
		return nil
	}
	return m.data
}

func (m *Replay) SetResult(data []interface{}) {
	if m != nil {
		m.data = data
	}
}

func (m *Replay) Clear() {
	if m == nil {
		return
	}
	m.data = nil
}
