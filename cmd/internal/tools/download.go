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
	"bufio"
	"errors"
	"fmt"
	"github.com/energye/energy/v2/cmd/internal/term"
	"github.com/energye/energy/v2/cmd/internal/tools/rawhttp"
	"github.com/pterm/pterm"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// DownloadFile 下载文件
//
//	如果文件存在大小一样不再下载
func DownloadFile(url, localPath, proxy string, callback func(totalLength, processLength int64)) error {
	var (
		fsize   int64
		buf     = make([]byte, 1024*10)
		written int64
	)
	tmpFilePath := localPath + ".download"

	//headers := map[string][]string{}
	//headers["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/128.0.0.0"}

	options := &rawhttp.Options{
		Timeout:                0,
		FollowRedirects:        true,
		MaxRedirects:           50,
		AutomaticHostHeader:    true,
		AutomaticContentLength: true,
		Proxy:                  proxy,
	}

	client := rawhttp.NewClient(options)
	resp, err := client.Get(url)

	//resp, err := client.DoRaw("GET", url, "", headers, nil)
	if err != nil {
		return err
	}
	fsize = resp.ContentLength
	term.Section.Println("Download File Size: ", fsize)
	fileExist := IsExistAndSize(localPath, fsize)
	if fileExist {
		term.Section.Println("File already exists")
		return nil
	}

	file, err := os.Create(tmpFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if resp.Body == nil {
		return errors.New("body is null")
	}
	defer resp.Body.Close()
	total := 100
	_, fileName := filepath.Split(localPath)
	p, err := pterm.DefaultProgressbar.WithMaxWidth(80).WithTotal(total).WithTitle("Download " + fileName).Start()
	if err != nil {
		return err
	}
	var stop = func() {
		if p != nil {
			p.Stop()
			p = nil
		}
	}
	defer stop()
	var (
		count int
		cn    int
		nw    int
		read  = bufio.NewReader(resp.Body)
	)
	for {
		nr, er := read.Read(buf)
		if nr > 0 {
			nw, err = file.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if callback != nil {
				callback(fsize, written)
			} else {
				process := int((float64(written) / float64(fsize)) * 100)
				if process > count {
					count = process
					p.Add(1)
					cn++
				}
			}
			if err != nil {
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	if cn < total {
		p.Add(total - cn)
	}
	stop()
	if err == nil {
		file.Sync()
		file.Close()
		err = os.Rename(tmpFilePath, localPath)
		if err != nil {
			return err
		}
		inf, err := os.Stat(localPath)
		if err != nil {
			return err
		}
		term.Section.Println("File Size: ", inf.Size())
		// todo test
		if strings.Contains(localPath, "liblcl") {
			data, _ := ioutil.ReadFile(localPath)
			fmt.Println("data:", string(data))
		}
	}
	return err
}
