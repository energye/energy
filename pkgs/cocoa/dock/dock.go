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

package dock

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa
#import <Cocoa/Cocoa.h>

static void Hide() {
  [[NSApplication sharedApplication] setActivationPolicy:NSApplicationActivationPolicyProhibited];
}

static void Show() {
  [[NSApplication sharedApplication] setActivationPolicy:NSApplicationActivationPolicyRegular];
}
*/
import "C"

func Hide() {
	C.Hide()
}

func Show() {
	C.Show()
}
