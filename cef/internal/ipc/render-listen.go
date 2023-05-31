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
	"github.com/energye/energy/v2/cef/ipc/argument"
	"github.com/energye/energy/v2/cef/ipc/context"
	"github.com/energye/energy/v2/cef/process"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/pkgs/channel"
	"github.com/energye/energy/v2/pkgs/json"
)

// listen
func (m *renderIPCChan) listen(context channel.IIPCContext, argumentList argument.IList) bool {
	if argumentList != nil {
		if argumentList.GetName() == InternalIPCGoExecuteGoEvent { // Go emit Go Event
			m.goExecuteGoEvent(context, argumentList)
			return true
		} else if argumentList.GetName() == InternalIPCGoExecuteJSEventReplay { // Go emit Go Event Replay
			m.goExecuteGoEventReplay(context, argumentList)
			return true
		}
	}
	return false
}

// goExecuteGoEventReplay
func (m *renderIPCChan) goExecuteGoEventReplay(ctx channel.IIPCContext, argumentList argument.IList) {
	emitCallback := CheckEmitCallback(argumentList.MessageId())
	if emitCallback != nil {
		var argumentJSONArray json.JSONArray
		if argumentList.JSON() != nil {
			argumentJSONArray = argumentList.JSON().JSONArray()
		}
		ipcContext := context.NewContext(process.BrowserId(), ctx.ChannelId(), false, argumentJSONArray)
		// Call listener function
		// Based on the currently defined event listening
		// 1. Callback Function - Context Mode
		// 2. Callback Function - Parameter List Method
		if ctxCallback := emitCallback.ContextCallback(); ctxCallback != nil {
			ctxCallback.Invoke(ipcContext)
		} else if argsCallback := emitCallback.ArgumentCallback(); argsCallback != nil {
			argsCallback.Invoke(ipcContext)
		}
		if ipcContext.ArgumentList() != nil {
			ipcContext.ArgumentList().Free()
		}
		ipcContext.Result(nil)
	}
}

// goExecuteGoEvent
func (m *renderIPCChan) goExecuteGoEvent(ctx channel.IIPCContext, argumentList argument.IList) {
	messageId := argumentList.MessageId()
	eventName := argumentList.GetEventName()
	var ipcContext context.IContext
	if eventCallback := CheckOnEvent(eventName); eventCallback != nil {
		var argumentJSONArray json.JSONArray
		if argumentList.JSON() != nil {
			argumentJSONArray = argumentList.JSON().JSONArray()
		}
		ipcContext = context.NewContext(process.BrowserId(), ctx.ChannelId(), messageId != 0, argumentJSONArray)
		// Call listener function
		// Based on the currently defined event listening
		// 1. Callback Function - Context Mode
		// 2. Callback Function - Parameter List Method
		if ctxCallback := eventCallback.ContextCallback(); ctxCallback != nil {
			ctxCallback.Invoke(ipcContext)
		} else if argsCallback := eventCallback.ArgumentCallback(); argsCallback != nil {
			argsCallback.Invoke(ipcContext)
		}
	}
	if messageId != 0 {
		replyMessage := &argument.List{
			Id:   messageId,
			Name: InternalIPCGoExecuteJSEventReplay,
		}
		if ipcContext != nil {
			replay := ipcContext.Replay()
			if replay.Result() != nil && len(replay.Result()) > 0 {
				replyMessage.Data = replay.Result()
			}
		}
		if ctx.ProcessId() == consts.PID_RENDER {
			m.ipc.SendToChannel(ctx.ChannelId(), replyMessage.Bytes())
		} else {
			m.ipc.Send(replyMessage.Bytes())
		}
		// free
		replyMessage.Reset()
	}
	if ipcContext != nil {
		if ipcContext.ArgumentList() != nil {
			ipcContext.ArgumentList().Free()
		}
		ipcContext.Result(nil)
	}
}
