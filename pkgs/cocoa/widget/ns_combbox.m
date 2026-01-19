#import "config.h"
#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>

void* NewCombobox(void* nsDelegate, const char *tooltip, const char **items, int count, ControlProperty property) {
    if (!nsDelegate) {
        NSLog(@"[ERROR] NewTextField 必要参数为空");
        return nil;
    }
    MainToolbarDelegate *delegate = (MainToolbarDelegate*)nsDelegate;
    NSString *tooltipStr = tooltip ? [NSString stringWithUTF8String:tooltip] : nil;
    NSComboBox *comboBox = [[NSComboBox alloc] init];
    comboBox.delegate = delegate;
    [comboBox setEditable:NO];
    // 添加选项
    for (int i = 0; i < count; i++) {
        [comboBox addItemWithObjectValue:[NSString stringWithUTF8String:items[i]]];
    }
    // 设置默认选择
    if (count > 0) {
        [comboBox selectItemAtIndex:0];
    }
    ConfigureControl(comboBox, tooltipStr, property);
    return (__bridge void*)(comboBox);
}

void ComboboxAddItem(void* nsCombobox, const char **items, int count) {

}

void ComboboxRemoveItem(void* nsCombobox, const char *item) {

}

void ComboboxRemoveItemIndex(void* nsCombobox, int index) {

}