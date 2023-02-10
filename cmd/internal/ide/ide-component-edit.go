package ide

import (
	"fmt"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

type IDEEdit struct {
	*IDEComponent
	Component *lcl.TEdit
}

func (m *IDEForm) CreateEdit() *IDEEdit {
	com := &IDEEdit{}
	com.IDEComponent = m.newIDEComponentContainer(true, 50, 50, 150, 24)
	com.Component = lcl.NewEdit(com.IDEComponent.componentParentPanel)
	com.Component.SetParent(com.IDEComponent.componentParentPanel)
	com.Component.SetAlign(types.AlClient)
	com.Component.SetOnMouseMove(com.IDEComponent.mouseMove)
	com.Component.SetOnMouseDown(com.IDEComponent.mouseDown)
	com.Component.SetOnMouseUp(com.IDEComponent.mouseUp)
	com.componentControl = com.Component
	m.addComponent(com.IDEComponent)
	com.componentType = ctLabel
	com.name = fmt.Sprintf("Edit%d", com.Id)
	com.componentParentPanel.SetCaption(com.name)
	//com.createAnchor(m.componentParentPanel)
	com.createAfter()
	return com
}
