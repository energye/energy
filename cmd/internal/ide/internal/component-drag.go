package internal

import (
	"fmt"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"time"
)

func (m *IDEComponent) MouseMove(sender lcl.IObject, shift types.TShiftState, x, y int32) {
	if m.IsDown {
		var borderRect types.TRect
		if m.BorderPanel != nil {
			borderRect = m.BorderPanel.BoundsRect()
		}
		componentParentRect := m.ParentToControl().BoundsRect()
		if m.IsBorder && m.ComponentType == CtForm { //mouse down BorderMargin > resize
			switch m.BorderHT {
			case HTRIGHT:
				tmpWidth := m.Ow + (x - m.Dx)
				if tmpWidth <= MinW {
					return
				}
				if m.BorderPanel != nil {
					m.BorderPanel.SetWidth(tmpWidth + Border)
				}
				m.ParentToControl().SetWidth(tmpWidth)
			case HTLEFT:
				tmpX := m.ParentToControl().Left() + (x - m.Dx)
				tmpWidth := m.Ow + (m.Ox - tmpX)
				if tmpWidth <= MinW {
					return
				}
				if m.BorderPanel != nil {
					m.BorderPanel.SetLeft(tmpX - Border/2)
					m.BorderPanel.SetWidth(tmpWidth + Border)
				}
				m.ParentToControl().SetLeft(tmpX)
				m.ParentToControl().SetWidth(tmpWidth)
			case HTTOP:
				tmpY := m.ParentToControl().Top() + (y - m.Dy)
				tmpHeight := m.Oh + (m.Oy - tmpY)
				if tmpHeight <= MinH {
					return
				}
				if m.BorderPanel != nil {
					m.BorderPanel.SetTop(tmpY - Border/2)
					m.BorderPanel.SetHeight(tmpHeight + Border)
				}
				m.ParentToControl().SetTop(tmpY)
				m.ParentToControl().SetHeight(tmpHeight)
			case HTBOTTOM:
				tmpHeight := m.Oh + (y - m.Dy)
				if tmpHeight <= MinH {
					return
				}
				if m.BorderPanel != nil {
					m.BorderPanel.SetHeight(tmpHeight + Border)
				}
				m.ParentToControl().SetHeight(m.Oh + (y - m.Dy))
			case HTTOPRIGHT:
				tmpY := componentParentRect.Top + (y - m.Dy)
				tmpHeight := m.Oh + (m.Oy - tmpY)
				tmpWidth := m.Ow + (x - m.Dx)
				if tmpWidth <= MinW || tmpHeight <= MinH {
					return
				}
				if m.BorderPanel != nil {
					m.BorderPanel.SetBounds(borderRect.Left, tmpY-Border/2, tmpWidth+Border, tmpHeight+Border)
				}
				m.ParentToControl().SetBounds(componentParentRect.Left, tmpY, tmpWidth, tmpHeight)
			case HTBOTTOMRIGHT:
				tmpWidth := m.Ow + (x - m.Dx)
				tmpHeight := m.Oh + (y - m.Dy)
				if tmpWidth <= MinW || tmpHeight <= MinH {
					return
				}
				if m.BorderPanel != nil {
					m.BorderPanel.SetBounds(borderRect.Left, borderRect.Top, tmpWidth+Border, tmpHeight+Border)
				}
				m.ParentToControl().SetBounds(componentParentRect.Left, componentParentRect.Top, tmpWidth, tmpHeight)
			case HTTOPLEFT:
				tmpX := componentParentRect.Left + (x - m.Dx)
				tmpWidth := m.Ow + (m.Ox - tmpX)
				tmpY := componentParentRect.Top + (y - m.Dy)
				tmpHeight := m.Oh + (m.Oy - tmpY)
				if tmpWidth <= MinW || tmpHeight <= MinH {
					return
				}
				if m.BorderPanel != nil {
					m.BorderPanel.SetBounds(tmpX-Border/2, tmpY-Border/2, tmpWidth+Border, tmpHeight+Border)
				}
				m.ParentToControl().SetBounds(tmpX, tmpY, tmpWidth, tmpHeight)
			case HTBOTTOMLEFT:
				tmpX := componentParentRect.Left + (x - m.Dx)
				tmpWidth := m.Ow + (m.Ox - tmpX)
				tmpHeight := m.Oh + (y - m.Dy)
				if tmpWidth <= MinW || tmpHeight <= MinH {
					return
				}
				if m.BorderPanel != nil {
					m.BorderPanel.SetBounds(tmpX-Border/2, borderRect.Top, tmpWidth+Border, tmpHeight+Border)
				}
				m.ParentToControl().SetBounds(tmpX, componentParentRect.Top, tmpWidth, tmpHeight)
			}
			return
		} else if m.IsComponentArea && m.ComponentType != CtForm { // mouse down Component area > move
			m.IsDClick = false
			tmpY := componentParentRect.Top + (y - m.Dy)
			tmpX := componentParentRect.Left + (x - m.Dx)
			if m.BorderPanel != nil {
				m.BorderPanel.SetBounds(tmpX-Border/2, tmpY-Border/2, borderRect.Width(), borderRect.Height())
			}
			m.ParentToControl().SetBounds(tmpX, tmpY, componentParentRect.Width(), componentParentRect.Height())
			return
		}
	}
	if m.ComponentType == CtForm {
		if m.IsBorder = x <= m.Ow && x >= m.Ow-BorderRange && y <= BorderRange; m.IsBorder && m.ComponentType != CtForm { // 右上
			m.ParentToControl().SetCursor(types.CrSizeSW)
			m.BorderHT = HTTOPRIGHT
		} else if m.IsBorder = x <= m.Ow && x >= m.Ow-BorderRange && y <= m.Oh && y >= m.Oh-BorderRange; m.IsBorder { // 右下
			m.ParentToControl().SetCursor(types.CrSizeSE)
			m.BorderHT = HTBOTTOMRIGHT
		} else if m.IsBorder = x <= BorderRange && y <= BorderRange; m.IsBorder && m.ComponentType != CtForm { //左上
			m.ParentToControl().SetCursor(types.CrSizeSE)
			m.BorderHT = HTTOPLEFT
		} else if m.IsBorder = x <= BorderRange && y >= m.Oh-BorderRange; m.IsBorder && m.ComponentType != CtForm { //左下
			m.ParentToControl().SetCursor(types.CrSizeSW)
			m.BorderHT = HTBOTTOMLEFT
		} else if m.IsBorder = x <= m.Ow && x >= m.Ow-BorderRange && y > BorderRange && y < m.Oh-BorderRange; m.IsBorder { //右
			m.ParentToControl().SetCursor(types.CrSizeW)
			m.BorderHT = HTRIGHT
		} else if m.IsBorder = x <= BorderRange && y > BorderRange && y < m.Oh-BorderRange; m.IsBorder && m.ComponentType != CtForm { //左
			m.ParentToControl().SetCursor(types.CrSizeW)
			m.BorderHT = HTLEFT
		} else if m.IsBorder = x > BorderRange && x < m.Ow-BorderRange && y <= BorderRange; m.IsBorder && m.ComponentType != CtForm { //上
			m.ParentToControl().SetCursor(types.CrSizeN)
			m.BorderHT = HTTOP
		} else if m.IsBorder = x > BorderRange && x < m.Ow-BorderRange && y >= m.Oh-BorderRange; m.IsBorder { //下
			m.ParentToControl().SetCursor(types.CrSizeN)
			m.BorderHT = HTBOTTOM
		} else {
			m.IsBorder = false
			m.ParentToControl().SetCursor(types.CrDefault)
		}
	}
	if m.Component != nil {
		switch m.Component.(type) {
		case lcl.IControl:
			m.Component.(lcl.IControl).SetCursor(m.ParentToControl().Cursor())
		default:
			m.ParentToControl().SetCursor(m.ParentToControl().Cursor())
		}
	}
}

func (m *IDEComponent) MouseDown(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if button == types.MbLeft {
		if time.Now().UnixMilli()-m.ClickTime.UnixMilli() < 500 {
			m.IsDClick = true
		} else {
			m.IsDClick = false
		}
		m.ClickTime = time.Now()
		m.Dx = x
		m.Dy = y
		if m.ComponentType == CtForm {
			Ide.forms[m.Id].Active.clearBorderColor()
			Ide.forms[m.Id].Active = nil
		} else {
			m.switchActive(m)
		}
		if !m.IsBorder && m.ComponentType != CtForm {
			m.IsComponentArea = true
			m.Anchor.hide()
			m.ParentToControl().SetCursor(types.CrSizeAll)
			if m.Component != nil {
				switch m.Component.(type) {
				case lcl.IControl:
					m.Component.(lcl.IControl).SetCursor(m.ParentToControl().Cursor())
				default:
					m.ParentToControl().SetCursor(m.ParentToControl().Cursor())
				}
			}
		}
		m.IsDown = true
	}
}

func (m *IDEComponent) MouseUp(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	m.IsDown = false
	m.IsBorder = false
	if m.IsComponentArea {
		m.Anchor.show()
		m.IsComponentArea = false
		if m.IsDClick {
			if m.DClick != nil {
				m.DClick(button, shift, x, y)
			}
			fmt.Println("双击自定义组件", m.Id, m.Name)
			return
		}
		m.Anchor.refreshAnchorsPoint()
		m.ParentToControl().SetCursor(types.CrDefault)
		if m.Component != nil {
			switch m.Component.(type) {
			case lcl.IControl:
				m.Component.(lcl.IControl).SetCursor(m.ParentToControl().Cursor())
			default:
				m.ParentToControl().SetCursor(m.ParentToControl().Cursor())
			}
		}
	}
	m.Ox, m.Oy, m.Ow, m.Oh = m.ParentToControl().Left(), m.ParentToControl().Top(), m.ParentToControl().Width(), m.ParentToControl().Height()
}
