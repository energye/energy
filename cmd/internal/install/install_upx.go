//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package install

import (
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/assets"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/consts"
	progressbar "github.com/energye/energy/v2/cmd/internal/progress-bar"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"os"
	"path/filepath"
)

func installUPX(c *command.Config) (string, func()) {
	if !c.Install.IUPX {
		return "", nil
	}
	if (consts.IsWindows && !consts.IsARM64) || (consts.IsLinux) {
		s := upxInstallPathName(c) // 安装目录
		if !tools.IsExist(s) {
			os.MkdirAll(s, 0755)
		}
		var upxName = "upx"
		if consts.IsWindows {
			upxName += ".exe"
		}
		targetFileName := filepath.Join(s, upxName) // 保存安装目录
		if targetFile, err := os.Create(targetFileName); err == nil {
			defer targetFile.Close()
			fmt.Println("extract file: ", upxName)
			bar := progressbar.NewBar(100)
			bar.SetNotice("\t")
			bar.HideRatio()
			fs, err := assets.UpxBytes()
			if err != nil {
				println("UPX Installed Error:", err.Error())
				return "", nil
			}
			stat, err := fs.Stat()
			if err != nil {
				println("UPX Installed Error:", err.Error())
				return "", nil
			}
			writeFile(fs, targetFile, stat.Size(), func(totalLength, processLength int64) {
				bar.PrintBar(int((float64(processLength) / float64(totalLength)) * 100))
			})
			bar.PrintBar(100)
			bar.PrintEnd()
			return s, func() {
				println("UPX Installed Successfully Version:", assets.UpxVersion)
			}
		} else {
			println("createWriteFile", err.Error())
			return "", nil
		}
	}
	return "", nil
}
