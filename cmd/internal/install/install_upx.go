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
	"github.com/energye/energy/v3/cmd/internal/assets"
	"github.com/energye/energy/v3/cmd/internal/command"
	"github.com/energye/energy/v3/cmd/internal/consts"
	"github.com/energye/energy/v3/cmd/internal/term"
	"github.com/energye/energy/v3/cmd/internal/tools"
	"github.com/pterm/pterm"
	"os"
	"path/filepath"
)

func installUPX(c *command.Config) (string, func()) {
	if !c.Install.IUPX {
		return "", nil
	}
	pterm.Println()
	term.Section.Println("Install UPX")
	s := upxInstallPathName(c) // 安装目录
	if !tools.IsExist(s) {
		os.MkdirAll(s, 0755)
	}
	var upxName = "upx"
	if consts.IsWindows {
		upxName += ".exe"
	}
	targetFileName := filepath.Join(s, upxName) // 保存安装目录

	if targetFile, err := os.OpenFile(targetFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm); err == nil {
		defer targetFile.Close()
		term.Section.Println("extract file: ", upxName)

		fs, err := assets.UpxBytes()
		if err != nil {
			term.Logger.Error("UPX Installed Error: " + err.Error())
			return "", nil
		}
		stat, err := fs.Stat()
		if err != nil {
			term.Logger.Error("UPX Installed Error: " + err.Error())
			return "", nil
		}

		var (
			total = 100
			count int
			cn    int
		)
		p, err := pterm.DefaultProgressbar.WithTotal(total).WithTitle("Write File " + upxName).Start()
		if err != nil {
			return "", nil
		}
		writeFile(fs, targetFile, stat.Size(), func(totalLength, processLength int64) {
			process := int((float64(processLength) / float64(totalLength)) * 100)
			if process > count {
				count = process
				p.Add(1)
				cn++
			}
		})
		if cn < total {
			p.Add(total - cn)
		}
		p.Stop()
		return s, func() {
			term.Logger.Info("UPX Installed Successfully ", term.Logger.Args("Version", assets.UpxVersion))
		}
	} else {
		term.Logger.Error("CreateWriteFile: " + err.Error())
		return "", nil
	}

}
