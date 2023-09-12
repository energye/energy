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
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// 下载go并配置安装
func installGolang(c *command.Config) (string, func()) {
	if !c.Install.IGolang {
		return "", nil
	}
	pterm.Println()
	term.Section.Println("Install Golang")
	s := goInstallPathName(c) // 安装目录
	exts := map[string]string{
		"darwin":  "tar.gz",
		"linux":   "tar.gz",
		"windows": "zip",
	}
	// 开始下载并安装Go开发环境
	version := consts.GolangDefaultVersion
	gos := runtime.GOOS
	arch := runtime.GOARCH
	ext := exts[gos]
	if !tools.IsExist(s) {
		term.Section.Println("Directory does not exist. Creating directory.", s)
		if err := os.MkdirAll(s, fs.ModePerm); err != nil {
			term.Section.Println("Failed to create goroot directory", err.Error())
			return "", nil
		}
	}
	fileName := fmt.Sprintf("go%s.%s-%s.%s", version, gos, arch, ext)
	savePath := filepath.Join(c.Install.Path, consts.FrameworkCache, fileName) // 下载保存目录
	var err error
	if !tools.IsExist(savePath) {
		// Go下载源, 格式只能是 [https://xxx.xxx.xx]/dl/go1.18.10.windows-arm64.zip
		downloadSource := strings.TrimSpace(c.EnergyCfg.Source.Golang)
		if downloadSource == "" {
			downloadSource = consts.GolangDownloadSource
		}
		downloadUrl := fmt.Sprintf(consts.GolangDownloadURL, downloadSource, fileName)
		term.Logger.Info("Golang Download URL: " + downloadUrl)
		term.Logger.Info("Golang Save Path: " + savePath)
		err = downloadGolang(downloadUrl, savePath, fileName, 0)
		if err != nil {
			term.Logger.Error("Download [" + fileName + "] failed: " + err.Error())
		} else {
			term.Logger.Info("Download [" + fileName + "] success")
		}
	}
	if err == nil {
		// 安装目录
		targetPath := s
		// 释放文件
		if consts.IsWindows {
			//zip
			if err = ExtractUnZip(savePath, targetPath, true); err != nil {
				term.Logger.Error(err.Error())
				return "", nil
			}
		} else {
			//tar
			if err = ExtractUnTar(savePath, targetPath); err != nil {
				term.Logger.Error(err.Error())
				return "", nil
			}
		}
		return targetPath, func() {
			term.Logger.Info("Golang Installed Successfully", term.Logger.Args("Version", version))
		}
	}
	return "", nil
}

func downloadGolang(downloadUrl, savePath, fileName string, count int) error {
	err := DownloadFile(downloadUrl, savePath, nil)
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
