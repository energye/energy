package assserv

import (
	"fmt"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"os"
	"path/filepath"
)

func StartServer() {
	wd, _ := os.Getwd()
	fmt.Println("wd", wd)
	server := assetserve.NewAssetsHttpServer()
	server.PORT = 22022
	server.LocalAssets = filepath.Join(wd, "examples", "tiny-browser", "cefclient", "assets")
	go server.StartHttpServer()
}
