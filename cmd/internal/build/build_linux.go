//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build linux
// +build linux

package build

import (
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/project"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	toolsCommand "github.com/energye/golcl/tools/command"
	"strings"
)

func build(c *command.Config) error {
	// 读取项目配置文件 energy.json 在main函数目录
	if proj, err := project.NewProject(c.Build.Path); err != nil {
		return err
	} else {
		// go build
		cmd := toolsCommand.NewCMD()
		cmd.Dir = proj.ProjectPath
		cmd.IsPrint = false
		term.Section.Println("Building", proj.OutputFilename)
		var args = []string{"build", "-ldflags", "-s -w", "-o", proj.OutputFilename}
		cmd.Command("go", args...)
		cmd.Command("strip", proj.OutputFilename)
		// upx
		if c.Build.Upx && tools.CommandExists("upx") {
			term.Section.Println("Upx compression")
			args = []string{"--best", "--no-color", "--no-progress", proj.OutputFilename}
			if c.Build.UpxFlag != "" {
				args = strings.Split(c.Build.UpxFlag, " ")
				args = append(args, proj.OutputFilename)
			}
			cmd.Command("upx", args...)
		}
		cmd.Close()
		if err == nil {
			term.Section.Println("Build Successfully")
		}
	}
	return nil
}
