package browse

import (
	"embed"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/consts"
)

var Assets embed.FS

// Control IDs for Views in the top-level Window.
const (
	ID_WINDOW int32 = iota
	ID_BROWSER_VIEW
	ID_BACK_BUTTON
	ID_FORWARD_BUTTON
	ID_STOP_BUTTON
	ID_RELOAD_BUTTON
	ID_URL_TEXTFIELD
	ID_MENU_BUTTON
	// Reserved range of top menu button IDs.
	ID_TOP_MENU_FIRST
	ID_TOP_MENU_LAST = ID_TOP_MENU_FIRST + 10
)
const kMenuBarGroupId int32 = 100

func LoadImage(png string) *cef.ICefImage {
	if png == "" {
		return nil
	}
	pngData, err := Assets.ReadFile("assets/icon/" + png)
	if err != nil {
		panic(err)
	}
	icon := cef.ImageRef.New()
	icon.AddPng(1, pngData)
	return icon
}

type ImageButton struct {
	btn     *cef.ICefMenuButton
	enable  *cef.ICefImage
	disable *cef.ICefImage
}

func CreateImageButton(text, tooltip, iconEnable, iconDisable string, id int32, delegate *cef.ICefMenuButtonDelegate) *ImageButton {
	button := new(ImageButton)
	button.enable = LoadImage(iconEnable)
	button.disable = LoadImage(iconDisable)
	button.btn = cef.MenuButtonRef.New(delegate, text)
	button.btn.SetID(id)
	button.btn.SetInkDropEnabled(true)
	button.btn.SetEnabled(true)    // 默认为关闭
	button.btn.SetFocusable(false) // 不要把焦点放在按钮上
	button.btn.SetMinimumSize(cef.TCefSize{})
	button.btn.SetMaximumSize(cef.TCefSize{Height: 40, Width: 40})
	button.btn.SetImage(consts.CEF_BUTTON_STATE_NORMAL, button.enable)
	button.btn.SetTooltipText(tooltip)
	return button
}
