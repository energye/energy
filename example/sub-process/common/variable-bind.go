package common

import (
	"fmt"
	"github.com/energye/energy/cef"
)

func VariableBind() {
	/*
		通用类型变量和对象类型变量创建的回调函数
		该回调函数-在主进程和渲染进程创建时调用
	*/
	//ar integer cef.JSValue
	cef.VariableBind.VariableCreateCallback(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, bind cef.IProvisionalBindStorage) {
		fmt.Println("GO变量和函数绑定回调 当前进程类型-ProcessType:", cef.Args.ProcessType())
		bind.NewString("stringv", "我是渲染进程里的值")
		bind.NewInteger("integerv", 1211111)
		bind.NewDouble("doublev", 11.0505)
		bind.NewBoolean("booleanv", true)
		bind.NewNull("nullv")
		bind.NewUndefined("undefinedv")
	})
}
