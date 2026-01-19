#import <Cocoa/Cocoa.h>

#ifdef __cplusplus
extern "C" {
#endif

// 事件类型, 用于区分普通通知事件, 还是特殊事件
typedef enum {
    TCCNotify = 0,
    TCCClicked = 1,
    TCCTextDidChange = 2,
    TCCTextDidEndEditing = 3,
    TCCSelectionChanged = 4,
    TCCSelectionDidChange = 5
} TccType;

// 颜色
typedef struct {
    CGFloat Red;
    CGFloat Green;
    CGFloat Blue;
    CGFloat Alpha;
}  Color;

// 工具栏配置选项
typedef struct {
    int             IsAllowsUserCustomization;  // bool
    int             IsAutoSavesConfiguration;   // bool
	int             Transparent;                // bool
	int             ShowsToolbarButton;         // bool
	NSUInteger      SeparatorStyle;
    NSUInteger      DisplayMode;
    NSUInteger      SizeMode;
    NSUInteger      Style;
} ToolbarConfiguration;

// 控件样式结构体
typedef struct {
    CGFloat         width;
    CGFloat         height;
    CGFloat         minWidth;
    CGFloat         maxWidth;
    NSBezelStyle    bezelStyle;
    NSControlSize   controlSize;
    NSFont          *font;
    int             IsNavigational;      // bool
    int             IsCenteredItem;      // bool
    NSInteger       VisibilityPriority;
} ControlProperty;

// 主线程回调函数-测试或使用
typedef void (*RunOnMainThreadCallback)(long id);

void RegisterRunOnMainThreadCallback(RunOnMainThreadCallback callback);
void ExecuteRunOnMainThread(long id);

// 保存全局回调
static RunOnMainThreadCallback _runOnMainThreadCallback = nil;

#ifdef __cplusplus
}
#endif