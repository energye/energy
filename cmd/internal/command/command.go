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

import "strings"

type OS string   // windows, macos, linux
type Arch string // 386, amd64, arm, arm64

type Config struct {
	Wd      string
	Install Install `command:"install" description:"install energy development dependency environment"`
	Package Package `command:"package" description:"energy application production and installation package"`
	Version Version `command:"version" description:"list all release version numbers of energy"`
	Env     Env     `command:"env" description:"display ENERGY_ HOME framework environment directory"`
	Init    Init    `command:"init" description:"initialize the energy application project"`
	Build   Build   `command:"build" description:"building an energy project"`
	Bindata Bindata `command:"bindata" description:"if the go version is less than 1.16, you can use bindata to embed static resources"`
	Gen     Gen     `command:"gen" description:"generate icons or syso commands"`
	Update  Update  `command:"update" description:"check and update the version"`
	Cli     Cli     `command:"cli" description:"energy cli"`
}

type Command struct {
	Run                    func(c *Config) error
	UsageLine, Short, Long string
}

type Update struct {
	Version string `short:"v" long:"version" description:"Upgrade energy to the specified version" default:""`
	Path    string `short:"p" long:"path" description:"Project path" default:""`
	WS      string `long:"ws" description:"Set this parameter when GTK2 is used on Linux" default:""`
}

type Cli struct {
	Version bool `short:"v" long:"version" description:"energy cli version"`
	Update  bool `short:"u" long:"update" description:"energy cli update"`
}

type Gen struct {
	Icon bool `long:"icon" description:"Used to generate application icons,  can convert .png to .ico Generate pixel size of: [256, 128, 64, 48, 32, 16]"`
	Syso bool `long:"syso" description:"Generate the application program xxx.syso, and when compiling the execution file, the execution file information can be written into it"`
	// 参数
	// icon, syso
	IconFilePath string `short:"p" long:"iconFilePath" description:"Icon file directory"`
	OutPath      string `short:"o" long:"outPath" description:"Save directory"`
	// syso
	Name             string `short:"n" long:"name" description:"Generate the syso file name and move it to the application name"`
	ManifestFilePath string `short:"m" long:"manifestFilePath" description:"Manifest file directory, if empty, will use the default template"`
	Arch             string `short:"a" long:"arch" description:"amd64 or i386 or arm64, if empty, the current system architecture"`
	InfoFilePath     string `short:"i" long:"infoFilePath" description:"Generate directory for syso information data files in JSON format."`
}

type Install struct {
	Path    string `short:"p" long:"path" description:"Installation directory Default current directory"`
	Version string `short:"v" long:"version" description:"Specifying a version number. x.x.x" default:"latest"`
	All     bool   `long:"all" description:"Skip select. Install All Software"`
	OS      OS     `long:"os" description:"Specify install OS: [windows, linux, macos], default current os"`
	Arch    Arch   `long:"arch" description:"Specify install ARCH: [386, amd64, arm, arm64], default current arch"`
	CEF     string `long:"cef" description:"Install system supports CEF version. options: latest, 109, 101, 87, 49" default:""`
	WS      string `long:"ws" description:"Set this parameter when GTK2 is used on Linux" default:""`
	IGolang bool   // 是否安装Golang
	ICEF    bool   // 是否安装CEF
	INSIS   bool   // 是否安装nsis
	IUPX    bool   // 是否安装upx
	I7z     bool   // 是否安装7z
}

type Package struct {
	Path        string `short:"p" long:"path" description:"Project path, default current path. Can be configured in energy.json" default:""`
	Clean       bool   `short:"c" long:"clean" description:"Clear configuration and regenerate the default configuration"`
	File        string `short:"f" long:"file" description:"Execution file name"`
	OutFileName string `short:"o" long:"outfile" description:"Installation package file name"`
	Pkgbuild    bool   `long:"pkg" description:"MacOS Using pkgbuild to create pkg development installation packages"`
}

type Env struct {
	Write string `short:"w" long:"write" description:"Set the configuration environment. set=key:value"`
	Get   string `short:"g" long:"get" description:"Get the configuration environment value. name"`
	List  bool   `short:"l" long:"list" description:"Lists the CEF Framework installed"`
	Use   string `long:"use" description:"Use the installed CEF Framework Version. ver:arch"`
}

type Version struct {
	All bool `short:"a" long:"all" description:"show all"`
}

type Init struct {
	Name    string `short:"n" long:"name" description:"Initialized project name"`
	ResLoad string `short:"r" long:"resload" description:"Resource loading method, 1: HTTP, 2: Local Load, default 1 HTTP"`
	Version string `short:"v" long:"version" description:"Specify the use of the energy release version. example: vx.x.x"`
	IGo     bool
	INSIS   bool
	IUPX    bool
	IEnv    bool
	INpm    bool
}

type Build struct {
	Path      string `short:"p" long:"path" description:"Project path, default current path. Can be configured in energy_[os].json" default:""`
	Upx       bool   `short:"u" long:"upx" description:"Set this parameter and install upx. Use upx to compress the execution file"`
	UpxFlag   string `long:"upxFlag" description:"Upx command line parameters" default:""`
	BuildArgs bool   `long:"buildargs" description:"go build [args]. Parameter tags should be in the last position, and the parameters following them will be passed as build parameters."`
	Libemfs   bool   `long:"libemfs" description:"Built in dynamic libraries to executable files, Copy liblcl to the built-in directory every compilation"`
	Out       string `short:"o" long:"out" description:"Build out file path"`
	OS        OS     `long:"os" description:"Build OS for windows | darwin | linux"`
	ARCH      Arch   `long:"arch" description:"Build ARCH for 386 | amd64 | arm | arm64"`
}

type Bindata struct {
	Debug      bool   `long:"debug" description:"Do not embed the assets, but provide the embedding API. Contents will still be loaded from disk."`
	Dev        bool   `long:"dev" description:"Similar to debug, but does not emit absolute paths. Expects a rootDir variable to already exist in the generated code's package."`
	Tags       string `long:"tags" description:"Optional set of build tags to include." default:""`
	Prefix     string `long:"prefix" description:"Optional path prefix to strip off asset names." default:""`
	Package    string `long:"pkg" description:"Package name to use in the generated code." default:"main"`
	NoMemCopy  bool   `long:"nomemcopy" description:"Use a .rodata hack to get rid of unnecessary memcopies. Refer to the documentation to see what implications this carries."`
	NoCompress bool   `long:"nocompress" description:"Assets will *not* be GZIP compressed when this flag is specified."`
	NoMetadata bool   `long:"nometadata" description:"Assets will not preserve size, mode, and modtime info."`
	FSSystem   bool   `long:"fs" description:"Whether generate instance http.FileSystem interface code."`
	Mode       uint   `long:"mode" description:"Optional file mode override for all files."`
	ModTime    int64  `long:"modtime" description:"Optional modification unix timestamp override for all files."`
	Output     string `long:"o" description:"Optional name of the output file to be generated." default:"./bindata.go"`
	Ignore     string `long:"Ignore" description:"Regex pattern to ignore." default:""`
	Paths      string `long:"paths" description:"Static resource directory, Multiple Catalogs: ./resource,./libs" default:""`
}

func (m OS) IsWindows() bool {
	return strings.ToLower(m.Value()) == "windows"
}

func (m OS) IsLinux() bool {
	return strings.ToLower(m.Value()) == "linux"
}

func (m OS) IsMacOS() bool {
	v := strings.ToLower(m.Value())
	return v == "macos" || v == "darwin"
}

func (m OS) Value() string {
	return string(m)
}

func (m Arch) Is386() bool {
	v := m.Value()
	return v == "386" || v == "i386" || v == "32"
}

func (m Arch) IsAMD64() bool {
	v := strings.ToLower(m.Value())
	return v == "amd64" || v == "64" || v == "x64"
}

func (m Arch) IsARM64() bool {
	v := strings.ToLower(m.Value())
	return v == "arm64"
}

func (m Arch) IsARM() bool {
	v := strings.ToLower(m.Value())
	return v == "arm"
}

func (m Arch) IsLoong64() bool {
	v := strings.ToLower(m.Value())
	return v == "loong64"
}

func (m Arch) Value() string {
	v := string(m)
	switch v {
	case "386", "i386", "32":
		return "32"
	case "amd64", "64", "x64":
		return "64"
	default:
		return v
	}
}
