package wv

import (
	"github.com/energye/lcl/emfs"
	"github.com/energye/wv/wv"
)

// Init 全局初始化, 需手动调用的函数
//
//	参数:
//	   libs 内置到应用程序的类库
//	   resources 内置到应用程序的资源文件
func Init(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	wv.Init(libs, resources)
}
