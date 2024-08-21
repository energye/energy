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

import (
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/tools/exec"
)

const (
	localProto  = "fs"          // 默认本地资源加载协议
	localDomain = "energy"      // 默认本地资源加载域
	localHome   = "/index.html" //
)

// LocalLoad
//
//	Respond directly by loading local or built-in resources
type LocalLoad struct {
	Scheme string // 自定义协议, 不建议使用 HTTP、HTTPS、FILE、FTP、ABOUT和DATA 默认: fs
	Domain string // 自定义域, 格式: xxx | xxx.xx | xxx.xxx.xxx， example, example.com, 默认: energy
	// 资源根目录, fs为空时: 本地目录(默认当前程序执行目录), fs不为空时: 默认值 resources, 使用内置加载
	// 本地目录规则: 空("")时当前目录, @当前目录开始(@/to/path)，或绝对目录.
	ResRootDir string        //
	FS         emfs.IEmbedFS // 内置加载资源对象, 不为nil时使用内置加载，默认: nil
	Proxy      IXHRProxy     // 数据请求代理, 在浏览器发送xhr请求时可通过该配置转发, 你可自定义实现该 IXHRProxy 接口
	Home       string        // 默认首页HTML文件名: /index.html , 默认: /index.html
	exePath    string
}

func (m *LocalLoad) defaultInit() {
	if m.Domain == "" {
		m.Domain = localDomain
	}
	if m.Scheme == "" {
		m.Scheme = localProto
	}
	if m.Home == "" {
		m.Home = localHome
	} else if m.Home[0] != '/' {
		m.Home = "/" + m.Home
	}
	m.exePath = exec.Path
	// 默认的资源目录
	if m.ResRootDir == "" {
		if m.FS != nil && m.ResRootDir == "" {
			// embed exe
			m.ResRootDir = "resources"
		} else {
			// local disk, exe current dir
			m.ResRootDir = exec.Path
		}
	}
	if m.Proxy != nil {
		if proxy, ok := m.Proxy.(*XHRProxy); ok {
			proxy.init()
		}
	}
}
