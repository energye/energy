package assserv

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"os"
)

var Assets embed.FS

func StartServer() {
	wd, _ := os.Getwd()
	fmt.Println("wd", wd)
	server := assetserve.NewAssetsHttpServer()
	server.PORT = 22022
	server.Assets = Assets
	server.AssetsFSName = "assets"
	go server.StartHttpServer()
}
