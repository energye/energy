package ide

import (
	"fmt"
	"github.com/energye/golcl/lcl"
)

type IDEButton struct {
	*IDEComponent
	Component *lcl.TButton
}

func (m *IDEForm) CreateButton() *IDEButton {
	com := &IDEButton{}
	com.IDEComponent = m.newIDEComponentContainer(false, 50, 50, 100, 24)
	com.Component = lcl.NewButton(m.parentToPanel())
	com.Component.SetParent(m.parentToPanel())
	com.Component.SetOnMouseMove(com.IDEComponent.mouseMove)
	com.Component.SetOnMouseDown(com.IDEComponent.mouseDown)
	com.Component.SetOnMouseUp(com.IDEComponent.mouseUp)
	com.component = com.Component
	com.componentType = ctButton
	m.addComponent(com.IDEComponent)
	com.name = fmt.Sprintf("Button%d", com.Id)
	com.createAfter()
	return com
}
