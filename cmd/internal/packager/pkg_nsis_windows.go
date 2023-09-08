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

package packager

import (
	"errors"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/assets"
	"github.com/energye/energy/v2/cmd/internal/project"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/energye/golcl/tools/command"
	"os"
	"path/filepath"
)

const (
	windowsNsis      = "windows/installer-nsis.nsi"
	windowsNsisTools = "windows/installer-tools.nsh"
)

func GeneraInstaller(projectData *project.Project) error {
	if !tools.CommandExists("makensis") {
		return errors.New("failed to create application installation program. Could not find the makensis command")
	}
	if err := windows(projectData); err != nil {
		return err
	}
	if err := makeNSIS(projectData); err != nil {
		return err
	}
	return nil
}

func windows(projectData *project.Project) error {
	// 创建构建输出目录
	buildOutDir := assets.BuildOutPath(projectData)
	buildOutDir = filepath.Join(buildOutDir, "windows")
	if !tools.IsExist(buildOutDir) {
		if err := os.MkdirAll(buildOutDir, 0755); err != nil {
			return fmt.Errorf("unable to create directory: %w", err)
		}
	}
	// 生成安装生成配置文件 nsis.nsi
	if nsisData, err := assets.ReadFile(projectData, assetsFSPath, windowsNsis); err != nil {
		return err
	} else {
		if err = assets.WriteFile(projectData, windowsNsis, nsisData); err != nil {
			return err
		}
	}
	// tools.nsh
	if toolsData, err := assets.ReadFile(projectData, assetsFSPath, windowsNsisTools); err != nil {
		return err
	} else {
		data := make(map[string]any)
		data["Name"] = projectData.Name
		data["ProjectPath"] = filepath.FromSlash(projectData.ProjectPath)
		data["FrameworkPath"] = filepath.FromSlash(projectData.FrameworkPath)
		data["Info"] = projectData.Info.FromSlash()
		if content, err := tools.RenderTemplate(string(toolsData), data); err != nil {
			return err
		} else if err = assets.WriteFile(projectData, windowsNsisTools, content); err != nil {
			return err
		}
	}
	return nil
}

// 使用nsis生成安装包
func makeNSIS(projectData *project.Project) error {
	var args []string
	cmd := command.NewCMD()
	cmd.IsPrint = false
	cmd.Dir = projectData.ProjectPath
	cmd.MessageCallback = func(bytes []byte, err error) {
		println("makensis:", string(bytes))
	}
	nsisScriptPath := filepath.Join(assets.BuildOutPath(projectData), windowsNsis)
	args = append(args, nsisScriptPath)
	cmd.Command("makensis", args...)
	return nil
}
