#import "notification.h"
#include <Foundation/Foundation.h>
#import <Cocoa/Cocoa.h>
#import <UserNotifications/UserNotifications.h>

#pragma mark - Version & Bundle Checks

bool isNotificationAvailable(void) {
    if (@available(macOS 10.15, *)) {
        return true;
    }
    return false;
}

bool checkBundleIdentifier(void) {
    NSString *bundleID = [[NSBundle mainBundle] bundleIdentifier];
    return (bundleID != nil && bundleID.length > 0);
}

#pragma mark - External Callbacks

extern void onCallbackResult(int channelID, bool success, const char* error);
extern void didReceiveNotificationResponse(const char *jsonPayload, const char* error);

#pragma mark - Notification Center Delegate

@interface EnergyNotificationDelegate : NSObject <UNUserNotificationCenterDelegate>
@end

@implementation EnergyNotificationDelegate

- (void)userNotificationCenter:(UNUserNotificationCenter *)center
       willPresentNotification:(UNNotification *)notification
         withCompletionHandler:(void (^)(UNNotificationPresentationOptions))completionHandler {

    UNNotificationPresentationOptions options;

    if (@available(macOS 11.0, *)) {
        // macOS Big Sur+: Use modern presentation options
        options = UNNotificationPresentationOptionList |
                  UNNotificationPresentationOptionBanner |
                  UNNotificationPresentationOptionSound;
    } else {
        // macOS Catalina (10.15): Use legacy alert option
        #pragma clang diagnostic push
        #pragma clang diagnostic ignored "-Wdeprecated-declarations"
        options = UNNotificationPresentationOptionAlert |
                  UNNotificationPresentationOptionSound;
        #pragma clang diagnostic pop
    }

    completionHandler(options);
}

- (void)userNotificationCenter:(UNUserNotificationCenter *)center
didReceiveNotificationResponse:(UNNotificationResponse *)response
         withCompletionHandler:(void (^)(void))completionHandler {

    // Build response payload
    NSMutableDictionary *payload = [NSMutableDictionary dictionaryWithCapacity:8];

    payload[@"id"] = response.notification.request.identifier;
    payload[@"actionIdentifier"] = response.actionIdentifier;
    payload[@"title"] = response.notification.request.content.title ?: @"";
    payload[@"body"] = response.notification.request.content.body ?: @"";

    // Optional fields
    if (response.notification.request.content.categoryIdentifier.length > 0) {
        payload[@"categoryIdentifier"] = response.notification.request.content.categoryIdentifier;
    }

    if (response.notification.request.content.subtitle.length > 0) {
        payload[@"subtitle"] = response.notification.request.content.subtitle;
    }

    if (response.notification.request.content.userInfo.count > 0) {
        payload[@"userInfo"] = response.notification.request.content.userInfo;
    }

    // Text input response
    if ([response isKindOfClass:[UNTextInputNotificationResponse class]]) {
        UNTextInputNotificationResponse *textInput = (UNTextInputNotificationResponse *)response;
        if (textInput.userText.length > 0) {
            payload[@"userText"] = textInput.userText;
        }
    }

    // Serialize to JSON and callback to Go
    NSError *jsonError = nil;
    NSData *jsonData = [NSJSONSerialization dataWithJSONObject:payload
                                                       options:NSJSONWritingPrettyPrinted
                                                         error:&jsonError];

    if (jsonError) {
        NSString *errorMsg = [NSString stringWithFormat:@"JSON serialization failed: %@",
                              jsonError.localizedDescription];
        didReceiveNotificationResponse(NULL, [errorMsg UTF8String]);
    } else {
        NSString *jsonString = [[NSString alloc] initWithData:jsonData
                                                     encoding:NSUTF8StringEncoding];
        didReceiveNotificationResponse([jsonString UTF8String], NULL);
    }

    completionHandler();
}

@end

#pragma mark - Singleton Management

static EnergyNotificationDelegate *g_notificationDelegate = nil;
static dispatch_once_t g_delegateInitToken;

bool initializeNotificationCenter(void) {
    __block BOOL success = YES;

    dispatch_once(&g_delegateInitToken, ^{
        g_notificationDelegate = [[EnergyNotificationDelegate alloc] init];

        UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
        center.delegate = g_notificationDelegate;

        NSLog(@"[ENERGY] Notification center delegate initialized");
    });

    if (!g_notificationDelegate) {
        success = NO;
        NSLog(@"[ENERGY] Failed to initialize notification delegate");
    }

    return success;
}

#pragma mark - Authorization Management

void requestNotificationAuthorization(int channelID) {
    if (!initializeNotificationCenter()) {
        onCallbackResult(channelID, false, "requestNotificationAuthorization.initializeNotificationCenter: Failed to initialize the notification center");
        return;
    }

    UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
    UNAuthorizationOptions requestedOptions = UNAuthorizationOptionAlert |
                                              UNAuthorizationOptionSound |
                                              UNAuthorizationOptionBadge;

    [center requestAuthorizationWithOptions:requestedOptions
                          completionHandler:^(BOOL granted, NSError * _Nullable error) {
        if (error) {
            NSString *errorMsg = [NSString stringWithFormat:@"Authorization failed: %@",
                                  error.localizedDescription];
            onCallbackResult(channelID, false, [errorMsg UTF8String]);
            NSLog(@"[ENERGY] Notification authorization error: %@", error);
        } else {
            onCallbackResult(channelID, granted, NULL);
            NSLog(@"[ENERGY] Notification authorization %@", granted ? @"granted" : @"denied");
        }
    }];
}

void checkNotificationAuthorization(int channelID) {
    if (!initializeNotificationCenter()) {
        onCallbackResult(channelID, false, "checkNotificationAuthorization.initializeNotificationCenter: Failed to initialize the notification center");
        return;
    }

    UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
    [center getNotificationSettingsWithCompletionHandler:^(UNNotificationSettings *settings) {
        BOOL authorized = (settings.authorizationStatus == UNAuthorizationStatusAuthorized);
        onCallbackResult(channelID, authorized, NULL);

        NSLog(@"[ENERGY] Notification status: %ld", (long)settings.authorizationStatus);
    }];
}

#pragma mark - Notification Content Builder

static UNMutableNotificationContent* buildNotificationContent(const char *title,
                                                               const char *subtitle,
                                                               const char *body,
                                                               const char *data_json,
                                                               NSError **outError) {
    UNMutableNotificationContent *content = [[UNMutableNotificationContent alloc] init];

    content.title = [NSString stringWithUTF8String:title ?: ""];
    content.body = [NSString stringWithUTF8String:body ?: ""];
    content.sound = [UNNotificationSound defaultSound];

    // Optional subtitle
    if (subtitle && strlen(subtitle) > 0) {
        content.subtitle = [NSString stringWithUTF8String:subtitle];
    }

    // Parse custom data
    if (data_json && strlen(data_json) > 0) {
        NSString *jsonStr = [NSString stringWithUTF8String:data_json];
        NSData *jsonData = [jsonStr dataUsingEncoding:NSUTF8StringEncoding];

        NSError *parseError = nil;
        NSDictionary *userData = [NSJSONSerialization JSONObjectWithData:jsonData
                                                                 options:0
                                                                   error:&parseError];

        if (parseError) {
            *outError = parseError;
            return nil;
        }

        if (userData) {
            content.userInfo = userData;
        }
    }

    return content;
}

#pragma mark - Send Notifications

void sendNotification(int channelID,
                      const char *identifier,
                      const char *title,
                      const char *subtitle,
                      const char *body,
                      const char *data_json) {

    if (!initializeNotificationCenter()) {
        onCallbackResult(channelID, false, "sendNotification.initializeNotificationCenter: Failed to initialize the notification center");
        return;
    }

    // Build content
    NSError *contentError = nil;
    UNMutableNotificationContent *content = buildNotificationContent(title, subtitle, body,
                                                                      data_json, &contentError);
    if (contentError) {
        NSString *errorMsg = [NSString stringWithFormat:@"Content creation failed: %@",
                              contentError.localizedDescription];
        onCallbackResult(channelID, false, [errorMsg UTF8String]);
        return;
    }

    // Create request with immediate trigger
    NSString *requestID = [NSString stringWithUTF8String:identifier ?: ""];
    UNNotificationRequest *request = [UNNotificationRequest requestWithIdentifier:requestID
                                                                            content:content
                                                                             trigger:nil];

    // Submit to notification center
    UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
    [center addNotificationRequest:request
                 withCompletionHandler:^(NSError * _Nullable error) {
        if (error) {
            NSString *errorMsg = [NSString stringWithFormat:@"Send failed: %@",
                                  error.localizedDescription];
            onCallbackResult(channelID, false, [errorMsg UTF8String]);
            NSLog(@"[ENERGY] Notification send error: %@", error);
        } else {
            onCallbackResult(channelID, true, NULL);
            NSLog(@"[ENERGY] Notification sent: %@", requestID);
        }
    }];
}

void sendNotificationWithActions(int channelID,
                                 const char *identifier,
                                 const char *title,
                                 const char *subtitle,
                                 const char *body,
                                 const char *categoryId,
                                 const char *data_json) {

    if (!initializeNotificationCenter()) {
        onCallbackResult(channelID, false, "sendNotificationWithActions.initializeNotificationCenter: Failed to initialize the notification center");
        return;
    }

    // Build content
    NSError *contentError = nil;
    UNMutableNotificationContent *content = buildNotificationContent(title, subtitle, body,
                                                                      data_json, &contentError);
    if (contentError) {
        NSString *errorMsg = [NSString stringWithFormat:@"Content creation failed: %@",
                              contentError.localizedDescription];
        onCallbackResult(channelID, false, [errorMsg UTF8String]);
        return;
    }

    // Attach category for actions
    if (categoryId && strlen(categoryId) > 0) {
        content.categoryIdentifier = [NSString stringWithUTF8String:categoryId];
    }

    // Create and submit request
    NSString *requestID = [NSString stringWithUTF8String:identifier ?: ""];
    UNNotificationRequest *request = [UNNotificationRequest requestWithIdentifier:requestID
                                                                            content:content
                                                                             trigger:nil];

    UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
    [center addNotificationRequest:request
                 withCompletionHandler:^(NSError * _Nullable error) {
        if (error) {
            NSString *errorMsg = [NSString stringWithFormat:@"Send with actions failed: %@",
                                  error.localizedDescription];
            onCallbackResult(channelID, false, [errorMsg UTF8String]);
        } else {
            onCallbackResult(channelID, true, NULL);
            NSLog(@"[ENERGY] Notification with actions sent: %@", requestID);
        }
    }];
}

#pragma mark - Category Registration

void registerNotificationCategory(int channelID,
                                  const char *categoryId,
                                  const char *actions_json,
                                  bool hasReplyField,
                                  const char *replyPlaceholder,
                                  const char *replyButtonTitle) {

    if (!initializeNotificationCenter()) {
        onCallbackResult(channelID, false, "registerNotificationCategory.initializeNotificationCenter: Failed to initialize the notification center");
        return;
    }

    NSString *categoryID = [NSString stringWithUTF8String:categoryId ?: ""];

    // Parse actions from JSON
    NSMutableArray<UNNotificationAction *> *actions = [NSMutableArray array];

    if (actions_json && strlen(actions_json) > 0) {
        NSString *jsonStr = [NSString stringWithUTF8String:actions_json];
        NSData *jsonData = [jsonStr dataUsingEncoding:NSUTF8StringEncoding];

        NSError *parseError = nil;
        NSArray *actionsArray = [NSJSONSerialization JSONObjectWithData:jsonData
                                                                options:0
                                                                  error:&parseError];

        if (parseError) {
            NSString *errorMsg = [NSString stringWithFormat:@"Actions parse failed: %@",
                                  parseError.localizedDescription];
            onCallbackResult(channelID, false, [errorMsg UTF8String]);
            return;
        }

        // Build action objects
        for (NSDictionary *actionDict in actionsArray) {
            NSString *actionID = actionDict[@"id"];
            NSString *actionTitle = actionDict[@"title"];
            BOOL destructive = [actionDict[@"destructive"] boolValue];

            if (!actionID || !actionTitle) continue;

            UNNotificationActionOptions opts = UNNotificationActionOptionNone;
            if (destructive) {
                opts |= UNNotificationActionOptionDestructive;
            }

            UNNotificationAction *action = [UNNotificationAction actionWithIdentifier:actionID
                                                                                title:actionTitle
                                                                              options:opts];
            [actions addObject:action];
        }
    }

    // Add text input action if requested
    if (hasReplyField && replyPlaceholder && replyButtonTitle) {
        NSString *placeholder = [NSString stringWithUTF8String:replyPlaceholder];
        NSString *buttonTitle = [NSString stringWithUTF8String:replyButtonTitle];

        UNTextInputNotificationAction *textInputAction =
            [UNTextInputNotificationAction actionWithIdentifier:@"TEXT_REPLY"
                                                          title:buttonTitle
                                                        options:UNNotificationActionOptionNone
                                           textInputButtonTitle:buttonTitle
                                           textInputPlaceholder:placeholder];
        [actions addObject:textInputAction];
    }

    // Create category
    UNNotificationCategory *category = [UNNotificationCategory categoryWithIdentifier:categoryID
                                                                                actions:actions
                                                                      intentIdentifiers:@[]
                                                                                options:UNNotificationCategoryOptionNone];

    // Update categories
    UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
    [center getNotificationCategoriesWithCompletionHandler:^(NSSet<UNNotificationCategory *> *existing) {
        NSMutableSet *updated = [NSMutableSet setWithSet:existing];

        // Remove old category if exists
        for (UNNotificationCategory *cat in updated) {
            if ([cat.identifier isEqualToString:categoryID]) {
                [updated removeObject:cat];
                break;
            }
        }

        // Add new category
        [updated addObject:category];
        [center setNotificationCategories:updated];

        onCallbackResult(channelID, true, NULL);
        NSLog(@"[ENERGY] Category registered: %@", categoryID);
    }];
}

void removeNotificationCategory(int channelID, const char *categoryId) {
    NSString *categoryID = [NSString stringWithUTF8String:categoryId ?: ""];

    UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
    [center getNotificationCategoriesWithCompletionHandler:^(NSSet<UNNotificationCategory *> *existing) {
        NSMutableSet *updated = [NSMutableSet setWithSet:existing];

        BOOL found = NO;
        for (UNNotificationCategory *cat in updated) {
            if ([cat.identifier isEqualToString:categoryID]) {
                [updated removeObject:cat];
                found = YES;
                break;
            }
        }

        if (found) {
            [center setNotificationCategories:updated];
            onCallbackResult(channelID, true, NULL);
            NSLog(@"[ENERGY] Category removed: %@", categoryID);
        } else {
            NSString *errorMsg = [NSString stringWithFormat:@"Category not found: %@", categoryID];
            onCallbackResult(channelID, false, [errorMsg UTF8String]);
        }
    }];
}

#pragma mark - Notification Cleanup

void removeAllPendingNotifications(void) {
    UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
    [center removeAllPendingNotificationRequests];
    NSLog(@"[ENERGY] All pending notifications cleared");
}

void removePendingNotification(const char *identifier) {
    if (!identifier) return;

    NSString *requestID = [NSString stringWithUTF8String:identifier];
    UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
    [center removePendingNotificationRequestsWithIdentifiers:@[requestID]];
    NSLog(@"[ENERGY] Pending notification removed: %@", requestID);
}

void removeAllDeliveredNotifications(void) {
    UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
    [center removeAllDeliveredNotifications];
    NSLog(@"[ENERGY] All delivered notifications cleared");
}

void removeDeliveredNotification(const char *identifier) {
    if (!identifier) return;

    NSString *requestID = [NSString stringWithUTF8String:identifier];
    UNUserNotificationCenter *center = [UNUserNotificationCenter currentNotificationCenter];
    [center removeDeliveredNotificationsWithIdentifiers:@[requestID]];
    NSLog(@"[ENERGY] Delivered notification removed: %@", requestID);
}
