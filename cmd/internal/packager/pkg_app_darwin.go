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
	"github.com/cyber-xxm/energy/v2/cmd/internal/assets"
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/env"
	"github.com/cyber-xxm/energy/v2/cmd/internal/project"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"github.com/cyber-xxm/energy/v2/cmd/internal/tools"
	cmd "github.com/cyber-xxm/energy/v2/cmd/internal/tools/cmd"
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

func GeneraInstaller(c *command.Config, proj *project.Project) error {
	appRoot := fmt.Sprintf("darwin/%s.app", getAppName(c, proj))
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
	if err := createAppInfoPList(c, proj, appRoot, getExeName(c, proj)); err != nil {
		return err
	}
	if err := createAppPkgInfo(proj, appRoot); err != nil {
		return err
	}
	if err := copyFrameworkFile(c, proj, appRoot); err != nil {
		return err
	}
	if err := copyHelperFile(c, proj, appRoot); err != nil {
		return err
	}
	if proj.PList.Pkgbuild {
		if err := pkgbuild(c, proj, appRoot); err != nil {
			return err
		}
	}
	return nil
}

func getAppName(c *command.Config, proj *project.Project) string {
	appName := proj.OutputFilename
	if c.Package.OutFileName != "" {
		appName = c.Package.OutFileName
	}
	return appName
}

func getExeName(c *command.Config, proj *project.Project) string {
	exeName := proj.Name
	if c.Package.File != "" {
		exeName = c.Package.File
	}
	return exeName
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

func copyFrameworkFile(c *command.Config, proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate app copy framework:",
		term.Logger.Args("company", proj.Info.CompanyName, "product", proj.Info.ProductName))
	buildOutDir := assets.BuildOutPath(proj)
	appDir := filepath.Join(buildOutDir, appRoot)
	// Contents
	contents := filepath.Join(appDir, appContents)
	exeDir := filepath.Join(proj.ProjectPath, getExeName(c, proj))
	if !tools.IsExist(exeDir) {
		return fmt.Errorf("execution file not found: %s", exeDir)
	}
	cefDir := env.GlobalDevEnvConfig.FrameworkPath()
	if !tools.IsExist(cefDir) {
		return fmt.Errorf("%s not found", cefDir)
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
func pkgbuild(c *command.Config, proj *project.Project, appRoot string) error {
	proj.AppType = project.AtApp
	proj.ProjectPath = projectPath
	buildOutDir := assets.BuildOutPath(proj)
	cmdWorkDir := filepath.Join(buildOutDir, "darwin")
	term.Logger.Info("Generate app pkgbuild", term.Logger.Args("cmd work dir", cmdWorkDir))
	// remove xxx.pkg
	os.Remove(filepath.Join(cmdWorkDir, fmt.Sprintf("%s.pkg", getAppName(c, proj))))
	cmd := cmd.NewCMD()
	//cmd.IsPrint = false
	cmd.Dir = cmdWorkDir
	cmd.MessageCallback = func(bytes []byte, err error) {
		msg := string(bytes)
		if msg != "" {
			println(msg)
		}
	}
	app := fmt.Sprintf("%s.app", getAppName(c, proj))
	pkg := fmt.Sprintf("%s.pkg", getAppName(c, proj))
	var args = []string{"--root", app,
		"--identifier", proj.PList.BundleIdentifier,
		"--version", proj.PList.BundleVersion,
		"--install-location", fmt.Sprintf("/Applications/%s", app), pkg}
	cmd.Command("pkgbuild", args...)
	cmd.Close()
	// remove xxx.app
	os.RemoveAll(filepath.Join(cmdWorkDir, app))
	return nil
}

func copyHelperFile(c *command.Config, proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate app copy cef helper")
	buildOutDir := assets.BuildOutPath(proj)
	appDir := filepath.Join(buildOutDir, appRoot)
	contents := filepath.Join(appDir, appContents)
	frameworksDir := filepath.Join(contents, appContentsFrameworks)
	var err error
	// 设置为helper选项
	proj.PList.LSUIElement = true
	proj.AppType = project.AtHelper
	cmd := cmd.NewCMD()
	cmd.IsPrint = false
	cmd.MessageCallback = func(bytes []byte, e error) {
		if e != nil {
			err = e
		}
	}
	// helper process
	var helperProcessFilename string
	// helper 进程是单独执行文件(非同一个 main exe)
	if proj.HelperFilePath != "" {
		term.Logger.Info("Copy helper process: " + proj.HelperFilePath)
		if !tools.IsExist(proj.HelperFilePath) {
			return errors.New("helper process binary executable file does not exist")
		}
		// copy helper process to xxx.app/Contents/MacOS/xxxHelper
		helperExeFile, err := os.Open(proj.HelperFilePath)
		if err != nil {
			return err
		}
		_, helperExeFilename := filepath.Split(proj.HelperFilePath)
		// xxx + Helper = xxx.app/Contents/MacOS/xxx Helper
		helperProcessFilename = fmt.Sprintf("%sHelper", helperExeFilename)
		helperFilePath := filepath.Join(contents, appContentsMacOS, helperProcessFilename)
		helperFile, err := os.OpenFile(helperFilePath, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return err
		}
		io.Copy(helperFile, helperExeFile)
		helperFile.Close()
	}
	for _, app := range helpers {
		proj.ProjectPath = frameworksDir
		helperAppRoot := fmt.Sprintf("%s %s.app", getExeName(c, proj), app)
		helperAppExeName := fmt.Sprintf("%s %s", getExeName(c, proj), app)
		// helper app
		if err = createApp(proj, helperAppRoot); err != nil {
			return err
		}
		// helper app Info.plist
		if err = createAppInfoPList(c, proj, helperAppRoot, helperAppExeName); err != nil {
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
			// for error
			return err
		}

		// CEF 版本大于 109 时，helper 进程使用 ln 软链接执行文件, 减小 .app 体积
		// 109 及以下版本 helper 进程 copy 执行文件, 不然启动 helper 进程失败，会增大 .app 体积
		isLinked := env.GlobalDevEnvConfig.CEFVersion() > 109
		helperWork := filepath.Join(proj.ProjectPath, helperAppRoot, appContents, appContentsMacOS)
		helperExeFilePath := filepath.Join(helperWork, helperAppExeName)
		// remove ln helper process file
		os.Remove(helperExeFilePath)
		if isLinked {
			// 使用 ln 方式
			var helperSource string
			if helperProcessFilename != "" {
				helperSource = fmt.Sprintf("../../../../MacOS/%s", helperProcessFilename)
			} else {
				helperSource = fmt.Sprintf("../../../../MacOS/%s", getExeName(c, proj))
			}
			cmd.Dir = helperWork
			cmd.Command("ln", "-s", helperSource, helperAppExeName)
			if err != nil {
				// for error
				return err
			}
		} else {
			// 复制文件到 helper 进程
			var sourceExec string
			if helperProcessFilename != "" {
				sourceExec = filepath.Join(contents, appContentsMacOS, helperProcessFilename)
			} else {
				sourceExec = filepath.Join(contents, appContentsMacOS, getExeName(c, proj))
			}
			sourceFile, err := os.Open(sourceExec)
			if err != nil {
				return err
			}
			helperFile, err := os.OpenFile(helperExeFilePath, os.O_CREATE|os.O_WRONLY, 0755)
			if err != nil {
				return err
			}
			io.Copy(helperFile, sourceFile)
			helperFile.Close()
			sourceFile.Close()
		}
	}
	cmd.Close()
	return nil
}

func createApp(proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate App Create Dir: " + appRoot)
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
		term.Logger.Info("\tCreate icons")
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
		// 图标名
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
		// 删除扩展名 .png
		name := icnsName[:len(icnsName)-4]
		// 生成图标
		cmd := cmd.NewCMD()
		cmd.Dir = tmpWorkDir
		cmd.IsPrint = false
		cmd.MessageCallback = func(bytes []byte, e error) {
			if e != nil {
				err = e
			}
		}
		for _, arg := range sipsCmds {
			cmd.Command("sips", strings.Split(fmt.Sprintf(arg, icnsName, "icon"), " ")...)
		}
		icnsOutName := fmt.Sprintf("%s.icns", name)
		icnsArgs := []string{"-c", "icns", "-o", icnsOutName, "icons.iconset"}
		cmd.Command("iconutil", icnsArgs...)
		cmd.Close()
		if err != nil {
			return err
		}
		proj.PList.Icon = filepath.Join(tmpWorkDir, icnsOutName)
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

func createAppInfoPList(c *command.Config, proj *project.Project, appRoot, appExeName string) error {
	term.Logger.Info("Generate app Info.plist")
	// Contents/Info.plist
	if plistData, err := assets.ReadFile(proj, assetsFSPath, darwinInfoPList); err != nil {
		return err
	} else {
		data := make(map[string]interface{})
		data["Executable"] = appExeName
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
