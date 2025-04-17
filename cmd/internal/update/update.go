//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package update

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
	cmd "github.com/cyber-xxm/energy/v2/cmd/internal/tools/cmd"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Update 更新版本
func Update(c *command.Config) (err error) {
	// 在当前目录使用了 energy 模块, 检查它，并更新到指定版本或最新版本
	update := &c.Update
	// 当前执行目录
	if update.Path == "" {
		update.Path = c.Wd
	}
	// 指定了版本号
	if update.Version != "" {
		if update.Version[0] != 'v' {
			update.Version = "v" + update.Version
		}
		if !tools.VerifyRelease(update.Version) {
			err := fmt.Sprintf("Incorrect version format '%v'. Example: v1.0.0", update.Version)
			return errors.New(err)
		}
	} else {
		// 从远程服务获取最新版本号
		// 尝试使用从远程服务获取最新版本号
		latestVersion, err := remotecfg.LatestVersion()
		if err == nil {
			update.Version = fmt.Sprintf("v%v.%v.%v", latestVersion.Major, latestVersion.Minor, latestVersion.Build)
		} else {
			// 从远程服务获取版本号失败
			return err
		}
	}
	// 检查当前执行目录使用的 energy 版本号, 同时也验证了是否 energy 项目
	currentVersion := GetCurrentModVersion(update)
	checkVersionOK := tools.VerifyRelease(currentVersion)
	// 判断结果是否为版本号 vx.x.x
	// 如果版本号验证失败，说明没使用 energy
	if currentVersion != update.Version && checkVersionOK {
		// 进入这个条件，说明使用了 energy 项目，并更新版本号
		UpdateCurrentModVersion(update)
		term.Logger.Info("Update energy finish")
	}
	// 更新当前使用的 CEF Framework
	// 更新逻辑: 当前 energy version 使用的 CEF Framework 支持的最新 LibLCL
	return UpdateCurretFrameworkLibLCL(update)
}

// 更新当前 energy version 使用的 CEF Framework 支持的最新 LibLCL
func UpdateCurretFrameworkLibLCL(u *command.Update) error {
	term.Logger.Info("Start updating LibLCL")
	// 解析当前使用 CEF Framework 信息
	cefVer := env.GlobalDevEnvConfig.CEFVersion()
	cefOS := command.OS(env.GlobalDevEnvConfig.OS())
	cefARCH := command.Arch(env.GlobalDevEnvConfig.Arch())
	term.Logger.Info("Current CEF Framework", term.Logger.Args("ENERGY", u.Version, "CEF", cefVer, "CEF-OS", cefOS, "CEF-ARCH", cefARCH))
	// 从远程服务获取配置信息
	baseConfig, err := remotecfg.ModeBaseConfig()
	if err != nil {
		return err
	}
	// 从远程服务获取 LibLCL 模块配置信息
	lclConfig, err := remotecfg.ModelLCLConfig()
	if err != nil {
		return err
	}
	// 从远程服务获取升级列表
	upgraeList, err := remotecfg.VersionUpgradeList()
	if err != nil {
		return err
	}
	// 去除 v
	version := u.Version[1:]

	upgInfo := upgraeList.Item(version)
	if upgInfo == nil {
		return errors.New("Get energy version[" + u.Version + "]The upgrade information failed because the version does not exist")
	}
	// 升级模块依赖
	dependInfo := upgInfo.DependenceModule
	// LCL 模块, 使用 liblcl[-cefver], 注意: 最新支持的 CEF 在 LibLCL 没有版本号
	// 先匹配到 CEF 的版本, 返回 [-cefver]
	matchCEFVer := func() string {
		var maxVer int // 支持 CEF 最大版本号
		for name, _ := range dependInfo.CEF {
			namever := strings.Split(name, "-")
			if len(namever) == 2 {
				v, _ := strconv.Atoi(namever[1])
				if v > maxVer {
					maxVer = v
				}
			}
		}
		if maxVer == cefVer {
			// CEF 的最大版本号，返回 "", 因为支持的最新版本在 liblcl.xxx.zip 不加版本号
			return ""
		} else {
			// 其它是指定支持版本，返回 -cefver
			return fmt.Sprintf("-%v", cefVer)
		}
	}
	// liblcl[-cefver]
	lclModule := "liblcl" + matchCEFVer()
	lclVersion, ok := dependInfo.LCL[lclModule]
	if !ok {
		return errors.New("Get LibLCL version[" + u.Version + "]upgrade information failed because the version does not exist")
	}
	// LibLCL 模块配置信息
	lclModuleConfig := lclConfig.Model(lclModule)
	if lclModuleConfig == nil {
		return errors.New("Get LibLCL Module[" + lclModule + "] config failure, module does not exist")
	}
	// 获取模块里的版本配置信息
	lclModuleItem := lclModuleConfig.Item(lclVersion)

	// 组装下载地址, 使用 sourceforge(sf)
	lclDownloadSource := baseConfig.DownloadSourceItem.LCL
	// 获得模块配置的下载源
	lclDownloadItem := lclDownloadSource.Item(lclModuleItem.DownloadSource)

	// https://xxxx/xxx/xxx/xxx/{version}/{module}.{OSARCH}.zip
	libOS := common.OS(cefOS, cefARCH)
	// 要是Linux默认是GTK3，否则需要传递 --ws gtk2 指定
	libModule := common.LibLCLLinuxUseGTK3(cefOS, u.WS, lclModule, strconv.Itoa(cefVer))
	downloadURL := lclDownloadItem.Url
	downloadURL = strings.ReplaceAll(downloadURL, "{version}", "v"+lclVersion)
	downloadURL = strings.ReplaceAll(downloadURL, "{module}", libModule)
	downloadURL = strings.ReplaceAll(downloadURL, "{OSARCH}", libOS)

	term.Logger.Info("LibLCL", term.Logger.Args("Module-Name", lclModule, "Module-Version", lclVersion, "Download-Source", downloadURL))
	// 开始下载 LibLCL
	// 保存目录
	savePath := filepath.Join(env.GlobalDevEnvConfig.Root, consts.FrameworkCache, common.UrlName(downloadURL))
	term.Logger.Info("Start Download", term.Logger.Args("URL", downloadURL, "Save-Path", savePath))
	err = tools.DownloadFile(downloadURL, savePath, env.GlobalDevEnvConfig.Proxy, nil)
	if err != nil {
		return err
	}
	term.Logger.Info("Download OK")
	// 解压
	targetPath := env.GlobalDevEnvConfig.FrameworkPath()
	term.Logger.Info("Start Unpack", term.Logger.Args("File-Path", savePath, "Target-Path", targetPath))
	err = tools.ExtractFiles(savePath, targetPath, nil)
	term.Logger.Info("Unpack OK")
	if err != nil {
		return err
	}
	term.Logger.Info("Update LibLCL finish")
	return nil
}

// 获取当前 go.mod 的 energy 版本
func GetCurrentModVersion(u *command.Update) string {
	comd := exec.Command("go", "list", "-m", "-f", "{{.Version}}", "github.com/cyber-xxm/energy/v2")
	if runtime.GOOS == "windows" {
		comd.SysProcAttr = cmd.HideWindow(true)
	}
	comd.Dir = u.Path
	output, err := comd.CombinedOutput()
	if err != nil {
		return ""
	}
	result := strings.TrimSpace(string(output))
	term.Logger.Info("Get energy version by mod.", term.Logger.Args("Result", result))
	return result
}

// 更新当前 go.mod 的 energy 版本
func UpdateCurrentModVersion(u *command.Update) {
	term.Logger.Info("Update energy by mod", term.Logger.Args("version", u.Version))
	// go get
	comd := exec.Command("go", "get", "github.com/cyber-xxm/energy/v2@"+u.Version)
	if runtime.GOOS == "windows" {
		comd.SysProcAttr = cmd.HideWindow(true)
	}
	comd.Dir = u.Path
	output, err := comd.CombinedOutput()
	if err != nil {
		return
	}
	term.Logger.Info(string(output))
	term.Logger.Info("go mod tidy")
	// go mod tidy
	comd = exec.Command("go", "mod", "tidy")
	comd.Dir = u.Path
	output, err = comd.CombinedOutput()
	if err != nil {
		return
	}
	term.Logger.Info(string(output))
}
