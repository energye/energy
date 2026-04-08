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
	"github.com/energye/energy/v3/pkgs/linux"
	"github.com/energye/lcl/api/imports"
)

var gdk3 *linux.DnyLibrary

func init() {
	gdk3 = linux.LibLoad(linux.Libgdk3)
	gdk3.Table = []*imports.Table{
		// screen
		imports.NewTable("gdk_screen_get_rgba_visual", 0),
		imports.NewTable("gdk_screen_is_composited", 0),
		// DragContext
		imports.NewTable("gdk_drag_context_list_targets", 0),
		imports.NewTable("gtk_drag_finish", 0),
		imports.NewTable("gdk_drag_status", 0),
		// Atom
		imports.NewTable("gdk_atom_name", 0),
		imports.NewTable("gdk_atom_intern", 0),
	}
	gdk3.SetLibClose()
	gdk3.MapperIndex()
}
