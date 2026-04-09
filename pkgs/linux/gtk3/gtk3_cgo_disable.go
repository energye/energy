//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !cgo

package gtk3

import (
	"github.com/energye/energy/v3/pkgs/linux/gtk3/nocgo"
	. "github.com/energye/energy/v3/pkgs/linux/gtk3/types"
	"unsafe"
)

func AsScrolledWindow(ptr unsafe.Pointer) IScrolledWindow {
	return nocgo.AsScrolledWindow(ptr)
}

func AsWindow(ptr unsafe.Pointer) IWindow {
	return nocgo.AsWindow(ptr)
}
func AsContainer(ptr unsafe.Pointer) IContainer {
	return nocgo.AsContainer(ptr)
}

func AsBox(ptr unsafe.Pointer) IBox {
	return nocgo.AsBox(ptr)
}

func AsMenuBar(ptr unsafe.Pointer) IMenuBar {
	return nocgo.AsMenuBar(ptr)
}

func AsLayout(ptr unsafe.Pointer) ILayout {
	return nocgo.AsLayout(ptr)
}

func AsWidget(ptr unsafe.Pointer) IWidget {
	return nocgo.AsWidget(ptr)
}
func AsDragContext(ptr unsafe.Pointer) IDragContext {
	return nocgo.AsDragContext(ptr)
}

func AsAtom(ptr unsafe.Pointer) IAtom {
	return nocgo.AsAtom(ptr)
}

func AsEventButton(ptr unsafe.Pointer) IEventButton {
	return nocgo.AsEventButton(ptr)
}

func AsEventCrossing(ptr unsafe.Pointer) IEventCrossing {
	return nocgo.AsEventCrossing(ptr)
}

func AsEntry(ptr unsafe.Pointer) IEntry {
	return nocgo.AsEntry(ptr)
}

func NewEntry() IEntry {
	return nocgo.NewEntry()
}

func GdkAtomIntern(atomName string, onlyIfExists bool) IAtom {
	return nocgo.GdkAtomIntern(atomName, onlyIfExists)
}

func NewCssProvider() ICssProvider {
	return nocgo.NewCssProvider()
}
