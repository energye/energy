//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build cgo

package gtk3

import (
	"github.com/energye/energy/v3/pkgs/linux/gtk3/cgo"
	. "github.com/energye/energy/v3/pkgs/linux/gtk3/types"
	"unsafe"
)

func AsScrolledWindow(ptr unsafe.Pointer) IScrolledWindow {
	return cgo.AsScrolledWindow(ptr)
}

func AsWindow(ptr unsafe.Pointer) IWindow {
	return cgo.AsWindow(ptr)
}
func AsContainer(ptr unsafe.Pointer) IContainer {
	return cgo.AsContainer(ptr)
}

func AsBox(ptr unsafe.Pointer) IBox {
	return cgo.AsBox(ptr)
}

func AsMenuBar(ptr unsafe.Pointer) IMenuBar {
	return cgo.AsMenuBar(ptr)
}

func AsWidget(ptr unsafe.Pointer) IWidget {
	return cgo.AsWidget(ptr)
}

func AsLayout(ptr unsafe.Pointer) ILayout {
	return cgo.AsLayout(ptr)
}

func AsDragContext(ptr unsafe.Pointer) IDragContext {
	return cgo.AsDragContext(ptr)
}

func AsAtom(ptr unsafe.Pointer) IAtom {
	return cgo.AsAtom(ptr)
}

func GdkAtomIntern(atomName string, onlyIfExists bool) IAtom {
	return cgo.GdkAtomIntern(atomName, onlyIfExists)
}

func NewCssProvider() ICssProvider {
	return cgo.NewCssProvider()
}
