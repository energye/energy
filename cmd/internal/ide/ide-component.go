package ide

import (
	"github.com/energye/energy/cef"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
	"time"
)

type dClick func(button types.TMouseButton, shift types.TShiftState, x, y int32)
type IDEComponent struct {
	form                              *IDEForm
	Id                                int
	name                              string
	anchor                            *anchor
	isUseBorder                       bool
	borderPanel                       *lcl.TPanel
	componentParentPanel              lcl.IComponent
	component                         lcl.IComponent
	componentType                     componentType
	isBorder, isDown, isComponentArea bool
	isDClick                          bool
	isResize                          bool
	clickTime                         time.Time
	borderHT                          int32
	ox, oy, ow, oh                    int32
	dx, dy                            int32
	dClick                            dClick
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
	isShow      bool
	dx, dy      int32
}

func (m *anchor) hide() {
	if m == nil || !m.isShow {
		return
	}
	m.top.Hide()
	m.bottom.Hide()
	m.left.Hide()
	m.right.Hide()
	m.topLeft.Hide()
	m.topRight.Hide()
	m.bottomLeft.Hide()
	m.bottomRight.Hide()
	m.isShow = false
}

func (m *anchor) show() {
	if m == nil || m.isShow {
		return
	}
	m.top.Show()
	m.bottom.Show()
	m.left.Show()
	m.right.Show()
	m.topLeft.Show()
	m.topRight.Show()
	m.bottomLeft.Show()
	m.bottomRight.Show()
	m.isShow = true
}

func (m *anchor) remove() {
	if m == nil {
		return
	}
	m.top.Free()
	m.bottom.Free()
	m.left.Free()
	m.right.Free()
	m.topLeft.Free()
	m.topRight.Free()
	m.bottomLeft.Free()
	m.bottomRight.Free()
}

func (m *IDEComponent) parentToPanel() *lcl.TPanel {
	return m.componentParentPanel.(*lcl.TPanel)
}

func (m *IDEComponent) parentToControl() lcl.IControl {
	return m.componentParentPanel.(lcl.IControl)
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
		//m.mouseMove(sender, shift, x, y)
		if m.isDown && m.isResize {
			var (
				x, y = x - m.anchor.dx, y - m.anchor.dy
				rect = m.parentToControl().BoundsRect()
			)
			switch ht {
			case HTRIGHT:
				tmpWidth := rect.Width() + x
				if tmpWidth <= minW {
					return
				}
				if m.borderPanel != nil {
					m.borderPanel.SetWidth(tmpWidth + border)
				}
				m.parentToControl().SetWidth(tmpWidth)
			case HTLEFT:
				tmpX := rect.Left + x
				tmpWidth := rect.Width() + (rect.Left - tmpX)
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
				tmpY := rect.Top + y
				tmpHeight := rect.Height() + (rect.Top - tmpY)
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
				tmpHeight := rect.Height() + y
				if tmpHeight <= minH {
					return
				}
				if m.borderPanel != nil {
					m.borderPanel.SetHeight(tmpHeight + border)
				}
				m.parentToControl().SetHeight(tmpHeight)
			case HTTOPRIGHT:
				tmpY := rect.Top + y
				tmpHeight := rect.Height() + (rect.Top - tmpY)
				tmpWidth := rect.Width() + x
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				if m.borderPanel != nil {
					m.borderPanel.SetTop(tmpY - border/2)
					m.borderPanel.SetHeight(tmpHeight + border)
					m.borderPanel.SetWidth(tmpWidth + border)
				}
				m.parentToControl().SetTop(tmpY)
				m.parentToControl().SetHeight(tmpHeight)
				m.parentToControl().SetWidth(tmpWidth)
			case HTTOPLEFT:
				tmpX := rect.Left + x
				tmpWidth := rect.Width() + (rect.Left - tmpX)
				tmpY := rect.Top + y
				tmpHeight := rect.Height() + (rect.Top - tmpY)
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				if m.borderPanel != nil {
					m.borderPanel.SetLeft(tmpX - border/2)
					m.borderPanel.SetWidth(tmpWidth + border)
					m.borderPanel.SetTop(tmpY - border/2)
					m.borderPanel.SetHeight(tmpHeight + border)
				}
				m.parentToControl().SetLeft(tmpX)
				m.parentToControl().SetWidth(tmpWidth)
				m.parentToControl().SetTop(tmpY)
				m.parentToControl().SetHeight(tmpHeight)
			case HTBOTTOMRIGHT:
				tmpWidth := rect.Width() + x
				tmpHeight := rect.Height() + y
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				if m.borderPanel != nil {
					m.borderPanel.SetWidth(tmpWidth + border)
					m.borderPanel.SetHeight(tmpHeight + border)
				}
				m.parentToControl().SetWidth(tmpWidth)
				m.parentToControl().SetHeight(tmpHeight)
			case HTBOTTOMLEFT:
				tmpX := rect.Left + x
				tmpWidth := rect.Width() + (rect.Left - tmpX)
				tmpHeight := rect.Height() + y
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				if m.borderPanel != nil {
					m.borderPanel.SetLeft(tmpX - border/2)
					m.borderPanel.SetWidth(tmpWidth + border)
					m.borderPanel.SetHeight(tmpHeight + border)
				}
				m.parentToControl().SetLeft(tmpX)
				m.parentToControl().SetWidth(tmpWidth)
				m.parentToControl().SetHeight(tmpHeight)
			default:
				return
			}
			m.refreshAnchorsPoint()
		}
	})
	point.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		m.isDown = true
		m.anchor.dx, m.anchor.dy = x, y
	})
	point.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		m.isDown = false
	})
	return point
}

func (m *IDEComponent) createAnchor() {
	owner := m.parentToControl().Parent()
	acr := &anchor{}
	acr.isShow = true
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
	if m.anchor.isShow {
		rect := m.parentToControl().BoundsRect()
		m.anchor.left.SetBounds(rect.Left-pointWC, rect.Top+rect.Height()/2-pointWC, pointW, pointW)
		m.anchor.top.SetBounds(rect.Left+rect.Width()/2-pointWC, rect.Top-pointWC, pointW, pointW)
		m.anchor.bottom.SetBounds(rect.Left+rect.Width()/2-pointWC, rect.Bottom-pointWC, pointW, pointW)
		m.anchor.right.SetBounds(rect.Right-pointWC, rect.Top+rect.Height()/2-pointWC, pointW, pointW)
		m.anchor.topLeft.SetBounds(rect.Left-pointWC, rect.Top-pointWC, pointW, pointW)
		m.anchor.topRight.SetBounds(rect.Right-pointWC, rect.Top-pointWC, pointW, pointW)
		m.anchor.bottomLeft.SetBounds(rect.Left-pointWC, rect.Bottom-pointWC, pointW, pointW)
		m.anchor.bottomRight.SetBounds(rect.Right-pointWC, rect.Bottom-pointWC, pointW, pointW)
	}
}

func (m *IDEComponent) setBorderColor(color types.TColor) {
	if m == nil || m.borderPanel == nil {
		return
	}
	m.borderPanel.SetColor(color)
}

func (m *IDEComponent) clearBorderColor() {
	if m == nil {
		return
	}
	m.form.active.anchor.hide()
	if m.componentType != ctForm && m.borderPanel != nil {
		if m.componentType == ctImage {
			m.borderPanel.SetColor(colors.ClGray)
		} else {
			m.borderPanel.SetColor(colors.ClSysDefault)
		}
	}
}

func (m *IDEComponent) createAfter() {
	if !m.isUseBorder {
		m.componentParentPanel = m.component
	}
	m.createAnchor()
	switch m.componentType {
	case ctEdit:
		m.parentToControl().SetName(m.name)
	default:
		m.component.SetName(m.name)
		m.parentToControl().SetCaption(m.name)
	}
	pm := lcl.NewPopupMenu(m.component)
	item := lcl.NewMenuItem(m.component)
	item.SetCaption("删除")
	item.SetOnClick(func(lcl.IObject) {
		m.form.RemoveComponent(m.Id)
	})
	pm.Items().Add(item)
	switch m.component.(type) {
	case lcl.IControl:
		m.component.(lcl.IControl).SetPopupMenu(pm)
		m.component.(lcl.IControl).SetHint(m.name)
		m.component.(lcl.IControl).SetShowHint(true)
	default:
		m.parentToControl().SetPopupMenu(pm)
		m.parentToControl().SetHint(m.name)
		m.parentToControl().SetShowHint(true)
	}
	m.switchActive(m)
}

func (m *IDEComponent) switchActive(active *IDEComponent) {
	if m.form.active != nil {
		m.form.active.clearBorderColor()
	}
	m.form.active = m
	if m.isUseBorder {
		m.form.active.setBorderColor(colors.ClBlack)
		m.form.active.anchor.show()
	} else {
		m.form.active.clearBorderColor()
	}
}
