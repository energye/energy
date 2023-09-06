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
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/energye/golcl/energy/homedir"
	toolsCommand "github.com/energye/golcl/tools/command"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func SetNSISEnv(nsisRoot string) {

}

func SetGoEnv(goRoot string) {
	goBin := filepath.Join(goRoot, "bin", "go")
	if !tools.IsExist(goBin) {
		println("\nError: Failed to set the Golang environment variable, not a correct Golang installation directory. ", goRoot)
		return
	}
	println("\nSetting Golang environment Variables to:", goRoot)
	cmd := toolsCommand.NewCMD()
	cmd.IsPrint = false
	cmd.MessageCallback = func(s []byte, e error) {
		msg := strings.TrimSpace(string(s))
		if msg != "" {
			fmt.Println("CMD:", msg)
		}
		if e != nil {
			fmt.Println("CMD Error:", e)
		}
	}
	defer cmd.Close()
	//export GOROOT=/home/yanghy/app/go
	//export GOCACHE=$GOROOT/go-build
	//export GOBIN=$GOROOT/bin
	//export PATH=$PATH:$GOBIN
	var exGoRoot = fmt.Sprintf("export GOROOT=%s", goRoot)
	var exGoCache = "export GOCACHE==$GOROOT/go-build"
	var exGoBin = "export GOBIN=$GOROOT/bin"
	var exPath = "export PATH=$PATH:$GOBIN"
	var exs = []string{exGoRoot, exGoCache, exGoBin}
	setPosixEnv(exs, exPath, "$GOBIN")
	println("\nHint: Reopen the cmd window for the Go command to take effect.")
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
		println("\nError: Setting the ENERGY_HOME environment variable failed and is not a correct CEF Framework installation directory. ", homePath)
		return
	}
	println("\nSetting environment Variables [ENERGY_HOME] to", homePath)
	cmd := toolsCommand.NewCMD()
	cmd.IsPrint = false
	cmd.MessageCallback = func(s []byte, e error) {
		msg := strings.TrimSpace(string(s))
		if msg != "" {
			fmt.Println("CMD:", msg)
		}
		if e != nil {
			fmt.Println("CMD Error:", e)
		}
	}
	defer cmd.Close()
	var energyHome = fmt.Sprintf("export %s=%s", consts.EnergyHomeKey, homePath)
	exs := []string{energyHome}
	setPosixEnv(exs, "", "")
	println("\nHint: Reopen the cmd window to make the environment variables take effect.")
}

func setPosixEnv(exs []string, binPath, bin string) {
	cmd := toolsCommand.NewCMD()
	cmd.IsPrint = false
	cmd.MessageCallback = func(s []byte, e error) {
		msg := strings.TrimSpace(string(s))
		if msg != "" {
			fmt.Println("CMD:", msg)
		}
		if e != nil {
			fmt.Println("CMD Error:", e)
		}
	}
	defer cmd.Close()
	var envFiles []string
	if consts.IsLinux {
		envFiles = []string{".profile", ".zshrc", ".bashrc"}
	} else if consts.IsDarwin {
		envFiles = []string{".profile", ".zshrc", ".bash_profile"}
	}
	homeDir, err := homedir.Dir()
	if err != nil {
		println(err.Error())
		return
	}
	var tempExts []string
	var isExport = func(line string) (string, bool) {
		for i, ex := range tempExts {
			exName := strings.Split(ex, "=")[0]
			if strings.Index(line, exName) == 0 {
				tempExts = append(tempExts[i+1:], tempExts[:i]...)
				return ex, true
			}
		}
		return "", false
	}
	var isExportPath = func(line string) bool {
		if strings.Index(line, "export PATH") == 0 {
			if strings.Contains(line, bin) {
				return true
			}
		}
		return false
	}
	for _, file := range envFiles {
		tempExts = exs[:]
		var fp = path.Join(homeDir, file)
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
					if binPath != "" && isExportPath(line) {
						currentPath = line
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
				if currentPath != "" { // 不为空说明现有的path里已经存在 bin, 放到最后一行
					newContent.WriteString(currentPath)
					newContent.WriteString("\n")
				} else if binPath != "" { //不为空并且现有的path里没有bin时，自己设置一个带有bin 的 path, 放到最后一行
					newContent.WriteString(binPath)
					newContent.WriteString("\n")
				}
				// 有就覆盖掉之前的, 要先关闭掉文件
				if err = f.Close(); err == nil {
					// 如果任何操作失败, 重新写入覆盖文件
					var oldWrite = func() {
						if f, err = os.OpenFile(fp, os.O_RDWR, 0666); err == nil {
							println("Restore files", file)
							f.WriteString(oldContent)
							f.Close()
						}
					}
					// 打开文件覆盖模式
					if newOpenFile, err := os.OpenFile(fp, os.O_RDWR|os.O_TRUNC, 0666); err == nil {
						// 写入新的环境配置
						if _, err := newOpenFile.Write(newContent.Bytes()); err == nil {
							println("Write files success.", file)
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
