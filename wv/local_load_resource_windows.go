//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

package wv

import (
	"fmt"
	"github.com/energye/energy/v3/mime"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/wv/wv"
	"io/ioutil"
	"net/url"
	"path/filepath"
)

var (
	assetsStream        lcl.IMemoryStream
	assetsStreamAdapter lcl.IStreamAdapter
)

func localLoadStreamCreate() {
	assetsStream = lcl.NewMemoryStream()
	assetsStreamAdapter = lcl.NewStreamAdapter(assetsStream, types.SoOwned)
}

func (m *LocalLoadResource) read(path string) ([]byte, error) {
	if localLoadRes.FS == nil {
		var rootPath string
		if localLoadRes.ResRootDir[0] == '@' {
			rootPath = filepath.Join(localLoadRes.exePath, localLoadRes.ResRootDir[1:])
		} else {
			rootPath = localLoadRes.ResRootDir
		}
		return ioutil.ReadFile(filepath.Join(rootPath, path))
	} else {
		return m.FS.ReadFile(localLoadRes.ResRootDir + path)
	}
}

func (m *LocalLoadResource) resourceRequested(browser wv.IWVBrowser, webView wv.ICoreWebView2, args wv.ICoreWebView2WebResourceRequestedEventArgs) {
	// 自定义协议资源加载
	args = wv.NewCoreWebView2WebResourceRequestedEventArgs(args)
	request := wv.NewCoreWebView2WebResourceRequestRef(args.Request())
	// 需要释放掉
	defer func() {
		request.Free()
		args.Free()
	}()
	// 重置 stream
	assetsStream.SetPosition(0)
	assetsStream.Clear()
	fmt.Println("回调函数 WVBrowser => SetOnWebResourceRequested")
	fmt.Println("回调函数 WVBrowser => TempURI:", request.URI(), request.Method())
	var (
		statusCode   int32 = 200
		reasonPhrase       = "OK"
		headers            = ""
	)
	reqUrl, err := url.Parse(request.URI())
	if err != nil {
		fmt.Println("加载本地资源-error:", err)
		statusCode = 404
		reasonPhrase = "not found"
		headers = "Content-Type: application/json"
	} else {
		data, err := m.read(reqUrl.Path)
		if err != nil {
			fmt.Println("加载本地资源-error:", err)
			statusCode = 404
			reasonPhrase = "not found"
			headers = "Content-Type: application/json"
		} else {
			headers = "Content-Type: " + mime.GetMimeType(reqUrl.Path)
			assetsStream.LoadFromBytes(data)
		}
	}
	fmt.Println("回调函数 WVBrowser => stream", assetsStream.Size())
	fmt.Println("回调函数 WVBrowser => adapter:", assetsStreamAdapter.StreamOwnership(), assetsStreamAdapter.Stream().Size())

	var response wv.ICoreWebView2WebResourceResponse
	environment := browser.CoreWebView2Environment()
	fmt.Println("回调函数 WVBrowser => Initialized():", environment.Initialized(), environment.BrowserVersionInfo())
	environment.CreateWebResourceResponse(assetsStreamAdapter, statusCode, reasonPhrase, headers, &response)
	args.SetResponse(response)
}
