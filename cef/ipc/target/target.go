//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// IPC event to specified receiving destination

package target

import "github.com/energye/cef/cef"

// Type
//
//	0: Trigger the JS event of the specified target process
//	1: Trigger TgGoSub events for the specified target sub process
//	2: Trigger TgGoMain events for the specified target main process
type Type int8

const (
	TgJs     Type = iota //JS Event
	TgGoSub              //GO Event sub
	TgGoMain             //GO Event main
)

// ITarget
//
// ipc.NewTarget() *Target
type ITarget interface {
	BrowserId() int32 // Browser Window ID
	ChannelId() int64 // IPC channelID, frameId or GO IPC channelID
	TargetType() Type // Target type default 0: Trigger JS event
	Window() IWindow  // Send IPC Chromium
}

// IProcessMessage
// Send IPC Chromium
type IProcessMessage interface {
	EmitRender(messageId int32, eventName string, target ITarget, data ...interface{}) bool
	SendProcessMessageForJSONBytes(messageName string, targetProcess cef.TCefProcessId, data []byte)
}

// IWindow for IPC
type IWindow interface {
	Target(targetType ...Type) ITarget // return IPC target
	IsClosing() bool                   // Whether the window is closed
	ProcessMessage() IProcessMessage   // process message, chromium
}

// IBrowserWindow
//
//	BrowserWindow for IPC
type IBrowserWindow interface {
	LookForMainWindow() (window IWindow) //select a new window, This window is the first one created and not closed
}

// Target Go IPC
//
//	receiving target of the event
type Target struct {
	window     IWindow
	browseId   int32
	channelId  int64
	targetType Type
}

// NewTarget Create a new Emit target
//
//	targetChromium: current window chromium, Use the main window chromium when nil
//	browserId: browser window ID
//	channelId: IPC channelID, frameId or GO IPC channelID
//	targetType: Optional parameter, target type default 0
//	  Type: TgJs:JS Event, TgGoSub:GO Sub Event, TgGoMain:GO Main Event
func NewTarget(targetChromium IWindow, browserId int32, channelId int64, targetType ...Type) ITarget {
	m := &Target{
		window:    targetChromium,
		browseId:  browserId,
		channelId: channelId,
	}
	if len(targetType) > 0 {
		m.targetType = targetType[0]
	}
	return m
}

// NewTargetMain Create a new Emit target Main Process
//
//	targetType: TgGoMain
func NewTargetMain() ITarget {
	return &Target{
		targetType: TgGoMain,
	}
}

// TargetType
//
//	target type
//	0: Trigger JS event
//	1: Trigger Go Event
func (m *Target) TargetType() Type {
	return m.targetType
}

// BrowserId
//
//	return BrowserId
func (m *Target) BrowserId() int32 {
	return m.browseId
}

// ChannelId
//
//	return ChannelId
func (m *Target) ChannelId() int64 {
	return m.channelId
}

// Window
//
//	return chromium ProcessMessage
func (m *Target) Window() IWindow {
	return m.window
}
