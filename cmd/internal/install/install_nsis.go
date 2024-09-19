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
	"github.com/energye/energy/v2/cmd/internal/remotecfg"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/pterm/pterm"
	"path/filepath"
)

func installNSIS(config *remotecfg.TConfig, c *command.Config) (string, func()) {
	if !c.Install.INSIS {
		return "", nil
	}
	pterm.Println()
	term.Section.Println("Install NSIS")
	// 下载并安装配置NSIS
	s := nsisInstallPathName(c) // 安装目录
	downloadItem := config.ModeBaseConfig.DownloadSourceItem.NSIS.Item(0)
	version := downloadItem.Version
	fileName := fmt.Sprintf("nsis.windows.386-%s.zip", version)
	downloadUrl := fmt.Sprintf(downloadItem.Url, fileName)
	savePath := filepath.Join(c.Install.Path, consts.FrameworkCache, fileName) // 下载保存目录
	var err error
	if !tools.IsExist(savePath) {
		term.Logger.Info("NSIS Download URL: " + downloadUrl)
		term.Logger.Info("NSIS Save Path: " + savePath)
		err = tools.DownloadFile(downloadUrl, savePath, nil)
		if err != nil {
			term.Logger.Error("Download [" + fileName + "] failed: " + err.Error())
		} else {
			term.Logger.Info("Download ["+fileName+"]", term.Logger.Args(fileName, "success"))
		}
	}
	if err == nil {
		// 安装目录
		targetPath := s
		// 释放文件
		//zip
		if err = tools.ExtractUnZip(savePath, targetPath, true); err != nil {
			term.Logger.Error(err.Error())
			return "", nil
		}
		// 安装nsis7z插件
		installNSIS7z(config, c)
		return targetPath, func() {
			term.Logger.Info("NSIS Installed Successfully", term.Logger.Args("Version", version))
		}
	}
	return "", nil
}
