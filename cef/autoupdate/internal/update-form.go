package internal

//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

import (
	"fmt"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

type callback func(m *UpdateForm)

var OnAfter func()
var OnCreate callback

type UpdateForm struct {
	*lcl.TForm
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
	m.EnabledMinimize(false)
	m.EnabledMaximize(false)
	m.SetFormStyle(types.FsSystemStayOnTop)
	m.SetPosition(types.PoDesktopCenter)
	//m.SetBorderStyle(types.BsSingle)
	m.SetBorderStyle(types.BsNone)
	//m.SetShowInTaskBar(types.StNever)
	m.SetWidth(450)
	m.SetHeight(230)
	if OnCreate != nil {
		OnCreate(m)
	}

}
