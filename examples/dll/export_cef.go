//go:build windows
// +build windows

package main

import "C"

import (
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/config"
	"github.com/cyber-xxm/energy/v2/cef/process"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/win"
	"github.com/energye/golcl/pkgs/libname"
	"path/filepath"
	"runtime"
)

//export initCEFApplication
func initCEFApplication() {
	defer func() {
		if err := recover(); err != nil {
			Println("Libenergy.dll initCEFApplication err: ", err)
		}
	}()
	Println("initApplication")
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	libname.LibName = filepath.Join(config.Get().FrameworkPath(), "liblcl.dll")
	cef.GlobalInit(nil, nil)
	Println("Libenergy.dll initCEFApplication - inits.Init process-type: ", process.Args.ProcessType())

	//创建应用
	app := cef.NewApplication()
	app.SetSingleProcess(true)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	// 主进程启动
	mainStart := app.StartMainProcess()
	Println("mainStart:", mainStart, process.Args.ProcessType())
	if mainStart {
		// 结束应用后释放资源
		api.SetReleaseCallback(func() {
			Println("Release")
			app.Free()
		})
		rtl.InitGoDll(0)

		moduleHandle := win.GetSelfModuleHandle()
		if moduleHandle > 0 {
			lcl.Application.Icon().SetHandle(win.LoadIcon(moduleHandle, 100))
		}
		if BW == nil {
			lcl.CreateResForm(lcl.Application, &BW)
			BW.SetShowInTaskBar(types.StAlways)
			//BW.SetFormStyle(types.FsStayOnTop)
		}
	}
}

//export cefFormShow
func cefFormShow() {
	Println("Libenergy.dll MainForm.Show")
	if BW != nil {
		BW.Show()
	}
}

//export cefFormClose
func cefFormClose() {
	Println("Libenergy.dll mainFormClose")
	if BW != nil {
		BW.Close()
	}
}

//export cefClose
func cefClose() {
	Println("Libenergy.dll cefClose")
	if BW != nil {
		BW.canClose = true
		BW.chromium.CloseBrowser(true)
	}
}

//export cefFormFree
func cefFormFree() {
	Println("Libenergy.dll mainFormFree")
	if BW != nil {
		BW.Close()
		//BW.Free()
		BW = nil
		lcl.Application.Terminate()
	}
}

//export setWindowHwnd
func setWindowHwnd(ptr uintptr) {
	Println("Libenergy.dll setWindowHwnd:", ptr)
	hostWindowHwnd = ptr
}
