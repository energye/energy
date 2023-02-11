package ide

import (
	"fmt"
	"github.com/energye/golcl/lcl"
)

type IDEOpenDialog struct {
	*IDEComponent
	Component *lcl.TOpenDialog
}

func (m *IDEForm) CreateDialogOpen() *IDEOpenDialog {
	com := &IDEOpenDialog{}
	com.IDEComponent = m.newIDEComponentContainer(true, 50, 50, 28, 28)
	com.Component = lcl.NewOpenDialog(com.IDEComponent.componentParentPanel)
	com.component = com.Component
	m.addComponent(com.IDEComponent)
	com.componentType = ctOpenDialog
	com.name = fmt.Sprintf("DialogOpen%d", com.Id)
	com.componentParentPanel.SetCaption(com.name)
	//com.createAnchor(m.componentParentPanel)
	com.createAfter()
	return com
}
