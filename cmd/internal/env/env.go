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
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/consts"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	"github.com/pterm/pterm"
	"os"
	"path/filepath"
	"strings"
)

func Env(c *command.Config) error {
	env := c.Env
	if env.List { // -l --list 列出当前安装的 CEF Framework 列表
		if err := printInstalledFrameworks(); err != nil {
			return err
		}
	} else if env.Use != "" { // --use 切换 CEF Framework
		if err := useInstalledFrameworks(env); err != nil {
			return err
		}
	} else if env.Get != "" { // -g --get 返回指定 name 值
		switch strings.ToLower(env.Get) {
		case "golang", "go", "goroot":
			term.Section.Println(GlobalDevEnvConfig.GoRoot)
		case "root":
			term.Section.Println(GlobalDevEnvConfig.Root, "=>", filepath.Join(GlobalDevEnvConfig.Root, consts.ENERGY))
		case "framework":
			term.Section.Println(GlobalDevEnvConfig.Framework, "=>", GlobalDevEnvConfig.FrameworkPath())
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
		case "version":
			term.Section.Println(GlobalDevEnvConfig.Version)
		}
	} else if env.Write != "" { // -w --write 写入指定 name 的值
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
			case "version":
				if val != "" {
					// 修复版本号
					if val[0] != 'v' {
						val = "v" + val
					}
					// 验证版本号格式
					if !tools.VerifyRelease(val) {
						err := fmt.Sprintf("Incorrect version format '%v'. Example: v1.0.0", val)
						return errors.New(err)
					}
				}
				GlobalDevEnvConfig.Version = val
			}
			GlobalDevEnvConfig.Update()
		}
	} else { // 默认展示当前环境配置
		PrintENV()
	}
	return nil
}

func chkOS(os command.OS) bool {
	return os.IsWindows() || os.IsLinux() || os.IsMacOS()
}

func chkARCH(arch command.Arch) bool {
	return arch.Is386() || arch.IsAMD64() || arch.IsARM() || arch.IsARM64() || arch.IsLoong64()
}

// 按完全匹配规则: CEF-[VER]_[OS]_[ARCH]
func getInstalledFrameworks() ([]string, error) {
	dirs, err := os.ReadDir(filepath.Join(GlobalDevEnvConfig.Root, consts.ENERGY))
	if err != nil {
		return nil, err
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
		isVER := tools.IsInt(verosarch[0])
		isOS := chkOS(command.OS(verosarch[1]))
		isARCH := chkARCH(command.Arch(verosarch[2]))
		if isName && isVER && isOS && isARCH {
			result = append(result, dirName)
		}
	}
	return result, nil
}

// 切换已安装的版本, 只做 CEF 版本和 ARCH 验证
func useInstalledFrameworks(env command.Env) error {
	// CEF-[VER]_[OS]_[ARCH]
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
		// CEF Framework: [VER]_[OS]_[ARCH]
		verosarch := strings.Split(split[1], "_")
		if len(verosarch) != 3 {
			continue
		}
		ver := verosarch[0]
		use := strings.Split(env.Use, ":")
		if len(use) == 1 {
			// 只验证版本号
			useVer := use[0]
			if useVer == ver {
				frameworkName = dir
				break
			}
		} else if len(use) == 2 {
			// 验证版本号和架构
			arch := command.Arch(verosarch[2])
			useVer := use[0]
			useArch := command.Arch(use[1])
			if useVer == ver && useArch.Value() == arch.Value() {
				frameworkName = dir
				break
			}
		}
	}
	if frameworkName != "" {
		GlobalDevEnvConfig.Framework = frameworkName
		GlobalDevEnvConfig.Update()
		msg := fmt.Sprintf("Now using CEF Framework %v", frameworkName)
		term.Logger.Info(msg)
		return nil
	} else {
		err := fmt.Sprintf("Not Installed %v. Use CLI: [energy install --cef %v]", env.Use, env.Use)
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

// 检查当前已安装且使用的框架是否正确
// Framework: CEF-[VER]_[OS]_[ARCH]
func CheckCEFDir() (isCEF bool, isLCL bool) {
	var libcef, liblcl = GetCurrentFrameworkLibName()
	if libcef == "" || liblcl == "" {
		return
	}
	isCEF = tools.IsExist(filepath.Join(GlobalDevEnvConfig.FrameworkPath(), libcef))
	isLCL = tools.IsExist(filepath.Join(GlobalDevEnvConfig.FrameworkPath(), liblcl))
	return
}

// 获取当前使用的框架的lib动态库名
// Framework: CEF-[VER]_[OS]_[ARCH]
func GetCurrentFrameworkLibName() (libcef, liblcl string) {
	framework := GlobalDevEnvConfig.Framework
	split := strings.Split(framework, "-")
	if len(split) != 2 {
		return
	}
	verosarch := strings.Split(split[1], "_")
	if len(verosarch) != 3 {
		return
	}
	isName := tools.Equals(split[0], "CEF")       // CEF
	isVER := tools.IsInt(verosarch[0])            // VER 100
	isOS := chkOS(command.OS(verosarch[1]))       // WINDOWS, LINUX, MACOS
	isARCH := chkARCH(command.Arch(verosarch[2])) // ARCH
	if isName && isVER && isOS && isARCH {
		OS := command.OS(verosarch[1])
		if OS.IsWindows() {
			return "libcef.dll", "liblcl.dll"
		} else if OS.IsLinux() {
			return "libcef.so", "liblcl.so"
		} else if OS.IsMacOS() {
			return "cef_sandbox.a", "liblcl.dylib"
		}
	}
	return "", ""
}

func PrintENV() {
	tableData := pterm.TableData{
		{"Name", "Value"},
	}
	tableData = append(tableData, []string{"Golang", GlobalDevEnvConfig.GoRoot})
	tableData = append(tableData, []string{"Version", GlobalDevEnvConfig.Version})
	tableData = append(tableData, []string{"Root", GlobalDevEnvConfig.Root})
	tableData = append(tableData, []string{"Framework", GlobalDevEnvConfig.Framework})
	tableData = append(tableData, []string{"NSIS", GlobalDevEnvConfig.NSIS})
	tableData = append(tableData, []string{"7z", GlobalDevEnvConfig.Z7Z})
	tableData = append(tableData, []string{"UPX", GlobalDevEnvConfig.UPX})
	tableData = append(tableData, []string{"Registry", GlobalDevEnvConfig.Registry})
	tableData = append(tableData, []string{"Proxy", GlobalDevEnvConfig.Proxy})
	pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("-").WithBoxed().WithData(tableData).Render()
}
