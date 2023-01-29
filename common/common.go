//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package common

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/decimal"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/api/dllimports"
	"math"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const (
	IntSize     = strconv.IntSize             //bit
	IntSize32   = 32                          //
	intSize64   = 64                          //
	isWindows   = runtime.GOOS == "windows"   //support
	isLinux     = runtime.GOOS == "linux"     //support
	isDarwin    = runtime.GOOS == "darwin"    //support
	isAndroid   = runtime.GOOS == "android"   //not support
	isIos       = runtime.GOOS == "ios"       //not support
	isPlan9     = runtime.GOOS == "plan9"     //not support
	isAix       = runtime.GOOS == "aix"       //not support
	isDragonfly = runtime.GOOS == "dragonfly" //not support
	isFreebsd   = runtime.GOOS == "freebsd"   //not support
	isHurd      = runtime.GOOS == "hurd"      //not support
	isIllumos   = runtime.GOOS == "illumos"   //not support
	isJs        = runtime.GOOS == "js"        //not support
	isNacl      = runtime.GOOS == "nacl"      //not support
	isNetbsd    = runtime.GOOS == "netbsd"    //not support
	isOpenbsd   = runtime.GOOS == "openbsd"   //not support
	isSolaris   = runtime.GOOS == "solaris"   //not support
	isZos       = runtime.GOOS == "zos"       //not support
)

func Proc(index int) dllimports.ProcAddr {
	return api.EnergyDefSyscallN(index)
}

func IsWindows() bool {
	return isWindows
}

func IsLinux() bool {
	return isLinux
}

func IsDarwin() bool {
	return isDarwin
}

func IsPlan9() bool {
	return isPlan9
}

func StrToInt64(value string) int64 {
	v, _ := strconv.ParseInt(value, 10, 64)
	return v
}
func StrToInt32(value string) int32 {
	v, _ := strconv.ParseInt(value, 10, 32)
	return int32(v)
}

func StrToFloat64(value string) float64 {
	v, _ := strconv.ParseFloat(value, 64)
	return v
}

func StrToFloat32(value string) float32 {
	v, _ := strconv.ParseFloat(value, 32)
	return float32(v)
}

func InterfaceToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

// 获取参数指针
func GetParamOf(index int, ptr uintptr) uintptr {
	return *(*uintptr)(unsafe.Pointer(ptr + uintptr(index)*unsafe.Sizeof(ptr)))
}

// 根据指定指针位置开始 偏移获取指针
func GetParamPtr(ptr uintptr, offset int) unsafe.Pointer {
	return unsafe.Pointer(ptr + uintptr(offset))
}

func GOValueReflectType(v interface{}) GO_VALUE_TYPE {
	if v == nil {
		return GO_VALUE_NIL
	}
	vType := reflect.TypeOf(v).Kind()
	switch vType {
	case reflect.String:
		return GO_VALUE_STRING
	case reflect.Int:
		return GO_VALUE_INT
	case reflect.Int8:
		return GO_VALUE_INT8
	case reflect.Int16:
		return GO_VALUE_INT16
	case reflect.Int32:
		return GO_VALUE_INT32
	case reflect.Int64:
		return GO_VALUE_INT64
	case reflect.Uint:
		return GO_VALUE_UINT
	case reflect.Uint8:
		return GO_VALUE_UINT8
	case reflect.Uint16:
		return GO_VALUE_UINT16
	case reflect.Uint32:
		return GO_VALUE_UINT32
	case reflect.Uint64:
		return GO_VALUE_UINT64
	case reflect.Uintptr:
		return GO_VALUE_UINTPTR
	case reflect.Float32:
		return GO_VALUE_FLOAT32
	case reflect.Float64:
		return GO_VALUE_FLOAT64
	case reflect.Bool:
		return GO_VALUE_BOOL
	case reflect.Struct:
		return GO_VALUE_STRUCT
	case reflect.Slice:
		return GO_VALUE_SLICE
	case reflect.Func:
		return GO_VALUE_FUNC
	case reflect.Ptr:
		return GO_VALUE_PTR
	default:
		return GO_VALUE_EXCEPTION
	}
}

func GOValueType(v string) GO_VALUE_TYPE {
	if v == "nil" {
		return GO_VALUE_NIL
	}
	switch v {
	case "string":
		return GO_VALUE_STRING
	case "int":
		return GO_VALUE_INT
	case "int8":
		return GO_VALUE_INT8
	case "int16":
		return GO_VALUE_INT16
	case "int32":
		return GO_VALUE_INT32
	case "int64":
		return GO_VALUE_INT64
	case "uint":
		return GO_VALUE_UINT
	case "uint8":
		return GO_VALUE_UINT8
	case "uint16":
		return GO_VALUE_UINT16
	case "uint32":
		return GO_VALUE_UINT32
	case "uint64":
		return GO_VALUE_UINT64
	case "uintptr":
		return GO_VALUE_UINTPTR
	case "float32":
		return GO_VALUE_FLOAT32
	case "float64":
		return GO_VALUE_FLOAT64
	case "bool":
		return GO_VALUE_BOOL
	case "struct":
		return GO_VALUE_STRUCT
	case "slice":
		return GO_VALUE_SLICE
	case "func":
		return GO_VALUE_FUNC
	case "ptr":
		return GO_VALUE_PTR
	default:
		return GO_VALUE_EXCEPTION
	}
}

func GOValueAssertType(v interface{}) GO_VALUE_TYPE {
	if v == nil {
		return GO_VALUE_NIL
	}
	switch v.(type) {
	case string:
		return GO_VALUE_STRING
	case int:
		return GO_VALUE_INT
	case int8:
		return GO_VALUE_INT8
	case int16:
		return GO_VALUE_INT16
	case int32:
		return GO_VALUE_INT32
	case int64:
		return GO_VALUE_INT64
	case uint:
		return GO_VALUE_UINT
	case uint8:
		return GO_VALUE_UINT8
	case uint16:
		return GO_VALUE_UINT16
	case uint32:
		return GO_VALUE_UINT32
	case uint64:
		return GO_VALUE_UINT64
	case uintptr:
		return GO_VALUE_UINTPTR
	case float32:
		return GO_VALUE_FLOAT32
	case float64:
		return GO_VALUE_FLOAT64
	case bool:
		return GO_VALUE_BOOL
	default:
		return GO_VALUE_EXCEPTION
	}
}

func JSValueAssertType(v interface{}) V8_JS_VALUE_TYPE {
	switch v.(type) {
	case string:
		return V8_VALUE_STRING
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return V8_VALUE_INT
	case float32, float64:
		return V8_VALUE_DOUBLE
	case bool:
		return V8_VALUE_BOOLEAN
	default:
		return V8_VALUE_EXCEPTION
	}
}

func JSValueType(v string) V8_JS_VALUE_TYPE {
	if v == "nil" {
		return V8_VALUE_NULL
	}
	switch v {
	case "string":
		return V8_VALUE_STRING
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr":
		return V8_VALUE_INT
	case "float32", "float64":
		return V8_VALUE_DOUBLE
	case "bool":
		return V8_VALUE_BOOLEAN
	case "struct":
		return V8_VALUE_OBJECT
	case "slice":
		return V8_VALUE_ARRAY
	case "func":
		return V8_VALUE_FUNCTION
	default:
		return V8_VALUE_EXCEPTION
	}
}
func ValueToBool(v interface{}) (bool, error) {
	switch v.(type) {
	case []byte:
		return ByteToBool(v.([]byte)[0]), nil
	case byte:
		return ByteToBool(v.(byte)), nil
	}
	vType := JSValueAssertType(v)
	switch vType {
	case V8_VALUE_INT:
		v := StrToInt32(InterfaceToString(v))
		if v == 0 {
			return false, nil
		} else {
			return true, nil
		}
	case V8_VALUE_DOUBLE:
		v := StrToFloat64(InterfaceToString(v))
		if v == 0 {
			return false, nil
		} else {
			return true, nil
		}
	case V8_VALUE_STRING:
		v := v.(string)
		if v == "" {
			return false, nil
		} else {
			return true, nil
		}
	case V8_VALUE_BOOLEAN:
		return v.(bool), nil
	default:
		return false, errors.New("转换bool类型失败")
	}
}

func ValueToFloat32(v interface{}) (float32, error) {
	vType := JSValueAssertType(v)
	switch vType {
	case V8_VALUE_INT:
		return StrToFloat32(InterfaceToString(v)), nil
	case V8_VALUE_DOUBLE:
		return StrToFloat32(InterfaceToString(v)), nil
	case V8_VALUE_STRING:
		return StrToFloat32(v.(string)), nil
	case V8_VALUE_BOOLEAN:
		v := v.(bool)
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, errors.New("转换float32类型失败")
	}
}

func ValueToFloat64(v interface{}) (float64, error) {
	switch v.(type) {
	case []byte:
		return BytesToFloat64(v.([]byte)), nil
	}
	vType := JSValueAssertType(v)
	switch vType {
	case V8_VALUE_INT:
		return StrToFloat64(InterfaceToString(v)), nil
	case V8_VALUE_DOUBLE:
		return StrToFloat64(InterfaceToString(v)), nil
	case V8_VALUE_STRING:
		return StrToFloat64(v.(string)), nil
	case V8_VALUE_BOOLEAN:
		v := v.(bool)
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, errors.New("转换float64类型失败")
	}
}

func ValueToInt32(v interface{}) (int32, error) {
	switch v.(type) {
	case []byte:
		return BytesToInt32(v.([]byte)), nil
	}
	vType := JSValueAssertType(v)
	switch vType {
	case V8_VALUE_INT:
		return StrToInt32(InterfaceToString(v)), nil
	case V8_VALUE_DOUBLE:
		return int32(math.Round(StrToFloat64(InterfaceToString(v)))), nil
	case V8_VALUE_STRING:
		return StrToInt32(v.(string)), nil
	case V8_VALUE_BOOLEAN:
		v := v.(bool)
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, errors.New("转换int32类型失败")
	}
}

func ValueToInt64(v interface{}) (int64, error) {
	vType := JSValueAssertType(v)
	switch vType {
	case V8_VALUE_INT:
		return StrToInt64(InterfaceToString(v)), nil
	case V8_VALUE_DOUBLE:
		return int64(math.Round(StrToFloat64(InterfaceToString(v)))), nil
	case V8_VALUE_STRING:
		return StrToInt64(v.(string)), nil
	case V8_VALUE_BOOLEAN:
		v := v.(bool)
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, errors.New("转换int64类型失败")
	}
}

func NumberUintPtrToInt(value uintptr, gov GO_VALUE_TYPE) interface{} {
	switch gov {
	case GO_VALUE_INT:
		return int(value)
	case GO_VALUE_INT8:
		return int8(value)
	case GO_VALUE_INT16:
		return int16(value)
	case GO_VALUE_INT32:
		return int32(value)
	case GO_VALUE_INT64:
		return int64(value)
	case GO_VALUE_UINT:
		return uint(value)
	case GO_VALUE_UINT8:
		return uint8(value)
	case GO_VALUE_UINT16:
		return uint16(value)
	case GO_VALUE_UINT32:
		return uint32(value)
	case GO_VALUE_UINT64:
		return uint64(value)
	case GO_VALUE_UINTPTR:
		return value
	default:
		return nil
	}
}

func NumberPtrToFloat(value unsafe.Pointer, gov GO_VALUE_TYPE) interface{} {
	switch gov {
	case GO_VALUE_FLOAT32:
		return *(*float64)(value)
	case GO_VALUE_FLOAT64:
		return *(*float64)(value)
	default:
		return nil
	}
}

func ValueToString(v interface{}) (string, error) {
	switch v.(type) {
	case []byte:
		return BytesToString(v.([]byte)), nil
	}
	vType := JSValueAssertType(v)
	switch vType {
	case V8_VALUE_INT, V8_VALUE_DOUBLE:
		return fmt.Sprintf("%v", vType), nil
	case V8_VALUE_STRING, V8_VALUE_NULL, V8_VALUE_UNDEFINED:
		return v.(string), nil
	case V8_VALUE_BOOLEAN:
		v := v.(bool)
		if v {
			return "true", nil
		} else {
			return "false", nil
		}
	default:
		return "", errors.New("转换string类型失败")
	}
}

func ValueToBytes(v interface{}) []byte {
	switch v.(type) {
	case []byte:
		return v.([]byte)
	case byte:
		return []byte{v.(byte)}

	}
	return nil
}

func ParamType(t string) (V8_JS_VALUE_TYPE, GO_VALUE_TYPE) {
	switch t {
	case "string":
		return V8_VALUE_STRING, GO_VALUE_STRING
	case "int":
		return V8_VALUE_INT, GO_VALUE_INT
	case "int8":
		return V8_VALUE_INT, GO_VALUE_INT8
	case "int16":
		return V8_VALUE_INT, GO_VALUE_INT16
	case "int32":
		return V8_VALUE_INT, GO_VALUE_INT32
	case "int64":
		return V8_VALUE_INT, GO_VALUE_INT64
	case "float32":
		return V8_VALUE_DOUBLE, GO_VALUE_FLOAT32
	case "float64":
		return V8_VALUE_DOUBLE, GO_VALUE_FLOAT64
	case "bool":
		return V8_VALUE_BOOLEAN, GO_VALUE_BOOL
	case "EefError":
		return V8_VALUE_EXCEPTION, GO_VALUE_EXCEPTION
	default:
		return -1, -1
	}
}

func FuncParamJsTypeStr(jsValue V8_JS_VALUE_TYPE) string {
	switch jsValue {
	case V8_VALUE_STRING:
		return "string"
	case V8_VALUE_INT:
		return "int"
	case V8_VALUE_DOUBLE:
		return "double"
	case V8_VALUE_BOOLEAN:
		return "boolean"
	case V8_VALUE_EXCEPTION:
		return "EefError"
	default:
		return ""
	}
}

func FuncParamGoTypeStr(jsValue GO_VALUE_TYPE) string {
	switch jsValue {
	case GO_VALUE_STRING:
		return "string"
	case GO_VALUE_INT:
		return "int"
	case GO_VALUE_INT8:
		return "int8"
	case GO_VALUE_INT16:
		return "int16"
	case GO_VALUE_INT32:
		return "int32"
	case GO_VALUE_INT64:
		return "int64"
	case GO_VALUE_FLOAT32:
		return "float32"
	case GO_VALUE_FLOAT64:
		return "float64"
	case GO_VALUE_BOOL:
		return "bool"
	case GO_VALUE_EXCEPTION:
		return "EefError"
	default:
		return ""
	}
}

func CopyBytePtr(bytePtr uintptr, low, high int) []byte {
	var size = high - low
	var data = make([]byte, size, size)
	for i := low; i < high; i++ {
		data[i-low] = *(*byte)(unsafe.Pointer(bytePtr + (uintptr(i))))
	}
	return data
}

func IntToBytes(i int) []byte {
	buf := bytes.NewBuffer([]byte{})
	if IntSize == IntSize32 {
		if err := binary.Write(buf, binary.BigEndian, int32(i)); err == nil {
			return buf.Bytes()
		}
	} else {
		if err := binary.Write(buf, binary.BigEndian, int64(i)); err == nil {
			return buf.Bytes()
		}
	}
	return nil
}

func Int8ToBytes(i int8) []byte {
	return []byte{byte(i)}
}

func Int16ToBytes(i int16) []byte {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, i); err == nil {
		return buf.Bytes()
	}
	return nil
}

func Int32ToBytes(i int32) []byte {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, i); err == nil {
		return buf.Bytes()
	}
	return nil
}

func Int64ToBytes(i int64) []byte {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, i); err == nil {
		return buf.Bytes()
	}
	return nil
}

func BytesToInt(b []byte) int {
	var i int64
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		return 0
	}
	return int(i)
}

func ByteToInt8(b byte) int8 {
	return int8(b)
}

func BytesToInt16(b []byte) int16 {
	var i int16
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		return 0
	}
	return i
}

func BytesToInt32(b []byte) int32 {
	var i int32
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		return 0
	}
	return i
}

func BytesToInt64(b []byte) int64 {
	var i int64
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		return 0
	}
	return i
}

func BytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

//func StringToBytes(data string) []byte {
//	return *(*[]byte)(unsafe.Pointer(&data))
//}

//String转换Bytes数组，isDStr转换DString 默认GoString
func StringToBytes(s string, isDStr ...bool) []byte {
	if len(isDStr) > 0 && isDStr[0] {
		temp := []byte(s)
		utf8StrArr := make([]byte, len(temp)+1)
		copy(utf8StrArr, temp)
		return utf8StrArr
	} else {
		return []byte(s)
	}
}

//Float64ToBytes Float64转byte
func Float64ToBytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

//BytesToFloat64 byte转Float64
func BytesToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

//Float32ToBytes Float64转byte
func Float32ToBytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

//BytesToFloat32 byte转Float64
func BytesToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func BoolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func ByteToBool(b byte) bool {
	if b == 1 {
		return true
	}
	return false
}

const (
	dSecond float64 = 3600
	dDay            = 24 * dSecond
)

var dBaseDateTime = time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)

func GoDateTimeToDDateTime(dateTime time.Time) float64 {
	date := float64(dateTime.Sub(dBaseDateTime).Milliseconds() / 1000 / 60 / 60 / 24)
	diHour := decimal.NewFromFloat(float64(dateTime.Hour()))
	diMinute := decimal.NewFromFloat(float64(dateTime.Minute())).Mul(decimal.NewFromFloat(60))
	diSecond := decimal.NewFromFloat(float64(dateTime.Second()))
	diTime := diHour.Mul(decimal.NewFromFloat(dSecond)).Add(diMinute).Add(diSecond).Div(decimal.NewFromFloat(dDay))
	var dTime, _ = diTime.Add(decimal.NewFromFloat(date)).Float64()
	return dTime
}

func DDateTimeToGoDateTime(dateTime float64) time.Time {
	dtStr := strings.Split(fmt.Sprintf("%v", dateTime), ".")
	dDate, _ := strconv.Atoi(dtStr[0])
	diDateTime := decimal.NewFromFloat(dateTime)
	diDate := decimal.NewFromFloat(float64(dDate))
	diTime := diDateTime.Sub(diDate)
	dTime, _ := diTime.Float64()
	gTime := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
	gTime = gTime.AddDate(0, 0, dDate)
	diTime = decimal.NewFromFloat(float64(time.Second)).Mul(decimal.NewFromFloat(dTime).Mul(decimal.NewFromFloat(dDay)))
	diTime = diTime.Add(decimal.NewFromFloat(dTime))
	gTime = gTime.Add(time.Duration(diTime.IntPart()))
	return gTime
}

func ArrayIndexOf[T any](array []T, a interface{}) int {
	if len(array) == 0 {
		return -1
	}
	var t any
	for i := 0; i < len(array); i++ {
		t = array[i]
		if t == a {
			return i
		}
	}
	return -1
}

//获取指针的指针的地址
func GetInstancePtr(ptr uintptr) unsafe.Pointer {
	ptr = *(*uintptr)(unsafe.Pointer(ptr))
	return unsafe.Pointer(ptr)
}

//GoroutineID 获取当前携程ID
func GoroutineID() (id uint64) {
	var buf [30]byte
	runtime.Stack(buf[:], false)
	for i := 10; buf[i] != ' '; i++ {
		id = id*10 + uint64(buf[i]&15)
	}
	return id
}
