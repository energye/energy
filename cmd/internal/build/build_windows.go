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
	toolsCommand "github.com/energye/golcl/tools/command"
	"path/filepath"
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
	if _, err = generaSYSO(iconPath, proj); err != nil {
		return err
	}
	// go build
	cmd := toolsCommand.NewCMD()
	cmd.Dir = proj.ProjectPath
	cmd.IsPrint = false
	term.Section.Println("Building", proj.OutputFilename)
	var args = []string{"build"}
	if c.Build.Args != "" {
		// go build args
		gbargs := strings.Split(c.Build.Args, " ")
		for i := range gbargs {
			args = append(args, gbargs[i])
		}
	}
	args = append(args, "-ldflags", "-s -w -H windowsgui")
	args = append(args, "-o", proj.OutputFilename)
	cmd.Command("go", args...)
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

// 生成syso图标
func generaSYSO(iconPath string, proj *project.Project) (string, error) {
	return gen.GeneraSYSO(proj.Name, iconPath, proj.Info.Manifest, proj.ProjectPath, "", proj.Info)
}

// 生成应用图标，如果配置的是png图标，把png转换ico
func generaICON(proj *project.Project) (string, error) {
	iconPath := proj.Info.Icon
	outPath := filepath.Join(assets.BuildOutPath(proj), "windows")
	return gen.GeneraICON(iconPath, outPath)
}
