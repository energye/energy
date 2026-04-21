//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package notification

import (
	"encoding/json"
	"github.com/ebitengine/purego/objc"
	"unsafe"
)

const NSUTF8StringEncoding = 4

func convertGoMapToNSDictionary(data map[string]any) objc.ID {
	if len(data) == 0 {
		return 0
	}
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return 0
	}
	jsonStr := GoStringToNSString(string(jsonBytes))
	if jsonStr == 0 {
		return 0
	}
	jsonData := jsonStr.Send(objc.RegisterName("dataUsingEncoding:"), NSUTF8StringEncoding)
	if jsonData == 0 {
		return 0
	}
	var errorPtr uintptr
	dict := objc.ID(objc.GetClass("NSJSONSerialization")).Send(
		objc.RegisterName("JSONObjectWithData:options:error:"),
		jsonData, uintptr(0), uintptr(unsafe.Pointer(&errorPtr)),
	)
	if errorPtr != 0 || dict == 0 {
		return 0
	}
	return dict
}

func convertNSDictionaryToGoMap(dict objc.ID) map[string]any {
	if dict == 0 {
		return nil
	}
	const NSJSONWritingPrettyPrinted = 0
	var errorPtr uintptr
	jsonData := objc.ID(objc.GetClass("NSJSONSerialization")).Send(
		objc.RegisterName("dataWithJSONObject:options:error:"),
		dict,
		uintptr(NSJSONWritingPrettyPrinted),
		uintptr(unsafe.Pointer(&errorPtr)),
	)
	if errorPtr != 0 || jsonData == 0 {
		return nil
	}
	jsonBytes := NSDataToGoBytes(jsonData)
	if len(jsonBytes) == 0 {
		return nil
	}
	var result map[string]any
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return nil
	}
	return result
}

func NSDataToGoBytes(nsData objc.ID) []byte {
	length := nsData.Send(objc.RegisterName("length"))
	bytes := nsData.Send(objc.RegisterName("bytes"))
	if bytes != 0 && length != 0 {
		return unsafe.Slice((*byte)(unsafe.Pointer(bytes)), length)
	}
	return nil
}

func GoStringToNSString(str string) objc.ID {
	return objc.ID(objc.GetClass("NSString")).Send(objc.RegisterName("stringWithUTF8String:"), str)
}

func NSStringToGoString(nsString objc.ID) string {
	typeStr := nsString.Send(objc.RegisterName("UTF8String"))
	utf8Length := nsString.Send(objc.RegisterName("lengthOfBytesUsingEncoding:"), 4)
	value := unsafe.String((*byte)(unsafe.Pointer(typeStr)), uintptr(utf8Length))
	return value
}
