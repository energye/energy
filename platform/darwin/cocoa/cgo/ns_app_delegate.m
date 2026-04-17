#import <Cocoa/Cocoa.h>
#include "ns_app_delegate.h"

@interface AppDelegate : NSObject <NSApplicationDelegate>
@property (nonatomic, strong) id<NSApplicationDelegate> originalDelegate;
@end

@implementation AppDelegate

- (void)application:(NSApplication *)application openURLs:(NSArray<NSURL *> *)urls {
    NSMutableArray *items = [NSMutableArray array];
    for (NSURL *url in urls) {
        [items addObject:[url absoluteString]];
    }
    NSData *jsonData = [NSJSONSerialization dataWithJSONObject:items options:0 error:nil];
    NSString *json = [[NSString alloc] initWithData:jsonData encoding:NSUTF8StringEncoding];
    GoOpenURLsCallback((char *)[json UTF8String]);

    if (self.originalDelegate &&
        [self.originalDelegate respondsToSelector:@selector(application:openURLs:)]) {
        [self.originalDelegate application:application openURLs:urls];
    }
}

- (BOOL)application:(NSApplication *)application
continueUserActivity:(NSUserActivity *)userActivity
  restorationHandler:(void (^)(NSArray<id<NSUserActivityRestoring>> * _Nullable))restorationHandler {
//    NSLog(@"AppDelegate activityType = %@", userActivity.activityType);
    if ([userActivity.activityType isEqualToString:NSUserActivityTypeBrowsingWeb]) {
        NSURL *url = userActivity.webpageURL;
        if (url) {
            NSArray *items = @[[url absoluteString]];
            NSData *jsonData =  [NSJSONSerialization dataWithJSONObject:items options:0 error:nil];
            NSString *json = [[NSString alloc] initWithData:jsonData encoding:NSUTF8StringEncoding];
            GoOpenURLsCallback((char *)[json UTF8String]);
        }
    }
    if (self.originalDelegate &&
        [self.originalDelegate respondsToSelector: @selector(application:continueUserActivity:restorationHandler:)]) {
        return (BOOL)[self.originalDelegate application:application
                         continueUserActivity:userActivity
                           restorationHandler:restorationHandler];
    }
    return YES;
}

- (id)forwardingTargetForSelector:(SEL)aSelector {
    //NSLog(@"AppDelegate forwardingTargetForSelector: %@", NSStringFromSelector(aSelector));
    if (self.originalDelegate && [self.originalDelegate respondsToSelector:aSelector]) {
        return self.originalDelegate;
    }
    return [super forwardingTargetForSelector:aSelector];
}

- (BOOL)respondsToSelector:(SEL)aSelector {
    //NSLog(@"AppDelegate respondsToSelector: %@", NSStringFromSelector(aSelector));
    if ([super respondsToSelector:aSelector]) return YES;
    if (self.originalDelegate && [self.originalDelegate respondsToSelector:aSelector]) return YES;
    return NO;
}

@end

void InitAppDelegate(void) {;
    NSApplication *app = [NSApplication sharedApplication];
    if (app.delegate == nil) return;
    AppDelegate *delegate = [AppDelegate new];
    delegate.originalDelegate = app.delegate;
    NSLog(@"old delegate class = %@", NSStringFromClass([delegate.originalDelegate class]));
    [NSApp setDelegate:delegate];
    NSLog(@"new delegate class = %@", NSStringFromClass([[NSApp delegate] class]));
}