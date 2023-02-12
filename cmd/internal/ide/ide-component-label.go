package ide

import (
	"fmt"
	"github.com/energye/golcl/lcl"
)

type IDELabel struct {
	*IDEComponent
	Component *lcl.TLabel
}

func (m *IDEForm) CreateLabel() *IDELabel {
	com := &IDELabel{}
	com.IDEComponent = m.newIDEComponentContainer(false, 50, 50, 100, 24)
	com.Component = lcl.NewLabel(m.parentToPanel())
	com.Component.SetParent(m.parentToPanel())
	com.Component.SetOnMouseMove(com.IDEComponent.mouseMove)
	com.Component.SetOnMouseDown(com.IDEComponent.mouseDown)
	com.Component.SetOnMouseUp(com.IDEComponent.mouseUp)
	com.component = com.Component
	com.componentType = ctLabel
	m.addComponent(com.IDEComponent)
	com.name = fmt.Sprintf("Label%d", com.Id)
	com.createAfter()
	com.isResize = false
	return com
}
