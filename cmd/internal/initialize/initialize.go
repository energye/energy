//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package initialize

import (
	"embed"
	"errors"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/tools"
	toolsCommand "github.com/energye/golcl/tools/command"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//go:embed assets
var assets embed.FS

func InitEnergyProject(c *command.Config) error {
	// 检查环境
	checkEnv(&c.Init)
	// 生成项目
	if err := generaProject(c); err != nil {
		return err
	}
	return nil
}

func generaProject(c *command.Config) error {
	projectPath := filepath.Join(c.Wd, c.Init.Name)
	println("Create Project:", c.Init.Name)
	if tools.IsExist(projectPath) {
		fmt.Printf("project dir %s exist, Do you want to delete and recreate? Y/n:  ", c.Init.Name)
		var s string
		fmt.Scan(&s)
		println()
		if strings.ToLower(s) != "y" {
			return errors.New("Failed to initialize project " + c.Init.Name)
		} else {
			var deleteFiles = []string{"energy.json", "resources", "main.go", "go.mod", "go.sum"}
			for _, f := range deleteFiles {
				path := filepath.Join(projectPath, f)
				if info, err := os.Lstat(path); err == nil {
					if info.IsDir() {
						os.RemoveAll(path)
					} else {
						os.ReadFile(path)
					}
				}
			}
		}
	} else {
		// 创建目录
		if err := os.Mkdir(projectPath, fs.ModePerm); err != nil {
			return err
		}
	}

	// 创建 energy.json template
	if energyText, err := assets.ReadFile("assets/energy.json"); err != nil {
		return err
	} else {
		data := make(map[string]any)
		data["Name"] = c.Init.Name
		if content, err := tools.RenderTemplate(string(energyText), data); err != nil {
			return err
		} else {
			path := filepath.Join(projectPath, "energy.json")
			if err = ioutil.WriteFile(path, content, 0666); err != nil {
				return err
			}
		}
	}

	// 创建 main.go
	if mainText, err := assets.ReadFile(fmt.Sprintf("assets/app_load_res.go.%s", c.Init.ResLoad)); err != nil {
		return err
	} else {
		path := filepath.Join(projectPath, "main.go")
		if err = ioutil.WriteFile(path, mainText, 0666); err != nil {
			return err
		}
	}

	// 创建 go.mod
	if modText, err := assets.ReadFile("assets/go.mod.t"); err != nil {
		return err
	} else {
		data := make(map[string]any)
		data["Name"] = c.Init.Name
		data["GoVersion"] = "1.18"
		data["EnergyVersion"] = "latest"
		if content, err := tools.RenderTemplate(string(modText), data); err != nil {
			return err
		} else {
			path := filepath.Join(projectPath, "go.mod")
			if err = ioutil.WriteFile(path, content, 0666); err != nil {
				return err
			}
		}
	}

	// cmd
	println("cmd run")
	cmd := toolsCommand.NewCMD()
	cmd.Dir = projectPath
	cmd.MessageCallback = func(bytes []byte, err error) {
		fmt.Println("CMD:", bytes, " error:", err)
	}
	// cmd go env -w GO111MODULE=on
	println("Enable Go mod management")
	cmd.Command("go", []string{"env", "-w", "GO111MODULE=on"}...)

	// cmd go env -w GOPROXY=https://goproxy.io,direct
	println("Configure mod agent")
	cmd.Command("go", []string{"env", "-w", "GOPROXY=https://goproxy.io,direct"}...)

	// cmd go mod tidy
	println("Update Energy dependencies, version:")
	cmd.Command("go", []string{"mod", "tidy"}...)
	cmd.Close()
	return nil
}

func checkEnv(init *command.Init) {
	println("Check the current environment and follow the prompts if there are any warnings")
	// 检查Go环境
	if !tools.CommandExists("go") {
		println("Warning: Golang development environment not installed, Download-URL: https://golang.google.cn/dl/")
	} else {
		println("\tGolang OK")
		init.IGo = true
	}
	// 检查nsis
	if !tools.CommandExists("makensis") {
		println(`Warning: NSIS not installed, Unable to create installation package through energy command line`)
	} else {
		println("\tNSIS OK")
		init.INSIS = true
	}
	// 检查ENERGY_HOME
	if !cef() {
		println(`Warning: Dependency framework CEF is not installed or configured to the ENERGY_HOME environment variable
	There are several ways to install, configure, or check the environment
		"energy install ." Installation and development environment
		"energy env ." check ENERGY_HOME are correctly, 
		"energy setenv -p /to/framework/path ." set ENERGY_HOME environment
`)
	} else {
		println("\tCEF Framework OK")
		init.IEnv = true
	}
	// 检查 node npm
	if !tools.CommandExists("npm") {
		println(`Warning: Installing node allows you to build the UI using, for example, a front-end framework (vue), Download URL: https://nodejs.org/`)
	} else {
		println("\tNode npm OK")
		init.INpm = true
	}
}

func cef() bool {
	var lib = func() string {
		if command.IsWindows {
			return "libcef.dll"
		} else if command.IsLinux {
			return "libcef.so"
		} else if command.IsDarwin {
			return "cef_sandbox.a"
		}
		return ""
	}()
	if lib != "" {
		return tools.IsExist(filepath.Join(os.Getenv("ENERGY_HOME"), lib))
	}
	return false
}
