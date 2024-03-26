package src

import (
	"fmt"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/pkgs/win"
	"github.com/energye/energy/v2/types"
	"github.com/energye/energy/v2/types/keys"
	"github.com/energye/energy/v2/types/messages"
)

type TMainForm struct {
	lcl.IForm
	ImgList  lcl.IImageList
	ActList  lcl.IActionList
	Tlbar    lcl.IToolBar
	Tlbtn    lcl.IToolButton
	Stbar    lcl.IStatusBar
	Btn      lcl.IButton
	Chk      lcl.ICheckBox
	Act      lcl.IAction
	MainMenu lcl.IMainMenu
}

var MainForm *TMainForm

func (f *TMainForm) OnFormCreate(sender lcl.IObject) {
	f.SetCaption("Hello")
	f.SetPosition(types.PoScreenCenter)
	f.EnabledMaximize(false)
	f.SetWidth(300)
	f.SetHeight(200)
	// 全局设置提示
	f.SetShowHint(true)

	// 动态创建
	f.initComponents()
	f.SetOnWndProc(func(msg *types.TMessage) {
		f.InheritedWndProc(msg)
		if msg.Msg == messages.WM_HOTKEY {
			v := f.Visible()
			if v {
				f.Hide()
			} else {
				f.Show()
			}
		}
	})
	f.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) {
		win.UnregisterHotKey(f.Handle(), 1)
	})
	win.RegisterHotKey(f.Handle(), 1, win.MOD_ALT|win.MOD_CONTROL, keys.VkH)
	go func() {
		lcl.RunOnMainThreadSync(func() {
			fmt.Println("主线程中运行 同步")
		})
		lcl.RunOnMainThreadAsync(func(id uint32) {
			fmt.Println("主线程中运行 异步", id)
		})
	}()
}

func (f *TMainForm) OnActExecute(sender lcl.IObject) {
	lcl.ShowMessage("点击了action")
}

func (f *TMainForm) OnActUpdate(sender lcl.IObject) {
	lcl.AsAction(sender).SetEnabled(f.Chk.Checked())
}

func (f *TMainForm) initComponents() {
	f.ImgList = lcl.NewImageList(f)

	if lcl.Application.Icon().Handle() != 0 {
		f.ImgList.AddIcon(lcl.Application.Icon())
	}

	f.ActList = lcl.NewActionList(f)
	f.ActList.SetImages(f.ImgList)

	// 顶部工具条
	f.Tlbar = lcl.NewToolBar(f)
	f.Tlbar.SetParent(f)
	f.Tlbar.SetImages(f.ImgList)

	f.Tlbtn = lcl.NewToolButton(f)
	f.Tlbtn.SetParent(f.Tlbar)

	// 底部状态条
	f.Stbar = lcl.NewStatusBar(f)
	f.Stbar.SetParent(f)
	f.Stbar.SetAutoHint(true)
	f.Stbar.SetSimplePanel(true)

	f.Btn = lcl.NewButton(f)
	f.Btn.SetParent(f)
	f.Btn.SetLeft(80)
	f.Btn.SetTop(f.Tlbar.Top() + f.Tlbar.Height() + 10)

	f.Chk = lcl.NewCheckBox(f)
	f.Chk.SetParent(f)
	f.Chk.SetCaption("action状态演示")
	f.Chk.SetLeft(f.Btn.Left())
	f.Chk.SetTop(f.Btn.Top() + f.Btn.Height() + 10)
	f.Chk.SetChecked(true)

	// action
	f.Act = lcl.NewAction(f)
	f.Act.SetCaption("action")
	f.Act.SetImageIndex(0)
	f.Act.SetHint("这是一个提示|长提示了")
	f.Act.SetOnExecute(f.OnActExecute)
	f.Act.SetOnUpdate(f.OnActUpdate)

	// mainMenu
	f.MainMenu = lcl.NewMainMenu(f)
	f.MainMenu.SetImages(f.ImgList)

	menu := lcl.NewMenuItem(f)
	menu.SetCaption("菜单")
	f.MainMenu.Items().Add(menu)
	subMenu := lcl.NewMenuItem(f)
	subMenu.SetAction(f.Act)
	menu.Add(subMenu)

	f.Btn.SetAction(f.Act)
	f.Tlbtn.SetAction(f.Act)
}
