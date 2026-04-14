#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>
#import "cocoa.h"

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

// 创建回调上下文
TCallbackContext* CreateCallbackContext(const NSString* identifier, const NSString* value, long index, void* owner, void* sender) {
    // 分配内存空间
    TCallbackContext* context = (TCallbackContext*)malloc(sizeof(TCallbackContext));
    if (!context) return NULL;  // 内存分配失败
    context->index = index;
    context->owner = owner;
    context->sender = sender;
    context->identifier = identifier ? strdup([identifier UTF8String]) : "";
    context->value = value ? strdup([value UTF8String]) : "";
    context->arguments = NULL;
    return context;
}

// 释放工具栏事件回调上下文
void FreeCallbackContext(TCallbackContext* context) {
    if (!context) return;
    // 释放字符串内存
    free((void*)context->identifier);
    free((void*)context->value);
    if(context->arguments){
        FreeGoArguments(context->arguments);
    }
    // 释放结构体
    free(context);
}