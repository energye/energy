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

import "github.com/energye/lcl/api/imports"

var gdk3 *dnyLibrary

func init() {
	gdk3 = libLoad(libgdk3)
	setLibClose(gdk3)
	gdk3.Table = []*imports.Table{
		// screen
		imports.NewTable("gdk_screen_get_rgba_visual", 0),
		imports.NewTable("gdk_screen_is_composited", 0),
	}
	gdk3.mapperIndex()
}
