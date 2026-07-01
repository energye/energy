//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build prod

package cef

import (
	"github.com/energye/cef/config"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/api/libname"
	"github.com/energye/lcl/tool"
	"path/filepath"
)

func loadLibENERGYRuntime() {
	api.SetOnLoadLibCallback(func() (lib imports.DLL, err error) {
		libPath := libname.LibName
		if libPath != "" {
			// Custom Load Directory
			// lib, err = imports.NewDLL(libPath)
		} else if tool.IsDarwin() {
			// MacOS Fixed Load Directory
			if libname.EnableUniversalBinary {
				libPath = "@executable_path/../Frameworks/" + libname.DarwinUniversalBinaryName
			} else {
				libPath = "@executable_path/../Frameworks/" + libname.GetDLLName()
			}
		} else {
			// Windows, Linux
			if cfg := config.GConfig; cfg != nil {
				libPath = filepath.Join(cfg.ChromiumPath(), libname.GetDLLName())
			}
			if !tool.IsExist(libPath) {
				// Finally Try Relative Directory
				libPath = libname.GetDLLName()
			}
		}
		if libPath != "" {
			libname.LibName = libPath
			lib, err = imports.NewDLL(libPath)
		}
		if lib == 0 {
			if err != nil {
				println("[ERROR] Load runtime libenergy", err.Error())
			}
			println("[ERROR] Path:", libname.LibName)
			panic(`Failed initialize runtime libenergy`)
		}
		return
	})
}
