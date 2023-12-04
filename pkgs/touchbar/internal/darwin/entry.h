#ifndef GO_TOUCH_BAR_ENTRY_H
#define GO_TOUCH_BAR_ENTRY_H

#import <Cocoa/Cocoa.h>

typedef enum AttachMode : NSUInteger {
    kMainWindow,
    kDebug
} AttachMode;

typedef struct InitResult {
  void* result;
  const char * err;
} InitResult;

typedef struct ErrorResult2 {
  const char * err;
} ErrorResult;

extern void handleEvent(void* me, char* event);

InitResult initTouchBar(AttachMode mode, const char* data, void* me);
ErrorResult runDebug(void* context);
ErrorResult updateTouchBar(void* context, const char* data);
ErrorResult destroyTouchBar(void* context);

#endif
