//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package exception
//
//	Underlying dynamic link library exception capture
//	Supports: Windows, MacOS.
package exception

import "github.com/cyber-xxm/energy/v2/cef/internal/exception"

// SetOnException 设置 liblcl -> CEF 低层异常捕获回调函数
//
//	Supports: Windows, MacOS.
func SetOnException(fn exception.Callback) {
	exception.HandlerInit(fn)
}
