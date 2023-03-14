//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// energy ipc 事件驱动监听
//
// 事件监听、事件触发
package ipc

import (
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"reflect"
	"sync"
)

var (
	browser *browserIPC
)

// emitContextCallback IPC 上下文件回调函数
type emitContextCallback func(context IContext)

// emitArgumentCallback 带有参数回调函数
type emitArgumentCallback any

// browserIPC 主进程 IPC
type browserIPC struct {
	onEvent               map[string]*callback
	emitCallback          map[int32]*reflect.Value
	emitCallbackMessageId int32
	onLock                sync.Mutex
	emitLock              sync.Mutex
	processMessage        IProcessMessage
}

func init() {
	if common.Args.IsMain() {
		browser = &browserIPC{onEvent: make(map[string]*callback), emitCallback: make(map[int32]*reflect.Value)}
	}
}

func SetProcessMessage(pm IProcessMessage) {
	if browser == nil || pm == nil {
		return
	}
	browser.processMessage = pm
}

//On
// IPC GO 监听事件, 上下文参数
func On(name string, fn emitContextCallback) {
	browser.addOnEvent(name, &callback{context: &contextCallback{callback: fn}})
}

//OnArguments
// IPC GO 监听事件, 自定义参数，仅支持对应 JavaScript 对应 Go 的常用类型
//
// 入参 不限制个数
//	 基本类型: int(int8 ~ uint64), bool, float(float32、float64), string
//
//   复杂类型: slice, map, struct
//
//   复杂类型限制示例: slice: []interface{}, map: map[string]interface{}
//
//   slice: 只能是 any 或 interface{}
//   map: key 只能 string 类型, value 基本类型+复杂类型
//   struct: 首字母大写, 字段类型匹配
//     type ArgsStructDemo struct {
//        Key1 string
//		  Key2 string
//		  Key3 string
//		  Key4 int
//		  Key5 float64
//		  Key6 bool
//		  Sub1  SubStructXXX
//		  Sub2  *SubStructXXX
//     }
//
// 出参
//
//   与同入参相同
func OnArguments(name string, fn emitArgumentCallback) {
	if browser != nil {
		v := reflect.ValueOf(fn)
		// f must be a function
		if v.Kind() != reflect.Func {
			return
		}
		browser.addOnEvent(name, &callback{argument: &argumentCallback{callback: &v}})
	}
}

//RemoveOn
// IPC GO 移除监听事件
func RemoveOn(name string) {
	if browser == nil || name == "" {
		return
	}
	browser.onLock.Lock()
	defer browser.onLock.Unlock()
	delete(browser.onEvent, name)
}

//Emit
// IPC GO 中触发 JS 监听的事件
func Emit(name string, argument ...any) {
	if browser == nil || name == "" || browser.processMessage == nil {
		return
	}
	browser.processMessage.SendProcessMessageForIPC(0, name, consts.PID_RENDER, nil, argument...)
}

//EmitAndCallback
// IPC GO 中触发 JS 监听的事件
func EmitAndCallback(name string, callback any, argument ...any) {
	if browser == nil || name == "" || browser.processMessage == nil {
		return
	}
	messageId := browser.addEmitCallback(callback)
	browser.processMessage.SendProcessMessageForIPC(messageId, name, consts.PID_RENDER, nil, argument...)
}

//EmitTarget
// IPC GO 中触发指定目标 JS 监听的事件
func EmitTarget(name string, target ITarget, argument ...any) {
	if browser == nil || name == "" || browser.processMessage == nil {
		return
	}
	browser.processMessage.SendProcessMessageForIPC(0, name, consts.PID_RENDER, target, argument...)
}

//EmitTargetAndCallback
// IPC GO 中触发指定目标 JS 监听的事件
func EmitTargetAndCallback(name string, target ITarget, callback any, argument ...any) {
	if browser == nil || name == "" || browser.processMessage == nil {
		return
	}
	messageId := browser.addEmitCallback(callback)
	browser.processMessage.SendProcessMessageForIPC(messageId, name, consts.PID_RENDER, target, argument...)
}

//CheckOnEvent
// IPC 检查 GO 中监听的事件是否存在, 并返回回调函数
func CheckOnEvent(name string) *callback {
	if browser == nil || name == "" {
		return nil
	}
	browser.onLock.Lock()
	defer browser.onLock.Unlock()
	if fn, ok := browser.onEvent[name]; ok {
		return fn
	}
	return nil
}

//CheckEmitCallback
// IPC 检查 GO Emit 回调函数是否存在，并返回回调函数
func CheckEmitCallback(id int32) *reflect.Value {
	if browser == nil {
		return nil
	}
	browser.emitLock.Lock()
	defer browser.emitLock.Unlock()
	if fn, ok := browser.emitCallback[id]; ok {
		delete(browser.emitCallback, id) //移除这个回调函数
		return fn                        //返回
	}
	return nil
}

// addOnEvent 添加监听事件
func (m *browserIPC) addOnEvent(name string, fn *callback) {
	if m == nil || name == "" || fn == nil {
		return
	}
	m.onLock.Lock()
	defer m.onLock.Unlock()
	m.onEvent[name] = fn
}

// emitOnEvent 触发监听事件
func (m *browserIPC) emitOnEvent(name string, argumentList IArrayValue) {
	if m == nil || name == "" || argumentList == nil {
		return
	}
	m.onLock.Lock()
	defer m.onLock.Unlock()
}

// addOnEvent 添加emit回调函数
func (m *browserIPC) addEmitCallback(fn any) int32 {
	if m == nil || fn == nil {
		return 0
	}
	rv := reflect.ValueOf(fn)
	//必须是一个函数
	if rv.Kind() != reflect.Func {
		return 0
	}
	m.emitLock.Lock()
	defer m.emitLock.Unlock()
	if m.emitCallbackMessageId == -1 {
		m.emitCallbackMessageId = 1
	} else {
		m.emitCallbackMessageId++
	}
	m.emitCallback[m.emitCallbackMessageId] = &rv
	return m.emitCallbackMessageId
}
