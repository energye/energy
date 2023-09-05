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
	println()
	println("Successfully initialized the energy application project:", c.Init.Name)
	println()
	println("  Run Application: go run main.go")
	println(`  Building Applications:
	Use GO: go build -ldflags "-s -w"
	Use Energy: energy build .
`)
	println(`  website:
	https://github.com/energye/energy
	https://energy.yanghy.cn
`)
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
			var deleteFiles = []string{"energy.json", "resources", "main.go", "go.mod", "go.sum", "resources/index.html", "README.md"}
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

	// 创建 resources/index.html
	if err := os.Mkdir(filepath.Join(projectPath, "resources"), fs.ModePerm); err != nil {
		return err
	} else {
		if indexText, err := assets.ReadFile("assets/index.html"); err != nil {
			return err
		} else {
			path := filepath.Join(projectPath, "resources", "index.html")
			if err = ioutil.WriteFile(path, indexText, 0666); err != nil {
				return err
			}
		}
	}

	// 创建 README.md
	if readmeText, err := assets.ReadFile("assets/README.md"); err != nil {
		return err
	} else {
		data := make(map[string]any)
		data["Name"] = c.Init.Name
		if content, err := tools.RenderTemplate(string(readmeText), data); err != nil {
			return err
		} else {
			path := filepath.Join(projectPath, "README.md")
			if err = ioutil.WriteFile(path, content, 0666); err != nil {
				return err
			}
		}
	}

	// cmd
	println("Run command-line")
	cmd := toolsCommand.NewCMD()
	cmd.Dir = projectPath
	cmd.MessageCallback = func(bytes []byte, err error) {
		s := string(bytes)
		if s != "" {
			println("\tCMD:", string(bytes))
		} else if err != nil {
			println("\tCMD-error:", err.Error())
		}
	}
	if c.Init.IGo {
		// cmd go env -w GO111MODULE=on
		println("Enable Go mod management")
		cmd.Command("go", []string{"env", "-w", "GO111MODULE=on"}...)

		// cmd go env -w GOPROXY=https://goproxy.io,direct
		println("Configure mod agent")
		cmd.Command("go", []string{"env", "-w", "GOPROXY=https://goproxy.io,direct"}...)

		// cmd go mod tidy
		println("Update Energy dependencies, version:")
		cmd.Command("go", []string{"mod", "tidy"}...)
	}
	cmd.Close()
	return nil
}

func checkEnv(init *command.Init) {
	println("Check the current environment and follow the prompts if there are any warnings")
	// 检查Go环境
	if !tools.CommandExists("go") {
		println("Warning: Golang development environment not installed, Download-URL: https://golang.google.cn/dl/")
	} else {
		var version string
		cmd := toolsCommand.NewCMD()
		cmd.IsPrint = false
		cmd.MessageCallback = func(bytes []byte, err error) {
			data := string(bytes)
			if strings.Index(data, "go version") != -1 {
				d := strings.Split(data, " ")
				if len(d) == 4 {
					version = d[2][2:] // x.x.x
				}
			}
		}
		cmd.Command("go", "version")
		cmd.Close()
		if version != "" {
			d := strings.Split(version, ".")
			if len(d) == 3 {
				if tools.ToInt(d[0]) < 1 || tools.ToInt(d[1]) < 18 {
					println(`Warning: current installed Go version should be greater than 1.18. version:`, version)
				}
			}
		}
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
	if !tools.CheckCEFDir() {
		println(`Warning: Dependency framework CEF is not installed or configured to the ENERGY_HOME environment variable
	There are several ways to install, configure, or check the environment
		"energy install ." Installation and development environment
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
