package src

import (
	"fmt"
	"github.com/energye/energy/cef/bind"
	"time"
)

var (
	JSString *bind.JSString
	JSInt    *bind.JSInteger
	JSBool   *bind.JSBoolean
	JSDouble *bind.JSDouble
)

func JSFunc(p1 string) string {
	fmt.Println("参数:", p1)
	return p1 + " Go返回的值: " + time.Now().String()
}
