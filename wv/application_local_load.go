//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package wv

import "github.com/energye/lcl/emfs"

// LocalLoad
//
//	本地&内置资源加载配置
//	然后使用 Build() 函数构建对象
type LocalLoad struct {
	Enable bool   // 设置是否启用本地资源缓存到内存, 默认true: 启用, 禁用时需要调用Disable函数
	Scheme string // 自定义协议, 不建议使用 HTTP、HTTPS、FILE、FTP、ABOUT和DATA 默认: fs
	Domain string // 自定义域, 格式: xxx | xxx.xx | xxx.xxx.xxx， example, example.com, 默认: energy
	// 资源根目录, fs为空时: 本地目录(默认当前程序执行目录), fs不为空时: 默认值 resources, 使用内置加载
	// 本地目录规则: 空("")时当前目录, @当前目录开始(@/to/path)，或绝对目录.
	ResRootDir string        //
	FS         emfs.IEmbedFS // 内置加载资源对象, 不为nil时使用内置加载，默认: nil
	Proxy      IXHRProxy     // 数据请求代理, 在浏览器发送xhr请求时可通过该配置转发, 你可自定义实现该 IXHRProxy 接口
	Home       string        // 默认首页HTML文件名: /index.html , 默认: /index.html
}
