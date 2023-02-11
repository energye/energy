package ide

import (
	"fmt"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
	"time"
)

func (m *IDEComponent) mouseMove(sender lcl.IObject, shift types.TShiftState, x, y int32) {
	if m.isDown {
		borderRect := m.borderPanel.BoundsRect()
		componentParentRect := m.componentParentPanel.BoundsRect()
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
				m.componentParentPanel.SetWidth(tmpWidth)
			case HTLEFT:
				tmpX := m.componentParentPanel.Left() + (x - m.dx)
				tmpWidth := m.ow + (m.ox - tmpX)
				if tmpWidth <= minW {
					return
				}
				m.borderPanel.SetLeft(tmpX - border/2)
				m.componentParentPanel.SetLeft(tmpX)
				m.borderPanel.SetWidth(tmpWidth + border)
				m.componentParentPanel.SetWidth(tmpWidth)
			case HTTOP:
				tmpY := m.componentParentPanel.Top() + (y - m.dy)
				tmpHeight := m.oh + (m.oy - tmpY)
				if tmpHeight <= minH {
					return
				}
				m.borderPanel.SetTop(tmpY - border/2)
				m.componentParentPanel.SetTop(tmpY)
				m.borderPanel.SetHeight(tmpHeight + border)
				m.componentParentPanel.SetHeight(tmpHeight)
			case HTBOTTOM:
				tmpHeight := m.oh + (y - m.dy)
				if tmpHeight <= minH {
					return
				}
				m.borderPanel.SetHeight(tmpHeight + border)
				m.componentParentPanel.SetHeight(m.oh + (y - m.dy))
			case HTTOPRIGHT:
				tmpY := componentParentRect.Top + (y - m.dy)
				tmpHeight := m.oh + (m.oy - tmpY)
				tmpWidth := m.ow + (x - m.dx)
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				m.borderPanel.SetBounds(borderRect.Left, tmpY-border/2, tmpWidth+border, tmpHeight+border)
				m.componentParentPanel.SetBounds(componentParentRect.Left, tmpY, tmpWidth, tmpHeight)
			case HTBOTTOMRIGHT:
				tmpWidth := m.ow + (x - m.dx)
				tmpHeight := m.oh + (y - m.dy)
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				m.borderPanel.SetBounds(borderRect.Left, borderRect.Top, tmpWidth+border, tmpHeight+border)
				m.componentParentPanel.SetBounds(componentParentRect.Left, componentParentRect.Top, tmpWidth, tmpHeight)
			case HTTOPLEFT:
				tmpX := componentParentRect.Left + (x - m.dx)
				tmpWidth := m.ow + (m.ox - tmpX)
				tmpY := componentParentRect.Top + (y - m.dy)
				tmpHeight := m.oh + (m.oy - tmpY)
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				m.borderPanel.SetBounds(tmpX-border/2, tmpY-border/2, tmpWidth+border, tmpHeight+border)
				m.componentParentPanel.SetBounds(tmpX, tmpY, tmpWidth, tmpHeight)
			case HTBOTTOMLEFT:
				tmpX := componentParentRect.Left + (x - m.dx)
				tmpWidth := m.ow + (m.ox - tmpX)
				tmpHeight := m.oh + (y - m.dy)
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				m.componentParentPanel.SetLeft(tmpX)
				m.componentParentPanel.SetWidth(tmpWidth)
				m.componentParentPanel.SetHeight(tmpHeight)

				m.borderPanel.SetBounds(tmpX-border/2, borderRect.Top, tmpWidth+border, tmpHeight+border)
				m.componentParentPanel.SetBounds(tmpX, componentParentRect.Top, tmpWidth, tmpHeight)
			}
			return
		} else if m.isComponentArea && m.componentType != ctForm { // mouse down component area > move
			m.isDClick = false
			tmpY := componentParentRect.Top + (y - m.dy)
			tmpX := componentParentRect.Left + (x - m.dx)

			m.borderPanel.SetBounds(tmpX-border/2, tmpY-border/2, borderRect.Width(), borderRect.Height())
			m.componentParentPanel.SetBounds(tmpX, tmpY, componentParentRect.Width(), componentParentRect.Height())
			return
		}
	}
	if m.componentType == ctForm {
		if m.isBorder = x <= m.ow && x >= m.ow-borderRange && y <= borderRange; m.isBorder && m.componentType != ctForm { // 右上
			m.componentParentPanel.SetCursor(types.CrSizeSW)
			m.borderHT = HTTOPRIGHT
		} else if m.isBorder = x <= m.ow && x >= m.ow-borderRange && y <= m.oh && y >= m.oh-borderRange; m.isBorder { // 右下
			m.componentParentPanel.SetCursor(types.CrSizeSE)
			m.borderHT = HTBOTTOMRIGHT
		} else if m.isBorder = x <= borderRange && y <= borderRange; m.isBorder && m.componentType != ctForm { //左上
			m.componentParentPanel.SetCursor(types.CrSizeSE)
			m.borderHT = HTTOPLEFT
		} else if m.isBorder = x <= borderRange && y >= m.oh-borderRange; m.isBorder && m.componentType != ctForm { //左下
			m.componentParentPanel.SetCursor(types.CrSizeSW)
			m.borderHT = HTBOTTOMLEFT
		} else if m.isBorder = x <= m.ow && x >= m.ow-borderRange && y > borderRange && y < m.oh-borderRange; m.isBorder { //右
			m.componentParentPanel.SetCursor(types.CrSizeW)
			m.borderHT = HTRIGHT
		} else if m.isBorder = x <= borderRange && y > borderRange && y < m.oh-borderRange; m.isBorder && m.componentType != ctForm { //左
			m.componentParentPanel.SetCursor(types.CrSizeW)
			m.borderHT = HTLEFT
		} else if m.isBorder = x > borderRange && x < m.ow-borderRange && y <= borderRange; m.isBorder && m.componentType != ctForm { //上
			m.componentParentPanel.SetCursor(types.CrSizeN)
			m.borderHT = HTTOP
		} else if m.isBorder = x > borderRange && x < m.ow-borderRange && y >= m.oh-borderRange; m.isBorder { //下
			m.componentParentPanel.SetCursor(types.CrSizeN)
			m.borderHT = HTBOTTOM
		} else {
			m.isBorder = false
			m.componentParentPanel.SetCursor(types.CrDefault)
		}
	}
	if m.component != nil {
		switch m.component.(type) {
		case lcl.IControl:
			m.component.(lcl.IControl).SetCursor(m.componentParentPanel.Cursor())
		default:
			m.componentParentPanel.SetCursor(m.componentParentPanel.Cursor())
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
			m.form.active.clearBorderColor()
			m.form.active = m
			if m.isUseBorder {
				m.form.active.setBorderColor(colors.ClBlack)
			} else {
				m.form.active.clearBorderColor()
			}
		}
		if !m.isBorder && m.componentType != ctForm {
			m.isComponentArea = true
			m.anchor.hide()
			m.componentParentPanel.SetCursor(types.CrSizeAll)
			if m.component != nil {
				switch m.component.(type) {
				case lcl.IControl:
					m.component.(lcl.IControl).SetCursor(m.componentParentPanel.Cursor())
				default:
					m.componentParentPanel.SetCursor(m.componentParentPanel.Cursor())
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
			fmt.Println("双击自定义组件", m.Id, m.name)
			return
		}
		m.refreshAnchorsPoint()
		m.componentParentPanel.SetCursor(types.CrDefault)
		if m.component != nil {
			switch m.component.(type) {
			case lcl.IControl:
				m.component.(lcl.IControl).SetCursor(m.componentParentPanel.Cursor())
			default:
				m.componentParentPanel.SetCursor(m.componentParentPanel.Cursor())
			}
		}
	}
	m.ox, m.oy, m.ow, m.oh = m.componentParentPanel.Left(), m.componentParentPanel.Top(), m.componentParentPanel.Width(), m.componentParentPanel.Height()
}
