package ide

import (
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

func (m *IDEComponent) parentToPanel() *lcl.TPanel {
	return m.componentParentPanel.(*lcl.TPanel)
}

func (m *IDEComponent) parentToControl() lcl.IControl {
	return m.componentParentPanel.(lcl.IControl)
}

func (m *IDEComponent) createAnchor() {
	owner := m.parentToControl().Parent()
	acr := &anchor{component: m}
	acr.isShow = true
	acr.top = acr.newAnchorPoint(owner, HTTOP)
	acr.bottom = acr.newAnchorPoint(owner, HTBOTTOM)
	acr.left = acr.newAnchorPoint(owner, HTLEFT)
	acr.right = acr.newAnchorPoint(owner, HTRIGHT)
	acr.topLeft = acr.newAnchorPoint(owner, HTTOPLEFT)
	acr.topRight = acr.newAnchorPoint(owner, HTTOPRIGHT)
	acr.bottomLeft = acr.newAnchorPoint(owner, HTBOTTOMLEFT)
	acr.bottomRight = acr.newAnchorPoint(owner, HTBOTTOMRIGHT)
	m.anchor = acr
	m.anchor.refreshAnchorsPoint()
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
