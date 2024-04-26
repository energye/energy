package crawling

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

// 示例来自 rod upload

// Upload 上传个试试？
func Upload(windowId int) {
	if window, ok := windows[windowId]; ok {
		page := window.energy.Page().MustWaitLoad()
		fmt.Println("TargetID:", page.TargetID)
		wd, _ := os.Getwd()
		wd = filepath.Join(wd, "devtools.go")
		fmt.Println("upload file:", wd)
		page.MustElement(`input[name="upload"]`).MustSetFiles(wd)
		page.MustElement(`input[name="submit"]`).MustClick()
	}
}

func UploadServer() string {
	// create http server and result channel
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprint(res, uploadHTML)
	})
	mux.HandleFunc("/upload", func(res http.ResponseWriter, req *http.Request) {
		f, _, err := req.FormFile("upload")
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		defer func() { _ = f.Close() }()

		buf, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		_, _ = fmt.Fprintf(res, resultHTML, string(buf))
	})
	l, _ := net.Listen("tcp4", "127.0.0.1:0")
	go func() { _ = http.Serve(l, mux) }()
	return "http://" + l.Addr().String()
}

const (
	uploadHTML = `<!doctype html>
<html>
<body>
  <form method="POST" action="/upload" enctype="multipart/form-data">
    <input name="upload" type="file"/>
    <input name="submit" type="submit"/>
  </form>
</body>
</html>`

	resultHTML = `<!doctype html>
<html>
<body>
  <div id="result">%v</div>
</body>
</html>`
)
