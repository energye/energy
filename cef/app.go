//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/cef/base"
	"github.com/energye/cef/config"
	engLCL "github.com/energye/energy/v3/lcl"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/tool/exec"
	"path/filepath"
)

var (
	// GApplication global application instance
	GApplication      *Application
	GCEFWorkScheduler ICEFWorkScheduler
)

type Application struct {
	ICEFApplication
}

func NewApplication() *Application {
	if GApplication == nil {
		GApplication = &Application{
			ICEFApplication: NewCEFApplication(),
		}
		base.SetGlobalCEFApplication(GApplication.Instance())
	}
	if !tool.IsDarwin() {
		// Set up CEF Framework
		// Non-macOS platforms require specifying the CEF framework directory;
		// no configuration is needed if the executable is located inside the CEF directory.
		cfg := config.GConfig
		if cfg != nil {
			libCef := func() string {
				if tool.IsWindows() {
					return "libcef.dll"
				} else if tool.IsLinux() {
					return "libcef.so"
				}
				return ""
			}()
			if libCef != "" {
				if frameworkDir := cfg.ChromiumPath(); tool.IsExist(filepath.Join(frameworkDir, libCef)) {
					GApplication.SetCEFFrameworkDir(frameworkDir)
				} else {
					execDir := exec.AppDir()
					if tool.IsExist(filepath.Join(execDir, libCef)) {
						GApplication.SetCEFFrameworkDir(execDir)
					}
				}
			}
		}
	}
	return GApplication
}

// SetCEFFrameworkDir Set unified key paths for CEF framework
func (m *Application) SetCEFFrameworkDir(path string) {
	m.ICEFApplication.SetFrameworkDirPath(path)
	m.ICEFApplication.SetResourcesDirPath(path)
	m.ICEFApplication.SetLocalesDirPath(filepath.Join(path, "locales"))
}

// SetMessageLoop 消息轮询, CEF Application 在不同的 OS 使用不同的配置
func (m *Application) SetMessageLoop() {
	if tool.IsDarwin() { // Darwin => LCL窗口
		if m.IsMainProcess() {
			GCEFWorkScheduler = NewWorkScheduler(nil)
			base.SetGlobalCEFWorkSchedule(GCEFWorkScheduler.Instance())
			m.SetOnScheduleMessagePumpWork(func(delayMs int64) {
				GCEFWorkScheduler.ScheduleMessagePumpWork(delayMs)
			})
		}
		m.SetOnScheduleMessagePumpWork(nil)
		m.SetExternalMessagePump(true)
		m.SetMultiThreadedMessageLoop(false)
	} else { // Windows, Linux => LCL窗口
		// TODO Linux 需要 Gtk3
		m.SetExternalMessagePump(false)
		m.SetMultiThreadedMessageLoop(true)
	}
}

func Run(forms ...lcl.IEngForm) {
	if GApplication == nil || !GApplication.IsValid() {
		println("CEF Application 实例未初始化")
		return
	}
	const (
		LOGSEVERITY_DISABLE = 99
	)
	GApplication.SetLogSeverity(LOGSEVERITY_DISABLE)
	GApplication.SetEnablePrintPreview(true)
	GApplication.SetEnablePrintPreview(true)
	if tool.IsDarwin() {
		base.AddCrDelegate()
		GApplication.InitLibLocationFromArgs()
		GApplication.SetExternalMessagePump(true)
		GApplication.SetMultiThreadedMessageLoop(false)
	}
	GApplication.SetMessageLoop()

	if GApplication.IsMainProcess() {
		isSuccess := GApplication.StartMainProcess()
		if isSuccess {
			api.SetOnReleaseCallback(func() {
				GApplication.ICEFApplication.Free()
			})
			// LCL Application
			engLCL.Run(forms...)
		}
	} else if tool.IsDarwin() && !GApplication.SingleProcess() && !GApplication.IsMainProcess() {
		GApplication.StartSubProcess()
		GApplication.ICEFApplication.Free()
	}
}
