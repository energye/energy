package browse

import (
	"embed"
	"github.com/energye/energy/v2/cef"
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
