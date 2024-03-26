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
	Btn1 lcl.IButton
}

var MainForm *TMainForm

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.CreateForm(&MainForm, true)
	lcl.Application.Run()
}

func (f *TMainForm) OnFormCreate(sender lcl.IObject) {
	f.ScreenCenter()
	f.SetCaption("taskDialog演示")

	f.Btn1 = lcl.NewButton(f)
	f.Btn1.SetParent(f)
	f.Btn1.SetCaption("TaskDialog")
	f.Btn1.SetLeft(10)
	f.Btn1.SetTop(10)
	f.Btn1.SetOnClick(f.OnBtn1Click)

}

func (f *TMainForm) OnFormDestroy(sender lcl.IObject) {

}

func (f *TMainForm) OnBtn1Click(sender lcl.IObject) {
	taskdlg := lcl.NewTaskDialog(f)
	defer taskdlg.Free()
	taskdlg.SetTitle("确认移除")
	taskdlg.SetCaption("询问")
	taskdlg.SetText("移除选择的项目？")
	taskdlg.SetExpandButtonCaption("展开按钮标题")
	taskdlg.SetExpandedText("展开的文本")

	taskdlg.SetFooterText("底部文本")

	rd := lcl.AsTaskDialogRadioButtonItem(taskdlg.RadioButtons().Add())
	rd.SetCaption("单选按钮1")
	rd = lcl.AsTaskDialogRadioButtonItem(taskdlg.RadioButtons().Add())
	rd.SetCaption("单选按钮2")
	rd = lcl.AsTaskDialogRadioButtonItem(taskdlg.RadioButtons().Add())
	rd.SetCaption("单选按钮3")

	taskdlg.SetCommonButtons(0) //rtl.Include(0, 0))
	btn := lcl.AsTaskDialogButtonItem(taskdlg.Buttons().Add())
	btn.SetCaption("移除")
	btn.SetModalResult(types.MrYes)

	btn = lcl.AsTaskDialogButtonItem(taskdlg.Buttons().Add())
	btn.SetCaption("保持")
	btn.SetModalResult(types.MrNo)

	taskdlg.SetMainIcon(types.TdiQuestion)

	if taskdlg.Execute() {
		if taskdlg.ModalResult() == types.MrYes {
			lcl.ShowMessage("项目已移除")

			fmt.Println(taskdlg.RadioButton().Caption())
		}
	}
}
