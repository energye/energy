//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	"encoding/json"
	"github.com/ebitengine/purego/objc"
	. "github.com/energye/energy/v3/platform/darwin/types"
	"unsafe"
)

type NSPasteboard struct {
	NSObject
}

func WrapNSPasteboard(data unsafe.Pointer) INSPasteboard {
	if data == nil {
		return nil
	}
	m := &NSPasteboard{}
	m.SetInstance(data)
	return m
}

func (m *NSPasteboard) Types() NSTypes {
	if m.Instance() == 0 {
		return nil
	}
	pasteboard := objc.ID(m.Instance())
	types := pasteboard.Send(objc.RegisterName("types"))
	if types == 0 {
		return nil
	}

	count := types.Send(objc.RegisterName("count"))
	if count == 0 {
		return nil
	}

	var result []string
	for i := uintptr(0); i < uintptr(count); i++ {
		typeObj := types.Send(objc.RegisterName("objectAtIndex:"), i)
		if typeObj != 0 {
			typeStr := typeObj.Send(objc.RegisterName("UTF8String"))
			if typeStr != 0 {
				result = append(result, goStringFromPtr(uintptr(typeStr)))
			}
		}
	}
	return result
}

func (m *NSPasteboard) PasteboardData() *PasteboardData {
	if m.Instance() == 0 {
		return nil
	}
	pasteboard := objc.ID(m.Instance())
	result := &PasteboardData{}
	readFilesAndURLs(result, pasteboard)
	readTexts(result, pasteboard)
	readImageData(result, pasteboard)
	return result
}

func readFilesAndURLs(data *PasteboardData, pasteboard objc.ID) {
	urlClass := objc.GetClass("NSURL")
	urlClasses := objc.ID(objc.GetClass("NSArray")).Send(
		objc.RegisterName("arrayWithObject:"),
		urlClass,
	)

	allURLs := pasteboard.Send(
		objc.RegisterName("readObjectsForClasses:options:"),
		urlClasses,
		objc.ID(objc.GetClass("NSDictionary")).Send(objc.RegisterName("dictionary")),
	)

	if allURLs == 0 {
		return
	}

	count := allURLs.Send(objc.RegisterName("count"))
	for i := uintptr(0); i < uintptr(count); i++ {
		url := allURLs.Send(objc.RegisterName("objectAtIndex:"), i)
		if url == 0 {
			continue
		}

		isFileURL := url.Send(objc.RegisterName("isFileURL"))
		if isFileURL != 0 {
			fileSystemPath := url.Send(objc.RegisterName("fileSystemRepresentation"))
			if fileSystemPath != 0 {
				pathStr := goStringFromPtr(uintptr(fileSystemPath))
				data.FilePaths = append(data.FilePaths, pathStr)
			}
		} else {
			absoluteString := url.Send(objc.RegisterName("absoluteString"))
			if absoluteString != 0 {
				urlStr := goStringFromObjcString(absoluteString)
				data.WebURLs = append(data.WebURLs, urlStr)
			}
		}
	}
}

func readTexts(data *PasteboardData, pasteboard objc.ID) {
	stringClass := objc.GetClass("NSString")
	stringClasses := objc.ID(objc.GetClass("NSArray")).Send(
		objc.RegisterName("arrayWithObject:"),
		stringClass,
	)

	texts := pasteboard.Send(
		objc.RegisterName("readObjectsForClasses:options:"),
		stringClasses,
		objc.ID(objc.GetClass("NSDictionary")).Send(objc.RegisterName("dictionary")),
	)

	if texts == 0 {
		return
	}

	count := texts.Send(objc.RegisterName("count"))
	for i := uintptr(0); i < uintptr(count); i++ {
		text := texts.Send(objc.RegisterName("objectAtIndex:"), i)
		if text != 0 {
			textStr := goStringFromObjcString(text)
			data.Texts = append(data.Texts, textStr)
		}
	}

	if len(data.Texts) > 0 {
		jsonData, _ := json.Marshal(data.Texts)
		data.PlainTexts = string(jsonData)
	}
}

func readImageData(data *PasteboardData, pasteboard objc.ID) {
	publicPNG := objc.ID(objc.GetClass("NSString")).Send(
		objc.RegisterName("stringWithUTF8String:"),
		"public.png",
	)

	imageData := pasteboard.Send(objc.RegisterName("dataForType:"), publicPNG)
	if imageData == 0 {
		tiffType := objc.ID(objc.GetClass("NSString")).Send(
			objc.RegisterName("stringWithUTF8String:"),
			"public.tiff",
		)
		imageData = pasteboard.Send(objc.RegisterName("dataForType:"), tiffType)
	}

	if imageData != 0 {
		length := imageData.Send(objc.RegisterName("length"))
		bytes := imageData.Send(objc.RegisterName("bytes"))
		if bytes != 0 && length != 0 {
			data.ImageData = goBytesFromPtr(uintptr(bytes), uintptr(length))
		}
	}
}

func goStringFromPtr(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&ptr))
}

func goStringFromObjcString(objcStr objc.ID) string {
	if objcStr == 0 {
		return ""
	}
	utf8Ptr := objcStr.Send(objc.RegisterName("UTF8String"))
	if utf8Ptr == 0 {
		return ""
	}
	return goStringFromPtr(uintptr(utf8Ptr))
}

func goBytesFromPtr(ptr uintptr, length uintptr) []byte {
	if ptr == 0 || length == 0 {
		return nil
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(ptr)), length)
}
