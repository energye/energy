package main

import (
	"github.com/energye/energy/v2/cef"
)

// 该示例采用tempdll方式
//  该方式不再需要手动或自动安装liblcl.xx
//  需要通过编译命令控制该方式的开启
//  编译命令 -tags="tempdll"
//   windows: go build -ldflags "-H windowsgui -s -w" -tags="tempdll"
//   linux: go build -ldflags "-s -w" -tags="tempdll"
//   macos: go build -ldflags "-s -w" -tags="tempdll"
//   可通过 upx 工具压缩编译好的执行文件

func main() {
	// SetDllSaveDirType 设置 liblcl 保存目录，默认系统临时目录
	//tempdll.TempDLL.SetDllSaveDirType(tempdll.TddEnergyHome)
	// 设置保存目录 DllSaveDirType = TddCustom 时生效
	//tempdll.TempDLL.SetDllSaveDir("/save/to/path/liblcl.xx")

	// 使用编译命令: go build -tags="tempdll"
	// 开发工具 Goland: 在运行配置中 Go tool arguments 中配置 -tags="tempdll"

	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	app := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	//运行应用
	cef.Run(app)
}
