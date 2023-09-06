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
	progressbar "github.com/energye/energy/v2/cmd/internal/progress-bar"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
)

// 下载go并配置安装
func installGolang(c *command.Config) (string, func()) {
	if !c.Install.IGolang {
		return "", nil
	}
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
		println("Directory does not exist. Creating directory.", s)
		if err := os.MkdirAll(s, fs.ModePerm); err != nil {
			println("Failed to create goroot directory", err.Error())
			return "", nil
		}
	}
	fileName := fmt.Sprintf("go%s.%s-%s.%s", version, gos, arch, ext)
	downloadUrl := fmt.Sprintf(consts.GolangDownloadURL, fileName)
	savePath := filepath.Join(c.Install.Path, consts.FrameworkCache, fileName) // 下载保存目录
	var err error
	println("Golang Download URL:", downloadUrl)
	println("Golang Save Path:", savePath)
	if !tools.IsExist(savePath) {
		// 已经存在不再下载
		bar := progressbar.NewBar(100)
		bar.SetNotice("\t")
		bar.HideRatio()
		err = downloadFile(downloadUrl, savePath, func(totalLength, processLength int64) {
			bar.PrintBar(int((float64(processLength) / float64(totalLength)) * 100))
		})
		if err != nil {
			bar.PrintEnd("Download [" + fileName + "] failed: " + err.Error())
		} else {
			bar.PrintEnd("Download [" + fileName + "] success")
		}
	}
	if err == nil {
		// 安装目录
		targetPath := s
		// 释放文件
		if consts.IsWindows {
			//zip
			if err = ExtractUnZip(savePath, targetPath, true); err != nil {
				println(err.Error())
				return "", nil
			}
		} else {
			//tar
			if err = ExtractUnTar(savePath, targetPath); err != nil {
				println(err.Error())
				return "", nil
			}
		}
		return targetPath, func() {
			println("Golang Installed Successfully Version:", version)
		}
	}
	return "", nil
}
