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
	"github.com/energye/energy/v2/cmd/internal/assets"
	"github.com/energye/energy/v2/cmd/internal/project"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools"
	"github.com/energye/golcl/tools/command"
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
)

func GeneraInstaller(proj *project.Project) error {
	appRoot := fmt.Sprintf("darwin/%s.app", proj.Name)
	buildOutDir := assets.BuildOutPath(proj)
	buildOutDir = filepath.Join(buildOutDir, appRoot)
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
	return nil
}

func createApp(proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate app create app dir.")
	buildOutDir := assets.BuildOutPath(proj)
	appDir := filepath.Join(buildOutDir, appRoot)
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
			if err != nil {
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
		src, err := os.Open(proj.PList.Icon)
		if err != nil {
			return err
		}
		defer src.Close()
		// Contents
		contents := filepath.Join(appDir, appContents)
		// Contents/Resources
		_, icnsName := filepath.Split(proj.PList.Icon)
		outIcnsFilePath := filepath.Join(contents, appContentsResources, icnsName)
		os.Remove(outIcnsFilePath)
		dst, err := os.OpenFile(outIcnsFilePath, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return err
		}
		defer dst.Close()
		_, err = io.Copy(dst, src)
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
		data := make(map[string]any)
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

func copyFrameworkFile(proj *project.Project, appRoot string) error {
	term.Logger.Info("Generate app copy framework:",
		term.Logger.Args("company", proj.Info.CompanyName, "product", proj.Info.ProductName))
	//buildOutDir := assets.BuildOutPath(proj)
	//appDir := filepath.Join(buildOutDir, appRoot)
	//contents := filepath.Join(appDir, appContents)
	//var copyFiles = func(src, dst string) error {
	//	if srcFile, err := os.Open(src); err != nil {
	//		return err
	//	} else {
	//		st, err := srcFile.Stat()
	//		if err != nil {
	//			return err
	//		}
	//		if st.IsDir() {
	//			srcFile.Close() //close
	//			var pathLen = len(src)
	//			err := filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
	//				if path == src { // current root
	//					return nil
	//				}
	//				outPath := path[pathLen:]
	//				// exclude file or dir
	//				for _, p := range proj.Dpkg.Exclude {
	//					if strings.Contains(outPath, p) {
	//						return nil
	//					}
	//				}
	//				targetPath := filepath.Join(dst, outPath)
	//				info, _ := d.Info()
	//				if d.IsDir() {
	//					return os.MkdirAll(targetPath, info.Mode())
	//				} else {
	//					if tools.IsExistAndSize(targetPath, info.Size()) {
	//						term.Logger.Info("\tcopy skip: " + outPath)
	//						return nil
	//					}
	//					srcFile, err := os.Open(path)
	//					if err != nil {
	//						return err
	//					}
	//					defer srcFile.Close()
	//					dstFile, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY, info.Mode())
	//					if err != nil {
	//						return err
	//					}
	//					defer dstFile.Close()
	//					term.Logger.Info("\tcopy: " + outPath)
	//					_, err = io.Copy(dstFile, srcFile)
	//					return err
	//				}
	//			})
	//			if err != nil {
	//				return err
	//			}
	//		} else {
	//			defer srcFile.Close()
	//			dstFilePath := filepath.Join(dst, st.Name())
	//			dstFile, err := os.OpenFile(dstFilePath, os.O_CREATE|os.O_WRONLY, st.Mode())
	//			if err != nil {
	//				return err
	//			}
	//			defer dstFile.Close()
	//			_, err = io.Copy(dstFile, srcFile)
	//			if err != nil {
	//				return err
	//			}
	//		}
	//	}
	//	return nil
	//}

	return nil
}
