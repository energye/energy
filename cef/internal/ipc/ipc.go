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
	"github.com/energye/energy/cef/ipc/callback"
	"github.com/energye/energy/cef/ipc/context"
	"github.com/energye/energy/cef/ipc/target"
	"github.com/energye/energy/cef/ipc/types"
	"github.com/energye/energy/cef/process"
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
	processMessage        types.IProcessMessage
}

// SyncChan
//	IPC synchronous receiving data channel
type SyncChan struct {
	lock           sync.Mutex  //
	isClose        bool        //is closed
	timer          *time.Timer //
	ResultSyncChan chan any    //receive synchronization chan, default nil
	delay          time.Duration
}

func init() {
	isMainProcess = process.Args.IsMain()
	isSubProcess = process.Args.IsRender()
	if isMainProcess || isSubProcess {
		browser = &browserIPC{onEvent: make(map[string]*callback.Callback), emitCallback: make(map[int32]*callback.Callback)}
	}
}

// createCallback
//	Create and return a callback function
func createCallback(fn any) *callback.Callback {
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
//	Set the process message object
//	without manually calling it
func SetProcessMessage(pm types.IProcessMessage) {
	if pm == nil {
		return
	}
	browser.processMessage = pm
}

//On
//  IPC GO Listening for events
func On(name string, fn any, options ...types.OnOptions) {
	if name == "" || fn == nil {
		return
	}
	var isOn = false
	if options != nil && len(options) > 0 {
		option := options[0]
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
			browser.addOnEvent(name, callbackFN)
		}
	}
}

//RemoveOn
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
//  Event that triggers listening
//  default to triggering the main process
func Emit(name string, argument ...any) {
	if name == "" || browser.processMessage == nil {
		return
	}
	browser.processMessage.EmitRender(0, name, nil, argument...)
}

// EmitAndCallback
//  Event that triggers listening
//  with callback function
//  default to the main process
func EmitAndCallback(name string, argument []any, fn any) {
	if name == "" || browser.processMessage == nil {
		return
	}
	messageId := browser.addEmitCallback(fn)
	if ok := browser.processMessage.EmitRender(messageId, name, nil, argument...); !ok {
		if messageId > 0 {
			removeEmitCallback(messageId)
		}
	}
}

// EmitTarget
//  Trigger an event for the specified target to listen to
func EmitTarget(name string, tag target.ITarget, argument ...any) {
	if name == "" {
		return
	}
	if tag != nil && tag.TargetType() == target.TgGo {
		if tag.ChannelId() > 0 {
			emitSendToGoChannel(0, tag.ChannelId(), name, argument)
			return
		}
	}
	if browser.processMessage == nil {
		return
	}
	browser.processMessage.EmitRender(0, name, tag, argument...)
}

// EmitTargetAndCallback
//  Trigger an event with a callback function for the specified target to listen on
func EmitTargetAndCallback(name string, tag target.ITarget, argument []any, fn any) {
	if name == "" {
		return
	}
	var messageId int32 = 0
	if tag != nil && tag.TargetType() == target.TgGo {
		if tag.ChannelId() > 0 {
			messageId = browser.addEmitCallback(fn)
			emitSendToGoChannel(messageId, tag.ChannelId(), name, argument)
			return
		}
	}
	if browser.processMessage == nil {
		return
	}
	messageId = browser.addEmitCallback(fn)
	if ok := browser.processMessage.EmitRender(messageId, name, tag, argument...); !ok {
		if messageId > 0 {
			removeEmitCallback(messageId)
		}
	}
}

// CheckOnEvent
//  IPC checks if the event listening in GO exists
//  returns the function and removes it
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
//  IPC checks if the GO Emit callback function exists
//  returns the function and removes it
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
//	Delete callback function for specified message ID
func removeEmitCallback(id int32) {
	browser.emitLock.Lock()
	defer browser.emitLock.Unlock()
	delete(browser.emitCallback, id)
}

// addOnEvent
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
//	Trigger listening event
func (m *browserIPC) emitOnEvent(name string, argumentList types.IArrayValue) {
	if m == nil || name == "" || argumentList == nil {
		return
	}
	m.onLock.Lock()
	defer m.onLock.Unlock()
}

// addOnEvent
//	Add emit callback function
func (m *browserIPC) addEmitCallback(fn any) int32 {
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

// SetDelayTime
//  Set maximum delay time in milliseconds
func (m *SyncChan) SetDelayTime(delay time.Duration) {
	m.delay = delay
}

// Stop Immediate stop delay
func (m *SyncChan) Stop() {
	if m.timer != nil {
		if !m.isClose {
			m.timer.Stop()
		}
		m.isClose = true
	}
}

// DelayWaiting
//	Synchronous message, delay, default 5000 milliseconds
func (m *SyncChan) DelayWaiting() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.isClose = false
	if m.ResultSyncChan == nil {
		m.ResultSyncChan = make(chan any)
		if m.delay == 0 {
			m.delay = 5 * time.Second
		}
	}
	if m.timer == nil {
		m.timer = time.AfterFunc(m.delay, func() {
			if !m.isClose {
				m.isClose = true
				m.ResultSyncChan <- nil
			}
		})
	} else {
		m.timer.Reset(m.delay)
	}
}
