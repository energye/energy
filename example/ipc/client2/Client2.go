package main

import (
	"fmt"
	"github.com/energye/energy/ipc"
)

func main() {
	ipc.UseNetIPCChannel = false
	ipc.IPC.CreateRenderIPC(1, 10)
	var i = 0
	ipc.IPC.Render().On("on-client", func(context ipc.IIPCContext) {
		var message = context.Arguments()
		var data = message.GetString(0)
		//context.Free()
		fmt.Println("on-client:", data)
		i++
	})

	for j := 0; j < 1; j++ {
		go func() {
			for {
				args := ipc.NewArgumentList()
				args.SetString(0, "数据:"+fmt.Sprintf("%d", i))
				ipc.IPC.Render().Emit("on-server", args)
				//cef.IPC.Render().EmitAndCallback("on-server", args, func(context cef.IIPCContext) {
				//	fmt.Println("客户端接收:", string(context.Message().Data()))
				//	//context.Free()
				//	i++
				//})
				args.Clear()
			}
		}()
	}
	select {}
}
