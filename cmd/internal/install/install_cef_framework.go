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
//	   Gtk2: CEF = 106
//	   Gtk3: CEF 最新版本
func installCEFFramework(config *remotecfg.TConfig, c *command.Config) (string, func()) {
	if !c.Install.ICEF {
		return "", nil
	}
	pterm.Println()
	term.Section.Println("Install CEF")

	extractOSConfig := config.ModeBaseConfig.Extract.Item(string(c.Install.OS))
	latestVersion := config.LatestVersion //最新版本号
	versionList, err := remotecfg.VersionUpgradeList()
	if err != nil {
		term.Logger.Error(err.Error())
		return "", nil
	}

	// 安装目录名称
	installPathName := cefInstallPathName(c)
	term.Section.Println("Install Path", installPathName)

	term.Section.Println("Start downloading CEF and Energy dependency")

	// 获取到当前安装版本
	var installVersion *remotecfg.TVersionsUpgrade
	if c.Install.Version == "latest" {
		// 获取最新版本号, latest=vx.x.x
		if v, ok := versionList[latestVersion.Version]; ok {
			installVersion = &v
		}
	} else {
		// 自己选择版本
		if v, ok := versionList[c.Install.Version]; ok {
			installVersion = &v
		}
	}
	term.Section.Println("Check version")
	if installVersion == nil {
		term.Logger.Error("Invalid version number " + c.Install.Version)
		return "", nil
	}
	// 找到相同版配置
	for {
		if installVersion.Identical != "" {
			if v, ok := versionList[installVersion.Identical]; ok {
				installVersion = &v
				break
			} else {
				term.Logger.Error("Incorrect version configuration. Identical: " + installVersion.Identical)
				return "", nil
			}
		} else {
			break
		}
	}
	if installVersion == nil {
		term.Logger.Error("Invalid version number " + c.Install.Version)
		return "", nil
	}
	// 当前版本 cef 和 liblcl 版本选择
	var (
		cefModuleName, liblclModuleName string
		// CEF 版本号
		cef = strings.ToLower(c.Install.CEF)
	)
	if cef == "" {
		if consts.IsWindows {
			// 判断 windows 当小于 windows 10，默认使用 CEF 109
			majorVersion, _, _ := versionNumber()
			if majorVersion < 10 {
				cef = consts.CEF109
			}
		}
	}
	// 匹配固定的几个模块名
	if cef == consts.CEF109 {
		cefModuleName = "cef-109"
		liblclModuleName = "liblcl-109"
	} else if cef == consts.CEF106 {
		cefModuleName = "cef-106"
		liblclModuleName = "liblcl-106"
	} else if cef == consts.CEF101 {
		cefModuleName = "cef-101"
		liblclModuleName = "liblcl-101"
	} else if cef == consts.CEF87 {
		cefModuleName = "cef-87"
		liblclModuleName = "liblcl-87"
	}
	term.Section.Println("Install CEF-VER: " + cef)

	// CEF 如未匹配到模块，找到当前支持CEF版本最大的版本号
	if cefModuleName == "" {
		var (
			cefDefault string
			number     int
		)
		for module, _ := range installVersion.DependenceModule.CEF {
			// module = "cef-xxx"
			if s := strings.Split(module, "-"); len(s) == 2 {
				n, _ := strconv.Atoi(s[1])
				if n >= number {
					number = n
					cefDefault = module
				}
			} else {
				// module = "cef"
				cefDefault = module
				break
			}
		}
		cefModuleName = cefDefault
		liblclModuleName = "liblcl" // 固定名前缀
	}
	// 判断模块名是否为空
	if liblclModuleName == "" || cefModuleName == "" {
		term.Logger.Error("The supported version is not matched, CEF Version: " + cef)
		return "", nil
	}
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
	libCEFOS := cefOS(c)
	// https://cef-builds.spotifycdn.com/cef_binary_{version}_{OSARCH}_minimal.tar.bz2
	// https://www.xxx.xxx/xxx/releases/download/{version}/cef_binary_{version}_{OSARCH}_minimal.7z
	downloadCefURL := cefDownloadItem.Url
	downloadCefURL = strings.ReplaceAll(downloadCefURL, "{version}", cefModuleVersion)
	downloadCefURL = strings.ReplaceAll(downloadCefURL, "{OSARCH}", libCEFOS)
	downloads[consts.CefKey] = &downloadInfo{fileName: urlName(downloadCefURL), downloadPath: filepath.Join(c.Install.Path, consts.FrameworkCache, urlName(downloadCefURL)), frameworkPath: installPathName, url: downloadCefURL, module: cefModuleName}

	// liblcl
	// 如果选定的cef 106，在linux会指定liblcl gtk2 版本, 其它系统和版本以默认的形式区分
	// 最后根据模块名称来确定使用哪个liblcl
	libEnergyOS := liblclOS(c)
	useLibLCLModuleNameGTK3 := linuxUseGTK3(c, liblclModuleName, moduleVersion)
	// https://www.xxx.xxx/xxx/releases/download/{version}/{module}.{OSARCH}.zip
	downloadEnergyURL := lclDownloadItem.Url
	downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{version}", "v"+liblclModuleVersion)
	downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{module}", useLibLCLModuleNameGTK3)
	downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{OSARCH}", libEnergyOS)
	downloads[consts.LiblclKey] = &downloadInfo{fileName: urlName(downloadEnergyURL), downloadPath: filepath.Join(c.Install.Path, consts.FrameworkCache, urlName(downloadEnergyURL)), frameworkPath: installPathName, url: downloadEnergyURL, module: liblclModuleName}

	// 在线下载框架二进制包
	for key, dl := range downloads {
		term.Section.Println("Download", key, ":", dl.url)
		err = tools.DownloadFile(dl.url, dl.downloadPath, nil)
		if err != nil {
			term.Logger.Error("Download [" + dl.fileName + "] " + err.Error())
			return "", nil
		}
		dl.success = err == nil
	}
	// 解压文件, 并根据配置提取文件
	term.Logger.Info("Unpack files")
	for key, di := range downloads {
		if di.success {
			if key == consts.CefKey {
				processBar, err := pterm.DefaultProgressbar.WithShowCount(false).WithShowPercentage(false).WithMaxWidth(1).Start()
				if err != nil {
					term.Logger.Error(err.Error())
					return "", nil
				}
				tarName, err := tools.UnBz2ToTar(di.downloadPath, func(totalLength, processLength int64) {
					processBar.UpdateTitle(fmt.Sprintf("Unpack file %s, process: %d", key, processLength)) // Update the title of the progressbar.
				})
				processBar.Stop()
				if err != nil {
					term.Logger.Error(err.Error())
					return "", nil
				}
				if err := extractFiles(key, tarName, di, extractOSConfig.CEF); err != nil {
					term.Logger.Error(err.Error())
					return "", nil
				}
			} else if key == consts.LiblclKey {
				if err := extractFiles(key, di.downloadPath, di, extractOSConfig.LCL); err != nil {
					term.Logger.Error(err.Error())
					return "", nil
				}
			}
			term.Section.Println("Unpack file", key, "success")
		}
	}
	return installPathName, func() {
		term.Logger.Info("CEF Installed Successfully", term.Logger.Args("Version", c.Install.Version, "liblcl", liblclModuleVersion))
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
	if keyName == consts.CefKey {
		//tar
		return tools.ExtractUnTar(sourcePath, di.frameworkPath, extractOSConfig...)
	} else if keyName == consts.LiblclKey {
		//zip
		return tools.ExtractUnZip(sourcePath, di.frameworkPath, false, extractOSConfig...)
	}
	return errors.New("not module")
}

// url文件名
func urlName(downloadUrl string) string {
	downloadUrl = downloadUrl[strings.LastIndex(downloadUrl, "/")+1:]
	return downloadUrl
}
