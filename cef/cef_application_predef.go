//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/v2/api"
	"github.com/energye/energy/v2/common"
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

var globalApp ICefApplication

type IApplication interface {
	ICefApplication
	registerDefaultEvent()
	initDefaultSettings()
}

type TApplication struct {
	TCefApplication
	ui              UITool
	specificVersion SpecificVersion // 特定版本 default -1
}

// SetGlobalCEFApp
//
//	设置全局 ICefApplication 实例，单例模式, 应在 NewCefApplication 后设置
func SetGlobalCEFApp(application ICefApplication) {
	api.CEFPreDef().SysCallN(12, application.Instance())
	globalApp = application
}

// DestroyGlobalCEFApp 销毁全局App实例
func DestroyGlobalCEFApp() {
	if globalApp != nil {
		api.CEFPreDef().SysCallN(13)
		globalApp.SetInstance(nil)
		globalApp = nil
	}
}

// NewApplication 创建CEF应用
//
//	初始化CEF时, 创建全局 Application 实例，多进程模式每个 application 配置都应该相同
//	disableRegisDefaultEvent = true 时不会注册默认事件
func NewApplication(disableRegisDefaultEvent ...bool) ICefApplication {
	if globalApp == nil {
		app := NewCefApplication()
		customApp := AsApplication(app)
		SetGlobalCEFApp(app)
		if len(disableRegisDefaultEvent) == 0 || !disableRegisDefaultEvent[0] {
			customApp.registerDefaultEvent()
		}
		customApp.initDefaultSettings()
	}
	return globalApp
}

// SpecificVersion
//
//	energy 对CEF一些版本做为特定版本，支持对旧的CEF所支持的系统或功能
func (m *TApplication) SpecificVersion() SpecificVersion {
	if m.specificVersion == SvINVALID {
		r1 := api.CEFPreDef().SysCallN(14)
		switch SpecificVersion(r1) {
		case SvCEF, SvCEF49, SvCEF87, SvCEF106, SvCEF109:
			m.specificVersion = SpecificVersion(r1)
		}
	}
	return m.specificVersion
}

// IsNotSpecVer 非针特定本，当前版本或当前最新版本
func (m *TApplication) IsNotSpecVer() bool {
	return m.SpecificVersion() == SvCEF
}

// IsSpecVer49 特定 WindowsXP
func (m *TApplication) IsSpecVer49() bool {
	return m.SpecificVersion() == SvCEF49
}

// IsSpecVer87 特定 Flash
func (m *TApplication) IsSpecVer87() bool {
	return m.SpecificVersion() == SvCEF87
}

// IsSpecVer106 特定 Linux GTK2
func (m *TApplication) IsSpecVer106() bool {
	return m.SpecificVersion() == SvCEF106
}

// IsSpecVer109 特定 7, 8/8.1 and Windows Server 2012
func (m *TApplication) IsSpecVer109() bool {
	return m.SpecificVersion() == SvCEF109
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
	if common.IsWindows() {
		m.ui = UitWin32
	} else if common.IsDarwin() {
		m.ui = UitCocoa
	} else if common.IsLinux() {
		cefVersion := strings.Split(m.LibCefVersion(), ".")
		if len(cefVersion) > 0 {
			major := common.StrToInt32(cefVersion[0])
			// cef version <= 106.1.1 default use gtk2
			if major <= 106 {
				m.ui = UitGtk2
			} else {
				// cef version > 106.1.1 default use gtk3
				m.ui = UitGtk3
			}
		} else {
			// default use gtk3
			m.ui = UitGtk3
		}
	} else {
		panic("Unsupported system, currently only supports Windows, Mac OS, and Linux")
	}
	if m.FrameworkDirPath() == "" {
		// 默认CEF框架目录
		// 当前执行文件所在目录或ENERGY_HOME环境配置目录
		lp := common.FrameworkDir()
		if lp != "" {
			m.SetFrameworkDirPath(lp)
		}
	}

	m.SetLocale(LANGUAGE_zh_CN)
	m.SetLogSeverity(LOGSEVERITY_DISABLE)
	m.SetEnablePrintPreview(true)
	//m.SetEnableGPU(true) 默认还是关闭GPU加速
	// 以下条件判断根据不同平台, 启动不同的窗口组件
	// ViewsFrameworkBrowserWindow 简称(VF)窗口组件, 同时支持 Windows/Linux/MacOSX
	// LCL 窗口组件,同时支持 Windows/MacOSX, CEF版本<=106.xx时支持GTK2, CEF版本 >= 107.xx时默认开启 GTK3 且不支持 GTK2 和 LCL提供的各种组件
	if common.IsLinux() { // Linux => (VF)View Framework 窗口
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
	} else if common.IsDarwin() { // Darwin => LCL窗口
		AddCrDelegate()
		GlobalWorkSchedulerCreate(nil)
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
	return m.ui == UitWin32
}

func (m *TApplication) IsUICocoa() bool {
	return m.ui == UitCocoa
}

func (m *TApplication) IsUIGtk2() bool {
	return m.ui == UitGtk2
}

func (m *TApplication) IsUIGtk3() bool {
	return m.ui == UitGtk3
}
