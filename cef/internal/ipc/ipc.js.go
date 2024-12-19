//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package ipc

var IPCJS = `
let __drag;
if (!__drag) {
    __drag = {
        cssDragProperty: "-webkit-app-region",
        cssDragValue: "drag",
        os: null,
        enableResize: true,
    };
}
(function () {
    let shouldDrag = false;
    let idcCursor = null;
    let frameWidth = 4;
    let frameHeight = 4;
    let frameCorner = 8;

    function isWindows() {
        return __drag.os === "windows";
    }

    function isDarwin() {
        return __drag.os === "darwin";
    }

    function isLinux() {
        return __drag.os === "linux";
    }

    function setCursor(cursor, ht) {
        if (idcCursor !== cursor) {
            document.documentElement.style.cursor = cursor || 'auto';
            idcCursor = cursor;
        }
    }

    function enableResize() {
        return __drag.enableResize;
    }

    function mouseDragResize(e) {
        let leftBorder = e.clientX < frameWidth;
        let topBorder = e.clientY < frameHeight;
        let rightBorder = window.outerWidth - e.clientX < frameWidth;
        let bottomBorder = window.outerHeight - e.clientY < frameHeight;
        let leftCorner = e.clientX < frameWidth + frameCorner;
        let topCorner = e.clientY < frameHeight + frameCorner;
        let rightCorner = window.outerWidth - e.clientX < frameWidth + frameCorner;
        let bottomCorner = window.outerHeight - e.clientY < frameHeight + frameCorner;
        if (!leftBorder && !topBorder && !rightBorder && !bottomBorder && idcCursor !== void 0) {
            setCursor();
        } else if (rightCorner && bottomCorner) {
            setCursor("se-resize", 17);
        } else if (leftCorner && bottomCorner) {
            setCursor("sw-resize", 16);
        } else if (leftCorner && topCorner) {
            setCursor("nw-resize", 13);
        } else if (topCorner && rightCorner) {
            setCursor("ne-resize", 14);
        } else if (leftBorder) {
            setCursor("w-resize", 10);
        } else if (topBorder) {
            setCursor("n-resize", 12);
        } else if (bottomBorder) {
            setCursor("s-resize", 15);
        } else if (rightBorder) {
            setCursor("e-resize", 11);
        }
    }

    function test(e) {
        let v = window.getComputedStyle(e.target)[__drag.cssDragProperty];
        if (v) {
            v = v.trim();
            if (v !== __drag.cssDragValue) {
                return false;
            }
            return e.detail === 1 || e.detail === 2;
        }
        return false;
    }

    function mouseMove(e) {
        if (shouldDrag) {
            if (isWindows() || isDarwin()) {
                shouldDrag = false;// && !IsDarwin();
            }
            energyExtension.mouseMove({x: e.screenX, y: e.screenY, ts: parseInt(e.timeStamp)});
        } else if (enableResize() && isWindows()) {
            mouseDragResize(e);
        }
    }

    function mouseUp(e) {
        shouldDrag = false;
    }

    function mouseDown(e) {
        if (enableResize() && idcCursor) {
            e.preventDefault();
            energyExtension.mouseResize(idcCursor);
        } else if (!(e.offsetX > e.target.clientWidth || e.offsetY > e.target.clientHeight) && test(e)) {
            shouldDrag = true;
            energyExtension.mouseDown({x: e.screenX, y: e.screenY, ts: parseInt(e.timeStamp)});
        } else {
            shouldDrag = false;
        }
    }

    function dblClick(e) {
        if (test(e)) {
            e.preventDefault();
            energyExtension.mouseDblClick();
        }
    }
    __drag.setup = function () {
        window.addEventListener("mousemove", mouseMove);
        window.addEventListener("mousedown", mouseDown);
        window.addEventListener("mouseup", mouseUp);
        window.addEventListener("dblclick", dblClick);
    }
    /**
     * 在JS里判断 x y 点是否在指定区域，以下步骤
     * 1. CreateRectRgn(x, y, x+w, y+h) 创建一个区域 x, y, w, h
     * 2. CreateRectRgn 创建当前区域
     * 3. CombineRgn 合并区域
     * 4. PtInRegion(rgn, x, y) 检查点是否在区域
     */
})();
`
