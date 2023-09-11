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
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/initialize"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/pterm/pterm"
	"os"
	"strings"
)

var CmdInit = &command.Command{
	UsageLine: "init -n [name]",
	Short:     "Initialized energy project",
	Long: `
	-n Initialized project name
	.  Execute default command

	Initialize energy golang project
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
	m.Name = strings.TrimSpace(m.Name)

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

	return initialize.InitEnergyProject(c)
}
