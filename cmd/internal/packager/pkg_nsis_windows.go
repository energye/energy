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
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/energye/golcl/tools/command"
	"os"
	"path/filepath"
)

const (
	windowsNsis      = "windows/installer-nsis.nsi"
	windowsNsisTools = "windows/installer-tools.nsh"
)

func GeneraInstaller(proj *project.Project) error {
	if !tools.CommandExists("makensis") {
		return errors.New("failed to create application installation program. Could not find the makensis command")
	}
	var err error
	// 创建构建输出目录
	buildOutDir := assets.BuildOutPath(proj)
	buildOutDir = filepath.Join(buildOutDir, "windows")
	if !tools.IsExist(buildOutDir) {
		if err := os.MkdirAll(buildOutDir, 0755); err != nil {
			return fmt.Errorf("unable to create directory: %w", err)
		}
	}
	// 7za 压缩 CEF
	comper := proj.NSIS.Compress
	switch comper {
	case "7z", "7za":
		proj.NSIS.UseCompress = tools.CommandExists(comper)
	}
	if proj.NSIS.UseCompress {
		if cef7zFile, err := compressCEF7za(proj); err != nil {
			return err
		} else {
			proj.NSIS.CompressFile = cef7zFile
		}
	}

	// 生成 nsis 脚本
	if err = windows(proj); err != nil {
		return err
	}

	// make
	var outInstall string
	if outInstall, err = makeNSIS(proj); err != nil {
		return err
	}
	term.Section.Println("Success \n\tInstall package:", outInstall)
	return nil
}

func compressCEF7za(proj *project.Project) (string, error) {
	term.Logger.Info("7za compress " + proj.NSIS.CompressName + ", This may take some time")
	buildWindowsPath := filepath.Join(assets.BuildOutPath(proj), "windows")
	outFilePath := filepath.FromSlash(filepath.Join(buildWindowsPath, proj.NSIS.CompressName))
	if proj.Clean {
		os.Remove(outFilePath)
	} else if tools.IsExist(outFilePath) {
		term.Logger.Info(proj.NSIS.CompressName + " file exist")
		return outFilePath, nil
	}

	wd := tools.CurrentExecuteDir()
	defer func() {
		os.Chdir(wd)
	}()
	err := os.Chdir(proj.FrameworkPath)
	if err != nil {
		return "", err
	}

	cmd := command.NewCMD()
	cmd.IsPrint = false
	var args = []string{"a", outFilePath, filepath.FromSlash(fmt.Sprintf("%s/*", proj.FrameworkPath))}
	for _, exc := range proj.NSIS.Exclude {
		args = append(args, "-xr!"+exc)
	}
	cmd.Command(proj.NSIS.Compress, args...)
	cmd.Close()
	return outFilePath, nil
}

func windows(proj *project.Project) error {
	term.Logger.Info("Generate NSIS script")
	// 生成安装生成配置文件 nsis.nsi
	if nsisData, err := assets.ReadFile(proj, assetsFSPath, windowsNsis); err != nil {
		return err
	} else {
		if err = assets.WriteFile(proj, windowsNsis, nsisData); err != nil {
			return err
		}
	}
	// tools.nsh
	if toolsData, err := assets.ReadFile(proj, assetsFSPath, windowsNsisTools); err != nil {
		return err
	} else {
		data := make(map[string]any)
		data["Name"] = proj.Name
		data["ProjectPath"] = filepath.FromSlash(proj.ProjectPath)
		data["FrameworkPath"] = filepath.FromSlash(proj.FrameworkPath)
		proj.Info.FromSlash()
		proj.NSIS.FromSlash()
		data["Info"] = proj.Info
		data["NSIS"] = proj.NSIS
		if content, err := tools.RenderTemplate(string(toolsData), data); err != nil {
			return err
		} else if err = assets.WriteFile(proj, windowsNsisTools, content); err != nil {
			return err
		}
	}
	return nil
}

// 使用nsis生成安装包
func makeNSIS(proj *project.Project) (string, error) {
	installPackage := proj.Name + "-installer.exe"
	term.Logger.Info("NSIS Making Installation, Almost complete", term.Logger.Args("Install Package", installPackage))
	var args []string
	cmd := command.NewCMD()
	cmd.IsPrint = false
	cmd.Dir = proj.ProjectPath
	cmd.MessageCallback = func(bytes []byte, err error) {
		println("makensis:", string(bytes))
	}
	nsisScriptPath := filepath.Join(assets.BuildOutPath(proj), windowsNsis)
	args = append(args, nsisScriptPath)
	cmd.Command("makensis", args...)
	outInstall := filepath.Join(filepath.Dir(nsisScriptPath), installPackage)
	return outInstall, nil
}
