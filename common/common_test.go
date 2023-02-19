//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package common

import (
	"fmt"
	"testing"
	"time"
)

func TestGoTimeToDTime(t *testing.T) {
	for i := 0; i < 1; i++ {
		var now = time.Now()
		fmt.Println("now", now)
		var dNow = GoDateTimeToDDateTime(now)
		fmt.Println("dNow", dNow)
		var gNow = DDateTimeToGoDateTime(dNow)
		fmt.Println("gNow", gNow)
		fmt.Println("==========")
	}
	//1/24 + (1/24)/60*x
	//x=60162037037
	fmt.Println(63422453704 * 1000 / int64(time.Second))
	fmt.Println(62710648148 * 1000 / int64(time.Second))
	sub := time.Now().Sub(time.Date(2022, 10, 5, 0, 0, 0, 0, time.UTC))
	fmt.Println(sub.Seconds())
}
