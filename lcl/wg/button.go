// Copyright © yanghy. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.

package wg

import (
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/colors"
	"path/filepath"
	"strings"
	"time"
)

// TextAlign 行间距, 多行文本
type TextAlign int8

const (
	TextAlignCenter TextAlign = iota // 居中对齐（默认）
	TextAlignRight                   // 右对齐
	TextAlignLeft                    // 左对齐（可选扩展）
)

// RoundedCorner 按钮圆角方向，默认四角
type RoundedCorner = int32

const (
	RcLeftTop RoundedCorner = iota
	RcRightTop
	RcLeftBottom
	RcRightBottom
)

// TRoundedCorners 按钮圆角方向，默认四角
type TRoundedCorners = types.TSet

// 图标默认边距
const iconMargin = 5

// TButtonState 按钮当着状态
type TButtonState = int32

const (
	BsDefault  TButtonState = iota // 默认状态
	BsEnter                        // 移入状态
	BsDown                         // 按下状态
	BsDisabled                     // 禁用状态
)

var (
	defaultButtonColor        = colors.RGBToColor(66, 133, 244)  // 淡蓝色
	defaultButtonColorDisable = colors.RGBToColor(200, 200, 200) // 浅灰色
)

// TButton 多功能自绘按钮
// 颜色状态: 默认颜色, 移入颜色, 按下颜色, 禁用颜色
// 当大小改变, 颜色改变 会重新绘制
type TButton struct {
	lcl.ICustomGraphicControl
	isDisable                          bool            // 是否禁用
	alpha                              byte            // 透明度 0 ~ 255
	radius                             int32           // 圆角度
	autoSize                           bool            // 自动大小
	text                               string          // 文本
	RoundedCorner                      TRoundedCorners // 按钮圆角方向，默认四角
	TextOffSetX, TextOffSetY           int32           // 文本显示偏移位置
	IconCloseOffSetX, IconCloseOffSetY int32           // 关闭按钮偏移位置
	TextAlign                          TextAlign       // 该校对齐
	TextLineSpacing                    int32           // 行间距 px
	// 图标
	iconFavorite       lcl.IPicture // 按钮前置图标, 靠左
	iconClose          lcl.IPicture // 按钮关闭图标, 靠右
	iconCloseHighlight lcl.IPicture // 按钮关闭图标移入高亮, 靠右
	isEnterClose       bool         // 鼠标是否移入关闭图标
	icon               lcl.IPicture // 按钮图标, 中间
	// 用户事件
	onCloseClick lcl.TNotifyEvent
	onPaint      lcl.TNotifyEvent
	onMouseEnter lcl.TNotifyEvent
	onMouseLeave lcl.TNotifyEvent
	onMouseDown  lcl.TMouseEvent
	onMouseUp    lcl.TMouseEvent
	// 默认颜色, 移入颜色, 按下颜色, 禁用颜色
	buttonState   TButtonState
	defaultColor  *TButtonColor
	enterColor    *TButtonColor
	downColor     *TButtonColor
	disabledColor *TButtonColor
	// 提示
	closeHintTimer *time.Timer
	closeHint      lcl.IHintWindow
	closeHintText  string
}

func NewButton(owner lcl.IComponent) *TButton {
	m := &TButton{ICustomGraphicControl: lcl.NewCustomGraphicControl(owner)}
	m.SetWidth(120)
	m.SetHeight(40)
	m.SetParentBackground(true)
	m.SetParentColor(true)
	m.Canvas().SetAntialiasingMode(types.AmOn)
	m.SetControlStyle(m.ControlStyle().Include(types.CsParentBackground))
	m.alpha = 255
	m.radius = 0
	m.ICustomGraphicControl.SetOnPaint(m.paint)
	m.ICustomGraphicControl.SetOnMouseEnter(m.Enter) // 进入
	m.ICustomGraphicControl.SetOnMouseLeave(m.Leave) // 移出
	m.ICustomGraphicControl.SetOnMouseDown(m.Down)   // 按下
	m.ICustomGraphicControl.SetOnMouseUp(m.Up)       // 抬起
	m.ICustomGraphicControl.SetOnMouseMove(m.move)
	m.RoundedCorner = types.NewSet(RcLeftTop, RcRightTop, RcLeftBottom, RcRightBottom)
	m.iconFavorite = lcl.NewPicture()
	m.iconClose = lcl.NewPicture()
	m.iconCloseHighlight = lcl.NewPicture()
	m.icon = lcl.NewPicture()
	m.iconFavorite.SetOnChange(m.iconChange)
	m.iconClose.SetOnChange(m.iconChange)
	m.iconCloseHighlight.SetOnChange(m.iconChange)
	m.icon.SetOnChange(m.iconChange)
	// 创建按钮颜色对象
	m.defaultColor = NewButtonColor()
	m.defaultColor.type_ = BsDefault
	m.enterColor = NewButtonColor()
	m.enterColor.type_ = BsEnter
	m.downColor = NewButtonColor()
	m.downColor.type_ = BsDown
	m.disabledColor = NewButtonColor()
	m.disabledColor.type_ = BsDisabled
	// 设置按钮颜色
	m.SetColor(defaultButtonColor)
	// 设置禁用颜色
	m.SetDisabledColor(defaultButtonColorDisable, defaultButtonColorDisable)
	// 启用边框
	m.SetBorderDirections(types.NewSet(BbdLeft, BbdTop, BbdRight, BbdBottom))
	// 边框宽度 1px
	m.SetBorderWidth(0, 1)

	m.closeHint = lcl.NewHintWindow(nil)
	// TODO WndProc
	//m.SetOnWndProc(func(theMessage *types.TLMessage) {
	//	m.InheritedWndProc(theMessage)
	//	fmt.Println(theMessage.Msg)
	//})
	// 销毁事件
	m.SetOnDestroy(func() {
		//fmt.Println("Graphic Button 释放资源")
		// 清空事件
		m.ICustomGraphicControl.SetOnPaint(nil)
		m.ICustomGraphicControl.SetOnMouseEnter(nil)
		m.ICustomGraphicControl.SetOnMouseLeave(nil)
		m.ICustomGraphicControl.SetOnMouseDown(nil)
		m.ICustomGraphicControl.SetOnMouseUp(nil)
		m.ICustomGraphicControl.SetOnMouseMove(nil)
		m.iconFavorite.SetOnChange(nil)
		m.iconClose.SetOnChange(nil)
		m.iconCloseHighlight.SetOnChange(nil)
		m.icon.SetOnChange(nil)
		m.SetOnDestroy(nil)
		// 释放持有资源
		m.iconFavorite.Free()
		m.iconClose.Free()
		m.iconCloseHighlight.Free()
		m.icon.Free()
		m.defaultColor.Free()
		m.enterColor.Free()
		m.downColor.Free()
		m.disabledColor.Free()
	})
	return m
}

func (m *TButton) SetCloseHintText(text string) {
	m.closeHintText = text
}

// ShowHint 显示按钮的提示信息
// text: 要显示的提示文本内容
func (m *TButton) ShowHint(text string) {
	if text == "" {
		return
	}
	if m.isEnterClose && m.closeHintTimer != nil {
		return
	}
	m.closeHintTimer = time.AfterFunc(time.Second/2, func() {
		if !m.isEnterClose {
			return
		}
		lcl.RunOnMainThreadAsync(func(id uint32) {
			if !m.isEnterClose {
				return
			}
			cursorPos := lcl.Mouse.CursorPos()
			hintRect := m.closeHint.CalcHintRect(0, text, 0)
			w, h := hintRect.Width(), hintRect.Height()
			hintRect.Left = cursorPos.X + 15
			hintRect.Top = cursorPos.Y + 15
			hintRect.SetWidth(w)
			hintRect.SetHeight(h)
			m.closeHint.ActivateHintWithRectStr(hintRect, text)
		})
	})
}

func (m *TButton) HideHint() {
	if m.closeHintTimer != nil {
		m.closeHintTimer.Stop()
		m.closeHintTimer = nil
		lcl.RunOnMainThreadAsync(func(id uint32) {
			m.closeHint.Hide()
		})
	}
}

func (m *TButton) Enter(sender lcl.IObject) {
	if m.isDisable || !m.IsValid() {
		return
	}
	m.buttonState = BsEnter
	m.Invalidate()
	if m.onMouseEnter != nil {
		m.onMouseEnter(sender)
	}
}

func (m *TButton) Leave(sender lcl.IObject) {
	if m.isDisable || !m.IsValid() {
		return
	}
	m.isEnterClose = false
	m.buttonState = BsDefault
	m.Invalidate()
	if m.onMouseLeave != nil {
		m.onMouseLeave(sender)
	}
}

func (m *TButton) Down(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, X int32, Y int32) {
	if m.isDisable || !m.IsValid() {
		return
	}
	m.HideHint()
	if !m.isCloseArea(X, Y) {
		m.buttonState = BsDown
		m.Invalidate()
		if m.onMouseDown != nil {
			m.onMouseDown(sender, button, shift, X, Y)
		}
	}
}

func (m *TButton) Up(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, X int32, Y int32) {
	if m.isDisable || !m.IsValid() {
		return
	}
	m.HideHint()
	if m.isCloseArea(X, Y) {
		if m.onCloseClick != nil {
			m.onCloseClick(sender)
		}
	} else {
		m.buttonState = BsEnter
		m.Invalidate()
		if m.onMouseUp != nil {
			m.onMouseUp(sender, button, shift, X, Y)
		}
	}
}

func (m *TButton) SetDisable(disable bool) {
	m.isDisable = disable
	if m.isDisable {
		m.buttonState = BsDisabled
	} else {
		m.buttonState = BsDefault
	}
	lcl.RunOnMainThreadAsync(func(id uint32) {
		m.Invalidate()
	})
}
func (m *TButton) iconChange(sender lcl.IObject) {
	if m.isDisable || !m.IsValid() {
		return
	}
	m.Invalidate()
}

func (m *TButton) isCloseArea(X int32, Y int32) bool {
	if m.isDisable || !m.IsValid() {
		return false
	}
	btnRect := m.ClientRect()
	closeW := m.iconClose.Width()
	closeH := m.iconClose.Height()
	closeX := btnRect.Width() - closeW - iconMargin
	closeY := btnRect.Height()/2 - closeH/2
	return X >= closeX && X <= btnRect.Width()-iconMargin && Y >= closeY && Y <= btnRect.Height()/2+closeH/2
}

func (m *TButton) move(sender lcl.IObject, shift types.TShiftState, X int32, Y int32) {
	if m.isDisable || !m.IsValid() {
		return
	}
	lcl.Screen.SetCursor(types.CrDefault)
	if m.isCloseArea(X, Y) {
		m.ShowHint(m.closeHintText)
		if !m.isEnterClose {
			m.isEnterClose = true
			m.Invalidate()
		}
		return
	} else if m.isEnterClose {
		m.isEnterClose = false
		m.Invalidate()
	}
	m.HideHint()
}

func (m *TButton) drawRoundedGradientButton(canvas lcl.ICanvas, rect types.TRect) {
	text := m.text
	var color *TButtonColor
	switch m.buttonState {
	case BsDefault:
		color = m.defaultColor
	case BsEnter:
		color = m.enterColor
	case BsDown:
		color = m.downColor
	case BsDisabled:
		color = m.disabledColor
	}
	if color == nil {
		return
	}
	color.tryPaint(m.RoundedCorner, rect, m.alpha, m.radius)

	// 绘制到目标画布
	canvas.DrawWithIntX2Graphic(rect.Left, rect.Top, color.bitMap)

	// 绘制按钮文字（在原始画布上绘制，确保文字不透明）
	brush := canvas.BrushToBrush()
	brush.SetStyle(types.BsClear)

	textMargin := int32(0) // 文本与图标的间距
	// 计算左图标占用的空间
	leftArea := int32(0)
	if m.iconFavorite.Width() > 0 {
		leftArea = iconMargin + m.iconFavorite.Width() + iconMargin // 左边距10 + 图标宽度 + 图标与文本间距10
		textMargin += iconMargin
	}
	// 计算右图标占用的空间
	rightArea := int32(0)
	if m.iconClose.Width() > 0 {
		rightArea = iconMargin + m.iconClose.Width() + iconMargin // 右边距10 + 图标宽度 + 图标与文本间距10
		textMargin += -iconMargin
	}

	// 计算文本可用宽度
	availWidth := rect.Width() - leftArea - rightArea
	if availWidth < 0 {
		availWidth = 0
	}

	lines := strings.Split(text, "\n")

	// 逐行处理：截断每行超长文本
	var processedLines []string
	var lineHeight int32 // 单行文本高度（默认取第一行高度，假设字体统一）
	// 获取单行文本高度
	tempSize := canvas.TextExtentWithStr(lines[0])
	lineHeight = tempSize.Cy

	// 逐行截断超长文本
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		// 截断 确保每行不超过可用宽度
		truncatedLine := truncateText(canvas, line, availWidth)
		processedLines = append(processedLines, truncatedLine)
	}
	lines = processedLines

	// 计算多行文本的整体位置（保持垂直居中）
	totalTextHeight := int32(len(lines)) * lineHeight // 总文本高度（无行间距）
	// 添加行间距
	totalTextHeight = totalTextHeight + (int32(len(lines))-1)*m.TextLineSpacing

	// 文本区域起始Y坐标（垂直居中）
	startY := rect.Top + m.TextOffSetY + (rect.Height()-totalTextHeight)/2
	textBaseX := rect.Left + m.TextOffSetX + textMargin
	for i, line := range lines {
		lineSize := canvas.TextExtentWithStr(line)
		var textX int32
		switch m.TextAlign {
		case TextAlignRight:
			textX = textBaseX + leftArea + (availWidth - lineSize.Cx)
		case TextAlignLeft:
			textX = textBaseX + leftArea
		default:
			textX = textBaseX + (rect.Width()-lineSize.Cx)/2
		}
		textY := startY + int32(i)*(lineHeight+m.TextLineSpacing) + (lineHeight-lineSize.Cy)/2
		canvas.TextOutWithIntX2Str(textX, textY, line)
	}

	// 左: 绘制图标 favorite
	favY := rect.Height()/2 - m.iconFavorite.Height()/2
	canvas.DrawWithIntX2Graphic(iconMargin, favY, m.iconFavorite.Graphic())

	// 右: 绘制图标 close
	iconClose := m.iconClose
	if m.isEnterClose {
		iconClose = m.iconCloseHighlight
	}
	closeX := rect.Width() - iconClose.Width() - iconMargin
	closeY := rect.Height()/2 - iconClose.Height()/2
	canvas.DrawWithIntX2Graphic(closeX, closeY, iconClose.Graphic())

	// 中间: 绘制图标 icon
	iconW, iconH := m.icon.Width(), m.icon.Height()
	iconX := rect.Left + (rect.Width()-iconW)/2
	iconY := rect.Top + (rect.Height()-iconH)/2
	canvas.DrawWithIntX2Graphic(iconX, iconY, m.icon.Graphic())
}

func (m *TButton) Disable() bool {
	return m.isDisable
}

func (m *TButton) SetCaption(value string) {
	m.SetText(value)
}

func (m *TButton) Caption() string {
	return m.text
}

func (m *TButton) SetText(value string) {
	m.text = value
	m.AutoSizeWidth()
}

// 自动大小, 根据文本宽自动调整按钮宽度
func (m *TButton) AutoSizeWidth() {
	if m.autoSize {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			if m.Canvas() != nil {
				leftArea := int32(0)
				if m.iconFavorite.Width() > 0 {
					leftArea = iconMargin + m.iconFavorite.Width() + iconMargin
				}
				rightArea := int32(0)
				if m.iconClose.Width() > 0 {
					rightArea = iconMargin + m.iconClose.Width() + iconMargin
				}
				textWidth := m.Canvas().TextWidthWithStr(m.text)
				width := textWidth + leftArea + rightArea + iconMargin*2
				if m.Width() != width {
					m.SetWidth(width)
				}
			}
		})
	} else {
		m.Invalidate()
	}
}

func (m *TButton) Text() string {
	return m.text
}

// SetAutoSize 设置按钮的自动大小属性
//
//	当启用自动大小时，按钮会根据其内容自动调整大小
//	Note: 当前需要在第一次设置文本之前设置生效
func (m *TButton) SetAutoSize(v bool) {
	m.autoSize = v
}

func (m *TButton) SetIconFavorite(filePath string) {
	if !m.IsValid() {
		return
	}
	m.iconFavorite.LoadFromFile(filePath)
}

func (m *TButton) SetIconFavoriteFormBytes(pngData []byte) {
	if !m.IsValid() {
		return
	}
	if pngData == nil {
		m.iconFavorite.Clear()
		return
	}
	mem := lcl.NewMemoryStream()
	defer mem.Free()
	lcl.StreamHelper.WriteBuffer(mem, pngData)
	mem.SetPosition(0)
	m.iconFavorite.LoadFromStream(mem)
}

func (m *TButton) SetIcon(filePath string) {
	if !m.IsValid() {
		return
	}
	m.icon.LoadFromFile(filePath)
}

func (m *TButton) SetIconFormBytes(pngData []byte) {
	if !m.IsValid() {
		return
	}
	if pngData == nil {
		m.icon.Clear()
		return
	}
	mem := lcl.NewMemoryStream()
	defer mem.Free()
	lcl.StreamHelper.WriteBuffer(mem, pngData)
	mem.SetPosition(0)
	m.icon.LoadFromStream(mem)
}

func (m *TButton) SetIconClose(filePath string) {
	if !m.IsValid() {
		return
	}
	path, name := filepath.Split(filePath)
	ns := strings.Split(name, ".")
	enterFilePath := filepath.Join(path, ns[0]+"_enter.png")
	m.iconClose.LoadFromFile(filePath)
	m.SetIconCloseHighlight(enterFilePath)
}

func (m *TButton) SetIconCloseHighlight(filePath string) {
	if !m.IsValid() {
		return
	}
	m.iconCloseHighlight.LoadFromFile(filePath)
}

func (m *TButton) SetIconCloseFormBytes(pngData []byte) {
	if !m.IsValid() {
		return
	}
	if pngData == nil {
		m.iconClose.Clear()
		return
	}
	mem := lcl.NewMemoryStream()
	defer mem.Free()
	lcl.StreamHelper.WriteBuffer(mem, pngData)
	mem.SetPosition(0)
	m.iconClose.LoadFromStream(mem)

}

func (m *TButton) SetIconCloseHighlightFormBytes(pngData []byte) {
	if !m.IsValid() {
		return
	}
	if pngData == nil {
		m.iconCloseHighlight.Clear()
		return
	}
	mem := lcl.NewMemoryStream()
	defer mem.Free()
	lcl.StreamHelper.WriteBuffer(mem, pngData)
	mem.SetPosition(0)
	m.iconCloseHighlight.LoadFromStream(mem)
}

// 绘制事件
func (m *TButton) paint(sender lcl.IObject) {
	if !m.IsValid() {
		return
	}
	canvas := m.Canvas()
	if canvas == nil || !canvas.IsValid() {
		return
	}
	m.drawRoundedGradientButton(canvas, m.ClientRect())
	if m.onPaint != nil {
		m.onPaint(sender)
	}
}
func (m *TButton) SetOnCloseClick(fn lcl.TNotifyEvent) {
	m.onCloseClick = fn
}

func (m *TButton) SetOnPaint(fn lcl.TNotifyEvent) {
	m.onPaint = fn
}

func (m *TButton) SetOnMouseDown(fn lcl.TMouseEvent) {
	m.onMouseDown = fn
}

func (m *TButton) SetOnMouseUp(fn lcl.TMouseEvent) {
	m.onMouseUp = fn
}

func (m *TButton) SetOnMouseEnter(fn lcl.TNotifyEvent) {
	m.onMouseEnter = fn
}

func (m *TButton) SetOnMouseLeave(fn lcl.TNotifyEvent) {
	m.onMouseLeave = fn
}

// SetDefaultColor 设置按钮的默认颜色
// start: 按钮默认状态下的起始颜色
// end: 按钮默认状态下的结束颜色
func (m *TButton) SetDefaultColor(start, end colors.TColor) {
	// 更新按钮默认颜色配置
	m.defaultColor.start = start
	m.defaultColor.end = end
	m.defaultColor.canPaint = true
}

func (m *TButton) DefaultColor() (start, end colors.TColor) {
	start = m.defaultColor.start
	end = m.defaultColor.end
	return
}

// SetEnterColor 设置按钮进入状态时的颜色渐变效果
// start: 渐变开始颜色
// end: 渐变结束颜色
func (m *TButton) SetEnterColor(start, end colors.TColor) {
	m.enterColor.start = start
	m.enterColor.end = end
	m.enterColor.canPaint = true
}

func (m *TButton) EnterColor() (start, end colors.TColor) {
	start = m.enterColor.start
	end = m.enterColor.end
	return
}

// SetDownColor 设置按钮按下状态时的颜色渐变效果
// start: 按下状态渐变起始颜色
// end: 按下状态渐变结束颜色
func (m *TButton) SetDownColor(start, end colors.TColor) {
	m.downColor.start = start
	m.downColor.end = end
	m.downColor.canPaint = true
}

func (m *TButton) DownColor() (start, end colors.TColor) {
	start = m.downColor.start
	end = m.downColor.end
	return
}

// SetDisabledColor 设置按钮禁用状态时的渐变颜色
// start: 渐变起始颜色
// end: 渐变结束颜色
func (m *TButton) SetDisabledColor(start, end colors.TColor) {
	m.disabledColor.start = start
	m.disabledColor.end = end
	m.disabledColor.canPaint = true
}

// DisabledColor 返回禁用颜色
func (m *TButton) DisabledColor() (start, end colors.TColor) {
	start = m.disabledColor.start
	end = m.disabledColor.end
	return
}

// SetColor 设置按钮的颜色渐变为同一颜色
func (m *TButton) SetColor(color colors.TColor) {
	m.SetColorGradient(color, color)
}

// SetColorGradient 设置按钮的颜色渐变效果
// start: 渐变起始颜色
// end: 渐变结束颜色
func (m *TButton) SetColorGradient(start, end colors.TColor) {
	m.SetDefaultColor(start, end)
	m.SetEnterColor(DarkenColor(start, 0.1), DarkenColor(end, 0.1))
	m.SetDownColor(DarkenColor(start, 0.2), DarkenColor(end, 0.2))
}

// SetBorderColor 设置按钮所有状态下的边框颜色
//
//	color - 指定的边框颜色值
//	该函数会同时设置按钮在默认、悬停、按下和禁用状态下的边框颜色
//	为统一的颜色值，实现按钮边框颜色的整体变更
func (m *TButton) SetBorderColor(direction TButtonBorderDirection, color colors.TColor) {
	m.defaultColor.SetBorderColor(direction, color)
	m.enterColor.SetBorderColor(direction, DarkenColor(color, 0.1))
	m.downColor.SetBorderColor(direction, DarkenColor(color, 0.2))
}

// SetBorderWidth 设置按钮的边框宽度
// width: 边框宽度值
func (m *TButton) SetBorderWidth(direction TButtonBorderDirection, width int32) {
	m.defaultColor.SetBorderWidth(direction, width)
	m.enterColor.SetBorderWidth(direction, width)
	m.downColor.SetBorderWidth(direction, width)
}

// SetBorderDirections 设置按钮的所有状态边框样式
//
//	borders: TButtonBorders类型，指定按钮的边框样式
//
// 该函数会同时设置按钮的默认、悬停、按下和禁用四种状态的边框样式
// 为相同的值，实现统一的边框外观效果
func (m *TButton) SetBorderDirections(directions TButtonBorderDirections) {
	m.defaultColor.Border.Direction = directions
	m.enterColor.Border.Direction = directions
	m.downColor.Border.Direction = directions
	m.disabledColor.Border.Direction = directions
	m.defaultColor.canPaint = true
	m.enterColor.canPaint = true
	m.downColor.canPaint = true
	m.disabledColor.canPaint = true
}

func (m *TButton) SetAlpha(alpha byte) {
	m.alpha = alpha
}

func (m *TButton) SetRadius(radius int32) {
	m.radius = radius
}

func (m *TButton) Free() {
	m.ICustomGraphicControl.Free()
}

// 文本截断函数（添加在文本末尾）
func truncateText(canvas lcl.ICanvas, text string, maxWidth int32) string {
	if maxWidth <= 0 {
		return ""
	}
	ellipsis := "..."
	ellipsisWidth := canvas.GetTextWidth(ellipsis)
	if ellipsisWidth > maxWidth {
		return ""
	}
	textWidth := canvas.GetTextWidth(text)
	if textWidth <= maxWidth {
		return text
	}
	// 二分查找截断位置
	runes := []rune(text)
	left, right := 0, len(runes)
	for left < right {
		mid := (left + right) / 2
		truncated := string(runes[:mid]) + ellipsis
		if canvas.GetTextWidth(truncated) <= maxWidth {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left == 0 {
		return ellipsis
	}
	return string(runes[:left-1]) + ellipsis
}
