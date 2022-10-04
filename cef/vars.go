//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"errors"
	"fmt"
	"github.com/energye/golcl/tools/homedir"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"unsafe"
)

var (
	nullptr        = unsafe.Pointer(uintptr(0)) //
	Separator      = string(filepath.Separator) //
	commonRootName = "gocobj"                   //ICEFv8Value 通用类型变量属性的所属默认对象名称
	objectRootName = "goobj"                    //ICEFv8Value 对象类型变量属性的所属默认对象名称
	ExePath        string                       //执行文件目录
	DownloadsDir   string                       //下载目录
	enableGPU      = false                      //启用GPU true启动 false不启用
	SingleProcess  = false                      //进程启动方式, true单进程 false多进程
	processName    PRCESS_TYPE                  //进程名称
	commonInstance = &CommonInstance{}          //通用实例
	homeDir, _     = homedir.Dir()              //系统用户目录
)

type initBindVariableCallback func(browser *ICefBrowser, frame *ICefFrame, bind IProvisionalBindStorage)

// 变量绑定
var VariableBind = &variableBind{bindStorage: &provisionalBindStorage{}, valuesBind: make(map[string]JSValue)}

func init() {
	ExePath = os.Args[0]
	ExePath = ExePath[:strings.LastIndex(ExePath, Separator)]
	ipcSock = fmt.Sprintf("%s%sgolcl%s%s", homeDir, Separator, Separator, memoryAddress)
}

type IProvisionalBindStorage interface {
	NewString(name, value string) *JSString
	NewInteger(name string, value int32) *JSInteger
	NewDouble(name string, value float64) *JSDouble
	NewBoolean(name string, value bool) *JSBoolean
	NewNull(name string) *JSNull
	NewUndefined(name string) *JSUndefined
	NewFunction(name string, fn interface{}) error
	NewObjects(objects ...interface{})
}

type provisionalBindStorage struct {
}

type variableBind struct {
	bindStorage              IProvisionalBindStorage  //
	initBindVariableCallback initBindVariableCallback //
	valuesBind               map[string]JSValue       //所有变量属性或函数的信息集合
}

func (m *variableBind) putValueBind(fullName string, value JSValue) {
	if _, ok := m.valuesBind[fullName]; !ok {
		m.valuesBind[fullName] = value
	}
}

func (m *variableBind) ValueBindCount() int {
	return len(m.valuesBind)
}

func (m *variableBind) getValueBind(fullName string) (JSValue, bool) {
	value, ok := m.valuesBind[fullName]
	return value, ok
}

func clearValueBind() {
	//valuesBind = make(map[uintptr]JSValue)
}

// Javascript 和 Go的函数和变量绑定声明初始函数
//
// 在JS中访问Go中的（函数,变量）或Go中访问JS的（函数,变量）需要在此回调函数中初始化
//
// 主进程回调 browser 和 frame 为 nil
func (m *variableBind) VariableCreateCallback(callback func(browser *ICefBrowser, frame *ICefFrame, bind IProvisionalBindStorage)) {
	m.initBindVariableCallback = callback
}

// 调用变量绑定回调函数
//
// 在主进程和渲染进程创建时调用
func (m *variableBind) callVariableBind(browser *ICefBrowser, frame *ICefFrame) {
	if m.initBindVariableCallback != nil {
		m.initBindVariableCallback(browser, frame, m.bindStorage)
	}
}

// ICEFv8Value NewString
func (m *provisionalBindStorage) NewString(name, value string) *JSString {
	jsValueBind := new(JSString)
	jsValueBind.rwLock = new(sync.Mutex)
	jsValueBind.name = name
	jsValueBind.value = value
	jsValueBind.valueType = V8_VALUE_STRING
	VariableBind.valueHandle(jsValueBind)
	return jsValueBind
}

// ICEFv8Value NewInteger
func (m *provisionalBindStorage) NewInteger(name string, value int32) *JSInteger {
	jsValueBind := new(JSInteger)
	jsValueBind.rwLock = new(sync.Mutex)
	jsValueBind.name = name
	jsValueBind.value = value
	jsValueBind.valueType = V8_VALUE_INT
	VariableBind.valueHandle(jsValueBind)
	return jsValueBind
}

// ICEFv8Value NewDouble
func (m *provisionalBindStorage) NewDouble(name string, value float64) *JSDouble {
	jsValueBind := new(JSDouble)
	jsValueBind.rwLock = new(sync.Mutex)
	jsValueBind.name = name
	jsValueBind.value = value
	jsValueBind.valueType = V8_VALUE_DOUBLE
	VariableBind.valueHandle(jsValueBind)
	return jsValueBind
}

// ICEFv8Value NewBool
func (m *provisionalBindStorage) NewBoolean(name string, value bool) *JSBoolean {
	jsValueBind := new(JSBoolean)
	jsValueBind.rwLock = new(sync.Mutex)
	jsValueBind.name = name
	jsValueBind.value = value
	jsValueBind.valueType = V8_VALUE_BOOLEAN
	VariableBind.valueHandle(jsValueBind)
	return jsValueBind
}

// ICEFv8Value NewNull
func (m *provisionalBindStorage) NewNull(name string) *JSNull {
	jsValueBind := new(JSNull)
	jsValueBind.rwLock = new(sync.Mutex)
	jsValueBind.name = name
	jsValueBind.value = "null"
	jsValueBind.valueType = V8_VALUE_NULL
	VariableBind.valueHandle(jsValueBind)
	return jsValueBind
}

// ICEFv8Value NewUndefined
func (m *provisionalBindStorage) NewUndefined(name string) *JSUndefined {
	jsValueBind := new(JSUndefined)
	jsValueBind.rwLock = new(sync.Mutex)
	jsValueBind.name = name
	jsValueBind.value = "undefined"
	jsValueBind.valueType = V8_VALUE_UNDEFINED
	VariableBind.valueHandle(jsValueBind)
	return jsValueBind
}

// ICEFv8Value NewFunction
//
// 绑定 定义的普通函数 prefix default struct name
func (m *provisionalBindStorage) NewFunction(name string, fn interface{}) error {
	if GOValueReflectType(fn) == GO_VALUE_FUNC {
		if info, err := checkFunc(reflect.TypeOf(fn), FN_TYPE_COMMON); err == nil {
			jsValueBind := new(JSFunction)
			jsValueBind.rwLock = new(sync.Mutex)
			jsValueBind.name = name
			jsValueBind.value = fn
			jsValueBind.valueType = V8_VALUE_FUNCTION
			jsValueBind.funcInfo = info
			VariableBind.valueHandle(jsValueBind)
			return nil
		} else {
			return err
		}
	}
	return errors.New("创建的函数不是函数类型")
}

// ICEFv8Value NewObjects
//
// 对象类型变量和对象函数绑定
func (m *provisionalBindStorage) NewObjects(objects ...interface{}) {
	bindObject(objects...)
}

// ICEFv8Value valueHandle
func (m *variableBind) valueHandle(jsValue JSValue) {
	jsValue.setPtr(unsafe.Pointer(&jsValue))
	jsValue.setInstance(uintptr(jsValue.Ptr()))
	jsValue.setEventId(uintptr(__bind_id()))
	jsValue.setThat(jsValue)
	VariableBind.putValueBind(jsValue.Name(), jsValue)
}
