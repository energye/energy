package internal

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
	"time"
)

type DClick func(button types.TMouseButton, shift types.TShiftState, x, y int32)

type IDEComponent struct {
	Form                              *IDEForm
	Id                                int
	Name                              string
	Anchor                            *anchor
	IsUseBorder                       bool
	BorderPanel                       *lcl.TPanel
	ComponentParentPanel              lcl.IComponent
	Component                         lcl.IComponent
	ComponentType                     ComponentType
	IsBorder, IsDown, IsComponentArea bool
	IsDClick                          bool
	IsResize                          bool
	ClickTime                         time.Time
	BorderHT                          int32
	Ox, Oy, Ow, Oh                    int32
	Dx, Dy                            int32
	DClick                            DClick
}

func (m *IDEComponent) ParentToPanel() *lcl.TPanel {
	return m.ComponentParentPanel.(*lcl.TPanel)
}

func (m *IDEComponent) ParentToControl() lcl.IControl {
	return m.ComponentParentPanel.(lcl.IControl)
}

func (m *IDEComponent) CreateAnchor() {
	owner := m.ParentToControl().Parent()
	acr := &anchor{Component: m}
	acr.IsShow = true
	acr.Top = acr.newAnchorPoint(owner, HTTOP)
	acr.Bottom = acr.newAnchorPoint(owner, HTBOTTOM)
	acr.Left = acr.newAnchorPoint(owner, HTLEFT)
	acr.Right = acr.newAnchorPoint(owner, HTRIGHT)
	acr.TopLeft = acr.newAnchorPoint(owner, HTTOPLEFT)
	acr.TopRight = acr.newAnchorPoint(owner, HTTOPRIGHT)
	acr.BottomLeft = acr.newAnchorPoint(owner, HTBOTTOMLEFT)
	acr.BottomRight = acr.newAnchorPoint(owner, HTBOTTOMRIGHT)
	m.Anchor = acr
	m.Anchor.refreshAnchorsPoint()
}

func (m *IDEComponent) setBorderColor(color types.TColor) {
	if m == nil || m.BorderPanel == nil {
		return
	}
	m.BorderPanel.SetColor(color)
}

func (m *IDEComponent) clearBorderColor() {
	if m == nil {
		return
	}
	m.Form.Active.Anchor.hide()
	if m.ComponentType != CtForm && m.BorderPanel != nil {
		if m.ComponentType == CtImage {
			m.BorderPanel.SetColor(colors.ClGray)
		} else {
			m.BorderPanel.SetColor(colors.ClSysDefault)
		}
	}
}

func (m *IDEComponent) createAfter() {
	if !m.IsUseBorder {
		m.ComponentParentPanel = m.Component
	}
	m.CreateAnchor()
	switch m.ComponentType {
	case CtEdit:
		m.ParentToControl().SetName(m.Name)
	default:
		m.Component.SetName(m.Name)
		m.ParentToControl().SetCaption(m.Name)
	}
	pm := lcl.NewPopupMenu(m.Component)
	item := lcl.NewMenuItem(m.Component)
	item.SetCaption("删除")
	item.SetOnClick(func(lcl.IObject) {
		m.Form.RemoveComponent(m.Id)
	})
	pm.Items().Add(item)
	switch m.Component.(type) {
	case lcl.IControl:
		m.Component.(lcl.IControl).SetPopupMenu(pm)
		m.Component.(lcl.IControl).SetHint(m.Name)
		m.Component.(lcl.IControl).SetShowHint(true)
	default:
		m.ParentToControl().SetPopupMenu(pm)
		m.ParentToControl().SetHint(m.Name)
		m.ParentToControl().SetShowHint(true)
	}
	m.switchActive(m)
}

func (m *IDEComponent) switchActive(active *IDEComponent) {
	if m.Form.Active != nil {
		m.Form.Active.clearBorderColor()
	}
	m.Form.Active = m
	if m.IsUseBorder {
		m.Form.Active.setBorderColor(colors.ClBlack)
		m.Form.Active.Anchor.show()
	} else {
		m.Form.Active.clearBorderColor()
	}
}
