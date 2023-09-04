//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----

package env

import (
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/energye/golcl/energy/homedir"
	toolsCommand "github.com/energye/golcl/tools/command"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func SetGoEnv(goRoot string) {
	var goexe = "go"
	if command.IsWindows {
		goexe += ".exe"
	}
	gobin := filepath.Join(goRoot, "bin", goexe)
	if !tools.IsExist(gobin) {
		println("\nError: Failed to set the Golang environment variable, not a correct Golang installation directory. ", goRoot)
		return
	}
	println("\nSetting Golang environment Variables")
	cmd := toolsCommand.NewCMD()
	cmd.MessageCallback = func(s []byte, e error) {
		fmt.Println("CMD:", string(s), " error:", e)
	}
	defer cmd.Close()
	if command.IsWindows {
		// setx
		// GOROOT=/to/go/path
		var args = []string{"GOROOT", goRoot}
		cmd.Command("setx", args...)
		// GOCACHE=%GOROOT%\go-build
		args = []string{"GOCACHE", "%GOROOT%\\go-build"}
		cmd.Command("setx", args...)
		// GOBIN=%GOROOT%\bin
		args = []string{"GOBIN", "%GOROOT%\\bin"}
		cmd.Command("setx", args...)
		// PATH=%GOROOT%\bin
		args = []string{"path", "%path%;%GOROOT%\\bin"}
		cmd.Command("setx", args...)
	} else {

	}
	println("\nHint: reopen the cmd window for the Go command to take effect.")
}

func SetEnergyHomeEnv(homePath string) {
	var cef string
	if command.IsWindows {
		cef = "libcef.dll"
	} else if command.IsLinux {
		cef = "libcef.so"
	} else if command.IsDarwin {
		cef = "cef_sandbox.a"
	}
	cefPath := filepath.Join(homePath, cef)
	if !tools.IsExist(cefPath) {
		println("\nError: Setting the ENERGY_HOME environment variable failed and is not a correct CEF Framework installation directory. ", homePath)
		return
	}
	println("\nSetting environment Variables [ENERGY_HOME] to", homePath)
	cmd := toolsCommand.NewCMD()
	cmd.MessageCallback = func(s []byte, e error) {
		fmt.Println("CMD:", s, " error:", e)
	}
	defer cmd.Close()
	if command.IsWindows {
		var args = []string{"/c", "setx", command.EnergyHomeKey, homePath}
		cmd.Command("cmd.exe", args...)
	} else {
		var envFiles []string
		var energyHomeKey = fmt.Sprintf("export %s", command.EnergyHomeKey)
		var energyHome = fmt.Sprintf("export %s=%s", command.EnergyHomeKey, homePath)
		if command.IsLinux {
			envFiles = []string{".profile", ".zshrc", ".bashrc"}
		} else if command.IsDarwin {
			envFiles = []string{".profile", ".zshrc", ".bash_profile"}
		}
		homeDir, _ := homedir.Dir()
		for _, file := range envFiles {
			var fp = path.Join(homeDir, file)
			cmd.Command("touch", fp)
			f, err := os.OpenFile(fp, os.O_RDWR|os.O_APPEND, 0666)
			if err == nil {
				var oldContent string
				if contentBytes, err := ioutil.ReadAll(f); err == nil {
					content := string(contentBytes)
					oldContent = content
					var lines = strings.Split(content, "\n")
					var exist = false
					for i := 0; i < len(lines); i++ {
						line := lines[i]
						if strings.Index(line, energyHomeKey) == 0 {
							content = strings.ReplaceAll(content, line, energyHome)
							exist = true
						}
					}
					if exist {
						if err := f.Close(); err == nil {
							var oldWrite = func() {
								if f, err = os.OpenFile(fp, os.O_RDWR, 0666); err == nil {
									f.WriteString(oldContent)
									f.Close()
								}
							}
							if newOpenFile, err := os.OpenFile(fp, os.O_RDWR|os.O_TRUNC, 0666); err == nil {
								if _, err := newOpenFile.WriteString(content); err == nil {
									newOpenFile.Close()
								} else {
									newOpenFile.Close()
									oldWrite()
								}
							} else {
								oldWrite()
							}
						}
					} else {
						f.WriteString("\n")
						f.WriteString(energyHome)
						f.WriteString("\n")
					}
				} else {
					f.Close()
				}
			}
		}
	}
}
