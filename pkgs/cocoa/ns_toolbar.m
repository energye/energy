#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>
#import "ns_window.h"
#import "ns_toolbar.h"

// 创建并配置窗口工具栏
void CreateToolbar(void* nsWindow, void* nsDelegate, ToolbarConfiguration config) {
	NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"CreateToolbar window nil");
        return;
    }
	TWindowDelegate* delegate = (TWindowDelegate*)nsDelegate;
    NSToolbar *toolbar = [[NSToolbar alloc] initWithIdentifier:@"ENERGY.ToolBar"];
    [toolbar autorelease];
    toolbar.showsBaselineSeparator = config.ShowSeparator;
    toolbar.delegate = delegate;
    window.toolbar = toolbar;
}
