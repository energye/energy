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
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/api/dllimports"
	"github.com/energye/golcl/lcl/types"
	"strings"
)

type ObjectPropertyType int8

const (
	ObjPropTypeClass ObjectPropertyType = iota
	ObjPropTypeMethod
	ObjPropTypeSet
	ObjPropTypeEnum
	ObjPropTypeField
)

type ObjectProperty struct {
	Type          ObjectPropertyType
	PropertyName  string
	PropertyType  string
	PropertyValue string
}

func init() {
	var lclExtImportDefs = []*dllimports.ImportTable{
		dllimports.NewEnergyImport("Ext_Panel_GetBevelColor", 0),
		dllimports.NewEnergyImport("Ext_Panel_SetBevelColor", 0),
		dllimports.NewEnergyImport("Ext_ReadStringProperty", 0),
		dllimports.NewEnergyImport("Ext_Form_Activate", 0),
		dllimports.NewEnergyImport("Ext_Form_Deactivate", 0),
		dllimports.NewEnergyImport("Ext_Assigned", 0),
	}
	imports.SetLClExtImportDefs(lclExtImportDefs)
}

const (
	Ext_Panel_GetBevelColor = iota
	Ext_Panel_SetBevelColor
	Ext_ReadStringProperty
	Ext_Form_Activate
	Ext_Form_Deactivate
	Ext_Assigned
)

func PanelBevelColor(panel *lcl.TPanel) types.TColor {
	r1, _, _ := imports.ExtProc(Ext_Panel_GetBevelColor).Call(panel.Instance())
	return types.TColor(r1)
}

func SetPanelBevelColor(panel *lcl.TPanel, colors types.TColor) {
	imports.ExtProc(Ext_Panel_SetBevelColor).Call(panel.Instance(), uintptr(colors))
}

func readObjectStringProperty(sender lcl.IObject) string {
	r1, _, _ := imports.ExtProc(Ext_ReadStringProperty).Call(lcl.CheckPtr(sender))
	return api.GoStr(r1)
}

func FormActivate(form *lcl.TForm) {
	imports.ExtProc(Ext_Form_Activate).Call(form.Instance())
}

func FormDeactivate(form *lcl.TForm) {
	imports.ExtProc(Ext_Form_Deactivate).Call(form.Instance())
}

func ReadObjectStringProperty(sender lcl.IObject) map[string]ObjectProperty {
	str := readObjectStringProperty(sender)
	strs := strings.Split(str, "\n")
	var result = map[string]ObjectProperty{}
	for _, prop := range strs {
		p := strings.Split(prop, "-")
		if len(p) == 4 {
			op := ObjectProperty{}
			op.Type = ObjectPropertyType(common.StrToInt32(p[0]))
			op.PropertyName = p[1]
			op.PropertyType = p[2]
			op.PropertyValue = p[3]
			result[op.PropertyName] = op
		}
	}
	return result
}

func Assigned(ptr uintptr) bool {
	r1, _, _ := imports.ExtProc(Ext_Assigned).Call(ptr)
	return api.GoBool(r1)
}
