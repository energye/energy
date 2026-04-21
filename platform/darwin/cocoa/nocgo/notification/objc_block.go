//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package notification

import (
	"errors"
	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
	"unsafe"
)

// blockDescriptor is the Go representation of an Objective-C block descriptor.
// It is a component to be referenced by blockDescriptor.
//
// The layout of this struct matches Block_literal_1 described in https://clang.llvm.org/docs/Block-ABI-Apple.html#high-level
type blockDescriptor struct {
	_         uintptr
	size      uintptr
	_         uintptr
	dispose   uintptr
	signature *uint8
}

// blockLayout is the Go representation of the structure abstracted by a block pointer.
// From the Objective-C point of view, a pointer to this struct is equivalent to an ID that
// references a block.
//
// The layout of this struct matches __block_literal_1 described in https://clang.llvm.org/docs/Block-ABI-Apple.html#high-level
type blockLayout struct {
	isa        objc.Class
	flags      uint32
	_          uint32
	invoke     uintptr
	descriptor *blockDescriptor
}

func BlockInvoke(block objc.Block, args ...uintptr) (uintptr, error) {
	layout := (*blockLayout)(unsafe.Pointer(block))
	if layout == nil || layout.invoke == 0 {
		return 0, errors.New("invoke pointer is nil")
	}
	args = append([]uintptr{uintptr(block)}, args...)
	r1, _, _ := purego.SyscallN(layout.invoke, args...)
	return r1, nil
}
