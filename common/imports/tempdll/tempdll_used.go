//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package tempdll

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/energy/emfs"
	"hash/crc32"
	"io/ioutil"
	"os"
	"path"
)

func init() {
	TempDLL = &temdll{
		dllFSDir: "libs",
	}
}

// CheckAndReleaseDLL
//
//	检查动态库并释放
func CheckAndReleaseDLL(dllName string) (filePath string, fileFullPath string, ok bool) {
	if TempDLL == nil || TempDLL.DllSaveDirType() == TddInvalid || emfs.GetLibsFS() == nil {
		return
	}
	// 在内置libs内读取
	liblclData, err := emfs.GetLibsFS().ReadFile(path.Join(TempDLL.dllFSDir, dllName))
	if err != nil {
		return
	}
	crc32Val := crc32.ChecksumIEEE(liblclData)

	// 动态库保存目录
	var tempDLLDir = fmt.Sprintf("%s/liblcl/%x", os.TempDir(), crc32Val)
	switch TempDLL.DllSaveDirType() {
	case TddCurrent:
		tempDLLDir = consts.ExeDir
	case TddEnergyHome:
		//tempDLLDir = os.Getenv(consts.ENERGY_HOME_KEY)// TODO
	case TddCustom:
		if TempDLL.DllSaveDir() != "" {
			tempDLLDir = TempDLL.DllSaveDir()
		}
	}

	// create liblcl: $tempdir/liblcl/{crc32}/liblcl.{ext}
	if !fileExists(tempDLLDir) {
		if err := os.MkdirAll(tempDLLDir, 0775); err != nil {
			return
		}
	}
	// 如果不是自定义设置目录，在这里设置到DllSaveDir方便使用时获取
	if TempDLL.DllSaveDirType() != TddCustom {
		TempDLL.SetDllSaveDir(tempDLLDir)
	}
	// test crc32
	// 防止liblcl被修改
	tempDLLFileName := fmt.Sprintf("%s/%s", tempDLLDir, dllName)
	if fileExists(tempDLLFileName) {
		bs, err := ioutil.ReadFile(tempDLLFileName)
		if err == nil {
			if crc32.ChecksumIEEE(bs) != crc32Val {
				os.Remove(tempDLLFileName)
			}
		}
	}
	if !fileExists(tempDLLFileName) {
		if err := releaseLib(tempDLLFileName, liblclData); err != nil {
			if os.Remove(tempDLLFileName) != nil {
				return
			}
		}
	}
	return tempDLLDir, tempDLLFileName, true
}

func releaseLib(destFileName string, input []byte) error {
	var file *os.File
	file, err := os.OpenFile(destFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(input)
	return err
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
