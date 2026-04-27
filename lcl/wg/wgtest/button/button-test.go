package main

import (
	"fmt"
	"github.com/energye/energy/v3/lcl/wg"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/colors"
	"os"
	"path/filepath"
	"time"
)

type TMainForm struct {
	lcl.TEngForm
	oldWndPrc uintptr
}

var MainForm TMainForm

var (
	wd, _       = os.Getwd()
	examplePath = filepath.Join(wd, "lcl", "wg", "wgtest", "button")
)

func main() {
	lcl.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.NewForm(&MainForm)
	lcl.Application.Run()
}

func (m *TMainForm) FormCreate(sender lcl.IObject) {
	m.SetCaption("ENERGY 自绘(自定义)按钮")
	m.SetPosition(types.PoScreenCenter)
	m.SetWidth(800)
	m.SetHeight(600)
	m.SetDoubleBuffered(true)
	m.SetColor(colors.RGBToColor(56, 57, 60))

	box := lcl.NewPanel(m)
	box.SetBevelInner(types.BvNone)
	box.SetBevelOuter(types.BvNone)
	//box.SetBounds(0, 0, 800, 600)
	box.SetAlign(types.AlClient)
	//box.SetParentColor(false)
	box.SetBorderStyleToBorderStyle(types.BsNone)
	box.SetParent(m)

	isCloseing := false
	m.SetOnCloseQuery(func(sender lcl.IObject, canClose *bool) {
		isCloseing = true
	})

	{
		click := func(sender lcl.IObject) {
			fmt.Println(lcl.AsGraphicControl(sender).Caption())
			m.SetCaption(time.Now().String())
		}
		cus := wg.NewButton(m)
		cus.SetShowHint(true)
		cus.RoundedCorner = cus.RoundedCorner.Exclude(wg.RcLeftBottom).Exclude(wg.RcRightBottom)
		cus.SetText("上圆角-禁用/启用")
		cus.SetHint("上圆角上圆角")
		cus.Font().SetSize(12)
		cus.Font().SetColor(colors.Cl3DFace)
		cus.SetRadius(5)
		cus.SetBorderColor(wg.BbdNone, colors.ClWhite)
		cus.SetBorderWidth(wg.BbdNone, 1)
		cusRect := types.TRect{Left: 50, Top: 50}
		cusRect.SetWidth(200)
		cusRect.SetHeight(40)
		cus.SetBoundsRect(cusRect)
		cus.SetDefaultColor(colors.RGBToColor(86, 88, 93), colors.RGBToColor(86, 88, 93))
		start, end := cus.DefaultColor()
		cus.SetEnterColor(wg.DarkenColor(start, 0.1), wg.DarkenColor(end, 0.1))
		cus.SetDownColor(wg.DarkenColor(start, 0.2), wg.DarkenColor(end, 0.2))
		cus.SetOnCloseClick(func(sender lcl.IObject) {
			fmt.Println("点击了 X")
		})
		cus.SetIconFavorite(filepath.Join(examplePath, "resources", "icon.png"))
		cus.SetIconClose(filepath.Join(examplePath, "resources", "close.png"))
		cus.SetOnClick(click)
		cus.SetOnDblClick(func(sender lcl.IObject) {
			println("DblClick")
		})
		cus.SetParent(box)

		cus2 := wg.NewButton(m)
		cus2.SetText("大圆角-渐变色有点长\n\n换行文本换行文本")
		cus2.SetBoundsRect(types.TRect{Left: 50, Top: 150, Right: 250, Bottom: 220})
		var (
			defaultStartColor = colors.RGBToColor(255, 100, 0)
			defaultEndColor   = colors.RGBToColor(69, 81, 143)
		)
		cus2.SetRadius(20)
		cus2.SetAlpha(255)
		cus2.SetBorderColor(wg.BbdNone, colors.ClRed)
		cus2.SetBorderWidth(wg.BbdNone, 1)
		cus2.Font().SetColor(colors.ClWhite)
		cus2.Font().SetSize(16)
		cus2.SetDefaultColor(defaultStartColor, defaultEndColor)
		cus2.SetEnterColor(wg.DarkenColor(defaultStartColor, 0.1), wg.DarkenColor(defaultEndColor, 0.1))
		cus2.SetDownColor(wg.DarkenColor(defaultStartColor, 0.2), wg.DarkenColor(defaultEndColor, 0.2))
		cus2.SetOnClick(click)
		cus2.SetParent(box)

		cus3 := wg.NewButton(m)
		cus3.SetText("小圆角")
		cus3.SetBoundsRect(types.TRect{Left: 50, Top: 250, Right: 250, Bottom: 320})
		cus3.Font().SetColor(colors.ClYellow)
		cus3.SetRadius(8)
		cus3.SetAlpha(255)
		cus3.SetDefaultColor(colors.RGBToColor(0, 180, 0), colors.RGBToColor(0, 100, 0))
		start, end = cus3.DefaultColor()
		cus3.SetEnterColor(wg.DarkenColor(start, 0.1), wg.DarkenColor(end, 0.1))
		cus3.SetDownColor(wg.DarkenColor(start, 0.2), wg.DarkenColor(end, 0.2))
		cus3.SetOnClick(click)
		cus3.SetParent(box)

		cus41 := wg.NewButton(m)
		cus41.SetText("红的-半透明")
		cus41.Font().SetColor(colors.ClWhite)
		cus41.Font().SetSize(12)
		cus41.SetColor(colors.ClRed)
		cusBtn := types.TRect{Left: 80, Top: 370}
		cusBtn.SetWidth(200)
		cusBtn.SetHeight(70)
		cus41.SetBoundsRect(cusBtn)
		cus41.SetRadius(35)
		cus41.SetAlpha(125)
		cus41.SetOnClick(click)
		cus41.SetParent(box)

		cus4 := wg.NewButton(m)
		cus4.SetText("大大圆角-半透明")
		cus4.Font().SetColor(colors.ClWhite)
		cus4.Font().SetSize(12)
		cus4.SetBoundsRect(types.TRect{Left: 50, Top: 350, Right: 250, Bottom: 420})
		cus4.SetRadius(35)
		cus4.SetAlpha(125)
		cus4.SetOnClick(click)
		cus4.SetParent(box)

		cus5 := wg.NewButton(m)
		cus5.SetText("X")
		cus5.Font().SetColor(colors.ClWhite)
		cus5.Font().SetSize(14)
		rect5 := types.TRect{Left: 50, Top: 450}
		rect5.SetSize(50, 50)
		cus5.SetBoundsRect(rect5)
		cus5.SetRadius(35)
		cus5.SetAlpha(255)
		cus5.SetOnClick(click)
		cus5.SetParent(box)

		cus6 := wg.NewButton(m)
		cus6.SetText("< X >")
		cus6.Font().SetColor(colors.ClWhite)
		cus6.Font().SetSize(14)
		rect6 := types.TRect{Left: 150, Top: 450}
		rect6.SetSize(50, 50)
		cus6.SetBoundsRect(rect6)
		cus6.SetRadius(5)
		cus6.SetAlpha(255)
		cus6.SetOnClick(click)
		cus6.SetParent(box)

		setDisable := func(button *wg.TButton) {
			button.SetDisable(!button.Disable())
			if button.Disable() {
				button.Font().SetColor(colors.ClBlack)
			} else {
				button.Font().SetColor(colors.ClWhite)
			}
		}
		//disableBtn := wg.NewButton(m)
		disableBtn := lcl.NewButton(m)
		disableBtn.SetCaption("禁用/启用 按钮")
		disableBtnRect := types.TRect{Left: 250, Top: 450}
		disableBtnRect.SetSize(150, 30)
		disableBtn.SetBoundsRect(disableBtnRect)
		disableBtn.SetOnClick(func(sender lcl.IObject) {
			setDisable(cus)
			//setDisable(cus2)
			//setDisable(cus3)
			//setDisable(cus6)
			//setDisable(cus4)
			//setDisable(cus5)
		})
		disableBtn.SetParent(box)
		{
			cus := wg.NewButton(m)
			cus.SetShowHint(true)
			cus.SetAutoSize(true)
			cus.SetText("自动宽1")
			cus.SetHint("自动宽1")
			cus.Font().SetSize(12)
			cus.Font().SetColor(colors.Cl3DFace)
			rect := types.TRect{Left: 50, Top: 10}
			rect.SetWidth(100)
			rect.SetHeight(30)
			cus.SetBoundsRect(rect)
			cus.SetAutoSize(true)
			//cus.SetStartColor(colors.RGBToColor(86, 88, 93))
			//cus.SetEndColor(colors.RGBToColor(86, 88, 93))
			cus.RoundedCorner = cus.RoundedCorner.Exclude(wg.RcLeftBottom).Exclude(wg.RcRightBottom)
			cus.SetOnCloseClick(func(sender lcl.IObject) {
				fmt.Println("点击了 X")
			})
			cus.SetIconFavorite(filepath.Join(examplePath, "resources", "icon.png"))
			cus.SetIconClose(filepath.Join(examplePath, "resources", "close.png"))
			cus.SetOnClick(click)
			cus.SetParent(box)
			go func() {
				i := 0
				str := "自DEF动CD宽AB"
				for {
					if i > 10 {
						time.Sleep(time.Second)
						if isCloseing {
							break
						}
						i = 0
						cus.SetText(str)
					} else {
						time.Sleep(time.Second / 10)
						if isCloseing {
							break
						}
						i++
						cus.SetText(cus.Text() + str)
					}
				}
			}()
		}
		{
			cus7 := wg.NewButton(m)
			cus7.SetAutoSize(true)
			cus7.SetText("自动宽2")
			cus7.Font().SetColor(colors.ClWhite)
			//cus7.SetStartColor(colors.RGBToColor(41, 42, 43))
			//cus7.SetEndColor(colors.RGBToColor(80, 81, 82))
			rect := types.TRect{Left: 50, Top: 110}
			rect.SetWidth(100)
			rect.SetHeight(30)
			cus7.SetBoundsRect(rect)
			cus7.SetRadius(8)
			cus7.SetAlpha(255)
			cus7.SetAutoSize(true)
			cus7.SetOnClick(click)
			cus7.SetIconFavorite(filepath.Join(examplePath, "resources", "icon.png"))
			cus7.SetIconClose(filepath.Join(examplePath, "resources", "close.png"))
			cus7.SetParent(box)
			go func() {
				i := 0
				str := "自ab动CD宽ef"
				for {
					if i > 10 {
						time.Sleep(time.Second)
						if isCloseing {
							break
						}
						i = 0
						cus7.SetText(str)
					} else {
						time.Sleep(time.Second / 5)
						if isCloseing {
							break
						}
						i++
						cus7.SetText(cus7.Text() + str)
					}
				}
			}()
		}
	}
	{
		if false {
			bgColors := []colors.TColor{colors.ClBlue, colors.ClRed, colors.ClGreen, colors.ClYellow}
			go func() {
				i := 0
				for {
					time.Sleep(time.Second)
					if isCloseing {
						break
					}
					lcl.RunOnMainThreadAsync(func(id uint32) {
						box.SetColor(bgColors[i])
					})
					i++
					if i >= len(bgColors) {
						i = 0
					}
				}
			}()
		}
	}
	{
		textAlignLeft := wg.NewButton(m)
		textAlignLeftRect := types.TRect{Left: 255, Top: 150}
		textAlignLeftRect.SetWidth(150)
		textAlignLeftRect.SetHeight(80)
		textAlignLeft.SetBoundsRect(textAlignLeftRect)
		textAlignLeft.Font().SetColor(colors.ClWhite)
		textAlignLeft.Font().SetSize(10)
		textAlignLeft.TextAlign = wg.TextAlignLeft
		textAlignLeft.TextLineSpacing = 8
		textAlignLeft.SetRadius(10)
		textAlignLeft.SetAlpha(255)
		textAlignLeft.SetText("左对齐\n换行文本换行文本\n行间距8px, 第三行")
		textAlignLeft.SetIconFavorite(filepath.Join(examplePath, "resources", "icon.png"))
		textAlignLeft.SetIconClose(filepath.Join(examplePath, "resources", "close.png"))
		textAlignLeft.SetParent(box)

		textAlignRight := wg.NewButton(m)
		textAlignRightRect := types.TRect{Left: 455, Top: 150}
		textAlignRightRect.SetWidth(150)
		textAlignRightRect.SetHeight(40)
		textAlignRight.SetBoundsRect(textAlignRightRect)
		textAlignRight.Font().SetColor(colors.ClWhite)
		textAlignRight.Font().SetSize(10)
		textAlignRight.TextAlign = wg.TextAlignRight
		textAlignRight.TextOffSetX = -5
		textAlignRight.SetRadius(10)
		textAlignRight.SetAlpha(255)
		textAlignRight.SetText("右对齐\n\n换行文本换行文本")
		textAlignRight.SetIconFavorite(filepath.Join(examplePath, "resources", "icon.png"))
		textAlignRight.SetIconClose(filepath.Join(examplePath, "resources", "close.png"))
		textAlignRight.SetParent(box)

		textAlignRight2 := wg.NewButton(m)
		textAlignRight2Rect := types.TRect{Left: 455, Top: 200}
		textAlignRight2Rect.SetWidth(150)
		textAlignRight2Rect.SetHeight(40)
		textAlignRight2.SetBoundsRect(textAlignRight2Rect)
		textAlignRight2.Font().SetColor(colors.ClWhite)
		textAlignRight2.Font().SetSize(10)
		textAlignRight2.TextAlign = wg.TextAlignRight
		textAlignRight2.TextOffSetX = -5
		textAlignRight2.SetRadius(0)
		textAlignRight2.SetAlpha(255)
		textAlignRight2.SetText("右对齐\n\n换行文本换行文本")
		textAlignRight2.SetParent(box)
	}

}
