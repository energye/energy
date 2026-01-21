#import "Cocoa/Cocoa.h"
#import <WebKit/WebKit.h>
#import <QuartzCore/QuartzCore.h>
#import "cocoa.h"
#import "ns_window.h"

@implementation TWindowDelegate

- (void)dealloc {
//    NSLog(@"TWindowDelegate dealloc");
    [[NSNotificationCenter defaultCenter] removeObserver:self];
    [super dealloc];
}

- (void)attachToWindow:(NSWindow *)window withCallback:(TEventCallback ) callback {
//    NSLog(@"TWindowDelegate attachToWindow");
    self.window = window;
    self._callback = callback;
    self.originalDelegate = window.delegate;
    window.delegate = self;
}

#pragma mark - Window Delegate Impl

- (id)windowWillReturnFieldEditor:(NSWindow *)sender toObject:(id)client {
//    NSLog(@"TWindowDelegate windowWillReturnFieldEditor");
    if (self.originalDelegate && [self.originalDelegate respondsToSelector:@selector(windowWillReturnFieldEditor:toObject:)]) {
        return [self.originalDelegate windowWillReturnFieldEditor:sender toObject:client];
    }
    return nil;
}

- (void)windowDidBecomeKey:(NSNotification *)notification {
//    NSLog(@"TWindowDelegate windowDidBecomeKey");
    if (self.originalDelegate && [self.originalDelegate respondsToSelector:@selector(windowDidBecomeKey:)]) {
        [self.originalDelegate windowDidBecomeKey:notification];
    }
}

- (void)windowDidResignKey:(NSNotification *)notification {
//    NSLog(@"TWindowDelegate windowDidResignKey");
    if (self.originalDelegate && [self.originalDelegate respondsToSelector:@selector(windowDidResignKey:)]) {
        [self.originalDelegate windowDidResignKey:notification];
    }
}

- (void)windowDidMiniaturize:(NSNotification *)notification {
//    NSLog(@"TWindowDelegate windowDidMiniaturize");
    if (self.originalDelegate && [self.originalDelegate respondsToSelector:@selector(windowDidMiniaturize:)]) {
        [self.originalDelegate windowDidMiniaturize:notification];
    }
}

- (void)windowDidDeminiaturize:(NSNotification *)notification {
//    NSLog(@"TWindowDelegate windowDidMiniaturize");
    if (self.originalDelegate && [self.originalDelegate respondsToSelector:@selector(windowDidDeminiaturize:)]) {
        [self.originalDelegate windowDidDeminiaturize:notification];
    }
}

- (void)windowDidResize:(NSNotification *)notification {
//    NSLog(@"TWindowDelegate windowDidResize");
    NSWindow *window = notification.object;
    if (self.originalDelegate && [self.originalDelegate respondsToSelector:@selector(windowDidResize:)]) {
        [self.originalDelegate windowDidResize:notification];
    }
}

- (void)windowWillEnterFullScreen:(NSNotification *)notification {
//    NSLog(@"TWindowDelegate windowWillEnterFullScreen");
    if (self.originalDelegate && [self.originalDelegate respondsToSelector:@selector(windowWillEnterFullScreen:)]) {
        [self.originalDelegate windowWillEnterFullScreen:notification];
    }
    if (self._callback) {
        NSString *eventId = [NSString stringWithFormat:@"%d_%p", TWindowEventEnterFullScreen, self.window];
        TCallbackContext *context = CreateCallbackContext(eventId, @"", -1, nil, self.window);
        GoArguments *result;
        @try{
            result = self._callback(context);
        } @finally {
            if (result) {
                FreeGoArguments(result);
            }
            FreeCallbackContext(context);
        }
    }
}

- (void)windowDidExitFullScreen:(NSNotification *)notification {
//     NSLog(@"TWindowDelegate windowDidExitFullScreen");
     if (self.originalDelegate && [self.originalDelegate respondsToSelector:@selector(windowDidExitFullScreen:)]) {
         [self.originalDelegate windowDidExitFullScreen:notification];
     }
    if (self._callback) {
        NSString *eventId = [NSString stringWithFormat:@"%d_%p", TWindowEventExitFullScreen, self.window];
        TCallbackContext *context = CreateCallbackContext(eventId, @"", -1, nil, self.window);
        GoArguments *result;
        @try{
            result = self._callback(context);
        } @finally {
            if (result) {
                FreeGoArguments(result);
            }
            FreeCallbackContext(context);
        }
    }
 }

- (NSApplicationPresentationOptions)window:(NSWindow *)window willUseFullScreenPresentationOptions:(NSApplicationPresentationOptions)proposedOptions {
//    NSLog(@"TWindowDelegate: windowWillUseFullScreenPresentationOptions");
    NSApplicationPresentationOptions options = NSApplicationPresentationAutoHideToolbar | NSApplicationPresentationAutoHideMenuBar | NSApplicationPresentationFullScreen;
    if (self._callback) {
        NSString *eventId = [NSString stringWithFormat:@"%d_%p", TWindowEventWillUseFullScreenPresentationOptions, self.window];
        TCallbackContext *context = CreateCallbackContext(eventId, @"", -1, nil, self.window);
        GoArguments *result;
        @try{
            result = self._callback(context);
            if(result){
                int resultOptions = GetIntFromGoArguments(result, 0);
                NSLog(@"willUseFullScreenPresentationOptions %d %d",  result->Count, resultOptions);
                for (int i = 0; i < result->Count; i++) {
                    //
                }
            }
        } @finally {
            if (result) {
                FreeGoArguments(result);
            }
            FreeCallbackContext(context);
        }
    }
    return options;
}

#pragma mark - Toolbar Delegate Impl

- (NSArray<NSString *> *)toolbarDefaultItemIdentifiers:(NSToolbar *)toolbar {
    NSLog(@"toolbarDefaultItemIdentifiers");
    NSMutableArray *identifiers = [NSMutableArray array];
    return identifiers;
}

- (NSArray<NSString *> *)toolbarAllowedItemIdentifiers:(NSToolbar *)toolbar {
    NSLog(@"toolbarAllowedItemIdentifiers");
    NSMutableArray *identifiers = [NSMutableArray array];
    return identifiers;
}

- (NSToolbarItem *)toolbar:(NSToolbar *)toolbar itemForItemIdentifier:(NSToolbarItemIdentifier)itemIdentifier willBeInsertedIntoToolbar:(BOOL)flag {
    NSLog(@"toolbarItemIdentifier: %@", itemIdentifier);
    if ([itemIdentifier isEqualToString:NSToolbarFlexibleSpaceItemIdentifier]) {
        return [[NSToolbarItem alloc] initWithItemIdentifier:NSToolbarFlexibleSpaceItemIdentifier];
    }
    if ([itemIdentifier isEqualToString:NSToolbarSpaceItemIdentifier]) {
        return [[NSToolbarItem alloc] initWithItemIdentifier:NSToolbarSpaceItemIdentifier];
    }
    return nil;
}

@end


// 公共函数

TWindowDelegate* CreateWindowDelegate(void* nsWindow, TEventCallback callback) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"CreateWindowDelegate window is nil");
        return nil;
    }
    TWindowDelegate *windowDelegate = [[TWindowDelegate alloc] init];
    [windowDelegate attachToWindow:window withCallback:callback];
    return windowDelegate;
}

// 最大化

void WindowMaximize(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"WindowMaximize window is nil");
        return;
    }
    //if (!window.isZoomed) {
        [window zoom:nil];
    //}
}

void WindowRestore(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"WindowRestore window is nil");
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
        NSLog(@"WindowMinimized window is nil");
        return;
    }
    if (![window isMiniaturized]) {
        [window miniaturize:nil];
    }
}

void WindowExitMinimized(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"WindowExitMinimized window is nil");
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
        NSLog(@"WindowEnterFullScreen window is nil");
        return;
    }
    [window toggleFullScreen:nil];
}

void WindowExitFullScreen(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"WindowExitFullScreen window is nil");
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
        NSLog(@"DragWindow currentMouseEvent is nil");
        return;
    }
    if (currentMouseEvent.type != NSEventTypeLeftMouseDown) {
        NSLog(@"DragWindow currentMouseEvent not left down");
        return;
    }
    //NSWindow* window = (NSWindow*)nsWindow;
	NSWindow* window = [currentMouseEvent window];
 	//NSWindow *window = [NSApp keyWindow];
    if (!window) {
        NSLog(@"DragWindow window is nil");
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

void WindowAddSubview(void* nsWindow, void* nsView, float x, float y, float width, float height) {
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
	[view setAutoresizingMask: NSViewWidthSizable | NSViewHeightSizable];
    CGRect viewFrame = CGRectMake(x, y, width, height);
	[view setFrame:viewFrame];
}