//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package ipc

import (
	"fmt"
	"github.com/energye/energy/commons"
	"testing"
	"time"
)

//emit方式发送到render
func TestRenderOnEmit(t *testing.T) {
	commons.Args.SetArgs("type", "renderer")
	UseNetIPCChannel = false
	IPC.render.SetOnEvent(func(event IEventOn) {
		event.On("renderOnTest", func(context IIPCContext) {
			args := NewArgumentList()
			fmt.Println("renderOnTest:", context.Arguments().Size())
			defer args.Clear()
			args.SetInt8(0, 1)
			args.SetInt16(1, 122)
			args.SetInt(2, 256)
			args.SetString(3, "字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串")
			args.SetString(4, "render on test")
			time.Sleep(time.Second / 1000)
			IPC.render.Emit("browserOnTest", args)
		})
		event.On("renderOnTest2", func(context IIPCContext) {
			args := NewArgumentList()
			fmt.Println("renderOnTest2:", context.Arguments().Size())
			defer args.Clear()
			args.SetInt8(0, 1)
			args.SetInt16(1, 122)
			args.SetInt(2, 256)
			args.SetString(3, "字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串")
			args.SetString(4, "render on test")
			args.SetString(5, "render on test")
			time.Sleep(time.Second / 1000)
			IPC.render.Emit("browserOnTest2", args)
		})
	})
	IPC.CreateRenderIPC(1, 20)
	IPC.render.emitConnect()
	IPC.render.Emit("browserOnTest", nil)
	IPC.render.Emit("browserOnTest2", nil)
	for {
		time.Sleep(time.Second)
	}
}

//响应的方式返回给render
func TestRenderOnResponse(t *testing.T) {
	commons.Args.SetArgs("type", "renderer")
	UseNetIPCChannel = false
	IPC.CreateRenderIPC(1, 20)
	IPC.render.emitConnect()
	IPC.render.EmitAndCallback("browserOnTest", nil, func(context IIPCContext) {
		fmt.Println("响应renderOnTest:", context.Arguments().Size())
	})
	IPC.render.EmitAndCallback("browserOnTest2", nil, func(context IIPCContext) {
		fmt.Println("响应renderOnTest2:", context.Arguments().Size())
	})
	for {
		time.Sleep(time.Second)
	}
}

//render 1
func TestRender1OnEmit(t *testing.T) {
	commons.Args.SetArgs("type", "renderer")
	UseNetIPCChannel = false
	IPC.render.SetOnEvent(func(event IEventOn) {
		event.On("testRender1OnEmit", func(context IIPCContext) {
			args := NewArgumentList()
			fmt.Println("testRender1OnEmit:", context.Arguments().Size())
			defer args.Clear()
			args.SetInt8(0, 1)
			args.SetInt16(1, 122)
			args.SetInt(2, 256)
			args.SetString(3, "字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串")
			args.SetString(4, "render on test")
			time.Sleep(time.Second / 1000)
			//IPC.render.Emit("browserOnTest", args)
		})
	})
	IPC.CreateRenderIPC(1, 20)
	IPC.render.emitConnect()
	IPC.render.Emit("TestBrowserOnByRender", nil)
	for {
		time.Sleep(time.Second)
	}
}

//render 2
func TestRender2OnEmit(t *testing.T) {
	commons.Args.SetArgs("type", "renderer")
	UseNetIPCChannel = false
	IPC.render.SetOnEvent(func(event IEventOn) {
		event.On("testRender2OnEmit", func(context IIPCContext) {
			args := NewArgumentList()
			fmt.Println("testRender2OnEmit:", context.Arguments().Size())
			defer args.Clear()
			args.SetInt8(0, 1)
			args.SetInt16(1, 122)
			args.SetInt(2, 256)
			args.SetString(3, "字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串字符串")
			args.SetString(4, "render on test")
			time.Sleep(time.Second / 1000)
			//IPC.render.Emit("browserOnTest", args)
		})
	})
	IPC.CreateRenderIPC(1, 20)
	IPC.render.emitConnect()
	IPC.render.Emit("TestBrowserOnByRender", nil)
	for {
		time.Sleep(time.Second)
	}
}
