package main

import (
	"fmt"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
)

type TMainForm struct {
	*lcl.TForm
}

var mainForm *TMainForm

func main() {
	inits.Init(nil, nil)
	lcl.RunApp(&mainForm)
}

func (f *TMainForm) OnFormCreate(object lcl.IObject) {
	f.SetCaption("TCheckListBox测试")
	f.ScreenCenter()
	f.EnabledMaximize(false)
	f.EnabledMinimize(false)

	chkListBox := lcl.NewCheckListBox(f)
	chkListBox.SetParent(f)
	chkListBox.SetAlign(types.AlClient)
	chkListBox.SetOnClickCheck(func(sender lcl.IObject) {
		fmt.Println("check单击。")
	})

	for i := 1; i < 100; i++ {
		chkListBox.Items().Add(fmt.Sprintf("第%d个项目", i))
	}
	button := lcl.NewButton(f)
	button.SetParent(f)
	button.SetAlign(types.AlBottom)
	button.SetCaption("项目启用")
	button.SetOnClick(func(sender lcl.IObject) {
		//fmt.Println("选中数：", chkListBox.Checked())
		chkListBox.SetItemEnabled(0, !chkListBox.ItemEnabled(0))
	})

	// 获取/设置项目启用
	//chkListBox.SetItemEnabled()
	//chkListBox.ItemEnabled()

	// 获取/设置项目状态
	//chkListBox.State()
	//chkListBox.SetState()

	// 获取/设置项目选中
	//chkListBox.Checked()
	//chkListBox.SetChecked()

	button = lcl.NewButton(f)
	button.SetParent(f)
	button.SetCaption("全选")
	button.SetAlign(types.AlBottom)
	button.SetOnClick(func(sender lcl.IObject) {
		chkListBox.CheckAll(types.CbChecked, true, true)
	})

	button = lcl.NewButton(f)
	button.SetParent(f)
	button.SetCaption("取消全选")
	button.SetAlign(types.AlBottom)
	button.SetOnClick(func(sender lcl.IObject) {
		chkListBox.CheckAll(types.CbUnchecked, true, true)
	})

}
