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
