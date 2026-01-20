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

type AppearanceNamed string

const (
	// NSAppearanceNameAqua - 标准浅色系统外观
	NSAppearanceNameAqua AppearanceNamed = "NSAppearanceNameAqua"
	// NSAppearanceNameDarkAqua - 标准深色系统外观
	NSAppearanceNameDarkAqua AppearanceNamed = "NSAppearanceNameDarkAqua"
	// NSAppearanceNameVibrantLight - 浅色生动外观
	NSAppearanceNameVibrantLight AppearanceNamed = "NSAppearanceNameVibrantLight"
	// NSAppearanceNameAccessibilityHighContrastAqua - 标准浅色系统外观的高对比度版本
	NSAppearanceNameAccessibilityHighContrastAqua AppearanceNamed = "NSAppearanceNameAccessibilityHighContrastAqua"
	// NSAppearanceNameAccessibilityHighContrastDarkAqua - 标准深色系统外观的高对比度版本
	NSAppearanceNameAccessibilityHighContrastDarkAqua AppearanceNamed = "NSAppearanceNameAccessibilityHighContrastDarkAqua"
	// NSAppearanceNameAccessibilityHighContrastVibrantLight - 浅色生动外观的高对比度版本
	NSAppearanceNameAccessibilityHighContrastVibrantLight AppearanceNamed = "NSAppearanceNameAccessibilityHighContrastVibrantLight"
	// NSAppearanceNameAccessibilityHighContrastVibrantDark - 深色生动外观的高对比度版本
	NSAppearanceNameAccessibilityHighContrastVibrantDark AppearanceNamed = "NSAppearanceNameAccessibilityHighContrastVibrantDark"
)

type ToolBar struct {
	ShowSeparator bool
}

type MacOS struct {
	AppearanceNamed  AppearanceNamed // 外观
	ToolBar          *ToolBar        // 自定义工具栏
	WindowRadius     float32         // 设置窗口四角圆角, Frameless = true 时有效
	FullSizeContent  bool            // 窗口内容填充整个窗口
	TitleTransparent bool            // 标题栏透明
	TitleHideText    bool            // 隐藏标题栏标题文本
}
