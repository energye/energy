//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package wv

import (
	engLCL "github.com/energye/energy/v3/lcl"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/lcl"
)

// Run Launch application
func Run(forms ...lcl.IEngForm) {
	if gApplication != nil {
		gApplication.Start()
	}
	api.SetOnReleaseCallback(func() {
		DestroyGlobalLoader()
	})
	engLCL.Run(forms...)
}
