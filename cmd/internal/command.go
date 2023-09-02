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
	Env     Env     `command:"env"`
	Init    Init    `command:"init"`
	Build   Build   `command:"build"`
}

type Command struct {
	Run                    func(c *CommandConfig) error
	UsageLine, Short, Long string
}

type Install struct {
	Path     string `short:"p" long:"path" description:"Installation directory Default current directory"`
	Version  string `short:"v" long:"version" description:"Specifying a version number"`
	Name     string `short:"n" long:"name" description:"Name of the frame after installation" default:"EnergyFramework"`
	Download string `short:"d" long:"download" description:"Download Source, 0:gitee or 1:github, Default empty" default:""`
	CEF      string `short:"c" long:"cef" description:"Install system supports CEF version, provide 4 options, default empty. default, windows7, gtk2, flash" default:""`
}

type Package struct {
	Path string `short:"p" long:"path" description:"project path"`
}

type Env struct {
}

type Setenv struct {
	Path string `short:"p" long:"path" description:"Energy framework dir"`
}

type Version struct {
	All bool `short:"a" long:"all" description:"show all"`
}

type Init struct {
	Name string `short:"n" long:"name" description:"Initialized project name"`
}

type Build struct {
}
