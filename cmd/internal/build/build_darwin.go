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
	"flag"
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/env"
	"github.com/cyber-xxm/energy/v2/cmd/internal/project"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	toolsCommand "github.com/cyber-xxm/energy/v2/cmd/internal/tools/cmd"
	"os"
	"runtime"
	"strings"
)

func build(c *command.Config, proj *project.Project) (err error) {
	// go build
	cmd := toolsCommand.NewCMD()
	cmd.Dir = proj.ProjectPath
	outputFilename := proj.OutputFilename
	if c.Build.Out != "" {
		outputFilename = c.Build.Out
	}
	term.Section.Println("Building", outputFilename)
	args := []string{"build"}
	if c.Build.BuildArgs {
		// go build args
		// 在 energy build 时，如果设置 go 的构建参数, 需要设置 --buildargs 标记，并且让其在 cli 命令最一个有效参数位置
		// 其之后参数都将做为 go build [args] 传递
		buildArgs := os.Args[tools.GetBuildArgsFlagIndex():]
		cmdLine := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		tags := cmdLine.String("tags", "", "")
		ldflags := cmdLine.String("ldflags", "", "")
		cmdLine.Parse(buildArgs)
		if tags != nil && *tags != "" {
			args = append(args, "-tags", "prod,"+*tags)
		} else {
			args = append(args, "-tags", "prod")
		}
		if ldflags != nil && *ldflags != "" {
			args = append(args, "-ldflags", "-s -w"+*ldflags)
		} else {
			args = append(args, "-ldflags", "-s -w")
		}
	} else {
		// 默认构建参数
		args = append(args, "-tags", "prod")
		args = append(args, "-ldflags", "-s -w")
	}
	args = append(args, "-trimpath")
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
