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

func (m *LocalLoadResource) read(path string) ([]byte, error) {
	if m.FS == nil {
		var rootPath string
		if m.ResRootDir[0] == '@' {
			rootPath = filepath.Join(m.exePath, m.ResRootDir[1:])
		} else {
			rootPath = m.ResRootDir
		}
		return ioutil.ReadFile(filepath.Join(rootPath, path))
	} else {
		return m.FS.ReadFile(m.ResRootDir + path)
	}
}

// released after the resource processing is complete
func (m *LocalLoadResource) releaseStream(path string) {
	if stream, ok := m.streams[path]; ok {
		stream.FreeAndNil()
		delete(m.streams, path)
	}
}

// current resource is set temp cache
func (m *LocalLoadResource) setTempStream(path string, stream lcl.IMemoryStream) {
	m.streams[path] = stream
}

func (m *LocalLoadResource) resourceRequested(browser wv.IWVBrowser, webView wv.ICoreWebView2, args wv.ICoreWebView2WebResourceRequestedEventArgs) {
	// temp object
	tempArgs := wv.NewCoreWebView2WebResourceRequestedEventArgs(args)
	defer tempArgs.FreeAndNil()
	request := tempArgs.Request()
	tempRequest := wv.NewCoreWebView2WebResourceRequestRef(request)
	defer tempRequest.FreeAndNil()
	var (
		statusCode    int32 = 200
		reasonPhrase        = "OK"
		headers             = ""
		data          []byte
		response      wv.ICoreWebView2WebResourceResponse
		stream        lcl.IMemoryStream
		streamAdapter lcl.IStreamAdapter
	)
	//m.freePreviousResourceStream()
	// request url, get file path
	reqUrl, err := url.Parse(tempRequest.URI())
	if err == nil {
		// read resource
		data, err = m.read(reqUrl.Path)
		fmt.Println("data-len:", len(data), err, reqUrl.Path)
		if err == nil {
			stream = lcl.NewMemoryStream()
			streamAdapter = lcl.NewStreamAdapter(stream, types.SoOwned)
			defer streamAdapter.Nil()
			stream.Write(data)
			// current resource is set temp cache
			// released after the resource processing is complete
			m.setTempStream(reqUrl.Path, stream)
			headers = "Content-Type: " + mime.GetMimeType(reqUrl.Path)
			environment := browser.CoreWebView2Environment()
			// success response resource
			environment.CreateWebResourceResponse(streamAdapter, statusCode, reasonPhrase, headers, &response)
		}

	}
	// No matter what error, return 404
	if err != nil {
		statusCode = 404
		reasonPhrase = "Not Found"
		environment := browser.CoreWebView2Environment()
		// empty response resource
		environment.CreateWebResourceResponse(nil, statusCode, reasonPhrase, headers, &response)
	}
	tempArgs.SetResponse(response)

	if response != nil {
		response.Nil()
	}
}
