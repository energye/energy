//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// auto update application

package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef/autoupdate"
	"github.com/energye/energy/v2/cef/autoupdate/internal"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api/dllimports"
)

var (
	//go:embed resources
	resources embed.FS
	// form
	updateForm *internal.UpdateForm
	version    = []*dllimports.ImportTable{
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("LibVersion", 0),
		dllimports.NewEnergyImport("LibBuildVersion", 0),
	}
)

func main() {
	imports.SetEnergyImportDefs(version)
	inits.Init(nil, &resources)
	autoupdate.IsCheckUpdate(true)
	autoupdate.CanUpdateLiblcl = func(model *autoupdate.Model, level int) {
		fmt.Println(*model)
		fmt.Println(model.Versions[model.Latest])
		fmt.Println(level)
		internal.OnCreate = func(m *internal.UpdateForm) {
			lcl.Application.Icon().LoadFromFSFile("resources/icon.ico")
			m.SetCaption("应用更新")
			ok := lcl.NewButton(m)
			ok.SetParent(m)
			ok.SetCaption("确定")
			ok.SetOnClick(func(sender lcl.IObject) {
				fmt.Println("确定")
			})
		}
		// create update form
		lcl.RunApp(&updateForm)
	}
	autoupdate.CheckUpdate()
}
