#import "Cocoa/Cocoa.h"
#import <WebKit/WebKit.h>
#import <QuartzCore/QuartzCore.h>
#import "browser_darwin.h"

// 设置webview透明属性
void SetWebviewTransparent(void* nsWebview, bool isTransparent) {
	WKWebView* webview = (WKWebView*)nsWebview;
    if (!webview) {
        NSLog(@"SetWebviewTransparent webview nil");
        return;
    }
   [webview setOpaque:NO];
   [webview setBackgroundColor:[NSColor clearColor]];
   [webview setValue:[NSNumber numberWithBool:isTransparent] forKey:@"drawsBackground"];
}
