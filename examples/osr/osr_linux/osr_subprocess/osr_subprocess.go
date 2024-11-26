package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/golcl/pkgs/libname"
)

// 必须构建子进程，并和主进程放置在同一目录
func main() {
	cef.GlobalInit(nil, nil)
	app := cef.CreateApplication()
	cef.SetApplication(app)
	app.SetWindowlessRenderingEnabled(true)
	app.SetExternalMessagePump(true)
	app.SetMultiThreadedMessageLoop(false)
	app.SetSetCurrentDir(true)
	fmt.Println("start sub", libname.LibName)
	startsub := app.StartSubProcess()
	fmt.Println("start sub", startsub)
}
