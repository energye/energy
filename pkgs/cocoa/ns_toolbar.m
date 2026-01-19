#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>
#import "ns_toolbar.h"

// 创建并配置窗口工具栏
void CreateToolbar(void* nsWindow, ToolbarConfiguration config) {
    NSLog(@"CreateToolbar");
	NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"CreateToolbar window nil");
        return;
    }
    NSToolbar *toolbar = [[NSToolbar alloc] initWithIdentifier:@"ENERGY.ToolBar"];
    [toolbar autorelease];
    toolbar.showsBaselineSeparator = config.ShowSeparator;
    window.toolbar = toolbar;
}
