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
	"os"
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
	postMessage              *tPostMessage
	onContextCreated         cef.TOnContextCreatedEvent
	onProcessMessageReceived cef.TOnProcessMessageReceivedEvent
	onRegisterCustomSchemes  cef.TOnRegisterCustomSchemesEvent
	onContextInitialized     cef.TOnContextInitializedEvent
	viewsWindows             []IViewsWindow
}

// IRunWindow is the accepted Run target type.
//
// Supported values are native LCL forms (lcl.IEngForm) and CEF views framework
// windows (objects implementing CreateTopLevelWindow).
type IRunWindow = any

type IViewsWindow interface {
	CreateTopLevelWindow()
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
		GApplication.initDefaultEvent()
		GApplication.SetLogSeverity(types.LOGSEVERITY_DISABLE)
		GApplication.SetEnablePrintPreview(true)
		//GApplication.SetAllowFileAccessFromFiles(true)
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
func (m *Application) messageLoop(kind browserKind) {
	if kind == bkEmbedded {
		if tool.IsDarwin() { // Darwin embed => LCL
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
		} else if tool.IsWindows() { //  Window embed >= LCL
			m.SetExternalMessagePump(false)
			m.SetMultiThreadedMessageLoop(true)
		} else if tool.IsLinux() { // Linux default GTK3 => views framework
			if api.Widget().IsGTK2() {
				// GTK2 => native LCL
				println("[ERROR] Linux CEF use GTK3")
				os.Exit(1)
			} else if api.Widget().IsGTK3() {
				m.SetExternalMessagePump(false)
				m.SetMultiThreadedMessageLoop(true)
				// Solution for "GPU unavailable error" on Linux.
				// https://bitbucket.org/chromiumembedded/cef/issues/2964/gpu-is-not-usable-error-during-cef
				m.SetDisableZygote(true)
			}
		}
	} else if kind == bkViews {
		m.SetExternalMessagePump(false)
		m.SetMultiThreadedMessageLoop(false)
		// Solution for "GPU unavailable error" on Linux.
		// https://bitbucket.org/chromiumembedded/cef/issues/2964/gpu-is-not-usable-error-during-cef
		m.SetDisableZygote(true)
	}
}

func (m *Application) IsMainProcess() bool {
	return m.ProcessType() == types.PtBrowser
}

// Run runs the application and starts the message loop
//
// Launches CEF application and selects window mode based on window instance type
// Uses embedding when window implements lcl.IEngForm
// Uses CEF Views Framework when window implements IViewsWindow
func Run(windows ...IRunWindow) {
	if GApplication == nil || !GApplication.IsValid() {
		println("[ERROR] CEF Application Instance is not initialized")
		return
	}

	processTypeStr := ProcessType(GApplication.ProcessType())

	if GApplication.IsMainProcess() {
		var kind browserKind
		embedWindows := make([]lcl.IEngForm, 0, len(windows))
		viewsWindows := make([]IViewsWindow, 0, len(windows))
		for _, window := range windows {
			if w, ok := window.(lcl.IEngForm); ok {
				embedWindows = append(embedWindows, w)
				continue
			}
			if w, ok := window.(IViewsWindow); ok {
				viewsWindows = append(viewsWindows, w)
				continue
			}
			logger.Debug("Application Run unsupported window:", fmt.Sprintf("%T", window))
		}
		if len(embedWindows) > 0 {
			kind = bkEmbedded
		} else if len(viewsWindows) > 0 {
			kind = bkViews
			GApplication.viewsWindows = viewsWindows
		}
		// selects window mode based on window instance type
		GApplication.messageLoop(kind)
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
			if kind == bkEmbedded {
				engLCL.Run(embedWindows...)
			} else if kind == bkViews {
				GApplication.RunMessageLoop()
			}
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
