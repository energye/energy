//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build linux

// Linux
// GTK2 & GTK3: Manually calling initialization
// Other: Automatic initialization and destruction

package lclwidget

import (
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
)

// CustomWidgetSetInitialization
// 自定义组件初始化 Linux GTK3
func CustomWidgetSetInitialization() {
	imports.Proc(def.Interface_CustomWidgetSetInitialization).Call()
}

// CustomWidgetSetFinalization
//  自定义组件销毁 Linux GTK3
func CustomWidgetSetFinalization() {
	imports.Proc(def.Interface_CustomWidgetSetFinalization).Call()
}
