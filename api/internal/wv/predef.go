//----------------------------------------
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import "github.com/energye/energy/v2/api/imports"

var wvImportDefs = []*imports.Table{
	// TWVLoader Event
	/* 0 */ imports.NewTable("WVLoader_SetOnEnvironmentCreated", 0),
	/* 1 */ imports.NewTable("WVLoader_SetOnInitializationError", 0),
	/* 2 */ imports.NewTable("WVLoader_SetOnGetCustomSchemes", 0),
	/* 3 */ imports.NewTable("WVLoader_SetOnNewBrowserVersionAvailable", 0),
	/* 4 */ imports.NewTable("WVLoader_SetOnBrowserProcessExited", 0),
	/* 5 */ imports.NewTable("WVLoader_SetOnProcessInfosChanged", 0),
	/* 6 */ imports.NewTable("WVLoader_SetOnProcessInfosChanged", 0),
	// TWVLoader Proc
	/* 7 */ imports.NewTable("WVLoader_GlobalWebView2Loader", 0),
}

// InitWVPreDefsImport 初始化CEF预定义api
func InitWVPreDefsImport(imp *imports.Imports) {
	imp.SetImportTable(wvImportDefs)
	imp.SetOk(true)
}
