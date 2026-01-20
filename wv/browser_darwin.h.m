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
