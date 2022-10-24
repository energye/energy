package main

import (
	"fmt"
	"github.com/energye/golcl/tools/winRes"
)

func main() {
	//rsrc()
	sysoRsrc()
}
func sysoRsrc() {
	err := syso()
	if err != nil {
		fmt.Println("syso-error:", err.Error())
		//syso失败的时候使用rsrc方式生成，该方式需要rsrc.exe执行文件
		rsrc()
	}
}
func rsrc() error {
	rsrc := winRes.NewRSRC()
	rsrc.RSRCPath = "E:\\SWT\\gopath\\src\\gitee.com\\snxamdf\\golcl\\exe"
	rsrc.CMDDir = "E:\\SWT\\gopath\\src\\github.com\\energye\\energy\\example\\browser-control"
	return rsrc.Gen()
}
func syso() error {
	syso := winRes.NewSYSO()
	syso.CMDDir = "E:\\SWT\\gopath\\src\\github.com\\energye\\energy\\example\\browser-control"
	syso.IconName = "icon.ico"
	return syso.RC()
}
