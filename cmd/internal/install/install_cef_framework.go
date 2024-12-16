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
	"errors"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/env"
	"github.com/energye/energy/v2/cmd/internal/remotecfg"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/pterm/pterm"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// 安装CEF和liblcl框架
//
//	 根据当前系统自动安装CEF
//	 Windows
//	   	系统版本号 > windows10: CEF > 109
//	   	系统版本号 < windows10: CEF = 109
//	 MacOS
//			CEF 最新版本
//	 Linux
//	   Gtk2: CEF 最新版本
//	   Gtk3: CEF 最新版本
func installCEFFramework(config *remotecfg.TConfig, cmdConfig *command.Config) (string, func()) {
	if !cmdConfig.Install.ICEF {
		return "", nil
	}
	pterm.Println()
	term.Section.Println("Install CEF")

	// 安装目录名称
	installPathName := config.GetFrameworkInstallPath(cmdConfig)
	if installPathName == "" {
		term.Logger.Error("Failed to obtain the frame installation directory")
		return "", nil
	}
	term.Section.Println("Install Path", installPathName)

	term.Section.Println("Start downloading CEF and Energy dependency")

	// 获取到当前安装版本
	installVersion, cefModuleName, liblclModuleName := config.GetInstallVersion(cmdConfig)
	// 判断模块名是否为空
	if installVersion == nil || liblclModuleName == "" || cefModuleName == "" {
		term.Logger.Error("The supported version is not matched, Version: " + cmdConfig.Install.Version)
		return "", nil
	}
	term.Section.Println("Install CEF-VER: " + strings.ToUpper(cefModuleName))

	// 当前安装版本的所有模块
	modules := installVersion.DependenceModule

	// 根据模块名拿到对应的模块配置
	var (
		// 当前安装版本
		cefModuleVersion    = modules.CEF[cefModuleName]
		liblclModuleVersion = modules.LCL[liblclModuleName]
		// 当前安装配置
		cefCfg  = config.ModelCEFConfig.Model(cefModuleName)
		lclCfg  = config.ModelLCLConfig.Model(liblclModuleName)
		cefItem = cefCfg.Item(cefModuleVersion)
		lclItem = lclCfg.Item(liblclModuleVersion)
		// 当前安装下载源
		cefDownloadItem = config.ModeBaseConfig.DownloadSourceItem.CEF.Item(cefItem.DownloadSource)
		lclDownloadItem = config.ModeBaseConfig.DownloadSourceItem.LCL.Item(lclItem.DownloadSource)
	)
	if cefModuleVersion == "" || liblclModuleVersion == "" {
		term.Logger.Error("CEF module " + cefModuleName + " is not configured in the current version")
		return "", nil
	}

	// CEF 版本号 109, 101 ...
	moduleVersion := cefModuleName
	if strings.Index(moduleVersion, "-") != -1 {
		moduleVersion = strings.Split(moduleVersion, "-")[1]
	}

	// 下载集合
	var downloads = make(map[string]*downloadInfo)

	// CEF 当前模块版本支持系统，如果支持返回下载地址
	libCEFOS := cefOS(cmdConfig)
	// https://cef-builds.spotifycdn.com/cef_binary_{version}_{OSARCH}_minimal.tar.bz2
	// https://www.xxx.xxx/xxx/releases/download/{version}/cef_binary_{version}_{OSARCH}_minimal.7z
	downloadCefURL := cefDownloadItem.Url
	downloadCefURL = strings.ReplaceAll(downloadCefURL, "{version}", cefModuleVersion)
	downloadCefURL = strings.ReplaceAll(downloadCefURL, "{OSARCH}", libCEFOS)
	downloads[consts.CefKey] = &downloadInfo{fileName: urlName(downloadCefURL), downloadPath: filepath.Join(cmdConfig.Install.Path, consts.FrameworkCache, urlName(downloadCefURL)), frameworkPath: installPathName, url: downloadCefURL, module: cefModuleName}

	// liblcl
	// 如果选定的cef 106，在linux会指定liblcl gtk2 版本, 其它系统和版本以默认的形式区分
	// 最后根据模块名称来确定使用哪个liblcl
	libEnergyOS := liblclOS(cmdConfig)
	useLibLCLModuleNameGTK3 := linuxUseGTK3(cmdConfig, liblclModuleName, moduleVersion)
	// https://www.xxx.xxx/xxx/releases/download/{version}/{module}.{OSARCH}.zip
	downloadEnergyURL := lclDownloadItem.Url
	downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{version}", "v"+liblclModuleVersion)
	downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{module}", useLibLCLModuleNameGTK3)
	downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{OSARCH}", libEnergyOS)
	downloads[consts.LiblclKey] = &downloadInfo{fileName: urlName(downloadEnergyURL), downloadPath: filepath.Join(cmdConfig.Install.Path, consts.FrameworkCache, urlName(downloadEnergyURL)), frameworkPath: installPathName, url: downloadEnergyURL, module: liblclModuleName}

	// 在线下载 CEF 框架二进制包
	for key, dl := range downloads {
		term.Section.Println("Download", key, ":", dl.url)
		err := tools.DownloadFile(dl.url, dl.downloadPath, env.GlobalDevEnvConfig.Proxy, nil)
		if err != nil {
			term.Logger.Error("Download [" + dl.fileName + "] " + err.Error())
			return "", nil
		}
		dl.success = err == nil
	}
	// 解压文件, 并根据配置提取文件
	term.Logger.Info("Unpack files")

	extractOSConfig := config.ModeBaseConfig.Extract.Item(string(cmdConfig.Install.OS))
	for key, di := range downloads {
		if di.success {
			if key == consts.CefKey {
				if filepath.Ext(di.downloadPath) == ".bz2" {
					processBar, err := pterm.DefaultProgressbar.WithShowCount(false).WithShowPercentage(false).WithMaxWidth(1).Start()
					if err != nil {
						term.Logger.Error(err.Error())
						return "", nil
					}
					// 解压 tar bz2
					tarSourcePath, err := tools.UnBz2ToTar(di.downloadPath, func(totalLength, processLength int64) {
						processBar.UpdateTitle(fmt.Sprintf("Unpack file %s, process: %d", key, processLength)) // Update the title of the progressbar.
					})
					processBar.Stop()
					if err != nil {
						term.Logger.Error(err.Error())
						return "", nil
					}
					// 释放文件
					if err := extractFiles(key, tarSourcePath, di, extractOSConfig.CEF); err != nil {
						term.Logger.Error(err.Error())
						return "", nil
					}
				} else {
					// 释放文件
					if err := extractFiles(key, di.downloadPath, di, extractOSConfig.CEF); err != nil {
						term.Logger.Error(err.Error())
						return "", nil
					}
				}
			} else if key == consts.LiblclKey {
				// 释放文件
				if err := extractFiles(key, di.downloadPath, di, extractOSConfig.LCL); err != nil {
					term.Logger.Error(err.Error())
					return "", nil
				}
			}
			term.Section.Println("Unpack file", key, "success")
		}
	}
	return config.GetFrameworkName(cmdConfig), func() {
		term.Logger.Info("CEF Installed Successfully", term.Logger.Args("Version", cmdConfig.Install.Version, "liblcl", liblclModuleVersion))
		if liblclModuleName == "" {
			term.Section.Println("hint: liblcl module", liblclModuleName, `is not configured in the current version`)
		}
	}
}

// 返回 CEF 支持的系统, 格式 [os][arch] 示例 windows32, windows64, macosx64 ...
func cefOS(c *command.Config) string {
	ins := c.Install
	os := command.OS(runtime.GOOS)
	arch := command.Arch(runtime.GOARCH)
	if ins.OS != "" {
		os = ins.OS
	}
	if ins.Arch != "" {
		arch = ins.Arch
	}
	if os.IsMacOS() {
		os = "macos"
		if arch.IsAMD64() {
			os += "x"
		}
	}
	if arch.Is386() {
		arch = "32"
	} else if arch.IsAMD64() {
		arch = "64"
	}
	return fmt.Sprintf("%v%v", os, arch)
}

// 返回 liblcl 在 linux 版本号大于 106 版本默认使用 GTK3
// 格式: liblcl[-ver][-GTK3]
func linuxUseGTK3(c *command.Config, liblclModuleName, moduleVersion string) string {
	ins := c.Install
	os := command.OS(runtime.GOOS)
	if ins.OS != "" {
		os = ins.OS
	}

	if os.IsLinux() && !tools.Equals(c.Install.WS, "GTK2") {
		isGT106 := false         // 标记大于106
		if moduleVersion == "" { // 空值是最新版本, 默认用GTK3
			isGT106 = true
		} else {
			if v, err := strconv.Atoi(moduleVersion); err == nil && v > 106 {
				isGT106 = true
			}
		}
		if isGT106 {
			return liblclModuleName + "-GTK3"
		}
	}
	return liblclModuleName
}

// 返回 CEF 支持的系统, 格式 [os][arch] 示例 windows32, windows64, macosx64 ...
func liblclOS(c *command.Config) string {
	var LibLCLFileNames = map[string]string{
		"windows32":    consts.Windows32,
		"windows64":    consts.Windows64,
		"windowsarm64": consts.WindowsARM64,
		"linuxarm":     consts.LinuxARM,
		"linuxarm64":   consts.LinuxARM64,
		"linux32":      consts.Linux32,
		"linux64":      consts.Linux64,
		"macosx64":     consts.MacOSX64,
		"macosarm64":   consts.MacOSARM64,
	}
	ins := c.Install
	os := command.OS(runtime.GOOS)
	arch := command.Arch(runtime.GOARCH)
	if ins.OS != "" {
		os = ins.OS
	}
	if ins.Arch != "" {
		arch = ins.Arch
	}
	if os.IsMacOS() {
		os = "macos"
		if arch.IsAMD64() {
			os += "x"
		}
	}
	if arch.Is386() {
		arch = "32"
	} else if arch.IsAMD64() {
		arch = "64"
	}
	return LibLCLFileNames[fmt.Sprintf("%v%v", os, arch)]
}

// 提取文件
func extractFiles(keyName, sourcePath string, di *downloadInfo, extractOSConfig []string) error {
	println("Extract", keyName, "sourcePath:", sourcePath, "targetPath:", di.frameworkPath)
	ext := filepath.Ext(sourcePath)
	switch ext {
	case ".tar":
		return tools.ExtractUnTar(sourcePath, di.frameworkPath, extractOSConfig...)
	case ".zip":
		return tools.ExtractUnZip(sourcePath, di.frameworkPath, false, extractOSConfig...)
	case ".7z":
		// 7z 直接解压目录内所有文件
	}
	return errors.New("not module")
}

// url文件名
func urlName(downloadUrl string) string {
	downloadUrl = downloadUrl[strings.LastIndex(downloadUrl, "/")+1:]
	return downloadUrl
}
