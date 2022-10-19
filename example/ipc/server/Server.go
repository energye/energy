package main

import (
	"fmt"
	"github.com/energye/energy/ipc"
)

func main() {
	ipc.UseNetIPCChannel = false
	ipc.IPC.StartBrowserIPC()
	var i = 0
	ipc.IPC.Browser().On("on-server", func(context ipc.IIPCContext) {
		var message = context.Arguments()
		var data = message.GetString(0)
		fmt.Println("on-server:", data)
		context.Response([]byte("服务端回复:" + fmt.Sprintf("%d", i)))
		ipc.IPC.Browser().EmitChannelId("on-client", 10, message)
		i++
	})
	select {}
}
