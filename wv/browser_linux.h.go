//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build linux

package wv

import (
	. "github.com/energye/energy/v3/platform/linux/types"
)

func (m *TWebview) beginResize(rootX, rootY, timestamp int32) {
	if m.resizeHT == "" || m.window == nil {
		return
	}
	gtkWin := m.window.GTKWindow()
	if gtkWin == nil {
		return
	}
	edge := getWindowEdge(m.resizeHT)
	gtkWin.BeginResizeDrag(edge, BUTTON_PRIMARY, int(rootX), int(rootY), uint32(timestamp))
}

func getWindowEdge(ht string) WindowEdge {
	switch ht {
	case "nw-resize":
		return GDK_WINDOW_EDGE_NORTH_WEST
	case "n-resize":
		return GDK_WINDOW_EDGE_NORTH
	case "ne-resize":
		return GDK_WINDOW_EDGE_NORTH_EAST
	case "w-resize":
		return GDK_WINDOW_EDGE_WEST
	case "e-resize":
		return GDK_WINDOW_EDGE_EAST
	case "sw-resize":
		return GDK_WINDOW_EDGE_SOUTH_WEST
	case "s-resize":
		return GDK_WINDOW_EDGE_SOUTH
	case "se-resize":
		return GDK_WINDOW_EDGE_SOUTH_EAST
	default:
		return GDK_WINDOW_EDGE_NORTH_WEST // fallback
	}
}
