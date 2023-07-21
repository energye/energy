//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package internal

type CommandConfig struct {
	Index   int
	Wd      string
	Install Install `command:"install"`
	Package Package `command:"package"`
	Version Version `command:"version"`
	Setenv  Setenv  `command:"setenv"`
}

type Install struct {
	Path     string `short:"p" long:"path" description:"Installation directory Default current directory"`
	Version  string `short:"v" long:"version" description:"Specifying a version number"`
	Name     string `short:"n" long:"name" description:"Name of the frame after installation" default:"EnergyFramework"`
	Download string `short:"d" long:"download" description:"Download Source, gitee or github, Default gitee" default:"gitee"`
}

type Package struct {
	Path string `short:"p" long:"path" description:"Package directory"`
	Mode string `short:"m" long:"mode" description:"Use mode to set online or offline, offline by default." default:"offline"`
	Out  string `short:"o" long:"out" description:"Output directory" default:"EnergyInstallPkg"`
}

type Setenv struct {
	Path string `short:"p" long:"path" description:"Energy framework dir"`
}

type Version struct {
	All bool `short:"a" long:"all" description:"show all"`
}

type Command struct {
	Run                    func(c *CommandConfig) error
	UsageLine, Short, Long string
}
