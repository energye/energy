package main

import (
	"fmt"
	"github.com/energye/energy/cef"
)

func main() {
	cef.UseNetIPCChannel = false
	cef.IPC.CreateRenderIPC(1, 11)
	var i = 0
	cef.IPC.Render().On("on-client", func(context cef.IIPCContext) {
		var message = context.Arguments()
		var data = message.GetString(0)
		//context.Free()
		fmt.Println("on-client:", data)
		i++
	})

	for j := 0; j < 1; j++ {
		go func() {
			for {
				args := cef.NewArgumentList()
				args.SetString(0, "数据:"+fmt.Sprintf("%d", i))
				cef.IPC.Render().Emit("on-server", args)
				cef.IPC.Render().EmitAndCallback("on-server", args, func(context cef.IIPCContext) {
					fmt.Println("客户端接收:", string(context.Message().Data()))
					//context.Free()
					i++
				})
			}
		}()
	}
	select {}
}
