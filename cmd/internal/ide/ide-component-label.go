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
	com.IDEComponent = m.newIDEComponentContainer(true, 50, 50, 170, 50)
	com.Component = lcl.NewLabel(com.IDEComponent.componentParentPanel)
	com.Component.SetParent(com.IDEComponent.componentParentPanel)
	com.Component.SetAlign(types.AlClient)
	com.Component.SetOnMouseMove(com.IDEComponent.mouseMove)
	com.Component.SetOnMouseDown(com.IDEComponent.mouseDown)
	com.Component.SetOnMouseUp(com.IDEComponent.mouseUp)
	com.component = com.Component
	m.addComponent(com.IDEComponent)
	com.componentType = ctLabel
	com.name = fmt.Sprintf("Label%d", com.Id)
	com.createAfter()
	//com.createAnchor(m.componentParentPanel)
	return com
}
