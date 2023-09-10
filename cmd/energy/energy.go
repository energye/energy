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
	"github.com/energye/energy/v2/cmd/internal"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/jessevdk/go-flags"
	"os"
)

type TermExit chan int

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
	termExit := make(TermExit)
	termRun(termExit)
	//<-termExit
}

func termRun(exit TermExit) {
	wd, _ := os.Getwd()
	cc := &command.Config{Wd: wd}
	parser := flags.NewParser(cc, flags.HelpFlag|flags.PassDoubleDash)
	if len(os.Args) < 2 {
		parser.WriteHelp(term.TermOut)
		//exit <- 1
		os.Exit(1)
	}
	if extraArgs, err := parser.ParseArgs(os.Args[1:]); err != nil {
		println(err.Error())
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
		}
		cmd := commands[cc.Index]
		if len(extraArgs) < 1 || extraArgs[len(extraArgs)-1] != "." {
			println(cmd.UsageLine, "\n", cmd.Long)
			//exit <- 1
			os.Exit(1)
		}
		term.Section.Println(cmd.Short)
		if err := cmd.Run(cc); err != nil {
			println(err.Error())
			//exit <- 1
			os.Exit(1)
		}
	}
}
