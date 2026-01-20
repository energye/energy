#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>
#import <WebKit/WebKit.h>
#import <QuartzCore/QuartzCore.h>

#include <stdlib.h>
#include <string.h>

#ifdef __cplusplus
extern "C" {
#endif

void WindowMaximize(void* nsWindow);
void WindowRestore(void* nsWindow);
void WindowMinimized(void* nsWindow);
void WindowExitMinimized(void* nsWindow);
void WindowEnterFullScreen(void* nsWindow);
void WindowExitFullScreen(void* nsWindow);
NSString* NewNSString(const char* string);
void SetWindowRadius(void* nsWindow, float radius);
void WindowAddSubview(void* nsWindow, void* nsView, float x, float y, float width, float height);
void SetWindowBackgroundColor(void* nsWindow, int r, int g, int b, int alpha);
void DragWindow(void* nsWindow);
NSVisualEffectView* SetWindowTransparent(void* nsWindow);
void SwitchFrostedMaterial(void* nsFrostedView, void* nsWindow, const char *nsAppearance);

#ifdef __cplusplus
}
#endif
