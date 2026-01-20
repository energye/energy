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
void AddSubviewWebview(void* nsWindow, void* nsWebview);
void UpdateWebviewBounds(void* nsWindow, void* nsWebview);

#ifdef __cplusplus
}
#endif
