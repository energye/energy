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
	. "github.com/energye/energy/v3/pkgs/linux/types"
)
import "unsafe"

// Atom is a representation of GDK's GdkAtom.
type Atom TAtom

func AsAtom(v unsafe.Pointer) IAtom {
	return Atom(v)
}

func (m Atom) Name() string {
	r := gdk3.SysCall("gdk_atom_name", uintptr(m))
	return GoStr(r)
}

func (m Atom) Atom() TAtom {
	return TAtom(m)
}

// GdkAtomIntern is a wrapper around gdk_atom_intern
func GdkAtomIntern(atomName string, onlyIfExists bool) IAtom {
	r := gdk3.SysCall("gdk_atom_intern", CStr(atomName), ToCBool(onlyIfExists))
	return Atom(r)
}
