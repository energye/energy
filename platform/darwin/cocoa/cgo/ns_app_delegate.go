package cgo

/*
#cgo CFLAGS: -mmacosx-version-min=10.15
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa

void InitAppDelegate(void);
*/
import "C"
import (
	"encoding/json"
	"fmt"
)

//export GoOpenURLsCallback
func GoOpenURLsCallback(cURLs *C.char) {
	var (
		urls   []string
		goUrls = C.GoString(cURLs)
	)
	_ = json.Unmarshal([]byte(goUrls), &urls)
	fmt.Println("打开文件:", urls)
}

func InitAppDelegate() {
	C.InitAppDelegate()
}
