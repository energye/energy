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

package cgo

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa

#import <Cocoa/Cocoa.h>
#import <WebKit/WebKit.h>

void AppDockHide() {
  [[NSApplication sharedApplication] setActivationPolicy:NSApplicationActivationPolicyProhibited];
}

void AppDockShow() {
  [[NSApplication sharedApplication] setActivationPolicy:NSApplicationActivationPolicyRegular];
}

void AppSetPresentationOptions(NSApplicationPresentationOptions opts) {
	[[NSApplication sharedApplication] setPresentationOptions:opts];
}

void AppSetMainMenu(void* nsMenu) {
	NSMenu* menu = (NSMenu*)nsMenu;
    if (!menu) {
        NSLog(@"AppSetMainMenu menu nil");
        return ;
    }
    NSApplication *app = [NSApplication sharedApplication];
    [app setMainMenu:menu];
}

*/
import "C"
import (
	. "github.com/energye/energy/v3/pkgs/darwin/types"
	"unsafe"
)

func AppSetPresentationOptions(options NSApplicationPresentationOptions) {
	C.AppSetPresentationOptions(C.NSApplicationPresentationOptions(C.uint(options)))
}

func AppDockHide() {
	C.AppDockHide()
}

func AppDockShow() {
	C.AppDockShow()
}

func AppSetMainMenu(nsMenu unsafe.Pointer) {
	C.AppSetMainMenu(nsMenu)
}
