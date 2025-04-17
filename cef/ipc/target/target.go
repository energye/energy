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

import "github.com/cyber-xxm/energy/v2/consts"

// ITarget
//
// ipc.NewTarget() *Target
type ITarget interface {
	BrowserId() int32  // Browser Window ID
	ChannelId() string // IPC channelID, frameId or GO IPC channelID
	Window() IWindow   // Send IPC Chromium
}

// IProcessMessage
// Send IPC Chromium
type IProcessMessage interface {
	EmitRender(messageId int32, eventName string, target ITarget, data ...interface{}) bool
	SendProcessMessageForJSONBytes(messageName string, targetProcess consts.CefProcessId, data []byte)
}

// IWindow for IPC
type IWindow interface {
	Target() ITarget                 // return IPC target
	IsClosing() bool                 // Whether the window is closed
	ProcessMessage() IProcessMessage // process message, chromium
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
	window    IWindow
	browseId  int32
	channelId string
}

// NewTarget Create a new Emit target
//
//	targetChromium: current window chromium, Use the main window chromium when nil
//	browserId: browser window ID
//	channelId: IPC channelID, frameId or GO IPC channelID
//	targetType: Optional parameter, target type default 0
//	  Type: TgJs:JS Event, TgGoSub:GO Sub Event, TgGoMain:GO Main Event
func NewTarget(targetChromium IWindow, browserId int32, channelId string) ITarget {
	m := &Target{
		window:    targetChromium,
		browseId:  browserId,
		channelId: channelId,
	}
	return m
}

// NewTargetMain Create a new Emit target Main Process
//
//	targetType: TgGoMain
func NewTargetMain() ITarget {
	return &Target{}
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
func (m *Target) ChannelId() string {
	return m.channelId
}

// Window
//
//	return chromium ProcessMessage
func (m *Target) Window() IWindow {
	return m.window
}
