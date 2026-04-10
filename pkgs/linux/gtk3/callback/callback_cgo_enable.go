//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build cgo

package callback

/*
#cgo pkg-config: glib-2.0 gobject-2.0 gtk+-3.0
#include <stdlib.h>
#include <gtk/gtk.h>
#include <stdint.h>
#include <callback_cgo_enable.go.h>


static void remove_signal_handler(GtkWidget* widget, gulong handler_id) {
  	g_print("尝试移除信号处理器: handler_id=%lu, widget=%p\n", handler_id, widget);
    if (handler_id > 0 && widget != NULL) {
        g_signal_handler_disconnect(widget, handler_id);
        g_print("移除信号处理器成功: handler_id=%lu\n", handler_id);
    }
}
*/
import "C"

import (
	"github.com/energye/energy/v3/pkgs/linux/gtk3/types"
	"reflect"
	"sync"
	"unsafe"
)

type ReturnType int

const (
	RET_VOID ReturnType = iota
	RET_GBOOLEAN
	RET_GINT
	RET_POINTER
)

type doMethodHandler struct {
	method  *reflect.Value
	in, out int
	ret     *reflect.Type
	retType ReturnType
}

type signalManager struct {
	mu       sync.Mutex
	handlers map[uint64]*doMethodHandler
	nextID   uint64
}

var signalMgr = &signalManager{
	handlers: make(map[uint64]*doMethodHandler),
	nextID:   1,
}

func (m *SignalHandlerID) Disconnect() {
	Disconnect(m.Widget, m.HandlerID, m.Id)
}

// Connect 注册信号
func Connect(instance unsafe.Pointer, signalName, trampolineName string, fn any, userData unsafe.Pointer) *SignalHandlerID {
	signalMgr.mu.Lock()
	callbackFn := compileCallback(fn)
	id := signalMgr.nextID
	signalMgr.handlers[id] = callbackFn
	signalMgr.nextID++
	signalMgr.mu.Unlock()
	if userData == nil {
		userData = unsafe.Pointer(uintptr(id))
	}

	cSignalName := C.CString(signalName)
	defer C.free(unsafe.Pointer(cSignalName))

	cTrampolineName := C.CString(trampolineName)
	defer C.free(unsafe.Pointer(cTrampolineName))

	cTrampoline := C.get_trampoline(cTrampolineName)
	if cTrampoline == nil {
		panic("No trampoline for given nArgs")
	}

	handlerID := C.go_g_signal_connect(C.gpointer(instance), cSignalName, cTrampoline, C.gpointer(userData))
	return &SignalHandlerID{
		Widget:    types.PGtkWidget(instance),
		HandlerID: types.GULong(handlerID),
		Id:        id,
	}
}

func Disconnect(widget types.PGtkWidget, handlerID types.GULong, id uint64) {
	if widget != 0 && handlerID > 0 {
		signalMgr.mu.Lock()
		if _, ok := signalMgr.handlers[id]; ok {
			C.remove_signal_handler((*C.GtkWidget)(unsafe.Pointer(widget)), C.gulong(handlerID))
			delete(signalMgr.handlers, id)
		}
		signalMgr.mu.Unlock()
	}
}

//export go_signal_handler_generic
func go_signal_handler_generic(call *C.SignalCall) {
	id := uint64(uintptr(call.user_data))

	signalMgr.mu.Lock()
	handler, ok := signalMgr.handlers[id]
	signalMgr.mu.Unlock()

	if !ok {
		return
	}
	args := make([]reflect.Value, int(call.nargs))
	for i := 0; i < int(call.nargs); i++ {
		args[i] = reflect.ValueOf(uintptr(call.args[i]))
	}
	// call method
	result := handler.method.Call(args)
	if handler.ret != nil {
		ret0 := result[0]
		switch handler.retType {
		case RET_POINTER: // pointer
			call.ret = unsafe.Pointer(ret0.Pointer())
		case RET_GBOOLEAN: // bool
			if ret0.Bool() {
				call.ret = unsafe.Pointer(uintptr(1))
			} else {
				call.ret = unsafe.Pointer(uintptr(0))
			}
		case RET_GINT: // int
			call.ret = unsafe.Pointer(uintptr(ret0.Int()))
		default:
			panic("handler.method.Call, unsupported return type: " + ret0.String())
		}
	}
}

// 底层方法
// fn: 只支持指针和普通类型参数和返回值, 需要通过上层包装返回复杂类型
func compileCallback(fn any) *doMethodHandler {
	val := reflect.ValueOf(fn)
	if val.Kind() != reflect.Func {
		panic("the type must be a function but was not")
	}
	if val.IsNil() {
		panic("function must not be nil")
	}
	ty := val.Type()
	gh := &doMethodHandler{}
	gh.method = &val
	gh.in = ty.NumIn()
	for i := 0; i < gh.in; i++ {
		in := ty.In(i)
		switch in.Kind() {
		case reflect.Struct, reflect.Interface, reflect.Func, reflect.Slice,
			reflect.Chan, reflect.Complex64, reflect.Complex128,
			reflect.String, reflect.Map, reflect.Invalid:
			panic("unsupported argument type: " + in.Kind().String())
		}
	}
	gh.out = ty.NumOut()
	switch {
	case gh.out == 1:
		out0 := ty.Out(0)
		gh.ret = &out0
		switch out0.Kind() {
		case reflect.Pointer, reflect.UnsafePointer: // pointer
			gh.retType = RET_POINTER
		case reflect.Bool: // bool
			gh.retType = RET_GBOOLEAN
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, // int
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr: // int
			gh.retType = RET_GINT
		default:
			panic("unsupported return type: " + ty.String())
		}
	case gh.out > 1:
		panic("callbacks can only have one return")
	}
	return gh
}
