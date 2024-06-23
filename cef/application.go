//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/cef/cef"
	"github.com/energye/cef/types"
	"github.com/energye/lcl/tools"
	"github.com/energye/lcl/tools/conv"
	"github.com/energye/lcl/tools/exec"
	"os"
	"path/filepath"
	"strings"
)

// *********************************************************
// ********************** ATTENTION ! **********************
// *********************************************************
// **                                                     **
// **  MANY OF THE EVENTS IN CEF4DELPHI COMPONENTS LIKE   **
// **  TCHROMIUM, TFMXCHROMIUM OR TApplication ARE     **
// **  EXECUTED IN A CEF THREAD BY DEFAULT.               **
// **                                                     **
// **  WINDOWS CONTROLS MUST BE CREATED AND DESTROYED IN  **
// **  THE SAME THREAD TO AVOID ERRORS.                   **
// **  SOME OF THEM RECREATE THE HANDLERS IF THEY ARE     **
// **  MODIFIED AND CAN CAUSE THE SAME ERRORS.            **
// **                                                     **
// **  DON'T CREATE, MODIFY OR DESTROY WINDOWS CONTROLS   **
// **  INSIDE THE CEF4DELPHI EVENTS AND USE               **
// **  SYNCHRONIZATION OBJECTS TO PROTECT VARIABLES AND   **
// **  FIELDS IF THEY ARE ALSO USED IN THE MAIN THREAD.   **
// **                                                     **
// **  READ THIS FOR MORE INFORMATION :                   **
// **  https://www.briskbard.com/index.php?pageid=cef     **
// **                                                     **
// **  USE OUR FORUMS FOR MORE QUESTIONS :                **
// **  https://www.briskbard.com/forum/                   **
// **                                                     **
// *********************************************************
// *********************************************************

// IApplication
//
//	自定义Application
type IApplication interface {
	cef.ICefApplication
	// 注册默认事件
	registerDefaultEvent()
	// 初始化默认设置
	initDefaultSettings()
}

type TApplication struct {
	cef.TCefApplication
	ui              cef.UITool
	specificVersion cef.SpecificVersion // 特定版本 default -1
}

// AsApplication Convert a pointer object to an existing class object
func AsApplication(obj interface{}) IApplication {
	instance := cef.GetInstance(obj)
	if instance == nil {
		return nil
	}
	application := new(TApplication)
	application.specificVersion = cef.SvINVALID
	application.ui = cef.UitInvalid
	cef.SetObjectInstance(application, instance)
	return application
}

// NewApplication 创建CEF应用
//
//	初始化CEF时, 创建全局 Application 实例，多进程模式每个 application 配置都应该相同
//	disableRegisDefaultEvent = true 时不会注册默认事件
func NewApplication(disableRegisDefaultEvent ...bool) cef.ICefApplication {
	if cef.GlobalCEFApp() == nil {
		app := cef.NewCefApplication()
		customApp := AsApplication(app)
		cef.SetGlobalCEFApp(app)
		if len(disableRegisDefaultEvent) == 0 || !disableRegisDefaultEvent[0] {
			customApp.registerDefaultEvent()
		}
		customApp.initDefaultSettings()
	}
	return cef.GlobalCEFApp()
}

// SpecificVersion
//
//	energy 对CEF一些版本做为特定版本，支持对旧的CEF所支持的系统或功能
func (m *TApplication) SpecificVersion() cef.SpecificVersion {
	if m.specificVersion == cef.SvINVALID {
		ver := cef.GetSpecificVersion()
		switch ver {
		case cef.SvCEF, cef.SvCEF49, cef.SvCEF87, cef.SvCEF106, cef.SvCEF109:
			m.specificVersion = ver
		}
	}
	return m.specificVersion
}

// IsNotSpecVer 非针特定本，当前版本或当前最新版本
func (m *TApplication) IsNotSpecVer() bool {
	return m.SpecificVersion() == cef.SvCEF
}

// IsSpecVer49 特定 WindowsXP
func (m *TApplication) IsSpecVer49() bool {
	return m.SpecificVersion() == cef.SvCEF49
}

// IsSpecVer87 特定 Flash
func (m *TApplication) IsSpecVer87() bool {
	return m.SpecificVersion() == cef.SvCEF87
}

// IsSpecVer106 特定 Linux GTK2
func (m *TApplication) IsSpecVer106() bool {
	return m.SpecificVersion() == cef.SvCEF106
}

// IsSpecVer109 特定 7, 8/8.1 and Windows Server 2012
func (m *TApplication) IsSpecVer109() bool {
	return m.SpecificVersion() == cef.SvCEF109
}

// registerDefaultEvent 注册默认事件
func (m *TApplication) registerDefaultEvent() {
	m.defaultSetOnContextCreated()
	//m.defaultSetOnProcessMessageReceived()
	//m.defaultSetOnWebKitInitialized()
	//m.defaultSetOnRegCustomSchemes()
	//m.defaultSetOnRenderLoadStart()
}

func (m *TApplication) defaultSetOnContextCreated() {
	//m.SetOnContextCreated(func(browse ICefBrowser, frame ICefFrame, context ICefv8Context) {
	//	var flag bool
	//	if m.onContextCreated != nil {
	//		flag = m.onContextCreated(browse, frame, context)
	//	}
	//	if !flag {
	//		appOnContextCreated(browse, frame, context)
	//	}
	//})
}

// initDefaultSettings 初始 energy 默认设置
func (m *TApplication) initDefaultSettings() {
	if tools.IsWindows() {
		m.ui = cef.UitWin32
	} else if tools.IsDarwin() {
		m.ui = cef.UitCocoa
	} else if tools.IsLinux() {
		cefVersion := strings.Split(m.LibCefVersion(), ".")
		if len(cefVersion) > 0 {
			major := conv.StrToInt32(cefVersion[0])
			// cef version <= 106.1.1 default use gtk2
			if major <= 106 {
				m.ui = cef.UitGtk2
			} else {
				// cef version > 106.1.1 default use gtk3
				m.ui = cef.UitGtk3
			}
		} else {
			// default use gtk3
			m.ui = cef.UitGtk3
		}
	} else {
		panic("Unsupported system, currently only supports Windows, Mac OS, and Linux")
	}
	if m.FrameworkDirPath() == "" {
		// 默认CEF框架目录
		// 当前执行文件所在目录或ENERGY_HOME环境配置目录
		lp := FrameworkDir()
		if lp != "" {
			m.SetFrameworkDirPath(lp)
		}
	}

	m.SetLocale(cef.LANGUAGE_zh_CN)
	m.SetLogSeverity(cef.LOGSEVERITY_DISABLE)
	m.SetEnablePrintPreview(true)
	//m.SetEnableGPU(true) 默认还是关闭GPU加速
	// 以下条件判断根据不同平台, 启动不同的窗口组件
	// ViewsFrameworkBrowserWindow 简称(VF)窗口组件, 同时支持 Windows/Linux/MacOSX
	// LCL 窗口组件,同时支持 Windows/MacOSX, CEF版本<=106.xx时支持GTK2, CEF版本 >= 107.xx时默认开启 GTK3 且不支持 GTK2 和 LCL提供的各种组件
	if tools.IsLinux() { // Linux => (VF)View Framework 窗口
		if m.IsUIGtk3() {
			// Linux CEF >= 107.xxx 版本以后，默认启用的GTK3，106及以前版本默认支持GTK2但无法正常输入中文
			// Linux 默认设置为false,将启用 ViewsFrameworkBrowserWindow 窗口
			m.SetExternalMessagePump(false)
			m.SetMultiThreadedMessageLoop(false)
		} else if m.IsUIGtk2() {
			// GTK2 默认支持LCL,但还未解决无法输入中文问题
			m.SetExternalMessagePump(false)
			m.SetMultiThreadedMessageLoop(true)
		}
		// 这是一个解决“GPU不可用错误”问题的方法 linux
		// https://bitbucket.org/chromiumembedded/cef/issues/2964/gpu-is-not-usable-error-during-cef
		m.SetDisableZygote(true)
	} else if tools.IsDarwin() { // Darwin => LCL窗口
		cef.AddCrDelegate()
		cef.GlobalWorkSchedulerCreate(nil)
		m.SetOnScheduleMessagePumpWork(nil)
		// MacOSX 在使用LCL窗口组件必须将 ExternalMessagePump=true 和 MultiThreadedMessageLoop=false
		// 或
		// 同 Linux 一样使用 ViewsFrameworkBrowserWindow 窗口组件
		m.SetExternalMessagePump(true)
		m.SetMultiThreadedMessageLoop(false)
	} else { // Windows => LCL窗口
		m.SetExternalMessagePump(false)
		m.SetMultiThreadedMessageLoop(true)
	}
}

func (m *TApplication) IsUIWin32() bool {
	return m.ui == cef.UitWin32
}

func (m *TApplication) IsUICocoa() bool {
	return m.ui == cef.UitCocoa
}

func (m *TApplication) IsUIGtk2() bool {
	return m.ui == cef.UitGtk2
}

func (m *TApplication) IsUIGtk3() bool {
	return m.ui == cef.UitGtk3
}

func libCef() string {
	if tools.IsWindows() {
		return "libcef.dll"
	} else if tools.IsLinux() {
		return "libcef.so"
	}
	return ""
}

// FrameworkDir
//
//	返回CEF框架目录, 以当前执行文件所在目录开始查找
//	如果当前执行文件目录未找到，再从ENERGY_HOME环境变量查找
//	Darwin 平台除外
func FrameworkDir() string {
	var lib = libCef() // 根据CEF libcef.xx 动态库
	if lib != "" {
		//当前目录
		if tools.IsExist(filepath.Join(exec.Dir, lib)) {
			return exec.Dir
		}
		//环境变量
		var env = os.Getenv(types.ENERGY_HOME_KEY)
		if tools.IsExist(filepath.Join(env, lib)) {
			return env
		}
	}
	return ""
}
