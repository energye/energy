//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build finalizerOn
// +build finalizerOn

package api

import "runtime"

func callGC() {
	runtime.GC()
}
