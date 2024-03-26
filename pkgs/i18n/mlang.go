//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// 多语言包，用于本地化操作

package i18n

import (
	"encoding/json"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/lcl/rtl"
	"github.com/energye/energy/v2/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// TLangItem
//
// 本地已经存在语言列表的项目定义
type TLangItem struct {
	Language struct {
		Id          int    // 2052
		Name        string // zh-CN
		Description string // 简体中文
		Author      string // ying32
		AuthorEmail string // 1444386932@qq.com

	} `json:"!language"`
}

var (
	// :: public

	// LocalLangs
	//
	// 本地已经添加了的语言列表。
	//
	// List of languages that have been added locally.
	LocalLangs = make(map[int]TLangItem, 0)

	// AppNodeName
	//
	// 默认应用的节点名称
	//
	// Node name applied by default.
	AppNodeName string

	// CurrentLang
	//
	// 当前语言
	//
	// Current language
	CurrentLang string

	// :: private

	//  语言存放目录
	//  language storage directory
	langsPath = extractFilePath(os.Args[0]) + "Langs" + string(filepath.Separator)

	//  强制显示的语言文件名
	//  Forced language filename to be displayed
	langSetFileName = langsPath + "lang.s"

	// 公共资源
	// Public resource
	commonResouces map[string]string

	// 当前app资源
	// Current app resources
	appResouces map[string]string

	// lib中的资源
	// Resources in the library
	libResouces map[string]string

	// 当前app节点信息
	// Current app node information
	appNode map[string]interface{}

	// 已经注册的Form
	// Registered Form
	regForms = make(map[uintptr]lcl.IComponent, 0)

	// 需要注册的资源
	// Resources that require registration
	regResources = make(map[string]*string, 0)

	// lib中注册的资源
	// Resources registered in the library
	regLibResources []types.TLibResource

	// 修改lib中资源的函数
	// Functions that modify resources in the library
	modifyLibResourceFN func(aPtr uintptr, aValue string)
)

func extractFilePath(path string) string {
	filename := filepath.Base(path)
	return path[:len(path)-len(filename)]
}

func parseLangFile(lang string) {
	filename := langsPath + lang + ".lang"
	if bs, err := ioutil.ReadFile(filename); err == nil {
		var temp interface{}
		if json.Unmarshal(bs, &temp) == nil {
			// 公共资源
			commonResouces = make(map[string]string, 0)
			appResouces = make(map[string]string, 0)
			libResouces = make(map[string]string, 0)

			appNode = make(map[string]interface{}, 0)

			// 共享资源
			if v, ok := temp.(map[string]interface{}); ok {
				for key, val := range v["!resources"].(map[string]interface{}) {
					commonResouces[key] = val.(string)
				}
			}

			// 共享的Lib中的资源
			if v, ok := temp.(map[string]interface{}); ok {
				for key, val := range v["!libresources"].(map[string]interface{}) {
					libResouces[key] = val.(string)
				}
			}

			// 当前app资源
			if v, ok := temp.(map[string]interface{}); ok {
				if node, ok := v[strings.ToLower(AppNodeName)]; ok {
					appNode = node.(map[string]interface{})
					if v, ok := appNode["!resources"]; ok {
						for key, val := range v.(map[string]interface{}) {
							appResouces[key] = val.(string)
						}
					}
				}
			}
		}
	}
}

// 翻译资源，这里不翻译UI上的资源，只是一些常量什么的
func translateStrings() {
	// 这里先翻译lib中的资源
	if len(regLibResources) > 0 && len(libResouces) > 0 {
		for _, item := range regLibResources {
			if v, ok := libResouces[item.Name]; ok {
				if modifyLibResourceFN != nil {
					modifyLibResourceFN(item.Ptr, v)
				}
			}
		}
	}

	// 没有待翻译的，不进行翻译
	if len(commonResouces) > 0 || len(appResouces) > 0 {
		for key, val := range regResources {
			if v, ok := appResouces[key]; ok {
				*val = v
			} else {
				if v, ok := commonResouces[key]; ok {
					*val = v
				}
			}
		}
	}
}

// InitDefaultLang
//  初始默认语言
//  Initial default language.
func InitDefaultLang() {
	slang := ReadSetLang()
	if slang != "" {
		ChangeLang(slang)
		return
	}
	if v, ok := LocalLangs[int(rtl.SysLocale.DefaultLCID)]; ok {
		ChangeLang(v.Language.Name)
	}
}

// ReadSetLang
//  读当前强制显示语言
//  Read the current mandatory display language.
func ReadSetLang() string {
	bs, err := ioutil.ReadFile(langSetFileName)
	if err == nil {
		return string(bs)
	}
	return ""
}

// WriteSetLang
//  写入强显示制语言
//  Write mandatory language.
func WriteSetLang(lang string) {
	ioutil.WriteFile(langSetFileName, []byte(lang), 0775)
}

// ChangeLang
//  改变语言
//  Change language.
func ChangeLang(lang string) {
	if lang == CurrentLang {
		return
	}
	CurrentLang = lang
	if AppNodeName == "" {
		AppNodeName = filepath.Base(os.Args[0])
		AppNodeName = AppNodeName[:len(AppNodeName)-len(filepath.Ext(AppNodeName))]
	}
	parseLangFile(CurrentLang)
	// 翻译语言
	translateStrings()

	// 重新翻译已注册的TForm
	for _, c := range regForms {
		InitComponentLang(c)
	}
}

// IdRes
//  通过key查询当前资源中的，顺序为当前app资源 -> 共享资源 -> lib资源
//  Query the current resources by key, the order is current app resources -> shared resources -> lib resources.
func IdRes(key string) string {
	if v, ok := appResouces[key]; ok {
		return v
	}
	if v, ok := commonResouces[key]; ok {
		return v
	}
	if v, ok := libResouces[key]; ok {
		return v
	}
	return ""
}

// InitComponentLang
//  初始一个Form的语言
//  The language of the initial Form.
func InitComponentLang(aOwner lcl.IComponent) {
	ptr := lcl.CheckPtr(aOwner)
	if ptr == 0 {
		return
	}
	if _, ok := regForms[ptr]; !ok {
		regForms[ptr] = aOwner
	}
	if node, ok := appNode[aOwner.Name()]; ok {
		for propName, propValue := range node.(map[string]interface{}) {
			propName = strings.Trim(propName, " ")
			propValue, _ := propValue.(string)
			if strings.Contains(propName, ".") {
				arr := strings.Split(propName, ".")
				if len(arr) > 1 {
					obj := aOwner.FindComponent(arr[0])
					if obj != nil && obj.IsValid() {
						switch len(arr) {
						case 2:
							rtl.SetPropertyValue(obj.Instance(), arr[1], propValue)
						case 3:
							rtl.SetPropertySecValue(obj.Instance(), arr[1], arr[2], propValue)
						}
					}
				}
			} else {
				rtl.SetPropertyValue(ptr, propName, propValue)
			}
		}
	}
}

// RegisterVarString
//  注册需要翻译的字符
//  Register characters that need to be translated.
func RegisterVarString(name string, value *string) {
	regResources[name] = value
}

func initLoadLocalLangsInfo() {
	_, err := os.Stat(langsPath)
	if os.IsNotExist(err) {
		return
	}
	filepath.Walk(langsPath, func(path string, info os.FileInfo, err error) error {
		if strings.ToLower(filepath.Ext(info.Name())) == ".lang" {
			bs, err := ioutil.ReadFile(path)
			if err == nil {
				item := TLangItem{}
				if json.Unmarshal(bs, &item) == nil {
					LocalLangs[item.Language.Id] = item
				}
			}
		}
		return nil
	})
}

func I18NInit() {
	// 首先设置lib中资源
	regLibResources = rtl.GetLibResourceItems()
	modifyLibResourceFN = rtl.ModifyLibResource
	initLoadLocalLangsInfo()
}
