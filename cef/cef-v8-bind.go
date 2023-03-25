//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 基于IPC的字段数据绑定
package cef

import "github.com/energye/energy/common"

const (
	internalBind   = "energy"
	internalV8Bind = "v8"
)

// isInternalBind 内部使用字段不能使用
func isInternalBind(name string) bool {
	return name == internalBind
}

// bindInit 初始化
func bindInit() {
	isSingleProcess := application.SingleProcess()
	if isSingleProcess {
		bindRender = &bindRenderProcess{}
		bindBrowser = &bindBrowserProcess{}
	} else {
		if common.Args.IsMain() {
			bindBrowser = &bindBrowserProcess{}
		} else {
			bindRender = &bindRenderProcess{}
		}
	}
}
