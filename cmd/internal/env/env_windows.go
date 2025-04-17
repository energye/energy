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

package env

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	toolsCommand "github.com/cyber-xxm/energy/v2/cmd/internal/tools/cmd"
	"path/filepath"
	"strings"
)

func SetGoEnv(goRoot string) {
	goBin := filepath.Join(goRoot, "bin", "go.exe")
	if !tools.IsExist(goBin) {
		term.Logger.Error("Error: Failed to set the Golang environment variable, not a correct Golang installation directory. " + goRoot)
		return
	}
	term.Logger.Info("Setting Golang environment Variables: ", term.Logger.Args("GOROOT", goRoot, "PATH", "%GOROOT%\\bin"))
	setWindowsEnv("GOROOT", goRoot)
	appendWindowsEnv("Path", "%GOROOT%\\bin")
	term.BoxPrintln("Hint: Restart the terminal and development tools for the commands to take effect.")
}

func SetEnergyCLIEnv(energyCLI string) {
	term.Logger.Info("Setting ENERGY CLI env: ", term.Logger.Args("PATH", energyCLI))
	appendWindowsEnv("Path", energyCLI)
	term.BoxPrintln("Hint: Restart the terminal and development tools for the commands to take effect.")
}

func setWindowsEnv(name, value string) {
	cmd := toolsCommand.NewCMD()
	cmd.IsPrint = false
	cmd.Command("setx", name, value)
	cmd.Close()
}

func appendWindowsEnv(name, value string) {
	regCurUser := tools.NewRegistryCurrentUser()
	defer regCurUser.Close()
	oldValue, err := regCurUser.Read(name)
	if err == nil {
		// 可变变量替换完整路径
		var fullValuePath = func(value []string) string {
			for i, val := range value {
				values := strings.Split(val, "\\")
				for i, vab := range values {
					vab = strings.TrimSpace(vab)
					if vab == "" || len(vab) <= 2 {
						continue
					}
					if vab[0] == '%' && vab[len(vab)-1] == '%' {
						vab = vab[1 : len(vab)-1]
						if v, err := regCurUser.Read(vab); err == nil {
							values[i] = v
						}
					}
				}
				value[i] = strings.Join(values, "\\")
			}
			return strings.Join(value, ";")
		}
		// 转换完整路径
		valueFull := fullValuePath([]string{value})
		oldValueFull := fullValuePath(strings.Split(oldValue, ";"))
		oldValues := strings.Split(oldValueFull, ";")
		// 检测当前设置变量如果已存在就跳出
		for _, oval := range oldValues {
			if strings.TrimSpace(valueFull) == strings.TrimSpace(oval) {
				return
			}
		}
		cmd := toolsCommand.NewCMD()
		cmd.IsPrint = false
		cmd.Command("setx", name, fmt.Sprintf("%s;%s", oldValue, value))
		cmd.Close()
	} else {
		// 没有就设置一个新的
		setWindowsEnv(name, value)
	}
}
