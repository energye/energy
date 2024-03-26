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

// IEnergyStringList Parent: ICefCustomStringList
//
//	CEF string maps are a set of key/value string pairs.
type IEnergyStringList interface {
	ICefCustomStringList
}

// TEnergyStringList Parent: TCefCustomStringList
//
//	CEF string maps are a set of key/value string pairs.
type TEnergyStringList struct {
	TCefCustomStringList
}

func NewEnergyStringList() IEnergyStringList {
	r1 := CEF().SysCallN(2174)
	return AsEnergyStringList(r1)
}

func EnergyStringListClass() TClass {
	ret := CEF().SysCallN(2173)
	return TClass(ret)
}
