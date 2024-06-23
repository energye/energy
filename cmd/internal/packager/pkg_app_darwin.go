//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin
// +build darwin

package packager

import (
	"errors"
	"fmt"
	"github.com/energye/energy/v3/cmd/internal/assets"
	"github.com/energye/energy/v3/cmd/internal/consts"
	"github.com/energye/energy/v3/cmd/internal/pkgs/command"
	"github.com/energye/energy/v3/cmd/internal/project"
	"github.com/energye/energy/v3/cmd/internal/term"
	"github.com/energye/energy/v3/cmd/internal/tools"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	appContents           = "Contents"
	appContentsFrameworks = "Frameworks"
	appContentsMacOS      = "MacOS"
	appContentsResources  = "Resources"
)

const (
	darwinInfoPList = "darwin/Info.plist"
)

const (
	macCefHelper          = "Helper"
	macCefHelperGpu       = "Helper (GPU)"
	macCefHelperPlugin    = "Helper (Plugin)"
	macCefHelpersRenderer = "Helper (Renderer)"
)

var (
	pkgInfo  = []byte{0x41, 0x50, 0x50, 0x4C, 0x3F, 0x3F, 0x3F, 0x3F, 0x0D, 0x0A}
	sipsCmds = []string{
		"-z 16 16 %s -o icons.iconset/%s_16x16.png",
		"-z 32 32 %s -o icons.iconset/%s_16x16@2x.png",
		"-z 32 32 %s -o icons.iconset/%s_32x32.png",
		"-z 64 64 %s -o icons.iconset/%s_32x32@2x.png",
		"-z 128 128 %s -o icons.iconset/%s_128x128.png",
		"-z 256 256 %s -o icons.iconset/%s_128x128@2x.png",
		"-z 256 256 %s -o icons.iconset/%s_256x256.png",
		"-z 512 512 %s -o icons.iconset/%s_256x256@2x.png",
		"-z 512 512 %s -o icons.iconset/%s_512x512.png",
		"-z 1024 1024 %s -o icons.iconset/%s_512x512@2x.png",
	}
	helpers = []string{
		macCefHelper,
		macCefHelperGpu,
		macCefHelperPlugin,
		macCefHelpersRenderer,
	}
	projectPath string
)

func GeneraInstaller(proj *project.Project) error {
	appRoot := fmt.Sprintf("darwin/%s.app", proj.Name)
	buildOutDir := assets.BuildOutPath(proj)
	buildOutDir = filepath.Join(buildOutDir, appRoot)
	projectPath = proj.ProjectPath
	if !tools.IsExist(buildOutDir) {
		if err := os.MkdirAll(buildOutDir, 0755); err != nil {
			return fmt.Errorf("unable to create directory: %w", err)
		}
	}
	if err := createApp(proj, appRoot); err != nil {
		return err
	}
	if err := generateICNS(proj, appRoot); err != nil {
		return err
	}
	if err := createAppInfoPList(proj, appRoot); err != nil {
		return err
	}
	if err := createAppPkgInfo(proj, appRoot); err != nil {
		return err
	}
	if err := copyFrameworkFile(proj, appRoot); err != nil {
		return err
	}
	if err := copyHelperFile(proj, appRoot); err != nil {
		return err
	}
	if proj.PList.Pkgbuild {
		if err := pkgbuild(proj, appRoot); err != nil {
			return err
		}
	}
	return nil
}

func copyFiles(proj *project.Project, src, dst string) error {
	if srcFile, err := os.Open(src); err != nil {
		return err
	} else {
		st, err := srcFile.Stat()
		if err != nil {
			return err
		}
		if st.IsDir() {
			srcFile.Close() //close
			var pathLen = len(src)
			err := filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
				if path == src { // current root
					return nil
				}
				outPath := path[pathLen:]
				// exclude file or dir
				for _, p := range proj.PList.Exclude {
					if strings.Contains(outPath, p) {
						return nil
					}
				}
				targetPath := filepath.Join(dst, outPath)
				info, _ := d.Info()
				if d.IsDir() {
					return os.MkdirAll(targetPath, info.Mode())
				} else {
					if tools.IsExistAndSize(targetPath, info.Size()) {
						//term.Logger.Info("\tcopy skip: " + outPath)
						return nil
					}
					srcFile, err := os.Open(path)
					if err != nil {
						return err
					}
					defer srcFile.Close()
					dstFile, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY, info.Mode())
					if err != nil {
						return err
					}
					defer dstFile.Close()
					//term.Logger.Info("\tcopy: " + outPath)
					_, err = io.Copy(dstFile, srcFile)
					return err
				}
			})
			if err != nil {
				return err
			}
		} else {
			defer srcFile.Close()
			dstFilePath := filepath.Join(dst, st.Name())
			dstFile, err := os.OpenFile(dstFilePath, os.O_CREATE|os.O_WRONLY, st.Mode())
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

func copyFrameworkFile(proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate app copy framework:",
		term.Logger.Args("company", proj.Info.CompanyName, "product", proj.Info.ProductName))
	buildOutDir := assets.BuildOutPath(proj)
	appDir := filepath.Join(buildOutDir, appRoot)
	// Contents
	contents := filepath.Join(appDir, appContents)
	exeDir := filepath.Join(proj.ProjectPath, proj.OutputFilename)
	if !tools.IsExist(exeDir) {
		return fmt.Errorf("execution file not found: %s", exeDir)
	}
	cefDir := os.Getenv(consts.EnergyHomeKey)
	if !tools.IsExist(cefDir) {
		return fmt.Errorf("%s not found: %s", consts.EnergyHomeKey, cefDir)
	}
	term.Logger.Info("Generate app copy:", term.Logger.Args("execution", exeDir))
	// Contents/MacOS/exe
	outExe := filepath.Join(contents, appContentsMacOS)
	if err := copyFiles(proj, exeDir, outExe); err != nil {
		return err
	}
	term.Logger.Info("Generate app copy:", term.Logger.Args("framework", cefDir))
	// Contents/Frameworks/cef
	outCEF := filepath.Join(contents, appContentsFrameworks)
	if err := copyFiles(proj, cefDir, outCEF); err != nil {
		return err
	}
	return nil
}

// pkgbuild --root demo.app --identifier com.demo.demo --version 1.0.0 --install-location /Applications/demo.app demo.pkg
func pkgbuild(proj *project.Project, appRoot string) error {
	proj.AppType = project.AtApp
	proj.ProjectPath = projectPath
	buildOutDir := assets.BuildOutPath(proj)
	cmdWorkDir := filepath.Join(buildOutDir, "darwin")
	term.Logger.Info("Generate app pkgbuild", term.Logger.Args("cmd work dir", cmdWorkDir))
	// remove xxx.pkg
	os.Remove(filepath.Join(cmdWorkDir, fmt.Sprintf("%s.pkg", proj.Name)))
	cmd := command.NewCMD()
	//cmd.IsPrint = false
	cmd.Dir = cmdWorkDir
	cmd.MessageCallback = func(bytes []byte, err error) {
		msg := string(bytes)
		if msg != "" {
			println(msg)
		}
	}
	var args = []string{"--root", fmt.Sprintf("%s.app", proj.Name),
		"--identifier", fmt.Sprintf("com.%s.%s", proj.PList.CompanyName, proj.PList.ProductName),
		"--version", proj.PList.CFBundleVersion,
		"--install-location", fmt.Sprintf("/Applications/%s.app", proj.Name),
		fmt.Sprintf("%s.pkg", proj.Name)}
	cmd.Command("pkgbuild", args...)
	cmd.Close()
	// remove xxx.app
	os.Remove(filepath.Join(buildOutDir, appRoot))
	return nil
}

func copyHelperFile(proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate app copy cef helper")
	buildOutDir := assets.BuildOutPath(proj)
	appDir := filepath.Join(buildOutDir, appRoot)
	contents := filepath.Join(appDir, appContents)
	frameworksDir := filepath.Join(contents, appContentsFrameworks)
	// 应用构建执行文件目录
	exeDir := filepath.Join(proj.ProjectPath, proj.OutputFilename)
	var err error
	exeFile, err := os.Open(exeDir)
	if err != nil {
		return err
	}
	defer exeFile.Close()
	// 设置为helper选项
	proj.PList.LSUIElement = true
	proj.AppType = project.AtHelper
	cmd := command.NewCMD()
	cmd.IsPrint = false
	cmd.MessageCallback = func(bytes []byte, e error) {
		if e != nil {
			err = e
		}
	}
	for _, app := range helpers {
		proj.ProjectPath = frameworksDir
		helperAppRoot := fmt.Sprintf("%s %s.app", proj.Name, app)
		// helper app
		if err = createApp(proj, helperAppRoot); err != nil {
			return err
		}
		// helper app Info.plist
		if err = createAppInfoPList(proj, helperAppRoot); err != nil {
			return err
		}
		// helper PkgInfo
		if err = createAppPkgInfo(proj, helperAppRoot); err != nil {
			return err
		}
		// helper ln liblcl.dylib
		cmd.Dir = filepath.Join(proj.ProjectPath, helperAppRoot, appContents, appContentsFrameworks)
		cmd.Command("ln", "-shf", "../../../liblcl.dylib", "liblcl.dylib")
		if err != nil {
			return err
		}
		// helper exe
		helperMacOSExe := filepath.Join(proj.ProjectPath, helperAppRoot, appContents, appContentsMacOS, fmt.Sprintf("%s %s", proj.OutputFilename, app))
		helperMacOSExeFile, err := os.OpenFile(helperMacOSExe, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return err
		}
		io.Copy(helperMacOSExeFile, exeFile)
		helperMacOSExeFile.Close()
		exeFile.Seek(0, 0)
	}
	cmd.Close()
	return nil
}

func createApp(proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate app create app dir: " + appRoot)
	buildOutDir := assets.BuildOutPath(proj)
	appDir := filepath.Join(buildOutDir, appRoot)
	os.Remove(appDir)
	// Contents
	contents := filepath.Join(appDir, appContents)
	if err := os.MkdirAll(contents, 0755); err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}
	// Contents/Frameworks
	contentsFrameworks := filepath.Join(contents, appContentsFrameworks)
	if err := os.MkdirAll(contentsFrameworks, 0755); err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}
	// Contents/MacOS
	contentsMacOS := filepath.Join(contents, appContentsMacOS)
	if err := os.MkdirAll(contentsMacOS, 0755); err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}
	// Contents/Resources
	contentsResources := filepath.Join(contents, appContentsResources)
	if err := os.MkdirAll(contentsResources, 0755); err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}
	return nil
}

func generateICNS(proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate app icns")
	buildOutDir := assets.BuildOutPath(proj)
	appDir := filepath.Join(buildOutDir, appRoot)
	iconExt := strings.ToLower(filepath.Ext(proj.PList.Icon))
	tmpWorkDir := filepath.Join(buildOutDir, "tmp")
	os.Remove(tmpWorkDir)
	if iconExt == ".png" {
		term.Logger.Info("\tcreate icnns")
		src, err := os.Open(proj.PList.Icon)
		if err != nil {
			return err
		}
		var closeSrc = func() {
			if src != nil {
				src.Close()
				src = nil
			}
		}
		defer closeSrc()
		_, icnsName := filepath.Split(proj.PList.Icon)
		os.MkdirAll(tmpWorkDir, fs.ModePerm)

		outIconsetPath := filepath.Join(tmpWorkDir, "icons.iconset")
		os.MkdirAll(outIconsetPath, fs.ModePerm)
		tmpIconFilePath := filepath.Join(tmpWorkDir, icnsName)
		dst, err := os.OpenFile(tmpIconFilePath, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return err
		}
		io.Copy(dst, src)
		dst.Close()
		closeSrc()
		name := icnsName[:len(icnsName)-4]
		// 生成图标
		cmd := command.NewCMD()
		cmd.Dir = tmpWorkDir
		cmd.IsPrint = false
		cmd.MessageCallback = func(bytes []byte, e error) {
			if e != nil {
				err = e
			}
		}
		for _, arg := range sipsCmds {
			cmd.Command("sips", strings.Split(fmt.Sprintf(arg, icnsName, name), " ")...)
		}
		cmd.Command("iconutil", strings.Split(fmt.Sprintf("-c icns icons.iconset -o %s.icns", name), " ")...)
		cmd.Close()
		if err != nil {
			return err
		}
		proj.PList.Icon = filepath.Join(tmpWorkDir, name+".icns")
	}
	iconExt = strings.ToLower(filepath.Ext(proj.PList.Icon))
	if iconExt == ".icns" {
		term.Logger.Info("\tcopy icns")
		// Contents
		contents := filepath.Join(appDir, appContents)
		// Contents/Resources
		_, icnsName := filepath.Split(proj.PList.Icon)
		outIcnsFilePath := filepath.Join(contents, appContentsResources)
		err := copyFiles(proj, proj.PList.Icon, outIcnsFilePath)
		if err != nil {
			return err
		}
		// 设置icon文件名称
		proj.PList.Icon = icnsName
	} else {
		return errors.New("app icon, only supports png or icns")
	}
	return nil
}

func createAppInfoPList(proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate app Info.plist")
	// Contents/Info.plist
	if plistData, err := assets.ReadFile(proj, assetsFSPath, darwinInfoPList); err != nil {
		return err
	} else {
		data := make(map[string]interface{})
		data["Name"] = proj.Name
		data["OutputFilename"] = proj.OutputFilename
		data["PList"] = proj.PList
		if content, err := tools.RenderTemplate(string(plistData), data); err != nil {
			return err
		} else {
			infoPListFile := filepath.Join(appRoot, appContents, "Info.plist")
			if err = assets.WriteFile(proj, infoPListFile, content); err != nil {
				return err
			}
		}
	}
	return nil
}

func createAppPkgInfo(proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate app PkgInfo")
	// Contents/PkgInfo
	infoPListFile := filepath.Join(appRoot, appContents, "PkgInfo")
	if err := assets.WriteFile(proj, infoPListFile, pkgInfo); err != nil {
		return err
	}
	return nil
}
