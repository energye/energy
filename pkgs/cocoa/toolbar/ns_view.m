#import "config.h"
#import "ns_view.h"

@implementation CustomView

- (instancetype)initWithFrame:(NSRect)frame {
    self = [super initWithFrame:frame];
    if (self) {
        _backgroundColor = [NSColor systemBlueColor]; // 默认颜色
        self.autoresizingMask = NSViewNotSizable;
    }
    return self;
}

- (void)drawRect:(NSRect)dirtyRect {
    [super drawRect:dirtyRect];
    [self.backgroundColor setFill];
    NSRectFill(dirtyRect);
}

- (NSSize)intrinsicContentSize {
    return NSMakeSize(200, 38); // 自定义宽度和高度
}

@end

void* NewCustomView() {
    // 测试还有问题
    NSRect frame = NSMakeRect(0, 0, 200, 38);
    CustomView *customView = [[CustomView alloc] initWithFrame:frame];
    customView.backgroundColor = [NSColor systemBlueColor]; // 设置填充颜色
    // 设置尺寸约束
    [customView.widthAnchor constraintEqualToConstant:200].active = YES;
    [customView.heightAnchor constraintEqualToConstant:38].active = YES;

    return (__bridge void*)(customView);
}