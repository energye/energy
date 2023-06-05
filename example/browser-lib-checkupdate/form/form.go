//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package form

import (
	"fmt"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
)

type callback func(m *UpdateForm)

var OnAfter func()
var OnCreate callback

type UpdateForm struct {
	*lcl.TForm
	TitlePanel          *lcl.TPanel
	UpdatePromptPanel   *lcl.TPanel
	UpdateProgressPanel *lcl.TPanel
}

func (m *UpdateForm) OnFormCreate(sender lcl.IObject) {
	fmt.Println("OnFormCreate")
	var after = true
	m.SetOnActivate(func(sender lcl.IObject) {
		if OnAfter != nil && after {
			fmt.Println("SetOnActivate")
			after = false
			OnAfter()
		}
	})
	if OnCreate != nil {
		OnCreate(m)
	}

}

func (m *UpdateForm) NewPanel() *lcl.TPanel {
	var result = lcl.NewPanel(m)
	result.SetParent(m)
	result.SetBevelInner(types.BvNone)
	result.SetBevelOuter(types.BvNone)
	result.SetBorderStyle(types.BsNone)
	result.SetColor(colors.ClWhite)
	return result
}
