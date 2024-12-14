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
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/env"
	"github.com/energye/energy/v2/cmd/internal/project"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"io"
	"os"
	"path"
)

const (
	assetsFSPath = "assets/build/"
)

func Build(c *command.Config) error {
	// 读取项目配置文件 energy.json 在main函数目录
	if proj, err := project.NewProject(c.Build.Path); err != nil {
		return err
	} else {
		// libemfs 标志开启后,如果使用了内置libs到执行文件, 项目energy.json配置将ENERGY_HOME环境变量下的liblcl库复制到内置目录
		// 这里仅仅是在编译时把liblcl复制到内置资源目录中
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
			libsrc := path.Join(env.GlobalDevEnvConfig.Framework, tools.GetDLLName())
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
		return build(c, proj)
	}
}
