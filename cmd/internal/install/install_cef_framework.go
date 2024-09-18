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
	"github.com/pterm/pterm"
	"path/filepath"
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
	latestVersion := config.LatestVersion
	versionList, err := remotecfg.VersionUpgradeList()
	if err != nil {
		term.Logger.Error(err.Error())
		return "", nil
	}

	// -c cef args value
	// default(empty), windows7, gtk2, flash
	cefVer := strings.ToLower(c.Install.CEFVer)
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
	// 相同版配置
	for {
		if installVersion.Identical != "" {
			if v, ok := versionList[installVersion.Identical]; ok {
				installVersion = &v
				break
			}
		}
	}
	if installVersion == nil {
		term.Logger.Error("Invalid version number " + c.Install.Version)
		return "", nil
	}

	// 当前版本 cef 和 liblcl 版本选择
	var (
		cefModuleName, liblclModuleName string
	)
	if consts.IsWindows {
		// 判断当前 windows 版本, 当小于windows 10，使用CEF109，CEF 109是最后一个支持windows7的版本
		majorVersion, _, _ := versionNumber()
		if majorVersion < 10 {
			cefVer = consts.Cef109
		}
	}
	// 判断是否指定了 cefVer 特定版本号版本号
	if cefVer == consts.Cef106 {
		cefModuleName = "cef-106" // CEF 106.1.1, linux gtk2
	} else if cefVer == consts.Cef109 {
		cefModuleName = "cef-109"       // CEF 109.1.18
		liblclModuleName = "liblcl-109" // liblcl 109, windows7
	} else if cefVer == consts.Cef87 {
		// cef 87 要和 liblcl 87 配对
		cefModuleName = "cef-87"       // CEF 87.1.14
		liblclModuleName = "liblcl-87" // liblcl 87, flash
	}
	// 如未指定CEF参数、或参数不正确，选择当前支持CEF版本最（新）大的版本号
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
	}
	// liblcl 在未指定版本时，它是空 ""
	if liblclModuleName == "" {
		liblclModuleName = "liblcl"
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
	moduleVersion := cefModuleName
	if strings.Index(moduleVersion, "-") != -1 {
		moduleVersion = strings.Split(moduleVersion, "-")[1]
	}
	// 下载集合
	var downloads = make(map[string]*downloadInfo)
	// 当前模块版本支持系统，如果支持返回下载地址
	libCEFOS, isSupport := cefOS(c, moduleVersion, cefItem.SupportOSArch)
	downloadCefURL := cefDownloadItem.Url
	downloadCefURL = strings.ReplaceAll(downloadCefURL, "{version}", cefModuleVersion)
	downloadCefURL = strings.ReplaceAll(downloadCefURL, "{OSARCH}", libCEFOS)
	downloads[consts.CefKey] = &downloadInfo{isSupport: isSupport, fileName: urlName(downloadCefURL), downloadPath: filepath.Join(c.Install.Path, consts.FrameworkCache, urlName(downloadCefURL)), frameworkPath: installPathName, url: downloadCefURL, module: cefModuleName}

	// liblcl
	// 如果选定的cef 106，在linux会指定liblcl gtk2 版本, 其它系统和版本以默认的形式区分
	// 最后根据模块名称来确定使用哪个liblcl
	libEnergyOS, isSupport := liblclOS(c, moduleVersion, lclItem.SupportOSArch)
	downloadEnergyURL := lclDownloadItem.Url
	downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{version}", "v"+liblclModuleVersion)
	downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{module}", liblclModuleName)
	downloadEnergyURL = strings.ReplaceAll(downloadEnergyURL, "{OSARCH}", libEnergyOS)
	downloads[consts.LiblclKey] = &downloadInfo{isSupport: isSupport, fileName: urlName(downloadEnergyURL), downloadPath: filepath.Join(c.Install.Path, consts.FrameworkCache, urlName(downloadEnergyURL)), frameworkPath: installPathName, url: downloadEnergyURL, module: liblclModuleName}

	// 在线下载框架二进制包
	for key, dl := range downloads {
		term.Section.Println("Download", key, ":", dl.url)
		if !dl.isSupport {
			term.Logger.Warn("Warn module is not built or configured [" + dl.module + "]")
			continue
		}
		err = DownloadFile(dl.url, dl.downloadPath, nil)
		if err != nil {
			term.Logger.Error("Download [" + dl.fileName + "] " + err.Error())
			return "", nil
		}
		dl.success = err == nil
	}
	// 解压文件, 并根据配置提取文件
	term.Logger.Info("Unpack files")
	for key, di := range downloads {
		if !di.isSupport {
			term.Logger.Warn("module is not built or configured [" + di.module + "]")
			continue
		}
		if di.success {
			if key == consts.CefKey {
				processBar, err := pterm.DefaultProgressbar.WithShowCount(false).WithShowPercentage(false).WithMaxWidth(1).Start()
				if err != nil {
					term.Logger.Error(err.Error())
					return "", nil
				}
				tarName, err := UnBz2ToTar(di.downloadPath, func(totalLength, processLength int64) {
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

func cefOS(c *command.Config, moduleVersion, buildSupportOSArch string) (string, bool) {
	archs := strings.Split(buildSupportOSArch, ",")
	var isSupport = func(goarch string) bool {
		for _, v := range archs {
			if goarch == v {
				return true
			}
		}
		return false
	}

	if c.Install.OS.IsWindows() { // windows arm for 64 bit, windows for 32/64 bit
		if c.Install.Arch.IsARM64() {
			return "windowsarm64", isSupport(consts.WindowsARM64)
		}
		if c.Install.Arch.Is386() {
			return "windows32", isSupport(consts.Windows32)
		}
		return "windows64", isSupport(consts.Windows64)
	} else if c.Install.OS.IsLinux() { //linux for 64 bit
		if c.Install.Arch.IsARM64() {
			if moduleVersion == consts.Cef106 {
				return "linuxarm64", isSupport(consts.LinuxARM64GTK2)
			}
			return "linuxarm64", isSupport(consts.LinuxARM64) || isSupport(consts.LinuxARM64GTK3)
		} else if c.Install.Arch.IsAMD64() {
			if moduleVersion == consts.Cef106 {
				return "linux64", isSupport(consts.Linux64GTK2)
			}
			return "linux64", isSupport(consts.Linux64) || isSupport(consts.Linux64GTK3)
		}
	} else if c.Install.OS.IsDarwin() { // macosx for 64 bit
		if c.Install.Arch.IsARM64() {
			return "macosarm64", isSupport(consts.MacOSARM64)
		} else if c.Install.Arch.IsAMD64() {
			return "macosx64", isSupport(consts.MacOSX64)
		}
	}
	//not support
	return fmt.Sprintf("%v %v", c.Install.OS, c.Install.Arch), false
}

var liblclFileNames = map[string]string{
	"windows386":     consts.Windows32,
	"windowsamd64":   consts.Windows64,
	"windowsarm64":   consts.WindowsARM64,
	"linuxarm64":     consts.LinuxARM64,
	"linuxarm64gtk2": consts.LinuxARM64GTK2,
	"linuxamd64":     consts.Linux64,
	"linuxamd64gtk2": consts.Linux64GTK2,
	"darwinamd64":    consts.MacOSX64,
	"darwinarm64":    consts.MacOSARM64,
}

func liblclName(c *command.Config, moduleVersion string) (string, bool) {
	var key string
	var isOld bool
	if c.Install.Arch.IsARM64() {
		if c.Install.OS.IsLinux() && moduleVersion == consts.Cef106 { // 只linux区别liblcl gtk2
			key = "linuxarm64gtk2"
		} else {
			key = fmt.Sprintf("%sarm64", c.Install.OS) // linux arm64, macos arm64
		}
	} else {
		if c.Install.OS.IsLinux() && moduleVersion == consts.Cef106 { // 只linux区别liblcl gtk2
			key = "linuxamd64gtk2"
		} else {
			key = fmt.Sprintf("%s%s", c.Install.OS, c.Install.Arch) // windows 386 amd64, linux amd64, macos amd64
		}
	}
	if key != "" {
		return liblclFileNames[key], isOld
	}
	return "", false
}

// 命名规则 OS+[ARCH]+BIT+[GTK2]
//
//	ARCH: 非必需, ARM 时填写, AMD为空
//	GTK2: 非必需, GTK2(Linux CEF 106) 时填写, 非Linux或GTK3时为空
func liblclOS(c *command.Config, moduleVersion, buildSupportOSArch string) (string, bool) {
	archs := strings.Split(buildSupportOSArch, ",")
	var goarch string
	if c.Install.OS.IsWindows() && c.Install.Arch.Is386() {
		goarch = "32" // windows32 = > windows386
	} else {
		goarch = string(c.Install.Arch)
	}
	noSuport := fmt.Sprintf("%v %v", c.Install.OS, goarch)
	var isSupport = func(goarch string) bool {
		for _, v := range archs {
			if goarch == v {
				return true
			}
		}
		return false
	}
	if name, isOld := liblclName(c, moduleVersion); isOld {
		if name == "" {
			return noSuport, false
		}
		return name, true
	} else {
		return name, isSupport(name)
	}
}

// 提取文件
func extractFiles(keyName, sourcePath string, di *downloadInfo, extractOSConfig []string) error {
	println("Extract", keyName, "sourcePath:", sourcePath, "targetPath:", di.frameworkPath)
	if keyName == consts.CefKey {
		//tar
		return ExtractUnTar(sourcePath, di.frameworkPath, extractOSConfig...)
	} else if keyName == consts.LiblclKey {
		//zip
		return ExtractUnZip(sourcePath, di.frameworkPath, false, extractOSConfig...)
	}
	return errors.New("not module")
}

// url文件名
func urlName(downloadUrl string) string {
	downloadUrl = downloadUrl[strings.LastIndex(downloadUrl, "/")+1:]
	return downloadUrl
}
