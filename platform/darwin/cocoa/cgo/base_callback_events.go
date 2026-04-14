//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cgo

import (
	"sync"
)

type TCallback struct {
	cb func(ctx *TCallbackContext) *GoArguments
}

func MakeNotifyEvent(cb TNotifyEvent) *TCallback {
	return &TCallback{
		cb: func(ctx *TCallbackContext) *GoArguments {
			return cb(ctx.Identifier, ctx.Owner, ctx.Sender)
		},
	}
}

func MakeTextChangeEvent(cb TTextEvent) *TCallback {
	return &TCallback{
		cb: func(ctx *TCallbackContext) *GoArguments {
			return cb(ctx.Identifier, ctx.Value, ctx.Owner, ctx.Sender)
		},
	}
}

func MakeTextCommitEvent(cb TTextEvent) *TCallback {
	return &TCallback{
		cb: func(ctx *TCallbackContext) *GoArguments {
			return cb(ctx.Identifier, ctx.Value, ctx.Owner, ctx.Sender)
		},
	}
}

func MakeDelegateEvent(cb TDelegateEvent) *TCallback {
	return &TCallback{
		cb: func(ctx *TCallbackContext) *GoArguments {
			return cb(ctx.Arguments, ctx.Owner, ctx.Sender)
		},
	}
}

// 事件列表
var (
	eventList = make(map[string]*TCallback)
	eventLock sync.Mutex
)

// RegisterEvent 事件注册，使用控件唯一标识 + 事件类型做为事件唯一id
func RegisterEvent(identifier string, fn *TCallback) {
	eventLock.Lock()
	defer eventLock.Unlock()
	eventList[identifier] = fn
}
