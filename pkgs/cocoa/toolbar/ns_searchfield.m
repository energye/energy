#import "config.h"
#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>

void* NewSearchField(void* nsDelegate, const char *placeholder, const char *tooltip, ControlProperty property) {
    if (!nsDelegate) {
        NSLog(@"[ERROR] NewTextField 必要参数为空");
        return nil;
    }
    MainToolbarDelegate *delegate = (MainToolbarDelegate*)nsDelegate;
    NSString *placeholderStr = placeholder ? [NSString stringWithUTF8String:placeholder] : nil;
    NSString *tooltipStr = tooltip ? [NSString stringWithUTF8String:tooltip] : nil;
    NSSearchField *searchField = [[NSSearchField alloc] init];
    searchField.placeholderString = placeholderStr;
    searchField.delegate = delegate;
    // 设置自动调整大小的属性
    [searchField setContentHuggingPriority:NSLayoutPriorityDefaultLow
                          forOrientation:NSLayoutConstraintOrientationHorizontal];
    [searchField setContentCompressionResistancePriority:NSLayoutPriorityDefaultLow
                                        forOrientation:NSLayoutConstraintOrientationHorizontal];
    ConfigureControl(searchField, tooltipStr, property);
    return (__bridge void*)(searchField);
}