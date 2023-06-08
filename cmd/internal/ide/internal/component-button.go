package internal

import (
	"fmt"
	"github.com/energye/golcl/lcl"
)

type IDEButton struct {
	*IDEComponent
	Button *lcl.TButton
}

func (m *IDEForm) CreateButton() *IDEButton {
	com := &IDEButton{}
	com.IDEComponent = m.newIDEComponentContainer(false, 50, 50, 100, 24)
	com.Button = lcl.NewButton(m.ParentToPanel())
	com.Button.SetParent(m.ParentToPanel())
	com.Button.SetOnMouseMove(com.IDEComponent.MouseMove)
	com.Button.SetOnMouseDown(com.IDEComponent.MouseDown)
	com.Button.SetOnMouseUp(com.IDEComponent.MouseUp)
	com.Component = com.Button
	com.ComponentType = CtButton
	m.addComponent(com.IDEComponent)
	com.Name = fmt.Sprintf("Button%d", com.Id)
	com.createAfter()
	return com
}
