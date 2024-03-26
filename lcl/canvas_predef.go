//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/energy/v2/types"
)

func (m *TCanvas) TextRect2(aRect *TRect, text string, textFormat TTextFormat) {
	style := m.TextStyle()
	if textFormat.In(TfSingleLine) {
		style.SingleLine = true
	}
	if textFormat.In(TfTop) {
		style.Layout = TlTop
	} else if textFormat.In(TfVerticalCenter) {
		style.Layout = TlCenter
	} else if textFormat.In(TfBottom) {
		style.Layout = TlBottom
	}
	if textFormat.In(TfNoClip) {
		style.Clipping = true
	}
	if textFormat.In(TfExpandTabs) {
		style.ExpandTabs = true
	}
	if !textFormat.In(TfHidePrefix) && !textFormat.In(TfNoPrefix) {
		style.ShowPrefix = true
	}
	if textFormat.In(TfWordBreak) {
		style.Wordbreak = true
	}
	if textFormat.In(TfEndEllipsis) {
		style.EndEllipsis = true
	}
	if textFormat.In(TfLeft) {
		style.Alignment = TaLeftJustify
	} else if textFormat.In(TfCenter) {
		style.Alignment = TaCenter
	} else if textFormat.In(TfRight) {
		style.Alignment = TaRightJustify
	}
	m.TextRect1(aRect, aRect.Left, aRect.Top, text, &style)
}
