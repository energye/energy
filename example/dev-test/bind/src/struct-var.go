package src

import "fmt"

var JSStructVarDemo *StructVarDemo

//定义结构类型的变量
//结构类型将属性和函数导出为JS可调用类型
//大写字母开头
type StructVarDemo struct {
	StringField   string
	IntField      int32
	BoolField     bool
	FloatField    float64
	noExportField string //小写字母无法绑定到js
}

//结构类型的函数导出
func (m *StructVarDemo) StringValue(p0 string) string {
	fmt.Println("结构类型绑定函数 StringValue 被调用 入参:", p0)
	return m.StringField + p0
}

//结构类型的函数导出
func (m *StructVarDemo) IntValue(intParam int32) int32 {
	fmt.Println("结构类型绑定函数 IntValue 被调用 入参:", intParam)
	return m.IntField + intParam
}

//定义导出函数参数最多为9个
func (m *StructVarDemo) FuncMaxParam(p1, p2, p3, p4, p5, p6, p7, p8, p9 string) {
	fmt.Println("结构类型绑定函数 FuncMaxParam 被调用 入参:", p1, p2, p3, p4, p5, p6, p7, p8, p9)
}
