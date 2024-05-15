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
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/pterm/pterm"
	"path/filepath"
)

func installNSIS7z(c *command.Config) (string, func()) {
	pterm.Println()
	term.Section.Println("Install NSIS7za")
	// 下载并安装配置NSIS7za
	s := nsisInstallPathName(c) // 安装目录
	version := consts.NSIS7zDownloadVersion
	fileName := fmt.Sprintf("nsis7z.windows.386-%s.zip", version)
	downloadUrl := fmt.Sprintf(consts.NSIS7zDownloadURL, fileName)
	savePath := filepath.Join(c.Install.Path, consts.FrameworkCache, fileName) // 下载保存目录
	var err error
	if !tools.IsExist(savePath) {
		term.Logger.Info("NSIS7z Download URL: " + downloadUrl)
		term.Logger.Info("NSIS7z Save Path: " + savePath)
		err = DownloadFile(downloadUrl, savePath, nil)
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
		// zip
		if err = ExtractUnZip(savePath, targetPath, false); err != nil {
			term.Logger.Error(err.Error())
			return "", nil
		}
		return targetPath, func() {
			term.Logger.Info("NSIS7z Installed Successfully", term.Logger.Args("Version", version))
		}
	}
	return "", nil
}
