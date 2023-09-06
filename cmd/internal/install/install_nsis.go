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
	"path/filepath"
	"runtime"
)

func installNSIS(c *command.Config) (string, func()) {
	if !c.Install.INSIS {
		return "", nil
	}
	if consts.IsWindows && runtime.GOARCH != "arm64" {
		// 下载并安装配置NSIS
		s := c.Install.Path // 安装目录
		version := consts.NSISDownloadVersion
		fileName := fmt.Sprintf("nsis.windows.386-%s.zip", version)
		downloadUrl := fmt.Sprintf(consts.NSISDownloadURL, fileName)
		savePath := filepath.Join(s, consts.FrameworkCache, fileName)
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
			targetPath := filepath.Join(s, "NSIS")
			// 释放文件
			//zip
			ExtractUnZip(savePath, targetPath, true)
			return targetPath, func() {
				println("NSIS Installed Successfully Version:", version)
			}
		}
	}
	return "", nil
}
