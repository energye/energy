#import <Cocoa/Cocoa.h>
#import <Foundation/Foundation.h>
#import <stdarg.h>
#import <stdbool.h>
#import <stdlib.h>
#import <objc/runtime.h>
#import <objc/message.h>
#import <Foundation/Foundation.h>

#ifdef __cplusplus
extern "C" {
#endif

// 简化的数据类型枚举
typedef enum {
    ArgsType_None,
    ArgsType_Int,
    ArgsType_Float,
    ArgsType_Bool,
    ArgsType_String,
    ArgsType_Struct,
    ArgsType_Object,
    ArgsType_Pointer
} GoArgumentsType;

// 简化的数据项结构
typedef struct {
    void* Value;
    GoArgumentsType Type;
} GoArgsItem;

// 主数据结构
typedef struct {
    GoArgsItem* Items;
    int Count;
} GoArguments;

// 销毁函数
void FreeGoArguments(GoArguments* data);
void GoFreeGoArguments(GoArguments *data);

// 创建函数
GoArguments* CreateGoArguments(int count, ...);

// 从 GoArguments 获取数据的函数
void* GetFromGoArguments(GoArguments* data, int index, GoArgumentsType expectedType);

// 获取参数项
GoArgsItem* GetItemFromGoArguments(GoArguments* data, int index);
// 获取参数的 int 值
int GetIntFromGoArguments(GoArguments* data, int index);
// 获取参数的 double 值
double GetFloatFromGoArguments(GoArguments* data, int index);
// 获取参数的 bool 值
bool GetBoolFromGoArguments(GoArguments* data, int index);
// 获取参数的 string 值
const char* GetStringFromGoArguments(GoArguments* data, int index);
// 获取参数的 NSString 值
NSString* GetNSStringFromGoArguments(GoArguments* data, int index);
// 获取参数的 struct 值
void* GetStructFromGoArguments(GoArguments* data, int index);
// 获取参数的 object 值
void* GetObjectFromGoArguments(GoArguments* data, int index);
// 获取参数的 pointer 值
void* GetPointerFromGoArguments(GoArguments* data, int index);

#ifdef __cplusplus
}
#endif