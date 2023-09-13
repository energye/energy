//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build linux
// +build linux

package packager

import (
	"errors"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/assets"
	"github.com/energye/energy/v2/cmd/internal/project"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	deb         = "DEBIAN"
	debControl  = deb + "/control"
	debPreinit  = deb + "/preinit"
	debPostinit = deb + "/postinit"
	debPrerm    = deb + "/prerm"
	debPostrm   = deb + "/postrm"
)

func GeneraInstaller(proj *project.Project) error {
	if !tools.CommandExists("dpkg") {
		return errors.New("failed to create application installation program. Could not find the dpkg command")
	}
	// 创建构建输出目录
	buildOutDir := assets.BuildOutPath(proj)
	buildOutDir = filepath.Join(buildOutDir, "linux")
	if !tools.IsExist(buildOutDir) {
		if err := os.MkdirAll(buildOutDir, 0755); err != nil {
			return fmt.Errorf("unable to create directory: %w", err)
		}
	}
	var err error
	if err = linux(proj); err != nil {
		return err
	}
	return nil
}

func linux(proj *project.Project) error {
	buildOutDir := assets.BuildOutPath(proj)
	buildOutDir = filepath.Join(buildOutDir, "linux")
	// app dir
	appDir := filepath.Join(buildOutDir, proj.Name)
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}
	debDir := filepath.Join(appDir, deb)
	if err := os.MkdirAll(debDir, 0755); err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}
	controlTemp := `Package: {{.CompanyName}}.{{.ProductName}}.{{.Name}}
Version: {{.ProductVersion}}
Section: Application
Priority: optional
Architecture: {{.Arch}}
Maintainer: {{.Maintainer}}
Description: {{.Comments}}
`
	data := make(map[string]any)
	data["Name"] = proj.Name
	data["CompanyName"] = proj.Info.CompanyName
	data["ProductName"] = proj.Info.ProductName
	data["ProductVersion"] = proj.Info.ProductVersion
	data["Arch"] = ""
	data["Maintainer"] = ""
	data["Comments"] = ""
	content, err := tools.RenderTemplate(controlTemp, data)
	if err != nil {
		return err
	}
	debControlFile := filepath.Join(appDir, debControl)
	ioutil.WriteFile(debControlFile, content, fs.ModePerm)
	return nil
}
