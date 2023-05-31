//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 基于IPC的字段数据绑定
package cef

import (
	"github.com/energye/energy/v2/cef/internal/ipc"
	"github.com/energye/energy/v2/cef/process"
	"time"
)

const (
	internalV8Bind     = "v8"
	internalGET        = "get"
	internalGETResult  = "get_result"
	internalSET        = "set"
	internalSETResult  = "set_result"
	internalCALL       = "call"
	internalCALLResult = "call_result"
	internalJSON       = "json"
	internalJSONResult = "json_result"
)

var internalBind = "energy"

const (
	internalGetFieldBind       = "field_bind"
	internalGetFieldBindResult = "field_bind_result"
)

// isInternalBind 内部字段不能使用
func isInternalBind(name string) bool {
	return name == internalBind || name == internalV8Bind || name == internalGetFieldBind || name == internalGetFieldBindResult || name == internalGET ||
		name == internalGETResult || name == internalSET || name == internalSETResult || name == internalCALL || name == internalCALLResult || name == internalJSON || name == internalJSONResult
}

// SetBindName 设置自定义绑定根对象名
func SetBindName(name string) {
	internalBind = name
}

// bindInit 初始化
func bindInit() {
	isSingleProcess := application.SingleProcess()
	if isSingleProcess {
		bindBrowser = &bindBrowserProcess{}
		bindRender = &bindRenderProcess{syncChan: &ipc.SyncChan{}}
		// TODO 单进程有些问题
		ipc.CreateBrowserIPC()                         // Go IPC browser
		ipc.CreateRenderIPC(0, time.Now().UnixMicro()) // Go IPC render
	} else {
		if process.Args.IsMain() {
			bindBrowser = &bindBrowserProcess{}
			ipc.CreateBrowserIPC() // Go IPC browser
		} else if process.Args.IsRender() {
			bindRender = &bindRenderProcess{syncChan: &ipc.SyncChan{}}
			ipc.CreateRenderIPC(0, time.Now().UnixMicro()) // Go IPC render
		}
	}
}
