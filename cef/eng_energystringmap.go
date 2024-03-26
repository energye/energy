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

// IEnergyStringMap Parent: ICefStringMap
//
//	CEF string maps are a set of key/value string pairs.
type IEnergyStringMap interface {
	ICefStringMap
}

// TEnergyStringMap Parent: TCefStringMap
//
//	CEF string maps are a set of key/value string pairs.
type TEnergyStringMap struct {
	TCefStringMap
}

func NewEnergyStringMap() IEnergyStringMap {
	r1 := CEF().SysCallN(2176)
	return AsEnergyStringMap(r1)
}

func EnergyStringMapClass() TClass {
	ret := CEF().SysCallN(2175)
	return TClass(ret)
}
