package common

import (
	"net"
)

// Port 获取并返回未使用的net socket端口
func Port() int {
	//主进程获取端口号
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		panic("Failed to Get unused Port number Error: " + err.Error())
	}
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic("Failed to Get unused Port number Error: " + err.Error())
	}
	defer listen.Close()
	return listen.Addr().(*net.TCPAddr).Port
}
