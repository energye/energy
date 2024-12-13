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
	toolsCommand "github.com/energye/energy/v2/cmd/internal/tools/cmd"
	"os"
	"strings"
)

func build(c *command.Config, proj *project.Project) (err error) {
	// go build
	cmd := toolsCommand.NewCMD()
	cmd.Dir = proj.ProjectPath
	cmd.IsPrint = false
	outputFilename := proj.OutputFilename
	if c.Build.Out != "" {
		outputFilename = c.Build.Out
	}
	term.Section.Println("Building", outputFilename)
	var args = []string{"build"}
	if c.Build.Args != "" {
		gbargs := strings.Split(c.Build.Args, " ")
		for i := range gbargs {
			args = append(args, gbargs[i])
		}
	}
	args = append(args, "-ldflags", "-s -w")
	args = append(args, "-o", outputFilename)
	// GOOS=windows GOARCH=386
	if c.Build.OS != "" {
		os.Setenv("GOOS", c.Build.OS)
	}
	if c.Build.ARCH != "" {
		os.Setenv("GOARCH", c.Build.ARCH)
	}
	cmd.Command("go", args...)
	cmd.Command("strip", outputFilename)
	// upx
	if c.Build.Upx && tools.CommandExists("upx") {
		term.Section.Println("Upx compression")
		args = []string{"--best", "--no-color", "--no-progress", outputFilename}
		if c.Build.UpxFlag != "" {
			args = strings.Split(c.Build.UpxFlag, " ")
			args = append(args, outputFilename)
		}
		cmd.Command("upx", args...)
	}
	cmd.Close()
	if err == nil {
		term.Section.Println("Build Successfully")
	}
	return nil
}
