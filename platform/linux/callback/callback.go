//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package callback

import (
	"github.com/energye/energy/v3/platform/linux/types"
)

// 用于动态不同函数签名
const (
	C_trampoline_2_void                      = "c_trampoline_2_void"
	C_trampoline_3_void                      = "c_trampoline_3_void"
	C_trampoline_4_void                      = "c_trampoline_4_void"
	C_trampoline_3_gboolean                  = "c_trampoline_3_gboolean"
	C_trampoline_4_gboolean                  = "c_trampoline_4_gboolean"
	C_trampoline_8_void_drag_data_received   = "c_trampoline_8_void_drag_data_received"
	C_trampoline_6_gboolean_drag_drop_motion = "c_trampoline_6_gboolean_drag_drop_motion"
	C_trampoline_4_void_drag_leave           = "c_trampoline_4_void_drag_leave"
)

type SignalHandlerID struct {
	Widget    types.PGtkWidget
	HandlerID types.GULong
	Id        uint64
}
