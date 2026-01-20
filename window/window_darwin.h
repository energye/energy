#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>
#import <WebKit/WebKit.h>
#import <QuartzCore/QuartzCore.h>

#include <stdlib.h>
#include <string.h>

#ifdef __cplusplus
extern "C" {
#endif

extern void GoLog(char* message);

void SetWindowRadius(void* nsWindow);
void SetWindowBackgroundColor(void* nsWindow, int r, int g, int b, int alpha);
void DragWindow(void* nsWindow);
NSVisualEffectView* SetWindowTransparent(void* nsWindow);
void SwitchFrostedMaterial(void* nsFrostedView, void* nsWindow, const char *nsAppearance);

#ifdef __cplusplus
}
#endif
