//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//  动态链接库内置到Go执行文件
//  如果启用该方式则根据存放目录类型生成动态库并加载
//  默认将liblcl释放放在系统临时目录

package tempdll

// TempDLL
//  目录释放配置
var TempDLL *temdll

// TempDllDIR
//  DLL存放目录
type TempDllDIR int8

const (
	TddInvalid    TempDllDIR = iota - 1 // 无效
	TddTmp                              // 系统临时目录
	TddCurrent                          // 当前执行文件目录
	TddEnergyHome                       // Energy环境变量目录, 如果为空，则为系统临时目录
	TddCustom                           // 自定义目录, 如果为空，则为系统临时目录
)

type temdll struct {
	dllSaveDirType TempDllDIR // 保存类型
	dllSaveDir     string     // 保存目录
	dllFSDir       string     // dll所在fs目录, 默认libs, 格式为: you/liblcl/path, 目录不允许\\出现
}

// SetDllSaveDirType
//  设置liblcl.xx动态库自动保存目录类型，默认为系统临时目录
//	如果值为 TddCustom 则必须设置保存目录(SetDllSavePath)，否则还是以系统临时目录
func (m *temdll) SetDllSaveDirType(dllSaveDirType TempDllDIR) {
	if m == nil {
		return
	}
	m.dllSaveDirType = dllSaveDirType
}

// SetDllSaveDir
//	设置动态库保存目录, 默认为空
//  dllSaveDirType = TddCustom 时有效
func (m *temdll) SetDllSaveDir(dllSaveDir string) {
	if m == nil {
		return
	}
	m.dllSaveDir = dllSaveDir
}

// DllSaveDirType
//  返回liblcl.xx动态库自动保存目录类型，默认为系统临时目录
//	如果值为 TddCustom 则必须设置保存目录(SetDllSavePath)，否则还是以系统临时目录
func (m *temdll) DllSaveDirType() TempDllDIR {
	if m == nil {
		return TddInvalid
	}
	return m.dllSaveDirType
}

// DllSaveDir
//	返回动态库保存目录, 默认为空
func (m *temdll) DllSaveDir() string {
	if m == nil {
		return ""
	}
	return m.dllSaveDir
}

// SetDllFSDir
//
// 设置dll所在内置资源目录
// 默认libs, 格式为: you/liblcl/path, 目录不允许\\
func (m *temdll) SetDllFSDir(dllFSDir string) {
	m.dllFSDir = dllFSDir
}
