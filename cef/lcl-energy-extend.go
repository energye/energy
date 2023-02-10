package cef

import (
	"github.com/energye/energy/common"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

func PanelBevelColor(panel *lcl.TPanel) types.TColor {
	r1, _, _ := common.Proc(internale_Ext_Panel_GetBevelColor).Call(panel.Instance())
	return types.TColor(r1)
}

func SetPanelBevelColor(panel *lcl.TPanel, colors types.TColor) {
	common.Proc(internale_Ext_Panel_SetBevelColor).Call(panel.Instance(), uintptr(colors))
}
