#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>
#import <WebKit/WebKit.h>
#import <QuartzCore/QuartzCore.h>

#include <stdlib.h>
#include <string.h>

#ifdef __cplusplus
extern "C" {
#endif

void SetWebviewTransparent(void* nsWebview, int isTransparent);
void UpdateWebviewBounds(void* nsWindow, void* nsWebview, float x, float y, float width, float height);
void WebViewBecomeFirstResponder(void* nsWebview);

void WebViewUndo(void* nsWebview);
void WebViewRedo(void* nsWebview);
void WebViewCut(void* nsWebview);
void WebViewCopy(void* nsWebview);
void WebViewPaste(void* nsWebview);
void WebViewSelectAll(void* nsWebview);

void WebViewRegisterPerformKeyEquivalentMethod(void* nsWebview);

#ifdef __cplusplus
}
#endif
