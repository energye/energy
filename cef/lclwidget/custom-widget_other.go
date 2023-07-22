//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !linux

// GTK2 & GTK3: Manually calling initialization
// Other: Automatic initialization and destruction

package lclwidget

// CustomWidgetSetInitialization
// 自定义组件初始化 other
func CustomWidgetSetInitialization() {
}

// CustomWidgetSetFinalization
//  自定义组件销毁 other
func CustomWidgetSetFinalization() {
}
