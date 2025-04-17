//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package tools

import (
	"bytes"
	"github.com/cyber-xxm/energy/v2/cmd/internal/consts"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"text/template"
)

var (
	exePath           string
	currentExecuteDir string
)

func init() {
	// 执行文件所在目录
	exePath, _ = filepath.Split(os.Args[0])
	// 当前执行目录，在其它目录执行目标执行文件时，返回当前执行目录
	currentExecuteDir, _ = os.Getwd()
}

func FixCMDName(name string) string {
	if consts.IsWindows {
		name += ".exe"
	}
	return name
}

// CommandExists 命令是否存在
func CommandExists(name string) bool {
	_, err := exec.LookPath(name)
	if err != nil {
		return false
	}
	return true
}

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	return v.(string)
}

func StrToInt64(value string) int64 {
	v, _ := strconv.ParseInt(value, 10, 64)
	return v
}
func StrToFloat64(value string) float64 {
	v, _ := strconv.ParseFloat(value, 64)
	return v
}

func StrToFloat32(value string) float32 {
	v, _ := strconv.ParseFloat(value, 32)
	return float32(v)
}

var numberReg = regexp.MustCompile("^\\d+$")

func IsInt(v string) bool {
	return numberReg.MatchString(v)
}

// 验证发行版本，入参: vx.x.x
func VerifyRelease(v string) bool {
	if v == "" {
		return false
	}
	// 验证版本号格式
	tmpVers := strings.Split(v[1:], ".")
	if len(tmpVers) != 3 {
		return false
	}
	if !IsInt(tmpVers[0]) || !IsInt(tmpVers[1]) || !IsInt(tmpVers[2]) {
		return false
	}
	return true
}

func ToInt(v interface{}) int {
	switch v.(type) {
	case string:
		return int(StrToInt64(v.(string)))
	case float32:
		return int(math.Round(float64(StrToFloat32(v.(string)))))
	case float64:
		return int(math.Round(StrToFloat64(v.(string))))
	case bool:
		if v.(bool) {
			return 1
		} else {
			return 0
		}
	case int:
		return v.(int)
	case int8:
		return int(v.(int8))
	case int16:
		return int(v.(int16))
	case int32:
		return int(v.(int32))
	case int64:
		return int(v.(int64))
	case uintptr:
		return int(v.(uintptr))
	default:
		return 0
	}
}

func Equals(s1, s2 string) bool {
	return strings.ToLower(s1) == strings.ToLower(s2)
}

func ToRNilString(v interface{}, new string) string {
	if v == nil {
		return new
	}
	return v.(string)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		} else if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

func IsExistAndSize(path string, size int64) bool {
	fInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else if os.IsExist(err) && fInfo.Size() == size {
			return true
		}
		return false
	}
	return fInfo.Size() == size
}

func RenderTemplate(templateText string, data map[string]interface{}) ([]byte, error) {
	tmpl, err := template.New("").Parse(templateText)
	if err != nil {
		return nil, err
	}
	var out bytes.Buffer
	if err = tmpl.Execute(&out, data); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// Compare compare2 < compare1 = true
func Compare(compare1, compare2 string) bool {
	if compare1[0] == 'v' {
		compare1 = compare1[1:]
	}
	if compare2[0] == 'v' {
		compare2 = compare2[1:]
	}
	compare1 = strings.Split(compare1, "-")[0]
	compare2 = strings.Split(compare2, "-")[0]
	cv, _ := strconv.Atoi(strings.Replace(compare1, ".", "", -1))
	ev, _ := strconv.Atoi(strings.Replace(compare2, ".", "", -1))
	if ev < cv {
		return true
	}
	return false
}

// ExePath
//
//	返回当前执行文件路径
func ExePath() string {
	return exePath
}

// CurrentExecuteDir
//
//	返回当前执行目录
func CurrentExecuteDir() string {
	return currentExecuteDir
}

var platformExtNames = map[string]string{
	"windows": ".dll",
	"linux":   ".so",
	"darwin":  ".dylib",
}

func GetDLLName() string {
	libName := "liblcl"
	if ext, ok := platformExtNames[runtime.GOOS]; ok {
		return libName + ext
	}
	return libName
}
