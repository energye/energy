//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEFWindowParent 组件

package cef

import (
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
)

// ICEFWindowParent 接口定义
//
//	Windows return TCEFWindowParent
//	MacOSX, Linux return TCEFLinkedWindowParent
type ICEFWindowParent interface {
	lcl.IWinControl
	Type() consts.TCefWindowHandleType         // 组件类型, Windows TCEFWindowParent 组件，MacOSX, Linux TCEFLinkedWindowParent 组件
	SetChromium(chromium IChromium, tag int32) // 设置 IChromium, 只 TCEFLinkedWindowParent 有效
	UpdateSize()                               // 更新组件大小
	CreateHandle()                             // 创建句柄
	SetOnEnter(fn lcl.TNotifyEvent)            // 进入事件
	SetOnExit(fn lcl.TNotifyEvent)             // 退出事件
	DestroyChildWindow() bool                  // 销毁子窗口
	RevertCustomAnchors()                      // 恢复到自定义四角锚点定位
	DefaultAnchors()                           // 恢复到默认四角锚点定位
	point() (x, y int32)                       // 坐标点
	size() (w, h int32)                        // 大小
}
