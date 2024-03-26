package form

import (
	"fmt"
	"github.com/energye/energy/v2/lcl"

	"github.com/energye/energy/v2/types/keys"

	"github.com/energye/energy/v2/types"
)

// ::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender lcl.IObject) {
	fmt.Println(Form1.Caption(), f.PixelsPerInch())

	// 遍历组件
	// 只要owner设置为Form的都可以通过这个方法来遍历。
	var i int32
	for i = 0; i < f.ComponentCount(); i++ {
		comp := f.Components(i)
		//fmt.Println(i, "=", comp.Name())
		if comp.Is().Memo() {
			fmt.Println(i, "=", comp.Name(), ", 继承自TMemo")
			mem := lcl.AsMemo(comp)
			mem.SetOnKeyUp(f.memoOnKeyup)
		}
	}
}

func (f *TForm1) memoOnKeyup(sender lcl.IObject, key *types.Char, shift types.TShiftState) {
	if shift.In(types.SsCtrl) && *key == /*keys.VkA*/ keys.VkB {
		lcl.AsMemo(sender).SelectAll()
	}
}

func (f *TForm1) OnActExitExecute(lcl.IObject) {
	lcl.Application.Terminate()
}

func (f *TForm1) OnButton2Click(lcl.IObject) {
	result := Form2.ShowModal()
	if result == types.MrOk {
		lcl.ShowMessage("Form2返回了OK")
	} else if result == types.MrClose || result == types.MrNone {
		lcl.ShowMessage("Form2返回了Close")
	} else if result == types.MrCancel {
		lcl.ShowMessage("Form2返回了Cancel")
	}
}

func (f *TForm1) OnActFileNewExecute(lcl.IObject) {
	lcl.ShowMessage("ActFileNew Execute.")
}

func (f *TForm1) OnFormCloseQuery(sender lcl.IObject, canClose *bool) {
	fmt.Println("关闭。")
}

func (f *TForm1) OnCheckBox1Click(sender lcl.IObject) {
	f.Button1.SetEnabled(f.CheckBox1.Checked())
}

func (f *TForm1) OnButton1Click(sender lcl.IObject) {
	//lcl.ShowMessage("Hello!")
	jpg := lcl.NewJPEGImage()
	defer jpg.Free()
	jpg.LoadFromFile("a.jpg")
}
