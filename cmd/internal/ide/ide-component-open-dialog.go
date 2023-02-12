package ide

import (
	"fmt"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

type IDEOpenDialog struct {
	*IDEComponent
	Component *lcl.TImage
}

func (m *IDEForm) CreateDialogOpen() *IDEOpenDialog {
	com := &IDEOpenDialog{}
	com.IDEComponent = m.newIDEComponentContainer(true, 50, 50, 28, 28)
	com.Component = lcl.NewImage(com.IDEComponent.parentToPanel())
	com.Component.SetParent(com.IDEComponent.parentToPanel())
	com.Component.SetAlign(types.AlClient)
	com.Component.SetOnMouseMove(com.IDEComponent.mouseMove)
	com.Component.SetOnMouseDown(com.IDEComponent.mouseDown)
	com.Component.SetOnMouseUp(com.IDEComponent.mouseUp)
	com.component = com.Component
	com.componentType = ctOpenDialog
	m.addComponent(com.IDEComponent)
	com.name = fmt.Sprintf("DialogOpen%d", com.Id)
	com.createAfter()
	return com
}
