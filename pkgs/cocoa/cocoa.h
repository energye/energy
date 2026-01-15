#import <Cocoa/Cocoa.h>

#ifdef __cplusplus
extern "C" {
#endif

// 主线程回调函数-测试或使用
typedef void (*RunOnMainThreadCallback)(long id);

void RegisterRunOnMainThreadCallback(RunOnMainThreadCallback callback);
void ExecuteRunOnMainThread(long id);

// 保存全局回调
static RunOnMainThreadCallback _runOnMainThreadCallback = nil;

#ifdef __cplusplus
}
#endif