//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package main energy command-line
package main

import (
	"github.com/energye/energy/v2/cmd/internal"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/env"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/jessevdk/go-flags"
	"os"
)

var commands = []*command.Command{
	/*0*/ nil,
	/*1*/ internal.CmdInstall,
	/*2*/ internal.CmdPackage,
	/*3*/ internal.CmdVersion,
	/*4*/ internal.CmdSetenv,
	/*5*/ internal.CmdEnv,
	/*6*/ internal.CmdInit,
	/*7*/ internal.CmdBuild,
	/*8*/ internal.CmdBindata,
	/*9*/ internal.CmdGen,
	/*10*/ internal.CmdUpgrade,
	/*11*/ internal.CmdHelp,
	/*12*/ internal.CmdCli,
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
		os.Exit(0)
	}
	if extraArgs, err := parser.ParseArgs(os.Args[1:]); err != nil {
		println(err.Error())
		return
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
		case "bindata":
			cc.Index = 8
		case "gen":
			cc.Index = 9
		case "upg":
			cc.Index = 10
		case "help":
			cc.Index = 11
		case "cli":
			cc.Index = 12
		}
		cmd := commands[cc.Index]
		// energy [cmd] help
		if len(extraArgs) > 0 {
			name := extraArgs[0]
			if name == "help" {
				term.Section.Println(cmd.UsageLine, "\n", cmd.Long)
				os.Exit(0)
			}
		}
		if cmd.Short != "" {
			term.Section.Println(cmd.Short)
		}
		readConfig(cc)
		if err := cmd.Run(cc); err != nil {
			term.Section.Println(err.Error())
			os.Exit(1)
		}
	}
}

func readConfig(c *command.Config) {
	c.EnergyCfg = *env.DevEnvReadUpdate("", "", "", "", "")
}
