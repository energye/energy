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
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/common"
	"github.com/cyber-xxm/energy/v2/cmd/internal/consts"
	"github.com/cyber-xxm/energy/v2/cmd/internal/env"
	"github.com/cyber-xxm/energy/v2/cmd/internal/remotecfg"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	"github.com/pterm/pterm"
	"path/filepath"
	"runtime"
	"strings"
	"time"
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
func installCEFFramework(config *remotecfg.TConfig, cmdConfig *command.Config) (string, func(), error) {
	if !cmdConfig.Install.ICEF {
		return "", nil, nil
	}
	pterm.Println()
	term.Section.Println("Install CEF")

	// 安装目录名称
	installPathName := config.GetFrameworkInstallPath(cmdConfig)
	if installPathName == "" {
		term.Logger.Error("Failed to obtain the frame installation directory")
		return "", nil, nil
	}
	term.Section.Println("Install Path", installPathName)

	term.Section.Println("Start downloading CEF and Energy dependency")

	// 获取到当前安装版本
	installVersion, cefModuleName, liblclModuleName, err := config.GetInstallVersion(cmdConfig)
	if err != nil {
		return "", nil, err
	}
	// 判断模块名是否为空
	if installVersion == nil || liblclModuleName == "" || cefModuleName == "" {
		return "", nil, errors.New("The supported version is not matched, Version: " + cmdConfig.Install.Version)
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
	)
	if cefModuleVersion == "" || liblclModuleVersion == "" {
		return "", nil, errors.New("CEF module " + cefModuleName + " is not configured in the current version")
	}

	// CEF 版本号 109, 101 ...
	moduleVersion := cefModuleName
	if strings.Index(moduleVersion, "-") != -1 {
		moduleVersion = strings.Split(moduleVersion, "-")[1]
	}

	// 下载集合
	var downloads = make(map[string]*downloadInfo)
	isCEF, isLCL := CheckBeingInstalledFramework(config, cmdConfig)
	if !isCEF {
		// CEF 当前模块版本支持系统
		libCEFOS := cefOS(cmdConfig)
		// https://cef-builds.spotifycdn.com/cef_binary_{version}_{OSARCH}_minimal.tar.bz2
		// https://www.xxx.xxx/xxx/releases/download/{version}/cef_binary_{version}_{OSARCH}_minimal.7z
		downloadCefURL := cefDownloadItem.Url
		downloadCefURL = strings.ReplaceAll(downloadCefURL, "{version}", cefModuleVersion)
		downloadCefURL = strings.ReplaceAll(downloadCefURL, "{OSARCH}", libCEFOS)
		downloads[consts.CefKey] = &downloadInfo{fileName: common.UrlName(downloadCefURL), downloadPath: filepath.Join(cmdConfig.Install.Path, consts.FrameworkCache, common.UrlName(downloadCefURL)), frameworkPath: installPathName, url: downloadCefURL, module: cefModuleName}
	}
	// LibLCL 动态链接库 下载地址, 在下面下载失败时自动切换下一个下载源使用
	var getLibLCLDownUrl func() string
	if !isLCL {
		// liblcl 当前模块版本支持系统
		libEnergyOS := common.LibLCLOS(cmdConfig)
		useLibLCLModuleNameGTK3 := common.LibLCLLinuxUseGTK3(cmdConfig.Install.OS, cmdConfig.Install.WS, liblclModuleName, moduleVersion)
		// https://www.xxx.xxx/xxx/releases/download/{version}/{module}.{OSARCH}.zip
		downSrcList := lclItem.DownloadSourceList[:]
		// 默认使用下载源
		useDownSrc := lclItem.DownloadSource
		// 删除当前使用下载源，以在失败情况下可以获取下一个
		removeDownSrcForListAndSetNext := func() {
			for i, v := range downSrcList {
				if v == useDownSrc {
					// 删除这个
					downSrcList = append(downSrcList[:i], downSrcList[i+1:]...)
					break
				}
			}
			// 重新设置下载源，如果下载失败
			if len(downSrcList) > 0 {
				useDownSrc = downSrcList[0]
			}
		}
		// 获取 LibLCL 下载源，返回空表示已经被全部使用过了
		getLibLCLDownUrl = func() string {
			if len(downSrcList) == 0 {
				return ""
			}
			lclDownloadItem := config.ModeBaseConfig.DownloadSourceItem.LCL.Item(useDownSrc)
			downloadURL := lclDownloadItem.Url
			downloadURL = strings.ReplaceAll(downloadURL, "{version}", "v"+liblclModuleVersion)
			downloadURL = strings.ReplaceAll(downloadURL, "{module}", useLibLCLModuleNameGTK3)
			downloadURL = strings.ReplaceAll(downloadURL, "{OSARCH}", libEnergyOS)
			// 删除 downSrcList 当前下载源，表示已经使用过
			// 将当前使用的在 下载集合里 删除掉
			removeDownSrcForListAndSetNext()
			return downloadURL
		}
		// 获取下载地址
		downloadEnergyURL := getLibLCLDownUrl()
		downloads[consts.LiblclKey] = &downloadInfo{fileName: common.UrlName(downloadEnergyURL), downloadPath: filepath.Join(cmdConfig.Install.Path, consts.FrameworkCache, common.UrlName(downloadEnergyURL)), frameworkPath: installPathName, url: downloadEnergyURL, module: liblclModuleName}
	}
	// 在线下载 CEF 框架二进制包
	var sortsKeys = []string{consts.LiblclKey, consts.CefKey}
	for i := 0; i < len(sortsKeys); i++ {
		key := sortsKeys[i]
		dl, ok := downloads[key]
		if ok {
			term.Section.Println("Download", key, ":", dl.url)
			err := tools.DownloadFile(dl.url, dl.downloadPath, env.GlobalDevEnvConfig.Proxy, nil)
			if key == consts.LiblclKey {
				if err != nil {
					// 失败尝试使用下个下载源
					if downURL := getLibLCLDownUrl(); downURL != "" {
						term.Logger.Error("Download", term.Logger.Args("ERROR", err.Error(), "Auto switch download source", downURL))
						dl.url = downURL
						i--
						continue
					} else {
						return "", nil, errors.New("Download [" + dl.fileName + "] " + err.Error())
					}
				}
			} else {
				if err != nil {
					return "", nil, errors.New("Download [" + dl.fileName + "] " + err.Error())
				}
			}
			dl.success = err == nil
		}
	}

	// 解压文件, 并根据配置提取文件
	term.Logger.Info("Unpack files")

	extractOSConfig := config.ModeBaseConfig.Extract.Item(cmdConfig.Install.OS)
	for key, di := range downloads {
		if di.success {
			if key == consts.CefKey {
				if filepath.Ext(di.downloadPath) == ".bz2" {
					processBar, err := pterm.DefaultProgressbar.WithShowCount(false).WithShowPercentage(false).WithMaxWidth(1).Start()
					if err != nil {
						return "", nil, err
					}
					// 解压 tar bz2
					beginTime := time.Now()
					tarSourcePath, err := tools.UnBz2ToTar(di.downloadPath, func(totalLength, processLength int64) {
						nowTime := time.Now()
						if nowTime.Sub(beginTime) >= time.Second { //1秒更新一次
							beginTime = nowTime
							processBar.UpdateTitle(fmt.Sprintf("Unpack file %s, process: %d", key, processLength)) // Update the title of the progressbar.
						}
					})
					processBar.Stop()
					if err != nil {
						return "", nil, err
					}
					// 释放文件
					if err := extractFiles(key, tarSourcePath, di, extractOSConfig.CEF); err != nil {
						return "", nil, err
					}
				} else {
					// 释放文件
					if err := extractFiles(key, di.downloadPath, di, extractOSConfig.CEF); err != nil {
						return "", nil, err
					}
				}
			} else if key == consts.LiblclKey {
				// 释放文件
				if err := extractFiles(key, di.downloadPath, di, extractOSConfig.LCL); err != nil {
					return "", nil, err
				}
			}
			term.Section.Println("Unpack file", key, "success")
		}
	}
	return config.GetFrameworkName(cmdConfig), func() {
		term.Logger.Info("CEF Installed Successfully", term.Logger.Args("Version", strings.ToUpper(cefModuleVersion), "LibLCL", liblclModuleVersion))
		if liblclModuleName == "" {
			term.Section.Println("Hint: LibLCL Module", liblclModuleName, `is not configured in the current version`)
		}
	}, nil
}

// 返回 CEF 支持的系统, 格式 [os][arch] 示例 windows32, windows64, macosx64 ...
func cefOS(c *command.Config) string {
	os := command.OS(runtime.GOOS)
	arch := command.Arch(runtime.GOARCH)
	if c.Install.OS != "" {
		os = c.Install.OS
	}
	if c.Install.Arch != "" {
		arch = c.Install.Arch
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

// 提取文件
func extractFiles(keyName, sourcePath string, di *downloadInfo, extractOSConfig []string) error {
	println("Extract", keyName, "sourcePath:", sourcePath, "targetPath:", di.frameworkPath)
	return tools.ExtractFiles(sourcePath, di.frameworkPath, extractOSConfig)
}
