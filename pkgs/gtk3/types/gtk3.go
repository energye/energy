//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package types

import "unsafe"

type IObject interface {
	Instance() uintptr
	Ref()
	Unref()
}

type IScreen interface {
	IObject
	GetRGBAVisual() IVisual
	IsComposited() bool
}

type IVisual interface {
	IObject
}

type IWidget interface {
	IObject
	GetScreen() IScreen
	SetVisual(visual IVisual)
	SetAppPaintable(paintable bool)
	GetName() string
	GetAllocation() IRectangle
	SetSizeRequest(width, height int)
	GetSizeRequest() (width, height int)
	GetStyleContext() IStyleContext
}

type IContainer interface {
	IWidget
	Add(w IWidget)
	Remove(w IWidget)
	CheckResize()
	GetChildren() IList
}

type IBin interface {
	IContainer
}

type IBox interface {
	IContainer
	PackStart(child IWidget, expand, fill bool, padding uint)
	PackEnd(child IWidget, expand, fill bool, padding uint)
}

type IStyleProvider interface {
	Instance() uintptr
}

type IList interface {
	Instance() uintptr
	Append(data uintptr) IList
	Prepend(data uintptr) IList
	Insert(data uintptr, position int) IList
	Length() uint
	NthDataRaw(n uint) unsafe.Pointer
	Next() IList
	Previous() IList
	First() IList
	Last() IList
	Free()
}

type IStyleContext interface {
	IObject
	AddClass(class_name string)
	RemoveClass(class_name string)
	AddProvider(provider IStyleProvider, prio uint)
}

type IWindow interface {
	IBin
	GetDefaultSize() (width, height int)
	SetDecorated(setting bool)
	Maximize()
	Unmaximize()
	Fullscreen()
	Unfullscreen()
	SetTitle(title string)
	GetTitle() string
}

type IMenuShell interface {
	IContainer
}

type IScrolledWindow interface {
	IBin
}

type IMenuBar interface {
	IMenuShell
}

type ILayout interface {
	IContainer
	Put(w IWidget, x, y int)
	Move(w IWidget, x, y int)
	SetSize(width, height uint)
	GetSize() (width, height uint)
}

type ICssProvider interface {
	IObject
	LoadFromPath(path string) error
	LoadFromData(data string) error
	ToString() string
}

type IRectangle interface {
	GetX() int
	SetX(x int)
	GetY() int
	SetY(y int)
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}
