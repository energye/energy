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
	"github.com/energye/energy/v3/mime"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/wv/wv"
	"io/ioutil"
	"net/url"
	"path/filepath"
)

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
	tempArgs := wv.NewCoreWebView2WebResourceRequestedEventArgs(args)
	request := tempArgs.Request()
	tempRequest := wv.NewCoreWebView2WebResourceRequestRef(request)
	var (
		statusCode          int32 = 200
		reasonPhrase              = "OK"
		headers                   = ""
		data                []byte
		response            wv.ICoreWebView2WebResourceResponse
		assetsStream        lcl.IMemoryStream
		assetsStreamAdapter lcl.IStreamAdapter
	)
	reqUrl, err := url.Parse(tempRequest.URI())
	if err == nil {
		data, err = m.read(reqUrl.Path)
		if err == nil {
			assetsStream = lcl.NewMemoryStream()
			assetsStreamAdapter = lcl.NewStreamAdapter(assetsStream, types.SoOwned)
			assetsStream.Write(data)
			assetsStream.SetPosition(0)
			headers = "Content-Type: " + mime.GetMimeType(reqUrl.Path)
			environment := browser.CoreWebView2Environment()
			environment.CreateWebResourceResponse(assetsStreamAdapter, statusCode, reasonPhrase, headers, &response)
		}
	}
	if err != nil {
		statusCode = 404
		reasonPhrase = "Not Found"
		environment := browser.CoreWebView2Environment()
		environment.CreateWebResourceResponse(nil, statusCode, reasonPhrase, headers, &response)
	}
	tempArgs.SetResponse(response)

	tempRequest.FreeAndNil()
	tempArgs.FreeAndNil()
	response.Nil()
	if assetsStreamAdapter != nil {
		assetsStreamAdapter.Nil()
	}
	if assetsStream != nil {
		assetsStream.Nil()
	}
}
