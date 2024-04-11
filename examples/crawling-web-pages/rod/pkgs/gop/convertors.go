package gop

import (
	"encoding/base64"
	"reflect"
	"time"
)

// SymbolPtr for Ptr
const SymbolPtr = "gop.Ptr"

// Ptr returns a pointer to v
func Ptr(v interface{}) interface{} {
	val := reflect.ValueOf(v)
	ptr := reflect.New(val.Type())
	ptr.Elem().Set(val)
	return ptr.Interface()
}

// SymbolCircular for Circular
const SymbolCircular = "gop.Circular"

// Circular reference of the path from the root
func Circular(path ...interface{}) interface{} {
	return nil
}

// SymbolBase64 for Base64
const SymbolBase64 = "gop.Base64"

// Base64 returns the []byte that s represents
func Base64(s string) []byte {
	b, _ := base64.StdEncoding.DecodeString(s)
	return b
}

// SymbolTime for Time
const SymbolTime = "gop.Time"

// Time from parsing s
func Time(s string, monotonic int) time.Time {
	t, _ := time.Parse(time.RFC3339Nano, s)
	return t
}

// SymbolDuration for Duration
const SymbolDuration = "gop.Duration"

// Duration from parsing s
func Duration(s string) time.Duration {
	d, _ := time.ParseDuration(s)
	return d
}

// SymbolJSONStr for JSONStr
const SymbolJSONStr = "gop.JSONStr"

// JSONStr returns the raw
func JSONStr(v interface{}, raw string) string {
	return raw
}

// SymbolJSONBytes for JSONBytes
const SymbolJSONBytes = "gop.JSONBytes"

// JSONBytes returns the raw as []byte
func JSONBytes(v interface{}, raw string) []byte {
	return []byte(raw)
}
