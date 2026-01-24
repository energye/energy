//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin

package cocoa

/*
#cgo darwin CFLAGS: -DDARWIN -x objective-c
#cgo darwin LDFLAGS: -framework WebKit -framework Cocoa

#import <WebKit/WebKit.h>
#import <Cocoa/Cocoa.h>

// 粘贴板数据结构体（Go侧对应）
typedef struct {
    const char* filePathsJSON; 		// 文件路径数组JSON
    const char* webURLsJSON;   		// 网页URL数组JSON
    const char* textsJSON;     		// 文本数组JSON
    const void* imageData;       	// 图片二进制数据
    NSUInteger imageLength;      	// 图片数据长度
} PasteboardData;

// 通用读取粘贴板数据
PasteboardData ReadPasteboardData(void* nsPasteboard) {
    PasteboardData data = {NULL, NULL, NULL, NULL, 0};
    @autoreleasepool {
        NSPasteboard* pboard = (NSPasteboard*)nsPasteboard;
        if (!pboard) return data;

        // 读取文件/URL
        NSArray<Class>* urlClasses = @[[NSURL class]];
        NSArray<NSURL*>* allURLs = [pboard readObjectsForClasses:urlClasses options:@{}];
        NSMutableArray<NSString*>* filePaths = [NSMutableArray array];
        NSMutableArray<NSString*>* webURLs = [NSMutableArray array];

        for (NSURL* url in allURLs) {
            if ([url isFileURL]) {
                const char* fsPath = [url fileSystemRepresentation];
                NSString* pathStr = [[NSString alloc] initWithCString:fsPath encoding:NSUTF8StringEncoding];
                [filePaths addObject:pathStr];
            } else {
                [webURLs addObject:url.absoluteString];
            }
        }

        // 序列化文件路径/URL为JSON
        if (filePaths.count > 0) {
            NSData* fpJSON = [NSJSONSerialization dataWithJSONObject:filePaths options:0 error:nil];
			const char* utf8Str = [[[NSString alloc] initWithData:fpJSON encoding:NSUTF8StringEncoding] UTF8String];
            data.filePathsJSON = strdup(utf8Str);
        }
        if (webURLs.count > 0) {
            NSData* wuJSON = [NSJSONSerialization dataWithJSONObject:webURLs options:0 error:nil];
			const char* utf8Str = [[[NSString alloc] initWithData:wuJSON encoding:NSUTF8StringEncoding] UTF8String];
            data.webURLsJSON = strdup(utf8Str);
        }

        // 读取文本
        NSArray<Class>* stringClasses = @[[NSString class]];
        NSArray<NSString*>* texts = [pboard readObjectsForClasses:stringClasses options:@{}];
        if (texts.count > 0) {
            NSData* textJSON = [NSJSONSerialization dataWithJSONObject:texts options:0 error:nil];
			const char* utf8Str = [[[NSString alloc] initWithData:textJSON encoding:NSUTF8StringEncoding] UTF8String];
            data.textsJSON = strdup(utf8Str);
        }

        // 读取图片数据
        NSData* imageData = [pboard dataForType:@"public.png"];
        if (!imageData) imageData = [pboard dataForType:NSPasteboardTypeTIFF];
        if (imageData) {
            data.imageData = imageData.bytes;
            data.imageLength = imageData.length;
        }
    }
    return data;
}

void FreePasteboardData(PasteboardData data) {
    if (data.filePathsJSON) {
        free((void*)data.filePathsJSON);
    }
    if (data.webURLsJSON) {
        free((void*)data.webURLsJSON);
    }
    if (data.textsJSON) {
        free((void*)data.textsJSON);
    }
}

const char* PasteboardTypes(void* nsPasteboard) {
	NSPasteboard* pasteboard = (NSPasteboard*)nsPasteboard;
	if (!pasteboard) {
        NSLog(@"PasteboardTypes pasteboard is nil");
        return NULL;
    }
    NSArray<NSPasteboardType> *types = [pasteboard types];
    if (!types || types.count == 0) {
        return NULL;
    }
    @autoreleasepool {
		NSError *error = nil;
		NSData *jsonData = [NSJSONSerialization dataWithJSONObject:types options:0 error:&error];
		if (error || !jsonData) return NULL;
		NSString *jsonStr = [[NSString alloc] initWithData:jsonData encoding:NSUTF8StringEncoding];
		return [jsonStr UTF8String];
    }
}

*/
import "C"
import (
	"encoding/json"
	"unsafe"
)

type TNSPasteboard struct {
	data unsafe.Pointer
}

func WrapNSPasteboard(data uintptr) *TNSPasteboard {
	if data == 0 {
		return nil
	}
	return &TNSPasteboard{data: unsafe.Pointer(data)}
}

func (m *TNSPasteboard) Types() []string {
	if m.data == nil {
		return nil
	}
	cData := C.PasteboardTypes(m.data)
	if cData == nil {
		return nil
	}
	data := C.GoString(cData)
	var result []string
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil
	}
	return result
}

type PasteboardData struct {
	FilePaths  []string
	WebURLs    []string
	Texts      []string
	PlainTexts string
	ImageData  []byte
}

func (m *TNSPasteboard) PasteboardData() *PasteboardData {
	if m.data == nil {
		return nil
	}
	cData := C.ReadPasteboardData(m.data)
	defer C.FreePasteboardData(cData)
	result := &PasteboardData{}
	// 解析文件路径
	if cData.filePathsJSON != nil {
		fpJSON := C.GoString(cData.filePathsJSON)
		_ = json.Unmarshal([]byte(fpJSON), &result.FilePaths)
	}
	// 解析网页URL
	if cData.webURLsJSON != nil {
		wuJSON := C.GoString(cData.webURLsJSON)
		_ = json.Unmarshal([]byte(wuJSON), &result.WebURLs)
	}
	// 解析文本
	if cData.textsJSON != nil {
		textJSON := C.GoString(cData.textsJSON)
		_ = json.Unmarshal([]byte(textJSON), &result.Texts)
	}
	// 解析图片数据
	if cData.imageData != nil && cData.imageLength > 0 {
		result.ImageData = C.GoBytes(cData.imageData, C.int(cData.imageLength))
	}
	return result
}
