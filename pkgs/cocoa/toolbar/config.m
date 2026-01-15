#import "config.h"
#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>

BOOL isARCMode() {
#if __has_feature(objc_arc)
    return YES;  // 支持ARC
#else
    return NO;   // 不支持ARC
#endif
}

// 获取字符串常量的值, 例如: C.NSToolbarSpaceItemIdentifier
const char* GetStringConstValue(const void* nsStringConst) {
    NSString* identifier = (__bridge NSString*)nsStringConst;
    return [identifier UTF8String];
}

// 创建工具栏事件回调上下文
ToolbarCallbackContext* CreateToolbarCallbackContext(const NSString* identifier, const NSString* value, long index, void* owner, void* sender) {
    // 分配内存空间
    ToolbarCallbackContext* context = (ToolbarCallbackContext*)malloc(sizeof(ToolbarCallbackContext));
    if (!context) return NULL;  // 内存分配失败
    context->type_ = TCCNotify;
    context->index = index;
    context->owner = owner;
    context->sender = sender;
    context->identifier = identifier ? strdup([identifier UTF8String]) : "";
    context->value = value ? strdup([value UTF8String]) : "";
    context->arguments = NULL;
    return context;
}

// 释放工具栏事件回调上下文
void FreeToolbarCallbackContext(ToolbarCallbackContext* context) {
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

// 工具栏代理 key 唯一
static char kToolbarDelegateKey;

@implementation MainToolbarDelegate

- (instancetype)init {
    self = [super init];
    if (self) {
        self.controls = [NSMutableDictionary dictionary];
        self.dynamicIdentifiers = [NSMutableArray array];
        self.controlProperty = [NSMutableDictionary dictionary];
        _callback = NULL;
        // 监听窗口大小变化
        [[NSNotificationCenter defaultCenter] addObserver:self selector:@selector(windowDidResize:) name:NSWindowDidResizeNotification object:nil];
    }
    return self;
}

- (void)dealloc {
    NSLog(@"MainToolbarDelegate dealloc 释放");
    [[NSNotificationCenter defaultCenter] removeObserver:self];
    [super dealloc];
}

// 窗口大小监听
- (void)windowDidResize:(NSNotification *)notification {
    NSWindow *window = notification.object;
    // NSLog(@"windowDidResize");
    ToolbarCallbackContext *context = CreateToolbarCallbackContext(@"__doWindowResize", @"", -1, nil, _window);
    GoArguments *result;
    @try{
        result = _callback(context);
    } @finally {
        if(result){
            FreeGoArguments(result);
        }
        FreeToolbarCallbackContext(context);
    }
}

- (void)addControl:(NSView *)control forIdentifier:(NSString *)identifier withProperty:(ControlProperty)property {
    NSLog(@"addControl");
    _controls[identifier] = control;
    // 存储控件样式
    NSValue *propertyValue = [NSValue value:&property withObjCType:@encode(ControlProperty)];
    _controlProperty[identifier] = propertyValue;
    // 存储控件 id
    if (![_dynamicIdentifiers containsObject:identifier]) {
        [_dynamicIdentifiers addObject:identifier];
    }
}

- (NSView *)controlForIdentifier:(NSString *)identifier {
    NSLog(@"controlForIdentifier");
    return _controls[identifier];
}

- (void)removeControlForIdentifier:(NSString *)identifier {
    NSLog(@"removeControlForIdentifier");
    // 从控件字典中移除
    [_controls removeObjectForKey:identifier];
    // 从样式字典中移除
    [_controlProperty removeObjectForKey:identifier];
    // 从标识符数组中移除
    [_dynamicIdentifiers removeObject:identifier];
}

- (void)setCallback:(ControlEventCallback)callback withWindow:(NSWindow *)window withToolbar:(NSToolbar *)toolbar {
    _callback = callback;
    _window = window;
    _toolbar = toolbar;
}

- (void)updateControlProperty:(NSString *)identifier withProperty:(ControlProperty)property {
    NSLog(@"updateControlProperty");
    NSView *control = [self controlForIdentifier:identifier];
    if (!control) return;

    // 更新存储的样式
    NSValue *propertyValue = [NSValue value:&property withObjCType:@encode(ControlProperty)];
    _controlProperty[identifier] = propertyValue;

    // 应用样式到控件
    if ([control isKindOfClass:[NSControl class]]) {
        NSControl *ctrl = (NSControl *)control;
        ctrl.controlSize = property.controlSize;

        // 宽度约束
        if (property.width > 0) {
            // 移除现有宽度约束
            for (NSLayoutConstraint *constraint in control.constraints) {
                if (constraint.firstAttribute == NSLayoutAttributeWidth) {
                    [control removeConstraint:constraint];
                    break;
                }
            }
            // 添加新宽度约束
            [control.widthAnchor constraintEqualToConstant:property.width].active = YES;
        }

        // 高度约束
        if (property.height > 0) {
            // 移除现有高度约束
            for (NSLayoutConstraint *constraint in control.constraints) {
                if (constraint.firstAttribute == NSLayoutAttributeHeight) {
                    [control removeConstraint:constraint];
                    break;
                }
            }
            // 添加新高度约束
            [control.heightAnchor constraintEqualToConstant:property.height].active = YES;
        }

        // 特定控件类型的样式
        if ([control isKindOfClass:[NSButton class]]) {
            NSButton *button = (NSButton *)control;
            button.bezelStyle = property.bezelStyle;
            if (property.font) {
                button.font = property.font;
            }
        } else if ([control isKindOfClass:[NSTextField class]] ||
                   [control isKindOfClass:[NSSearchField class]] ||
                   [control isKindOfClass:[NSComboBox class]]) {
            NSTextField *textField = (NSTextField *)control;
            if (property.font) {
                textField.font = property.font;
            }
        }
    }
}

#pragma mark - Toolbar Delegate

- (NSArray<NSString *> *)toolbarDefaultItemIdentifiers:(NSToolbar *)toolbar {
    NSLog(@"toolbarDefaultItemIdentifiers");
    ToolbarCallbackContext *context = CreateToolbarCallbackContext(@"__doToolbarDefaultItemIdentifiers", @"", -1, _window, _toolbar);
    GoArguments *result;
    NSMutableArray *identifiers = [NSMutableArray array];
    @try{
        result = _callback(context);
        if(result){
            for (int i = 0; i < result->Count; i++) {
                NSString *id = GetNSStringFromGoArguments(result, i);
                [identifiers addObject:id];
            }
        }
    } @finally {
        if(result){
           FreeGoArguments(result);
        }
        FreeToolbarCallbackContext(context);
    }
    return identifiers;
}

- (NSArray<NSString *> *)toolbarAllowedItemIdentifiers:(NSToolbar *)toolbar {
    NSLog(@"toolbarAllowedItemIdentifiers");
    ToolbarCallbackContext *context = CreateToolbarCallbackContext(@"__doToolbarAllowedItemIdentifiers", @"", -1, _window, _toolbar);
    GoArguments *result;
    NSMutableArray *identifiers = [NSMutableArray array];
    @try{
        result = _callback(context);
        if(result){
            for (int i = 0; i < result->Count; i++) {
                NSString *id = GetNSStringFromGoArguments(result, i);
                [identifiers addObject:id];
            }
        }
    } @finally {
        if(result){
           FreeGoArguments(result);
        }
        FreeToolbarCallbackContext(context);
    }
    return identifiers;
}

//- (BOOL)toolbar:(NSToolbar *)toolbar canHideItem:(NSToolbarItem *)item {
//    NSLog(@"toolbar-canHideItem");
//    // 对搜索框项目返回 NO，禁止隐藏
//    if ([item.itemIdentifier isEqualToString:@"SiteAddrSearch"]) {
//        return NO;
//    }
//    return YES; // 其他项目允许隐藏
//}

- (NSToolbarItem *)toolbar:(NSToolbar *)toolbar
     itemForItemIdentifier:(NSToolbarItemIdentifier)itemIdentifier
 willBeInsertedIntoToolbar:(BOOL)flag {
    NSLog(@"toolbarItemIdentifier: %@", itemIdentifier);
    // 处理系统项
    if ([itemIdentifier isEqualToString:NSToolbarFlexibleSpaceItemIdentifier]) {
        return [[NSToolbarItem alloc] initWithItemIdentifier:NSToolbarFlexibleSpaceItemIdentifier];
    }
    if ([itemIdentifier isEqualToString:NSToolbarSpaceItemIdentifier]) {
        return [[NSToolbarItem alloc] initWithItemIdentifier:NSToolbarSpaceItemIdentifier];
    }

    ToolbarCallbackContext *context = CreateToolbarCallbackContext(@"__doDelegateToolbar", @"", -1, _window, _toolbar);
    GoArguments *result;
    @try{
        context->arguments = CreateGoArguments(1, itemIdentifier);
        result = _callback(context);
        if(result){
            NSView *control = (NSView *)GetObjectFromGoArguments(result, 0);
            if (control) {
                NSLog(@"doDelegateToolbar control 获取成功");
                if ([control isKindOfClass:[NSControl class]]) { // 控件类
                    ControlProperty *property = (ControlProperty *)GetStructFromGoArguments(result, 1);
                    NSLog(@"doDelegateToolbar control %d %d", property->IsNavigational, property->IsCenteredItem);
                    NSToolbarItem *item = [[NSToolbarItem alloc] initWithItemIdentifier:itemIdentifier];
                    item.view = control;
                    item.navigational = property->IsNavigational; // 导航模式 靠左
                    if (property->IsCenteredItem) {
                        toolbar.centeredItemIdentifier = item.itemIdentifier;  // 设置为居中项
                    }
                    item.visibilityPriority = property->VisibilityPriority; // 可见优先1级
//                    [self updateControlProperty:itemIdentifier withProperty:property];
                    //ConfigureControl((NSControl *)control, itemIdentifier, *property);
                    return item;
                } else if ([control isKindOfClass:[NSView class]]) { // 视图类
                    NSLog(@"doDelegateToolbar view");
                    NSToolbarItem *item = [[NSToolbarItem alloc] initWithItemIdentifier:itemIdentifier];
                    item.minSize = NSMakeSize(200, 38);
                    item.maxSize = NSMakeSize(200, 38);
                    item.view = control;
                    return item;
                } else {
                     NSLog(@"doDelegateToolbar control 类型判断失败");
                }
            } else {
                NSLog(@"doDelegateToolbar control 获取失败");
            }
        } else {
             NSLog(@"doDelegateToolbar 返回值为空");
         }
    } @finally {
        if(result){
           FreeGoArguments(result);
        }
        FreeToolbarCallbackContext(context);
    }
    return nil;
}

#pragma mark - 事件处理

- (void)searchFieldDidStartSearching:(NSSearchField *)sender {
    NSLog(@"搜索开始: %@", sender.stringValue);
    // 在这里处理搜索开始时的逻辑
}

- (void)searchFieldDidEndSearching:(NSSearchField *)sender {
    NSLog(@"搜索结束");
    // 在这里处理搜索结束时的逻辑
}

- (void)buttonClicked:(NSButton *)sender {
    NSLog(@"buttonClicked");
    if (_callback) {
        NSString *identifier = GetBindControlObjectIdentifier(sender);
        if (identifier) {
            ToolbarCallbackContext *context = CreateToolbarCallbackContext(identifier, @"", -1, _window, sender);
            GoArguments *result;
            @try{
                context->type_ = TCCNotify;
                result = _callback(context);
            } @finally {
                if(result){
                    FreeGoArguments(result);
                }
                FreeToolbarCallbackContext(context);
            }
        }
    }
}

- (void)comboBoxSelectionChanged:(NSComboBox *)sender {
    NSLog(@"comboBoxSelectionChanged");
    if (_callback) {
        NSString *identifier = GetBindControlObjectIdentifier(sender);
        if (identifier) {
            NSInteger selectedIndex = [sender indexOfSelectedItem];
            ToolbarCallbackContext *context = CreateToolbarCallbackContext(identifier, [sender stringValue], selectedIndex, _window, sender);
            GoArguments *result;
            @try{
                context->type_ = TCCSelectionChanged;
                result = _callback(context);
            } @finally {
                if(result){
                    FreeGoArguments(result);
                }
                FreeToolbarCallbackContext(context);
            }
        }
    }
}

// 用户选择发生变化时触发
- (void)comboBoxSelectionDidChange:(NSNotification *)notification {
    NSLog(@"comboBoxSelectionChanged");
    if (_callback) {
        id control = notification.object;
        NSString *identifier = GetBindControlObjectIdentifier(control);
        if (identifier) {
            NSInteger selectedIndex = [control indexOfSelectedItem];
            ToolbarCallbackContext *context = CreateToolbarCallbackContext(identifier, [control stringValue], selectedIndex, _window, control);
            GoArguments *result;
            @try{
                context->type_ = TCCSelectionDidChange;
                result = _callback(context);
            } @finally {
                if(result){
                    FreeGoArguments(result);
                }
                FreeToolbarCallbackContext(context);
            }
        }
    }
}

- (void)controlTextDidChange:(NSNotification *)notification {
    NSLog(@"controlTextDidChange");
    if (_callback) {
        id control = notification.object;
        NSString *identifier = GetBindControlObjectIdentifier(control);
        if (identifier) {
            ToolbarCallbackContext *context = CreateToolbarCallbackContext(identifier, [control stringValue], -1, _window, control);
            GoArguments *result;
            @try{
                context->type_ = TCCTextDidChange;
                result = _callback(context);
            } @finally {
                if(result){
                    FreeGoArguments(result);
                }
                FreeToolbarCallbackContext(context);
            }
        }
    }
}

- (void)controlTextDidEndEditing:(NSNotification *)notification {
    NSLog(@"controlTextDidEndEditing");
    if (_callback) {
        id control = notification.object;
        NSString *identifier = GetBindControlObjectIdentifier(control);
        if (identifier) {
            ToolbarCallbackContext *context = CreateToolbarCallbackContext(identifier, [control stringValue], -1, _window, control);
            GoArguments *result;
            @try{
                context->type_ = TCCTextDidEndEditing;
                result = _callback(context);
            } @finally {
                if(result){
                    FreeGoArguments(result);
                }
                FreeToolbarCallbackContext(context);
            }
        }
    }
}

@end

#pragma mark - 公共函数实现


// 初始化函数
__attribute__((constructor))
static void initializeDelegateMap() {
    // NSLog(@"initializeDelegateMap");
}

// 设置窗口背景色
void SetWindowBackgroundColor(unsigned long nsWindowHandle, Color color) {
    NSWindow *window = (__bridge NSWindow *)(void *)nsWindowHandle;
    NSColor *bgColor = [NSColor colorWithCalibratedRed:color.Red
                                                 green:color.Green
                                                  blue:color.Blue
                                                 alpha:color.Alpha];
    window.backgroundColor = bgColor;
//    NSView *contentView = window.contentView;
//    contentView.wantsLayer = YES;
//    contentView.layer.backgroundColor = bgColor.CGColor;
}

// 创建并配置窗口工具栏
void CreateToolbar(unsigned long nsWindowHandle, ToolbarConfiguration config, ControlEventCallback callback, void **outToolbarDelegate, void** outToolbar) {
    NSLog(@"CreateToolbar");
    NSWindow *window = (__bridge NSWindow *)(void *)nsWindowHandle;

    // 创建工具栏代理
    MainToolbarDelegate *toolbarDelegate = [[MainToolbarDelegate alloc] init];
    // 创建工具栏
    NSToolbar *toolbar = [[NSToolbar alloc] initWithIdentifier:@"ENERGY.ToolBar"];
    // 设置实例到当前代理对象
    [toolbarDelegate setCallback:callback withWindow:window withToolbar:toolbar];
    toolbar.delegate = toolbarDelegate;
    // 设置显示模式
    window.titlebarAppearsTransparent = config.Transparent;

    window.showsToolbarButton = config.ShowsToolbarButton;
    window.toolbarStyle = config.Style;
    window.titlebarSeparatorStyle = config.SeparatorStyle;
    toolbar.allowsUserCustomization = config.IsAllowsUserCustomization;
    toolbar.autosavesConfiguration = config.IsAutoSavesConfiguration;
    toolbar.displayMode = config.DisplayMode;
    toolbar.sizeMode = config.SizeMode; //NSToolbarSizeModeRegular; // 或 NSToolbarSizeModeSmall
    window.toolbar = toolbar;

    // 保留委托对象
    objc_setAssociatedObject(window, &kToolbarDelegateKey, toolbarDelegate, OBJC_ASSOCIATION_RETAIN_NONATOMIC);

    if (outToolbarDelegate) {
        *outToolbarDelegate = (__bridge void*)(toolbarDelegate);
    }
    if (outToolbar) {
        *outToolbar = (__bridge void*)(toolbar);
    }
}

// 向 toolbar 添加控件
void ToolbarAddItem(void* nsDelegate, void* nsToolbar, void* nsControl, const char *identifier, ControlProperty property) {
    if (!nsDelegate || !nsToolbar || !nsControl || !identifier) {
        NSLog(@"[ERROR] AddToolbarControl 必要入参为空");
        return;
    }
    MainToolbarDelegate *delegate = (MainToolbarDelegate*)nsDelegate;
    NSToolbar *toolbar = (NSToolbar*)nsToolbar;
    NSView *view = (NSView*)nsControl;
    NSString *idStr = [NSString stringWithUTF8String:identifier];
    if (!toolbar || !delegate || !view || !idStr) {
        NSLog(@"[ERROR] AddToolbarControl 必要转换参数为空");
        return;
    }
    // 添加到委托 维护, 工具栏获取时使用
    [delegate addControl:view forIdentifier:idStr withProperty:property];
    // 添加到工具栏
    [toolbar insertItemWithItemIdentifier:idStr atIndex:toolbar.items.count];
}

long ToolbarItemCount(void* nsToolbar) {
    if (!nsToolbar) {
        NSLog(@"[ERROR] AddToolbarControl 必要入参为空");
        return 0;
    }
    NSToolbar *toolbar = (NSToolbar*)nsToolbar;
    return toolbar.items.count;
}

#pragma mark - 动态控件创建函数

void AddToolbarCombobox(unsigned long nsWindowHandle, const char *identifier, const char **items, int count, ControlProperty property) {
    NSWindow *window = (__bridge NSWindow *)(void *)nsWindowHandle;
    MainToolbarDelegate *delegate = objc_getAssociatedObject(window, &kToolbarDelegateKey);
    NSString *idStr = [NSString stringWithUTF8String:identifier];
    // 创建下拉框
    NSComboBox *comboBox = [[NSComboBox alloc] init];
    comboBox.delegate = delegate;
    comboBox.controlSize = property.controlSize;
    [comboBox setEditable:NO];
    if (property.font) {
        comboBox.font = property.font;
    }
    // 添加选项
    for (int i = 0; i < count; i++) {
        [comboBox addItemWithObjectValue:[NSString stringWithUTF8String:items[i]]];
    }
    // 设置默认选择
    if (count > 0) {
        [comboBox selectItemAtIndex:0];
    }
    // 设置尺寸约束
    if (property.width > 0) {
        [comboBox.widthAnchor constraintEqualToConstant:property.width].active = YES;
    }
    if (property.height > 0) {
        [comboBox.heightAnchor constraintEqualToConstant:property.height].active = YES;
    }
    // 关联标识符
    objc_setAssociatedObject(comboBox, @"identifier", idStr, OBJC_ASSOCIATION_RETAIN);
    // 添加到委托
    [delegate addControl:comboBox forIdentifier:idStr withProperty:property];
    // 添加到工具栏
    [window.toolbar insertItemWithItemIdentifier:idStr atIndex:window.toolbar.items.count];
}

void AddToolbarCustomView(unsigned long nsWindowHandle, const char *identifier, ControlProperty property) {
    NSWindow *window = (__bridge NSWindow *)(void *)nsWindowHandle;
    MainToolbarDelegate *delegate = objc_getAssociatedObject(window, &kToolbarDelegateKey);
    NSString *idStr = [NSString stringWithUTF8String:identifier];
    // 创建自定义容器
    NSView *container = [[NSView alloc] init];
    // 设置尺寸约束
    if (property.width > 0) {
        [container.widthAnchor constraintEqualToConstant:property.width].active = YES;
    }
    if (property.height > 0) {
        [container.heightAnchor constraintEqualToConstant:property.height].active = YES;
    }
    // 关联标识符
    objc_setAssociatedObject(container, @"identifier", idStr, OBJC_ASSOCIATION_RETAIN);
    // 添加到委托
    [delegate addControl:container forIdentifier:idStr withProperty:property];
    // 添加到工具栏
    [window.toolbar insertItemWithItemIdentifier:idStr atIndex:window.toolbar.items.count];
}

#pragma mark - 工具栏管理函数

void RemoveToolbarItem(unsigned long nsWindowHandle, const char *identifier) {
    NSWindow *window = (__bridge NSWindow *)(void *)nsWindowHandle;
    MainToolbarDelegate *delegate = objc_getAssociatedObject(window, &kToolbarDelegateKey);

    if (!delegate) return;

    NSString *idStr = [NSString stringWithUTF8String:identifier];

    // 从委托中移除控件
    [delegate removeControlForIdentifier:idStr];

    // 从工具栏中移除项
    NSUInteger index = [window.toolbar.items indexOfObjectPassingTest:^BOOL(NSToolbarItem * _Nonnull obj, NSUInteger idx, BOOL * _Nonnull stop) {
        return [obj.itemIdentifier isEqualToString:idStr];
    }];

    if (index != NSNotFound) {
        [window.toolbar removeItemAtIndex:index];
    }
}

void UpdateToolbarItemProperty(unsigned long nsWindowHandle, const char *identifier, ControlProperty property) {
    NSWindow *window = (__bridge NSWindow *)(void *)nsWindowHandle;
    MainToolbarDelegate *delegate = objc_getAssociatedObject(window, &kToolbarDelegateKey);

    if (!delegate) return;

    NSString *idStr = [NSString stringWithUTF8String:identifier];
    [delegate updateControlProperty:idStr withProperty:property];
}

void InsertToolbarItemAtIndex(unsigned long nsWindowHandle, const char *identifier, int index) {
    NSWindow *window = (__bridge NSWindow *)(void *)nsWindowHandle;
    MainToolbarDelegate *delegate = objc_getAssociatedObject(window, &kToolbarDelegateKey);

    if (!delegate) return;

    NSString *idStr = [NSString stringWithUTF8String:identifier];

    // 确保索引在有效范围内
    NSUInteger itemCount = window.toolbar.items.count;
    NSUInteger insertIndex = MIN(MAX(index, 0), itemCount);

    // 从当前位置移除（如果存在）
    NSUInteger currentIndex = [window.toolbar.items indexOfObjectPassingTest:^BOOL(NSToolbarItem * _Nonnull obj, NSUInteger idx, BOOL * _Nonnull stop) {
        return [obj.itemIdentifier isEqualToString:idStr];
    }];

    if (currentIndex != NSNotFound) {
        [window.toolbar removeItemAtIndex:currentIndex];
        // 如果当前索引在插入索引之前，需要调整插入索引
        if (currentIndex < insertIndex) {
            insertIndex--;
        }
    }

    // 插入到新位置
    [window.toolbar insertItemWithItemIdentifier:idStr atIndex:insertIndex];
}


// 循环工具栏每项获取 NSControl，通过代理获取有问题啊。
NSView* GetToolbarControl(unsigned long nsWindowHandle, const char *identifier) {
    NSWindow *window = (__bridge NSWindow *)(void *)nsWindowHandle;
    NSString *idStr = [NSString stringWithUTF8String:identifier];

    // 使用代理获取 controls
    MainToolbarDelegate *delegate = objc_getAssociatedObject(window, &kToolbarDelegateKey);
    NSView *control = [delegate controlForIdentifier:idStr];
    if (!control) return nil;
    return control;
}

#pragma mark - 控件管理函数

const char *GetToolbarControlValue(unsigned long nsWindowHandle, const char *identifier) {
    NSView *control = GetToolbarControl(nsWindowHandle, identifier);
    if (!control) return NULL;
    NSString *idStr = [NSString stringWithUTF8String:identifier];

    if ([control isKindOfClass:[NSTextField class]]) {
        return [[(NSTextField *)control stringValue] UTF8String];
    }
    else if ([control isKindOfClass:[NSComboBox class]]) {
        return [[(NSComboBox *)control stringValue] UTF8String];
    }
    else if ([control isKindOfClass:[NSSearchField class]]) {
        return [[(NSSearchField *)control stringValue] UTF8String];
    }
    return NULL;
}

void SetToolbarControlValue(unsigned long nsWindowHandle, const char *identifier, const char *value) {
    NSView *control = GetToolbarControl(nsWindowHandle, identifier);
    if (!control) return;
    NSString *valueStr = [NSString stringWithUTF8String:value];
    if ([control isKindOfClass:[NSTextField class]]) {
        [(NSTextField *)control setStringValue:valueStr];
    }
    else if ([control isKindOfClass:[NSComboBox class]]) {
        [(NSComboBox *)control setStringValue:valueStr];
    }
    else if ([control isKindOfClass:[NSSearchField class]]) {
        [(NSSearchField *)control setStringValue:valueStr];
    }
}

void SetToolbarControlEnabled(unsigned long nsWindowHandle, const char *identifier, bool enabled) {
    NSView *control = GetToolbarControl(nsWindowHandle, identifier);
    if (!control) return;
    if ([control isKindOfClass:[NSControl class]]) {
        [(NSControl *)control setEnabled:(BOOL)enabled];
    }
}

void SetToolbarControlHidden(unsigned long nsWindowHandle, const char *identifier, bool hidden) {
    NSView *control = GetToolbarControl(nsWindowHandle, identifier);
    if (!control) {
        NSLog(@"获取 NSView(control)失败");
        return;
    }
    if ([control isKindOfClass:[NSControl class]]) {
        [(NSControl *)control setHidden:(BOOL)hidden];
    }
}
