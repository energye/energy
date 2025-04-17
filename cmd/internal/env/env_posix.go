//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !windows
// +build !windows

package env

import (
	"bytes"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cmd/internal/consts"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	toolsCommand "github.com/cyber-xxm/energy/v2/cmd/internal/tools/cmd"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools/homedir"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func homeKey(homeKey string) string {
	return "$" + homeKey
}

func SetGoEnv(goRoot string) {
	goBin := filepath.Join(goRoot, "bin", "go")
	if !tools.IsExist(goBin) {
		term.Logger.Error("Failed to set the Golang environment variable, not a correct Golang installation directory. " + goRoot)
		return
	}
	goCMD := filepath.Join(goRoot, "bin")
	term.Logger.Info("Setting Golang env: ", term.Logger.Args("PATH", "export PATH=$PATH:"+goCMD))
	setEnvToPath(goCMD)
	term.BoxPrintln("Hint: Restart the terminal and development tools for the commands to take effect.")
}

func SetEnergyCLIEnv(energyCLI string) {
	term.Logger.Info("Setting ENERGY CLI env: ", term.Logger.Args("PATH", "export PATH=$PATH:"+energyCLI))
	setEnvToPath(energyCLI)
	term.BoxPrintln("Hint: Restart the terminal and development tools for the commands to take effect.")
}

func envfiles() (result []string) {
	if consts.IsLinux {
		result = []string{".profile", ".zshrc", ".bashrc"}
	} else if consts.IsDarwin {
		result = []string{".profile", ".zshrc", ".bash_profile"}
	}
	return
}

func SourceEnvFiles() {
	term.Logger.Info("Refresh Environment Variables")
	cmd := toolsCommand.NewCMD()
	cmd.IsPrint = false
	defer cmd.Close()
	homeDir, err := homedir.Dir()
	if err != nil {
		term.Logger.Error(err.Error())
		return
	}
	var envFiles = envfiles()
	for _, file := range envFiles {
		var fp = filepath.Join(homeDir, file)
		// bash
		cmd.Command("bash", "-c", fmt.Sprintf("source %s", fp))
		// zsh
		cmd.Command("zsh", "-c", fmt.Sprintf("source %s", fp))
	}
}

func setEnvToPath(env string) {
	env = filepath.ToSlash(env)
	cmd := toolsCommand.NewCMD()
	cmd.IsPrint = false
	defer cmd.Close()
	var envFiles = envfiles()
	homeDir, err := homedir.Dir()
	if err != nil {
		term.Logger.Error(err.Error())
		return
	}
	// 拆分当前配置，确保是环境变量并且没有空字符
	var exportSplit = func(s string) (export, exportName, value string) {
		s1 := strings.Split(strings.TrimSpace(s), "=")
		if len(s1) < 2 {
			return "", "", ""
		}
		s2 := strings.Split(strings.TrimSpace(s1[0]), " ")
		for _, v := range s2 {
			v = strings.TrimSpace(v)
			if v != "" && export == "" {
				export = strings.ToLower(v)
			} else if v != "" && exportName == "" {
				exportName = strings.ToUpper(v)
			}
		}
		value = strings.TrimSpace(s1[1])
		return
	}
	// 检查path变量
	var isExportPath = func(line string) (bool, string) {
		export, exportName, value := exportSplit(line)
		return tools.Equals(export, "export") && tools.Equals(exportName, "PATH"), value
	}
	for _, file := range envFiles {
		var fp = filepath.Join(homeDir, file)
		cmd.Command("touch", fp)
		f, err := os.OpenFile(fp, os.O_RDWR|os.O_APPEND, 0666)
		if err == nil {
			if contentBytes, err := ioutil.ReadAll(f); err == nil {
				lines := strings.Split(string(contentBytes), "\n")
				isWritePath := true // env 是否写入到文件
				for i := 0; i < len(lines); i++ {
					line := strings.TrimRightFunc(lines[i], func(r rune) bool {
						if r == '\n' || r == '\r' {
							return true
						}
						return false
					})
					// is path
					if isPath, pathValue := isExportPath(line); isPath {
						// path 变量没有 env
						if tools.Equals(pathValue, "$PATH:"+env) {
							// 已经有了
							isWritePath = false
							break
						}
					}
				}
				// isWritePath == true 时，在原来的配置最后一行追加一个 PATH
				if err = f.Close(); err == nil {
					if isWritePath {
						// 打开文件覆盖模式
						if newOpenFile, err := os.OpenFile(fp, os.O_RDWR|os.O_TRUNC, 0666); err == nil {
							// 写入新的环境配置
							contentBuf := bytes.NewBuffer(contentBytes)
							contentBuf.WriteString("\n")
							contentBuf.WriteString("export PATH=$PATH:" + env)
							contentBuf.WriteString("\n")
							if _, err := newOpenFile.Write(contentBuf.Bytes()); err == nil {
								term.Logger.Info("Write files success. " + file)
							} else {
								term.Logger.Info("Write files ERROR: " + err.Error())
							}
							newOpenFile.Close()
						}
					}
				}
			} else {
				f.Close()
			}
		}
	}
}
