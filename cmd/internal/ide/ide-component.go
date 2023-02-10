package ide

import (
	"github.com/energye/energy/cef"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
)

type IDEComponent struct {
	form                              *IDEForm
	Id                                int
	name                              string
	anchor                            *anchor
	borderPanel                       *lcl.TPanel
	isUseBorder                       bool
	componentParentPanel              *lcl.TPanel
	componentControl                  lcl.IControl
	componentType                     componentType
	isBorder, isDown, isComponentArea bool
	borderHT                          int32
	ox, oy, ow, oh                    int32
	dx, dy                            int32
}

type anchor struct {
	top         *lcl.TPanel
	bottom      *lcl.TPanel
	left        *lcl.TPanel
	right       *lcl.TPanel
	topLeft     *lcl.TPanel
	topRight    *lcl.TPanel
	bottomLeft  *lcl.TPanel
	bottomRight *lcl.TPanel
}

func (m *IDEComponent) newAnchorPoint(owner lcl.IWinControl, ht int32) *lcl.TPanel {
	point := lcl.NewPanel(owner)
	point.SetParent(owner)
	point.SetBevelInner(types.BvSpace)
	point.SetBevelOuter(types.BvNone)
	cef.SetPanelBevelColor(point, colors.ClBlack)
	point.SetColor(colors.ClTeal)
	point.SetOnMouseMove(func(sender lcl.IObject, shift types.TShiftState, x, y int32) {
		m.borderHT = ht
		switch ht {
		case HTTOP, HTBOTTOM:
			point.SetCursor(types.CrSizeN)
		case HTLEFT, HTRIGHT:
			point.SetCursor(types.CrSizeW)
		case HTTOPRIGHT, HTBOTTOMLEFT:
			point.SetCursor(types.CrSizeSW)
		case HTTOPLEFT, HTBOTTOMRIGHT:
			point.SetCursor(types.CrSizeSE)
		default:
			point.SetCursor(types.CrDefault)
		}
		m.mouseMove(sender, shift, x, y)
	})
	point.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		m.mouseDown(sender, button, shift, x, y)
	})
	point.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		m.mouseUp(sender, button, shift, x, y)
	})
	return point
}

func (m *IDEComponent) createAnchor(owner lcl.IWinControl) {
	acr := &anchor{}
	acr.top = m.newAnchorPoint(owner, HTTOP)
	acr.bottom = m.newAnchorPoint(owner, HTBOTTOM)
	acr.left = m.newAnchorPoint(owner, HTLEFT)
	acr.right = m.newAnchorPoint(owner, HTRIGHT)
	acr.topLeft = m.newAnchorPoint(owner, HTTOPLEFT)
	acr.topRight = m.newAnchorPoint(owner, HTTOPRIGHT)
	acr.bottomLeft = m.newAnchorPoint(owner, HTBOTTOMLEFT)
	acr.bottomRight = m.newAnchorPoint(owner, HTBOTTOMRIGHT)
	m.anchor = acr
	m.refreshAnchorsPoint()
}

func (m *IDEComponent) refreshAnchorsPoint() {
	if m.anchor == nil {
		return
	}
	rect := m.componentParentPanel.BoundsRect()
	m.anchor.top.SetBounds(rect.Left+rect.Width()/2-pointWC, rect.Top-pointWC, pointW, pointW)
	m.anchor.bottom.SetBounds(rect.Left+rect.Width()/2-pointWC, rect.Bottom-pointWC, pointW, pointW)
	m.anchor.left.SetBounds(rect.Left-pointWC, rect.Top+rect.Height()/2-pointWC, pointW, pointW)
	m.anchor.right.SetBounds(rect.Right-pointWC, rect.Top+rect.Height()/2-pointWC, pointW, pointW)
}

func (m *IDEComponent) mouseMove(sender lcl.IObject, shift types.TShiftState, x, y int32) {
	if m.isDown {
		if m.isBorder { //mouse down borderMargin > resize
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
				tmpY := m.componentParentPanel.Top() + (y - m.dy)
				tmpHeight := m.oh + (m.oy - tmpY)
				tmpWidth := m.ow + (x - m.dx)
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				m.borderPanel.SetTop(tmpY - border/2)
				m.componentParentPanel.SetTop(tmpY)
				m.borderPanel.SetHeight(tmpHeight + border)
				m.componentParentPanel.SetHeight(tmpHeight)
				m.borderPanel.SetWidth(tmpWidth + border)
				m.componentParentPanel.SetWidth(tmpWidth)
			case HTBOTTOMRIGHT:
				tmpWidth := m.ow + (x - m.dx)
				tmpHeight := m.oh + (y - m.dy)
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				m.borderPanel.SetWidth(tmpWidth + border)
				m.componentParentPanel.SetWidth(tmpWidth)
				m.borderPanel.SetHeight(tmpHeight + border)
				m.componentParentPanel.SetHeight(tmpHeight)
			case HTTOPLEFT:
				tmpX := m.componentParentPanel.Left() + (x - m.dx)
				tmpWidth := m.ow + (m.ox - tmpX)
				tmpY := m.componentParentPanel.Top() + (y - m.dy)
				tmpHeight := m.oh + (m.oy - tmpY)
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				m.borderPanel.SetLeft(tmpX - border/2)
				m.borderPanel.SetWidth(tmpWidth + border)
				m.componentParentPanel.SetLeft(tmpX)
				m.componentParentPanel.SetWidth(tmpWidth)
				m.borderPanel.SetTop(tmpY - border/2)
				m.borderPanel.SetHeight(tmpHeight + border)
				m.componentParentPanel.SetTop(tmpY)
				m.componentParentPanel.SetHeight(tmpHeight)
			case HTBOTTOMLEFT:
				tmpX := m.componentParentPanel.Left() + (x - m.dx)
				tmpWidth := m.ow + (m.ox - tmpX)
				tmpHeight := m.oh + (y - m.dy)
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				m.borderPanel.SetLeft(tmpX - border/2)
				m.borderPanel.SetWidth(tmpWidth + border)
				m.borderPanel.SetHeight(tmpHeight + border)
				m.componentParentPanel.SetLeft(tmpX)
				m.componentParentPanel.SetWidth(tmpWidth)
				m.componentParentPanel.SetHeight(tmpHeight)
			}
			//rect := m.componentParentPanel.BoundsRect()
			//fx, fy, fw, fh = rect.Left, rect.Top, rect.Width(), rect.Height()
			//Ide.formsSyncSize(m.Id)
			//m.refreshAnchorsPoint()
			return
		} else if m.isComponentArea && m.componentType != ctForm { // mouse down componentControl area > move
			tmpY := m.componentParentPanel.Top() + (y - m.dy)
			tmpX := m.componentParentPanel.Left() + (x - m.dx)
			m.borderPanel.SetLeft(tmpX - border/2)
			m.borderPanel.SetTop(tmpY - border/2)
			m.componentParentPanel.SetTop(tmpY)
			m.componentParentPanel.SetLeft(tmpX)
			//rect := m.componentParentPanel.BoundsRect()
			//fx, fy, fw, fh = rect.Left, rect.Top, rect.Width(), rect.Height()
			//Ide.formsSyncSize(m.Id)
			//m.refreshAnchorsPoint()
			return
		}
	}
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
	if m.componentControl != nil {
		m.componentControl.SetCursor(m.componentParentPanel.Cursor())
	}
}

func (m *IDEComponent) mouseDown(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if button == types.MbLeft {
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
			m.componentParentPanel.SetCursor(types.CrSizeAll)
			if m.componentControl != nil {
				m.componentControl.SetCursor(m.componentParentPanel.Cursor())
			}
		}
		m.isDown = true
	}
}

func (m *IDEComponent) mouseUp(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	m.isDown = false
	if m.isComponentArea {
		m.isComponentArea = false
		m.componentParentPanel.SetCursor(types.CrDefault)
		if m.componentControl != nil {
			m.componentControl.SetCursor(m.componentParentPanel.Cursor())
		}
	}
	m.ox, m.oy, m.ow, m.oh = m.componentParentPanel.Left(), m.componentParentPanel.Top(), m.componentParentPanel.Width(), m.componentParentPanel.Height()
}

func (m *IDEComponent) setBorderColor(color types.TColor) {
	if m == nil {
		return
	}
	m.borderPanel.SetColor(color)
}

func (m *IDEComponent) clearBorderColor() {
	if m == nil {
		return
	}
	if m.componentType != ctForm {
		if m.componentType == ctImage {
			m.borderPanel.SetColor(colors.ClGray)
		} else {
			m.borderPanel.SetColor(colors.ClSysDefault)
		}
	}
}

func (m *IDEComponent) createAfter() {
	m.componentParentPanel.SetCaption(m.name)
	pm := lcl.NewPopupMenu(m.componentControl)
	item := lcl.NewMenuItem(m.componentControl)
	item.SetCaption("删除")
	item.SetOnClick(func(lcl.IObject) {
		m.form.RemoveComponent(m.Id)
	})
	pm.Items().Add(item)
	m.componentControl.SetPopupMenu(pm)
}
