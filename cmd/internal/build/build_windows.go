//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

package build

import (
	"errors"
	"flag"
	"github.com/cyber-xxm/energy/v2/cmd/internal/assets"
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/env"
	"github.com/cyber-xxm/energy/v2/cmd/internal/gen"
	"github.com/cyber-xxm/energy/v2/cmd/internal/project"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	toolsCommand "github.com/cyber-xxm/energy/v2/cmd/internal/tools/cmd"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// 构建windows执行程序
//
//	exe生成图标
//	编译go
//	upx
func build(c *command.Config, proj *project.Project) (err error) {
	var (
		iconPath string
	)
	if iconPath, err = generaICON(proj); err != nil {
		return err
	}
	if _, err = generaSYSO(iconPath, c, proj); err != nil {
		return err
	}
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
			args = append(args, "-ldflags", "-s -w -H windowsgui "+*ldflags)
		} else {
			args = append(args, "-ldflags", "-s -w -H windowsgui")
		}
	} else {
		// 默认构建参数
		args = append(args, "-tags", "prod")
		args = append(args, "-ldflags", "-s -w -H windowsgui")
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
	gocmd := env.GlobalDevEnvConfig.GoCMD()
	if gocmd != "" {
		cmd.Command(gocmd, args...)
	} else {
		return errors.New("no Go command found")
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
			cmd.Command(upxcmd, args...)
		} else {
			term.Logger.Error("No UPX command found")
		}
	}
	cmd.Close()
	if err == nil {
		term.Section.Println("Build Successfully")
	}

	return nil
}

// 生成syso图标
func generaSYSO(iconPath string, c *command.Config, proj *project.Project) (string, error) {
	arch := runtime.GOARCH
	if c.Build.ARCH != "" {
		arch = string(c.Build.ARCH)
	}
	return gen.GeneraSYSO(proj.Name, iconPath, proj.Info.Manifest, proj.ProjectPath, arch, proj.Info)
}

// 生成应用图标，如果配置的是png图标，把png转换ico
func generaICON(proj *project.Project) (string, error) {
	iconPath := proj.Info.Icon
	outPath := filepath.Join(assets.BuildOutPath(proj), "windows")
	return gen.GeneraICON(iconPath, outPath)
}
