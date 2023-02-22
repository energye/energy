package src

import (
	"fmt"
	"github.com/energye/energy/cef"
	"time"
)

var (
	JSString *cef.JSString
	JSInt    *cef.JSInteger
	JSBool   *cef.JSBoolean
	JSDouble *cef.JSDouble
)

func JSFunc(p1 string) string {
	fmt.Println("参数:", p1)
	return p1 + " Go返回的值: " + time.Now().String()
}
