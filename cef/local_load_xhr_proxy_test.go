package cef

import (
	"bytes"
	"log"
	"net/http"
	"testing"
)

func TestXHRProxyClientSSLSend(t *testing.T) {
	proxy := &XHRProxy{
		IP: "energy.yanghy.cn",
	}
	proxy.init()
	httpRequest, err := http.NewRequest("GET", "https://energy.yanghy.cn/api/energy/download", nil)
	if err != nil {
		log.Println("new request:", err.Error())
		return
	}
	httpResponse, err := proxy.HttpClient.Client.Do(httpRequest)
	if err != nil {
		log.Println("do:", err.Error())
		return
	}
	defer httpResponse.Body.Close()
	buf := new(bytes.Buffer)
	c, err := buf.ReadFrom(httpResponse.Body)
	if err != nil {
		log.Println("read from:", err.Error())
		return
	}
	log.Println(c, err, httpResponse.StatusCode)
	log.Println(buf.String())
}
