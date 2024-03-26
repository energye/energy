package main

import (
	"fmt"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
	"net/http"
	_ "net/http/pprof"
)

type TMainForm struct {
	*lcl.TForm
	Button1 lcl.IButton
}

type TForm1 struct {
	*lcl.TForm
	Button1 lcl.IButton
}

var (
	mainForm *TMainForm
	form1    *TForm1
)

func main() {
	go http.ListenAndServe("localhost:8080", nil)
	lcl.DEBUG = true
	inits.Init(nil, nil)
	lcl.RunApp(&mainForm, &form1)
}

func (m *TMainForm) OnFormCreate(sender lcl.IObject) {
	fmt.Println("TMainForm OnFormCreate")
	m.SetOnWndProc(func(msg *types.TMessage) {
		m.InheritedWndProc(msg)
		fmt.Println("msg", msg)
	})
	m.SetCaption("Hello")
	m.EnabledMaximize(false)
	m.WorkAreaCenter()
	m.SetWidth(600)
	m.SetHeight(600)
	m.Button1 = lcl.NewButton(m)
	m.Button1.SetParent(m)
	m.Button1.SetCaption("窗口1")
	m.Button1.SetLeft(50)
	m.Button1.SetTop(50)
	m.Button1.SetOnClick(m.OnButton1Click)
}

func (f *TMainForm) OnFormCloseQuery(Sender lcl.IObject, CanClose *bool) {
	*CanClose = lcl.MessageDlg("是否退出？", types.MtConfirmation, types.MbYes, types.MbNo) == types.IdYes
}

func (f *TMainForm) OnButton1Click(object lcl.IObject) {
	form1.Show()
	fmt.Println("清除事件")
	//f.Button1.SetOnClick(nil)
	f.Button1.SetOnClick(f.OnButton1Click)
	fmt.Println("更换事件")
	f.Button1.SetOnClick(f.OnButton2Click)
}

func (f *TMainForm) OnButton2Click(object lcl.IObject) {
	fmt.Println("换成button2click事件了啊")
}

// ---------- Form1 ----------------

func (f *TForm1) OnFormCreate(sender lcl.IObject) {
	fmt.Println("TForm1 OnFormCreate")
	f.Button1 = lcl.NewButton(f)
	fmt.Println("f.Button1:", f.Button1.Instance())
	f.Button1.SetParent(f)
	f.Button1.SetCaption("我是按钮")
	f.Button1.SetOnClick(f.OnButton1Click)
}

func (f *TForm1) OnButton1Click(object lcl.IObject) {
	lcl.ShowMessage("Click")
}
