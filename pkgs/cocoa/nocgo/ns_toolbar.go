//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	"github.com/ebitengine/purego/objc"
	. "github.com/energye/energy/v3/pkgs/cocoa/types"
)

type NSToolBar struct {
	NSObject
}

// NewToolBar 为窗口创建并配置工具栏
//
// 参数:
//   - window: 目标 NSWindow 对象
//   - delegate: 工具栏代理对象（通常与窗口代理相同），负责提供工具栏项
//   - config: 工具栏配置选项
func NewToolBar(window *NSWindow, delegate *NSWindowDelegate, config ToolbarConfiguration) {
	if window == nil {
		return
	}
	nsWindow := objc.ID(window.Instance())

	toolbarClass := objc.GetClass("NSToolbar")
	toolbar := objc.ID(toolbarClass).Send(objc.RegisterName("alloc"))
	toolbar = toolbar.Send(objc.RegisterName("initWithIdentifier:"), "ENERGY.ToolBar")

	// 设置是否显示基线分隔符
	showSep := uintptr(0)
	if config.ShowSeparator {
		showSep = 1
	}
	toolbar.Send(objc.RegisterName("setShowsBaselineSeparator:"), showSep)

	// 设置代理
	if delegate != nil {
		toolbar.Send(objc.RegisterName("setDelegate:"), delegate.Self())
	}

	// 将工具栏设置到窗口
	nsWindow.Send(objc.RegisterName("setToolbar:"), toolbar)
}
