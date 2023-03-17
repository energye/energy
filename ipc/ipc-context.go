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
	"github.com/energye/energy/pkgs/json"
)

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
	ArgumentList() json.JSONArray //参数列表
	BrowserId() int32             //事件所属 browser id
	FrameId() int64               //事件所属 frame id
	Replay() IReplay              //回复, 触发事件返回结果是 IContext 时,该字段为nil
	Result(data ...interface{})   //返回结果
}

// Context IPC 事件上下文
type Context struct {
	browserId int32
	frameId   int64
	argument  json.JSONArray
	replay    IReplay
}

func NewContext(browserId int32, frameId int64, isReplay bool, argument json.JSONArray) IContext {
	ctx := &Context{
		browserId: browserId,
		frameId:   frameId,
		argument:  argument,
	}
	if isReplay {
		ctx.replay = new(Replay)
	}
	return ctx
}

// ArgumentList 参数列表
func (m *Context) ArgumentList() json.JSONArray {
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
