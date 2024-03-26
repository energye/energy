//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// 移植来自Delphi

package version

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"

	"io/ioutil"
)

const (
	sVersionStr = "%s (Version %d.%d.%d)"
	//sCName          = "hw.optional.x86_64"
	pListName       = "/System/Library/CoreServices/SystemVersion.plist"
	pListServerName = "/System/Library/CoreServices/ServerVersion.plist"
)

type cFDataRef struct {
	XMLName xml.Name `xml:"plist"`
	Dict    struct {
		XMLName xml.Name `xml:"dict"`
		Key     []string `xml:"key"`
		String  []string `xml:"string"`
	}
}

func openDataFromURL(listName string) (bool, []byte) {
	bs, err := ioutil.ReadFile(listName)
	if err != nil {
		return false, nil
	}
	return true, bs
}

func initOSVersion() {
	OSVersion.Platform = PfMacOS
	OSVersion.Name = "Mac OS X"

	validData, lData := openDataFromURL(pListServerName)
	if !validData {
		validData, lData = openDataFromURL(pListName)
	}
	if validData {
		lDataRef := cFDataRef{}
		if err := xml.Unmarshal(lData, &lDataRef); err == nil {
			for i, key := range lDataRef.Dict.Key {
				if key == "ProductName" {
					OSVersion.Name = lDataRef.Dict.String[i]
				}
				if key == "ProductVersion" {
					items := strings.Split(lDataRef.Dict.String[i], ".")
					if len(items) > 0 {
						OSVersion.Major, _ = strconv.Atoi(items[0])
					}
					if len(items) > 1 {
						OSVersion.Minor, _ = strconv.Atoi(items[1])
					}
					if len(items) > 2 {
						OSVersion.ServicePackMajor, _ = strconv.Atoi(items[2])
					}
					break
				}
			}
		}
	}
	OSVersion.fmtVerString = fmt.Sprintf(sVersionStr, OSVersion.Name, OSVersion.Major, OSVersion.Minor, OSVersion.ServicePackMajor)
}
