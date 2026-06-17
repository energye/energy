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
	"errors"
	"fmt"
	"github.com/energye/cef/base"
	"github.com/energye/cef/cef"
	"github.com/energye/cef/cef/types"
	"github.com/energye/cef/config"
	"github.com/energye/energy/v3/application"
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
	gProcessType   = map[types.TCefProcessType]string{
		types.PtBrowser:  "Browser",
		types.PtRenderer: "Renderer",
		types.PtZygote:   "Zygote",
		types.PtGPU:      "GPU",
		types.PtUtility:  "Utility",
		types.PtBroker:   "Broker",
		types.PtCrashpad: "Crashpad",
		types.PtOther:    "Other",
	}
)

type Application struct {
	cef.ICefApplication
	application.Application
}

// Init CEF Global initialization, invoked at application startup in main
func Init() *Application {
	lcl.Init()
	base.Init()
	return NewApplication()
}

// CheckLibRuntimeVersion Check runtime library and CEF version compatibility
func CheckLibRuntimeVersion() error {
	major, _, _, _ := cef.LibVersion()
	//major = cef.CEFVersion // TODO test
	if cef.CEFVersion != major {
		e := fmt.Sprintf("CEF version does not match the runtime library. Current CEF version: %v, Runtime library version: %v",
			cef.CEFVersion, major)
		return errors.New(e)
	}
	return nil
}

func ProcessType(pt types.TCefProcessType) string {
	return gProcessType[pt]
}

func NewApplication() *Application {
	if GApplication == nil {
		chkErr := CheckLibRuntimeVersion()
		if chkErr != nil {
			println("[ERROR]", chkErr.Error())
			return nil
		}
		GApplication = &Application{
			ICefApplication: cef.NewApplication(),
		}
		application.GApplication = &GApplication.Application
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

// messageLoop Message polling. CEF Application uses OS-specific configurations
func (m *Application) messageLoop() {
	if tool.IsDarwin() { // Darwin => LCL
		if m.IsMainProcess() {
			base.AddCrDelegate()
		}
		m.InitLibLocationFromArgs()

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

func (m *Application) IsMainThread() bool {
	return api.CurrentThreadId() == api.MainThreadId()
}

func Run(forms ...lcl.IEngForm) {
	if GApplication == nil || !GApplication.IsValid() {
		println("[ERROR] CEF Application Instance is not initialized")
		return
	}
	GApplication.initDefaultEvent()
	GApplication.messageLoop()
	GApplication.SetLogSeverity(types.LOGSEVERITY_DISABLE)
	GApplication.SetEnablePrintPreview(true)

	processTypeStr := ProcessType(GApplication.ProcessType())

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
		logger.Debug("Application StartProcess 'darwin' for sub.executable processType:", processTypeStr)
		GApplication.StartSubProcess()
		GApplication.ICefApplication.Free()
	} else if !GApplication.IsMainProcess() {
		var startSubSuccess bool
		subProcessPath := GApplication.BrowserSubprocessPath()
		if subProcessPath != "" {
			logger.Debug("Application StartProcess for sub.executable processType:", processTypeStr, "subprocessPath:", subProcessPath)
			startSubSuccess = GApplication.StartSubProcess()
		} else {
			logger.Debug("Application StartProcess for main.executable processType:", processTypeStr)
			startSubSuccess = GApplication.StartMainProcess()
		}
		if startSubSuccess {
			GApplication.ICefApplication.Free()
		}
	}
}
