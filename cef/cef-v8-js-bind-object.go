//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// V8 JSValue 对象类型实现
package cef

import (
	"bytes"
	"fmt"
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/imports"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl/api"
	"reflect"
	"strconv"
	"unsafe"
)

var objectSti = new(structTypeInfo)

// 函数详情 1.参数个数 2.每个参数类型 3.返回参数类型
type funcInfo struct {
	InNum          int32   `json:"inNum"`
	InParam        []*vt   `json:"inParam"`
	OutNum         int32   `json:"outNum"`
	OutParam       []*vt   `json:"outParam"`
	OutParamErrIdx int32   `json:"outParamErrIdx"`
	OutParamIdx    int32   `json:"outParamIdx"`
	FnType         FN_TYPE `json:"fnType"`
}

type vt struct {
	Jsv V8_JS_VALUE_TYPE `json:"jsv"`
	Gov GO_VALUE_TYPE    `json:"gov"`
}

// 结构类型信息，结构的字段信息和方法函数信息
type structTypeInfo struct {
	StructsObject map[string]*structObjectInfo //分析后的 有关系的结构信息
}

// 分析后的 结构信息
type structObjectInfo struct {
	Instance            uintptr                      `json:"instance"`
	ParentInstance      uintptr                      `json:"parentInstance"`
	Parent              *structObjectInfo            `json:"-"`
	Ptr                 unsafe.Pointer               `json:"-"`
	ObjName             string                       `json:"objName"` //对象名称
	FullObjName         string                       `json:"fullObjName"`
	FieldsInfo          map[string]*structFieldInfo  `json:"fieldsInfo"` //字段信息 key=字段名 value=字段类型
	FuncsInfo           map[string]*structFuncInfo   `json:"funcsInfo"`
	SubStructObjectInfo map[string]*structObjectInfo `json:"subStructObjectInfo"` //子对象信息
}

// 结构的字段信息
type structFieldInfo struct {
	EventId    uintptr        `json:"event_id"`
	ValueType  *vt            `json:"valueType"` //字段类型，go 和 js
	FieldValue *reflect.Value `json:"-"`         //用于字段值修改和获取
}

type structFuncInfo struct {
	EventId uintptr `json:"event_id"`
	*funcInfo
	Method reflect.Value `json:"-"`
}

type cefObject struct {
	Id          uintptr
	ParentId    uintptr
	Name        uintptr //string
	FullObjName uintptr //string
	FieldLen    uintptr
	Fields      uintptr //array
	FuncLen     uintptr
	Funcs       uintptr //array
}

func (m *vt) ToValueTypeString() string {
	govs := common.FuncParamGoTypeStr(m.Gov)
	jsvs := common.FuncParamJsTypeStr(m.Jsv)
	return fmt.Sprintf("GO=%s,JS:=%s", govs, jsvs)
}

// IsGoIntAuto 判断Go 所有 int 类型
func (m *vt) IsGoIntAuto() bool {
	switch m.Gov {
	case GO_VALUE_INT, GO_VALUE_INT8, GO_VALUE_INT16, GO_VALUE_INT32, GO_VALUE_INT64, GO_VALUE_UINT, GO_VALUE_UINT8, GO_VALUE_UINT16, GO_VALUE_UINT32, GO_VALUE_UINT64:
		return true
	}
	return false
}

// IsGoFloatAuto 判断Go 所有 float 类型
func (m *vt) IsGoFloatAuto() bool {
	switch m.Gov {
	case GO_VALUE_FLOAT32, GO_VALUE_FLOAT64:
		return true
	}
	return false
}

// ICefV8Context BindObject
// 对应Go，不支持字段的类型修改（包括对象类型）,不支持删除和增加字段变更，支持字段值修改。和获取。
func bindObject(objects ...interface{}) {
	objectSti.StructsObject = make(map[string]*structObjectInfo, len(objects))
	for i := 0; i < len(objects); i++ {
		object := objects[i]
		typ := reflect.TypeOf(object)
		if typ.Kind() == reflect.Ptr {
			var value = reflect.ValueOf(object)
			var objTyp = reflect.ValueOf(object).Type().Elem()
			objectSti.StructsObject[objTyp.Name()] = &structObjectInfo{FieldsInfo: make(map[string]*structFieldInfo), SubStructObjectInfo: make(map[string]*structObjectInfo)}
			objectSti.StructsObject[objTyp.Name()].analysisObjectField(objTyp, typ, value)
		} else {
			logger.Error("结构对象非指针类型:", typ.Name())
		}
	}
	objectSti._objectToCefObject()
}

// _objectToCefObject
// 对象字段和函数统计
// 对象转换 go to cef
func (m *structTypeInfo) _objectToCefObject() {
	for _, info := range m.StructsObject {
		m._subInfoToCefObject(info)
	}
}

func (m *structTypeInfo) _subInfoToCefObject(info *structObjectInfo) {
	m._infoTo(info)
	if len(info.SubStructObjectInfo) > 0 {
		for _, subInfo := range info.SubStructObjectInfo {
			m._subInfoToCefObject(subInfo)
		}
	}
}

func (m *structTypeInfo) _infoTo(info *structObjectInfo) {
	var (
		fieldLen = len(info.FieldsInfo)
		fieldPtr uintptr
		funcLen  = len(info.FuncsInfo)
		funcPtr  uintptr
		cofs     []*valueBindInfo
		fns      []*valueBindInfo
		i        = 0
	)
	if fieldLen > 0 {
		//字段
		cofs = make([]*valueBindInfo, fieldLen, fieldLen)
		for fieldName, fi := range info.FieldsInfo {
			cofs[i] = &valueBindInfo{
				Name:     api.PascalStr(fieldName),
				EventId:  fi.EventId,
				BindType: uintptr(fi.ValueType.Jsv),
			}
			i++
			m.createObjectFieldVariable(info.FullObjName, fieldName, fi)
		}
		fieldPtr = uintptr(unsafe.Pointer(&cofs[0]))
		i = 0

	}
	if funcLen > 0 {
		fns = make([]*valueBindInfo, funcLen, funcLen)
		for fnName, fn := range info.FuncsInfo {
			var inParamBuf bytes.Buffer
			var outParamBuf bytes.Buffer
			for j, inParamType := range fn.InParam {
				if j > 0 {
					inParamBuf.WriteString(",")
				}
				inParamBuf.WriteString(strconv.Itoa(int(inParamType.Jsv)))
			}
			for j, outParanType := range fn.OutParam {
				if j > 0 {
					outParamBuf.WriteString(",")
				}
				outParamBuf.WriteString(strconv.Itoa(int(outParanType.Jsv)))
			}
			fns[i] = &valueBindInfo{
				Name:           api.PascalStr(fnName),
				EventId:        fn.EventId,
				BindType:       uintptr(V8_VALUE_FUNCTION),
				FnInNum:        uintptr(fn.InNum),
				FnInParamType:  api.PascalStr(inParamBuf.String()),
				FnOutNum:       uintptr(fn.OutNum),
				FnOutParamType: api.PascalStr(outParamBuf.String()),
			}
			i++
			m.createObjectFuncVariable(info.FullObjName, fnName, fn)
		}
		funcPtr = uintptr(unsafe.Pointer(&fns[0]))
	}
	co := &cefObject{
		Id:          info.Instance,
		ParentId:    info.ParentInstance,
		Name:        api.PascalStr(info.ObjName),
		FullObjName: api.PascalStr(info.FullObjName),
		FieldLen:    uintptr(fieldLen),
		Fields:      fieldPtr,
		FuncLen:     uintptr(funcLen),
		Funcs:       funcPtr,
	}
	imports.Proc(internale_CEFV8ValueRef_ObjectValueBindInfo).Call(uintptr(unsafe.Pointer(co)))
}

// 创建 结构对象的字段变量
func (m *structTypeInfo) createObjectFieldVariable(fullParentName, fieldName string, sfi *structFieldInfo) {
	newV8Value(sfi.EventId, fullParentName, fieldName, sfi.FieldValue, nil, sfi.ValueType.Jsv, IS_OBJECT)
}

// 创建 结构对象的函数变量
func (m *structTypeInfo) createObjectFuncVariable(fullParentName, funcName string, sfi *structFuncInfo) {
	newV8Value(sfi.EventId, fullParentName, funcName, nil, sfi, V8_VALUE_FUNCTION, IS_OBJECT)
}

// 分析对象的字段
func (m *structObjectInfo) analysisObjectField(typ reflect.Type, typPtr reflect.Type, value reflect.Value) {
	if m.Parent == nil {
		m.ObjName = typ.Name()
		m.FullObjName = m.ObjName
		m.FullObjName = m.ObjName
	}
	m.Instance = value.Elem().UnsafeAddr()
	m.Ptr = unsafe.Pointer(m.Instance)
	//字段信息遍历
	for i := 0; i < typ.NumField(); i++ {
		field := value.Elem().Field(i)
		fieldType := field.Type()
		fieldName := typ.Field(i).Name
		var b, isPrt = false, false
		//取出类型，同时判断出是循环引用的类型
		if fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
			isPrt = true
		}
		//结构对象,循环引用的对象不被支持
		if isPrt && fieldType.Kind() == reflect.Struct && !field.IsZero() {
			subSoi := &structObjectInfo{FieldsInfo: make(map[string]*structFieldInfo), SubStructObjectInfo: make(map[string]*structObjectInfo)}
			subSoi.Parent = m
			subSoi.ParentInstance = m.Instance
			subSoi.ObjName = fieldName
			subSoi.FullObjName = fmt.Sprintf("%s.%s", subSoi.Parent.FullObjName, subSoi.ObjName)
			m.SubStructObjectInfo[fieldName] = subSoi
			//是结构对象分析
			if isPrt {
				subSoi.analysisObjectField(fieldType, field.Type(), field)
			} else {
				subSoi.analysisObjectField(fieldType, fieldType, field)
			}
			b = true
		} else if fieldType.Kind() == reflect.Slice { //数组

		} else if typ.Field(i).Type.Kind() != reflect.Ptr && fieldType.Kind() != reflect.Struct && typ.Field(i).IsExported() {
			//过滤掉指针类型和非导出大写字段，和不是结构类型
			b = true
		} else if fieldType.Kind() == reflect.Struct && field.IsZero() {
			logger.Debug("字段类型-对象,", fieldName, " 未初始化, 忽略JS绑定映射.")
		}
		if b { //b=true可以正常解析映射
			filedValue := value.Elem().FieldByName(fieldName)
			t := fieldType.Kind().String()
			m.FieldsInfo[fieldName] = &structFieldInfo{
				EventId: uintptr(__bind_id()),
				ValueType: &vt{
					Jsv: common.JSValueType(t),
					Gov: common.GOValueType(t),
				},
				FieldValue: &filedValue,
			}
		}
	}
	m.analysisObjectMethod(typPtr, value)
}

// 分析对象的函数方法
// 不符合js类型的函数的参数，不会被解析成js函数
func (m *structObjectInfo) analysisObjectMethod(typPtr reflect.Type, value reflect.Value) {
	m.FuncsInfo = make(map[string]*structFuncInfo)
	for idx := 0; idx < typPtr.NumMethod(); idx++ {
		method := typPtr.Method(idx)
		if method.IsExported() {
			if fi, err := checkFunc(method.Type, FN_TYPE_OBJECT); err == nil {
				methodValue := value.MethodByName(method.Name)
				m.FuncsInfo[method.Name] = &structFuncInfo{
					EventId:  uintptr(__bind_id()),
					funcInfo: fi,
					Method:   methodValue,
				}
			} else {
				panic(err.Error())
			}
		}
	}
}
