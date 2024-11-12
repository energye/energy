//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//

package install

import (
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/env"
	"github.com/energye/energy/v2/cmd/internal/remotecfg"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/pterm/pterm"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
)

type downloadInfo struct {
	fileName      string
	frameworkPath string
	downloadPath  string
	url           string
	success       bool
	isSupport     bool
	module        string
}

type softEnf struct {
	name      string
	desc      string
	installed bool
	yes       func()
}

func Install(cliConfig *command.Config) error {
	config, err := remotecfg.BaseConfig(cliConfig.EnergyCfg)
	if err != nil {
		return err
	}
	// 设置默认参数
	defaultInstallConfig(cliConfig)
	// 检查环境
	willInstall := checkInstallEnv(cliConfig)
	var (
		goRoot                      string
		goSuccessCallback           func()
		cefFrameworkRoot            string
		cefFrameworkSuccessCallback func()
		nsisRoot                    string
		nsisSuccessCallback         func()
		upxRoot                     string
		upxSuccessCallback          func()
		z7zRoot                     string
		z7zSuccessCallback          func()
	)
	if len(willInstall) > 0 {
		// energy 依赖的软件框架
		tableData := pterm.TableData{
			{"Software", "Installed", "Support Platform Description"},
		}
		var options []string
		for _, se := range willInstall {
			var installed string
			if se.installed {
				installed = "Yes"
			} else {
				installed = "No"
				options = append(options, se.name)
				//se.yes()
			}
			var data = []string{se.name, installed, se.desc}
			tableData = append(tableData, data)
		}
		term.Section.Println("Energy Development Environment Framework Dependencies")
		err := pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("-").WithBoxed().WithData(tableData).Render()
		if err != nil {
			return err
		}
		// 选择
		if len(options) > 0 {
			printer := term.DefaultInteractiveMultiselect.WithOnInterruptFunc(func() {
				os.Exit(1)
			}).WithOptions(options)
			printer.CheckmarkANSI()
			printer.DefaultText = "Optional Installation"
			printer.Filter = false
			selectedOptions, err := printer.Show()
			if err != nil {
				return err
			}
			term.Section.Printfln("Selected : %s", pterm.Green(selectedOptions))
			for _, option := range selectedOptions {
				for _, wi := range willInstall {
					if option == wi.name {
						wi.yes()
						break
					}
				}
			}
			if len(selectedOptions) > 0 {
				// 初始配置和安装目录
				if err := initInstall(cliConfig); err != nil {
					return err
				}
			}
		}
	}
	// 检查保存框架二进制文件缓存目录是否存在
	saveCachePath := filepath.Join(cliConfig.Install.Path, consts.FrameworkCache)
	if !tools.IsExist(saveCachePath) {
		term.Section.Println("Creating directory.", saveCachePath)
		if err := os.MkdirAll(saveCachePath, fs.ModePerm); err != nil {
			return err
		}
	}
	// 安装Go开发环境
	goRoot, goSuccessCallback = installGolang(config, cliConfig)
	// 设置 go 环境变量
	if goRoot != "" {
		env.SetGoEnv(goRoot)
	}

	// 安装CEF二进制框架
	cefFrameworkRoot, cefFrameworkSuccessCallback = installCEFFramework(config, cliConfig)
	// 设置 energy cef 环境变量
	if cefFrameworkRoot != "" && cliConfig.Install.IsSame {
		env.SetEnergyHomeEnv(cefFrameworkRoot)
	}

	// 安装nsis安装包制作工具, 仅windows - amd64
	nsisRoot, nsisSuccessCallback = installNSIS(config, cliConfig)
	// 设置nsis环境变量
	if nsisRoot != "" {
		env.SetNSISEnv(nsisRoot)
	}

	// 安装upx, 内置, 仅windows, linux
	upxRoot, upxSuccessCallback = installUPX(config, cliConfig)
	// 设置upx环境变量
	if upxRoot != "" {
		env.SetUPXEnv(upxRoot)
	}

	// 安装7za
	z7zRoot, z7zSuccessCallback = install7z(config, cliConfig)
	// 设置7za环境变量
	if z7zRoot != "" {
		env.Set7zaEnv(z7zRoot)
	}
	env.SourceEnvFiles()

	// success 输出
	if nsisSuccessCallback != nil || goSuccessCallback != nil || upxSuccessCallback != nil || cefFrameworkSuccessCallback != nil || z7zSuccessCallback != nil {
		term.BoxPrintln("Hint: Reopen the cmd window for command to take effect.")
	}

	if nsisSuccessCallback != nil {
		nsisSuccessCallback()
	}
	if cefFrameworkSuccessCallback != nil {
		cefFrameworkSuccessCallback()
	}
	if goSuccessCallback != nil {
		goSuccessCallback()
	}
	if upxSuccessCallback != nil {
		upxSuccessCallback()
	}
	if z7zSuccessCallback != nil {
		z7zSuccessCallback()
	}

	//cfg := env.DevEnvReadUpdate(goRoot, cefFrameworkRoot, nsisRoot, upxRoot, z7zRoot)
	copyEnergyCMD(goRoot)

	return nil
}

// 复制energy命令工具到GOROOT/bin目录
func copyEnergyCMD(goRoot string) {
	term.Logger.Info("Copy energy command-line to GOROOT/bin")
	if goRoot == "" {
		goRoot = os.Getenv("GOROOT")
		if goRoot == "" {
			term.Logger.Warn("Install Copy energy command-line, GOROOT not found, Incorrect configuration or please restart term")
			return
		}
	}
	exe, err := os.Executable()
	if err != nil {
		term.Logger.Error(err.Error())
		return
	}
	energyName := "energy"
	if consts.IsWindows {
		energyName += ".exe"
	}
	energyBin := filepath.Join(goRoot, "bin", energyName)
	if filepath.ToSlash(exe) == filepath.ToSlash(energyBin) {
		term.Logger.Info("Current energy cli")
		return
	}
	if tools.IsExist(energyBin) {
		os.Remove(energyBin)
	}
	src, err := os.Open(exe)
	if err != nil {
		term.Logger.Error(err.Error())
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(energyBin, os.O_CREATE|os.O_WRONLY, fs.ModePerm)
	if err != nil {
		term.Logger.Error(err.Error())
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		term.Logger.Error(err.Error())
		return
	}
}

func cefInstallPathName(c *command.Config) string {
	if c.Install.IsSame {
		return filepath.Join(c.Install.Path, consts.ENERGY, c.Install.Name)
	} else {
		return filepath.Join(c.Install.Path, consts.ENERGY, fmt.Sprintf("%s_%s%s", c.Install.Name, c.Install.OS, c.Install.Arch))
	}
}

func goInstallPathName(c *command.Config) string {
	return filepath.Join(c.Install.Path, consts.ENERGY, "go")
}

func nsisInstallPathName(c *command.Config) string {
	return filepath.Join(c.Install.Path, consts.ENERGY, "nsis")
}

func upxInstallPathName(c *command.Config) string {
	return filepath.Join(c.Install.Path, consts.ENERGY, "upx")
}

func z7zInstallPathName(c *command.Config) string {
	return filepath.Join(c.Install.Path, consts.ENERGY, "7za")
}

func nsisCanInstall() bool {
	return consts.IsWindows && !consts.IsARM64
}

func upxCanInstall() bool {
	return (consts.IsWindows && !consts.IsARM64) || consts.IsLinux || consts.IsDarwin
}

func z7zCanInstall() bool {
	return consts.IsWindows && !consts.IsARM64
}

// 检查当前环境
//
//	golang, nsis, cef, upx
//	golang: all os
//	nsis: windows
//	cef: all os
//	upx: windows amd64, 386, linux amd64, arm64
func checkInstallEnv(c *command.Config) (result []*softEnf) {
	result = make([]*softEnf, 0)
	var check = func(chkInstall func() (string, bool), name string, yes func()) {
		desc, ok := chkInstall()
		result = append(result, &softEnf{name: name, desc: desc, installed: ok, yes: yes})
	}

	// go
	check(func() (string, bool) {
		return "All", tools.CommandExists("go")
	}, "Golang", func() {
		c.Install.IGolang = true //yes callback
	})

	// cef
	var cefName = "CEF Framework"
	if !c.Install.IsSame {
		cefName = fmt.Sprintf("CEF Framework %s%s", c.Install.OS, c.Install.Arch)
	}
	check(func() (string, bool) {
		if c.Install.IsSame {
			// 检查环境变量是否配置
			return "All", tools.CheckCEFDir()
		}
		// 非当前系统架构时检查一下目标安装路径是否已经存在
		var lib = func() string {
			if c.Install.OS.IsWindows() {
				return "libcef.dll"
			} else if c.Install.OS.IsLinux() {
				return "libcef.so"
			} else if c.Install.OS.IsDarwin() {
				return "cef_sandbox.a"
			}
			return ""
		}()
		s := filepath.Join(cefInstallPathName(c), lib)
		return "All", tools.IsExist(s)
	}, cefName, func() {
		c.Install.ICEF = true //yes callback
	})
	if nsisCanInstall() {
		// nsis
		check(func() (string, bool) {
			if nsisCanInstall() {
				return "Windows AMD", tools.CommandExists("makensis")
			} else {
				return "Non Windows AMD skipping NSIS.", true
			}
		}, "NSIS", func() {
			c.Install.INSIS = true //yes callback
		})
	}
	if upxCanInstall() {
		// upx
		check(func() (string, bool) {
			if upxCanInstall() {
				if consts.IsDarwin {
					if tools.CommandExists("upx") {
						return "All", true
					}
					return "Install: brew install upx", true
				}
				return "All", tools.CommandExists("upx")
			} else {
				return "Unsupported platform UPX.", true
			}
		}, "UPX", func() {
			c.Install.IUPX = true //yes callback
		})
	}
	if z7zCanInstall() {
		// 7za
		check(func() (string, bool) {
			if z7zCanInstall() {
				return "Windows AMD", tools.CommandExists("7za")
			} else {
				return "Non Windows skipping UPX.", true
			}
		}, "7za", func() {
			c.Install.I7za = true //yes callback
		})
	}
	if consts.IsLinux {
		// gtk2
		// gtk3
		// dpkg
	}
	if consts.IsDarwin {
		//
	}
	return
}

func defaultInstallConfig(c *command.Config) {
	if c.Install.Path == "" {
		// current dir
		c.Install.Path = c.Wd
	}
	if c.Install.Version == "" {
		// latest
		c.Install.Version = "latest"
	}
	if c.Install.OS == "" {
		c.Install.OS = command.OS(runtime.GOOS)
	}
	if c.Install.Arch == "" {
		c.Install.Arch = command.Arch(runtime.GOARCH)
	}
	if string(c.Install.OS) == runtime.GOOS && string(c.Install.Arch) == runtime.GOARCH {
		c.Install.IsSame = true
	}
}

func initInstall(c *command.Config) (err error) {
	// 创建安装目录
	err = os.MkdirAll(c.Install.Path, fs.ModePerm) // framework root
	if err != nil {
		return
	}
	if c.Install.ICEF {
		err = os.MkdirAll(cefInstallPathName(c), fs.ModePerm) //cef
		if err != nil {
			return
		}
	}
	if c.Install.IGolang {
		err = os.MkdirAll(goInstallPathName(c), fs.ModePerm) // go
		if err != nil {
			return
		}
	}
	if c.Install.INSIS {
		if nsisCanInstall() {
			err = os.MkdirAll(nsisInstallPathName(c), fs.ModePerm) // nsis
			if err != nil {
				return
			}
		}
	}
	if c.Install.IUPX {
		if upxCanInstall() {
			err = os.MkdirAll(upxInstallPathName(c), fs.ModePerm) //upx
			if err != nil {
				return
			}
		}
	}
	if c.Install.I7za {
		if z7zCanInstall() {
			err = os.MkdirAll(z7zInstallPathName(c), fs.ModePerm) //upx
			if err != nil {
				return
			}
		}
	}
	// framework download cache
	err = os.MkdirAll(filepath.Join(c.Install.Path, consts.FrameworkCache), fs.ModePerm)
	return err
}
