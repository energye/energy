//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	"fmt"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl/rtl"
)

func main() {
	inits.Init(nil, nil)
	fmt.Println("MainThreadId: ", rtl.MainThreadId())
	fmt.Println("CurrentThreadId: ", rtl.CurrentThreadId())
	go func() {
		fmt.Println("CurrentThreadId2: ", rtl.CurrentThreadId())
	}()
	var s string
	fmt.Scan(&s)
}
