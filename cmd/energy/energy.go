package main

import (
	"fmt"
	"github.com/energye/energy/cmd"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	wd, _ := os.Getwd()
	commandConfig := &cmd.CommandConfig{}
	parser := flags.NewParser(commandConfig, flags.HelpFlag|flags.PassDoubleDash)
	if len(os.Args) < 2 {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
	if ret, err := parser.ParseArgs(os.Args[1:]); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	} else {
		fmt.Println(ret)
		fmt.Println(commandConfig.Install.Path)
		fmt.Println(parser.Active.Name)
		switch parser.Active.Name {
		case "install":
			cmd.CmdInstall(wd, commandConfig.Install)
		case "package":
			cmd.CmdPackage(wd, commandConfig.Package)
		}
	}
}
