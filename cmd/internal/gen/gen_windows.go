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

package gen

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cmd/internal/assets"
	"github.com/cyber-xxm/energy/v2/cmd/internal/project"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools/winicon"
	"github.com/tc-hib/winres"
	"github.com/tc-hib/winres/version"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// GeneraICON 生成图标，如果配置的是png图标，把png转换ico
func GeneraICON(iconFilePath, outPath string) (string, error) {
	if !tools.IsExist(iconFilePath) {
		return "", errors.New("file does not exist: " + iconFilePath)
	}
	if outPath == "" {
		outPath = tools.CurrentExecuteDir()
	}
	iconExt := filepath.Ext(iconFilePath)
	if strings.ToLower(iconExt) == ".png" {
		// png => ico
		content, err := ioutil.ReadFile(iconFilePath)
		if err != nil {
			return "", err
		}
		outPath = filepath.Join(outPath, "icon.ico")
		output, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return "", err
		}
		defer output.Close()
		err = winicon.GenerateIcon(bytes.NewBuffer(content), output, []int{256, 128, 64, 48, 32, 16})
		if err != nil {
			return "", err
		}
		return outPath, nil
	}
	return iconFilePath, nil
}

// GeneraSYSO 生成 syso
func GeneraSYSO(exeName, iconFilePath, manifestFilePath, outPath, arch string, info project.Info) (string, error) {
	if exeName == "" {
		exeName = "energy-demo"
	}
	if outPath == "" {
		outPath = tools.CurrentExecuteDir()
	}
	rs := &winres.ResourceSet{}
	iconFile, err := os.Open(iconFilePath)
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
	data := make(map[string]interface{})
	data["Info"] = info

	// Manifest
	var manifestData []byte
	if manifestFilePath != "" {
		manifestData, err = ioutil.ReadFile(manifestFilePath)
	}
	if manifestData == nil || err != nil {
		manifestData, err = assets.Assets().ReadFile("assets/build/windows/app.exe.manifest")
		if err != nil {
			return "", err
		}
	}

	// 生成 manifest
	manifestData, err = tools.RenderTemplate(string(manifestData), data)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(filepath.Join(outPath, fmt.Sprintf("%s.manifest", exeName)), manifestData, fs.ModePerm)
	if err != nil {
		return "", err
	}
	xmlData, err := winres.AppManifestFromXML(manifestData)
	if err != nil {
		return "", err
	}
	rs.SetManifest(xmlData)
	// versionInfo
	versionInfo, err := assets.Assets().ReadFile("assets/build/windows/version.info.json")
	if err != nil {
		return "", err
	}
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
	var defaultArch = runtime.GOARCH
	if arch != "" {
		defaultArch = arch
	}
	targetFile := filepath.Join(outPath, fmt.Sprintf("%s-%s_%v.syso", exeName, runtime.GOOS, defaultArch))
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
	targetArch, supported := archs[defaultArch]
	if !supported {
		return targetFile, fmt.Errorf("arch '%s' not supported", defaultArch)
	}
	err = rs.WriteObject(fout, targetArch)
	if err != nil {
		return targetFile, err
	}
	return targetFile, nil
}
