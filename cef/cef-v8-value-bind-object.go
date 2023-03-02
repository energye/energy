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

// 对象变量绑定
var objectTI = &objectTypeInfo{ObjectInfos: make(map[string]*objectInfo), FieldInfos: make(map[string]*fieldInfo), FunctionInfos: make(map[string]*functionInfo)}

// 函数详情 1.参数个数 2.每个参数类型 3.返回参数类型
type funcInfo struct {
	InNum          int32   `json:"inNum"`          //入参个数
	InParam        []*VT   `json:"inParam"`        //入参类型
	OutNum         int32   `json:"outNum"`         //出参个数
	OutParam       []*VT   `json:"outParam"`       //出参类型
	OutParamErrIdx int32   `json:"outParamErrIdx"` //出参错误位置
	OutParamIdx    int32   `json:"outParamIdx"`    //出参位置
	FnType         FN_TYPE `json:"fnType"`         //函数类型, 直接定义函数或对象函数
}

// VT 值类型
type VT struct {
	Jsv V8_JS_VALUE_TYPE `json:"jsv"`
	Gov GO_VALUE_TYPE    `json:"gov"`
}

// objectTypeInfo 结构类型描述，结构的字段描述和方法函数描述
type objectTypeInfo struct {
	isBind        bool                     //是否已绑定过
	ObjectInfos   map[string]*objectInfo   `json:"objectInfos"`   //分析后有关系的结构描述
	FieldInfos    map[string]*fieldInfo    `json:"fieldInfos"`    //字段描述 key=字段名 value=字段描述
	FunctionInfos map[string]*functionInfo `json:"functionInfos"` //函数描述 key=函数名 value=函数描述
}

// objectInfo 分析后的 结构描述
type objectInfo struct {
	isCreateJSValue bool                     //是否已创建 JSValue
	Instance        uintptr                  `json:"instance"`       //对象地址
	ParentInstance  uintptr                  `json:"parentInstance"` //父对象地址
	Parent          *objectInfo              `json:"-"`              //父对象结构描述
	Ptr             unsafe.Pointer           `json:"-"`              //对象指针
	ObjName         string                   `json:"objName"`        //对象名称
	FullObjName     string                   `json:"fullObjName"`    //对象全路径名称: structName.FieldName.XXX
	FieldInfos      map[string]*fieldInfo    `json:"fieldInfos"`     //字段描述 key=字段名 value=字段描述
	FunctionInfos   map[string]*functionInfo `json:"functionInfos"`  //函数描述 key=函数名 value=函数描述
	SubObjectInfo   map[string]*objectInfo   `json:"subObjectInfo"`  //子对象描述
}

// fieldInfo 字段描述
type fieldInfo struct {
	EventId    uintptr     `json:"event_id"`
	ValueType  *VT         `json:"valueType"` //字段类型，go 和 js
	FieldValue interface{} `json:"-"`         //用于字段值修改和获取 *reflect.Value | *field
}

// functionInfo 结构的函数描述
type functionInfo struct {
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

func (m *VT) ToString() string {
	gov := common.FuncParamGoTypeStr(m.Gov)
	jsv := common.FuncParamJsTypeStr(m.Jsv)
	return fmt.Sprintf("GO=%s JS=%s", gov, jsv)
}

// IsGoIntAuto 判断Go 所有 int 类型
func (m *VT) IsGoIntAuto() bool {
	switch m.Gov {
	case GO_VALUE_INT, GO_VALUE_INT8, GO_VALUE_INT16, GO_VALUE_INT32, GO_VALUE_INT64, GO_VALUE_UINT, GO_VALUE_UINT8, GO_VALUE_UINT16, GO_VALUE_UINT32, GO_VALUE_UINT64:
		return true
	}
	return false
}

// IsGoFloatAuto 判断Go 所有 float 类型
func (m *VT) IsGoFloatAuto() bool {
	switch m.Gov {
	case GO_VALUE_FLOAT32, GO_VALUE_FLOAT64:
		return true
	}
	return false
}

// bindObject ICefV8Context
// 对应Go，不支持字段的类型修改（包括对象类型）,不支持删除和增加字段变更，支持字段值修改。和获取。
func bindObject(objects ...interface{}) {
	for i := 0; i < len(objects); i++ {
		object := objects[i]
		typ := reflect.TypeOf(object)
		if typ.Kind() == reflect.Ptr {
			var value = reflect.ValueOf(object)
			var objTyp = reflect.ValueOf(object).Type().Elem()
			objectTI.ObjectInfos[objTyp.Name()] = &objectInfo{FieldInfos: make(map[string]*fieldInfo), SubObjectInfo: make(map[string]*objectInfo)}
			objectTI.ObjectInfos[objTyp.Name()].analysisObjectField(objTyp, typ, value)
		} else {
			logger.Error("结构对象非指针类型:", typ.Name())
		}
	}
	objectTI.objectToCefObject()
}

// objectToCefObject
// 对象转换 go to v8
func (m *objectTypeInfo) objectToCefObject() {
	for _, info := range m.ObjectInfos {
		m.subInfoToCefObject(info)
	}
}

//subInfoToCefObject
func (m *objectTypeInfo) subInfoToCefObject(info *objectInfo) {
	m.infoTo(info)
	if len(info.SubObjectInfo) > 0 {
		for _, subInfo := range info.SubObjectInfo {
			m.subInfoToCefObject(subInfo)
		}
	}
}

func (m *objectTypeInfo) infoTo(info *objectInfo) {
	var (
		fieldLen = len(info.FieldInfos)
		fieldPtr uintptr
		funcLen  = len(info.FunctionInfos)
		funcPtr  uintptr
		cofs     []*valueBindInfo
		fns      []*valueBindInfo
		i        = 0
	)
	if fieldLen > 0 {
		//字段
		cofs = make([]*valueBindInfo, fieldLen, fieldLen)
		for fieldName, fi := range info.FieldInfos {
			cofs[i] = &valueBindInfo{
				Name:     api.PascalStr(fieldName),
				EventId:  fi.EventId,
				BindType: uintptr(fi.ValueType.Jsv),
			}
			i++
			if !info.isCreateJSValue {
				m.createObjectFieldVariable(info.FullObjName, fieldName, fi)
			}
		}
		fieldPtr = uintptr(unsafe.Pointer(&cofs[0]))
		i = 0

	}
	if funcLen > 0 {
		fns = make([]*valueBindInfo, funcLen, funcLen)
		for fnName, fn := range info.FunctionInfos {
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
			fmt.Println("fnName", fnName, info.FullObjName, info.ObjName)
			i++
			if !info.isCreateJSValue {
				m.createObjectFuncVariable(info.FullObjName, fnName, fn)
			}
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
func (m *objectTypeInfo) createObjectFieldVariable(fullParentName, fieldName string, sfi *fieldInfo) {
	newV8Value(sfi.EventId, fullParentName, fieldName, sfi.FieldValue, nil, sfi.ValueType, IS_OBJECT)
}

// 创建 结构对象的函数变量
func (m *objectTypeInfo) createObjectFuncVariable(fullParentName, funcName string, sfi *functionInfo) {
	newV8Value(sfi.EventId, fullParentName, funcName, nil, sfi, &VT{Jsv: V8_VALUE_FUNCTION, Gov: GO_VALUE_FUNC}, IS_OBJECT)
}

// 分析对象的字段
func (m *objectInfo) analysisObjectField(typ reflect.Type, typPtr reflect.Type, value reflect.Value) {
	if m.Parent == nil {
		m.ObjName = typ.Name()
		m.FullObjName = m.ObjName
	}
	m.Instance = value.Elem().UnsafeAddr()
	m.Ptr = unsafe.Pointer(m.Instance)
	//字段描述遍历
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
			subSoi := &objectInfo{FieldInfos: make(map[string]*fieldInfo), SubObjectInfo: make(map[string]*objectInfo)}
			subSoi.Parent = m
			subSoi.ParentInstance = m.Instance
			subSoi.ObjName = fieldName
			subSoi.FullObjName = fmt.Sprintf("%s.%s", subSoi.Parent.FullObjName, subSoi.ObjName)
			m.SubObjectInfo[fieldName] = subSoi
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
		if b { //true 可以解析绑定
			filedValue := value.Elem().FieldByName(fieldName)
			t := fieldType.Kind()
			gov, jsv := common.FieldReflectType(t)
			m.FieldInfos[fieldName] = &fieldInfo{
				EventId: uintptr(__bind_id()),
				ValueType: &VT{
					Jsv: jsv,
					Gov: gov,
				},
				FieldValue: &filedValue,
			}
		}
	}
	//解析当前对象函数
	m.analysisObjectFunction(typPtr, value)
}

// 分析对象的函数
// 不符合js类型的函数的参数，不会被解析成js函数
func (m *objectInfo) analysisObjectFunction(typPtr reflect.Type, value reflect.Value) {
	m.FunctionInfos = make(map[string]*functionInfo)
	for idx := 0; idx < typPtr.NumMethod(); idx++ {
		method := typPtr.Method(idx)
		if method.IsExported() {
			if fi, err := checkFunc(method.Type, FN_TYPE_OBJECT); err == nil {
				methodValue := value.MethodByName(method.Name)
				m.FunctionInfos[method.Name] = &functionInfo{
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

//bind
//直接绑定
func (m *objectTypeInfo) bind(value JSValue) {
	if value.IsFunction() {
		m.FunctionInfos[value.Name()] = &functionInfo{
			EventId:  value.getEventId(),
			funcInfo: value.getFuncInfo(),
			Method:   reflect.ValueOf(value.getValue()),
		}
	} else {
		m.FieldInfos[value.Name()] = &fieldInfo{
			EventId:    uintptr(__bind_id()),
			ValueType:  value.ValueType(),
			FieldValue: value.getValue(),
		}
	}
}

//bindTo
//直接绑定 -> 到CEF
func (m *objectTypeInfo) bindToCEF() {
	fmt.Println("ProcessType:", common.Args.ProcessType())
	tmpObjectInfo := &objectInfo{
		isCreateJSValue: true, //已创建完 JSValue
		FullObjName:     "",   //FullObjName & ObjName 值一样，绑定到根对象
		ObjName:         "",   //FullObjName & ObjName 值一样，绑定到根对象
		FieldInfos:      m.FieldInfos,
		FunctionInfos:   m.FunctionInfos,
	}
	m.infoTo(tmpObjectInfo)
}
