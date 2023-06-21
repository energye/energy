//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// gtk3 Manually calling initialization

package gtk

import (
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
)

func CustomWidgetSetInitializationGtk3() {
	imports.Proc(def.GTK3_CustomWidgetSetInitialization).Call()
}

func CustomWidgetSetFinalizationGtk3() {
	imports.Proc(def.GTK3_CustomWidgetSetFinalization).Call()
}
