#import "Cocoa/Cocoa.h"
#import <WebKit/WebKit.h>
#import <QuartzCore/QuartzCore.h>
#import "browser_darwin.h"

// 设置webview透明属性
void SetWebviewTransparent(void* nsWebview, int isTransparent) {
	WKWebView* webview = (WKWebView*)nsWebview;
    if (!webview) {
        NSLog(@"SetWebviewTransparent webview nil");
        return;
    }
//   NSLog(@"SetWebviewTransparent %d", isTransparent);
   [webview setValue:[NSNumber numberWithBool:isTransparent] forKey:@"drawsBackground"];
}

void UpdateWebviewBounds(void* nsWindow, void* nsWebview, float x, float y, float width, float height) {
	NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"UpdateWebviewBounds window nil");
        return;
    }
	WKWebView* webview = (WKWebView*)nsWebview;
    if (!webview) {
        NSLog(@"WindowAddSubview webview nil");
        return;
    }
	NSView* contentView = window.contentView;
//	CGRect viewFrame = [contentView bounds];
    CGRect viewFrame = CGRectMake(x, y, width, height);
//	NSLog(@"UpdateWebviewBounds contentView bounds: %@", NSStringFromRect(viewFrame));
	[webview setFrame:viewFrame];
}