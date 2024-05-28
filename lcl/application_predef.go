//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/energy/v2/api"
	"reflect"
)

/*
 TApplication.CreateForm 一般不建议使用NewForm，而优先使用CreateForm

  -------------------- 用法一--------------------------------------
  1、mainForm := lcl.Application.CreateForm()    // 直接返回，不推荐使用
  例:
    mainForm := lcl.Application.CreateForm()
    mainForm.SetOnClick(func(sender lcl.IObject) {
        lcl.ShowMessage("msg")
    })


  -------------------- 用法二--------------------------------------
  2、lcl.Application.CreateForm(&mainForm)       // 无资源或者自动加载对应类名的资源，当无资源时只会绑定窗口的事件，不会绑定子组件事件，有资源则自动绑定所有事件
  例：
    type TMainForm struct {
        *lcl.TForm
    }

    var mainForm *TMainForm
    lcl.Application.CreateForm(&mainForm)

    func (f *TMainForm)OnFormCreate(sender lcl.IObject) {
        fmt.Println("FormCreate")
    }

    func (f *TMainForm)OnFormClick(sender lcl.IObject) {
        lcl.ShowMessage("click")
    }


  -------------------- 用法三--------------------------------------
  3、lcl.Application.CreateForm(&mainForm, true) // 无资源或者自动加载对应类名的资源，当第二个参数为true时在OnFormCreate调用完后会绑定子组件事件(当查找到对应的资源则第二个参数无效)
  例：
    type TMainForm struct {
        *lcl.TForm
        Btn1 *lcl.TButton
    }

    var mainForm *TMainForm
    lcl.Application.CreateForm(&mainForm, true)

    func (f *TMainForm)OnFormCreate(sender lcl.IObject) {
        fmt.Println("FormCreate")
        f.Btn1 = lcl.NewButton(f)
        f.Btn1.SetParent(f)
    }

    func (f *TMainForm)OnFormClick(sender lcl.IObject) {
        lcl.ShowMessage("click")
    }

    func (f *TMainForm)OnBtn1Click(Sender lcl.IObject) {
        lcl.ShowMessage("Btn1 Click")
    }
*/

// CreateForm
//
// 创建一个TForm。
//
// Create a TForm.
func (m *TApplication) CreateForm(forms ...IForm) IForm {
	//runtime.LockOSThread()
	//defer runtime.UnlockOSThread()
	size := len(forms)
	if size == 0 {
		return AsForm(Application_CreateForm(m.Instance()))
	}
	for i := 0; i < size; i++ {
		form := forms[i]
		var (
			mainForm        = Application.MainForm()
			isMain          = mainForm == nil || mainForm.Instance() == 0 // 0 | nil = main
			v               = reflect.ValueOf(form)
			createParamsPtr uintptr
		)
		if !isMain {
			createParamsPtr = v.Pointer()
		}
		// OnCreate 实现回调
		if _, ok := form.(IOnCreate); ok {
			addRequestFormCreateMap(createParamsPtr, form)
		}
		// CreateParams 实现回调
		if _, ok := form.(IOnCreateParams); ok {
			addRequestCreateParamsMap(createParamsPtr, form)
		}

		formPtr := Application_CreateForm(m.Instance())
		form.SetInstance(unsafePointer(formPtr))
		if !isMain {
			Form_SetGoPtr(formPtr, createParamsPtr)
		}
	}
	return nil
}

func (m *TApplication) CreateResForm(forms ...IForm) {

}

// Run
//
// 运行APP。
//
// Run the app.
func (m *TApplication) Run() {
	Application_Run(m.Instance())
}

// Initialize
//
// 初始APP信息。
//
// Initial APP information.
func (m *TApplication) Initialize() {
	CustomWidgetSetInitialization()
	Application_Initialize(m.Instance())
}

// SetRunLoopReceived 这里只是测试，实际Go并未用得着他
func (m *TApplication) SetRunLoopReceived(proc uintptr) {
	Application_SetRunLoopReceived(m.Instance(), proc)
}
