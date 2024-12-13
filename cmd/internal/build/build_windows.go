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
	"github.com/energye/energy/v2/cmd/internal/assets"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/gen"
	"github.com/energye/energy/v2/cmd/internal/project"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	toolsCommand "github.com/energye/energy/v2/cmd/internal/tools/cmd"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	windowManifest    = "windows/app.exe.manifest"
	windowVersionInfo = "windows/version.info.json"
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
	cmd.IsPrint = false
	outputFilename := proj.OutputFilename
	if c.Build.Out != "" {
		outputFilename = c.Build.Out
	}
	term.Section.Println("Building", outputFilename)
	var args = []string{"build"}
	if c.Build.Args != "" {
		// go build args
		gbargs := strings.Split(c.Build.Args, " ")
		for i := range gbargs {
			args = append(args, gbargs[i])
		}
	}
	args = append(args, "-ldflags", "-s -w -H windowsgui")
	args = append(args, "-o", outputFilename)
	// GOOS=windows GOARCH=386
	if c.Build.OS != "" {
		os.Setenv("GOOS", c.Build.OS)
	}
	if c.Build.ARCH != "" {
		os.Setenv("GOARCH", c.Build.ARCH)
	}
	cmd.Command("go", args...)
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

// 生成syso图标
func generaSYSO(iconPath string, c *command.Config, proj *project.Project) (string, error) {
	arch := runtime.GOARCH
	if c.Build.ARCH != "" {
		arch = c.Build.ARCH
	}
	return gen.GeneraSYSO(proj.Name, iconPath, proj.Info.Manifest, proj.ProjectPath, arch, proj.Info)
}

// 生成应用图标，如果配置的是png图标，把png转换ico
func generaICON(proj *project.Project) (string, error) {
	iconPath := proj.Info.Icon
	outPath := filepath.Join(assets.BuildOutPath(proj), "windows")
	return gen.GeneraICON(iconPath, outPath)
}
