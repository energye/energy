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

void AddSubviewWebview(void* nsWindow, void* nsWebview) {
	NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"WindowAddSubview window nil");
        return;
    }
	WKWebView* webview = (WKWebView*)nsWebview;
    if (!webview) {
        NSLog(@"WindowAddSubview webview nil");
        return;
    }
	NSView* contentView = window.contentView;
	[contentView addSubview:webview];
	[webview setAutoresizingMask: NSViewWidthSizable | NSViewHeightSizable];
	CGRect contentViewBounds = [contentView bounds];
	[webview setFrame:contentViewBounds];

	// [frostedView setAutoresizingMask:NSViewWidthSizable | NSViewHeightSizable];
}