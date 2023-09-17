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
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/energye/golcl/energy/homedir"
	toolsCommand "github.com/energye/golcl/tools/command"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func homeKey(homeKey string) string {
	return "$" + homeKey
}

func SetUPXEnv(upxRoot string) {
	upx := filepath.Join(upxRoot, "upx")
	if !tools.IsExist(upx) {
		term.Logger.Error("Failed to set the UPX environment variable, not a correct UPX installation directory. " + upxRoot)
		return
	}
	term.Logger.Info("Setting UPX environment Variables: ", term.Logger.Args(consts.UPXHomeKey, upxRoot))
	var exUpxRoot = fmt.Sprintf("export %s=%s", consts.UPXHomeKey, upxRoot)
	var exPath = "export PATH=$PATH:" + homeKey(consts.UPXHomeKey)
	var exs = []string{exUpxRoot}
	setPosixEnv(exs, exPath, homeKey(consts.UPXHomeKey))
	term.BoxPrintln("Hint: Reopen the cmd window for the upx command to take effect.")
}

func SetNSISEnv(nsisRoot string) {

}

func Set7zaEnv(z7zRoot string) {

}

func SetGoEnv(goRoot string) {
	goBin := filepath.Join(goRoot, "bin", "go")
	if !tools.IsExist(goBin) {
		term.Logger.Error("Error: Failed to set the Golang environment variable, not a correct Golang installation directory. " + goRoot)
		return
	}
	term.Logger.Info("Setting Golang environment Variables: ", term.Logger.Args("GOROOT", goRoot, "GOCACHE", "%GOROOT%\\go-build", "GOBIN", "%GOROOT%\\bin"))
	//export GOROOT=/home/yanghy/app/go
	//export GOCACHE=$GOROOT/go-build
	//export GOBIN=$GOROOT/bin
	//export PATH=$PATH:$GOBIN
	var exGoRoot = fmt.Sprintf("export GOROOT=%s", goRoot)
	var exGoCache = "export GOCACHE=$GOROOT/go-build"
	var exGoBin = "export GOBIN=$GOROOT/bin"
	var exPath = "export PATH=$PATH:$GOBIN"
	var exs = []string{exGoRoot, exGoCache, exGoBin}
	setPosixEnv(exs, exPath, homeKey("GOBIN"))
	term.BoxPrintln("Hint: Reopen the cmd window for the Go command to take effect.")
}

func SetEnergyHomeEnv(homePath string) {
	var cef string
	if consts.IsLinux {
		cef = "libcef.so"
	} else if consts.IsDarwin {
		cef = "cef_sandbox.a"
	}
	cefPath := filepath.Join(homePath, cef)
	if !tools.IsExist(cefPath) {
		term.Logger.Error("Error: Setting ENERGY_HOME environment variable failed and is not a correct CEF Framework installation directory. " + homePath)
		return
	}
	term.Logger.Info("Setting ENERGY environment Variables [ENERGY_HOME] to " + homePath)
	var energyHome = fmt.Sprintf("export %s=%s", consts.EnergyHomeKey, homePath)
	exs := []string{energyHome}
	setPosixEnv(exs, "", "")
	term.BoxPrintln("Hint: Reopen the cmd window to make the environment variables take effect.")
}

func envfiles() (result []string) {
	if consts.IsLinux {
		result = []string{".profile", ".zshrc", ".bashrc"}
	} else if consts.IsDarwin {
		result = []string{".profile", ".zshrc", ".bash_profile"}
	}
	return result
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

func setPosixEnv(exs []string, binPath, bin string) {
	cmd := toolsCommand.NewCMD()
	cmd.IsPrint = false
	defer cmd.Close()
	var envFiles = envfiles()
	homeDir, err := homedir.Dir()
	if err != nil {
		term.Logger.Error(err.Error())
		return
	}
	var tempExts []string
	// 拆分当前配置，确保是环境变量并且没有空字符
	var exportSplit = func(s string) (export, exportName string) {
		s1 := strings.Split(strings.TrimSpace(s), "=")
		if len(s1) < 2 {
			return "", ""
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
		return
	}
	// 检查当前要导出的变量
	var isExport = func(line string) (string, bool) {
		export, exportName := exportSplit(line)
		for i, ex := range tempExts {
			exExport, exExportName := exportSplit(ex)
			if exExport == export && exportName == exExportName {
				tempExts = append(tempExts[i+1:], tempExts[:i]...)
				return ex, true
			}
		}
		return "", false
	}
	// 检查path变量
	var isExportPath = func(line string) bool {
		export, exportName := exportSplit(line)
		return export == "export" && exportName == "PATH"
	}
	// 检查bin path变量
	var isExportBinPath = func(line string) (isPath, isBin bool) {
		isPath = isExportPath(line)
		if isPath {
			isBin = strings.Contains(line, bin)
		}
		return
	}
	for _, file := range envFiles {
		tempExts = exs[:]
		var fp = filepath.Join(homeDir, file)
		cmd.Command("touch", fp)
		f, err := os.OpenFile(fp, os.O_RDWR|os.O_APPEND, 0666)
		if err == nil {
			var oldContent string
			if contentBytes, err := ioutil.ReadAll(f); err == nil {
				content := string(contentBytes)
				oldContent = content
				var newContent = new(bytes.Buffer)
				var lines = strings.Split(content, "\n")
				var currentPath string
				for i := 0; i < len(lines); i++ {
					line := strings.TrimRightFunc(lines[i], func(r rune) bool {
						if r == '\n' || r == '\r' {
							return true
						}
						return false
					})
					// path里是否存在要设置的bin
					if isPath, isBin := isExportBinPath(line); binPath != "" && isBin {
						currentPath = line
					} else if isPath {
						newContent.WriteString(line)
						newContent.WriteString("\n")
					} else {
						// 其它变量, 判断是否存在，如果存在替换, 否则将原来的添加进来
						if ex, ok := isExport(line); ok {
							newContent.WriteString(ex)
						} else {
							newContent.WriteString(line)
						}
						if i < len(lines)-1 {
							newContent.WriteString("\n")
						}
					}
				}
				for _, ext := range tempExts {
					newContent.WriteString(ext)
					newContent.WriteString("\n")
				}
				// 最后一行设置 path
				if currentPath != "" { // 不为空说明现有的path里已经存在 bin, 放到最后一行
					newContent.WriteString(currentPath)
					newContent.WriteString("\n")
				} else if binPath != "" { //不为空并且现有的path里没有bin时，自己设置一个带有bin 的 新path, 放到最后一行
					newContent.WriteString(binPath)
					newContent.WriteString("\n")
				}
				// 有就覆盖掉之前的, 要先关闭掉文件
				if err = f.Close(); err == nil {
					// 如果任何操作失败, 重新写入覆盖文件
					var oldWrite = func() {
						if f, err = os.OpenFile(fp, os.O_RDWR, 0666); err == nil {
							term.Logger.Info("Restore files. " + file)
							f.WriteString(oldContent)
							f.Close()
						}
					}
					// 打开文件覆盖模式
					if newOpenFile, err := os.OpenFile(fp, os.O_RDWR|os.O_TRUNC, 0666); err == nil {
						// 写入新的环境配置
						if _, err := newOpenFile.Write(newContent.Bytes()); err == nil {
							term.Logger.Info("Write files success. " + file)
							newOpenFile.Close()
						} else {
							//写入失败，把老的内容还原
							newOpenFile.Close()
							oldWrite()
						}
					} else {
						//打开失败，把老的内容还原
						oldWrite()
					}
				}
			} else {
				f.Close()
			}
		}
	}
}
