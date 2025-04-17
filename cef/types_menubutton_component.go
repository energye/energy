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

// TCEFMenuButtonComponent
type TCEFMenuButtonComponent struct {
	*TCEFLabelButtonComponent
}

// MenuButtonComponentRef -> TCEFMenuButtonComponent
var MenuButtonComponentRef menuButtonComponent

type menuButtonComponent uintptr

func (*menuButtonComponent) New() *TCEFMenuButtonComponent {
	var result uintptr
	imports.Proc(def.MenuButtonComponent_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &TCEFMenuButtonComponent{&TCEFLabelButtonComponent{&TCEFButtonComponent{&TCEFViewComponent{
			instance: getInstance(result),
		}}}}
	}
	return nil
}

func (m *TCEFMenuButtonComponent) CreateMenuButton(text string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuButtonComponent_CreateMenuButton).Call(m.Instance(), api.PascalStr(text))
}

func (m *TCEFMenuButtonComponent) ShowMenu(menuModel *ICefMenuModel, screenPoint TCefPoint, anchorPosition consts.TCefMenuAnchorPosition) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuButtonComponent_ShowMenu).Call(m.Instance(), menuModel.Instance(), uintptr(unsafe.Pointer(&screenPoint)), uintptr(anchorPosition))
}

func (m *TCEFMenuButtonComponent) TriggerMenu() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuButtonComponent_TriggerMenu).Call(m.Instance())
}

func (m *TCEFMenuButtonComponent) SetOnMenuButtonPressed(fn menuButtonOnMenuButtonPressed) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuButtonComponent_SetOnMenuButtonPressed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}
