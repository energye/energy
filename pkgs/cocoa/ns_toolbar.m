#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>
#import "ns_toolbar.h"

// 创建并配置窗口工具栏
void CreateToolbar(unsigned long nsWindowHandle, ToolbarConfiguration config) {
    NSLog(@"CreateToolbar");
    NSWindow *window = (__bridge NSWindow *)(void *)nsWindowHandle;
    NSToolbar *toolbar = [[NSToolbar alloc] initWithIdentifier:@"ENERGY.ToolBar"];
    [toolbar autorelease];
    toolbar.showsBaselineSeparator = config.ShowSeparator;
    window.toolbar = toolbar;
}
