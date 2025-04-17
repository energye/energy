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
	"github.com/cyber-xxm/energy/v2/cmd/internal/assets"
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/env"
	"github.com/cyber-xxm/energy/v2/cmd/internal/project"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	cmd "github.com/cyber-xxm/energy/v2/cmd/internal/tools/cmd"
	"os"
	"path/filepath"
	"strings"
)

const (
	windowsNsis      = "windows/installer-nsis.nsi"
	windowsNsisTools = "windows/installer-tools.nsh"
)

func GeneraInstaller(c *command.Config, proj *project.Project) error {
	if env.GlobalDevEnvConfig.NSISCMD() == "" {
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
		proj.NSIS.UseCompress = env.GlobalDevEnvConfig.Z7ZCMD() != ""
	}
	if proj.NSIS.UseCompress {
		if cef7zFile, err := compressCEF7za(proj); err != nil {
			return err
		} else {
			proj.NSIS.CompressFile = cef7zFile
		}
	}

	// 生成 nsis 脚本
	if err = windows(c, proj); err != nil {
		return err
	}

	// make
	var outInstall string
	if outInstall, err = makeNSIS(c, proj); err != nil {
		return err
	}
	term.Section.Println("Success \n\tInstall package:", outInstall)
	return nil
}

func compressCEF7za(proj *project.Project) (string, error) {
	term.Logger.Info("7z compress " + proj.NSIS.CompressName + ", This may take some time")
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

	cmd := cmd.NewCMD()
	cmd.IsPrint = false
	var args = []string{"a", outFilePath, filepath.FromSlash(fmt.Sprintf("%s/*", proj.FrameworkPath))}
	for _, exc := range proj.NSIS.Exclude {
		args = append(args, "-xr!"+exc)
	}
	cmd.Command(env.GlobalDevEnvConfig.Z7ZCMD(), args...)
	cmd.Close()
	return outFilePath, nil
}

func windows(c *command.Config, proj *project.Project) error {
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
		exeName := proj.Name
		if c.Package.File != "" {
			exeName = c.Package.File
			if strings.LastIndex(exeName, ".") != -1 {
				exeName = exeName[:strings.LastIndex(exeName, ".")]
			}
		}
		installerFileName := installFileName(c, proj)

		data := make(map[string]interface{})
		data["Name"] = proj.Name
		data["ExeName"] = exeName
		data["InstallerFileName"] = installerFileName
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

func installFileName(c *command.Config, proj *project.Project) string {
	installPackage := proj.OutputFilename
	if c.Package.OutFileName != "" {
		installPackage = c.Package.OutFileName
		if !strings.HasSuffix(installPackage, ".exe") {
			installPackage += ".exe"
		}
	}
	return installPackage
}

// 使用nsis生成安装包
func makeNSIS(c *command.Config, proj *project.Project) (string, error) {
	installPackage := installFileName(c, proj)
	term.Logger.Info("NSIS Making Installation, Almost complete", term.Logger.Args("Install Package", installPackage))
	var args []string
	cmd := cmd.NewCMD()
	cmd.IsPrint = false
	cmd.Dir = proj.ProjectPath
	cmd.MessageCallback = func(bytes []byte, err error) {
		println("makensis:", string(bytes))
	}
	nsisScriptPath := filepath.Join(assets.BuildOutPath(proj), windowsNsis)
	args = append(args, nsisScriptPath)
	cmd.Command(env.GlobalDevEnvConfig.NSISCMD(), args...)
	outInstall := filepath.Join(filepath.Dir(nsisScriptPath), installPackage)
	return outInstall, nil
}
