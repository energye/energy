#ifndef GO_TOUCH_BAR_WINDOW_CONTROLLER_H
#define GO_TOUCH_BAR_WINDOW_CONTROLLER_H

#import <Cocoa/Cocoa.h>

@interface WindowController : NSWindowController {}
- (id)initWithData:(const char *)data andHandler:(void (^)(char *))handler error:(NSError**)error;
- (NSError*)update:(NSTouchBar*)touchBar withData:(const char *)data;
@end

#endif
