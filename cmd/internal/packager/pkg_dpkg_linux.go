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
	"github.com/energye/energy/v2/cmd/internal/consts"
	"github.com/energye/energy/v2/cmd/internal/project"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

const (
	deb               = "DEBIAN"
	debControl        = deb + "/control"
	debPreinit        = deb + "/preinit"
	debPostinit       = deb + "/postinit"
	debPrerm          = deb + "/prerm"
	debPostrm         = deb + "/postrm"
	usrSharApps       = "usr/share/applications"
	optCompanyProduct = "opt/%s/%s"
)

const (
	linuxDebControl = "linux/control"
	linuxAppDesktop = "linux/app.desktop"
)

func GeneraInstaller(proj *project.Project) error {
	if !tools.CommandExists("dpkg") {
		return errors.New("failed to create application installation program. Could not find the dpkg command")
	}
	// 创建构建输出目录
	appRoot := fmt.Sprintf("linux/%s-%s", proj.Name, proj.Info.ProductVersion)
	buildOutDir := assets.BuildOutPath(proj)
	buildOutDir = filepath.Join(buildOutDir, appRoot)
	if !tools.IsExist(buildOutDir) {
		if err := os.MkdirAll(buildOutDir, 0755); err != nil {
			return fmt.Errorf("unable to create directory: %w", err)
		}
	}
	var err error
	if err = linuxControl(proj, appRoot); err != nil {
		return err
	}
	if err = linuxDesktop(proj, appRoot); err != nil {
		return err
	}
	if err = linuxOpt(proj, appRoot); err != nil {
		return err
	}
	return nil
}

func opt(proj *project.Project) string {
	return filepath.Join("/opt", proj.Info.CompanyName, proj.Info.ProductName)
}

func linuxOpt(proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate dpkg copy:",
		term.Logger.Args("company", proj.Info.CompanyName, "product", proj.Info.ProductName, "opt",
			fmt.Sprintf("/opt/%s/%s", proj.Info.CompanyName, proj.Info.ProductName)))
	buildOutDir := assets.BuildOutPath(proj)
	appDir := filepath.Join(buildOutDir, appRoot)
	// app/opt/[company]/[product]
	optDir := filepath.Join(appDir, fmt.Sprintf(optCompanyProduct, proj.Info.CompanyName, proj.Info.ProductName))
	if err := os.MkdirAll(optDir, 0755); err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}
	exeDir := filepath.Join(proj.ProjectPath, proj.OutputFilename)
	if !tools.IsExist(exeDir) {
		return fmt.Errorf("execution file not found: %s", exeDir)
	}
	term.Logger.Info("Generate dpkg execution " + exeDir)
	cefDir := os.Getenv(consts.EnergyHomeKey)
	if !tools.IsExist(cefDir) {
		return fmt.Errorf("%s not found: %s", consts.EnergyHomeKey, cefDir)
	}
	term.Logger.Info("Generate dpkg framework " + cefDir)
	var copyFiles = func(src, dst string) error {
		if srcFile, err := os.Open(src); err != nil {
			return err
		} else {
			defer srcFile.Close()
			st, err := srcFile.Stat()
			if err != nil {
				return err
			}
			if st.IsDir() {

			} else {
				dstFilePath := filepath.Join(dst, st.Name())
				dstFile, err := os.OpenFile(dstFilePath, os.O_CREATE|os.O_WRONLY, 0755)
				if err != nil {
					return err
				}
				defer dstFile.Close()
				_, err = io.Copy(dstFile, srcFile)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
	term.Logger.Info("Generate dpkg copy:", term.Logger.Args("execution", exeDir))
	if err := copyFiles(exeDir, optDir); err != nil {
		return nil
	}
	term.Logger.Info("Generate dpkg copy:", term.Logger.Args("framework", cefDir))
	if err := copyFiles(cefDir, optDir); err != nil {
		return nil
	}
	return nil
}

func linuxDesktop(proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate dpkg desktop")
	buildOutDir := assets.BuildOutPath(proj)
	appDir := filepath.Join(buildOutDir, appRoot)
	// app/usr/share/applications
	apps := filepath.Join(appDir, usrSharApps)
	if err := os.MkdirAll(apps, 0755); err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}
	if desktopData, err := assets.ReadFile(proj, assetsFSPath, linuxAppDesktop); err != nil {
		return err
	} else {
		optDir := opt(proj)
		_, icon := filepath.Split(proj.Info.Icon)
		data := make(map[string]any)
		data["Name"] = proj.Name
		data["Exec"] = filepath.Join(optDir, proj.Name)
		data["Icon"] = filepath.Join(optDir, icon)
		data["Comments"] = proj.Info.Comments
		if content, err := tools.RenderTemplate(string(desktopData), data); err != nil {
			return err
		} else {
			debControlFile := filepath.Join(appRoot, usrSharApps, fmt.Sprintf("%s.desktop", proj.Name))
			if err = assets.WriteFile(proj, debControlFile, content); err != nil {
				return err
			}
		}
	}
	return nil
}

func linuxControl(proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate dpkg control")
	buildOutDir := assets.BuildOutPath(proj)
	appDir := filepath.Join(buildOutDir, appRoot)
	// DEBIAN app/DEBIAN
	debDir := filepath.Join(appDir, deb)
	if err := os.MkdirAll(debDir, 0755); err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}
	if controlData, err := assets.ReadFile(proj, assetsFSPath, linuxDebControl); err != nil {
		return err
	} else {
		data := make(map[string]any)
		data["Arch"] = runtime.GOARCH
		data["Info"] = proj.Info
		data["Author"] = proj.Author
		data["Dpkg"] = proj.Dpkg
		if content, err := tools.RenderTemplate(string(controlData), data); err != nil {
			return err
		} else {
			debControlFile := filepath.Join(appRoot, debControl)
			if err = assets.WriteFile(proj, debControlFile, content); err != nil {
				return err
			}
		}
	}
	return nil
}
