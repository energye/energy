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
	"bytes"
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
		//export GOROOT=/home/yanghy/app/go
		//export GOCACHE=$GOROOT/go-build
		//export GOBIN=$GOROOT/bin
		//export PATH=$PATH:$GOBIN
		//var exGoRoot = fmt.Sprintf("export GOROOT=%s", goRoot)
		//var exGoCache = "export GOCACHE==$GOROOT/go-build"
		//var exGoBin = "export GOBIN=$GOROOT/bin"
		//var exPath = "export PATH=$PATH:$GOBIN"
		//var envFiles []string
		//if command.IsLinux {
		//	envFiles = []string{".profile", ".zshrc", ".bashrc"}
		//} else if command.IsDarwin {
		//	envFiles = []string{".profile", ".zshrc", ".bash_profile"}
		//}
		//homeDir, err := homedir.Dir()
		//if err != nil {
		//	println(err.Error())
		//	return
		//}
		//for _, file := range envFiles {
		//	var fp = path.Join(homeDir, file)
		//	cmd.Command("touch", fp)
		//	f, err := os.OpenFile(fp, os.O_RDWR|os.O_APPEND, 0666)
		//	if err == nil {
		//		var oldContent string
		//		if contentBytes, err := ioutil.ReadAll(f); err == nil {
		//			content := string(contentBytes)
		//			oldContent = content
		//			var lines = strings.Split(content, "\n")
		//			var exist = false
		//			for i := 0; i < len(lines); i++ {
		//				line := lines[i]
		//				if strings.Index(line, energyHomeKey) == 0 {
		//					content = strings.ReplaceAll(content, line, energyHome)
		//					exist = true
		//				}
		//			}
		//		}
		//	}
		//}
	}
	println("\nHint: Reopen the cmd window for the Go command to take effect.")
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
		homeDir, err := homedir.Dir()
		if err != nil {
			println(err.Error())
			return
		}
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
						f.Close()
					}
				} else {
					f.Close()
				}
			}
		}
	}
	println("\nHint: Reopen the cmd window to make the environment variables take effect.")
}

func setEnv(exs []string, exGoBinPath, bin string) {
	cmd := toolsCommand.NewCMD()
	cmd.MessageCallback = func(s []byte, e error) {
		fmt.Println("CMD:", string(s), " error:", e)
	}
	defer cmd.Close()
	var envFiles []string
	if command.IsLinux {
		envFiles = []string{".profile", ".zshrc", ".bashrc"}
	} else if command.IsDarwin {
		envFiles = []string{".profile", ".zshrc", ".bash_profile"}
	}
	homeDir, err := homedir.Dir()
	// test
	homeDir = "E:\\app"
	envFiles = []string{".profile", ".zshrc", ".bashrc"}
	if err != nil {
		println(err.Error())
		return
	}
	var isExport = func(line string) (string, bool) {
		for _, ex := range exs {
			exName := strings.Split(ex, "=")[0]
			if strings.Index(line, exName) == 0 {
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
				var exist = false
				var gobin = false
				for i := 0; i < len(lines); i++ {
					line := strings.TrimSpace(lines[i])
					if line == "" {
						continue
					}
					if exGoBinPath != "" && isExportPath(line) {
						//是path，并且有gobin
						newContent.WriteString(exGoBinPath)
						gobin = true
					} else {
						// 其它变量, 判断是否存在，如果存在替换, 否则将原来的添加进来
						if ex, ok := isExport(line); ok {
							newContent.WriteString(ex)
							exist = true
						} else {
							newContent.WriteString(line)
						}
					}
					newContent.WriteString("\n")
				}
				if exist {
					// 如果path里没有gobin, 添加一个
					if exGoBinPath != "" && !gobin {
						newContent.WriteString(exGoBinPath)
						newContent.WriteString("\n")
					}
					// 有就覆盖掉之前的, 要先关闭掉文件
					if err := f.Close(); err == nil {
						// 如果关闭文件流成功, 重新写入覆盖文件
						var oldWrite = func() {
							if f, err = os.OpenFile(fp, os.O_RDWR, 0666); err == nil {
								f.WriteString(oldContent)
								f.Close()
							}
						}
						// 打开覆盖模式
						if newOpenFile, err := os.OpenFile(fp, os.O_RDWR|os.O_TRUNC, 0666); err == nil {
							// 写入，如果失败，把老的写入
							if _, err := newOpenFile.Write(newContent.Bytes()); err == nil {
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
					// 都没有全添加一次
					for _, ex := range exs {
						f.WriteString("\n")
						f.WriteString(ex)
						f.WriteString("\n")
					}
					// 如果path里没有gobin, 添加一个
					if exGoBinPath != "" && !gobin {
						f.WriteString(exGoBinPath)
						f.WriteString("\n")
					}
					f.Close()
				}
			} else {
				f.Close()
			}
		}
	}
}
