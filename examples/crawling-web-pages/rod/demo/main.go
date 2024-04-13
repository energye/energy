// Package main ...
package main

import (
	"flag"
	rod "github.com/energye/energy/v2/examples/crawling-web-pages/rod"
	"log"
)

var flagDevToolWsURL = flag.String("devtools-ws-url", "", "DevTools WebSocket URL")

// This example demonstrates how to connect to an existing Chrome DevTools
// instance using a remote WebSocket URL.
func main() {
	flag.Parse()
	// chrome.exe --headless --remote-debugging-port=8777
	*flagDevToolWsURL = "ws://127.0.0.1:9222/devtools/browser/41c8af06-f898-4e32-bb5c-a65506066a3f"
	if *flagDevToolWsURL == "" {
		log.Fatal("must specify -devtools-ws-url")
	}
	browse := rod.New()
	browse = browse.ControlURL(*flagDevToolWsURL)
	browse = browse.MustConnect()
	page := browse.MustPage("https://www.baidu.com")
	page.MustElement("#wrapper").MustWaitVisible()
	log.Println("Body of duckduckgo.com starts with:")
	log.Println(page.MustHTML()[0:100])
}
