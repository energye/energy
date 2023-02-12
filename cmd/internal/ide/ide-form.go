package ide

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
)

type IDEForm struct {
	count int
	IDEComponent
	active     *IDEComponent
	components map[int]*IDEComponent
	tabSheet   *lcl.TTabSheet
}

func (m *IDEForm) addComponent(component *IDEComponent) int {
	m.count++
	m.components[m.count] = component
	component.Id = m.count
	return m.count
}

func (m *IDEForm) RemoveComponent(index int) {
	if component, ok := m.components[index]; ok {
		if component.borderPanel != nil {
			m.componentFrees(component.borderPanel)
		}
		if component.componentParentPanel != nil {
			m.componentFrees(component.componentParentPanel)
		}
		component.anchor.remove()
	}
	delete(m.components, index)
	m.active = nil
	println("剩余组件个数", len(m.components))
}

func (m *IDEForm) componentFrees(control lcl.IComponent) {
	control.Free()
}

func (m *IDEForm) newIDEComponentContainer(useBorder bool, left, top, width, height int32) *IDEComponent {
	ideComponent := &IDEComponent{}
	ideComponent.isResize = true
	if useBorder {
		ideComponent.borderPanel = lcl.NewPanel(m.parentToPanel())
		ideComponent.borderPanel.SetParent(m.parentToPanel())
		ideComponent.borderPanel.SetDoubleBuffered(true)
		ideComponent.borderPanel.SetBevelInner(types.BvNone)
		ideComponent.borderPanel.SetBevelOuter(types.BvNone)
		ideComponent.borderPanel.SetBorderStyle(types.BsNone)
		if useBorder {
			ideComponent.borderPanel.SetBounds(left-border, top-border, width+border, height+border)
			ideComponent.borderPanel.SetColor(colors.ClBlack)
		} else {
			ideComponent.borderPanel.SetBounds(left, top, width, height)
		}

		ideComponent.componentParentPanel = lcl.NewPanel(m.parentToPanel())
		ideComponent.parentToPanel().SetParent(m.parentToPanel())
		ideComponent.parentToPanel().SetDoubleBuffered(true)
		ideComponent.parentToPanel().SetBevelInner(types.BvNone)
		ideComponent.parentToPanel().SetBevelOuter(types.BvNone)
		ideComponent.parentToPanel().SetBorderStyle(types.BsNone)
		ideComponent.parentToPanel().SetColor(colors.ClSysDefault)
		if useBorder {
			ideComponent.parentToPanel().SetBounds(left-border/2, top-border/2, width, height)
		} else {
			ideComponent.parentToPanel().SetBounds(left, top, width, height)
		}
		//ideComponent.parentToPanel().SetOnMouseMove(ideComponent.mouseMove)
		//ideComponent.parentToPanel().SetOnMouseDown(ideComponent.mouseDown)
		//ideComponent.parentToPanel().SetOnMouseUp(ideComponent.mouseUp)
		ideComponent.ox, ideComponent.oy, ideComponent.ow, ideComponent.oh = ideComponent.parentToPanel().Left(), ideComponent.parentToPanel().Top(), ideComponent.parentToPanel().Width(), ideComponent.parentToPanel().Height()
	}
	ideComponent.form = m
	ideComponent.isUseBorder = useBorder
	return ideComponent
}
