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
#cgo darwin CFLAGS: -DDARWIN -x objective-c
#cgo darwin LDFLAGS: -framework WebKit -framework Cocoa

#import <WebKit/WebKit.h>
#import <Cocoa/Cocoa.h>

NSPasteboard* DraggingPasteboard(void* nsDraggingInfo) {
	id<NSDraggingInfo> dragInfo = (id<NSDraggingInfo>)nsDraggingInfo;
	if (!dragInfo) {
        NSLog(@"DraggingPasteboard dragInfo is nil");
        return nil;
    }
    return [dragInfo draggingPasteboard];
}

NSPoint DraggingLocation(void* nsDraggingInfo) {
	id<NSDraggingInfo> dragInfo = (id<NSDraggingInfo>)nsDraggingInfo;
	if (!dragInfo) {
        NSLog(@"DraggingPasteboard dragInfo is nil");
        return NSMakePoint(0, 0);
    }
    return [dragInfo draggingLocation];
}

*/
import "C"

import (
	. "github.com/energye/energy/v3/platform/darwin/types"
	"github.com/energye/lcl/types"
	"unsafe"
)

type NSDraggingInfo struct {
	data unsafe.Pointer
}

func WrapNSDraggingInfo(data unsafe.Pointer) INSDraggingInfo {
	if data == nil {
		return nil
	}
	return &NSDraggingInfo{data: data}
}

func (m *NSDraggingInfo) DraggingPasteboard() INSPasteboard {
	if m.data == nil {
		return nil
	}
	cResult := C.DraggingPasteboard(m.data)
	return WrapNSPasteboard(unsafe.Pointer(cResult))
}

func (m *NSDraggingInfo) DraggingLocation() (point types.TPoint) {
	if m.data == nil {
		return
	}
	dl := C.DraggingLocation(m.data)
	point.X = int32(dl.x)
	point.Y = int32(dl.y)
	return
}
