package src

import (
	"fmt"
	"github.com/energye/energy/cef"
)

func IPCInit() {
	//渲染进程 IPC事件
	cef.IPC.Render().SetOnEvent(func(event cef.IEventOn) {
		fmt.Println("渲染进程IPC事件注册")
		//渲染进程监听的事件
		event.On("sub-process-on-event", func(context cef.IIPCContext) {
			fmt.Println("sub-process-on-event")
			//渲染进程处理程序....
			context.Response([]byte("返回结果"))
		})
	})
}
