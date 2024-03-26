//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import "github.com/energye/energy/v2/lcl"

// AsStream Convert a pointer object to an existing class object
func AsStream(obj uintptr) IStream {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	stream := new(Stream)
	SetObjectInstance(stream, instance)
	return stream
}

// AsUnknown Convert a pointer object to an existing class object
func AsUnknown(obj uintptr) IUnknown {
	return lcl.AsUnknown(obj)
}

// AsDataObject Convert a pointer object to an existing class object
func AsDataObject(obj uintptr) IDataObject {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	dataObject := new(DataObject)
	SetObjectInstance(dataObject, instance)
	return dataObject
}

// AsCoreWebView2WebResourceRequest Convert a pointer object to an existing class object
func AsCoreWebView2WebResourceRequest(obj uintptr) ICoreWebView2WebResourceRequestRef {
	return AsCoreWebView2WebResourceRequestRef(obj)
}
