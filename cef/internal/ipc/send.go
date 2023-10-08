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
	"github.com/energye/energy/v2/cef/ipc/target"
)

// emitSendToChannel
//  trigger the specified target Go channel event
func emitSendToGoChannel(messageId int32, tag target.ITarget, eventName string, arguments []any) {
	message := &argument.List{
		Id:        messageId,
		Name:      InternalIPCGoExecuteGoEvent,
		EventName: eventName,
		Data:      arguments,
	}
	if isMainProcess {
		BrowserChan().IPC().Send(tag.ChannelId(), message.Bytes())
	} else {
		message.BId = RenderChan().BrowserId()
		if tag.TargetType() == target.TgGoSub {
			RenderChan().IPC().SendToChannel(tag.ChannelId(), message.Bytes())
		} else if tag.TargetType() == target.TgGoMain {
			RenderChan().IPC().Send(message.Bytes())
		}
	}
	message.Reset()
}
