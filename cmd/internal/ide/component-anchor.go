package ide

import (
	"github.com/energye/energy/v2/pkgs/ext"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
)

type anchor struct {
	Component   *IDEComponent
	Top         *lcl.TPanel
	Bottom      *lcl.TPanel
	Left        *lcl.TPanel
	Right       *lcl.TPanel
	TopLeft     *lcl.TPanel
	TopRight    *lcl.TPanel
	BottomLeft  *lcl.TPanel
	BottomRight *lcl.TPanel
	IsShow      bool
	Dx, Dy      int32
}

func (m *anchor) hide() {
	if m == nil || !m.IsShow {
		return
	}
	m.Top.Hide()
	m.Bottom.Hide()
	m.Left.Hide()
	m.Right.Hide()
	m.TopLeft.Hide()
	m.TopRight.Hide()
	m.BottomLeft.Hide()
	m.BottomRight.Hide()
	m.IsShow = false
}

func (m *anchor) show() {
	if m == nil || m.IsShow {
		return
	}
	m.Top.Show()
	m.Bottom.Show()
	m.Left.Show()
	m.Right.Show()
	m.TopLeft.Show()
	m.TopRight.Show()
	m.BottomLeft.Show()
	m.BottomRight.Show()
	m.IsShow = true
}

func (m *anchor) remove() {
	if m == nil {
		return
	}
	m.Top.Free()
	m.Bottom.Free()
	m.Left.Free()
	m.Right.Free()
	m.TopLeft.Free()
	m.TopRight.Free()
	m.BottomLeft.Free()
	m.BottomRight.Free()
}

func (m *anchor) refreshAnchorsPoint() {
	if m == nil {
		return
	}
	if m.IsShow {
		rect := m.Component.ParentToControl().BoundsRect()
		m.Left.SetBounds(rect.Left-PointWC, rect.Top+rect.Height()/2-PointWC, PointW, PointW)
		m.Top.SetBounds(rect.Left+rect.Width()/2-PointWC, rect.Top-PointWC, PointW, PointW)
		m.Bottom.SetBounds(rect.Left+rect.Width()/2-PointWC, rect.Bottom-PointWC, PointW, PointW)
		m.Right.SetBounds(rect.Right-PointWC, rect.Top+rect.Height()/2-PointWC, PointW, PointW)
		m.TopLeft.SetBounds(rect.Left-PointWC, rect.Top-PointWC, PointW, PointW)
		m.TopRight.SetBounds(rect.Right-PointWC, rect.Top-PointWC, PointW, PointW)
		m.BottomLeft.SetBounds(rect.Left-PointWC, rect.Bottom-PointWC, PointW, PointW)
		m.BottomRight.SetBounds(rect.Right-PointWC, rect.Bottom-PointWC, PointW, PointW)
	}
}

func (m *anchor) newAnchorPoint(owner lcl.IWinControl, ht int32) *lcl.TPanel {
	point := lcl.NewPanel(owner)
	point.SetParent(owner)
	point.SetBevelInner(types.BvSpace)
	point.SetBevelOuter(types.BvNone)
	ext.SetPanelBevelColor(point, colors.ClBlack)
	point.SetColor(colors.ClTeal)
	point.SetOnMouseMove(func(sender lcl.IObject, shift types.TShiftState, x, y int32) {
		m.Component.BorderHT = ht
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
		//m.MouseMove(sender, shift, x, y)
		if m.Component.IsDown && m.Component.IsResize {
			var (
				x, y = x - m.Component.Anchor.Dx, y - m.Component.Anchor.Dy
				rect = m.Component.ParentToControl().BoundsRect()
			)
			switch ht {
			case HTRIGHT:
				tmpWidth := rect.Width() + x
				if tmpWidth <= MinW {
					return
				}
				if m.Component.BorderPanel != nil {
					m.Component.BorderPanel.SetWidth(tmpWidth + Border)
				}
				m.Component.ParentToControl().SetWidth(tmpWidth)
			case HTLEFT:
				tmpX := rect.Left + x
				tmpWidth := rect.Width() + (rect.Left - tmpX)
				if tmpWidth <= MinW {
					return
				}
				if m.Component.BorderPanel != nil {
					m.Component.BorderPanel.SetLeft(tmpX - Border/2)
					m.Component.BorderPanel.SetWidth(tmpWidth + Border)
				}
				m.Component.ParentToControl().SetLeft(tmpX)
				m.Component.ParentToControl().SetWidth(tmpWidth)
			case HTTOP:
				tmpY := rect.Top + y
				tmpHeight := rect.Height() + (rect.Top - tmpY)
				if tmpHeight <= MinH {
					return
				}
				if m.Component.BorderPanel != nil {
					m.Component.BorderPanel.SetTop(tmpY - Border/2)
					m.Component.BorderPanel.SetHeight(tmpHeight + Border)
				}
				m.Component.ParentToControl().SetTop(tmpY)
				m.Component.ParentToControl().SetHeight(tmpHeight)
			case HTBOTTOM:
				tmpHeight := rect.Height() + y
				if tmpHeight <= MinH {
					return
				}
				if m.Component.BorderPanel != nil {
					m.Component.BorderPanel.SetHeight(tmpHeight + Border)
				}
				m.Component.ParentToControl().SetHeight(tmpHeight)
			case HTTOPRIGHT:
				tmpY := rect.Top + y
				tmpHeight := rect.Height() + (rect.Top - tmpY)
				tmpWidth := rect.Width() + x
				if tmpWidth <= MinW || tmpHeight <= MinH {
					return
				}
				if m.Component.BorderPanel != nil {
					m.Component.BorderPanel.SetTop(tmpY - Border/2)
					m.Component.BorderPanel.SetHeight(tmpHeight + Border)
					m.Component.BorderPanel.SetWidth(tmpWidth + Border)
				}
				m.Component.ParentToControl().SetTop(tmpY)
				m.Component.ParentToControl().SetHeight(tmpHeight)
				m.Component.ParentToControl().SetWidth(tmpWidth)
			case HTTOPLEFT:
				tmpX := rect.Left + x
				tmpWidth := rect.Width() + (rect.Left - tmpX)
				tmpY := rect.Top + y
				tmpHeight := rect.Height() + (rect.Top - tmpY)
				if tmpWidth <= MinW || tmpHeight <= MinH {
					return
				}
				if m.Component.BorderPanel != nil {
					m.Component.BorderPanel.SetLeft(tmpX - Border/2)
					m.Component.BorderPanel.SetWidth(tmpWidth + Border)
					m.Component.BorderPanel.SetTop(tmpY - Border/2)
					m.Component.BorderPanel.SetHeight(tmpHeight + Border)
				}
				m.Component.ParentToControl().SetLeft(tmpX)
				m.Component.ParentToControl().SetWidth(tmpWidth)
				m.Component.ParentToControl().SetTop(tmpY)
				m.Component.ParentToControl().SetHeight(tmpHeight)
			case HTBOTTOMRIGHT:
				tmpWidth := rect.Width() + x
				tmpHeight := rect.Height() + y
				if tmpWidth <= MinW || tmpHeight <= MinH {
					return
				}
				if m.Component.BorderPanel != nil {
					m.Component.BorderPanel.SetWidth(tmpWidth + Border)
					m.Component.BorderPanel.SetHeight(tmpHeight + Border)
				}
				m.Component.ParentToControl().SetWidth(tmpWidth)
				m.Component.ParentToControl().SetHeight(tmpHeight)
			case HTBOTTOMLEFT:
				tmpX := rect.Left + x
				tmpWidth := rect.Width() + (rect.Left - tmpX)
				tmpHeight := rect.Height() + y
				if tmpWidth <= MinW || tmpHeight <= MinH {
					return
				}
				if m.Component.BorderPanel != nil {
					m.Component.BorderPanel.SetLeft(tmpX - Border/2)
					m.Component.BorderPanel.SetWidth(tmpWidth + Border)
					m.Component.BorderPanel.SetHeight(tmpHeight + Border)
				}
				m.Component.ParentToControl().SetLeft(tmpX)
				m.Component.ParentToControl().SetWidth(tmpWidth)
				m.Component.ParentToControl().SetHeight(tmpHeight)
			default:
				return
			}
			m.Component.Anchor.refreshAnchorsPoint()
		}
	})
	point.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		m.Component.IsDown = true
		m.Component.Anchor.Dx, m.Component.Anchor.Dy = x, y
	})
	point.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		m.Component.IsDown = false
	})
	return point
}
