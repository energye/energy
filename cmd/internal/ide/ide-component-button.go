package ide

import (
	"fmt"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

type IDEButton struct {
	*IDEComponent
	Component *lcl.TButton
}

func (m *IDEForm) CreateButton() *IDEButton {
	com := &IDEButton{}
	com.IDEComponent = m.newIDEComponentContainer(false, 50, 50, 100, 24)
	com.Component = lcl.NewButton(com.componentParentPanel)
	com.Component.SetParent(com.componentParentPanel)
	com.Component.SetAlign(types.AlClient)
	com.Component.SetOnMouseMove(com.IDEComponent.mouseMove)
	com.Component.SetOnMouseDown(com.IDEComponent.mouseDown)
	com.Component.SetOnMouseUp(com.IDEComponent.mouseUp)
	com.componentControl = com.Component
	m.addComponent(com.IDEComponent)
	com.name = fmt.Sprintf("Button%d", com.Id)
	com.componentType = ctButton
	com.Component.SetCaption(com.name)
	com.createAfter()
	return com
}
