package main

import (
	"fmt"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
	"math/rand"
	"time"
)

type TMainForm struct {
	lcl.TForm
}

type TForm1 struct {
	lcl.TForm
}

var MainForm TMainForm
var Form1 TForm1

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetScaled(true)
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&MainForm, &Form1)
	lcl.Application.Run()
}

func (m *TMainForm) OnFormCreate(sender lcl.IObject) {
	fmt.Println("main create")
	m.SetCaption("Main")
	m.SetPosition(types.PoScreenCenter)
	btn := lcl.NewButton(m)
	btn.SetParent(m)
	btn.SetCaption("test")
	btn.SetOnClick(func(sender lcl.IObject) {
		rand.Seed(time.Now().UnixNano())
		Form1.SetLeft(rand.Int31n(100))
		Form1.SetTop(rand.Int31n(100))
		Form1.Show()
		Form1.SetFocus()
	})
	//m.SetOnWndProc(func(msg *types.TMessage) {
	//	m.InheritedWndProc(msg)
	//	fmt.Println("OnWndProc:", msg.Msg)
	//})
}

func (m *TMainForm) CreateParams(params *types.TCreateParams) {
	fmt.Println("调用此过程  TMainForm.CreateParams")
}

func (m *TForm1) OnFormCreate(sender lcl.IObject) {
	fmt.Println("form1 create")
	m.SetCaption("Form1")
}

func (m *TForm1) CreateParams(params *types.TCreateParams) {
	fmt.Println("调用此过程 TForm1.CreateParams")
}
