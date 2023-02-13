package cef

import (
	"github.com/energye/energy/common"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"strings"
)

type ObjectPropertyType int8

const (
	OptClass ObjectPropertyType = iota
	OptMethod
	OptSet
	OptEnum
	OptField
)

type ObjectProperty struct {
	Type          ObjectPropertyType
	PropertyName  string
	PropertyType  string
	PropertyValue string
}

func PanelBevelColor(panel *lcl.TPanel) types.TColor {
	r1, _, _ := common.Proc(internale_Ext_Panel_GetBevelColor).Call(panel.Instance())
	return types.TColor(r1)
}

func SetPanelBevelColor(panel *lcl.TPanel, colors types.TColor) {
	common.Proc(internale_Ext_Panel_SetBevelColor).Call(panel.Instance(), uintptr(colors))
}

func readObjectStringProperty(sender lcl.IObject) string {
	r1, _, _ := common.Proc(internale_Ext_ReadStringProperty).Call(lcl.CheckPtr(sender))
	return api.GoStr(r1)
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
