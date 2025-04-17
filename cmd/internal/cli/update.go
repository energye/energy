//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cli

import (
	"github.com/cyber-xxm/energy/v2/cmd/internal/consts"
	"github.com/cyber-xxm/energy/v2/cmd/internal/env"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// OnlineUpdate 在线更新 cli
func OnlineUpdate(downloadURL string) error {
	term.Section.Println("Start downloading")
	path, err := os.Executable()
	if err != nil {
		return err
	}
	dUrl, err := url.Parse(downloadURL)
	if err != nil {
		return err
	}
	uPath := dUrl.Path
	fileName := uPath[strings.LastIndex(uPath, "/"):]
	path, _ = filepath.Split(path)
	savePath := filepath.Join(path, fileName)
	err = tools.DownloadFile(downloadURL, savePath, env.GlobalDevEnvConfig.Proxy, nil)
	if err != nil {
		return err
	}
	err = tools.ExtractUnZip(savePath, path, false)
	if err != nil {
		return err
	}
	//os.Remove(savePath)
	cliName := consts.ENERGY
	zipCliName := CliFileName()
	if consts.IsWindows {
		zipCliName += ".exe"
		cliName += ".exe"
	}

	if consts.IsWindows {
		args := []string{"/c", "del", cliName, "&", "ren", zipCliName, cliName}
		term.Section.Println("Run command:", args)
		cmd := exec.Command("cmd.exe", args...)
		cmd.Dir = path
		err = cmd.Start()
	} else {
		err = os.Remove(filepath.Join(path, cliName))
		if err != nil {
			return err
		}
		err = os.Rename(filepath.Join(path, zipCliName), filepath.Join(path, cliName))
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return nil
}
