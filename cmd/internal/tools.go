//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package internal

import (
	"github.com/energye/energy/v2/common"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	return v.(string)
}

func ToInt(v interface{}) int {
	return common.ValueToInt(v)
}

func ToRNilString(v interface{}, new string) string {
	if v == nil {
		return new
	}
	return v.(string)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		} else if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

// http 请求
func httpRequestGET(url string) ([]byte, error) {
	client := new(http.Client)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// Compare compare2 < compare1 = true
func Compare(compare1, compare2 string) bool {
	if compare1[0] == 'v' {
		compare1 = compare1[1:]
	}
	if compare2[0] == 'v' {
		compare2 = compare2[1:]
	}
	compare1 = strings.Split(compare1, "-")[0]
	compare2 = strings.Split(compare2, "-")[0]
	cv := strings.Split(compare1, ".")
	ev := strings.Split(compare2, ".")
	c0, _ := strconv.Atoi(cv[0])
	c1, _ := strconv.Atoi(cv[1])
	c2, _ := strconv.Atoi(cv[2])
	e0, _ := strconv.Atoi(ev[0])
	e1, _ := strconv.Atoi(ev[1])
	e2, _ := strconv.Atoi(ev[2])
	if e0 < c0 {
		return true
	}
	if e1 < c1 {
		return true
	}
	if e2 < c2 {
		return true
	}
	return false
}
