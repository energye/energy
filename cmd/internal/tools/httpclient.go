//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package tools

import (
	"github.com/energye/rawhttp"
	"io/ioutil"
	"time"
)

func Get(url string, proxy string) ([]byte, error) {
	options := &rawhttp.Options{
		Timeout:                30 * time.Second,
		FollowRedirects:        true,
		MaxRedirects:           10,
		AutomaticHostHeader:    true,
		AutomaticContentLength: true,
		Proxy:                  proxy,
	}
	client := rawhttp.NewClient(options)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
