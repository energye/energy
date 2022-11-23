package src

import (
	"fmt"
	"github.com/energye/energy/cef"
)

var (
	JSString *cef.JSString
	JSInt    *cef.JSInteger
	JSBool   *cef.JSBoolean
	JSDouble *cef.JSDouble
)

func JSFunc(p1 string) string {
	fmt.Println("Go中执行JSFunc 参数:", p1)
	var ret string
	//类型判断
	if JSString.IsString() {
		ret = JSString.Value()
	} else if JSString.IsInteger() {
		//web js 中改变成integer类型
		intVal, _ := JSString.IntegerValue()
		ret = fmt.Sprintf("%d", intVal)
	}
	fmt.Println("JSString:", ret, "JSInt:", JSInt.Value(), "JSBool:", JSBool.Value(), "JSDouble:", JSDouble.Value())
	return p1 + " Go返回的值: " + ret
}
