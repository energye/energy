//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//----------------------------------------
// 加载文件或者内存中的窗口资源文件功能
// 需要配合窗口设计器使用。
// 设计器是独立于govcl的，设计器的目的是用于简化窗口的创建，设计器不开源。
//
// 例：
//    type TMainForm struct {
//        *lcl.TForm  // 不要写名称，使用隐式的
//        Button1 *lcl.TButton // 注意，设计器中一定要首字母大写，否则不能使用反射填充
//    }
//
//	  var mainForm *TMainForm
//
//    func main() {
//        ...
//        // 文件加载方式
//        lcl.Application.CreateFormFromFile(filename, &mainForm)
//        // 流加载方式
//        // mem := lcl.NewMemoryStream()
//        // mem.Write([]byte(...))
//        // lcl.Application.CreateFormFromStream(mem, &mainForm)
//        // mem.Free()
//        // 资源加载方式
//        //lcl.Application.CreateFormFromResourceName("ResName", &mainForm)
//        ...
//     }

package lcl

import (
	"fmt"
	"reflect"
	"unsafe"
)

func setFiledVal(field reflect.Value, component IComponent) {
	if field.IsValid() && field.CanSet() {
		if field.Kind() == reflect.Interface {
			// TODO 需要动态创建小部件对象, 根据接口类型
			//fVal := AsButton(instance)
			field.Set(reflect.ValueOf(component))
		} else {
			instance := component.Instance()
			fVal := reflect.New(field.Type().Elem())
			var findInstance func(pVal reflect.Value)
			findInstance = func(pVal reflect.Value) {
				var num int
				if pVal.Kind() == reflect.Pointer {
					num = pVal.Elem().NumField()
				} else {
					num = pVal.NumField()
				}
				for i := 0; i < num; i++ {
					var val reflect.Value
					if pVal.Kind() == reflect.Pointer {
						val = pVal.Elem().Field(i)
					} else {
						val = pVal.Field(i)
					}
					if val.Kind() == reflect.UnsafePointer || val.Kind() == reflect.Struct {
						if val.Type().Field(i).Name == "instance" {
							*(*uintptr)(unsafe.Pointer(val.UnsafeAddr())) = instance
							return
						} else if val.Kind() == reflect.Struct {
							findInstance(val)
						}
					}
				}
			}
			findInstance(fVal)
			field.Set(fVal)
		}
	}
}

func newGoFormInstance(out interface{}) reflect.Value {
	// out是一个 **TXXForm的变量指针，未进行分配内存，表现形式为 **TXXX，每使用一个Elem()减少一个
	vt := reflect.TypeOf(out).Elem()
	v := reflect.New(vt.Elem())
	// 将实例化后的值填充到out指针变量中，这里要能修改的需要使用Elem()方法获取
	reflect.ValueOf(out).Elem().Set(v)
	return v
}

// fullFiledVal 动态创设置字段值
func fullFiledVal(f IComponent, goInstance reflect.Value, fullSubComponent, afterBindSubComponentsEvents bool) {
	if !DEBUG {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("err:", err)
			}
		}()
	}
	// 获取指针类型
	vPtr := goInstance.Elem()
	// 检查是否有效，并且可以被设置
	if vPtr.IsValid() && vPtr.CanSet() {
		// 如果没有名字，就指定一个名字，名字以当前类，如果首个为T则去除
		if f.Name() == "" {
			newName := vPtr.Type().Name()
			if newName[0] == 'T' {
				newName = newName[1:]
			}
			f.SetName(newName)
		}
		// TForm|IForm/TFrame|IFrame，默认的, 使用隐式嵌入
		className := "TForm"
		if f.ClassName() == "TFrame" {
			className = "TFrame"
		} else if f.ClassName() == "IFrame" {
			className = "IFrame"
		}
		field := vPtr.FieldByName(className)
		if className == "TForm" && (!field.IsValid() || !field.CanSet()) {
			field = vPtr.FieldByName("IForm")
		}
		setFiledVal(field, f)
		// fullComponent 只有当是从资源加载的才进行填充操作。
		if fullSubComponent {
			var ci int32
			for ci = 0; ci < f.ComponentCount(); ci++ {
				// 通过资源文件中的组件名字查找
				setFiledVal(vPtr.FieldByName(f.Components(ci).Name()), f.Components(ci))
			}
		}
		// 关联事件
		autoBindEvents(goInstance, f, fullSubComponent, afterBindSubComponentsEvents)
	}
}

// 共用的一个从资源中加载构建对象
//func resObjectBuild(typ int, owner IComponent, appInst uintptr, fields ...interface{}) IComponent {
//	if !DEBUG {
//		defer func() {
//			if err := recover(); err != nil {
//				fmt.Println("resCreateForm Error: ", err)
//			}
//		}()
//	}
//	runtime.LockOSThread()
//	defer runtime.UnlockOSThread()
//
//	var fullSubComponent bool
//	var afterBindSubComponentsEvents bool
//	var field1 interface{}
//	var goInstance reflect.Value
//
//	// 检测是否为MainForm，通过判断 指定方法为nil。
//	mainForm := Application.MainForm()
//	isMainForm := mainForm == nil || mainForm.Instance() == 0
//	instancePtr := uintptr(0)
//	// 不检查一些了，也不做最初版本的兼容
//	if len(fields) > 0 {
//		goInstance = newGoFormInstance(fields[0])
//		if !isMainForm {
//			instancePtr = goInstance.Pointer()
//		}
//		// 固定名称  CreateParams
//		cFunc := goInstance.MethodByName("CreateParams")
//		if cFunc.IsValid() {
//			//addToRequestCreateParamsMap(instancePtr, cFunc)
//		}
//	}
//
//	var resObj IComponent
//
//	switch typ {
//	case 0:
//		// 由参数的个数决定，创建窗口时是否使用缩放，此值需要 lcl.Application.SetScaled(true) 后才能生效。
//		resObj = AsForm(Application_CreateForm(appInst))
//	case 1:
//		resObj = AsForm(Form_Create2(CheckPtr(owner)))
//	case 2:
//		// TFrame
//		resObj = NewFrame(owner)
//	}
//
//	// 不为TFrame和MainForm时才设置这个
//	if typ != 2 && !isMainForm {
//		// 设置条件
//		Form_SetGoPtr(resObj.Instance(), instancePtr)
//	}
//
//	//条件设置用
//	bindSubs := func(sub, afterSub bool) {
//		fullSubComponent = sub
//		afterBindSubComponentsEvents = afterSub
//	}
//
//	// 查找并构建Form
//	findAndBuildForm := func(field interface{}) error {
//		res, err := findFormResource(field)
//		// 找到了对应的Form资源
//		if err == nil {
//			bindSubs(true, false)
//			loadFormResourceStream(*res.Data, resObj)
//		}
//		return err
//	}
//
//	switch len(fields) {
//	case 1:
//		field1 = fields[0]
//		bindSubs(false, false)
//		// 查找并构建Form
//		if findAndBuildForm(field1) != nil {
//			// 没有找到对应的资源，估计是手动创建的，将这个永远设置为true
//			bindSubs(false, true)
//		}
//	case 2:
//		switch fields[1].(type) {
//		// 当第二个参数为bool时，表示不填充子组件，为true表示之后绑定事件
//		case bool:
//			field1 = fields[0]
//			bindSubs(false, false) //fields[1].(bool)
//			// 查找并构建Form
//			if findAndBuildForm(field1) != nil {
//				// 没有找到对应的资源，估计是手动创建的
//				// 如果指定为false则不绑定
//				bindSubs(false, fields[1].(bool))
//			}
//		default:
//			// 第二个参数类型不为bool时，填充子组件为true，之后绑定事件为false
//			field1 = fields[1]
//			bindSubs(true, false)
//			switch fields[0].(type) {
//			case string:
//				ResFormLoadFromFile(fields[0].(string), CheckPtr(resObj))
//			case []byte:
//				loadFormResourceStream(fields[0].([]byte), resObj)
//			}
//		}
//	default:
//		return resObj
//	}
//	fullFiledVal(resObj, goInstance, fullSubComponent, afterBindSubComponentsEvents)
//	return nil
//}

// 从Stream中加载Form资源
//func loadFormResourceStream(data []byte, obj IComponent) {
//	mem := NewMemoryStream()
//	defer mem.Free()
//	mem.Write(data)
//	ResFormLoadFromStream(CheckPtr(mem), CheckPtr(obj))
//}
