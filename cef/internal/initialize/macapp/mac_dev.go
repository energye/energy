// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

//go:build darwin && !prod
// +build darwin,!prod

package macapp

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/energye/energy/v2/cef/config"
	"github.com/energye/golcl/energy/consts"
	"github.com/energye/golcl/pkgs/libname"
	"github.com/energye/golcl/tools/command"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	mac_cef_helper          = "Helper"
	mac_cef_helper_Gpu      = "Helper (GPU)"
	mac_cef_helper_Plugin   = "Helper (Plugin)"
	mac_cef_helper_Renderer = "Helper (Renderer)"
	cef                     = "Chromium Embedded Framework.framework"
)

var helpers = []string{
	mac_cef_helper,
	mac_cef_helper_Gpu,
	mac_cef_helper_Plugin,
	mac_cef_helper_Renderer,
}

func init() {
	MacApp.execName = os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]
	MacApp.macContentsDir = os.Args[0] + ".app/Contents"
	MacApp.macOSDir = MacApp.macContentsDir + "/MacOS"
	MacApp.macResources = MacApp.macContentsDir + "/Resources"
	MacApp.execFile = MacApp.macOSDir + "/" + MacApp.execName
	MacApp.plistFileName = MacApp.macContentsDir + "/Info.plist"
	MacApp.pkgInfoFileName = MacApp.macContentsDir + "/PkgInfo"
	MacApp.macAppFrameworksDir = MacApp.macContentsDir + "/Frameworks"
	MacApp.lclLibFileName = MacApp.macContentsDir + "/Frameworks/liblcl.dylib" //liblcl to frameworks
	MacApp.lsUIElement = "false"
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

/*
 * 拷贝文件夹,同时拷贝文件夹中的文件
 * srcPath 需要拷贝的文件夹路径: /test
 * destPath 拷贝到的位置: /backup/
 */
func copyDir(srcPath string, destPath string) error {
	if srcInfo, err := os.Stat(srcPath); err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		if !srcInfo.IsDir() {
			e := errors.New("srcPath不是一个正确的目录！")
			return e
		}
	}
	if destInfo, err := os.Stat(destPath); err != nil {
		return err
	} else {
		if !destInfo.IsDir() {
			e := errors.New("destInfo不是一个正确的目录！")
			return e
		}
	}
	sps := len(srcPath)
	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			copyFile(path, filepath.Join(destPath, path[sps:]))
		}
		return nil
	})
	return err
}

// 生成目录并拷贝文件
func copyFile(src, dest string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()
	destSplitPathDirs := strings.Split(dest, "/")
	destSplitPath := ""
	for index, dir := range destSplitPathDirs {
		if index < len(destSplitPathDirs)-1 {
			destSplitPath = destSplitPath + dir + "/"
			b := fileExists(destSplitPath)
			if b == false {
				if err := os.Mkdir(destSplitPath, 0755); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	dstFile, err := os.Create(dest)
	if err != nil {
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}

func Init() {
	if strings.Contains(os.Args[0], ".app/Contents/MacOS") {
		return
	}
	cfg := config.Get()
	// CEF 版本大于 109 时，helper 进程使用 ln 软链接执行文件
	// 109 及以下版本 helper 进程 copy 执行文件, 不然启动 helper 进程失败
	MacApp.IsLinked(cfg.Version() > 109)
	MacApp.SetBaseCefFrameworksDir(cfg.FrameworkPath())
	MacApp.isMain = true
	// 创建 xxx.app
	if MacApp.createMacOSApp(MacApp) {
		MacApp.copyDylib()
		MacApp.createCEFHelper()
		MacApp.runMacOSApp()
	}
}

func (m *macApp) SetBaseCefFrameworksDir(s string) {
	m.cefFrameworksDir = s
}

func (m *macApp) SetBrowseSubprocessPath(s string) {
	m.browseSubprocessPath = s
}

func (m *macApp) IsLinked(s bool) {
	m.isLinked = s
}

func (m *macApp) createCEFHelper() {
	if m.cefFrameworksDir == "" {
		m.cefFrameworksDir = os.Getenv("ENERGY_HOME")
	}
	if !fileExists(m.cefFrameworksDir) {
		panic("cef frameworks 不存在: " + m.cefFrameworksDir)
	}
	if !fileExists(m.macAppFrameworksDir) {
		os.Mkdir(m.macAppFrameworksDir, 0755)
	}
	if m.browseSubprocessPath != "" {
		if !fileExists(m.browseSubprocessPath) {
			panic("子进程执行文件不存在: " + m.browseSubprocessPath)
		}
		// copy helper process to xxx.app/Contents/MacOS/xxxHelper
		helperName := m.execFile[strings.LastIndex(m.execFile, "/")+1:] + "Helper"
		helperPath := m.execFile[:strings.LastIndex(m.execFile, "/")]
		helperFilePath := filepath.Join(helperPath, helperName)
		if _, err := copyFile(m.browseSubprocessPath, helperFilePath); err == nil {
			os.Chmod(helperFilePath, 0755)
		}
		m.browseSubprocessPath = helperFilePath
	}
	b := fileExists(m.macAppFrameworksDir + consts.Separator + cef)
	if !b {
		copyDir(m.cefFrameworksDir, m.macAppFrameworksDir)
	}
	for _, app := range helpers {
		var execName = m.execName
		hPath := fmt.Sprintf("%s/%s %s.app", m.macAppFrameworksDir, execName, app)
		if !fileExists(hPath) {
			if err := os.MkdirAll(hPath, 0755); err != nil {
				panic("创建cef helper失败: " + err.Error())
			}
		}
		helper := &macApp{}
		helper.browseSubprocessPath = m.browseSubprocessPath
		helper.execName = fmt.Sprintf("%s %s", m.execName, app)
		helper.macContentsDir = hPath + "/Contents"
		helper.macOSDir = helper.macContentsDir + "/MacOS"
		helper.macResources = helper.macContentsDir + "/Resources"
		helper.execFile = helper.macOSDir + "/" + helper.execName
		helper.plistFileName = helper.macContentsDir + "/Info.plist"
		helper.pkgInfoFileName = helper.macContentsDir + "/PkgInfo"
		helper.macAppFrameworksDir = helper.macContentsDir + "/Frameworks"
		helper.lsUIElement = "true"
		helper.isLinked = m.isLinked
		// 创建 helper 进程
		m.createMacOSApp(helper)
		cmd := command.NewCMD()
		cmd.Dir = helper.macAppFrameworksDir
		cmd.IsPrint = false
		cmd.Command("ln", "-shf", "../../../liblcl.dylib", "liblcl.dylib")
	}
}

func (m *macApp) runMacOSApp() {
	cmd := exec.Command(m.execFile, os.Args...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func (m *macApp) copyDylib() {
	var libPath = libname.LibPath(libname.GetDLLName())
	// 文件不存在，复制
	if !fileExists(m.lclLibFileName) {
		if fileExists(libPath) {
			copyFile(libPath, m.lclLibFileName)
		}
	} else {
		// 文件存在，对比后更新
		if fileExists(libPath) {
			f1, _ := os.Stat(libPath)
			f2, _ := os.Stat(m.lclLibFileName)
			if f1.Size() != f2.Size() {
				copyFile(libPath, m.lclLibFileName)
			}
		}
	}
}

// 以一个Mac下的app形式运行
// 调试下使用，正式发布的时候虽然可 以不用去掉，但也不咋好
func (*macApp) createMacOSApp(m *macApp) bool {
	if !fileExists(m.macOSDir) {
		if err := os.MkdirAll(m.macOSDir, 0755); err != nil {
			return false
		}
	}
	if !fileExists(m.macResources) {
		os.MkdirAll(m.macResources, 0755)
	}
	if !fileExists(m.macAppFrameworksDir) {
		os.MkdirAll(m.macAppFrameworksDir, 0755)
	}
	resName := fmt.Sprintf("%s/%s.icns", m.macResources, m.execName)
	if !fileExists(resName) {
		ioutil.WriteFile(resName, macOSAppIcon, 0666)
	}
	if !fileExists(m.plistFileName) {
		datas := map[string]string{
			"execName":    m.execName,
			"devRegion":   "China", // China English
			"locale":      "zh_CN", //os.Getenv("LANG"),
			"copyright":   "copyright xxxx",
			"LSUIElement": m.lsUIElement,
		}
		buff := bytes.NewBuffer([]byte{})
		tmp := template.New("file")
		tmp.Parse(infoplist)
		tmp.Execute(buff, datas)
		ioutil.WriteFile(m.plistFileName, buff.Bytes(), 0666)
	}
	if !fileExists(m.pkgInfoFileName) {
		ioutil.WriteFile(m.pkgInfoFileName, pkgInfo, 0666)
	}
	if m.browseSubprocessPath != "" && !m.isMain {
		// 子进程是单独的执行文件
		createHelperExe(m.browseSubprocessPath, m.execFile, m.isLinked)
	} else if m.isMain {
		// 主进程执行文件
		if _, err := copyFile(os.Args[0], m.execFile); err == nil {
			os.Chmod(m.execFile, 0755)
			return true
		}
	} else {
		// 子进程和主进程使用同一个执行文件
		createHelperExe(os.Args[0], m.execFile, m.isLinked)
	}
	return false
}

// mainExec 是 xxx.app/Contents/MacOS/mainExec
func createHelperExe(mainExec, execFile string, isLinked bool) {
	if isLinked {
		workDir := execFile
		lnDest := execFile
		// 工作目录
		workDir = workDir[:strings.LastIndex(workDir, "/")]
		// 执行文件名
		mainExec = mainExec[strings.LastIndex(mainExec, "/")+1:]
		lnDest = lnDest[strings.LastIndex(lnDest, "/")+1:]
		lnSource := fmt.Sprintf("../../../../MacOS/%s", mainExec)
		cmd := command.NewCMD()
		cmd.Dir = workDir
		cmd.IsPrint = false
		cmd.Command("ln", "-s", lnSource, lnDest)
	} else {
		if _, err := copyFile(mainExec, execFile); err == nil {
			os.Chmod(execFile, 0755)
		}
	}
}
