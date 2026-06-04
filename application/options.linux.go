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

type Linux struct {
	HardwareGPU THardwareGPU
	Theme       Theme // 跟随系统主题设置, SystemDefault, Dark, Light
	// TODO add titlebar
}
