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
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	toolsCommand "github.com/energye/energy/v2/tools/command"
	"path/filepath"
	"strings"
)

func homeKey(homeKey string) string {
	return "%" + homeKey + "%"
}

func SetUPXEnv(upxRoot string) {
	upx := filepath.Join(upxRoot, "upx.exe")
	if !tools.IsExist(upx) {
		term.Logger.Error("Failed to set the UPX environment variable, not a correct UPX installation directory. " + upxRoot)
		return
	}
	term.Logger.Info("Setting UPX environment Variables: ", term.Logger.Args(consts.UPXHomeKey, upxRoot))
	setWindowsEnv(consts.UPXHomeKey, upxRoot)
	appendWindowsEnv("Path", homeKey(consts.UPXHomeKey))
	term.BoxPrintln("Hint: Reopen the cmd window for the upx command to take effect.")
}

func SetNSISEnv(nsisRoot string) {
	makensis := filepath.Join(nsisRoot, "makensis.exe")
	if !tools.IsExist(makensis) {
		term.Logger.Error("Error: Failed to set the NSIS environment variable, not a correct NSIS installation directory. " + nsisRoot)
		return
	}
	term.Logger.Info("Setting NSIS environment Variables: ", term.Logger.Args(consts.NSISHomeKey, nsisRoot))
	//regCurUser := tools.NewRegistryCurrentUser()
	//defer regCurUser.Close()
	//regCurUser.Set(consts.NSISHomeKey, nsisRoot)
	//regCurUser.Append("Path", "%NSIS_HOME%")
	//cmd
	setWindowsEnv(consts.NSISHomeKey, nsisRoot)
	appendWindowsEnv("Path", homeKey(consts.NSISHomeKey))
	term.BoxPrintln("Hint: Reopen the cmd window for the makensis command to take effect.")
}

func Set7zaEnv(z7zRoot string) {
	makensis := filepath.Join(z7zRoot, "7za.exe")
	if !tools.IsExist(makensis) {
		term.Logger.Error("Error: Failed to set the 7za environment variable, not a correct 7za installation directory. " + z7zRoot)
		return
	}
	term.Logger.Info("Setting 7za environment Variables: ", term.Logger.Args(consts.Z7ZHomeKey, z7zRoot))
	setWindowsEnv(consts.Z7ZHomeKey, z7zRoot)
	appendWindowsEnv("Path", homeKey(consts.Z7ZHomeKey))
	term.BoxPrintln("Hint: Reopen the cmd window for the 7za command to take effect.")
}

func SetGoEnv(goRoot string) {
	goBin := filepath.Join(goRoot, "bin", "go.exe")
	if !tools.IsExist(goBin) {
		term.Logger.Error("Error: Failed to set the Golang environment variable, not a correct Golang installation directory. " + goRoot)
		return
	}
	term.Logger.Info("Setting Golang environment Variables: ", term.Logger.Args("GOROOT", goRoot, "GOCACHE", "%GOROOT%\\go-build", "GOBIN", "%GOROOT%\\bin"))
	//regCurUser := tools.NewRegistryCurrentUser()
	//defer regCurUser.Close()
	//regCurUser.Set("GOROOT", goRoot)
	//regCurUser.Set("GOCACHE", "%GOROOT%\\go-build")
	//regCurUser.Set("GOBIN", "%GOROOT%\\bin")
	//regCurUser.Append("Path", "%GOBIN%")
	// cmd
	setWindowsEnv("GOROOT", goRoot)
	setWindowsEnv("GOCACHE", "%GOROOT%\\go-build")
	setWindowsEnv("GOBIN", "%GOROOT%\\bin")
	appendWindowsEnv("Path", "%GOROOT%\\bin")
	term.BoxPrintln("Hint: Reopen the cmd window for the Go command to take effect.")
}

func SetEnergyHomeEnv(homePath string) {
	cefPath := filepath.Join(homePath, "libcef.dll")
	if !tools.IsExist(cefPath) {
		term.Logger.Error("Error: Setting ENERGY_HOME environment variable failed and is not a correct CEF Framework installation directory. " + homePath)
		return
	}
	term.Logger.Info("Setting ENERGY environment Variables [ENERGY_HOME] to " + homePath)
	//regCurUser := tools.NewRegistryCurrentUser()
	//defer regCurUser.Close()
	//regCurUser.Set(consts.EnergyHomeKey, homePath)
	//cmd
	setWindowsEnv(consts.EnergyHomeKey, homePath)
	term.BoxPrintln("Hint: Reopen the cmd window to make the environment variables take effect.")
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

func SourceEnvFiles() {

}
