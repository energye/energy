//go:build windows
// +build windows

package main

import "C"
import (
	"github.com/cyber-xxm/energy/v2/cef/config"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/messages"
	"github.com/energye/golcl/pkgs/libname"
	"path/filepath"
	"runtime"

	"github.com/energye/golcl/lcl/rtl"

	"github.com/energye/golcl/lcl"

	"github.com/energye/golcl/lcl/win"
)

//export mainFormShow
func mainFormShow() {
	defer func() {
		if err := recover(); err != nil {
			Println("err: ", err)
		}
	}()
	Println("Libenergy.dll MainForm.Show")
	if MainForm != nil {
		MainForm.Show()
	}
}

//export mainFormClose
func mainFormClose() {
	Println("Libenergy.dll mainFormClose")
	if MainForm != nil {
		MainForm.Close()
	}
}

//export mainFormFree
func mainFormFree() {
	Println("Libenergy.dll mainFormFree")
	if MainForm != nil {
		MainForm.Close()
		MainForm.Free()
		MainForm = nil
		lcl.Application.Terminate()
		api.EnergyLibRelease()
	}
}

//export initApplication
func initApplication() {
	defer func() {
		if err := recover(); err != nil {
			Println("Libenergy.dll initApplication err: ", err)
		}
	}()
	Println("initApplication")
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	libname.LibName = filepath.Join(config.Get().FrameworkPath(), "liblcl.dll")
	inits.Init(nil, nil)
	Println("Libenergy.dll initApplication - inits.Init")

	rtl.InitGoDll(0)

	moduleHandle := win.GetSelfModuleHandle()
	if moduleHandle > 0 {
		lcl.Application.Icon().SetHandle(win.LoadIcon(moduleHandle, 100))
	}
	if MainForm == nil {
		lcl.CreateResForm(lcl.Application, &MainForm)
		MainForm.SetShowInTaskBar(types.StAlways)
		MainForm.SetFormStyle(types.FsStayOnTop)
	}
}

//export messageProc
func messageProc(iMessage uint32, wParam uintptr, lParam uintptr) uintptr {
	if iMessage == messages.WM_DESTROY {
		Println("Libenergy.dll WM_DESTROY")
	}
	return 1
}
