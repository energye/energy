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
	"github.com/energye/cef/cef"
	"github.com/energye/cef/cef/types"
	"github.com/energye/cef/config"
	engLCL "github.com/energye/energy/v3/lcl"
	"github.com/energye/energy/v3/logger"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/tool/exec"
	"path/filepath"
)

var (
	// GApplication global application instance
	GApplication   *Application
	GWorkScheduler cef.ICEFWorkScheduler
)

type Application struct {
	cef.ICefApplication
}

// Init CEF Global initialization, invoked at application startup in main
func Init() *Application {
	lcl.Init()
	base.Init()
	return NewApplication()
}

func NewApplication() *Application {
	if GApplication == nil {
		GApplication = &Application{
			ICefApplication: cef.NewApplication(),
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
	m.ICefApplication.SetFrameworkDirPath(path)
	m.ICefApplication.SetResourcesDirPath(path)
	m.ICefApplication.SetLocalesDirPath(filepath.Join(path, "locales"))
}

// SetMessageLoop 消息轮询, CEF Application 在不同的 OS 使用不同的配置
func (m *Application) SetMessageLoop() {
	if tool.IsDarwin() { // Darwin => LCL窗口
		if m.IsMainProcess() {
			GWorkScheduler = cef.NewWorkScheduler(nil)
			base.SetGlobalCEFWorkSchedule(GWorkScheduler.Instance())
			m.SetOnScheduleMessagePumpWork(func(delayMs int64) {
				GWorkScheduler.ScheduleMessagePumpWork(delayMs)
			})
		}
		m.SetExternalMessagePump(true)
		m.SetMultiThreadedMessageLoop(false)
	} else { // Windows, Linux => LCL
		// TODO Linux Gtk3
		m.SetExternalMessagePump(false)
		m.SetMultiThreadedMessageLoop(true)
	}
}

func (m *Application) IsMainProcess() bool {
	return m.ProcessType() == types.PtBrowser
}

func Run(forms ...lcl.IEngForm) {
	if GApplication == nil || !GApplication.IsValid() {
		println("[ERROR] CEF Application Instance is not initialized")
		return
	}

	GApplication.SetLogSeverity(types.LOGSEVERITY_DISABLE)
	GApplication.SetEnablePrintPreview(true)
	if tool.IsDarwin() {
		base.AddCrDelegate()
		GApplication.InitLibLocationFromArgs()
		GApplication.SetExternalMessagePump(true)
		GApplication.SetMultiThreadedMessageLoop(false)
	}
	GApplication.SetMessageLoop()

	if GApplication.IsMainProcess() {
		mainSuccess := GApplication.StartMainProcess()
		logger.Debug("Application StartMainProcess:", mainSuccess)
		if mainSuccess {
			api.SetOnReleaseCallback(func() {
				logger.Debug("Release Callback")
				if GWorkScheduler != nil && GWorkScheduler.IsValid() {
					GWorkScheduler.Free()
				}
				GApplication.ICefApplication.Free()
			})
			engLCL.Run(forms...)
		}
	} else if tool.IsDarwin() && !GApplication.SingleProcess() && !GApplication.IsMainProcess() {
		logger.Debug("Application StartProcess 'darwin' for sub. processType:", GApplication.ProcessType())
		GApplication.StartSubProcess()
		GApplication.ICefApplication.Free()
	} else if !GApplication.IsMainProcess() {
		var startSubSuccess bool
		subProcessPath := GApplication.BrowserSubprocessPath()
		if subProcessPath != "" {
			logger.Debug("Application StartProcess for sub. processType:", GApplication.ProcessType(), "subprocessPath:", subProcessPath)
			startSubSuccess = GApplication.StartSubProcess()
		} else {
			logger.Debug("Application StartProcess for main. processType:", GApplication.ProcessType())
			startSubSuccess = GApplication.StartMainProcess()
		}
		if startSubSuccess {
			GApplication.ICefApplication.Free()
		}
	}
}
