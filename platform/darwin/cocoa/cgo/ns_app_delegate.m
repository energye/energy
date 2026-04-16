#import <Cocoa/Cocoa.h>

extern void GoOpenURLsCallback(char* path);

static id oldDelegate = nil;

@interface AppDelegate : NSObject <NSApplicationDelegate>
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

    if (oldDelegate &&
        [oldDelegate respondsToSelector:@selector(application:openURLs:)]) {
        [oldDelegate application:application openURLs:urls];
    }
}

- (BOOL)application:(NSApplication *)application
continueUserActivity:(NSUserActivity *)userActivity
 restorationHandler:(void (^)(NSArray<id<NSUserActivityRestoring>> * _Nullable))restorationHandler {
    if ([userActivity.activityType isEqualToString:NSUserActivityTypeBrowsingWeb]) {
        NSURL *url = userActivity.webpageURL;
        if (url) {
            NSArray *items = @[[url absoluteString]];
            NSData *jsonData =  [NSJSONSerialization dataWithJSONObject:items options:0 error:nil];
            NSString *json = [[NSString alloc] initWithData:jsonData encoding:NSUTF8StringEncoding];
            GoOpenURLsCallback((char *)[json UTF8String]);
        }
    }
    if (oldDelegate &&
        [oldDelegate respondsToSelector:
            @selector(application:continueUserActivity:restorationHandler:)]) {
        return (BOOL)[oldDelegate application:application
                         continueUserActivity:userActivity
                           restorationHandler:restorationHandler];
    }

    return YES;
}

- (id)forwardingTargetForSelector:(SEL)aSelector {
    //NSLog(@"AppDelegate forwardingTargetForSelector: %@", NSStringFromSelector(aSelector));
    if (oldDelegate && [oldDelegate respondsToSelector:aSelector]) {
        return oldDelegate;
    }
    return [super forwardingTargetForSelector:aSelector];
}

- (BOOL)respondsToSelector:(SEL)aSelector {
    //NSLog(@"AppDelegate respondsToSelector: %@", NSStringFromSelector(aSelector));
    if ([super respondsToSelector:aSelector]) return YES;
    if (oldDelegate && [oldDelegate respondsToSelector:aSelector]) return YES;
    return NO;
}

@end

void InitAppDelegate(void) {
    oldDelegate = [NSApp delegate];
    //NSLog(@"old delegate class = %@", NSStringFromClass([oldDelegate class]));
    AppDelegate *delegate = [AppDelegate new];
    [NSApp setDelegate:delegate];
    //NSLog(@"new delegate class = %@", NSStringFromClass([[NSApp delegate] class]));
}