package ide

import (
	"fmt"
	"github.com/energye/golcl/lcl"
)

type IDEEdit struct {
	*IDEComponent
	Component *lcl.TEdit
}

func (m *IDEForm) CreateEdit() *IDEEdit {
	com := &IDEEdit{}
	com.IDEComponent = m.newIDEComponentContainer(false, 50, 50, 150, 24)
	com.Component = lcl.NewEdit(m.parentToPanel())
	com.Component.SetParent(m.parentToPanel())
	com.Component.SetOnMouseMove(com.IDEComponent.mouseMove)
	com.Component.SetOnMouseDown(com.IDEComponent.mouseDown)
	com.Component.SetOnMouseUp(com.IDEComponent.mouseUp)
	com.component = com.Component
	com.componentType = ctEdit
	m.addComponent(com.IDEComponent)
	com.name = fmt.Sprintf("Edit%d", com.Id)
	com.createAfter()
	return com
}
