#import "cocoa_helpers.h"
#import <Foundation/Foundation.h>

// 释放继承链内存
void freeInheritanceChain(InheritanceChain* chain) {
    if (chain && chain->classNames) {
        for (int i = 0; i < chain->count; i++) {
            free((void*)chain->classNames[i]);
        }
        free(chain->classNames);
        chain->classNames = NULL;
        chain->count = 0;
    }
}

// 获取对象的完整继承链（包括自身）
InheritanceChain getObjectInheritanceChain(void* objHandle) {
    InheritanceChain chain = {NULL, 0};
    if (!objHandle) return chain;

    // 从对象获取类
    Class currentClass = object_getClass((NSObject*)objHandle);
    if (!currentClass) return chain;

    // 遍历继承链
    while (currentClass) {
        // 获取当前类名
        const char* className = class_getName(currentClass);
        if (className) {
            // 复制类名（避免指针失效）
            char* copiedName = strdup(className);
            if (copiedName) {
                // 扩展数组容量
                const char** newClassNames = realloc(chain.classNames, sizeof(const char*) * (chain.count + 1));
                if (newClassNames) {
                    chain.classNames = newClassNames;
                    chain.classNames[chain.count] = copiedName;
                    chain.count++;
                } else {
                    free(copiedName);
                }
            }
        }

        // 获取父类，继续遍历
        currentClass = class_getSuperclass(currentClass);
    }

    return chain;
}

// 获取对象的类名
const char* getObjectClassName(void* objHandle) {
    if (!objHandle) return "nil";
    NSObject* obj = (NSObject*)objHandle;
    return [obj.className UTF8String];
}

// 检查对象是否为指定类（或其子类）的实例
BOOL isObjectOfClass(void* objHandle, const char* className) {
    if (!objHandle || !className) return NO;
    NSObject* obj = (NSObject*)objHandle;
    Class targetClass = NSClassFromString([NSString stringWithUTF8String:className]);
    return [obj isKindOfClass:targetClass];
}
