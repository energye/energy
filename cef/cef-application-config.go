//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/golcl/lcl/api"
)

// Application 支持的配置
type tCefApplicationConfig struct {
	frameworkDirPath     uintptr //string cef框架根目录
	resourcesDirPath     uintptr //string
	localesDirPath       uintptr //string
	cache                uintptr //string
	userDataPath         uintptr //string
	language             uintptr //string 语言设置
	localesRequired      uintptr //string 默认空,检查所有的语言环境 逗号分隔
	logFile              uintptr //string
	mainBundlePath       uintptr //string 只对 darwin 启作用
	browseSubprocessPath uintptr //string 只对 非darwin 启作用
	logSeverity          uintptr //uint32
	noSandbox            uintptr //bool
	disableZygote        uintptr //bool 只对 linux 启作用
	enableGPU            uintptr //bool
	singleProcess        uintptr //bool 进程启动模式,默认false true:单进程 false:多进程
	useMockKeyChain      uintptr //bool
	checkCEFFiles        uintptr //bool
	remoteDebuggingPort  uintptr //int32
}

//创建应用全局配置
func NewApplicationConfig() *tCefApplicationConfig {
	m := &tCefApplicationConfig{}
	m.SetFrameworkDirPath(empty)
	m.SetResourcesDirPath(empty)
	m.SetLocalesDirPath(empty)
	m.SetCache(empty)
	m.SetUserDataPath(empty)
	m.SetLanguage(LANGUAGE_zh_CN)
	m.SetLocalesRequired(empty)
	m.SetLogFile(empty)
	m.SetMainBundlePath(empty)
	m.SetBrowseSubprocessPath(empty)
	m.SetLogSeverity(LOGSEVERITY_DISABLE)
	m.SetEnableGPU(enableGPU)
	m.SetSingleProcess(false)
	m.SetUseMockKeyChain(false)
	m.SetNoSandbox(true)
	m.SetDisableZygote(true)
	m.SetCheckCEFFiles(false)
	m.SetRemoteDebuggingPort(0)
	return m
}

//设置 Chromium Framework 编译好的二进制包根目录，默认当前目录
func (m *tCefApplicationConfig) SetFrameworkDirPath(s string) *tCefApplicationConfig {
	m.frameworkDirPath = api.GoStrToDStr(s)
	return m
}

//设置资源目录，默认当前目录
func (m *tCefApplicationConfig) SetResourcesDirPath(s string) *tCefApplicationConfig {
	m.resourcesDirPath = api.GoStrToDStr(s)
	return m
}

//设置本地语言目录，默认当前目录
func (m *tCefApplicationConfig) SetLocalesDirPath(s string) *tCefApplicationConfig {
	m.localesDirPath = api.GoStrToDStr(s)
	return m
}

//设置缓存目录，默认当前目录
func (m *tCefApplicationConfig) SetCache(s string) *tCefApplicationConfig {
	m.cache = api.GoStrToDStr(s)
	return m
}

//设置用户数据目录，默认当前目录
func (m *tCefApplicationConfig) SetUserDataPath(s string) *tCefApplicationConfig {
	m.userDataPath = api.GoStrToDStr(s)
	return m
}

//设置进程模型，作用于linux-默认禁用
func (m *tCefApplicationConfig) SetDisableZygote(s bool) *tCefApplicationConfig {
	m.disableZygote = api.GoBoolToDBool(s)
	return m
}

//设置关闭沙盒-默认关闭
func (m *tCefApplicationConfig) SetNoSandbox(s bool) *tCefApplicationConfig {
	m.noSandbox = api.GoBoolToDBool(s)
	return m
}

//设置开启关闭GPU加速
func (m *tCefApplicationConfig) SetEnableGPU(s bool) *tCefApplicationConfig {
	enableGPU = s
	m.enableGPU = api.GoBoolToDBool(s)
	return m
}

//设置进程模式，true:单进程模式
func (m *tCefApplicationConfig) SetSingleProcess(s bool) *tCefApplicationConfig {
	SingleProcess = s
	m.singleProcess = api.GoBoolToDBool(s)
	return m
}

//设置使用模拟key chain
func (m *tCefApplicationConfig) SetUseMockKeyChain(s bool) *tCefApplicationConfig {
	m.useMockKeyChain = api.GoBoolToDBool(s)
	return m
}

//检测CEF文件默认不检测
func (m *tCefApplicationConfig) SetCheckCEFFiles(s bool) *tCefApplicationConfig {
	m.checkCEFFiles = api.GoBoolToDBool(s)
	return m
}

//设置语言
func (m *tCefApplicationConfig) SetLanguage(s LANGUAGE) *tCefApplicationConfig {
	m.language = api.GoStrToDStr(string(s))
	return m
}

//设置必备的本地语言支持，逗号分隔的字符串 s="zh-CN,en-US" ,默认情况下 en-US 是必须的
func (m *tCefApplicationConfig) SetLocalesRequired(s string) *tCefApplicationConfig {
	m.localesRequired = api.GoStrToDStr(s)
	return m
}

//设置日志文件目录
func (m *tCefApplicationConfig) SetLogFile(s string) *tCefApplicationConfig {
	m.logFile = api.GoStrToDStr(s)
	return m
}

//设置主程序绑定目录 作用于macos
func (m *tCefApplicationConfig) SetMainBundlePath(s string) *tCefApplicationConfig {
	m.mainBundlePath = api.GoStrToDStr(s)
	return m
}

//设置子进程执行文件目录，一搬用于主进程过于复杂启动慢，需要独立出子进程
func (m *tCefApplicationConfig) SetBrowseSubprocessPath(s string) *tCefApplicationConfig {
	m.browseSubprocessPath = api.GoStrToDStr(s)
	return m
}

//设置日志级别
func (m *tCefApplicationConfig) SetLogSeverity(s LOG) *tCefApplicationConfig {
	m.logSeverity = uintptr(s)
	return m
}

//设置远程调式端口 (1024 ~ 65535)
func (m *tCefApplicationConfig) SetRemoteDebuggingPort(s int32) *tCefApplicationConfig {
	m.remoteDebuggingPort = uintptr(s)
	return m
}

// SetCommonRootName
//
//通用类型所属对象名定义 默认值:gocobj
//
//默认值  gocobj
func (m *tCefApplicationConfig) SetCommonRootName(name string) {
	if name == "" {
		name = commonRootName
	} else {
		commonRootName = name
	}
	_CEFV8ValueRef_SetCommonRootName(name)
}

// SetObjectRootName
//
//对象类型所属对象名定义 默认值:goobj
//
//默认值 goobj
func (m *tCefApplicationConfig) SetObjectRootName(name string) {
	if name == "" {
		name = objectRootName
	} else {
		objectRootName = name
	}
	_CEFV8ValueRef_SetObjectRootName(name)
}
