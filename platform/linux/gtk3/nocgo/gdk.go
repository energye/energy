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

var gdk3 *linux.DnyLibrary

func init() {
	gdk3 = linux.LibLoad(linux.Libgdk3)
	gdk3.Table = []*imports.Table{
		// screen
		imports.NewTable("gdk_screen_get_default", 0),
		imports.NewTable("gdk_screen_get_rgba_visual", 0),
		imports.NewTable("gdk_screen_get_system_visual", 0),
		imports.NewTable("gdk_screen_is_composited", 0),
		imports.NewTable("gdk_screen_get_root_window", 0),
		imports.NewTable("gdk_screen_get_display", 0),
		imports.NewTable("gdk_screen_get_resolution", 0),
		imports.NewTable("gdk_screen_set_resolution", 0),
		// DragContext
		imports.NewTable("gdk_drag_context_list_targets", 0),
		imports.NewTable("gtk_drag_finish", 0),
		imports.NewTable("gdk_drag_status", 0),
		// Atom
		imports.NewTable("gdk_atom_name", 0),
		imports.NewTable("gdk_atom_intern", 0),
		// Event
		imports.NewTable("gdk_event_free", 0),
		imports.NewTable("gdk_event_get_scancode", 0),
		// window
		imports.NewTable("gdk_window_set_decorations", 0),
		imports.NewTable("gdk_window_get_width", 0),
		imports.NewTable("gdk_window_get_height", 0),
		imports.NewTable("gdk_window_get_root_origin", 0),
		imports.NewTable("gdk_window_get_origin", 0),
		imports.NewTable("gdk_window_get_device_position", 0),
		imports.NewTable("gdk_window_set_override_redirect", 0),
		imports.NewTable("gdk_window_get_state", 0),
		// rectangle
		imports.NewTable("gdk_rectangle_intersect", 0),
		imports.NewTable("gdk_rectangle_union", 0),
		// Display
		imports.NewTable("gdk_display_open", 0),
		imports.NewTable("gdk_display_get_default", 0),
		imports.NewTable("gdk_display_get_name", 0),
		imports.NewTable("gdk_display_get_default_screen", 0),
		imports.NewTable("gdk_display_device_is_grabbed", 0),
		imports.NewTable("gdk_display_beep", 0),
		imports.NewTable("gdk_display_sync", 0),
		imports.NewTable("gdk_display_flush", 0),
		imports.NewTable("gdk_display_close", 0),
		imports.NewTable("gdk_display_is_closed", 0),
		imports.NewTable("gdk_display_get_event", 0),
		imports.NewTable("gdk_display_peek_event", 0),
		imports.NewTable("gdk_display_put_event", 0),
		imports.NewTable("gdk_display_has_pending", 0),
		imports.NewTable("gdk_display_set_double_click_time", 0),
		imports.NewTable("gdk_display_set_double_click_distance", 0),
		imports.NewTable("gdk_display_supports_cursor_color", 0),
		imports.NewTable("gdk_display_supports_cursor_alpha", 0),
		imports.NewTable("gdk_display_get_default_cursor_size", 0),
		imports.NewTable("gdk_display_get_maximal_cursor_size", 0),
		imports.NewTable("gdk_display_get_default_group", 0),
		imports.NewTable("gdk_display_supports_selection_notification", 0),
		imports.NewTable("gdk_display_request_selection_notification", 0),
		imports.NewTable("gdk_display_supports_clipboard_persistence", 0),
		imports.NewTable("gdk_display_supports_shapes", 0),
		imports.NewTable("gdk_display_supports_input_shapes", 0),
		imports.NewTable("gdk_display_notify_startup_complete", 0),
		imports.NewTable("gdk_display_get_default_seat", 0),
		imports.NewTable("gdk_x11_display_get_xdisplay", 0),
		// DeviceManager
		imports.NewTable("gdk_device_manager_get_display", 0),
		// DisplayManager
		imports.NewTable("gdk_display_manager_get", 0),
		imports.NewTable("gdk_display_manager_get_default_display", 0),
		imports.NewTable("gdk_display_manager_set_default_display", 0),
		imports.NewTable("gdk_display_manager_open_display", 0),
		// Seat
		imports.NewTable("gdk_seat_get_pointer", 0),
		// Device
		imports.NewTable("gdk_device_get_position", 0),
		// Visual
		imports.NewTable("gdk_visual_get_blue_pixel_details", 0),
		imports.NewTable("gdk_visual_get_depth", 0),
		imports.NewTable("gdk_visual_get_green_pixel_details", 0),
		imports.NewTable("gdk_visual_get_red_pixel_details", 0),
		imports.NewTable("gdk_visual_get_visual_type", 0),
		imports.NewTable("gdk_visual_get_screen", 0),
		// RGBA
		imports.NewTable("gdk_rgba_parse", 0),
		imports.NewTable("gdk_rgba_to_string", 0),
		imports.NewTable("gdk_rgba_free", 0),
		imports.NewTable("gdk_rgba_equal", 0),
		imports.NewTable("gdk_rgba_hash", 0),
	}
	gdk3.SetLibClose()
	gdk3.MapperIndex()
}
