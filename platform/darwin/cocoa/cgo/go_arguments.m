#import "go_arguments.h"

// 释放 GoArguments 及其所有元素
void FreeGoArguments(GoArguments* data) {
    if (!data) return;
    //NSLog(@"FreeGoArguments 释放个数: %d", data->Count);
    // 释放所有数据项
    for (int i = 0; i < data->Count; i++) {
        @try {
            GoArgsItem item = data->Items[i];
            //NSLog(@"FreeGoArguments: 释放第%d个参数, 类型: %d", i, item.Type);
            // 根据类型释放内存
            switch (item.Type) {
                case ArgsType_Int:
                case ArgsType_Float:
                case ArgsType_Bool:
                    if (item.Value) {
                        free(item.Value);
                        //NSLog(@"FreeGoArguments: 释放基本类型 - 索引=%d，Type=%d，Value=%p", i, item.Type, item.Value);
                        item.Value = NULL;
                    }
                    break;
                case ArgsType_String:
                    //NSLog(@"FreeGoArguments 释放字符串");
                    if (item.Value) {
                        free(item.Value);
                        //NSLog(@"FreeGoArguments: 释放字符串 - 索引=%d，Type=%d，Value=%p", i, item.Type, item.Value);
                        item.Value = NULL;
                    }
                    break;
                case ArgsType_Struct:
                    if (item.Value) {
                        free(item.Value);
                        item.Value = NULL;
                    }
                    break;
                case ArgsType_Object: { // 只在OC创建
                    id obj = (id)item.Value;
                    if (obj) {
                        [obj release];
                        //NSLog(@"FreeGoArguments: 释放对象 - 索引=%d，Type=%d，对象=%@（%p）", i, item.Type, obj, obj);
                        item.Value = NULL;
                    }
                    break;
                }
                case ArgsType_Pointer: // 只在OC创建
                    //NSLog(@"FreeGoArguments: 跳过指针释放 - 索引=%d，Type=%d，Value=%p", i, item.Type, item.Value);
                    item.Value = NULL;
                    break;
                default:
                    //NSLog(@"FreeGoArguments: 未处理类型 - 索引=%d，Type=%d，Value=%p", i, item.Type, item.Value);
                    item.Value = NULL;
                    break;
            }
        }
        @catch (NSException *e) {
            NSLog(@"FreeGoArguments: 释放第%d个参数时发生异常：%@，原因：%@", i, e.name, e.reason);
        }
    }
    // 释放数组本身
    if (data->Items) {
        free(data->Items);
        //NSLog(@"FreeGoArguments: 释放参数数组内存，地址=%p", data->Items);
    }
    // 释放 GoArguments 结构
    free(data);
    //NSLog(@"FreeGoArguments: 释放GoArguments结构体，地址=%p", data);
}

// 通用添加函数
/*
传递4个参数：整数、字符串、布尔值、浮点数
GoArguments* args = CreateGoArguments(
    5,                  // 参数数量为4
    @(123),             // 整数（int）
    @"mixed types",     // 字符串（NSString）
    @(NO),              // 布尔值（BOOL）
    @(3.14159)          // 浮点数（double
    [NSValue valueWithPointer:buffer]; // Pointer
);
*/
GoArguments* CreateGoArguments(int count, ...) {
    GoArguments* data = malloc(sizeof(GoArguments));
    data->Count = count;
    data->Items = malloc(sizeof(GoArgsItem) * data->Count);

    va_list args;
    va_start(args, count); // 正确初始化：第二个参数是可变参数前的最后一个命名参数（count）

    // 按参数数量count遍历，确保获取所有参数
    for (int i = 0; i < count; i++) {
        id arg = va_arg(args, id); // 逐个获取参数
        GoArgsItem item;
        // 自动类型推断
        if ([arg isKindOfClass:[NSNumber class]]) {
            NSNumber* number = (NSNumber*)arg;
            const char* objCType = [number objCType];
            if (strcmp(objCType, @encode(BOOL)) == 0 ||
                     strcmp(objCType, @encode(bool)) == 0 ||
                     strcmp(objCType, @encode(char)) == 0 ) {
                // 布尔类型
                bool* value = malloc(sizeof(bool));
                *value = [number boolValue];
                item.Value = value;
                item.Type = ArgsType_Bool;
            }
            else if (strcmp(objCType, @encode(int)) == 0 ||
                strcmp(objCType, @encode(long)) == 0 ||
                strcmp(objCType, @encode(NSInteger)) == 0) {
                // 整数类型
                int* value = malloc(sizeof(int));
                *value = [number intValue];
                item.Value = value;
                item.Type = ArgsType_Int;
            }
            else if (strcmp(objCType, @encode(float)) == 0 ||
                    strcmp(objCType, @encode(double)) == 0) {
                // 浮点数类型
                double* value = malloc(sizeof(double));
                *value = [number doubleValue];
                item.Value = value;
                item.Type = ArgsType_Float;
            }
            else {
                // 无法识别的NSNumber，当作对象处理
                item.Value = (void*)[arg retain];
                item.Type = ArgsType_Object;
            }
        }
        else if ([arg isKindOfClass:[NSString class]]) {
            // 字符串类型
            NSString* string = (NSString*)arg;
            char* value = strdup([string UTF8String]);
            item.Value = value;
            item.Type = ArgsType_String;
        }
        else if ([arg isKindOfClass:[NSValue class]] &&
                 strcmp([(NSValue*)arg objCType], @encode(void*)) == 0) {
            // 指针类型（包装在 NSValue 中）
            void* value;
            [(NSValue*)arg getValue:&value];
            item.Value = value;
            item.Type = ArgsType_Pointer;
        }
        else { // 对象类型
            item.Value = (void*)[arg retain];
//            item.Value = (__bridge void*)(arg);
            item.Type = ArgsType_Object;
        }
        // 添加到数组
        data->Items[i] = item;
    }
    va_end(args);
    return data;
}

// 从 GoArguments 获取 Item 函数
GoArgsItem* GetItemFromGoArguments(GoArguments* data, int index) {
    if (!data || index < 0 || index >= data->Count) return NULL;
    return &data->Items[index];
}

// 从 GoArguments 获取数据的通用函数
void* GetFromGoArguments(GoArguments* data, int index, GoArgumentsType expectedType) {
    if (!data || index < 0 || index >= data->Count) return NULL;
    GoArgsItem item = data->Items[index];
    if (item.Type != expectedType) return NULL;
    return item.Value;
}

int GetIntFromGoArguments(GoArguments* data, int index) {
    int* value = (int*)GetFromGoArguments(data, index, ArgsType_Int);
    return value ? *value : 0;
}

double GetFloatFromGoArguments(GoArguments* data, int index) {
    double* value = (double*)GetFromGoArguments(data, index, ArgsType_Float);
    return value ? *value : 0.0f;
}

bool GetBoolFromGoArguments(GoArguments* data, int index) {
    bool* value = (bool*)GetFromGoArguments(data, index, ArgsType_Bool);
    return value ? *value : false;
}

const char* GetStringFromGoArguments(GoArguments* data, int index) {
    return (const char*)GetFromGoArguments(data, index, ArgsType_String);
}

NSString* GetNSStringFromGoArguments(GoArguments* data, int index) {
    const char* cStr = GetStringFromGoArguments(data, index);
    if (!cStr) {
        return nil;
    }
    NSString* ocStr = [NSString stringWithUTF8String:cStr];
    // free((void*)cStr); // 不在这里释放，统一释放
    return ocStr;
}

void* GetStructFromGoArguments(GoArguments* data, int index) {
    return GetFromGoArguments(data, index, ArgsType_Struct);
}

void* GetObjectFromGoArguments(GoArguments* data, int index) {
    return GetFromGoArguments(data, index, ArgsType_Object);
}

void* GetPointerFromGoArguments(GoArguments* data, int index) {
    return GetFromGoArguments(data, index, ArgsType_Pointer);
}