//go:build !windows

package syso

const (
	Syso = false
)

var SysoBytesAMD64 []byte
var SysoBytes386 []byte
