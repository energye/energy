//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package wv

const energyProcessMessage = "processMessage"
const energyApplicationName = "energy"
const energyApplicationVersion = "3.0"

var (
	gRegisterSchemeCache = make(map[string]bool)
)

func setRegisterSchemeCache(scheme string) bool {
	if _, ok := gRegisterSchemeCache[scheme]; !ok {
		gRegisterSchemeCache[scheme] = true
		return false
	}
	return true
}

type TCustomSchemes struct {
	schemes []TCustomScheme
}

type TCustomScheme struct {
	Scheme string
	Domain string
}

func (m TCustomSchemes) Add(scheme, domain string) {
	m.schemes = append(m.schemes, TCustomScheme{scheme, domain})
}

type TApplicationOnCustomSchemesEvent func(customSchemes *TCustomSchemes)
