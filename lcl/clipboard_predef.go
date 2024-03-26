//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/energy/v2/api"
	"github.com/energye/energy/v2/types"
)

func SetClipboard(newClipboard IObject) IClipboard {
	return AsClipboard(DSetClipboard(CheckPtr(newClipboard)))
}

func RegisterClipboardFormat(aFormat string) types.TClipboardFormat {
	return DRegisterClipboardFormat(aFormat)
}

func PredefinedClipboardFormat(aFormat types.TPredefinedClipboardFormat) types.TClipboardFormat {
	return DPredefinedClipboardFormat(aFormat)
}
