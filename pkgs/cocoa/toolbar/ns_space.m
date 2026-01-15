#import "config.h"
#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>

void AddToolbarFlexibleSpace(void* nsToolbar) {
    if (!nsToolbar) {
        NSLog(@"[ERROR] AddToolbarFlexibleSpace 必要参数为空");
        return;
    }
    NSToolbar *toolbar = (NSToolbar*)nsToolbar;
    [toolbar insertItemWithItemIdentifier:NSToolbarFlexibleSpaceItemIdentifier atIndex:toolbar.items.count];
}

void AddToolbarSpace(void* nsToolbar) {
    if (!nsToolbar) {
        NSLog(@"[ERROR] AddToolbarSpace 必要参数为空");
        return;
    }
    NSToolbar *toolbar = (NSToolbar*)nsToolbar;
    [toolbar insertItemWithItemIdentifier:NSToolbarSpaceItemIdentifier atIndex:toolbar.items.count];
}

//void AddToolbarSpaceByWidth(unsigned long nsWindowHandle, CGFloat width) {
//    NSWindow *window = (__bridge NSWindow *)(void *)nsWindowHandle;
//    MainToolbarDelegate *delegate = objc_getAssociatedObject(window, &kToolbarDelegateKey);
//
//    // 创建固定空格标识符
//    NSString *spaceIdentifier = [NSString stringWithFormat:@"FixedSpace_%.0f", width];
//
//    // 创建固定宽度的视图
//    NSView *spaceView = [[NSView alloc] initWithFrame:NSMakeRect(0, 0, width, 1)];
//    spaceView.translatesAutoresizingMaskIntoConstraints = NO;  // 关闭自动尺寸调整
//    [spaceView.widthAnchor constraintEqualToConstant:width].active = YES;
//
//    // 添加到委托
//    [delegate addControl:spaceView forIdentifier:spaceIdentifier withProperty:CreateDefaultControlProperty()];
//
//    // 添加到工具栏
//    [window.toolbar insertItemWithItemIdentifier:spaceIdentifier atIndex:window.toolbar.items.count];
//}