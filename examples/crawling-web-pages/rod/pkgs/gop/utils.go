package gop

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// GetPrivateField via field index
// TODO: we can use a LRU cache for the copy of the values, but it might be trivial for just testing.
func GetPrivateField(v reflect.Value, i int) reflect.Value {
	if v.Kind() != reflect.Struct {
		panic("expect v to be a struct")
	}

	copied := reflect.New(v.Type()).Elem()
	copied.Set(v)
	f := copied.Field(i)
	return reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
}

// GetPrivateFieldByName is similar with GetPrivateField
func GetPrivateFieldByName(v reflect.Value, name string) reflect.Value {
	if v.Kind() != reflect.Struct {
		panic("expect v to be a struct")
	}

	copied := reflect.New(v.Type()).Elem()
	copied.Set(v)
	f := copied.FieldByName(name)
	return reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
}

// compare returns the float value of x minus y
func compare(x, y interface{}) int {
	return strings.Compare(fmt.Sprintf("%#v", x), fmt.Sprintf("%#v", y))
}

func formatUintptr(p uintptr) string {
	return "0x" + strconv.FormatUint(uint64(p), 16)
}

func wrapComment(s string) string {
	return "/* " + s + " */"
}

func formatLenCap(l, c int) string {
	if c >= 0 {
		return fmt.Sprintf("/* len=%d cap=%d */", l, c)
	}
	return fmt.Sprintf("/* len=%d */", l)
}
