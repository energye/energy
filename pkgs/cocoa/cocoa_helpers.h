#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>
#include <stdlib.h>
#include <string.h>

// 防止 C++ 名称修饰
#ifdef __cplusplus
extern "C" {
#endif

// 存储类名列表的结构体
typedef struct {
    const char** classNames;
    int count;
} InheritanceChain;

// 释放继承链内存
void freeInheritanceChain(InheritanceChain* chain);

// 获取对象的完整继承链（包括自身）
InheritanceChain getObjectInheritanceChain(void* objHandle);

// 获取对象的类名
const char* getObjectClassName(void* objHandle);

// 检查对象是否为指定类（或其子类）的实例
BOOL isObjectOfClass(void* objHandle, const char* className);

#ifdef __cplusplus
}
#endif
