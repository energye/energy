package main

import (
	"embed"
	"github.com/cyber-xxm/energy/v2/cef"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
)

/*
该示例采用tempdll方式
   该方式在Go编译时将执行文件内置到exe中
   前提条件
		1. 在Go main函数初始化全局配置时[cef.GlobalInit(libs, resources)]设置libs内置对象参数
		2. 内置资源根目录名默认libs, 也可通过 tempdll.TempDLL.SetDllFSDir("assets/libs") 自定义
*/

/*
有2种内置文件对象
	1. Go版本大于1.16可直接使用 embed.FS
	2. Go版本小于1.16不支持embed.FS, 使用energy提供的bindata方式内置进去
*/

/*
编译
energy build --libemfs 自动每次编译时都复制ENERGY_HOME目录的liblcl到内置目录
go build 手动复制liblcl到内置目录
*/

//go:embed libs
var libs embed.FS

func main() {
	// SetDllSaveDirType 设置 liblcl 保存目录，默认系统临时目录
	//tempdll.TempDLL.SetDllSaveDirType(tempdll.TddCurrent)
	// 设置保存目录 DllSaveDirType = TddCustom 时生效
	// tempdll.TempDLL.SetDllSaveDir("/save/to/path/liblcl.dll")
	// 设置liblcl所在FS目录
	// tempdll.TempDLL.SetDllFSDir("assets/libs")

	//全局初始化 每个应用都必须调用的
	// 在此处设置 libs 文件内嵌对象参数
	cef.GlobalInit(libs, nil)
	//创建应用
	app := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	//运行应用
	cef.Run(app)
}
