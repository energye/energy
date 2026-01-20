#import "Cocoa/Cocoa.h"
#import <WebKit/WebKit.h>
#import <QuartzCore/QuartzCore.h>
#import "cocoa.h"
#import "ns_window.h"

// 最大化

void WindowMaximize(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"窗口不可用或不可调整大小，无法执行最大化");
        return;
    }
    //if (!window.isZoomed) {
        [window zoom:nil];
    //}
}

void WindowRestore(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"窗口为 nil");
        return;
    }
    //if (window.isZoomed) {
        [window zoom:nil];
    //}
}

// 最小化

void WindowMinimized(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"窗口为 nil");
        return;
    }
    if (![window isMiniaturized]) {
        [window miniaturize:nil];
    }
}

void WindowExitMinimized(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"窗口为 nil");
        return;
    }
    if ([window isMiniaturized]) {
        [window deminiaturize:nil];
    }
}

// 全屏

void WindowEnterFullScreen(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"窗口为 nil");
        return;
    }
    [window toggleFullScreen:nil];
}

void WindowExitFullScreen(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"窗口为 nil");
        return;
    }
    [window toggleFullScreen:nil];
}

NSString* NewNSString(const char* string) {
    if (string != nil) {
        return [NSString stringWithUTF8String:string];
    }
    return nil;
}

// 设置无边框窗口
void SetWindowRadius(void* nsWindow, float radius) {
    NSWindow* window = (NSWindow*)nsWindow;
	NSView* contentView = window.contentView;
	[contentView setWantsLayer:YES]; // view.setWantsLayer(true);
	CALayer* layer = contentView.layer; // contentView.layer
	window.backgroundColor = [NSColor clearColor]; // window.setBackgroundColor(NSColor.clearColor);
	layer.backgroundColor = [NSColor whiteColor].CGColor; // layer.setBackgroundColor(NSColor.whiteColor.CGColor);
	layer.cornerRadius = radius; // layer.setCornerRadius(8.0);
	layer.masksToBounds = YES;
}

// 设置窗口背景色
void SetWindowBackgroundColor(void* nsWindow, int r, int g, int b, int alpha) {
	[(NSWindow*)nsWindow setBackgroundColor:[NSColor colorWithCalibratedRed:r/255.0 green:g/255.0 blue:b/255.0 alpha:alpha/255.0]];
}

// 拖拽窗口
void DragWindow(void* nsWindow) {
    NSEvent *currentMouseEvent = [NSApp currentEvent];

    if (!currentMouseEvent) {
        NSLog(@"DragWindow 获取当前事件失败：事件为 nil");
        return;
    }
    if (currentMouseEvent.type != NSEventTypeLeftMouseDown) {
        NSLog(@"DragWindow 获取当前事件失败：非左键按下事件");
        return;
    }
    //NSWindow* window = (NSWindow*)nsWindow;
	NSWindow* window = [currentMouseEvent window];
 	//NSWindow *window = [NSApp keyWindow];
    if (!window) {
        NSLog(@"DragWindow 获取当前事件窗口失败");
        return;
    }
    [window performWindowDragWithEvent:currentMouseEvent];
}

// 设置窗口透明属性
NSVisualEffectView* SetWindowTransparent(void* nsWindow) {
	NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"SetWindowTransparent window nil");
        return nil;
    }
	NSView *contentView = [window contentView];

	NSVisualEffectView *frostedView = [[NSVisualEffectView alloc] initWithFrame:[contentView bounds]];
	[frostedView setAutoresizingMask:NSViewWidthSizable | NSViewHeightSizable];
	[frostedView setBlendingMode:NSVisualEffectBlendingModeBehindWindow];
	[frostedView setState:NSVisualEffectStateActive];

	[contentView addSubview:frostedView positioned:NSWindowBelow relativeTo:nil];
    return frostedView;
}

// 切换窗口磨砂材质
void SwitchFrostedMaterial(void* nsFrostedView, void* nsWindow, const char *nsAppearance) {
	NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"SwitchFrostedMaterial window nil");
        return;
    }
	NSVisualEffectView* frostedView = (NSVisualEffectView*)nsFrostedView;
    if (!frostedView) {
        NSLog(@"SwitchFrostedMaterial frostedView nil");
        return;
    }
    NSString* appearance = NewNSString(nsAppearance);
    if (appearance != nil) {
        NSAppearance *nsAppearanceName = [NSAppearance appearanceNamed:appearance];
        [window setAppearance:nsAppearanceName];
    }
}

void WindowAddSubview(void* nsWindow, void* nsView) {
	NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"WindowAddSubview window nil");
        return;
    }
	NSView* view = (NSView*)nsView;
    if (!view) {
        NSLog(@"WindowAddSubview view nil");
        return;
    }
	NSView* contentView = window.contentView;
	[contentView addSubview:view];
	[view setAutoresizingMask: NSViewWidthSizable|NSViewHeightSizable];
	CGRect contentViewBounds = [contentView bounds];
	[view setFrame:contentViewBounds];
}