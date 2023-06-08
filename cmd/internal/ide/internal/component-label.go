package internal

import (
	"fmt"
	"github.com/energye/golcl/lcl"
)

type IDELabel struct {
	*IDEComponent
	Label *lcl.TLabel
}

func (m *IDEForm) CreateLabel() *IDELabel {
	com := &IDELabel{}
	com.IDEComponent = m.newIDEComponentContainer(false, 50, 50, 100, 24)
	com.Label = lcl.NewLabel(m.ParentToPanel())
	com.Label.SetParent(m.ParentToPanel())
	com.Label.SetOnMouseMove(com.IDEComponent.MouseMove)
	com.Label.SetOnMouseDown(com.IDEComponent.MouseDown)
	com.Label.SetOnMouseUp(com.IDEComponent.MouseUp)
	com.Component = com.Label
	com.ComponentType = CtLabel
	m.addComponent(com.IDEComponent)
	com.Name = fmt.Sprintf("Label%d", com.Id)
	com.createAfter()
	com.IsResize = false
	return com
}
