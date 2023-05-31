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
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

// ICEFWindowParent 接口定义
type ICEFWindowParent interface {
	lcl.IWinControl
	Instance() uintptr
	Type() consts.TCefWindowHandleType         // 组件类型, Windows TCEFWindowParent 组件，MacOSX, Linux TCEFLinkedWindowParent 组件
	SetChromium(chromium IChromium, tag int32) // 设置 IChromium, 只 TCEFLinkedWindowParent 有效
	UpdateSize()                               // 更新组件大小
	HandleAllocated() bool                     // 处理所有
	CreateHandle()                             // 创建句柄
	SetOnEnter(fn lcl.TNotifyEvent)            // 进入事件
	SetOnExit(fn lcl.TNotifyEvent)             // 退出事件
	DestroyChildWindow() bool                  // 销毁子窗口
	Free()                                     // 释放
	Handle() types.HWND                        // 组件句柄
	Name() string                              // 获取组件名称
	SetName(value string)                      // 设置组件名称
	SetParent(value lcl.IWinControl)           // 设置控件父容器
	RevertCustomAnchors()                      // 恢复到自定义四角锚点定位
	DefaultAnchors()                           // 恢复到默认四角锚点定位
	Align() types.TAlign                       // 获取控件自动调整
	SetAlign(value types.TAlign)               // 设置控件自动调整
	Anchors() types.TAnchors                   // 获取四个角位置的锚点
	SetAnchors(value types.TAnchors)           // 设置四个角位置的锚点
	Visible() bool                             // 获取控件可视
	SetVisible(value bool)                     // 设置控件可视
	Enabled() bool                             // 获取是否启用
	SetEnabled(value bool)                     // 设置是否启用
	Left() int32                               // 获取左边距
	SetLeft(value int32)                       // 设置左边距
	Top() int32                                // 获取上边距
	SetTop(value int32)                        // 设置上边距
	Width() int32                              // 获取宽度
	SetWidth(value int32)                      // 设置宽度
	Height() int32                             // 获取高度
	SetHeight(value int32)                     // 设置高度
	BoundsRect() (result types.TRect)          // 获取矩形边界
	SetBoundsRect(value types.TRect)           // 设置矩形边界
	point() (x, y int32)                       // 坐标点
	size() (w, h int32)                        // 大小
}

// NewCEFWindow 创建CEFWindowParent
//
// # Windows return TCEFWindowParent
//
// MacOSX, Linux return TCEFLinkedWindowParent
func NewCEFWindow(owner lcl.IComponent) ICEFWindowParent {
	if common.IsWindows() {
		return NewCEFWindowParent(owner)
	} else {
		return NewCEFLinkedWindowParent(owner)
	}
}
