package src

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/cef/ipc/callback"
	"github.com/cyber-xxm/energy/v2/cef/ipc/context"
	"github.com/energye/golcl/lcl/rtl/version"
	"time"
)

func BrowserProcessStart(b bool) {
	ipc.On("os-info", func(context context.IContext) {
		fmt.Println("os-info", version.OSVersion.ToString())
		context.Result(version.OSVersion.ToString())
	})
	ipc.On("minimize", func(channel callback.IChannel) {
		if win := cef.BrowserWindow.GetWindowInfo(channel.BrowserId()); win != nil {
			win.Minimize()
		}
	})
	ipc.On("maximize", func(channel callback.IChannel) bool {
		if win := cef.BrowserWindow.GetWindowInfo(channel.BrowserId()); win != nil {
			win.Maximize()
			return true
		}
		return false
	})
	ipc.On("close", func(channel callback.IChannel) {
		if win := cef.BrowserWindow.GetWindowInfo(channel.BrowserId()); win != nil {
			win.CloseBrowserWindow()
		}
	})
	go func() {
		for {
			time.Sleep(time.Second)
			ipc.EmitAndCallback("testasync", nil, func(data string) {
				fmt.Println("data:", data)
			})
		}
	}()
}

func BrowserInit(event *cef.BrowserEvent, window cef.IBrowserWindow) {
	//
}
