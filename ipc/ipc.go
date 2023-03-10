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
	"fmt"
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"reflect"
	"sync"
)

var (
	browser *browserIPC
)

type EmitContextCallback func(context IContext)
type EmitArgumentCallback any

// browserIPC 主进程 IPC
type browserIPC struct {
	onEvent map[string]*callback
	lock    sync.Mutex
}

type callback struct {
	context  *contextCallback
	argument *argumentCallback
}

type contextCallback struct {
	callback EmitContextCallback //带上下文的回调函数
}

type argumentCallback struct {
	callback        *reflect.Value         //带参数的回调函数
	inArgumentType  []consts.GO_VALUE_TYPE //入参类型
	outArgumentType []consts.GO_VALUE_TYPE //出参类型
}

func (m *callback) ContextCallback() EmitContextCallback {
	if m.context == nil {
		return nil
	}
	return m.context.callback
}

func (m *callback) ArgumentCallback() *argumentCallback {
	if m.argument == nil {
		return nil
	}
	return m.argument
}

func (m *argumentCallback) Callback() *reflect.Value {
	return m.callback
}

func (m *argumentCallback) InArgumentType() []consts.GO_VALUE_TYPE {
	return m.inArgumentType
}

func (m *argumentCallback) OutArgumentType() []consts.GO_VALUE_TYPE {
	return m.outArgumentType
}

func init() {
	if common.Args.IsMain() {
		browser = &browserIPC{onEvent: make(map[string]*callback)}
	}
}

//On
// IPC GO 监听事件, 上下文参数
func On(name string, fn EmitContextCallback) {
	if browser != nil {
		browser.addOnEvent(name, &callback{context: &contextCallback{callback: fn}})
	}
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
//   map: key只能string类型, value 只能是 any 或 interface{}
//   struct: 首字母大写, 字段类型匹配
//     type ArgsStructDemo struct {
//        Key1 string `json:"key1"`
//		  Key2 string `json:"key2"`
//		  Key3 string
//		  Key4 int
//		  Key5 float64
//		  Key6 bool
//     }
//
// 出参
//
//   与同入参相同
func OnArguments(name string, fn EmitArgumentCallback) {
	if browser != nil {
		v := reflect.ValueOf(fn)
		// f must be a function
		if v.Kind() != reflect.Func {
			return
		}
		vt := v.Type()
		inArgumentType := make([]consts.GO_VALUE_TYPE, vt.NumIn())
		outArgumentType := make([]consts.GO_VALUE_TYPE, vt.NumOut())
		for i := 0; i < len(inArgumentType); i++ {
			it := vt.In(i)
			gvt, _ := common.FieldReflectType(it)
			inArgumentType[i] = gvt
			fmt.Println("gvt", gvt)
		}
		for i := 0; i < len(outArgumentType); i++ {
			ot := vt.Out(i)
			gvt, _ := common.FieldReflectType(ot)
			inArgumentType[i] = gvt
		}
		browser.addOnEvent(name, &callback{argument: &argumentCallback{callback: &v, inArgumentType: inArgumentType, outArgumentType: outArgumentType}})
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
func CheckOnEvent(name string) *callback {
	if browser == nil || name == "" {
		return nil
	}
	browser.lock.Lock()
	defer browser.lock.Unlock()
	if fn, ok := browser.onEvent[name]; ok {
		return fn
	}
	return nil
}

// addOnEvent 添加监听事件
func (m *browserIPC) addOnEvent(name string, fn *callback) {
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
