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
	progressbar "github.com/energye/energy/v2/cmd/internal/progress-bar"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// 下载go并配置安装
func installGolang(c *command.Config) string {
	if tools.CommandExists("go") {
		println("Golang installed")
		return ""
	}
	print("Golang not installed, do you want to install Golang? Y/n: ")
	var s string
	if strings.ToLower(c.Install.All) != "y" {
		fmt.Scanln(&s)
		if strings.ToLower(s) != "y" {
			println("Golang install exit")
			return ""
		}
	}
	s = c.Install.Path // 安装目录
	exts := map[string]string{
		"darwin":  "tar.gz",
		"linux":   "tar.gz",
		"windows": "zip",
	}
	// 开始下载并安装Go开发环境
	version := command.GolangDefaultVersion
	gos := runtime.GOOS
	arch := runtime.GOARCH
	ext := exts[gos]
	if !tools.IsExist(s) {
		println("Directory does not exist. Creating directory.", s)
		if err := os.MkdirAll(s, fs.ModePerm); err != nil {
			println("Failed to create goroot directory", err.Error())
			return ""
		}
	}
	fileName := fmt.Sprintf("go%s.%s-%s.%s", version, gos, arch, ext)
	downloadUrl := fmt.Sprintf(command.GolangDownloadURL, fileName)
	savePath := filepath.Join(s, command.FrameworkCache, fileName)
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
		// 使用 go 名字做为 go 安装目录
		targetPath := filepath.Join(s, "go")
		// 释放文件
		if command.IsWindows {
			//zip
			ExtractUnZip(savePath, targetPath, true)
		} else {
			//tar
			ExtractUnTar(savePath, targetPath)
		}
		return targetPath
	}
	return ""
}
