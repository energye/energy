// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

package window

import (
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
)

func CallFormCreate(target lcl.IFormHook, sender lcl.IObject) {
	lcl.CallFormCreate(target, sender)
}

func CallFormShow(target lcl.IFormHook, sender lcl.IObject) {
	lcl.CallFormShow(target, sender)
}

func CallFormCloseQuery(target lcl.IFormHook, sender lcl.IObject, canClose *bool) bool {
	return lcl.CallFormCloseQuery(target, sender, canClose)
}

func CallFormClose(target lcl.IFormHook, sender lcl.IObject, closeAction *types.TCloseAction) bool {
	return lcl.CallFormClose(target, sender, closeAction)
}
