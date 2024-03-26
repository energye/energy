package main

import (
	"github.com/energye/energy/v2/examples/lcl/basicResForm/form"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
)

// windows : stdcall;  其他平台： cdecl
//var runLoopReceivedProc = syscall.NewCallback(runLoop)

// data 为 Application 实例
/*func runLoop(data uintptr) uintptr {
	// 这里不能产生异常，否则程序会崩溃

	return 0
	//lcl.Application.HandleMessage()
	//return 1 // 返回1表示不由Lazarus处理后面的，必须要加上Application.HandleMessage()，一般情况下最好不自己处理
}*/

func main() {
	//lcl.DEBUG = true
	//lcl.Application.SetRunLoopReceived(runLoopReceivedProc)
	inits.Init(nil, nil)
	lcl.RunApp(&form.Form1, &form.Form2)
}
