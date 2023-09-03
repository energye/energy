//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package packager

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/project"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/energye/golcl/tools/command"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"text/template"
)

//go:embed assets
var assets embed.FS

const (
	windowsNsis      = "windows/installer-nsis.nsi"
	windowsNsisTools = "windows/installer-tools.nsh"
)

func GeneraNSISInstaller(projectData *project.Project) error {
	switch runtime.GOOS {
	case "windows":
		if !tools.CommandExists("makensis") {
			return errors.New("failed to create application installation program. Could not find the makensis command")
		}
		if err := windows(projectData); err != nil {
			return err
		}
		if err := makeNSIS(projectData); err != nil {
			return err
		}
	case "linux":
	case "darwin":
	default:
		return errors.New("unsupported system")
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
	// 生成安装生成配置文件
	if nsisData, err := readFile(projectData, windowsNsis); err != nil {
		return err
	} else {
		if err = writeFile(projectData, windowsNsis, nsisData); err != nil {
			return err
		}
	}
	if toolsData, err := readFile(projectData, windowsNsisTools); err != nil {
		return err
	} else {
		tmpl, err := template.New("tools").Parse(string(toolsData))
		if err != nil {
			return err
		}
		data := make(map[string]any)
		data["Name"] = projectData.Name
		data["Info"] = projectData.Info
		var out bytes.Buffer
		if err = tmpl.Execute(&out, data); err != nil {
			return err
		}
		if err = writeFile(projectData, windowsNsisTools, out.Bytes()); err != nil {
			return err
		}
	}
	return nil
}

// 使用nsis生成安装包
func makeNSIS(projectData *project.Project) error {
	var args []string
	cmd := command.NewCMD()
	cmd.Dir = projectData.ProjectPath
	cmd.MessageCallback = func(bytes []byte, err error) {
		println("makensis:", string(bytes))
	}
	nsisScriptPath := filepath.Join(buildOutPath(projectData), windowsNsis)

	var binary string
	if runtime.GOOS == "windows" {
		binary = filepath.Join(projectData.ProjectPath, projectData.Name+".exe")
	} else {
		binary = filepath.Join(projectData.ProjectPath, projectData.Name)
	}

	args = append(args, "-DARG_ENERGY_BINARY="+binary)
	if projectData.Info.License != "" {
		// 授权信息文本目录: ..\LICENSE.txt
		args = append(args, "-DARG_ENERGY_PAGE_LICENSE="+projectData.Info.License)
	}
	if projectData.Info.Language != "" {
		// default English
		// 可选多种语言: SimpChinese, 参考目录: NSIS\Contrib\Language files
		args = append(args, "-DARG_ENERGY_LANGUAGE="+projectData.Info.Language)
	}
	//框架目录
	args = append(args, "-DARG_ENERGY_CEF_FRAMEWORK="+projectData.FrameworkPath)
	args = append(args, nsisScriptPath)
	cmd.Command("makensis", args...)

	return nil
}

// 返回根据配置的资源目录
func assetsPath(projectData *project.Project, file string) string {
	return filepath.ToSlash(filepath.Join(projectData.BuildAssetsDir, file))
}

// 返回固定的构建输出目录 $current/build
func buildOutPath(projectData *project.Project) string {
	return filepath.Join(projectData.ProjectPath, "build")
}

// ReadFile
//  读取文件，根据项目配置先在本地目录读取，如果读取失败，则在内置资源目录读取
func readFile(projectData *project.Project, file string) ([]byte, error) {
	localFilePath := assetsPath(projectData, file)
	content, err := os.ReadFile(localFilePath)
	if errors.Is(err, fs.ErrNotExist) {
		content, err = fs.ReadFile(assets, localFilePath)
		if err != nil {
			return nil, err
		}
		if err := writeFile(projectData, file, content); err != nil {
			return nil, fmt.Errorf("unable to create file in build folder: %s", err)
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
	if !tools.IsExist(filepath.Dir(targetPath)) {
		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			return fmt.Errorf("unable to create directory: %w", err)
		}
	}
	if err := os.WriteFile(targetPath, content, 0644); err != nil {
		return err
	}
	return nil
}
