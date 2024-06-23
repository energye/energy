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
	"github.com/energye/energy/v3/cmd/internal/command"
	toolsCommand "github.com/energye/energy/v3/cmd/internal/pkgs/command"
	"github.com/energye/energy/v3/cmd/internal/project"
	"github.com/energye/energy/v3/cmd/internal/term"
	"github.com/energye/energy/v3/cmd/internal/tools"
	"strings"
)

func build(c *command.Config, proj *project.Project) (err error) {
	// go build
	cmd := toolsCommand.NewCMD()
	cmd.Dir = proj.ProjectPath
	cmd.IsPrint = false
	term.Section.Println("Building", proj.OutputFilename)
	var args = []string{"build"}
	if c.Build.Args != "" {
		gbargs := strings.Split(c.Build.Args, " ")
		for i := range gbargs {
			args = append(args, gbargs[i])
		}
	}
	args = append(args, "-ldflags", "-s -w")
	args = append(args, "-o", proj.OutputFilename)
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
	return nil
}
