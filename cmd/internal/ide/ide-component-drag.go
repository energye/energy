package ide

import (
	"fmt"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"time"
)

func (m *IDEComponent) mouseMove(sender lcl.IObject, shift types.TShiftState, x, y int32) {
	if m.isDown {
		var borderRect types.TRect
		if m.borderPanel != nil {
			borderRect = m.borderPanel.BoundsRect()
		}
		componentParentRect := m.parentToControl().BoundsRect()
		if m.isBorder && m.componentType == ctForm { //mouse down borderMargin > resize
			switch m.borderHT {
			case HTRIGHT:
				tmpWidth := m.ow + (x - m.dx)
				if tmpWidth <= minW {
					return
				}
				if m.borderPanel != nil {
					m.borderPanel.SetWidth(tmpWidth + border)
				}
				m.parentToControl().SetWidth(tmpWidth)
			case HTLEFT:
				tmpX := m.parentToControl().Left() + (x - m.dx)
				tmpWidth := m.ow + (m.ox - tmpX)
				if tmpWidth <= minW {
					return
				}
				if m.borderPanel != nil {
					m.borderPanel.SetLeft(tmpX - border/2)
					m.borderPanel.SetWidth(tmpWidth + border)
				}
				m.parentToControl().SetLeft(tmpX)
				m.parentToControl().SetWidth(tmpWidth)
			case HTTOP:
				tmpY := m.parentToControl().Top() + (y - m.dy)
				tmpHeight := m.oh + (m.oy - tmpY)
				if tmpHeight <= minH {
					return
				}
				if m.borderPanel != nil {
					m.borderPanel.SetTop(tmpY - border/2)
					m.borderPanel.SetHeight(tmpHeight + border)
				}
				m.parentToControl().SetTop(tmpY)
				m.parentToControl().SetHeight(tmpHeight)
			case HTBOTTOM:
				tmpHeight := m.oh + (y - m.dy)
				if tmpHeight <= minH {
					return
				}
				if m.borderPanel != nil {
					m.borderPanel.SetHeight(tmpHeight + border)
				}
				m.parentToControl().SetHeight(m.oh + (y - m.dy))
			case HTTOPRIGHT:
				tmpY := componentParentRect.Top + (y - m.dy)
				tmpHeight := m.oh + (m.oy - tmpY)
				tmpWidth := m.ow + (x - m.dx)
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				if m.borderPanel != nil {
					m.borderPanel.SetBounds(borderRect.Left, tmpY-border/2, tmpWidth+border, tmpHeight+border)
				}
				m.parentToControl().SetBounds(componentParentRect.Left, tmpY, tmpWidth, tmpHeight)
			case HTBOTTOMRIGHT:
				tmpWidth := m.ow + (x - m.dx)
				tmpHeight := m.oh + (y - m.dy)
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				if m.borderPanel != nil {
					m.borderPanel.SetBounds(borderRect.Left, borderRect.Top, tmpWidth+border, tmpHeight+border)
				}
				m.parentToControl().SetBounds(componentParentRect.Left, componentParentRect.Top, tmpWidth, tmpHeight)
			case HTTOPLEFT:
				tmpX := componentParentRect.Left + (x - m.dx)
				tmpWidth := m.ow + (m.ox - tmpX)
				tmpY := componentParentRect.Top + (y - m.dy)
				tmpHeight := m.oh + (m.oy - tmpY)
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				if m.borderPanel != nil {
					m.borderPanel.SetBounds(tmpX-border/2, tmpY-border/2, tmpWidth+border, tmpHeight+border)
				}
				m.parentToControl().SetBounds(tmpX, tmpY, tmpWidth, tmpHeight)
			case HTBOTTOMLEFT:
				tmpX := componentParentRect.Left + (x - m.dx)
				tmpWidth := m.ow + (m.ox - tmpX)
				tmpHeight := m.oh + (y - m.dy)
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				if m.borderPanel != nil {
					m.borderPanel.SetBounds(tmpX-border/2, borderRect.Top, tmpWidth+border, tmpHeight+border)
				}
				m.parentToControl().SetBounds(tmpX, componentParentRect.Top, tmpWidth, tmpHeight)
			}
			return
		} else if m.isComponentArea && m.componentType != ctForm { // mouse down component area > move
			m.isDClick = false
			tmpY := componentParentRect.Top + (y - m.dy)
			tmpX := componentParentRect.Left + (x - m.dx)
			if m.borderPanel != nil {
				m.borderPanel.SetBounds(tmpX-border/2, tmpY-border/2, borderRect.Width(), borderRect.Height())
			}
			m.parentToControl().SetBounds(tmpX, tmpY, componentParentRect.Width(), componentParentRect.Height())
			return
		}
	}
	if m.componentType == ctForm {
		if m.isBorder = x <= m.ow && x >= m.ow-borderRange && y <= borderRange; m.isBorder && m.componentType != ctForm { // 右上
			m.parentToControl().SetCursor(types.CrSizeSW)
			m.borderHT = HTTOPRIGHT
		} else if m.isBorder = x <= m.ow && x >= m.ow-borderRange && y <= m.oh && y >= m.oh-borderRange; m.isBorder { // 右下
			m.parentToControl().SetCursor(types.CrSizeSE)
			m.borderHT = HTBOTTOMRIGHT
		} else if m.isBorder = x <= borderRange && y <= borderRange; m.isBorder && m.componentType != ctForm { //左上
			m.parentToControl().SetCursor(types.CrSizeSE)
			m.borderHT = HTTOPLEFT
		} else if m.isBorder = x <= borderRange && y >= m.oh-borderRange; m.isBorder && m.componentType != ctForm { //左下
			m.parentToControl().SetCursor(types.CrSizeSW)
			m.borderHT = HTBOTTOMLEFT
		} else if m.isBorder = x <= m.ow && x >= m.ow-borderRange && y > borderRange && y < m.oh-borderRange; m.isBorder { //右
			m.parentToControl().SetCursor(types.CrSizeW)
			m.borderHT = HTRIGHT
		} else if m.isBorder = x <= borderRange && y > borderRange && y < m.oh-borderRange; m.isBorder && m.componentType != ctForm { //左
			m.parentToControl().SetCursor(types.CrSizeW)
			m.borderHT = HTLEFT
		} else if m.isBorder = x > borderRange && x < m.ow-borderRange && y <= borderRange; m.isBorder && m.componentType != ctForm { //上
			m.parentToControl().SetCursor(types.CrSizeN)
			m.borderHT = HTTOP
		} else if m.isBorder = x > borderRange && x < m.ow-borderRange && y >= m.oh-borderRange; m.isBorder { //下
			m.parentToControl().SetCursor(types.CrSizeN)
			m.borderHT = HTBOTTOM
		} else {
			m.isBorder = false
			m.parentToControl().SetCursor(types.CrDefault)
		}
	}
	if m.component != nil {
		switch m.component.(type) {
		case lcl.IControl:
			m.component.(lcl.IControl).SetCursor(m.parentToControl().Cursor())
		default:
			m.parentToControl().SetCursor(m.parentToControl().Cursor())
		}
	}
}

func (m *IDEComponent) mouseDown(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if button == types.MbLeft {
		if time.Now().UnixMilli()-m.clickTime.UnixMilli() < 500 {
			m.isDClick = true
		} else {
			m.isDClick = false
		}
		m.clickTime = time.Now()
		m.dx = x
		m.dy = y
		if m.componentType == ctForm {
			for _, form := range Ide.forms {
				if form.active != nil {
					form.active.clearBorderColor()
				}
			}
		} else {
			m.switchActive(m)
		}
		if !m.isBorder && m.componentType != ctForm {
			m.isComponentArea = true
			m.anchor.hide()
			m.parentToControl().SetCursor(types.CrSizeAll)
			if m.component != nil {
				switch m.component.(type) {
				case lcl.IControl:
					m.component.(lcl.IControl).SetCursor(m.parentToControl().Cursor())
				default:
					m.parentToControl().SetCursor(m.parentToControl().Cursor())
				}
			}
		}
		m.isDown = true
	}
}

func (m *IDEComponent) mouseUp(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	m.isDown = false
	m.isBorder = false
	if m.isComponentArea {
		m.anchor.show()
		m.isComponentArea = false
		if m.isDClick {
			if m.dClick != nil {
				m.dClick(button, shift, x, y)
			}
			fmt.Println("双击自定义组件", m.Id, m.name)
			return
		}
		m.refreshAnchorsPoint()
		m.parentToControl().SetCursor(types.CrDefault)
		if m.component != nil {
			switch m.component.(type) {
			case lcl.IControl:
				m.component.(lcl.IControl).SetCursor(m.parentToControl().Cursor())
			default:
				m.parentToControl().SetCursor(m.parentToControl().Cursor())
			}
		}
	}
	m.ox, m.oy, m.ow, m.oh = m.parentToControl().Left(), m.parentToControl().Top(), m.parentToControl().Width(), m.parentToControl().Height()
}
