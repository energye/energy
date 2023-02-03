//go:build !windows
// +build !windows

package traydemo

import (
	"github.com/energye/energy/cef"
)

// 仅适用windows
//
// LCL + [CEF] 托盘 只适用 windows 基于html 和 ipc 实现功能
//
//推荐在windows或macosx中使用
func LCLCefTrayDemo(browserWindow cef.IBrowserWindow) {

}

// 仅适用windows
//
// LCL + [VF] 托盘 只适用 windows 基于html 和 ipc 实现功能
//
// VF组件托盘，无法使用LCL相关组件
func LCLVFTrayDemo(browserWindow cef.IBrowserWindow) {

}
