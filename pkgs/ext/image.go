//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package ext

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
)

type TImage struct {
	*lcl.TImage
}

func NewImage(owner lcl.IComponent) *TImage {
	m := new(TImage)
	m.TImage = lcl.NewImage(owner)
	return m
}

func (m *TImage) SetOnPaint(fn lcl.TNotifyEvent) {
	if m.IsValid() {
		api.Image_SetOnPaint(m.Instance(), fn)
	}
}

func (m *TImage) SetOnPictureChanged(fn lcl.TNotifyEvent) {
	if m.IsValid() {
		api.Image_SetOnPictureChanged(m.Instance(), fn)
	}
}
