//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 创建 energy 项目

package internal

import (
	"errors"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/env"
	"github.com/cyber-xxm/energy/v2/cmd/internal/initialize"
	"github.com/cyber-xxm/energy/v2/cmd/internal/remotecfg"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	"github.com/pterm/pterm"
	"os"
	"strings"
)

var CmdInit = &command.Command{
	UsageLine: "init -n [name]",
	Short:     "Initialized energy project",
	Long: `
	Initialize energy golang project
	-n Initialized project name
`,
}

func init() {
	CmdInit.Run = runInit
}

func runInit(c *command.Config) error {
	m := &c.Init
	if strings.TrimSpace(m.Name) == "" {
		for strings.TrimSpace(m.Name) == "" {
			print("Project Name: ")
			fmt.Scan(&m.Name)
			println()
		}
	}
	// 创建的项目名
	m.Name = strings.TrimSpace(m.Name)
	// 设置 go.mod 里使用的版本号
	if m.Version == "" {
		// 尝试全局默认版本号
		m.Version = env.GlobalDevEnvConfig.Version
	}
	if m.Version == "" {
		// 尝试使用从远程服务获取最新版本号
		latestVersion, err := remotecfg.LatestVersion()
		if err == nil {
			m.Version = fmt.Sprintf("v%v.%v.%v", latestVersion.Major, latestVersion.Minor, latestVersion.Build)
		} else {
			//最后获取失败, 要求输入版本号
			term.Logger.Error("Failed to retrieve version information from the remote service.")
			term.Logger.Error(err.Error())
			var inputVersion string
			for strings.TrimSpace(inputVersion) == "" {
				print("Enter energy release. (vx.x.x): ")
				fmt.Scan(&inputVersion)
				println()
			}
			m.Version = inputVersion
		}
	}
	// 修复版本号
	m.Version = strings.ToLower(m.Version)
	if m.Version[0] != 'v' {
		m.Version = "v" + m.Version
	}
	// 验证版本号格式, vx.x.x
	if !tools.VerifyRelease(m.Version) {
		err := fmt.Sprintf("Incorrect version format '%v'. Example: v1.0.0", m.Version)
		return errors.New(err)
	}
	if m.ResLoad == "" || (m.ResLoad != "1" && m.ResLoad != "2") {
		options := []string{"HTTP", "Local Load"}
		printer := term.DefaultInteractiveSelect.WithOnInterruptFunc(func() {
			os.Exit(1)
		}).WithOptions(options)
		printer.CheckmarkANSI()
		printer.DefaultText = "Resource Loading. Default HTTP"
		printer.Filter = false
		selectedOption, err := printer.Show()
		if err != nil {
			return err
		}
		pterm.Info.Printfln("Selected: %s", pterm.Green(selectedOption))
		if selectedOption == "" || selectedOption == "HTTP" {
			m.ResLoad = "1"
		} else if selectedOption == "Local Load" {
			m.ResLoad = "2"
		}
	}

	return initialize.InitEnergyProject(c)
}
