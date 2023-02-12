package ide

import (
	"github.com/energye/energy/cef"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
)

type anchor struct {
	component   *IDEComponent
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

func (m *anchor) refreshAnchorsPoint() {
	if m == nil {
		return
	}
	if m.isShow {
		rect := m.component.parentToControl().BoundsRect()
		m.left.SetBounds(rect.Left-pointWC, rect.Top+rect.Height()/2-pointWC, pointW, pointW)
		m.top.SetBounds(rect.Left+rect.Width()/2-pointWC, rect.Top-pointWC, pointW, pointW)
		m.bottom.SetBounds(rect.Left+rect.Width()/2-pointWC, rect.Bottom-pointWC, pointW, pointW)
		m.right.SetBounds(rect.Right-pointWC, rect.Top+rect.Height()/2-pointWC, pointW, pointW)
		m.topLeft.SetBounds(rect.Left-pointWC, rect.Top-pointWC, pointW, pointW)
		m.topRight.SetBounds(rect.Right-pointWC, rect.Top-pointWC, pointW, pointW)
		m.bottomLeft.SetBounds(rect.Left-pointWC, rect.Bottom-pointWC, pointW, pointW)
		m.bottomRight.SetBounds(rect.Right-pointWC, rect.Bottom-pointWC, pointW, pointW)
	}
}

func (m *anchor) newAnchorPoint(owner lcl.IWinControl, ht int32) *lcl.TPanel {
	point := lcl.NewPanel(owner)
	point.SetParent(owner)
	point.SetBevelInner(types.BvSpace)
	point.SetBevelOuter(types.BvNone)
	cef.SetPanelBevelColor(point, colors.ClBlack)
	point.SetColor(colors.ClTeal)
	point.SetOnMouseMove(func(sender lcl.IObject, shift types.TShiftState, x, y int32) {
		m.component.borderHT = ht
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
		if m.component.isDown && m.component.isResize {
			var (
				x, y = x - m.component.anchor.dx, y - m.component.anchor.dy
				rect = m.component.parentToControl().BoundsRect()
			)
			switch ht {
			case HTRIGHT:
				tmpWidth := rect.Width() + x
				if tmpWidth <= minW {
					return
				}
				if m.component.borderPanel != nil {
					m.component.borderPanel.SetWidth(tmpWidth + border)
				}
				m.component.parentToControl().SetWidth(tmpWidth)
			case HTLEFT:
				tmpX := rect.Left + x
				tmpWidth := rect.Width() + (rect.Left - tmpX)
				if tmpWidth <= minW {
					return
				}
				if m.component.borderPanel != nil {
					m.component.borderPanel.SetLeft(tmpX - border/2)
					m.component.borderPanel.SetWidth(tmpWidth + border)
				}
				m.component.parentToControl().SetLeft(tmpX)
				m.component.parentToControl().SetWidth(tmpWidth)
			case HTTOP:
				tmpY := rect.Top + y
				tmpHeight := rect.Height() + (rect.Top - tmpY)
				if tmpHeight <= minH {
					return
				}
				if m.component.borderPanel != nil {
					m.component.borderPanel.SetTop(tmpY - border/2)
					m.component.borderPanel.SetHeight(tmpHeight + border)
				}
				m.component.parentToControl().SetTop(tmpY)
				m.component.parentToControl().SetHeight(tmpHeight)
			case HTBOTTOM:
				tmpHeight := rect.Height() + y
				if tmpHeight <= minH {
					return
				}
				if m.component.borderPanel != nil {
					m.component.borderPanel.SetHeight(tmpHeight + border)
				}
				m.component.parentToControl().SetHeight(tmpHeight)
			case HTTOPRIGHT:
				tmpY := rect.Top + y
				tmpHeight := rect.Height() + (rect.Top - tmpY)
				tmpWidth := rect.Width() + x
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				if m.component.borderPanel != nil {
					m.component.borderPanel.SetTop(tmpY - border/2)
					m.component.borderPanel.SetHeight(tmpHeight + border)
					m.component.borderPanel.SetWidth(tmpWidth + border)
				}
				m.component.parentToControl().SetTop(tmpY)
				m.component.parentToControl().SetHeight(tmpHeight)
				m.component.parentToControl().SetWidth(tmpWidth)
			case HTTOPLEFT:
				tmpX := rect.Left + x
				tmpWidth := rect.Width() + (rect.Left - tmpX)
				tmpY := rect.Top + y
				tmpHeight := rect.Height() + (rect.Top - tmpY)
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				if m.component.borderPanel != nil {
					m.component.borderPanel.SetLeft(tmpX - border/2)
					m.component.borderPanel.SetWidth(tmpWidth + border)
					m.component.borderPanel.SetTop(tmpY - border/2)
					m.component.borderPanel.SetHeight(tmpHeight + border)
				}
				m.component.parentToControl().SetLeft(tmpX)
				m.component.parentToControl().SetWidth(tmpWidth)
				m.component.parentToControl().SetTop(tmpY)
				m.component.parentToControl().SetHeight(tmpHeight)
			case HTBOTTOMRIGHT:
				tmpWidth := rect.Width() + x
				tmpHeight := rect.Height() + y
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				if m.component.borderPanel != nil {
					m.component.borderPanel.SetWidth(tmpWidth + border)
					m.component.borderPanel.SetHeight(tmpHeight + border)
				}
				m.component.parentToControl().SetWidth(tmpWidth)
				m.component.parentToControl().SetHeight(tmpHeight)
			case HTBOTTOMLEFT:
				tmpX := rect.Left + x
				tmpWidth := rect.Width() + (rect.Left - tmpX)
				tmpHeight := rect.Height() + y
				if tmpWidth <= minW || tmpHeight <= minH {
					return
				}
				if m.component.borderPanel != nil {
					m.component.borderPanel.SetLeft(tmpX - border/2)
					m.component.borderPanel.SetWidth(tmpWidth + border)
					m.component.borderPanel.SetHeight(tmpHeight + border)
				}
				m.component.parentToControl().SetLeft(tmpX)
				m.component.parentToControl().SetWidth(tmpWidth)
				m.component.parentToControl().SetHeight(tmpHeight)
			default:
				return
			}
			m.component.anchor.refreshAnchorsPoint()
		}
	})
	point.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		m.component.isDown = true
		m.component.anchor.dx, m.component.anchor.dy = x, y
	})
	point.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		m.component.isDown = false
	})
	return point
}
