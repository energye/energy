//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	"github.com/energye/energy/v3/platform/linux"
	"github.com/energye/lcl/api/imports"
)

var cairo *linux.DnyLibrary

func init() {
	cairo = linux.LibLoad(linux.Libcairo)
	cairo.Table = []*imports.Table{
		imports.NewTable("cairo_status", 0),
	}
	cairo.SetLibClose()
	cairo.MapperIndex()
}
