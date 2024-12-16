//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package env

import (
	"errors"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/pterm/pterm"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Env(c *command.Config) error {
	env := c.Env
	if env.Write == "" && env.Get == "" && !env.List && env.Use == "" {
		PrintENV()
	} else if env.List {
		if err := printInstalledFrameworks(); err != nil {
			return err
		}
	} else if env.Use != "" {
		if err := useInstalledFrameworks(env.Use); err != nil {
			return err
		}
	} else if env.Get != "" {
		switch strings.ToLower(env.Get) {
		case "golang", "go", "goroot":
			term.Section.Println(GlobalDevEnvConfig.GoRoot)
		case "framework":
			term.Section.Println(GlobalDevEnvConfig.FrameworkPath())
		case "nsis":
			term.Section.Println(GlobalDevEnvConfig.NSIS)
		case "7z":
			term.Section.Println(GlobalDevEnvConfig.Z7Z)
		case "upx":
			term.Section.Println(GlobalDevEnvConfig.UPX)
		case "proxy":
			term.Section.Println(GlobalDevEnvConfig.Proxy)
		case "registry":
			term.Section.Println(GlobalDevEnvConfig.Registry)
		}
	} else if env.Write != "" {
		keyval := strings.Split(env.Write, ":")
		if len(keyval) > 1 {
			key := strings.TrimSpace(keyval[0])
			val := strings.TrimSpace(strings.Join(keyval[1:], ":"))
			switch strings.ToLower(key) {
			case "golang", "go", "goroot":
				GlobalDevEnvConfig.GoRoot = val
			case "framework":
				GlobalDevEnvConfig.Framework = val
			case "root":
				GlobalDevEnvConfig.Root = val
			case "nsis":
				GlobalDevEnvConfig.NSIS = val
			case "7z":
				GlobalDevEnvConfig.Z7Z = val
			case "upx":
				GlobalDevEnvConfig.UPX = val
			case "proxy":
				GlobalDevEnvConfig.Proxy = val
			case "registry":
				GlobalDevEnvConfig.Registry = val
			}
			GlobalDevEnvConfig.Update()
		}
	}
	return nil
}

// 按完全匹配规则: CEF-[VER]_[OS]_[ARCH]
func getInstalledFrameworks() ([]string, error) {
	dirs, err := os.ReadDir(filepath.Join(GlobalDevEnvConfig.Root, consts.ENERGY))
	if err != nil {
		return nil, err
	}
	numberReg := regexp.MustCompile("^\\d+$")
	chkOS := func(os command.OS) bool {
		return os.IsWindows() || os.IsLinux() || os.IsMacOS()
	}
	chkARCH := func(arch command.Arch) bool {
		return arch.Is386() || arch.IsAMD64() || arch.IsARM() || arch.IsARM64()
	}
	var result []string
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		dirName := dir.Name()
		split := strings.Split(dirName, "-")
		if len(split) != 2 {
			continue
		}
		verosarch := strings.Split(split[1], "_")
		if len(verosarch) != 3 {
			continue
		}
		isName := tools.Equals(split[0], "CEF")
		isVER := numberReg.MatchString(verosarch[0])
		isOS := chkOS(command.OS(verosarch[1]))
		isARCH := chkARCH(command.Arch(verosarch[2]))
		if isName && isVER && isOS && isARCH {
			result = append(result, dirName)
		}
	}
	return result, nil
}

// 切换已安装的版本, 只做 CEF 版本号验证
func useInstalledFrameworks(ver string) error {
	dirs, err := getInstalledFrameworks()
	if err != nil {
		return err
	}
	var frameworkName string
	for _, dir := range dirs {
		split := strings.Split(dir, "-")
		if len(split) != 2 {
			continue
		}
		verosarch := strings.Split(split[1], "_")
		if len(verosarch) != 3 {
			continue
		}
		// 先只验证版本号，不验证系统和架构
		if ver == verosarch[0] {
			frameworkName = dir
			break
		}
	}
	if frameworkName != "" {
		GlobalDevEnvConfig.Framework = frameworkName
		GlobalDevEnvConfig.Update()
		msg := fmt.Sprintf("Now using CEF Framework %v", ver)
		term.Logger.Info(msg)
		return nil
	} else {
		err := fmt.Sprintf("Not Installed %v. Use CLI: [energy install --cef %v]", ver, ver)
		return errors.New(err)
	}
}

// 列出 目录内安装的 CEF Framework
func printInstalledFrameworks() error {
	dirs, err := getInstalledFrameworks()
	if err != nil {
		return err
	}
	tableData := pterm.TableData{
		{"  ENERGY CEF Framework"},
	}
	for _, dir := range dirs {
		if tools.Equals(dir, GlobalDevEnvConfig.Framework) {
			dir = "* " + dir
		} else {
			dir = "  " + dir
		}
		tableData = append(tableData, []string{dir})
	}
	pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("-").WithBoxed().WithData(tableData).Render()
	return nil
}

func CheckCEFDir() bool {
	var lib = func() string {
		if consts.IsWindows {
			return "libcef.dll"
		} else if consts.IsLinux {
			return "libcef.so"
		} else if consts.IsDarwin {
			return "cef_sandbox.a"
		}
		return ""
	}()
	if lib != "" {
		return tools.IsExist(filepath.Join(GlobalDevEnvConfig.FrameworkPath(), lib))
	}
	return false
}

func PrintENV() {
	tableData := pterm.TableData{
		{"Name", "Directory"},
	}
	tableData = append(tableData, []string{"Golang", GlobalDevEnvConfig.GoRoot})
	tableData = append(tableData, []string{"Root", GlobalDevEnvConfig.Root})
	tableData = append(tableData, []string{"Framework", GlobalDevEnvConfig.Framework})
	tableData = append(tableData, []string{"NSIS", GlobalDevEnvConfig.NSIS})
	tableData = append(tableData, []string{"7z", GlobalDevEnvConfig.Z7Z})
	tableData = append(tableData, []string{"UPX", GlobalDevEnvConfig.UPX})
	tableData = append(tableData, []string{"Registry", GlobalDevEnvConfig.Registry})
	tableData = append(tableData, []string{"Proxy", GlobalDevEnvConfig.Proxy})
	pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("-").WithBoxed().WithData(tableData).Render()
}
