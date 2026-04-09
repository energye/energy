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
	"unsafe"
)

var ptrSize = unsafe.Sizeof(uintptr(0))

var gtk3 *linux.DnyLibrary

func init() {
	gtk3 = linux.LibLoad(linux.Libgtk3)
	gtk3.Table = []*imports.Table{
		// window
		imports.NewTable("gtk_window_new", 0),
		imports.NewTable("gtk_window_get_default_size", 0),
		imports.NewTable("gtk_window_set_decorated", 0),
		imports.NewTable("gtk_window_maximize", 0),
		imports.NewTable("gtk_window_unmaximize", 0),
		imports.NewTable("gtk_window_fullscreen", 0),
		imports.NewTable("gtk_window_unfullscreen", 0),
		imports.NewTable("gtk_window_set_title", 0),
		imports.NewTable("gtk_window_get_title", 0),
		// container
		imports.NewTable("gtk_container_add", 0),
		imports.NewTable("gtk_container_remove", 0),
		imports.NewTable("gtk_container_check_resize", 0),
		imports.NewTable("gtk_container_get_children", 0),
		// widget
		imports.NewTable("gtk_widget_get_screen", 0),
		imports.NewTable("gtk_widget_set_visual", 0),
		imports.NewTable("gtk_widget_set_app_paintable", 0),
		imports.NewTable("gtk_widget_get_name", 0),
		imports.NewTable("gtk_widget_get_allocation", 0),
		imports.NewTable("gtk_widget_set_size_request", 0),
		imports.NewTable("gtk_widget_get_size_request", 0),
		imports.NewTable("gtk_widget_get_style_context", 0),
		// box
		imports.NewTable("gtk_box_pack_start", 0),
		imports.NewTable("gtk_box_pack_end", 0),
		// Layout
		imports.NewTable("gtk_layout_put", 0),
		imports.NewTable("gtk_layout_move", 0),
		imports.NewTable("gtk_layout_set_size", 0),
		imports.NewTable("gtk_layout_get_size", 0),
		// ScrolledWindow
		// CssProvider
		imports.NewTable("gtk_css_provider_new", 0),
		imports.NewTable("gtk_css_provider_load_from_path", 0),
		imports.NewTable("gtk_css_provider_load_from_data", 0),
		imports.NewTable("gtk_css_provider_to_string", 0),
		// StyleContext
		imports.NewTable("gtk_style_context_add_class", 0),
		imports.NewTable("gtk_style_context_remove_class", 0),
		imports.NewTable("gtk_style_context_add_provider", 0),
		// SelectionData
		imports.NewTable("gtk_selection_data_get_length", 0),
		imports.NewTable("gtk_selection_data_get_data_with_length", 0),
		imports.NewTable("gtk_selection_data_set", 0),
		imports.NewTable("gtk_selection_data_get_text", 0),
		imports.NewTable("gtk_selection_data_set_text", 0),
		imports.NewTable("gtk_selection_data_set_uris", 0),
		imports.NewTable("gtk_selection_data_get_uris", 0),
		imports.NewTable("gtk_selection_data_free", 0),
		// Entry
		imports.NewTable("gtk_entry_new", 0),
	}
	gtk3.SetLibClose()
	gtk3.MapperIndex()
}

func ToCBool(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

func ToGoBool(b uintptr) bool {
	return b != 0
}
