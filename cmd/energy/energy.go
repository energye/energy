//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package main

import (
	"fmt"
	"github.com/energye/energy/cmd"
	"github.com/jessevdk/go-flags"
	"os"
)

var commands = []*cmd.Command{
	nil,
	cmd.CmdInstall,
	cmd.CmdPackage,
	cmd.CmdVersion,
}

func main() {
	wd, _ := os.Getwd()
	cc := &cmd.CommandConfig{Wd: wd}
	parser := flags.NewParser(cc, flags.HelpFlag|flags.PassDoubleDash)
	if len(os.Args) < 2 {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
	if extraArgs, err := parser.ParseArgs(os.Args[1:]); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	} else {
		switch parser.Active.Name {
		case "install", "i":
			cc.Index = 1
		case "package", "p":
			cc.Index = 2
		case "version", "v":
			cc.Index = 3
		}
		command := commands[cc.Index]
		if len(extraArgs) < 1 || extraArgs[len(extraArgs)-1] != "." {
			fmt.Fprintf(os.Stderr, "%s\n%s", command.UsageLine, command.Long)
			os.Exit(1)
		}
		fmt.Println("Energy executing:", command.Short)
		if err := command.Run(cc); err != nil {
			fmt.Fprint(os.Stderr, err.Error()+"\n")
			os.Exit(1)
		}
	}
}
