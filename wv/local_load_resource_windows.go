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
	// temp object
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
	// request url, get file path
	reqUrl, err := url.Parse(tempRequest.URI())
	if err == nil {
		// read resource
		data, err = m.read(reqUrl.Path)
		if err == nil {
			assetsStream = lcl.NewMemoryStream()
			assetsStreamAdapter = lcl.NewStreamAdapter(assetsStream, types.SoOwned)
			assetsStream.Write(data)
			headers = "Content-Type: " + mime.GetMimeType(reqUrl.Path)
			environment := browser.CoreWebView2Environment()
			// success response resource
			environment.CreateWebResourceResponse(assetsStreamAdapter, statusCode, reasonPhrase, headers, &response)
			environment.Nil()
		}
	}
	// No matter what error, return 404
	if err != nil {
		statusCode = 404
		reasonPhrase = "Not Found"
		environment := browser.CoreWebView2Environment()
		// empty response resource
		environment.CreateWebResourceResponse(nil, statusCode, reasonPhrase, headers, &response)
		environment.Nil()
	}
	tempArgs.SetResponse(response)

	tempRequest.FreeAndNil()
	tempArgs.FreeAndNil()
	response.Nil()
	request.Nil()
	if assetsStreamAdapter != nil {
		assetsStreamAdapter.Nil()
	}
	if assetsStream != nil {
		assetsStream.Nil()
	}
}
