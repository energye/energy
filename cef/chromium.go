//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/lcl/lcl"
)

type TChromium struct {
	ICEFWindowParent
	ICEFChromium
}

func NewChromium(owner lcl.IWinControl) *TChromium {
	m := &TChromium{}
	m.ICEFChromium = NewCEFChromium(owner)
	m.ICEFWindowParent = NewCEFWindowParent(m.ICEFChromium, owner)

	const (
		HpDisableNonProxiedUDP = 3
		STATE_DISABLED         = 2
	)
	m.SetWebRTCIPHandlingPolicy(HpDisableNonProxiedUDP)
	m.SetWebRTCMultipleRoutes(STATE_DISABLED)
	m.SetWebRTCNonproxiedUDP(STATE_DISABLED)

	return m
}
