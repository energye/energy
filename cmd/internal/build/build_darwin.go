//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin
// +build darwin

package build

import (
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/env"
	"github.com/energye/energy/v2/cmd/internal/project"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	toolsCommand "github.com/energye/energy/v2/cmd/internal/tools/cmd"
	"os"
	"runtime"
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
		os.Setenv("GOOS", string(c.Build.OS))
	}
	if c.Build.ARCH != "" {
		os.Setenv("GOARCH", string(c.Build.ARCH))
	}
	if c.Build.OS != "" || c.Build.ARCH != "" {
		if !tools.Equals(c.Build.OS.Value(), runtime.GOOS) || !tools.Equals(c.Build.ARCH.Value(), runtime.GOARCH) {
			os.Setenv("CGO_ENABLED", "1")
		}
	}

	gocmd := env.GlobalDevEnvConfig.GoCMD()
	if gocmd != "" {
		cmd.Command(gocmd, args...)
	} else {
		term.Logger.Error("No Go command found")
	}
	if c.Build.OS.IsMacOS() {
		cmd.Command("strip", outputFilename)
	}
	// upx
	if c.Build.Upx {
		upxcmd := env.GlobalDevEnvConfig.UPXCMD()
		if upxcmd != "" {
			term.Section.Println("Upx compression")
			args = []string{"--best", "--no-color", "--no-progress", outputFilename}
			if c.Build.UpxFlag != "" {
				args = strings.Split(c.Build.UpxFlag, " ")
				args = append(args, outputFilename)
			}
			cmd.Command("upx", args...)
		} else {
			term.Logger.Error("upx command not found", term.Logger.Args("install-upx", "brew install upx"))
		}
	}

	cmd.Close()
	if err == nil {
		term.Section.Println("Build Successfully")
	}
	return nil
}
