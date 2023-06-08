package ide

import (
	"fmt"
	"github.com/energye/golcl/lcl"
)

type IDEEdit struct {
	*IDEComponent
	Edit *lcl.TEdit
}

func (m *IDEForm) CreateEdit() *IDEEdit {
	com := &IDEEdit{}
	com.IDEComponent = m.newIDEComponentContainer(false, 50, 50, 150, 24)
	com.Edit = lcl.NewEdit(m.ParentToPanel())
	com.Edit.SetParent(m.ParentToPanel())
	com.Edit.SetOnMouseMove(com.IDEComponent.MouseMove)
	com.Edit.SetOnMouseDown(com.IDEComponent.MouseDown)
	com.Edit.SetOnMouseUp(com.IDEComponent.MouseUp)
	com.Component = com.Edit
	com.ComponentType = CtEdit
	m.addComponent(com.IDEComponent)
	com.Name = fmt.Sprintf("Edit%d", com.Id)
	com.createAfter()
	return com
}
