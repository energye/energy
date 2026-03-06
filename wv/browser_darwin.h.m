#import "Cocoa/Cocoa.h"
#import <WebKit/WebKit.h>
#import <objc/message.h>
#import <Foundation/Foundation.h>
#import "browser_darwin.h"


// 设置webview透明属性
void SetWebviewTransparent(void* nsWebview, int isTransparent) {
	WKWebView* webview = (WKWebView*)nsWebview;
    if (!webview) {
        NSLog(@"SetWebviewTransparent webview nil");
        return;
    }
//   NSLog(@"SetWebviewTransparent %d", isTransparent);
   [webview setValue:[NSNumber numberWithBool:isTransparent] forKey:@"drawsBackground"];
}

void UpdateWebviewBounds(void* nsWindow, void* nsWebview, float x, float y, float width, float height) {
	NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        NSLog(@"UpdateWebviewBounds window nil");
        return;
    }
	WKWebView* webview = (WKWebView*)nsWebview;
    if (!webview) {
        NSLog(@"WindowAddSubview webview nil");
        return;
    }
	NSView* contentView = window.contentView;
//	CGRect viewFrame = [contentView bounds];
    CGRect viewFrame = CGRectMake(x, y, width, height);
//	NSLog(@"UpdateWebviewBounds contentView bounds: %@", NSStringFromRect(viewFrame));
	[webview setFrame:viewFrame];
}

void WebViewBecomeFirstResponder(void* nsWebview) {
	WKWebView* webview = (WKWebView*)nsWebview;
    if (!webview) {
        NSLog(@"WebViewBecomeFirstResponder webview nil");
        return;
    }
    [webview becomeFirstResponder];
}

void WebViewUndo(void* nsWebview) {
	WKWebView* webview = (WKWebView*)nsWebview;
    if (!webview) {
        NSLog(@"WebViewUndo webview nil");
        return;
    }
    if (webview.undoManager) {
        NSUndoManager *undoManager = webview.undoManager;
        if ([undoManager canUndo]) {
            [undoManager undo];
        }
    }
}

void WebViewRedo(void* nsWebview) {
    WKWebView* webview = (WKWebView*)nsWebview;
    if (!webview) {
        NSLog(@"WebViewRedo webview nil");
        return;
    }
    if (webview.undoManager) {
        NSUndoManager *undoManager = webview.undoManager;
        if ([undoManager canRedo]) {
            [undoManager redo];
        }
    }
}

void WebViewCut(void* nsWebview) {
    WKWebView* webview = (WKWebView*)nsWebview;
    if (!webview) {
        NSLog(@"WebViewCut webview nil");
        return;
    }
    SEL selector = sel_registerName("cut:");
    ((void (*)(id, SEL, id))objc_msgSend)(webview, selector, webview);
}

void WebViewCopy(void* nsWebview) {
    WKWebView* webview = (WKWebView*)nsWebview;
    if (!webview) {
        NSLog(@"WebViewCopy webview nil");
        return;
    }
    SEL selector = sel_registerName("copy:");
    ((void (*)(id, SEL, id))objc_msgSend)(webview, selector, webview);
}

void WebViewPaste(void* nsWebview) {
    WKWebView* webview = (WKWebView*)nsWebview;
    if (!webview) {
        NSLog(@"WebViewPaste webview nil");
        return;
    }
    SEL selector = sel_registerName("paste:");
    ((void (*)(id, SEL, id))objc_msgSend)(webview, selector, webview);
}

void WebViewSelectAll(void* nsWebview) {
    WKWebView* webview = (WKWebView*)nsWebview;
    if (!webview) {
        NSLog(@"WebViewSelectAll webview nil");
        return;
    }
    SEL selector = sel_registerName("selectAll:");
    ((void (*)(id, SEL, id))objc_msgSend)(webview, selector, webview);
}

// webview 快捷键
static BOOL Webview_performKeyEquivalent(id self, SEL _cmd, NSEvent *event) {
    //NSLog(@"Webview_performKeyEquivalent");
    NSUInteger modifierFlags = [event modifierFlags];
    NSInteger keyCode = [event keyCode];
    NSUInteger cmd = [event modifierFlags] & NSEventModifierFlagCommand;
    NSUInteger shift = modifierFlags & NSEventModifierFlagShift;
    NSUInteger option = modifierFlags & NSEventModifierFlagOption;
    //NSLog(@"Webview_performKeyEquivalent cmd: %d shift: %d keyCode: %d", (int)cmd, (int)shift, (int)keyCode);
    SEL selector = NULL;
    if (cmd && !shift && !option) {
        switch (keyCode) {
            case 0:  // Cmd+A → 全选
                selector = sel_registerName("selectAll:");
                break;
            case 8:  // Cmd+C → 复制
                selector = sel_registerName("copy:");
                break;
            case 9:  // Cmd+V → 粘贴
                selector = sel_registerName("paste:");
                break;
            case 7:  // Cmd+X → 剪切
                selector = sel_registerName("cut:");
                break;
            case 6:  // Cmd+Z → 撤销
                WebViewUndo(self);
                return YES;
            case 15:  // Cmd+R → 刷新网页
                selector = sel_registerName("reload:");
                break;
            case 35: // Cmd+P → 打印网页
                //selector = sel_registerName("print:");
                break;
        }
    } else if (cmd && shift && !option) {
         switch (keyCode) {
             case 6:  // Cmd+Shift+Z → 重做
                WebViewRedo(self);
                return YES;
         }
    }
    if (selector) {
        // 在UI线程运行
        //dispatch_async(dispatch_get_main_queue(), ^{
            // ...
        //});
        if (selector && [self respondsToSelector:selector]) {
            ((void (*)(id, SEL, id))objc_msgSend)(self, selector, self);
        } else {
            NSLog(@"performKeyEquivalent - not support method：%@", NSStringFromSelector(selector));
        }
        return YES;
    }
    return NO;
}

void WebViewRegisterPerformKeyEquivalentMethod(void* nsWebview) {
	WKWebView* webview = (WKWebView*)nsWebview;
    if (!webview) {
        NSLog(@"WebViewRegisterPerformKeyMethod webview nil");
        return;
    }
    Class webviewClass = [webview class];
    const char* methodType = "B@:@"; // B=BOOL, @=id(self), :=SEL(_cmd), @=NSEvent*
    SEL targetSel = sel_registerName("performKeyEquivalent:");
    class_addMethod(webviewClass, targetSel, (IMP)Webview_performKeyEquivalent, methodType);
}

void WebViewEvaluateScriptCallback(void* nsWebview, int callbackID, const char* cScript, CGoEvaluateScriptCallback cGoCallback) {
	WKWebView* webview = (WKWebView*)nsWebview;
    if (!webview) {
        NSLog(@"WebViewEvaluateScriptCallback webview nil");
        return;
    }
    NSString *script = [NSString stringWithUTF8String:cScript];
    // OpaqueCBlock
    void (^completionHandler)(id, NSError *) = ^(id result, NSError *error) {
        const char* resStr = NULL;
        const char* errStr = NULL;
        if (result != nil) {
            if ([result isKindOfClass:[NSString class]]) {
                resStr = [result UTF8String];
            } else {
                NSData *jsonData = [NSJSONSerialization dataWithJSONObject:result options:0 error:nil];
                if (jsonData) {
                    NSString *jsonStr = [[NSString alloc] initWithData:jsonData encoding:NSUTF8StringEncoding];
                    resStr = [jsonStr UTF8String];
                }
            }
        }
        if (error != nil) {
            errStr = [error.localizedDescription UTF8String];
        }
        // Go 回调函数
        cGoCallback(callbackID, resStr, errStr);
    };
    // UI 线程
    dispatch_async(dispatch_get_main_queue(), ^{
        [webview evaluateJavaScript:script completionHandler:completionHandler];
    });
}

NSPoint ConvertPoint(void* nsWebview, float x, float y) {
	WKWebView* webview = (WKWebView*)nsWebview;
    if (!webview) {
        NSLog(@"ConvertPoint webview is nil");
        return NSMakePoint(0, 0);
    }
    NSPoint inPoint = NSMakePoint(x, y);
    NSPoint pointInView = [webview convertPoint:inPoint fromView:nil];
    return pointInView;
}