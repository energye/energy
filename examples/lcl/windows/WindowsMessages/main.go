//go:build windows
// +build windows

package main

import (
	"fmt"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/pkgs/win"
	"github.com/energye/energy/v2/types"
	"syscall"
)

var (
	oldWndPrc uintptr
)

type TMainForm struct {
	lcl.TForm
}

var mainForm TMainForm

func (f *TMainForm) FormCreate(object lcl.IObject) {
	f.SetCaption("Windows Messages")
	f.SetWidth(300)
	f.SetHeight(200)
	f.ScreenCenter()

	newWndProc := syscall.NewCallback(WndProc)
	oldWndPrc = win.SetWindowLongPtr(mainForm.Handle(), win.GWL_WNDPROC, newWndProc)
	fmt.Println("newWndProc:", newWndProc)
	fmt.Println("oldWndPro:", oldWndPrc)

	btn := lcl.NewButton(&mainForm)
	btn.SetParent(&mainForm)
	btn.SetCaption("按钮1")
	btn.SetLeft(50)
	btn.SetTop(50)
}

func (f *TMainForm) OnFormDestroy(object lcl.IObject) {
	fmt.Println("FormOnDestroy")
	// 完成后要恢复的
	win.SetWindowLongPtr(mainForm.Handle(), win.GWL_WNDPROC, oldWndPrc)
}

func main() {
	inits.Init(nil, nil)
	lcl.RunApp(&mainForm)
}

func WndProc(hWnd uintptr, message uint32, wParam, lParam uintptr) uintptr {
	switch message {
	case win.WM_SYSCOMMAND:
		switch wParam {
		case win.SC_MAXIMIZE:
			fmt.Println("最大化")
		case win.SC_MINIMIZE:
			fmt.Println("最小化")
		case win.SC_RESTORE:
			fmt.Println("恢复")
		case win.SC_CLOSE:
			fmt.Println("关闭")
		}
	}
	return win.CallWindowProc(oldWndPrc, types.HWND(hWnd), message, wParam, lParam)
}
