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
		rsrc()
	}
}
func rsrc() error {
	rsrc := winRes.NewRSRC()
	rsrc.RSRCPath = "E:\\SWT\\gopath\\src\\gitee.com\\snxamdf\\golcl\\exe"
	rsrc.CMDDir = "E:\\SWT\\gopath\\src\\swt-lazarus\\demo17-dll-load\\demo-golang-dll-01-chromium"
	return rsrc.Gen()
}
func syso() error {
	syso := winRes.NewSYSO()
	syso.CMDDir = "E:\\SWT\\gopath\\src\\swt-lazarus\\demo17-dll-load\\demo-golang-dll-01-chromium"
	syso.IconName = "icon.ico"
	return syso.RC()
}
