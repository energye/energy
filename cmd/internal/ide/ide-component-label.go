package ide

import (
	"fmt"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

type IDELabel struct {
	*IDEComponent
	Component *lcl.TLabel
}

func (m *IDEForm) CreateLabel() *IDELabel {
	com := &IDELabel{}
	com.IDEComponent = m.newIDEComponentContainer(false, 50, 50, 100, 24)
	com.Component = lcl.NewLabel(com.IDEComponent.componentParentPanel)
	com.Component.SetParent(com.IDEComponent.componentParentPanel)
	com.Component.SetAlign(types.AlClient)
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
