//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build tempdll

// 编译命令: go build -tags="tempdll"

package tempdll

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/pkgs/libname"
	"github.com/energye/liblclbinres"
	"hash/crc32"
	"io"
	"io/ioutil"
	"os"
)

func init() {
	TempDLL = &temdll{}
}

// CheckAndReleaseDLL
//  检查动态库并释放
func CheckAndReleaseDLL() (string, bool) {
	if TempDLL == nil || TempDLL.DllSaveDirType() == TddInvalid {
		return "", false
	}
	// 动态库保存目录
	var tempDLLDir string
	switch TempDLL.DllSaveDirType() {
	case TddTmp: // default
		tempDLLDir = fmt.Sprintf("%s/liblcl/%x", os.TempDir(), liblclbinres.CRC32Value)
	case TddCurrent:
		tempDLLDir = consts.ExePath
	case TddEnergyHome:
		tempDLLDir = os.Getenv(consts.ENERGY_HOME_KEY)
	case TddCustom:
		if TempDLL.DllSaveDir() != "" {
			tempDLLDir = TempDLL.DllSaveDir()
		} else {
			tempDLLDir = fmt.Sprintf("%s/liblcl/%x", os.TempDir(), liblclbinres.CRC32Value)
		}
	default:
		tempDLLDir = fmt.Sprintf("%s/liblcl/%x", os.TempDir(), liblclbinres.CRC32Value)
	}
	if tempDLLDir == "" {
		tempDLLDir = fmt.Sprintf("%s/liblcl/%x", os.TempDir(), liblclbinres.CRC32Value)
	}

	// create liblcl: $tempdir/liblcl/{crc32}/liblcl.{ext}
	if !fileExists(tempDLLDir) {
		if err := os.MkdirAll(tempDLLDir, 0775); err != nil {
			return "", false
		}
	}
	// 设置到tempDllDir
	// 使用tempdll将最优先从该目录加载
	libname.SetTempDllDir(tempDLLDir)
	tempDLLFileName := fmt.Sprintf("%s/%s", tempDLLDir, libname.GetDLLName())
	// test crc32
	if fileExists(tempDLLFileName) {
		bs, err := ioutil.ReadFile(tempDLLFileName)
		if err == nil {
			if crc32.ChecksumIEEE(bs) != liblclbinres.CRC32Value {
				os.Remove(tempDLLFileName)
			}
		}
	}
	if !fileExists(tempDLLFileName) {
		if err := zlibUnCompressToFile(tempDLLFileName, liblclbinres.LCLBinRes); err != nil {
			if os.Remove(tempDLLFileName) != nil {
				return "", false
			}
		}
	}
	return tempDLLFileName, true
}

func zlibUnCompressToFile(destFileName string, input []byte) error {
	r, err := zlib.NewReader(bytes.NewReader(input))
	if err != nil {
		return err
	}
	defer r.Close()
	fi, err := os.Create(destFileName)
	if err != nil {
		return err
	}
	defer fi.Close()
	_, err = io.Copy(fi, r)
	if err != nil {
		return nil
	}
	return nil
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
