#import <Foundation/Foundation.h>

#ifdef __cplusplus
extern "C" {
#endif

bool isNotificationAvailable(void);
bool checkBundleIdentifier(void);
bool initializeNotificationCenter(void);
void requestNotificationAuthorization(int channelID);
void checkNotificationAuthorization(int channelID);
void sendNotification(int channelID, const char *identifier, const char *title, const char *subtitle, const char *body, const char *data_json);
void sendNotificationWithActions(int channelID, const char *identifier, const char *title, const char *subtitle, const char *body, const char *categoryId, const char *actions_json);
void registerNotificationCategory(int channelID, const char *categoryId, const char *actions_json, bool hasReplyField, const char *replyPlaceholder, const char *replyButtonTitle);
void removeNotificationCategory(int channelID, const char *categoryId);
void removeAllPendingNotifications(void);
void removePendingNotification(const char *identifier);
void removeAllDeliveredNotifications(void);
void removeDeliveredNotification(const char *identifier);

#ifdef __cplusplus
}
#endif