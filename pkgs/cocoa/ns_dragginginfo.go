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
	"github.com/energye/lcl/types"
	"github.com/energye/wv/darwin"
	"unsafe"
)

// NSDragOperation A group of constants that represent which operations the dragging source can perform on dragging items.
type NSDragOperation = int32

const (
	// NSDragOperationCopy A constant that indicates the drag can copy the data that the image represents.
	NSDragOperationCopy = NSDragOperation(C.NSDragOperationCopy)
	// NSDragOperationLink A constant that indicates the drag can share the data.
	NSDragOperationLink = NSDragOperation(C.NSDragOperationLink)
	// NSDragOperationGeneric A constant that indicates the destination can define the drag operation.
	NSDragOperationGeneric = NSDragOperation(C.NSDragOperationGeneric)
	// NSDragOperationPrivate A constant that indicates the source and destination negotiate the drag operation privately.
	NSDragOperationPrivate = NSDragOperation(C.NSDragOperationPrivate)
	// NSDragOperationMove A constant that indicates the drag can move the data.
	NSDragOperationMove = NSDragOperation(C.NSDragOperationMove)
	// NSDragOperationDelete A constant that indicates the drag can delete the data.
	NSDragOperationDelete = NSDragOperation(C.NSDragOperationDelete)
	// NSDragOperationEvery A constant that indicates that drag can perform all of the drag operations.
	NSDragOperationEvery = NSDragOperation(C.NSDragOperationEvery)
	// NSDragOperationNone A constant that indicates that the drag cannot perform any operations.
	NSDragOperationNone = NSDragOperation(C.NSDragOperationNone)
)

type TNSDraggingInfo struct {
	data unsafe.Pointer
}

func WrapNSDraggingInfo(data darwin.NSDraggingInfoProtocol) *TNSDraggingInfo {
	if data == 0 {
		return nil
	}
	return &TNSDraggingInfo{data: unsafe.Pointer(data)}
}

func (m *TNSDraggingInfo) DraggingPasteboard() *TNSPasteboard {
	if m.data == nil {
		return nil
	}
	cResult := C.DraggingPasteboard(m.data)
	return WrapNSPasteboard(uintptr(unsafe.Pointer(cResult)))
}

func (m *TNSDraggingInfo) DraggingLocation() (point types.TPoint) {
	if m.data == nil {
		return
	}
	dl := C.DraggingLocation(m.data)
	point.X = int32(dl.x)
	point.Y = int32(dl.y)
	return
}
