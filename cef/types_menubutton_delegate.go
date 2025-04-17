//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefMenuButtonDelegate
// include/capi/views/cef_menu_button_delegate_capi.h (cef_menu_button_delegate_t)
type ICefMenuButtonDelegate struct {
	*ICefButtonDelegate
}

// MenuButtonDelegateRef -> ICefMenuModelDelegate
var MenuButtonDelegateRef menuButtonDelegate

type menuButtonDelegate uintptr

func (*menuButtonDelegate) New() *ICefMenuButtonDelegate {
	var result uintptr
	imports.Proc(def.MenuButtonDelegateRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMenuButtonDelegate{&ICefButtonDelegate{&ICefViewDelegate{
			instance: getInstance(result),
		}}}
	}
	return nil
}

func (*menuButtonDelegate) NewForCustom(menuButton *TCEFMenuButtonComponent) *ICefMenuButtonDelegate {
	if !menuButton.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.MenuButtonDelegateRef_CreateForCustom).Call(menuButton.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMenuButtonDelegate{&ICefButtonDelegate{&ICefViewDelegate{
			instance: getInstance(result),
			ct:       consts.CtOther,
		}}}
	}
	return nil
}

// / Called when |button| is pressed. Call ICefMenuButton.ShowMenu() to
// / show a popup menu at |screen_point|. When showing a custom popup such as a
// / window keep a reference to |button_pressed_lock| until the popup is hidden
// / to maintain the pressed button state.
func (m *ICefMenuButtonDelegate) SetOnMenuButtonPressed(fn menuButtonOnMenuButtonPressed) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.MenuButtonDelegate_SetOnMenuButtonPressed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ICefMenuButtonPressedLock
// include/capi/views/cef_menu_button_delegate_capi.h (cef_menu_button_pressed_lock_t)
type ICefMenuButtonPressedLock struct {
	base TCefBaseRefCounted
}

// Instance 实例
func (m *ICefMenuButtonPressedLock) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.base.instance)
}

type menuButtonOnMenuButtonPressed func(menuButton *ICefMenuButton, screenPoint TCefPoint, buttonPressedLock *ICefMenuButtonPressedLock)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case menuButtonOnMenuButtonPressed:
			button := &ICefMenuButton{&ICefLabelButton{&ICefButton{&ICefView{instance: getPtr(0)}}}}
			screenPoint := *(*TCefPoint)(getPtr(1))
			buttonPressedLock := &ICefMenuButtonPressedLock{base: TCefBaseRefCounted{instance: getPtr(3)}}
			fn.(menuButtonOnMenuButtonPressed)(button, screenPoint, buttonPressedLock)
		default:
			return false
		}
		return true
	})
}
