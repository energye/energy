package main

import (
	"fmt"
	"github.com/energye/golcl/energy/tools/winRes"
	"os"
	"path/filepath"
)

// 应用程序图标生成器
func main() {
	sysoRsrc()
}

// 使用 syso 和 rsrc 生成
func sysoRsrc() {
	err := syso() // 使用 syso
	if err != nil {
		fmt.Println("syso-error:", err.Error())
		//syso失败的时候使用rsrc方式生成，该方式需要rsrc.exe执行文件
		rsrc()
	}
}

func wd() string {
	d, _ := os.Getwd()
	return d
}

// 使用 rsrc 生成图标
//  基于 github.com/akavel/rsrc
func rsrc() error {
	// 创建 rsrc 对象
	rsrc := winRes.NewRSRC()
	// 设置 RSRC rvtf 执行文件 目录
	rsrc.RSRCPath = "/rsrc/exe/path"
	// 设置命令行执行目录, 该目录是要生成图标的应用根目录
	rsrc.CMDDir = filepath.Join(wd(), "example", "browser-control")
	return rsrc.Gen()
}

// 使用 syso 生成图标
//  基于 windres
func syso() error {
	// 创建 syso 对象
	syso := winRes.NewSYSO()
	// 设置命令行执行目录, 该目录是要生成图标的应用根目录
	// 资源在 resources 目录中读取
	syso.CMDDir = filepath.Join(wd(), "example", "browser-lib-checkupdate")
	syso.IconName = "icon.ico"
	return syso.RC()
}
