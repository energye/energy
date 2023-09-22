//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package main

import (
	"encoding/json"
	"github.com/energye/energy/v2/cmd/internal"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/energye/golcl/energy/homedir"
	"github.com/jessevdk/go-flags"
	"os"
	"path/filepath"
)

var commands = []*command.Command{
	nil,
	internal.CmdInstall,
	internal.CmdPackage,
	internal.CmdVersion,
	internal.CmdSetenv,
	internal.CmdEnv,
	internal.CmdInit,
	internal.CmdBuild,
}

func main() {
	term.GoENERGY()
	termRun()
}

func termRun() {
	wd := tools.CurrentExecuteDir()
	cc := &command.Config{Wd: wd}
	parser := flags.NewParser(cc, flags.HelpFlag|flags.PassDoubleDash)
	if len(os.Args) < 2 {
		parser.WriteHelp(term.TermOut)
		//exit <- 1
		os.Exit(1)
	}
	if extraArgs, err := parser.ParseArgs(os.Args[1:]); err != nil {
		if extraArgs[0] == "-v" || extraArgs[0] == "v" {
			term.Section.Println(" ", term.CliVersion)
		} else {
			println(err.Error())
		}
		//exit <- 1
		os.Exit(1)
	} else {
		switch parser.Active.Name {
		case "install":
			cc.Index = 1
		case "package":
			cc.Index = 2
		case "version":
			cc.Index = 3
		case "setenv":
			cc.Index = 4
		case "env":
			cc.Index = 5
		case "init":
			cc.Index = 6
		case "build":
			cc.Index = 7
		case "v":
			term.Section.Println(" ", term.CliVersion)
			return
		}
		cmd := commands[cc.Index]
		if len(extraArgs) < 1 || extraArgs[len(extraArgs)-1] != "." {
			term.Section.Println(cmd.UsageLine, "\n", cmd.Long)
			//exit <- 1
			os.Exit(1)
		}
		term.Section.Println(cmd.Short)
		readConfig(cc)
		if err := cmd.Run(cc); err != nil {
			term.Section.Println(err.Error())
			//exit <- 1
			os.Exit(1)
		}
	}
}

func readConfig(c *command.Config) {
	home, err := homedir.Dir()
	if err != nil {
		term.Section.Println(err.Error())
		return
	}
	energyDir := filepath.Join(home, ".energy")
	if !tools.IsExist(energyDir) {
		err = os.MkdirAll(energyDir, os.ModePerm)
		if err != nil {
			term.Section.Println(err.Error())
			return
		}
	}
	config := filepath.Join(energyDir, "energy.json")
	if !tools.IsExist(config) {
		cfg := command.EnergyConfig{
			Source: command.DownloadSource{
				Golang: consts.GolangDownloadSource,
				CEF:    "",
			},
		}
		cfgJSON, err := json.MarshalIndent(&cfg, "", "\t")
		if err != nil {
			term.Section.Println(err.Error())
			return
		}
		if err := os.WriteFile(config, cfgJSON, 0644); err != nil {
			term.Section.Println(err.Error())
			return
		}
		c.EnergyCfg = cfg
	} else {
		cfgJSON, err := os.ReadFile(config)
		if err != nil {
			term.Section.Println(err.Error())
			return
		}
		cfg := command.EnergyConfig{}
		err = json.Unmarshal(cfgJSON, &cfg)
		if err != nil {
			term.Section.Println(err.Error())
			return
		}
		c.EnergyCfg = cfg
	}
}
