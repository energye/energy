package main

import (
	"fmt"
	"github.com/energye/energy/cef"
)

func main() {
	cef.UseNetIPCChannel = false
	cef.IPC.StartBrowserIPC()
	var i = 0
	cef.IPC.Browser().On("on-server", func(context cef.IIPCContext) {
		var message = context.Arguments()
		var data = message.GetString(0)
		fmt.Println("on-server:", data)
		context.Response([]byte("服务端回复:" + fmt.Sprintf("%d", i)))
		cef.IPC.Browser().EmitChannelId("on-client", 10, message)
		i++
	})
	select {}
}
