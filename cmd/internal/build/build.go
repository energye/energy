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
	"github.com/energye/energy/v2/cmd/internal/project"
)

const (
	assetsFSPath = "assets/build/"
)

func Build(c *command.Config) error {
	// 读取项目配置文件 energy.json 在main函数目录
	if proj, err := project.NewProject(c.Build.Path); err != nil {
		return err
	} else {
		return build(c, proj)
	}
}
