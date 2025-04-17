//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package context

import (
	"github.com/cyber-xxm/energy/v2/pkgs/json"
)

// IReplay IPC Replay Interface
type IReplay interface {
	Result() []interface{}
	SetResult(data []interface{})
	Clear()
}

// Replay IPC Replay
type Replay struct {
	data []interface{}
}

// IContext
//
//	Inter process IPC communication callback context
type IContext interface {
	ArgumentList() json.JSONArray //ArgumentList
	BrowserId() int32             //Event ownership: browser id
	FrameId() string              //Event ownership: frame id
	Replay() IReplay              //Replay, When the trigger event returns IContext, this field is nil
	Result(data ...interface{})   //callback function return Result
}

// Context IPC Event context
type Context struct {
	browserId int32
	frameId   string
	argument  json.JSONArray
	replay    IReplay
}

// NewContext create IPC message Replay Context
func NewContext(browserId int32, frameId string, isReplay bool, argument json.JSONArray) IContext {
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

// ArgumentList Argument List JSONArray
func (m *Context) ArgumentList() json.JSONArray {
	return m.argument
}

func (m *Context) BrowserId() int32 {
	return m.browserId
}

func (m *Context) FrameId() string {
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
