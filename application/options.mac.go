//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package application

import (
	. "github.com/energye/energy/v3/pkgs/darwin/types"
)

type ToolBar struct {
	ShowSeparator bool
}

type MacOS struct {
	AppearanceName    AppearanceName // 外观
	ToolBar           *ToolBar       // 自定义工具栏
	WindowRadius      float32        // 设置窗口四角圆角, Frameless = true 时有效
	FullSizeContent   bool           // 窗口内容填充整个窗口
	TitleTransparent  bool           // 标题栏透明
	TitleHideText     bool           // 隐藏标题栏标题文本
	UseWindowDelegate bool           // 使用自定义 window 代理
}
