//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package common CEF Util
package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/cyber-xxm/energy/v2/pkgs/decimal"
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

// Concat 字符串拼接
func Concat(str ...string) string {
	var c = strings.Builder{}
	for _, v := range str {
		c.WriteString(v)
	}
	return c.String()
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

// InterfaceToString 接口转 string
func InterfaceToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

// GetParamOf 获取参数指针
func GetParamOf(index int, ptr uintptr) uintptr {
	return *(*uintptr)(unsafe.Pointer(ptr + uintptr(index)*unsafe.Sizeof(ptr)))
}

// GetParamPtr 根据指定指针位置开始 偏移获取指针
func GetParamPtr(ptr uintptr, offset int) unsafe.Pointer {
	return unsafe.Pointer(ptr + uintptr(offset))
}

// ValueToBool bool
func ValueToBool(v interface{}) bool {
	switch v.(type) {
	case []byte:
		bv := v.([]byte)
		if len(bv) == 1 {
			return ByteToInt8(bv[0]) > 0
		} else if len(bv) == 2 {
			return BytesToInt16(bv) > 0
		} else if len(bv) == 4 {
			return BytesToInt32(bv) > 0
		} else if len(bv) == 8 {
			return BytesToInt64(bv) > 0
		}
		return len(bv) > 0
	}
	switch v.(type) {
	case string:
		return len(v.(string)) > 0
	case float32:
		return v.(float32) > 0
	case float64:
		return v.(float64) > 0
	case bool:
		return v.(bool)
	case int:
		return v.(int) > 0
	case int8:
		return v.(int8) > 0
	case int16:
		return v.(int16) > 0
	case int32:
		return v.(int32) > 0
	case int64:
		return v.(int64) > 0
	case uintptr:
		return v.(uintptr) > 0
	default:
		return false
	}
}

func ValueToFloat64(v interface{}) float64 {
	switch v.(type) {
	case []byte:
		bv := v.([]byte)
		if len(bv) == 4 {
			return float64(BytesToFloat32(bv))
		} else if len(bv) == 8 {
			return BytesToFloat64(bv)
		}
		return 0.0
	}
	switch v.(type) {
	case string:
		return StrToFloat64(v.(string))
	case float32:
		return float64(v.(float32))
	case float64:
		return v.(float64)
	case bool:
		if v.(bool) {
			return 1
		} else {
			return 0
		}
	case int:
		return float64(v.(int))
	case int8:
		return float64(v.(int8))
	case int16:
		return float64(v.(int16))
	case int32:
		return float64(v.(int32))
	case int64:
		return float64(v.(int64))
	case uintptr:
		return float64(v.(uintptr))
	default:
		return 0
	}
}

func ValueToInt(v interface{}) int {
	switch v.(type) {
	case []byte:
		bv := v.([]byte)
		if len(bv) == 1 {
			return int(ByteToInt8(bv[0]))
		} else if len(bv) == 2 {
			return int(BytesToInt16(bv))
		} else if len(bv) == 4 {
			return int(BytesToInt32(bv))
		} else if len(bv) == 8 {
			return int(BytesToInt64(bv))
		}
		return 0
	}
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

func UIntToBytes(i uint) []byte {
	buf := bytes.NewBuffer([]byte{})
	if IntSize == IntSize32 {
		if err := binary.Write(buf, binary.BigEndian, uint32(i)); err == nil {
			return buf.Bytes()
		}
	} else {
		if err := binary.Write(buf, binary.BigEndian, uint64(i)); err == nil {
			return buf.Bytes()
		}
	}
	return nil
}

func Int8ToBytes(i int8) []byte {
	return []byte{byte(i)}
}

func UInt8ToBytes(i uint8) []byte {
	return []byte{byte(i)}
}

func Int16ToBytes(i int16) []byte {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, i); err == nil {
		return buf.Bytes()
	}
	return nil
}

func UInt16ToBytes(i uint16) []byte {
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

func UInt32ToBytes(i uint32) []byte {
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

func UInt64ToBytes(i uint64) []byte {
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

func BytesToUInt(b []byte) uint {
	var i uint64
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		return 0
	}
	return uint(i)
}

func ByteToInt8(b byte) int8 {
	return int8(b)
}

func ByteToUInt8(b byte) uint8 {
	return uint8(b)
}

func BytesToInt16(b []byte) int16 {
	var i int16
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	if err != nil {
		return 0
	}
	return i
}

func BytesToUInt16(b []byte) uint16 {
	var i uint16
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

func BytesToUInt32(b []byte) uint32 {
	var i uint32
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

func BytesToUInt64(b []byte) uint64 {
	var i uint64
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

// String转换Bytes数组，isDStr转换DString 默认GoString
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

// Float64ToBytes Float64转byte
func Float64ToBytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

// BytesToFloat64 byte转Float64
func BytesToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

// Float32ToBytes Float64转byte
func Float32ToBytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

// BytesToFloat32 byte转Float64
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
	date := float64(dateTime.Sub(dBaseDateTime).Nanoseconds() / 1000000 / 60 / 60 / 24)
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

// ArrayIndexOf 返回a在array数组的下标, a必须和array元素的类型相同
func ArrayIndexOf(array interface{}, a interface{}) int {
	switch array.(type) {
	case []string:
		arrs := array.([]string)
		if len(arrs) == 0 {
			return -1
		}
		eq := a.(string)
		for i, v := range arrs {
			if v == eq {
				return i
			}
		}
		//case []int: //其它待增加
	}
	return -1
}

// GetInstancePtr 获取指针的指针的地址
func GetInstancePtr(ptr uintptr) unsafe.Pointer {
	ptr = *(*uintptr)(unsafe.Pointer(ptr))
	return unsafe.Pointer(ptr)
}

// GoroutineID 获取当前携程ID
func GoroutineID() (id uint64) {
	var buf [30]byte
	runtime.Stack(buf[:], false)
	for i := 10; buf[i] != ' '; i++ {
		id = id*10 + uint64(buf[i]&15)
	}
	return id
}

func GoStr(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}
	resultString := (*reflect.StringHeader)(unsafe.Pointer(ptr))
	if resultString == nil || resultString.Len <= 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(resultString))

}

func string2bytes1(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	var b []byte
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pbytes.Data = stringHeader.Data
	pbytes.Len = stringHeader.Len
	pbytes.Cap = stringHeader.Len
	return b
}
