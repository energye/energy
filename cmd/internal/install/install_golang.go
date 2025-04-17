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
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/consts"
	"github.com/cyber-xxm/energy/v2/cmd/internal/env"
	"github.com/cyber-xxm/energy/v2/cmd/internal/remotecfg"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	"github.com/pterm/pterm"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// 下载go并配置安装
func installGolang(rtmConfig *remotecfg.TConfig, cmdConfig *command.Config) (string, func(), error) {
	if !cmdConfig.Install.IGolang {
		return "", nil, nil
	}
	pterm.Println()
	term.Section.Println("Install Golang")
	installPath := goInstallPathName(cmdConfig) // 安装目录
	exts := map[string]string{
		"darwin":  "tar.gz",
		"linux":   "tar.gz",
		"windows": "zip",
	}
	golang := rtmConfig.ModeBaseConfig.DownloadSourceItem.GoLang.Item(0)
	// 开始下载并安装Go开发环境
	version := golang.Version
	gos := runtime.GOOS
	arch := runtime.GOARCH
	ext := exts[gos]
	if !tools.IsExist(installPath) {
		term.Section.Println("Creating directory.", installPath)
		if err := os.MkdirAll(installPath, fs.ModePerm); err != nil {
			term.Section.Println("Failed to create goroot directory", err.Error())
			return "", nil, err
		}
	} else {
		// 检查 golang 的 go cmd 是否存在
		gocmd := filepath.Join(installPath, "bin", "go")
		if tools.IsExist(gocmd) {
			return installPath, func() {
				term.Logger.Info("Golang has been installed")
			}, nil
		}
	}
	// 下载文件名
	fileName := fmt.Sprintf("go%s.%s-%s.%s", version, gos, arch, ext)
	// 下载保存目录
	saveFileCachPath := filepath.Join(cmdConfig.Install.Path, consts.FrameworkCache, fileName)
	var err error
	if !tools.IsExist(saveFileCachPath) {
		// Go下载源, 格式只能是 https://xxx.xxx.xx/dl/%s
		downloadUrl := fmt.Sprintf(golang.Url, fileName)
		term.Logger.Info("Golang Download URL: " + downloadUrl)
		term.Logger.Info("Golang Save Path: " + saveFileCachPath)
		err = downloadGolang(downloadUrl, saveFileCachPath, fileName, 0)
		if err != nil {
			term.Logger.Error("Download [" + fileName + "] failed: " + err.Error())
			return "", nil, err
		} else {
			term.Logger.Info("Download [" + fileName + "] success")
		}
	}
	if err == nil {
		// 安装目录
		// 释放文件
		if consts.IsWindows {
			//zip
			if err = tools.ExtractUnZip(saveFileCachPath, installPath, true); err != nil {
				return "", nil, err
			}
		} else {
			//tar
			if err = tools.ExtractUnTar(saveFileCachPath, installPath); err != nil {
				term.Logger.Error(err.Error())
				return "", nil, err
			}
		}
		return installPath, func() {
			term.Logger.Info("Golang Installed Successfully", term.Logger.Args("Version", version))
		}, nil
	}
	return "", nil, err
}

func downloadGolang(downloadUrl, savePath, fileName string, count int) error {
	err := tools.DownloadFile(downloadUrl, savePath, env.GlobalDevEnvConfig.Proxy, nil)
	if err != nil && count < 5 {
		// 失败尝试5次，每次递增一秒等待
		n := count + 1
		term.Logger.Error(err.Error())
		term.Logger.Error(fmt.Sprintf("Download failed. %d second retry", n), term.Logger.Args("count", fmt.Sprintf("%d/5", n)))
		time.Sleep(time.Second * time.Duration(n))
		return downloadGolang(downloadUrl, savePath, fileName, n)
	}
	return err
}
