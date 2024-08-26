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
	localProto  = "fs"
	localDomain = "energy"
	localHome   = "/index.html"
)

// LocalLoad
//
//	Respond directly by loading local or built-in resources
type LocalLoad struct {
	Scheme string // User-defined protocol default: fs
	Domain string // The format is xxx | xxx.xx | xxx.xxx.xxx, example, example.com. default: energy
	// When fs is empty: local directory (default current program execution directory), when fs is not empty: default resources, using built-in load
	// Local directory rules: current directory when empty (""), @start of current directory (@/to/path), or absolute directory.
	ResRootDir string        // Resource root directory
	FS         emfs.IEmbedFS // Built-in load resource object, use built-in load when not nil, default: nil
	Proxy      IXHRProxy     // Proxy, which can be forwarded through this configuration when the browser sends xhr requests, you can customize the implementation of the IXHRProxy interface
	Home       string        // Home resources, default: /index.html
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
	// default resource dir
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
