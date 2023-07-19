//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//  动态链接库内置到Go字节码方式
//  如果启用该方式则根据存放目录类型生成动态库并加载
//  默认存放在系统临时目录
//  通过编译命令 `-tags` 参数指定开启该方式 `go build -tags="tempdll"`

package tempdll

// TempDLL
//  通过编译命令 `-tags` 参数控制该变量的初始化 `go build -tags="tempdll"`
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
	dllSaveDirType TempDllDIR
	dllSaveDir     string
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
//	设置动态库保存目录, 默认为空
func (m *temdll) DllSaveDir() string {
	if m == nil {
		return ""
	}
	return m.dllSaveDir
}
