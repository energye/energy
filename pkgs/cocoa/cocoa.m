#import "cocoa.h"
#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>

#pragma mark - UI 线程执行函数

// UI 线程执行
void RegisterRunOnMainThreadCallback(RunOnMainThreadCallback callback) {
    _runOnMainThreadCallback = callback;
}

// 在主线程执行回调的桥接函数
void ExecuteRunOnMainThread(long id) {
    if ([NSThread isMainThread]) {
        _runOnMainThreadCallback(id);
    } else {
        dispatch_async(dispatch_get_main_queue(), ^{
            _runOnMainThreadCallback(id);
        });
    }
}