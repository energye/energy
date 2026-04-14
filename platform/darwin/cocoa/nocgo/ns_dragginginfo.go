//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	"github.com/ebitengine/purego/objc"
	. "github.com/energye/energy/v3/platform/darwin/types"
	"github.com/energye/lcl/types"
	"unsafe"
)

type NSDraggingInfo struct {
	instance unsafe.Pointer
}

func WrapNSDraggingInfo(data unsafe.Pointer) INSDraggingInfo {
	if data == nil {
		return nil
	}
	return &NSDraggingInfo{instance: data}
}

func (m *NSDraggingInfo) DraggingPasteboard() INSPasteboard {
	if m.instance == nil {
		return nil
	}
	draggingInfo := objc.ID(m.instance)
	pasteboard := draggingInfo.Send(objc.RegisterName("draggingPasteboard"))
	if pasteboard == 0 {
		return nil
	}
	return WrapNSPasteboard(unsafe.Pointer(pasteboard))
}

func (m *NSDraggingInfo) DraggingLocation() (point types.TPoint) {
	if m.instance == nil {
		return
	}
	draggingInfo := objc.ID(m.instance)
	location := objc.Send[CGPoint](draggingInfo, objc.RegisterName("draggingLocation"))
	point.X = int32(location.X)
	point.Y = int32(location.Y)
	return
}
