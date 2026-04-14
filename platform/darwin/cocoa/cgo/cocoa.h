#import <Cocoa/Cocoa.h>
#import <go_arguments.h>

#ifdef __cplusplus
extern "C" {
#endif

// 窗口事件唯一标识
typedef enum {
    TWindowEventEnterFullScreen = 10000,
    TWindowEventExitFullScreen = 10001,
    TWindowEventWillUseFullScreenPresentationOptions = 10002,
    TWindowEventDidResize = 10003
} TWindowEvent;

// 颜色
typedef struct {
    CGFloat Red;
    CGFloat Green;
    CGFloat Blue;
    CGFloat Alpha;
}  Color;

typedef struct {
    long    type_;              // 事件类型, 用于区分普通通知事件, 还是特殊事件 1: 点击事件 2: 文本改变事件 3:文本提交事件 4:下拉框回车/离开焦点事件 5:下拉框选择事件
    const   char *identifier;   // 控件标识
    const   char *value;        // 控件值
    long    index;              // 值索引
    void    *owner;             // 控件所属对象
    void    *sender;            // 控件
    GoArguments *arguments;
} TCallbackContext;

typedef GoArguments* (*TEventCallback)(TCallbackContext *context);
TCallbackContext* CreateCallbackContext(const NSString* identifier, const NSString* value, long index, void* owner, void* sender);
void FreeCallbackContext(TCallbackContext* context);

// 工具栏配置选项
typedef struct {
    int             ShowSeparator;  // bool
} ToolbarConfiguration;

// 主线程回调函数-测试或使用
typedef void (*RunOnMainThreadCallback)(long id);

void RegisterRunOnMainThreadCallback(RunOnMainThreadCallback callback);
void ExecuteRunOnMainThread(long id);

// 保存全局回调
static RunOnMainThreadCallback _runOnMainThreadCallback = nil;

#ifdef __cplusplus
}
#endif