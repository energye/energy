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
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/pterm/pterm"
	"path/filepath"
	"strings"
)

func Env(c *command.Config) error {
	env := c.Env
	if env.Write == "" && env.Get == "" {
		PrintENV()
	} else if env.Get != "" {
		switch strings.ToLower(env.Get) {
		case "golang", "go", "goroot":
			term.Section.Println(GlobalDevEnvConfig.GoRoot)
		case "framework":
			term.Section.Println(GlobalDevEnvConfig.Framework)
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
		case "cef":
			term.Section.Println(GlobalDevEnvConfig.CEF)
		case "ver":
			term.Section.Println(GlobalDevEnvConfig.VER)
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
			case "cef":
				GlobalDevEnvConfig.CEF = val
			case "ver":
				GlobalDevEnvConfig.VER = val
			}
			GlobalDevEnvConfig.Update()
		}
	}
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
		return tools.IsExist(filepath.Join(GlobalDevEnvConfig.Framework, lib))
	}
	return false
}

func PrintENV() {
	tableData := pterm.TableData{
		{"Name", "Directory"},
	}
	tableData = append(tableData, []string{"CEF", GlobalDevEnvConfig.CEF})
	tableData = append(tableData, []string{"VER", GlobalDevEnvConfig.VER})
	tableData = append(tableData, []string{"Golang", GlobalDevEnvConfig.GoRoot})
	tableData = append(tableData, []string{"Framework", GlobalDevEnvConfig.Framework})
	tableData = append(tableData, []string{"NSIS", GlobalDevEnvConfig.NSIS})
	tableData = append(tableData, []string{"7z", GlobalDevEnvConfig.Z7Z})
	tableData = append(tableData, []string{"UPX", GlobalDevEnvConfig.UPX})
	tableData = append(tableData, []string{"Registry", GlobalDevEnvConfig.Registry})
	tableData = append(tableData, []string{"Proxy", GlobalDevEnvConfig.Proxy})
	pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("-").WithBoxed().WithData(tableData).Render()
}
