#import "config.h"
#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>

void* NewButton(void* nsDelegate, const char *title, const char *tooltip, ControlProperty property) {
    if (!nsDelegate || !title) {
        NSLog(@"[ERROR] NewButton 必要参数为空");
        return nil;
    }
    MainToolbarDelegate *delegate = (MainToolbarDelegate*)nsDelegate;
    NSString *titleStr = [NSString stringWithUTF8String:title];
    NSString *tooltipStr = tooltip ? [NSString stringWithUTF8String:tooltip] : nil;
    NSButton *button = [NSButton buttonWithTitle:titleStr target:delegate action:@selector(buttonClicked:)];
    button.bezelStyle = property.bezelStyle;
    ConfigureControl(button, tooltipStr, property);
    return (__bridge void*)(button);
}

void* NewImageButton(void* nsDelegate, NSImage *buttonImage, const char *tooltip, ControlProperty property) {
    MainToolbarDelegate *delegate = (MainToolbarDelegate*)nsDelegate;
    NSString *tooltipStr = tooltip ? [NSString stringWithUTF8String:tooltip] : nil;
    NSButton *button = [NSButton buttonWithImage:buttonImage
                                          target:delegate
                                          action:@selector(buttonClicked:)];
    button.bezelStyle = property.bezelStyle;
    button.imagePosition = NSImageOnly;
    ConfigureControl(button, tooltipStr, property);
    return button;
}

void* NewImageButtonFormImage(void* nsDelegate, const char *image, const char *tooltip, ControlProperty property) {
    NSString *imageNameStr = [NSString stringWithUTF8String:image];
    NSImage *buttonImage = nil;
    // 首先尝试从文件路径加载图像
    NSURL *imageURL = [NSURL fileURLWithPath:imageNameStr];
    if (imageURL) {
        buttonImage = [[NSImage alloc] initWithContentsOfURL:imageURL];
    }
    // 如果文件加载失败，尝试使用系统符号
    if (!buttonImage) {
        buttonImage = [NSImage imageWithSystemSymbolName:imageNameStr accessibilityDescription:nil];
    }
    // 如果仍然没有图像，使用默认图像
    if (!buttonImage) {
        buttonImage = [NSImage imageNamed:NSImageNameActionTemplate];
    }
    return NewImageButton(nsDelegate, buttonImage, tooltip, property);
}

void* NewImageButtonFormBytes(void* nsDelegate, const uint8_t* data, size_t length, const char *tooltip, ControlProperty property) {
    NSImage *buttonImage = imageFromBytes(data, length);
    return NewImageButton(nsDelegate, buttonImage, tooltip, property);
}

// 动态更换按钮图片（使用NSImage）
void SetButtonImage(void* buttonPtr, NSImage* newImage) {
    if (!buttonPtr || !newImage) {
        NSLog(@"[ERROR] SetButtonImage 参数为空");
        return;
    }

    NSButton *button = (__bridge NSButton *)buttonPtr;
    dispatch_async(dispatch_get_main_queue(), ^{
        [button setImage:newImage];
    });
}

// 动态更换按钮图片（使用图像文件路径或系统符号名称）
void SetButtonImageFromPath(void* buttonPtr, const char* imagePath) {
    if (!buttonPtr || !imagePath) {
        NSLog(@"[ERROR] SetButtonImageFromPath 参数为空");
        return;
    }
    NSString *imageNameStr = [NSString stringWithUTF8String:imagePath];
    NSImage *newImage = nil;
    // 首先尝试从文件路径加载图像
    NSURL *imageURL = [NSURL fileURLWithPath:imageNameStr];
    if (imageURL) {
        newImage = [[NSImage alloc] initWithContentsOfURL:imageURL];
    }
    // 如果文件加载失败，尝试使用系统符号
    if (!newImage) {
        if (@available(macOS 11.0, *)) {
            newImage = [NSImage imageWithSystemSymbolName:imageNameStr accessibilityDescription:nil];
        }
    }
    // 如果仍然没有图像，尝试从应用程序资源加载
    if (!newImage) {
        newImage = [NSImage imageNamed:imageNameStr];
    }
    // 如果仍然没有图像，使用默认图像
    if (!newImage) {
        newImage = [NSImage imageNamed:NSImageNameActionTemplate];
        NSLog(@"[WARNING] 使用默认图像替代: %@", imageNameStr);
    }

    SetButtonImage(buttonPtr, newImage);
}

// 动态更换按钮图片（使用字节数据）
void SetButtonImageFromBytes(void* buttonPtr, const uint8_t* data, size_t length) {
    if (!buttonPtr || !data || length == 0) {
        NSLog(@"[ERROR] SetButtonImageFromBytes 参数为空");
        return;
    }
    NSImage *newImage = imageFromBytes(data, length);
    if (!newImage) {
        NSLog(@"[ERROR] 无法从字节数据创建图像");
        return;
    }
    SetButtonImage(buttonPtr, newImage);
}
