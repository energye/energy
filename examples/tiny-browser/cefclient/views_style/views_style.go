package views_style

import (
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/types"
)

var (
	g_background_color       types.TCefColor = 0
	g_background_hover_color types.TCefColor = 0
	g_text_color             types.TCefColor = 0
)

func ApplyTo(menuModel *cef.ICefMenuModel) {
	// 所有文本，除了非悬停加速器获得相同的颜色。
	menuModel.SetColorAt(-1, consts.CEF_MENU_COLOR_TEXT, g_text_color)
	menuModel.SetColorAt(-1, consts.CEF_MENU_COLOR_TEXT_HOVERED, g_text_color)
	menuModel.SetColorAt(-1, consts.CEF_MENU_COLOR_TEXT_ACCELERATOR, g_background_hover_color)
	menuModel.SetColorAt(-1, consts.CEF_MENU_COLOR_TEXT_ACCELERATOR_HOVERED, g_text_color)
	// 递归地给子菜单上色。
	for i := 0; i < int(menuModel.GetCount()); i++ {
		if menuModel.GetTypeAt(uint32(i)) == consts.MENUITEMTYPE_SUBMENU {
			ApplyTo(menuModel.GetSubMenuAt(uint32(i)))
		}
	}
}

func WindowApplyTo(window *cef.ICefWindow) {
	// 自定义默认背景颜色。
	//window.SetThemeColor(CEF_ColorPrimaryBackground, g_background_color)
	//
	//// 自定义默认菜单颜色。
	//window.SetThemeColor(CEF_ColorMenuBackground, g_background_color)
	//window.SetThemeColor(CEF_ColorMenuItemBackgroundHighlighted, g_background_hover_color)
	//window.SetThemeColor(CEF_ColorMenuItemBackgroundSelected, g_background_hover_color)
	//window.SetThemeColor(CEF_ColorMenuSeparator, g_text_color)
	//window.SetThemeColor(CEF_ColorMenuItemForeground, g_text_color)
	//window.SetThemeColor(CEF_ColorMenuItemForegroundHighlighted, g_text_color)
	//window.SetThemeColor(CEF_ColorMenuItemForegroundSelected, g_text_color)
	//
	//// 自定义默认文本字段颜色。
	//window.SetThemeColor(CEF_ColorTextfieldBackground, g_background_color)
	//window.SetThemeColor(CEF_ColorTextfieldOutline, g_text_color)
	//
	//// 自定义默认的Chrome工具栏颜色。
	//window.SetThemeColor(CEF_ColorToolbar, g_background_color)
	//window.SetThemeColor(CEF_ColorToolbarText, g_text_color)
	//window.SetThemeColor(CEF_ColorToolbarButtonIcon, g_text_color)
	//window.SetThemeColor(CEF_ColorToolbarButtonText, g_text_color)
	//window.SetThemeColor(CEF_ColorLocationBarBackground, g_background_color)
	//window.SetThemeColor(CEF_ColorLocationBarBackgroundHovered, g_background_hover_color)
	//window.SetThemeColor(CEF_ColorOmniboxText, g_text_color)
}
