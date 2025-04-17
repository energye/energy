//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package i18n Multilingual resources
//
//	Resource usage file loading
//	File name format: locale.[lang].json | locale.[lang].ini => locale.en-US.json | locale.zh-CN.ini
package i18n

import (
	"encoding/json"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/energy/emfs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var (
	currentLang  consts.LANGUAGE
	resources    map[string]string             // 资源-优先
	resourcesVar = make(map[string]*string, 0) // 变量资源-切换时被同步
	lPath        string                        // 本地加载
	lFS          emfs.IEmbedFS                 // 内置加载-优先
	lFSPath      string                        // 内置加载-资源所在目录
)

// SetLocalPath
//
//	设置本地资源在本地目录加载
func SetLocalPath(localPath string) {
	lPath = localPath
}

// SetLocalFS
//
//	设置本地资源在FS中加载-优先
//	lFS: 内置对象
//	lFSPath: 资源所在目录 to/path
func SetLocalFS(localFS emfs.IEmbedFS, localFSPath string) {
	lFS = localFS
	lFSPath = localFSPath
}

// Switch
//
//	 切换语言
//		默认先在内置FS中加载
func Switch(lang consts.LANGUAGE) {
	if currentLang == lang {
		// 如果当前语言和切换语言一样
		return
	}
	currentLang = lang
	resources = make(map[string]string, 0)
	if lang == consts.LANGUAGE_zh_CN {
		initZhCn()
	} else {
		initEnUs()
	}
	// 在json文件中加载资源
	loadJSONConvertResource()
	// 在ini文件中加载资源
	loadINIConvertResource()
	// 加载后的资源同步到变量资源
	if len(resources) > 0 {
		for name, value := range resourcesVar {
			if v, ok := resources[name]; ok {
				*value = v
			}
		}
	}
}

// loadJSONConvertResource
//
//	加载JSON格式并转换资源
func loadJSONConvertResource() bool {
	jsonFileName := "locale." + string(currentLang) + ".json"
	//加载资源
	if contentBytes := loadResource(jsonFileName); contentBytes != nil {
		var temp interface{} // map
		if json.Unmarshal(contentBytes, &temp) == nil {
			if v, ok := temp.(map[string]interface{}); ok {
				for name, value := range v {
					resources[name] = value.(string)
				}
				return true
			}
		}
	}
	return false
}

// loadINIConvertResource
//
//	加载INI格式并转换资源
func loadINIConvertResource() bool {
	iniFileName := "locale." + string(currentLang) + ".ini"
	//加载资源
	if contentBytes := loadResource(iniFileName); contentBytes != nil {
		var temp = string(contentBytes)
		res := strings.Split(temp, "\r\n")
		for _, line := range res {
			nv := strings.Split(line, "=")
			if len(nv) == 2 {
				resources[nv[0]] = nv[1]
			}
		}
		return true
	}
	return false
}

// loadResource
//
//	加载资源-优先在内置FS中加载
func loadResource(fileName string) []byte {
	//加载资源
	if lFS != nil {
		var path string
		if lFSPath != "" {
			path = lFSPath + "/" + fileName
		} else {
			path = fileName
		}
		if contentBytes, err := lFS.ReadFile(path); err == nil {
			return contentBytes
		}
	} else if lPath != "" {
		path := filepath.Join(lPath, fileName)
		if contentBytes, err := ioutil.ReadFile(path); err == nil {
			return contentBytes
		}
	}
	return nil
}

// RegisterResource
//
//	注册资源, 在代码中手动设置在静态资源
func RegisterResource(name, value string) {
	resources[name] = value
}

// RegisterVarResource
//
//	注册变量资源, 在代码中手动设置, 切换资源时变量值会被同步
func RegisterVarResource(name string, value *string) {
	resourcesVar[name] = value
}

// Resource
//
//	返回资源-优先在静态资源中查找
func Resource(name string) string {
	if v, ok := resources[name]; ok {
		return v
	} else if v, ok := resourcesVar[name]; ok {
		return *v
	}
	return ""
}
