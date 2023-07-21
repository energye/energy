//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package internal

import "os"

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	return v.(string)
}

func ToRNilString(v interface{}, new string) string {
	if v == nil {
		return new
	}
	return v.(string)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		} else if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}
