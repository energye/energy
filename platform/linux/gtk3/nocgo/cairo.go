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
		imports.NewTable("cairo_destroy", 0),
		imports.NewTable("cairo_save", 0),
		imports.NewTable("cairo_restore", 0),
		imports.NewTable("cairo_set_source_rgb", 0),
		imports.NewTable("cairo_set_source_rgba", 0),
		imports.NewTable("cairo_set_line_width", 0),
		imports.NewTable("cairo_get_line_width", 0),
		imports.NewTable("cairo_clip", 0),
		imports.NewTable("cairo_clip_preserve", 0),
		imports.NewTable("cairo_clip_extents", 0),
		imports.NewTable("cairo_in_clip", 0),
		imports.NewTable("cairo_reset_clip", 0),
		imports.NewTable("cairo_rectangle", 0),
		imports.NewTable("cairo_arc", 0),
		imports.NewTable("cairo_arc_negative", 0),
		imports.NewTable("cairo_line_to", 0),
		imports.NewTable("cairo_curve_to", 0),
		imports.NewTable("cairo_move_to", 0),
		imports.NewTable("cairo_fill", 0),
		imports.NewTable("cairo_fill_preserve", 0),
		imports.NewTable("cairo_fill_extents", 0),
		imports.NewTable("cairo_in_fill", 0),
		imports.NewTable("cairo_close_path", 0),
		imports.NewTable("cairo_new_path", 0),
		imports.NewTable("cairo_get_current_point", 0),
		imports.NewTable("cairo_paint", 0),
		imports.NewTable("cairo_paint_with_alpha", 0),
		imports.NewTable("cairo_stroke", 0),
		imports.NewTable("cairo_stroke_preserve", 0),
		imports.NewTable("cairo_stroke_extents", 0),
		imports.NewTable("cairo_in_stroke", 0),
		imports.NewTable("cairo_copy_page", 0),
		imports.NewTable("cairo_show_page", 0),
		imports.NewTable("cairo_set_fill_rule", 0),
		imports.NewTable("cairo_get_fill_rule", 0),
		imports.NewTable("cairo_set_line_cap", 0),
		imports.NewTable("cairo_get_line_cap", 0),
		imports.NewTable("cairo_set_line_join", 0),
		imports.NewTable("cairo_get_line_join", 0),
		imports.NewTable("cairo_set_operator", 0),
		imports.NewTable("cairo_get_operator", 0),
		imports.NewTable("cairo_set_dash", 0),
		imports.NewTable("cairo_get_dash_count", 0),
		imports.NewTable("cairo_get_dash", 0),
		imports.NewTable("cairo_set_tolerance", 0),
		imports.NewTable("cairo_get_tolerance", 0),
		imports.NewTable("cairo_set_miter_limit", 0),
		imports.NewTable("cairo_get_miter_limit", 0),
		imports.NewTable("cairo_push_group", 0),
		imports.NewTable("cairo_pop_group_to_source", 0),
	}
	cairo.SetLibClose()
	cairo.MapperIndex()
}
