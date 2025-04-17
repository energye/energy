//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package build

import (
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/consts"
	"github.com/cyber-xxm/energy/v2/cmd/internal/env"
	"github.com/cyber-xxm/energy/v2/cmd/internal/project"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	"io"
	"os"
	"path"
)

const (
	assetsFSPath = "assets/build/"
)

func Build(c *command.Config) error {
	// 读取项目配置文件 config/energy_[os].json 在main函数目录
	if proj, err := project.NewProject(c.Build.Path); err != nil {
		return err
	} else {
		// libemfs 标志开启后,如果使用了内置libs到执行文件, 项目 config/energy_[os].json 配置 frameworkPath 量下的 liblcl 动态库复制到内置目录
		// 在编译时把 liblcl 复制到内置资源目录中
		if !consts.IsDarwin && c.Build.Libemfs {
			// 复制liblcl到内置目录
			emfsPath := path.Join(proj.ProjectPath, proj.LibEMFS)
			dllPath := path.Join(emfsPath, tools.GetDLLName())
			if tools.IsExist(dllPath) {
				os.Remove(dllPath)
			} else {
				os.MkdirAll(emfsPath, os.ModePerm)
			}
			// copy
			libsrc := path.Join(env.GlobalDevEnvConfig.FrameworkPath(), tools.GetDLLName())
			src, err := os.Open(libsrc)
			if err != nil {
				return err
			}
			defer src.Close()
			st, _ := src.Stat()
			dst, err := os.OpenFile(dllPath, os.O_RDWR|os.O_CREATE, st.Mode())
			if err != nil {
				return err
			}
			defer dst.Close()
			_, err = io.Copy(dst, src)
			if err != nil {
				return err
			}
		}
		// 默认值
		if c.Build.OS == "" {
			c.Build.OS = command.OS(env.GlobalDevEnvConfig.OS())
		}
		if c.Build.ARCH == "" {
			c.Build.ARCH = command.Arch(env.GlobalDevEnvConfig.Arch())
		}
		term.Section.Println("Build env GOOS:", c.Build.OS, "GOARCH:", c.Build.ARCH)
		return build(c, proj)
	}
}
