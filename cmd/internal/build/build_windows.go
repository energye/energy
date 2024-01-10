//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

package build

import (
	"bytes"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/assets"
	"github.com/energye/energy/v2/cmd/internal/command"
	"github.com/energye/energy/v2/cmd/internal/project"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/energye/energy/v2/pkgs/winicon"
	toolsCommand "github.com/energye/golcl/tools/command"
	"github.com/tc-hib/winres"
	"github.com/tc-hib/winres/version"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	windowManifest    = "windows/app.exe.manifest"
	windowVersionInfo = "windows/version.info.json"
)

// 构建windows执行程序
//
//	exe生成图标
//	编译go
//	upx
func build(c *command.Config, proj *project.Project) (err error) {
	var (
		iconPath string
		syso     string
	)
	if iconPath, err = generaICON(proj); err != nil {
		return err
	}
	var delSyso = func() {
		if syso != "" {
			os.Remove(syso)
			syso = ""
		}
	}
	defer delSyso()
	if syso, err = generaSYSO(iconPath, proj); err != nil {
		return err
	}
	// go build
	cmd := toolsCommand.NewCMD()
	cmd.Dir = proj.ProjectPath
	cmd.IsPrint = false
	term.Section.Println("Building", proj.OutputFilename)
	var args = []string{"build"}
	if c.Build.Args != "" {
		// go build args
		gbargs := strings.Split(c.Build.Args, " ")
		for i := range gbargs {
			args = append(args, gbargs[i])
		}
	}
	args = append(args, "-ldflags", "-s -w -H windowsgui")
	args = append(args, "-o", proj.OutputFilename)
	cmd.Command("go", args...)
	delSyso()
	// upx
	if c.Build.Upx && tools.CommandExists("upx") {
		term.Section.Println("Upx compression")
		args = []string{"--best", "--no-color", "--no-progress", proj.OutputFilename}
		if c.Build.UpxFlag != "" {
			args = strings.Split(c.Build.UpxFlag, " ")
			args = append(args, proj.OutputFilename)
		}
		cmd.Command("upx", args...)
	}
	cmd.Close()
	if err == nil {
		term.Section.Println("Build Successfully")
	}

	return nil
}

// 生成syso图标
func generaSYSO(iconPath string, proj *project.Project) (string, error) {
	rs := &winres.ResourceSet{}
	iconFile, err := os.Open(iconPath)
	if err != nil {
		return "", err
	}
	defer iconFile.Close()
	// icon
	ico, err := winres.LoadICO(iconFile)
	if err != nil {
		return "", fmt.Errorf("couldn't load icon from icon.ico: %w", err)
	}
	err = rs.SetIcon(winres.RT_ICON, ico)
	if err != nil {
		return "", err
	}
	// Manifest
	var manifestData []byte
	if proj.Info.Manifest != "" {
		manifestData, err = ioutil.ReadFile(proj.Info.Manifest)
	}
	if manifestData == nil || err != nil {
		manifestData, err = assets.ReadFile(proj, assetsFSPath, windowManifest)
	}
	if err != nil {
		return "", err
	}
	xmlData, err := winres.AppManifestFromXML(manifestData)
	if err != nil {
		return "", err
	}
	rs.SetManifest(xmlData)
	// versionInfo
	versionInfo, err := assets.ReadFile(proj, assetsFSPath, windowVersionInfo)
	if err != nil {
		return "", err
	}
	data := make(map[string]interface{})
	data["Info"] = proj.Info
	versionInfo, err = tools.RenderTemplate(string(versionInfo), data)
	if err != nil {
		return "", err
	}
	if len(versionInfo) != 0 {
		var v version.Info
		if err := v.UnmarshalJSON(versionInfo); err != nil {
			return "", err
		}
		rs.SetVersionInfo(v)
	}
	targetFile := filepath.Join(proj.ProjectPath, fmt.Sprintf("%s-%s.syso", proj.Name, runtime.GOOS))
	fout, err := os.Create(targetFile)
	if err != nil {
		return "", err
	}
	defer fout.Close()
	archs := map[string]winres.Arch{
		"amd64": winres.ArchAMD64,
		"arm64": winres.ArchARM64,
		"386":   winres.ArchI386,
	}
	targetArch, supported := archs[runtime.GOARCH]
	if !supported {
		return targetFile, fmt.Errorf("arch '%s' not supported", runtime.GOARCH)
	}
	err = rs.WriteObject(fout, targetArch)
	if err != nil {
		return targetFile, err
	}
	return targetFile, nil
}

// 生成应用图标，如果配置的是png图标，把png转换ico
func generaICON(proj *project.Project) (string, error) {
	iconPath := proj.Info.Icon
	if !tools.IsExist(iconPath) {
		return "", fs.ErrNotExist
	}
	iconExt := filepath.Ext(iconPath)
	if strings.ToLower(iconExt) == ".png" {
		// png => ico
		content, err := ioutil.ReadFile(iconPath)
		if err != nil {
			return "", err
		}
		iconPath = filepath.Join(assets.BuildOutPath(proj), "windows", "icon.ico")
		output, err := os.OpenFile(iconPath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return "", err
		}
		defer output.Close()
		err = winicon.GenerateIcon(bytes.NewBuffer(content), output, []int{256, 128, 64, 48, 32, 16})
		if err != nil {
			return "", err
		}
	}
	return iconPath, nil
}
