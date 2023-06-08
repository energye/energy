package ide

import (
	"fmt"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

type IDEImage struct {
	*IDEComponent
	Image *lcl.TImage
}

func (m *IDEForm) CreateImage() *IDEImage {
	com := &IDEImage{}
	com.IDEComponent = m.newIDEComponentContainer(true, 50, 50, 170, 50)
	com.Image = lcl.NewImage(com.IDEComponent.ParentToPanel())
	com.Image.SetParent(com.IDEComponent.ParentToPanel())
	com.Image.SetAlign(types.AlClient)
	com.Image.SetOnMouseMove(com.IDEComponent.MouseMove)
	com.Image.SetOnMouseDown(com.IDEComponent.MouseDown)
	com.Image.SetOnMouseUp(com.IDEComponent.MouseUp)
	com.Component = com.Image
	com.ComponentType = CtImage
	m.addComponent(com.IDEComponent)
	com.Name = fmt.Sprintf("Image%d", com.Id)
	com.createAfter()
	return com
}
