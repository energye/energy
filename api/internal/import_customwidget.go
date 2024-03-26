//----------------------------------------
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// 自定义组件初始化

package internal

import "github.com/energye/energy/v2/api/imports"

// CustomWidgetInterface for Linux
var customWidgetImportDefs = []*imports.Table{
	imports.NewTable("Interface_CustomWidgetSetInitialization", 0),
	imports.NewTable("Interface_CustomWidgetSetFinalization", 0),
}

// InitCustomWidgetImport 初始化自定义组件初始化api
func InitCustomWidgetImport(imp *imports.Imports) {
	imp.SetOk(true)
	imp.SetImportTable(customWidgetImportDefs)
}
