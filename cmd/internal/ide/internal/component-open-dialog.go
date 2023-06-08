package internal

import (
	"fmt"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

type IDEOpenDialog struct {
	*IDEComponent
	Image *lcl.TImage
}

func (m *IDEForm) CreateDialogOpen() *IDEOpenDialog {
	com := &IDEOpenDialog{}
	com.IDEComponent = m.newIDEComponentContainer(true, 50, 50, 28, 28)
	com.Image = lcl.NewImage(com.IDEComponent.ParentToPanel())
	com.Image.SetParent(com.IDEComponent.ParentToPanel())
	com.Image.SetAlign(types.AlClient)
	com.Image.SetOnMouseMove(com.IDEComponent.MouseMove)
	com.Image.SetOnMouseDown(com.IDEComponent.MouseDown)
	com.Image.SetOnMouseUp(com.IDEComponent.MouseUp)
	com.Component = com.Image
	com.ComponentType = CtOpenDialog
	m.addComponent(com.IDEComponent)
	com.Name = fmt.Sprintf("DialogOpen%d", com.Id)
	com.createAfter()
	com.DClick = func(button types.TMouseButton, shift types.TShiftState, x, y int32) {
		fmt.Println("双击了 打开文件窗口")
	}
	return com
}
