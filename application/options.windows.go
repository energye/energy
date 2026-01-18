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

type Theme uintptr

const (
	SystemDefault Theme = iota // 跟随系统主题设置
	Dark                       // 强制使用深色主题
	Light                      // 强制使用浅色主题
)

type TBackdropType int32

const (
	BtAuto    TBackdropType = iota // 自动选择最合适的背景材质
	BtNone                         // 无特殊背景效果，纯色或普通背景
	BtMica                         // 色调的轻度模糊背景
	BtAcrylic                      // Acrylic（亚克力）材质：高斯模糊 + 噪点纹理
	BtTabbed                       // Tabbed Acrylic：专为标签页设计的 Acrylic 变体
)

type ThemeSetting struct {
	DarkTitleBar           int32
	DarkTitleBarInactive   int32
	DarkTitleText          int32
	DarkTitleTextInactive  int32
	DarkBorder             int32
	DarkBorderInactive     int32
	LightTitleBar          int32
	LightTitleBarInactive  int32
	LightTitleText         int32
	LightTitleTextInactive int32
	LightBorder            int32
	LightBorderInactive    int32
}
