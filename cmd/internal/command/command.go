//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package command

type Config struct {
	Index   int
	Wd      string
	Install Install `command:"install" description:"install energy development dependency environment"`
	Package Package `command:"package" description:"energy application production and installation package"`
	Version Version `command:"version" description:"list all release version numbers of energy"`
	Setenv  Setenv  `command:"setenv" description:"set ENERGY_ HOME framework environment"`
	Env     Env     `command:"env" description:"display ENERGY_ HOME framework environment directory"`
	Init    Init    `command:"init" description:"initialize the energy application project"`
	Build   Build   `command:"build" description:"building an energy project"`
}

type Command struct {
	Run                    func(c *Config) error
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
	Name    string `short:"n" long:"name" description:"Initialized project name"`
	ResLoad string `short:"r" long:"resload" description:"Resource loading method, 1: HTTP, 2: Local Load, default 1 HTTP"`
	IGo     bool
	INSIS   bool
	IEnv    bool
	INpm    bool
}

type Build struct {
}
