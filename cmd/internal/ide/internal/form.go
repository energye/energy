package internal

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
)

type IDEForm struct {
	IDEComponent
	Count      int
	Active     *IDEComponent
	Components map[int]*IDEComponent
	TabSheet   *lcl.TTabSheet
}

func (m *IDEForm) addComponent(component *IDEComponent) int {
	m.Count++
	m.Components[m.Count] = component
	component.Id = m.Count
	return m.Count
}

func (m *IDEForm) RemoveComponent(index int) {
	if component, ok := m.Components[index]; ok {
		if component.BorderPanel != nil {
			m.componentFrees(component.BorderPanel)
		}
		if component.ComponentParentPanel != nil {
			m.componentFrees(component.ComponentParentPanel)
		}
		component.Anchor.remove()
	}
	delete(m.Components, index)
	m.Active = nil
	println("剩余组件个数", len(m.Components))
}

func (m *IDEForm) componentFrees(control lcl.IComponent) {
	control.Free()
}

func (m *IDEForm) newIDEComponentContainer(useBorder bool, left, top, width, height int32) *IDEComponent {
	ideComponent := &IDEComponent{}
	ideComponent.IsResize = true
	if useBorder {
		ideComponent.BorderPanel = lcl.NewPanel(m.ParentToPanel())
		ideComponent.BorderPanel.SetParent(m.ParentToPanel())
		ideComponent.BorderPanel.SetDoubleBuffered(true)
		ideComponent.BorderPanel.SetBevelInner(types.BvNone)
		ideComponent.BorderPanel.SetBevelOuter(types.BvNone)
		ideComponent.BorderPanel.SetBorderStyle(types.BsNone)
		if useBorder {
			ideComponent.BorderPanel.SetBounds(left-Border, top-Border, width+Border, height+Border)
			ideComponent.BorderPanel.SetColor(colors.ClBlack)
		} else {
			ideComponent.BorderPanel.SetBounds(left, top, width, height)
		}

		ideComponent.ComponentParentPanel = lcl.NewPanel(m.ParentToPanel())
		ideComponent.ParentToPanel().SetParent(m.ParentToPanel())
		ideComponent.ParentToPanel().SetDoubleBuffered(true)
		ideComponent.ParentToPanel().SetBevelInner(types.BvNone)
		ideComponent.ParentToPanel().SetBevelOuter(types.BvNone)
		ideComponent.ParentToPanel().SetBorderStyle(types.BsNone)
		//ideComponent.ParentToPanel().SetColor(colors.ClSysDefault)
		if useBorder {
			ideComponent.ParentToPanel().SetBounds(left-Border/2, top-Border/2, width, height)
		} else {
			ideComponent.ParentToPanel().SetBounds(left, top, width, height)
		}
		//ideComponent.ParentToPanel().SetOnMouseMove(ideComponent.MouseMove)
		//ideComponent.ParentToPanel().SetOnMouseDown(ideComponent.MouseDown)
		//ideComponent.ParentToPanel().SetOnMouseUp(ideComponent.MouseUp)
		ideComponent.Ox, ideComponent.Oy, ideComponent.Ow, ideComponent.Oh = ideComponent.ParentToPanel().Left(), ideComponent.ParentToPanel().Top(), ideComponent.ParentToPanel().Width(), ideComponent.ParentToPanel().Height()
	}
	ideComponent.Form = m
	ideComponent.IsUseBorder = useBorder
	return ideComponent
}
