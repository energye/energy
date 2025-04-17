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
	"fmt"
	"github.com/cyber-xxm/energy/v2/cmd/internal/assets"
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/consts"
	"github.com/cyber-xxm/energy/v2/cmd/internal/env"
	"github.com/cyber-xxm/energy/v2/cmd/internal/project"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	toolsCommand "github.com/cyber-xxm/energy/v2/cmd/internal/tools/cmd"
	"github.com/pterm/pterm"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func InitEnergyProject(c *command.Config) error {
	// 检查环境
	checkEnv(&c.Init)
	// 生成项目
	if err := generaProject(c); err != nil {
		return err
	}
	pterm.Println()
	term.Section.Println("Successfully initialized the energy application project:", c.Init.Name)
	term.Logger.Info("Website", term.Logger.Args("Github", "https://github.com/energye/energy", "ENERGY", "https://energye.github.io"))
	term.Section.Println("Run Application")
	tableData := pterm.TableData{
		{"command"}, {"go run main.go"},
	}
	err := pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("-").WithBoxed().WithData(tableData).Render()
	if err != nil {
		return err
	}
	term.Section.Println("Building Application")
	tableData = pterm.TableData{
		{"name", "command"},
	}
	tableData = append(tableData, []string{"go", `go build -ldflags "-s -w"`})
	tableData = append(tableData, []string{"energy", `energy build`})
	err = pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("-").WithBoxed().WithData(tableData).Render()
	if err != nil {
		return err
	}
	term.Section.Println("Make install package")
	tableData = pterm.TableData{
		{"command"}, {"energy package"},
	}
	err = pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("-").WithBoxed().WithData(tableData).Render()
	if err != nil {
		return err
	}
	return nil
}

func generaProject(c *command.Config) error {
	pterm.Println()
	projectPath := filepath.Join(c.Wd, c.Init.Name)
	term.Logger.Info("Create Project", term.Logger.Args("Name", c.Init.Name))
	if tools.IsExist(projectPath) {
		term.Logger.Warn(fmt.Sprintf("Project dir `%s` exist, delete init default files.", c.Init.Name))
		deleteFiles := []string{ /*"resources", */ "main.go", "go.mod", "go.sum", "resources/index.html", "README.md"}
		for _, fileName := range consts.EnergyProjectConfig {
			deleteFiles = append(deleteFiles, project.PlatformConfigFile(fileName))
		}
		for _, f := range deleteFiles {
			path := filepath.Join(projectPath, f)
			if info, err := os.Lstat(path); err == nil {
				if info.IsDir() {
					err = os.RemoveAll(path)
				} else {
					err = os.Remove(path)
				}
			}
		}
	} else {
		// 创建目录
		if err := os.Mkdir(projectPath, fs.ModePerm); err != nil {
			return err
		}
	}

	// 读取assets内的文件
	var createFile = func(readFilePath, outFilePath string, data map[string]interface{}, perm fs.FileMode, replace ...string) error {
		// 创建 energy.json template
		if fileData, err := assets.ReadFile(nil, "", readFilePath); err != nil {
			return err
		} else {
			if data != nil {
				if fileData, err = tools.RenderTemplate(string(fileData), data); err != nil {
					return err
				}
			}
			path := filepath.Join(projectPath, outFilePath)
			if len(replace) > 0 {
				sh := strings.NewReplacer(replace...)
				fileData = []byte(sh.Replace(string(fileData)))
			}
			os.MkdirAll(filepath.Dir(path), fs.ModePerm)
			if err = ioutil.WriteFile(path, fileData, perm); err != nil {
				return err
			}
		}
		return nil
	}

	if consts.IsLinux && (consts.IsARM64 || consts.IsLoong64) {
		if err := createFile("assets/initialize/run.sh", "run.sh", nil, 0755, "\r", ""); err != nil {
			return err
		}
	}
	// 创建 energy.json template
	// 默认配置
	data := make(map[string]interface{})
	data["Name"] = c.Init.Name
	data["ProjectPath"] = filepath.ToSlash(projectPath)
	data["FrameworkPath"] = filepath.ToSlash(env.GlobalDevEnvConfig.FrameworkPath())
	data["OutputFilename"] = c.Init.Name
	data["CompanyName"] = c.Init.Name
	data["ProductName"] = c.Init.Name
	// MacOS
	data["BundleName"] = c.Init.Name
	data["BundleDisplayName"] = c.Init.Name
	data["BundleIdentifier"] = c.Init.Name
	for _, fileName := range consts.EnergyProjectConfig {
		energyTemp := fmt.Sprintf("assets/%s", fileName)
		outPath := project.PlatformConfigFile(fileName)
		if err := createFile(energyTemp, outPath, data, 0666); err != nil {
			return err
		}
	}

	// 创建 main.go
	if err := createFile(fmt.Sprintf("assets/initialize/main.go.%s", c.Init.ResLoad), "main.go", nil, 0666); err != nil {
		return err
	}

	term.Logger.Info("ENERGY Release: " + c.Init.Version)

	// 创建 go.mod
	data = make(map[string]interface{})
	data["Name"] = c.Init.Name
	data["GoVersion"] = "1.18"
	data["EnergyVersion"] = c.Init.Version
	if err := createFile("assets/initialize/go.mod.t", "go.mod", data, 0666); err != nil {
		return err
	}

	// 创建 resources/index.html
	if err := os.MkdirAll(filepath.Join(projectPath, "resources"), fs.ModePerm); err != nil {
		return err
	}
	if err := createFile("assets/initialize/index.html", filepath.Join("resources", "index.html"), nil, 0666); err != nil {
		return err
	}
	if err := createFile("assets/icon.ico", filepath.Join("resources", "icon.ico"), nil, 0666); err != nil {
		return err
	}
	if err := createFile("assets/icon.png", filepath.Join("resources", "icon.png"), nil, 0666); err != nil {
		return err
	}

	// 创建 README.md
	data = make(map[string]interface{})
	data["Name"] = c.Init.Name
	if err := createFile("assets/initialize/README.md", "README.md", data, 0666); err != nil {
		return err
	}

	// cmd
	term.Section.Println("Config Go Environment")
	cmd := toolsCommand.NewCMD()
	cmd.IsPrint = false
	cmd.Dir = projectPath
	if c.Init.IGo {
		// cmd go env -w GO111MODULE=on
		term.Logger.Info("Enable Go mod management", term.Logger.Args("command-line", "go env -w GO111MODULE=on"))
		cmd.Command("go", "env", "-w", "GO111MODULE=on")

		// cmd go env -w GOPROXY=https://goproxy.io,direct
		term.Logger.Info("Configure mod proxy", term.Logger.Args("command-line", "go env -w GOPROXY=https://goproxy.io,direct"))
		cmd.Command("go", "env", "-w", "GOPROXY=https://goproxy.io,direct")

		// cmd go mod tidy
		term.Logger.Info("Update Energy dependencies", term.Logger.Args("command-line", "go mod tidy", "version", c.Init.Version))
		cmd.Command("go", "mod", "tidy")
	}
	cmd.Close()
	return nil
}

func checkEnv(init *command.Init) {
	term.Logger.Info("Check the current environment and follow the prompts if there are any")
	// 检查Go环境
	if env.GlobalDevEnvConfig.GoCMD() == "" {
		term.Logger.Warn("Golang development environment not installed, Download-URL: ", term.Logger.Args("Download-URL", "https://golang.google.cn/dl/", "ENERGY CLI", "energy install"))
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
					term.Logger.Warn(`Current installed Go version should be greater than 1.18.`, term.Logger.Args("Version", version))
				}
			}
		}
		term.Logger.Info("Golang OK")
		init.IGo = true
	}
	if consts.IsWindows {
		// 检查nsis
		if env.GlobalDevEnvConfig.NSISCMD() == "" {
			term.Logger.Warn(`NSIS not installed, Unable to create installation package through energy command line.`)
		} else {
			term.Logger.Info(`NSIS OK`)
			init.INSIS = true
		}
	}
	if !consts.IsDarwin {
		// 检查upx
		if env.GlobalDevEnvConfig.UPXCMD() == "" {
			term.Logger.Warn(`UPX not installed`)
		} else {
			term.Logger.Info(`UPX OK`)
			init.IUPX = true
		}
	}
	// 检查当前已安装且使用的框架是否正确
	isCEF, isLCL := env.CheckCEFDir()
	if !isCEF {
		term.Logger.Warn(`CEF Framework is not installed
	Installing using ENERGY CLI`, term.Logger.Args("ENERGY CLI", "energy install"))
	}
	if !isLCL {
		term.Logger.Warn(`LibLCL is not installed
	Installing using ENERGY CLI`, term.Logger.Args("ENERGY CLI", "energy install"))
	}
	if isCEF && isLCL {
		term.Logger.Info("CEF Framework OK")
		init.IEnv = true
	}
	// 检查 node npm
	if !tools.CommandExists("npm") {
		term.Logger.Warn(`Installing node allows you to build the UI using, for example, a front-end framework (vue)`, term.Logger.Args("Download URL", "https://nodejs.org/"))
	} else {
		term.Logger.Info("Node npm OK")
		init.INpm = true
	}
}
