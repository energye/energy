package ide

import (
	"fmt"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

type IDEImage struct {
	*IDEComponent
	Component *lcl.TImage
}

func (m *IDEForm) CreateImage() *IDEImage {
	com := &IDEImage{}
	com.IDEComponent = m.newIDEComponentContainer(true, 50, 50, 170, 50)
	com.Component = lcl.NewImage(com.IDEComponent.componentParentPanel)
	com.Component.SetParent(com.IDEComponent.componentParentPanel)
	com.Component.SetAlign(types.AlClient)
	com.Component.SetOnMouseMove(com.IDEComponent.mouseMove)
	com.Component.SetOnMouseDown(com.IDEComponent.mouseDown)
	com.Component.SetOnMouseUp(com.IDEComponent.mouseUp)
	com.component = com.Component
	m.addComponent(com.IDEComponent)
	com.name = fmt.Sprintf("Image%d", com.Id)
	com.componentType = ctImage
	com.createAfter()
	//com.createAnchor(m.componentParentPanel)
	return com
}
