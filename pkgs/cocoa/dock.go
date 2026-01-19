//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin

package cocoa

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa
#import <Cocoa/Cocoa.h>

static void DockHide() {
  [[NSApplication sharedApplication] setActivationPolicy:NSApplicationActivationPolicyProhibited];
}

static void DockShow() {
  [[NSApplication sharedApplication] setActivationPolicy:NSApplicationActivationPolicyRegular];
}
*/
import "C"

func DockHide() {
	C.DockHide()
}

func DockShow() {
	C.DockShow()
}
