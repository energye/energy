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
	"io/fs"
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
	buildOutDir := buildOutPath(projectData)
	if !tools.IsExist(buildOutDir) {
		if err := os.MkdirAll(buildOutDir, 0755); err != nil {
			return fmt.Errorf("unable to create directory: %w", err)
		}
	}
	// 生成安装生成配置文件 nsis.nsi
	if nsisData, err := readFile(projectData, windowsNsis); err != nil {
		return err
	} else {
		if err = writeFile(projectData, windowsNsis, nsisData); err != nil {
			return err
		}
	}
	// tools.nsh
	if toolsData, err := readFile(projectData, windowsNsisTools); err != nil {
		return err
	} else {
		data := make(map[string]any)
		data["Name"] = projectData.Name
		data["ProjectPath"] = filepath.FromSlash(projectData.ProjectPath)
		data["FrameworkPath"] = filepath.FromSlash(projectData.FrameworkPath)
		data["Info"] = projectData.Info.FromSlash()
		if content, err := tools.RenderTemplate(string(toolsData), data); err != nil {
			return err
		} else if err = writeFile(projectData, windowsNsisTools, content); err != nil {
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
	nsisScriptPath := filepath.Join(buildOutPath(projectData), windowsNsis)

	//var binary string
	//if consts.IsWindows {
	//	binary = filepath.Join(projectData.ProjectPath, projectData.Name+".exe")
	//} else {
	//	binary = filepath.Join(projectData.ProjectPath, projectData.Name)
	//}

	//args = append(args, "-DARG_ENERGY_BINARY="+binary)
	//if projectData.Info.InstallPack.License != "" {
	//	// 授权信息文本目录: ..\LICENSE.txt
	//	args = append(args, "-DARG_ENERGY_PAGE_LICENSE="+projectData.Info.License)
	//}
	//if projectData.Info.Language != "" {
	//	// default English
	//	// 可选多种语言: SimpChinese, 参考目录: NSIS\Contrib\Language files
	//	args = append(args, "-DARG_ENERGY_LANGUAGE="+projectData.Info.Language)
	//}
	////框架目录
	//args = append(args, "-DARG_ENERGY_CEF_FRAMEWORK="+projectData.FrameworkPath)
	args = append(args, nsisScriptPath)
	cmd.Command("makensis", args...)

	return nil
}

// 返回配置资源目录
func assetsPath(projectData *project.Project, file string) string {
	return filepath.ToSlash(filepath.Join(projectData.AssetsDir, file))
}

// 返回固定构建输出目录 $current/build
func buildOutPath(projectData *project.Project) string {
	return filepath.Join(projectData.ProjectPath, "build")
}

// ReadFile
//  读取文件，根据项目配置先在本地目录读取，如果读取失败，则在内置资源目录读取
func readFile(projectData *project.Project, file string) ([]byte, error) {
	localFilePath := assetsPath(projectData, file)
	content, err := os.ReadFile(localFilePath)
	if errors.Is(err, fs.ErrNotExist) {
		content, err = assets.ReadFile(assetsFSPath + file)
		if err != nil {
			return nil, err
		}
		return content, nil
	}

	return content, err
}

// 写文件
func writeFile(projectData *project.Project, file string, content []byte) error {
	buildOutDir := buildOutPath(projectData)
	if !tools.IsExist(buildOutDir) {
		if err := os.MkdirAll(buildOutDir, 0755); err != nil {
			return fmt.Errorf("unable to create directory: %w", err)
		}
	}
	targetPath := filepath.Join(buildOutDir, file)
	if !projectData.Clean {
		if tools.IsExist(targetPath) {
			return nil
		}
	}
	if err := os.WriteFile(targetPath, content, 0644); err != nil {
		return err
	}
	return nil
}
