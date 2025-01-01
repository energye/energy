//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package update

import (
	"errors"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/remotecfg"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	cmd "github.com/energye/energy/v2/cmd/internal/tools/cmd"
	"os/exec"
	"runtime"
	"strings"
)

// Update 更新版本
func Update(c *command.Config) (err error) {
	// 在当前目录使用了 energy 模块, 检查它，并更新到指定版本或最新版本
	update := &c.Update
	// 当前执行目录
	if update.Path == "" {
		update.Path = c.Wd
	}
	// 指定了版本号
	if update.Version != "" {
		if update.Version[0] != 'v' {
			update.Version = "v" + update.Version
		}
		if !tools.VerifyRelease(update.Version) {
			err := fmt.Sprintf("Incorrect version format '%v'. Example: v1.0.0", update.Version)
			return errors.New(err)
		}
	} else {
		// 从远程服务获取最新版本号
		// 尝试使用从远程服务获取最新版本号
		latestVersion, err := remotecfg.LatestVersion()
		if err == nil {
			update.Version = fmt.Sprintf("v%v.%v.%v", latestVersion.Major, latestVersion.Minor, latestVersion.Build)
		} else {
			// 从远程服务获取版本号失败
			return err
		}
	}
	// 检查当前执行目录使用的 energy 版本号, 同时也验证了是否 energy 项目
	currentVersion := strings.TrimSpace(GetCurrentModVersion(update))
	// 判断结果是否为版本号 vx.x.x
	// 如果版本号验证失败，说明没使用 energy
	if currentVersion != update.Version && tools.VerifyRelease(currentVersion) {
		// 使用了 energy 项目，更新版本号
		go UpdateCurrentModVersion(update)
	}
	// 更新当前使用的 CEF Framework
	// 更新逻辑: 当前 energy version 使用的 CEF Framework 支持的最新 LibLCL
	return UpdateCurretFrameworkLibLCL(update)
}

// 更新当前 energy version 使用的 CEF Framework 支持的最新 LibLCL
func UpdateCurretFrameworkLibLCL(u *command.Update) error {
	// 从远程服务获取配置信息

	return nil
}

// 获取当前 go.mod 的 energy 版本
func GetCurrentModVersion(u *command.Update) string {
	comd := exec.Command("go", "list", "-m", "-f", "{{.Version}}", "github.com/energye/energy/v2")
	if runtime.GOOS == "windows" {
		comd.SysProcAttr = cmd.HideWindow(true)
	}
	comd.Dir = u.Path
	output, err := comd.CombinedOutput()
	if err != nil {
		return ""
	}
	return string(output)
}

// 更新当前 go.mod 的 energy 版本
func UpdateCurrentModVersion(u *command.Update) {
	// go get
	comd := exec.Command("go", "get", "github.com/energye/energy/v2@"+u.Version)
	if runtime.GOOS == "windows" {
		comd.SysProcAttr = cmd.HideWindow(true)
	}
	comd.Dir = u.Path
	output, err := comd.CombinedOutput()
	if err != nil {
		return
	}
	term.Logger.Info(string(output))
	// go mod tidy
	comd = exec.Command("go", "mod", "tidy")
	comd.Dir = u.Path
	output, err = comd.CombinedOutput()
	if err != nil {
		return
	}
	term.Logger.Info(string(output))
}
