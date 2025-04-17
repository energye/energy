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
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/consts"
	"github.com/cyber-xxm/energy/v2/cmd/internal/env"
	"github.com/cyber-xxm/energy/v2/cmd/internal/remotecfg"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
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
	module        string
}

type softEnf struct {
	name      string
	desc      string
	installed bool
	yes       func()
}

func Install(cmdConfig *command.Config) error {
	rmtConfig, err := remotecfg.BaseConfig()
	if err != nil {
		return err
	}
	// 设置默认参数
	defaultInstallConfig(cmdConfig)
	// 检查本地环境
	willInstall := checkLocalInstallEnv(rmtConfig, cmdConfig)
	// 检查本地环境当前安装版本
	_, _, _, err = rmtConfig.GetInstallVersion(cmdConfig)
	if err != nil {
		return err
	}
	var (
		goRoot                      string
		goSuccessCallback           func()
		cefFrameworkName            string
		cefFrameworkSuccessCallback func()
		nsisRoot                    string
		nsisSuccessCallback         func()
		upxRoot                     string
		upxSuccessCallback          func()
		z7zRoot                     string
		z7zSuccessCallback          func()
	)

	// 当前系统架构将安装的框架和软件包
	if len(willInstall) > 0 {
		// energy 依赖的软件框架
		tableData := pterm.TableData{
			{"Software", "Installed", "Support Platform"},
		}
		var options []string
		for _, se := range willInstall {
			var installed string
			if se.installed {
				installed = "Installed"
			} else {
				installed = "Not Installed"
				options = append(options, se.name)
			}
			var data = []string{se.name, installed, se.desc}
			tableData = append(tableData, data)
		}
		term.Section.Println("Energy Development Environment Framework Dependencies")
		err := pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("-").WithBoxed().WithData(tableData).Render()
		if err != nil {
			return err
		}
		// 选择操作
		if len(options) > 0 {
			printer := term.DefaultInteractiveMultiselect.WithOnInterruptFunc(func() {
				os.Exit(1)
			}).WithOptions(options)
			printer.CheckmarkANSI()
			printer.DefaultText = "Optional Installation"
			printer.Filter = false
			var selectedOptions []string
			var err error
			// All 跳过输入选项，直接安装当前系统架构支持的框架和所有软件包
			if cmdConfig.Install.All {
				selectedOptions = options
			} else {
				selectedOptions, err = printer.Show()
				if err != nil {
					return err
				}
				term.Section.Printfln("Selected : %s", pterm.Green(selectedOptions))
			}
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
				if err := initInstall(rmtConfig, cmdConfig); err != nil {
					return err
				}
			}
		}
	}

	// 检查目录 框架下载缓存是否存在
	saveCachePath := filepath.Join(cmdConfig.Install.Path, consts.FrameworkCache)
	if !tools.IsExist(saveCachePath) {
		term.Section.Println("Creating Download Cache Directory:", saveCachePath)
		if err := os.MkdirAll(saveCachePath, fs.ModePerm); err != nil {
			return err
		}
	}

	// 安装Go开发环境
	goRoot, goSuccessCallback, err = installGolang(rmtConfig, cmdConfig)
	if err != nil {
		return err
	}
	// 设置 go 环境变量
	if goRoot != "" {
		env.GlobalDevEnvConfig.GoRoot = goRoot
		env.SetGoEnv(goRoot)
	}

	// 安装CEF二进制框架
	cefFrameworkName, cefFrameworkSuccessCallback, err = installCEFFramework(rmtConfig, cmdConfig)
	if err != nil {
		return err
	}
	if cefFrameworkName != "" {
		// 设置 CEF 框架, 此处是框架名不是目录
		env.GlobalDevEnvConfig.Framework = cefFrameworkName
	}

	// 安装nsis安装包制作工具, 仅windows - amd64
	nsisRoot, nsisSuccessCallback = installNSIS(rmtConfig, cmdConfig)
	// 设置nsis环境变量
	if nsisRoot != "" {
		env.GlobalDevEnvConfig.NSIS = nsisRoot
	}

	// 安装upx, 内置, 仅windows, linux
	upxRoot, upxSuccessCallback = installUPX(cmdConfig)
	// 设置upx环境变量
	if upxRoot != "" {
		env.GlobalDevEnvConfig.UPX = upxRoot
	}

	// 安装7za
	z7zRoot, z7zSuccessCallback = install7z(rmtConfig, cmdConfig)
	// 设置7za环境变量
	if z7zRoot != "" {
		env.GlobalDevEnvConfig.Z7Z = z7zRoot
	}

	// success 输出
	if nsisSuccessCallback != nil || goSuccessCallback != nil || upxSuccessCallback != nil || cefFrameworkSuccessCallback != nil || z7zSuccessCallback != nil {
		term.BoxPrintln("Hint: Restart the terminal and development tools for the commands to take effect.")
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

	env.GlobalDevEnvConfig.Update()

	copyEnergyCMD(rmtConfig, cmdConfig)

	return nil
}

// 复制 energy 命令工具到 energy 目录
func copyEnergyCMD(rmtConfig *remotecfg.TConfig, cmdConfig *command.Config) {
	term.Logger.Info("Install ENERGY CLI")

	energyRoot := filepath.Join(env.GlobalDevEnvConfig.Root, consts.ENERGY)
	energyCLIDir := filepath.Join(energyRoot, "cli")
	exeDir, err := os.Executable()
	if err != nil {
		term.Logger.Error(err.Error())
		return
	}
	cliCMD := "energy"
	if consts.IsWindows {
		cliCMD += ".exe"
	}
	energyCLI := filepath.Join(energyCLIDir, cliCMD)
	var csize, tsize int64 = 0, 0
	cexe, err := os.Stat(exeDir)
	if err == nil {
		csize = cexe.Size()
	}
	texe, err := os.Stat(energyCLI)
	if err == nil {
		tsize = texe.Size()
	}
	if filepath.ToSlash(exeDir) == filepath.ToSlash(energyCLI) && csize == tsize {
		term.Logger.Info("Current ENERGY CLI, Do not replace.")
		return
	}
	if tools.IsExist(energyCLI) {
		os.Remove(energyCLI)
	}
	os.MkdirAll(energyCLIDir, 0755)
	src, err := os.Open(exeDir)
	if err != nil {
		term.Logger.Error(err.Error())
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(energyCLI, os.O_CREATE|os.O_WRONLY, fs.ModePerm)
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
	env.SetEnergyCLIEnv(energyCLIDir)
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
	return filepath.Join(c.Install.Path, consts.ENERGY, "7z")
}

func nsisCanInstall() bool {
	return consts.IsWindows
}

func upxCanInstall() bool {
	return (consts.IsWindows && !consts.IsARM64) || consts.IsLinux
}

func z7zCanInstall() bool {
	return consts.IsWindows
}

// 检查当前环境
//
//	golang, nsis, cef, upx
//	golang: all os
//	nsis: windows
//	cef: all os
//	upx: windows amd64, 386, linux amd64, arm64
func checkLocalInstallEnv(config *remotecfg.TConfig, cmdConfig *command.Config) (result []*softEnf) {
	result = make([]*softEnf, 0)
	var check = func(chkInstall func() (string, bool), name string, yes func()) {
		desc, ok := chkInstall()
		result = append(result, &softEnf{name: name, desc: desc, installed: ok, yes: yes})
	}

	// go
	check(func() (string, bool) {
		return "All", env.GlobalDevEnvConfig.GoCMD() != ""
	}, "Golang", func() {
		cmdConfig.Install.IGolang = true //yes callback
	})

	// cef
	check(func() (string, bool) {
		// 当前安装的框架检查
		cef, lcl := CheckBeingInstalledFramework(config, cmdConfig)
		return "All", cef && lcl
	}, "CEF Framework", func() {
		cmdConfig.Install.ICEF = true //yes callback
	})

	// NSIS
	if nsisCanInstall() {
		check(func() (string, bool) {
			if nsisCanInstall() {
				return "Windows", env.GlobalDevEnvConfig.NSISCMD() != ""
			} else {
				return "Non Windows skipping NSIS.", true
			}
		}, "NSIS", func() {
			cmdConfig.Install.INSIS = true //yes callback
		})
	}

	// upx
	if upxCanInstall() {
		check(func() (string, bool) {
			if upxCanInstall() {
				return "All", env.GlobalDevEnvConfig.UPXCMD() != ""
			} else {
				return "Unsupported platform UPX.", true
			}
		}, "UPX", func() {
			cmdConfig.Install.IUPX = true //yes callback
		})
	}

	// 7z
	if z7zCanInstall() {
		check(func() (string, bool) {
			if z7zCanInstall() {
				return "Windows", env.GlobalDevEnvConfig.Z7ZCMD() != ""
			} else {
				return "Non Windows skipping UPX.", true
			}
		}, "7z", func() {
			cmdConfig.Install.I7z = true //yes callback
		})
	}

	// linux
	if consts.IsLinux {
		// gtk2
		// gtk3
		// dpkg
	}

	// macos
	if consts.IsDarwin {
		//
	}
	return
}

// 检查当前正在安装的框架lib库是否正确
func CheckBeingInstalledFramework(config *remotecfg.TConfig, cmdConfig *command.Config) (isCEF, isLCL bool) {
	// 当前安装的框架检查
	var libcef, liblcl = func() (string, string) {
		if cmdConfig.Install.OS.IsWindows() {
			return "libcef.dll", "liblcl.dll"
		} else if cmdConfig.Install.OS.IsLinux() {
			return "libcef.so", "liblcl.so"
		} else if cmdConfig.Install.OS.IsMacOS() {
			return "cef_sandbox.a", "liblcl.dylib"
		}
		return "", ""
	}()
	if libcef == "" || liblcl == "" {
		return
	}
	// 检查目录是否已安装
	// 根据安装目录（按规则取名）验证
	isCEF = tools.IsExist(filepath.Join(config.GetFrameworkInstallPath(cmdConfig), libcef))
	isLCL = tools.IsExist(filepath.Join(config.GetFrameworkInstallPath(cmdConfig), liblcl))
	return
}

func defaultInstallConfig(c *command.Config) {
	if c.Install.Path == "" {
		// 安装目录, 当前目录 或 .energy 存在的目录
		if tools.IsExist(env.GlobalDevEnvConfig.Root) {
			c.Install.Path = env.GlobalDevEnvConfig.Root
		} else {
			c.Install.Path = c.Wd
		}
	} else {
		env.GlobalDevEnvConfig.Root = c.Install.Path
		env.GlobalDevEnvConfig.Update()
	}
	if c.Install.OS == "" {
		c.Install.OS = command.OS(runtime.GOOS)
	}
	if c.Install.Arch == "" {
		c.Install.Arch = command.Arch(runtime.GOARCH)
	}
}

func initInstall(rmtConfig *remotecfg.TConfig, c *command.Config) (err error) {
	// 创建安装目录
	err = os.MkdirAll(c.Install.Path, fs.ModePerm) // framework root
	if err != nil {
		return
	}
	if c.Install.ICEF {
		err = os.MkdirAll(rmtConfig.GetFrameworkInstallPath(c), fs.ModePerm) //cef
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
	if c.Install.I7z {
		if z7zCanInstall() {
			err = os.MkdirAll(z7zInstallPathName(c), fs.ModePerm) //upx
			if err != nil {
				return
			}
		}
	}
	return err
}
