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
)

// emitSendToChannel
//  trigger the specified target Go channel event
func emitSendToGoChannel(messageId int32, channelId int64, eventName string, arguments []any) {
	message := &argument.List{
		Id:        messageId,
		Name:      InternalIPCGoExecuteGoEvent,
		EventName: eventName,
		Data:      arguments,
	}
	if isMainProcess {
		BrowserChan().IPC().Send(channelId, message.Bytes())
	} else {
		message.BId = RenderChan().BrowserId()
		RenderChan().IPC().SendToChannel(channelId, message.Bytes())
	}
	message.Reset()
}
