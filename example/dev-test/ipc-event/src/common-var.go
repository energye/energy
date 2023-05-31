package src

import (
	"fmt"
	"time"
)

func JSFunc(p1 string) string {
	fmt.Println("参数:", p1)
	return p1 + " Go返回的值: " + time.Now().String()
}
