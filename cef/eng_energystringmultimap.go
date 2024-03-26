//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// IEnergyStringMultiMap Parent: ICefStringMultimap
//
//	CEF string multimaps are a set of key/value string pairs.
//	More than one value can be assigned to a single key.
type IEnergyStringMultiMap interface {
	ICefStringMultimap
}

// TEnergyStringMultiMap Parent: TCefStringMultimap
//
//	CEF string multimaps are a set of key/value string pairs.
//	More than one value can be assigned to a single key.
type TEnergyStringMultiMap struct {
	TCefStringMultimap
}

func NewEnergyStringMultiMap() IEnergyStringMultiMap {
	r1 := CEF().SysCallN(2178)
	return AsEnergyStringMultiMap(r1)
}

func EnergyStringMultiMapClass() TClass {
	ret := CEF().SysCallN(2177)
	return TClass(ret)
}
