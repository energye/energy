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
	"testing"
)

// emit方式发送到render
func TestBrowserOn(t *testing.T) {
	UseNetIPCChannel = false
	IPC.browser.SetOnEvent(func(event IEventOn) {
		event.On("browserOnTest", func(context IIPCContext) {
			args := NewArgumentList()
			defer args.Clear()
			args.SetInt8(0, 1)
			args.SetInt16(1, 122)
			args.SetInt(2, 256)
			args.SetString(3, "字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串")
			args.SetString(4, "browser on test")
			IPC.browser.Emit("renderOnTest", args)
		})
		event.On("browserOnTest2", func(context IIPCContext) {
			fmt.Println("browserOnTest2:", context.Arguments().Size())
			args := NewArgumentList()
			defer args.Clear()
			args.SetInt8(0, 1)
			args.SetInt16(1, 122)
			args.SetInt(2, 256)
			args.SetString(3, "字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串")
			args.SetString(4, "browser on test")
			args.SetString(5, "render on test")
			args.SetString(6, "render on test")
			IPC.browser.Emit("renderOnTest2", args)
		})
	})
	IPC.StartBrowserIPC()
}

// 响应的方式返回给render
func TestBrowserOnResponse(t *testing.T) {
	UseNetIPCChannel = false
	IPC.browser.SetOnEvent(func(event IEventOn) {
		event.On("browserOnTest", func(context IIPCContext) {
			fmt.Println("browserOnTest:", context.Arguments().Size())
			args := NewArgumentList()
			defer args.Clear()
			args.SetInt8(0, 1)
			args.SetInt16(1, 122)
			args.SetInt(2, 256)
			args.SetString(3, "字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串")
			args.SetString(4, "browser on test")
			context.Response(args.Package())
		})
		event.On("browserOnTest2", func(context IIPCContext) {
			fmt.Println("browserOnTest2:", context.Arguments().Size())
			args := NewArgumentList()
			defer args.Clear()
			args.SetInt8(0, 1)
			args.SetInt16(1, 122)
			args.SetInt(2, 256)
			args.SetString(3, "字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串")
			args.SetString(4, "browser on test")
			args.SetString(5, "render on test")
			args.SetString(6, "render on test")
			context.Response(args.Package())
		})
	})
	IPC.StartBrowserIPC()
}

// emit方式发送到指定的render channel channelId
func TestBrowserOnByRender(t *testing.T) {
	UseNetIPCChannel = false
	IPC.browser.SetOnEvent(func(event IEventOn) {
		event.On("TestBrowserOnByRender", func(context IIPCContext) {
			fmt.Println("TestBrowserOnByRender:", context.Arguments().Size(), "channelId:", context.ChannelId())
			args := NewArgumentList()
			defer args.Clear()
			args.SetInt8(0, 1)
			args.SetInt16(1, 122)
			args.SetInt(2, 256)
			args.SetString(3, "字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串")
			args.SetString(4, "browser on TestBrowserOnByRender")
			args.SetFloat64(5, 555.666)
			IPC.browser.EmitChannelId(fmt.Sprintf("testRender%dOnEmit", context.ChannelId()), context.ChannelId(), args)
		})
	})
	IPC.StartBrowserIPC()
}
