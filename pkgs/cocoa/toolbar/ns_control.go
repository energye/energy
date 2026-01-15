//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin

package toolbar

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa
#include "config.h"
*/
import "C"

type IControl interface {
	IView
	Owner() *NSToolBar
	Property() *ControlProperty
}

type Control struct {
	item     ItemBase
	owner    *NSToolBar
	instance Pointer
	property *ControlProperty
}

func (m *Control) Identifier() string {
	return m.item.Identifier
}

func (m *Control) Instance() Pointer {
	return m.instance
}

func (m *Control) Owner() *NSToolBar {
	return m.owner
}

func (m *Control) Property() *ControlProperty {
	return m.property
}

func (m *Control) SetBindControlObjectIdentifier() {
	var cID *C.char
	cID = C.CString(m.Identifier())
	defer C.free(Pointer(cID))
	C.SetBindControlObjectIdentifier(m.instance, cID)
}

// SetEnable 设置控件启用状态
func (m *Control) SetEnable(v bool) {
	C.SetControlEnable(m.instance, C.BOOL(v))
}

// Enable 获取控件启用状态
func (m *Control) Enable() bool {
	return bool(C.GetControlEnable(m.instance))
}

// SetHidden 设置控件隐藏状态
func (m *Control) SetHidden(v bool) {
	C.SetControlHidden(m.instance, C.BOOL(v))
}

// Hidden 获取控件隐藏状态
func (m *Control) Hidden() bool {
	return bool(C.GetControlHidden(m.instance))
}

// SetAlpha 设置控件透明度
func (m *Control) SetAlpha(alpha float64) {
	C.SetControlAlphaValue(m.instance, C.CGFloat(alpha))
}

// Alpha 获取控件透明度
func (m *Control) Alpha() float64 {
	return float64(C.GetControlAlphaValue(m.instance))
}

// FadeIn 淡入动画效果
func (m *Control) FadeIn(duration float64) {
	// 确保控件可见
	m.SetHidden(false)

	// 使用 Go 的 goroutine 模拟动画效果
	// 注意：在实际应用中，你可能需要使用更复杂的动画实现
	go func() {
		steps := int(duration * 60) // 假设 60 FPS
		for i := 0; i <= steps; i++ {
			alpha := float64(i) / float64(steps)
			m.SetAlpha(alpha)
			// 这里应该使用更精确的时间控制
			// time.Sleep(time.Duration(duration*1000/float64(steps)) * time.Millisecond)
		}
	}()
}

// FadeOut 淡出动画效果
func (m *Control) FadeOut(duration float64) {
	// 使用 Go 的 goroutine 模拟动画效果
	go func() {
		steps := int(duration * 60) // 假设 60 FPS
		for i := steps; i >= 0; i-- {
			alpha := float64(i) / float64(steps)
			m.SetAlpha(alpha)
			// 这里应该使用更精确的时间控制
			// time.Sleep(time.Duration(duration*1000/float64(steps)) * time.Millisecond)
		}
		// 动画完成后隐藏控件
		m.SetHidden(true)
		// 恢复透明度以便下次显示
		m.SetAlpha(1.0)
	}()
}

func (m *Control) SetFocus(v bool) bool {
	return bool(C.SetControlFocus(m.instance, C.BOOL(v)))
}
