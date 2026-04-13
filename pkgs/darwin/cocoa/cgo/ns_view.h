#import <Cocoa/Cocoa.h>

#ifdef __cplusplus
extern "C" {
#endif

@interface CustomView : NSView

@property (strong, nonatomic) NSColor *backgroundColor;

@end

void* NewCustomView();

#ifdef __cplusplus
}
#endif