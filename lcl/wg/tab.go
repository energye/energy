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
	"github.com/energye/widget/assets"
	"strconv"
	"time"
)

type TCloseEvent func(page *TPage, canClose *bool)

var (
	activeColor     = colors.RGBToColor(60, 70, 80)
	defaultColor    = colors.RGBToColor(86, 88, 100)
	defaultPrefix   = "Tab"
	defaultHeight   = int32(25)
	scrollBtnWidth  = int32(20)
	scrollBtnHeight = int32(25)
	scrollBtnMargin = int32(4)
	scrollStep      = int32(15)
)

type TTab struct {
	lcl.ICustomPanel                   //
	pages             []*TPage         // 页列表
	totalTabWidth     int32            // 页签总宽度
	deleting          bool             // 正在删除中 page
	scrollLeftBtn     *TButton         // tab 滚动导航按钮 左滚动
	scrollRightBtn    *TButton         // tab 滚动导航按钮 右滚动
	scrollOffset      int32            // tab 滚动导航按钮 偏移坐标
	scrollTimer       *time.Timer      // tab 滚动连续
	triggerScrollStop bool             // 触发滚动是否停止
	onChange          lcl.TNotifyEvent //
	Margin            int32            // 两个 tab 之间的距离
}

type TPage struct {
	lcl.ICustomPanel
	tabSheet     lcl.ICustomPage
	active       bool     // 是否激活
	show         bool     // 是否显示
	tab          *TTab    // 所属的tab
	button       *TButton // 按钮
	onShow       lcl.TNotifyEvent
	onHide       lcl.TNotifyEvent
	onClose      TCloseEvent
	onClick      lcl.TNotifyEvent
	activeColor  types.TColor //
	defaultColor types.TColor //
}

func NewTab(owner lcl.IComponent) *TTab {
	tab := &TTab{}
	tab.ICustomPanel = lcl.NewCustomPanel(owner)
	tab.SetBevelInner(types.BvNone)
	tab.SetBevelOuter(types.BvNone)
	//tab.SetColor(colors.ClRed)
	tab.SetBorderStyleToBorderStyle(types.BsNone)
	tab.initScrollBtn()
	return tab
}

// 初始化滚动导航按钮
func (m *TTab) initScrollBtn() {
	m.scrollLeftBtn = NewButton(m)
	m.scrollRightBtn = NewButton(m)

	m.scrollLeftBtn.SetIconFormBytes(assets.Tab("scroll-left.png"))
	m.scrollLeftBtn.SetWidth(scrollBtnWidth)
	m.scrollLeftBtn.SetHeight(scrollBtnHeight)
	m.scrollLeftBtn.SetLeft(2)
	//m.scrollLeftBtn.SetTop(2)
	m.scrollLeftBtn.SetRadius(1)
	m.scrollLeftBtn.SetBorderDirections(types.NewSet())
	m.scrollLeftBtn.SetColor(LightenColor(colors.ClGray, 0.2))
	m.scrollLeftBtn.SetParent(m)

	m.scrollRightBtn.SetIconFormBytes(assets.Tab("scroll-right.png"))
	m.scrollRightBtn.SetWidth(scrollBtnWidth)
	m.scrollRightBtn.SetHeight(scrollBtnHeight)
	//m.scrollRightBtn.SetTop(2)
	m.scrollRightBtn.SetRadius(1)
	m.scrollRightBtn.SetBorderDirections(types.NewSet())
	m.scrollRightBtn.SetColor(LightenColor(colors.ClGray, 0.2))
	m.scrollRightBtn.SetParent(m)

	scrollBtnMouseUp := func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, X int32, Y int32) {
		m.triggerScrollStop = true
	}
	m.scrollLeftBtn.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, X int32, Y int32) {
		m.triggerScrollStop = false
		m.triggerScrollLoop(time.Second/2, 1)
	})
	m.scrollLeftBtn.SetOnMouseUp(scrollBtnMouseUp)
	m.scrollRightBtn.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, X int32, Y int32) {
		m.triggerScrollStop = false
		m.triggerScrollLoop(time.Second/2, 2)
	})
	m.scrollRightBtn.SetOnMouseUp(scrollBtnMouseUp)
}

func (m *TTab) NewPage() *TPage {
	page := new(TPage)
	page.tab = m
	page.activeColor = activeColor
	page.defaultColor = defaultColor
	button := NewButton(m)
	//button.SetAutoSize(true)
	//button.SetShowHint(true)
	button.SetCaption(defaultPrefix + strconv.Itoa(len(m.pages)))
	button.Font().SetSize(9)
	button.Font().SetColor(colors.Cl3DFace)
	button.RoundedCorner = button.RoundedCorner.Exclude(RcLeftBottom).Exclude(RcRightBottom)
	button.SetRadius(0)
	button.SetAlpha(255)
	button.SetHeight(defaultHeight)
	button.SetDefaultColor(defaultColor, defaultColor)
	button.SetEnterColor(DarkenColor(defaultColor, 0.1), DarkenColor(defaultColor, 0.1))
	button.SetDownColor(DarkenColor(defaultColor, 0.2), DarkenColor(defaultColor, 0.2))
	button.SetBorderColor(BbdNone, DarkenColor(defaultColor, 0.3))
	button.SetParent(m)
	page.button = button

	tabRect := m.ClientRect()
	sheet := lcl.NewCustomPanel(m)
	sheet.SetBevelInner(types.BvNone)
	sheet.SetBevelOuter(types.BvNone)
	sheet.SetBorderStyleToBorderStyle(types.BsNone)
	sheet.SetTop(button.Height())
	sheet.SetHeight(tabRect.Height() - button.Height())
	sheet.SetWidth(tabRect.Width())
	sheet.SetAlign(types.AlCustom)
	sheet.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight, types.AkBottom))
	sheet.SetParent(m)
	page.ICustomPanel = sheet

	tabSheet := lcl.NewCustomPage(m)
	tabSheet.SetParent(sheet)
	page.tabSheet = tabSheet

	m.pages = append(m.pages, page) // 添加到页列表
	page.initEvent()                // 初始化事件
	page.SetActive(false)

	// 事件处理
	tabSheet.SetOnShow(func(sender lcl.IObject) {
		if m.onChange != nil {
			m.onChange(page)
		}
		if page.onShow != nil {
			page.onShow(page)
		}
	})
	tabSheet.SetOnHide(func(sender lcl.IObject) {
		if page.onHide != nil {
			page.onHide(sender)
		}
	})
	return page
}

func (m *TTab) SetOnChange(fn lcl.TNotifyEvent) {
	m.onChange = fn
}

func (m *TTab) ScrollLeft() *TButton {
	return m.scrollLeftBtn
}

func (m *TTab) ScrollRight() *TButton {
	return m.scrollRightBtn
}

func (m *TTab) EnableScrollButton(value bool) {
	if m.scrollLeftBtn != nil {
		m.scrollLeftBtn.SetVisible(value)
	}
	if m.scrollRightBtn != nil {
		m.scrollRightBtn.SetVisible(value)
	}
}

func (m *TTab) triggerScrollLoop(afterTime time.Duration, scrollLeftOrRight int32) {
	if m.triggerScrollStop {
		if m.scrollTimer != nil {
			m.scrollTimer.Stop()
			m.scrollTimer = nil
		}
		return
	}
	lcl.RunOnMainThreadAsync(func(id uint32) {
		if scrollLeftOrRight == 1 {
			m.scrollLeft()
		} else {
			m.scrollRight()
		}
	})
	m.scrollTimer = time.AfterFunc(afterTime, func() {
		m.triggerScrollLoop(time.Second/30, scrollLeftOrRight)
	})
}

// 向左滚动
func (m *TTab) scrollLeft() {
	scrollLeft := int32(0)
	if m.scrollLeftBtn.Visible() {
		scrollLeft = scrollBtnWidth + scrollBtnMargin
	}
	if m.scrollOffset+scrollLeft < scrollLeft {
		m.scrollOffset += scrollStep
		m.RecalculatePosition()
	} else {
		m.triggerScrollStop = true
	}
}

// 向右滚动
func (m *TTab) scrollRight() {
	width := m.Width()
	widths := m.totalTabWidth + scrollBtnWidth + scrollBtnMargin
	if widths > width {
		m.scrollOffset += -scrollStep
		m.RecalculatePosition()
	} else {
		m.triggerScrollStop = true
	}
}

// RecalculatePosition 重新计算位置, 在隐藏/移除时使用
func (m *TTab) RecalculatePosition() {
	widths := m.scrollOffset + m.Margin
	if m.scrollLeftBtn != nil && m.scrollLeftBtn.Visible() {
		widths += scrollBtnWidth + scrollBtnMargin
	}
	for _, page := range m.pages {
		if page.button.Visible() {
			page.button.AutoSizeWidth()
			br := page.button.BoundsRect()
			width := br.Width()
			br.Left = widths
			br.SetWidth(width)
			page.button.SetBoundsRect(br)
			widths += br.Width() + m.Margin
		}
	}
	m.totalTabWidth = widths
	// 滚动导航按钮 位置调整
	m.scrollBtnPosition()
}

// 滚动导航按钮 位置调整
func (m *TTab) scrollBtnPosition() {
	if m.scrollLeftBtn != nil && m.scrollLeftBtn.Visible() {
		m.scrollLeftBtn.SetLeft(2)
		m.scrollLeftBtn.BringToFront()
	}
	if m.scrollRightBtn != nil && m.scrollRightBtn.Visible() {
		m.scrollRightBtn.SetLeft(m.Width() - scrollBtnWidth - 2)
		m.scrollRightBtn.BringToFront()
	}
}

// HideAllActivated 隐藏所有激活页面
func (m *TTab) HideAllActivated() {
	for _, page := range m.pages {
		if page.active {
			page.SetActive(false)
		}
	}
}

// 删除指定 page
func (m *TTab) RemovePage(removePage *TPage) {
	removeIndex := -1 // 存放当前删除page的索引
	var (
		newPages []*TPage // 替换为实际元素类型
		found    bool
	)
	for i, page := range m.pages {
		if page == removePage {
			removeIndex = i
			newPages = make([]*TPage, 0, len(m.pages)-1)
			newPages = append(newPages, m.pages[:i]...)
			newPages = append(newPages, m.pages[i+1:]...)
			m.pages = newPages //append(m.pages[:i], m.pages[i+1:]...)
			found = true
			break
		}
	}
	// 重新计算 button 位置
	m.RecalculatePosition()
	if removePage.active && found {
		// 根据删除索引获取要显示的 page
		var showPage *TPage
		if removeIndex != -1 && removeIndex < len(m.pages) {
			showPage = m.pages[removeIndex] // 显示当前索引的 page, 也就是删除后的下一个
		} else if len(m.pages) > 0 {
			showPage = m.pages[0] // 显示第一个 page
		}
		if showPage != nil {
			lcl.RunOnMainThreadAsync(func(id uint32) {
				if showPage.IsValid() {
					showPage.SetActive(true)
				}
			})
		}
	}
	m.deleting = false
}

func (m *TTab) Pages() []*TPage {
	return m.pages
}

func (m *TPage) Free() {
	m.button.Free()
	m.ICustomPanel.Free()
	m.tabSheet.Free()
}

// 删除掉自己
func (m *TPage) Remove() {
	m.button.SetOnClick(nil)
	m.button.SetOnCloseClick(nil)
	// 先隐藏掉
	m.button.Hide()
	m.ICustomPanel.Hide()
	m.tabSheet.Hide()
	// 在page里删除自己
	m.tab.RemovePage(m)
	// 最后释放掉
	m.Free()
	m.tab = nil
}

func (m *TPage) SetCaption(name string) {
	m.button.SetCaption(name)
	lcl.RunOnMainThreadAsync(func(id uint32) {
		m.tab.RecalculatePosition()
	})
}

func (m *TPage) SetActiveColor(color types.TColor) {
	m.activeColor = color
}
func (m *TPage) SetDefaultColor(color types.TColor) {
	m.defaultColor = color
}

func (m *TPage) Active() bool {
	return m.active
}

// 激活自己, 会取消其它激活的
func (m *TPage) SetActive(active bool) {
	m.active = active
	if active {
		m.button.SetDefaultColor(m.activeColor, m.activeColor)
		m.ICustomPanel.Show()
		m.tabSheet.Show()
	} else {
		m.button.SetDefaultColor(m.defaultColor, m.defaultColor)
		m.ICustomPanel.Hide()
		m.tabSheet.Hide()
	}
	m.button.Invalidate()
}

// 隐藏自己, button 和 page 同时隐藏
func (m *TPage) Hide() {
	m.button.Hide()
	m.SetActive(false)
	m.tab.RecalculatePosition()
}

// 显示自己, button 和 page 同时显示
func (m *TPage) Show() {
	m.button.Show()
	m.SetActive(true)
	m.tab.RecalculatePosition()
}

func (m *TPage) SetOnShow(fn lcl.TNotifyEvent) {
	m.onShow = fn
}

func (m *TPage) SetOnHide(fn lcl.TNotifyEvent) {
	m.onHide = fn
}

func (m *TPage) SetOnClose(fn TCloseEvent) {
	m.onClose = fn
}

func (m *TPage) SetOnClick(fn lcl.TNotifyEvent) {
	m.onClick = fn
}

// 是否进入关闭按钮
func (m *TPage) IsEnterClose() bool {
	return m.button.isEnterClose
}

func (m *TPage) initEvent() {
	m.button.SetOnClick(func(sender lcl.IObject) {
		// 关闭按钮位置, 不触发事件
		if m.IsEnterClose() {
			return
		}
		m.tab.HideAllActivated()
		m.SetActive(true)
		if m.onClick != nil {
			m.onClick(sender)
		}
	})
	m.button.SetOnCloseClick(func(sender lcl.IObject) {
		if m.tab.deleting {
			return
		}
		m.tab.deleting = true
		m.Close()
	})
	m.SetOnResize(func(sender lcl.IObject) {
		// 滚动导航按钮 位置调整
		m.tab.scrollBtnPosition()
	})
}

func (m *TPage) Button() *TButton {
	return m.button
}

func (m *TPage) Close() {
	m.button.isEnterClose = true
	lcl.RunOnMainThreadAsync(func(id uint32) {
		var canClose = true
		if m.onClose != nil {
			m.onClose(m, &canClose)
		}
		if canClose {
			m.Remove()
		} else {
			m.tab.deleting = false
		}
	})
}
