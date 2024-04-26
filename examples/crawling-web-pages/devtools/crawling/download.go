package crawling

import (
	"fmt"
	"github.com/go-rod/rod/lib/proto"
	"os"
	"path/filepath"
)

// 以下代码来自 rod download

func Download(windowId int) {
	if window, ok := windows[windowId]; ok {
		page := window.energy.Page().MustWaitLoad()
		fmt.Println("TargetID:", page.TargetID)
		wd, _ := os.Getwd()
		wait := window.energy.WaitDownload(wd)
		go window.energy.EachEvent(func(e *proto.PageDownloadProgress) bool {
			completed := "(unknown)"
			if e.TotalBytes != 0 {
				completed = fmt.Sprintf("%0.2f%%", e.ReceivedBytes/e.TotalBytes*100.0)
			}
			fmt.Printf("state: %s, completed: %s\n", e.State, completed)
			return e.State == proto.PageDownloadProgressStateCompleted
		})()
		page.MustElement(`#win64`).MustClick()
		res := wait()
		fmt.Printf("wrote %s\n", filepath.Join(wd, res.GUID))
	}
}
