#import "config.h"
#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>

// 从字节数组创建NSImage
NSImage* imageFromBytes(const uint8_t* data, size_t length) {
    if (!data || length == 0) {
        return nil;
    }
    NSData* imageData = [NSData dataWithBytes:data length:length];
    if (!imageData) {
        return nil;
    }
    NSImage* image = [[NSImage alloc] initWithData:imageData];
    return image;
}