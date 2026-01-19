#import "config.h"
#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>

// 创建默认控件样式
ControlProperty CreateDefaultControlProperty() {
    ControlProperty property;
    property.width = 0; // 0表示自动大小
    property.height = 0;
    property.minWidth = 0;
    property.maxWidth = 0;
    property.bezelStyle = NSBezelStyleTexturedRounded;
    property.controlSize = NSControlSizeRegular;
    property.font = nil;
    property.VisibilityPriority = NSToolbarItemVisibilityPriorityStandard;
    return property;
}

// 创建自定义控件样式
ControlProperty CreateControlProperty(CGFloat width, CGFloat height, NSBezelStyle bezelStyle, NSControlSize controlSize, void *font) {
    ControlProperty property;
    property.width = width;
    property.height = height;
    property.bezelStyle = bezelStyle;
    property.controlSize = controlSize;
    property.font = (__bridge NSFont *)font;
    return property;
}

// 通用函数：通过NSControl设置控件属性（适用于按钮、文本框等）
void ConfigureControl(NSControl *control, NSString *tooltipStr, ControlProperty property) {
    control.controlSize = property.controlSize;
    if (tooltipStr) {
        control.toolTip = tooltipStr;
    }
    if (property.font) {
        control.font = property.font;
    }
    if (property.width > 0) {
        [control.widthAnchor constraintEqualToConstant:property.width].active = YES;
    }
    if (property.height > 0) {
        [control.heightAnchor constraintEqualToConstant:property.height].active = YES;
    }
    // 最小和最大宽度约束
    if (property.minWidth > 0) {
        NSLayoutConstraint *minWidthConstraint = [control.widthAnchor constraintGreaterThanOrEqualToConstant:property.minWidth];
        minWidthConstraint.priority = NSLayoutPriorityDefaultHigh;
        minWidthConstraint.active = YES;
    }
    if (property.maxWidth > 0) {
        NSLayoutConstraint *maxWidthConstraint = [control.widthAnchor constraintLessThanOrEqualToConstant:property.maxWidth];
        maxWidthConstraint.priority = NSLayoutPriorityDefaultHigh;
        maxWidthConstraint.active = YES;
    }
}

void SetBindControlObjectIdentifier(void* nsObject, const char *identifier) {
    if (!nsObject || !identifier) {
        NSLog(@"[ERROR] BindControlObjectIdentifier 必要参数为空");
        return;
    }
    NSObject *object = (__bridge NSControl *)nsObject;
    NSString *idStr = [NSString stringWithUTF8String:identifier];
    objc_setAssociatedObject(object, @"identifier", idStr, OBJC_ASSOCIATION_RETAIN);
}

NSString *GetBindControlObjectIdentifier(NSObject* sender) {
    if (!sender) {
        NSLog(@"[ERROR] GetBindControlObjectIdentifier 必要参数为空");
        return nil;
    }
    NSString *identifier = objc_getAssociatedObject(sender, @"identifier");
    return identifier;
}

// 设置控件启用状态
void SetControlEnable(void* control, BOOL enable) {
    if (!control) {
        NSLog(@"[ERROR] SetControlEnable: 控件指针为空");
        return;
    }
    NSControl *nsControl = (__bridge NSControl *)control;
    dispatch_async(dispatch_get_main_queue(), ^{
        [nsControl setEnabled:enable];
    });
}

// 获取控件启用状态
BOOL GetControlEnable(void* control) {
    if (!control) {
        NSLog(@"[ERROR] GetControlEnable: 控件指针为空");
        return NO;
    }
    NSControl *nsControl = (__bridge NSControl *)control;
    return [nsControl isEnabled];
}

// 设置控件显示/隐藏状态
void SetControlHidden(void* control, BOOL hidden) {
    if (!control) {
        NSLog(@"[ERROR] SetControlHidden: 控件指针为空");
        return;
    }

    NSControl *nsControl = (__bridge NSControl *)control;
    dispatch_async(dispatch_get_main_queue(), ^{
        [nsControl setHidden:hidden];
    });
}

// 获取控件显示/隐藏状态
BOOL GetControlHidden(void* control) {
    if (!control) {
        NSLog(@"[ERROR] GetControlHidden: 控件指针为空");
        return YES; // 默认返回隐藏状态
    }

    NSControl *nsControl = (__bridge NSControl *)control;
    return [nsControl isHidden];
}

// 设置控件透明度（可选扩展）
void SetControlAlphaValue(void* control, CGFloat alpha) {
    if (!control) {
        NSLog(@"[ERROR] SetControlAlphaValue: 控件指针为空");
        return;
    }

    NSControl *nsControl = (__bridge NSControl *)control;
    dispatch_async(dispatch_get_main_queue(), ^{
        [nsControl setAlphaValue:alpha];
    });
}

// 获取控件透明度（可选扩展）
CGFloat GetControlAlphaValue(void* control) {
    if (!control) {
        NSLog(@"[ERROR] GetControlAlphaValue: 控件指针为空");
        return 0.0;
    }

    NSControl *nsControl = (__bridge NSControl *)control;
    return [nsControl alphaValue];
}

// 设置控件焦点
BOOL SetControlFocus(void* control, BOOL focus) {
    if (!control) {
        NSLog(@"[ERROR] SetControlFocus: 控件指针为空");
        return false;
    }
    NSControl *nsControl = (__bridge NSControl *)control;
    if(focus) {
        BOOL success = [nsControl becomeFirstResponder];
        if (success) {
//            NSLog(@"成功获取焦点");
        }
        return success;
    }else{
        BOOL success = [nsControl resignFirstResponder];
        if (success) {
//            NSLog(@"成功失去焦点");
        }
        return success;
    }
    return false;
}