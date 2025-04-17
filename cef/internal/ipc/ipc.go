//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// IPC - Based on pkgs IPC, CEF Internal implementation
// event listeners
// event triggered

package ipc

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc/callback"
	"github.com/cyber-xxm/energy/v2/cef/ipc/context"
	"github.com/cyber-xxm/energy/v2/cef/ipc/target"
	"github.com/cyber-xxm/energy/v2/cef/ipc/types"
	"github.com/cyber-xxm/energy/v2/cef/process"
	"reflect"
	"sync"
	"time"
)

var (
	isMainProcess bool
	isSubProcess  bool
	browser       *browserIPC
)

// EmitContextCallback IPC context Callback Function
type EmitContextCallback func(context context.IContext)

// browserIPC browser process IPC
type browserIPC struct {
	onEvent               map[string]*callback.Callback
	emitCallback          map[int32]*callback.Callback
	emitCallbackMessageId int32
	onLock                sync.Mutex
	emitLock              sync.Mutex
	window                target.IWindow
	browserWindow         target.IBrowserWindow
}

// WaitChan
//
//	IPC synchronous receiving data channel
type WaitChan struct {
	count   int32
	Pending *sync.Map
}

func init() {
	isMainProcess = process.Args.IsMain()
	isSubProcess = process.Args.IsRender()
	if isMainProcess || isSubProcess {
		browser = &browserIPC{onEvent: make(map[string]*callback.Callback), emitCallback: make(map[int32]*callback.Callback)}
	}
}

// createCallback
//
//	Create and return a callback function
func createCallback(fn interface{}) *callback.Callback {
	switch fn.(type) {
	case func(context context.IContext):
		return &callback.Callback{Context: &callback.ContextCallback{Callback: fn.(func(context context.IContext))}}
	default:
		v := reflect.ValueOf(fn)
		// fn must be a function
		if v.Kind() != reflect.Func {
			return nil
		}
		return &callback.Callback{Argument: &callback.ArgumentCallback{Callback: &v}}
	}
}

// SetProcessMessage
//
//	Set the process message object
//	without manually calling it
func SetProcessMessage(pm target.IWindow) {
	if pm == nil {
		return
	}
	browser.window = pm
}

// SetBrowserWindow
// Set BrowserWindow on initialization
func SetBrowserWindow(bw target.IBrowserWindow) {
	if browser.browserWindow == nil {
		browser.browserWindow = bw
	}
}

// On
//
//	IPC GO Listening for events
func On(name string, fn interface{}, options ...types.OnOptions) {
	if name == "" || fn == nil {
		return
	}
	var (
		isOn   = false
		option *types.OnOptions
	)
	if options != nil && len(options) > 0 && !cef.Application().SingleProcess() {
		option = &options[0]
		if option.OnType == types.OtAll {
			isOn = true
		} else if option.OnType == types.OtMain && isMainProcess {
			isOn = true
		} else if option.OnType == types.OtSub && isSubProcess {
			isOn = true
		}
	} else if isMainProcess || isSubProcess {
		isOn = true
	}
	if isOn {
		if callbackFN := createCallback(fn); callbackFN != nil {
			if option != nil {
				callbackFN.IsAsync = option.Mode == types.MAsync
			}
			browser.addOnEvent(name, callbackFN)
		}
	}
}

// RemoveOn
// IPC GO Remove listening events
func RemoveOn(name string) {
	if name == "" {
		return
	}
	browser.onLock.Lock()
	defer browser.onLock.Unlock()
	delete(browser.onEvent, name)
}

// Emit
//
//	Event that triggers listening
//	default to triggering the main process
func Emit(name string, argument ...interface{}) bool {
	if name == "" || browser.window == nil {
		return false
	}
	// When the window is closed
	if browser.window == nil || browser.window.IsClosing() {
		// This window is the first one created and not closed
		SetProcessMessage(browser.browserWindow.LookForMainWindow())
	}
	browser.window.ProcessMessage().EmitRender(0, name, nil, argument...)
	return true
}

// EmitAndCallback
//
//	Event that triggers listening
//	with callback function
//	default to the main process
func EmitAndCallback(name string, argument []interface{}, fn interface{}) bool {
	if name == "" || browser.window == nil {
		return false
	}
	// When the window is closed
	if browser.window == nil || browser.window.IsClosing() {
		// This window is the first one created and not closed
		SetProcessMessage(browser.browserWindow.LookForMainWindow())
	}
	messageId := browser.addEmitCallback(fn)
	if ok := browser.window.ProcessMessage().EmitRender(messageId, name, nil, argument...); !ok {
		//fail in send
		if messageId > 0 {
			removeEmitCallback(messageId)
		}
		return false
	}
	return true
}

// EmitTarget
//
//	Trigger an event for the specified target to listen to
func EmitTarget(name string, tag target.ITarget, argument ...interface{}) bool {
	if name == "" {
		return false
	}
	var window = tag.Window()
	if window == nil {
		// default window
		window = browser.window
	}
	if window == nil {
		return false
	}
	if window.IsClosing() {
		return false
	}
	window.ProcessMessage().EmitRender(0, name, tag, argument...)
	return true
}

// EmitTargetAndCallback
//
//	Trigger an event with a callback function for the specified target to listen on
func EmitTargetAndCallback(name string, tag target.ITarget, argument []interface{}, fn interface{}) bool {
	if name == "" {
		return false
	}
	var window = tag.Window()
	if window == nil {
		// default window
		window = browser.window
	}
	if window == nil {
		return false
	}
	if window.IsClosing() {
		return false
	}
	messageId := browser.addEmitCallback(fn)
	if ok := window.ProcessMessage().EmitRender(messageId, name, tag, argument...); !ok {
		if messageId > 0 {
			removeEmitCallback(messageId)
		}
		return false
	}
	return true
}

// CheckOnEvent
//
//	IPC checks if the event listening in GO exists
//	returns the function and removes it
func CheckOnEvent(name string) *callback.Callback {
	if name == "" {
		return nil
	}
	browser.onLock.Lock()
	defer browser.onLock.Unlock()
	if fn, ok := browser.onEvent[name]; ok {
		return fn
	}
	return nil
}

// CheckEmitCallback
//
//	IPC checks if the GO Emit callback function exists
//	returns the function and removes it
func CheckEmitCallback(id int32) *callback.Callback {
	browser.emitLock.Lock()
	defer browser.emitLock.Unlock()
	if fn, ok := browser.emitCallback[id]; ok {
		delete(browser.emitCallback, id)
		return fn
	}
	return nil
}

// removeEmitCallback
//
//	Delete callback function for specified message ID
func removeEmitCallback(id int32) {
	browser.emitLock.Lock()
	defer browser.emitLock.Unlock()
	delete(browser.emitCallback, id)
}

// addOnEvent
//
//	Add listening event
//	callback function
//	  1. context 2.argument list
func (m *browserIPC) addOnEvent(name string, fn *callback.Callback) {
	if m == nil || name == "" || fn == nil {
		return
	}
	m.onLock.Lock()
	defer m.onLock.Unlock()
	m.onEvent[name] = fn
}

// emitOnEvent
//
//	Trigger listening event
func (m *browserIPC) emitOnEvent(name string, argumentList types.IArrayValue) {
	if m == nil || name == "" || argumentList == nil {
		return
	}
	m.onLock.Lock()
	defer m.onLock.Unlock()
}

// addOnEvent
//
//	Add emit callback function
func (m *browserIPC) addEmitCallback(fn interface{}) int32 {
	if m == nil || fn == nil {
		return 0
	}
	m.emitLock.Lock()
	defer m.emitLock.Unlock()
	if callbackFN := createCallback(fn); callbackFN != nil {
		if m.emitCallbackMessageId == -1 {
			m.emitCallbackMessageId = 1
		} else {
			m.emitCallbackMessageId++
		}
		m.emitCallback[m.emitCallbackMessageId] = callbackFN
		return m.emitCallbackMessageId
	}
	return 0
}

// 计时器
type delayer struct {
	stop  bool
	timer *time.Timer
}

// Stop Immediate stop delay
func (m *delayer) Stop() {
	if m.timer != nil {
		if !m.stop {
			m.timer.Stop()
		}
		m.stop = true
		m.timer = nil
	}
}

// NextMessageId 返回下一个消息ID
func (m *WaitChan) NextMessageId() (id int32) {
	id = m.count
	m.count++
	return
}

// NewDelayer 创建一个计时器
func (m *WaitChan) NewDelayer(messageId int32, delay time.Duration) *delayer {
	md := new(delayer)
	md.timer = time.AfterFunc(delay, func() {
		if !md.stop {
			md.stop = true
			m.Done(messageId, nil)
		}
	})
	return md
}

func (m *WaitChan) Done(messageId int32, data interface{}) {
	if val, ok := m.Pending.Load(messageId); ok {
		val.(func(result interface{}))(data)
		m.Pending.Delete(messageId)
	}
}
