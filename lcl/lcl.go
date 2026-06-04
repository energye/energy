//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/lcl"
)

func SetOnBeforeRun(fn func()) {
	lcl.SetOnBeforeRun(fn)
}

func Init() {
	lcl.Init()
}

func Run(forms ...lcl.IEngForm) {
	lcl.Run(forms...)
}

func Application() lcl.IApp {
	return lcl.Application
}
