//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestArgsSet(t *testing.T) {
	var longStrBuf string
	for i := 0; i < 10; i++ {
		longStrBuf += "GoCEF字符串" + fmt.Sprintf("[%d]", i)
	}
	fmt.Println("longStrBuf：", longStrBuf)
	args := &argumentList{}
	args.SetInt(0, 2)
	args.SetInt(1, 2)
	fmt.Println(args.Package(), args.GetInt(0), args.GetInt(1))
	var data = args.Package()[:]
	fmt.Println(data)
	args1 := &argumentList{}
	args1.UnPackage(data)
	fmt.Println(args1.Package(), args1.GetInt(0), args1.GetInt(1))
	fmt.Println("------")
	var getVal = func(args IArgumentList) {
		fmt.Println("0=", args.GetString(0))
		fmt.Println("1=", args.GetInt(1))
		fmt.Println("2=", args.GetInt8(2))
		fmt.Println("3=", args.GetInt16(3))
		fmt.Println("4=", args.GetInt32(4))
		fmt.Println("5=", args.GetInt64(5))
		fmt.Println("6=", args.GetBool(6))
		fmt.Println("7=", args.GetBool(7))
		fmt.Println("8=", args.GetFloat32(8))
		fmt.Println("9=", args.GetFloat64(9))
	}
	args = &argumentList{}
	args.SetString(0, longStrBuf)
	args.SetInt(1, 100009)
	args.SetInt8(2, 127)
	args.SetInt16(3, 16330)
	args.SetInt32(4, 90000001)
	args.SetInt64(5, 99999999999)
	args.SetBool(6, true)
	args.SetBool(7, false)
	args.SetFloat32(8, 9999.9999)
	args.SetFloat64(9, 999999.9999)
	getVal(args)
	data = args.Package()
	fmt.Println(data)
	args1 = &argumentList{}
	args1.UnPackage(data)
	getVal(args1)
	data = args1.Package()
	fmt.Println(data)
}

func TestSize(t *testing.T) {
	fmt.Println("不同int类型占用的字节数大小：")
	var i1 int = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5
	var i6 uint = 6
	var i7 uint8 = 7
	var i8 uint16 = 8
	var i9 uint32 = 9
	var i10 uint64 = 10
	var i11 float32 = 11.0
	var i12 float64 = 12.0
	fmt.Printf("int    : %v\n", unsafe.Sizeof(i1))
	fmt.Printf("int8   : %v\n", unsafe.Sizeof(i2))
	fmt.Printf("int16  : %v\n", unsafe.Sizeof(i3))
	fmt.Printf("int32  : %v\n", unsafe.Sizeof(i4))
	fmt.Printf("int64  : %v\n", unsafe.Sizeof(i5))
	fmt.Printf("uint  : %v\n", unsafe.Sizeof(i6))
	fmt.Printf("uint8  : %v\n", unsafe.Sizeof(i7))
	fmt.Printf("uint16  : %v\n", unsafe.Sizeof(i8))
	fmt.Printf("uint32  : %v\n", unsafe.Sizeof(i9))
	fmt.Printf("uint64  : %v\n", unsafe.Sizeof(i10))
	fmt.Printf("float32  : %v\n", unsafe.Sizeof(i11))
	fmt.Printf("float64  : %v\n", unsafe.Sizeof(i12))
}
