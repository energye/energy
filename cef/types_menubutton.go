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
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefMenuButton
// include/capi/views/cef_menu_button_capi.h (cef_menu_button_t)
type ICefMenuButton struct {
	*ICefLabelButton
}

// MenuButtonRef -> ICefMenuButton
var MenuButtonRef menuButton

type menuButton uintptr

func (*menuButton) New(delegate *ICefMenuButtonDelegate, text string) *ICefMenuButton {
	var result uintptr
	imports.Proc(def.MenuButtonRef_New).Call(delegate.Instance(), api.PascalStr(text), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMenuButton{&ICefLabelButton{&ICefButton{&ICefView{instance: getInstance(result)}}}}
	}
	return nil
}

func (m *ICefMenuButton) ShowMenu(menuModel *ICefMenuModel, screenPoint TCefPoint, anchorPosition consts.TCefMenuAnchorPosition) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuButton_ShowMenu).Call(m.Instance(), menuModel.Instance(), uintptr(unsafe.Pointer(&screenPoint)), uintptr(anchorPosition))
}

func (m *ICefMenuButton) TriggerMenu() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuButton_TriggerMenu).Call(m.Instance())
}
