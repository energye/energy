//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

(function () {
    const MT_DRAG_DROP_ENTER = 100;
    const MT_DRAG_DROP_LEAVE = MT_DRAG_DROP_ENTER + 1;
    const MT_DRAG_DROP_OVER = MT_DRAG_DROP_LEAVE + 1;

    const DragTypeNo = 0
    const DragTypeFile = DragTypeNo + 1; // file list
    const DragTypeData = DragTypeFile + 1;

    const docElement = document.documentElement;
    const formats = ['text/plain', 'Files'];
    let dragCount = 0;
    let isDrag = false;

    const processMessage = window?.energy?.processMessage;
    const postMessageWithAdditionalObjects = function (message, additionalObjects) {
        if (window.chrome?.webview?.postMessageWithAdditionalObjects) {
            window.chrome.webview.postMessageWithAdditionalObjects(message, additionalObjects);
        }
    }
    const makePayload = window?.energy?.makePayload;

    const isFormats = function (event) {
        let matchedFormat = null;
        for (const format of formats) {
            if (event.dataTransfer?.types.includes(format)) {
                matchedFormat = format;
                break;
            }
        }
        if (matchedFormat !== null) {
            let result = {
                type: DragTypeNo,
                format: matchedFormat
            };
            if (matchedFormat === formats[0]) {
                result.type = DragTypeData;
            } else if (matchedFormat === formats[1]) {
                result.type = DragTypeFile;
            }
            return result;
        }
        return null;
    }

    // 拖拽进入事件
    docElement.addEventListener('dragenter', (event) => {
        const format = isFormats(event);
        if (format == null) {
            return;
        }
        dragCount++;
        if (dragCount === 1 && !isDrag) {
            isDrag = true;
            const data = {type: format.type, x: event.clientX, y: event.clientY};
            const payload = makePayload(MT_DRAG_DROP_ENTER, "", data, 0);
            processMessage(payload);
        }
    }, {passive: true});

    // 拖拽离开事件
    docElement.addEventListener('dragleave', (event) => {
        const format = isFormats(event);
        if (format == null) {
            return;
        }
        dragCount--;
        if (dragCount <= 0 && isDrag) {
            dragCount = 0;
            isDrag = false;
            const data = {type: format.type, x: event.clientX, y: event.clientY};
            const payload = makePayload(MT_DRAG_DROP_LEAVE, "", data, 0);
            processMessage(payload);
        }
    }, {passive: true});

    // 拖拽结束事件（包含 放置完成 和 未放置直接离开）
    docElement.addEventListener('drop', (event) => {
        const format = isFormats(event);
        if (format == null) {
            return;
        }
        if (isDrag) {
            dragCount = 0;
            isDrag = false;
            const data = {type: format.type, x: event.clientX, y: event.clientY, text: ""};
            const dataTransfer = event.dataTransfer;
            if (format.type === DragTypeFile) {
                const files = [];
                for (const item of dataTransfer.items) {
                    if (item.kind === 'file') {
                        const file = item.getAsFile();
                        files.push(file);
                    }
                }
                const payload = makePayload(MT_DRAG_DROP_OVER, "", data, 0);
                postMessageWithAdditionalObjects(payload, files);
            } else if (format.type === DragTypeData) {
                data.text = dataTransfer.getData(format.format);
                const payload = makePayload(MT_DRAG_DROP_OVER, "", data, 0);
                processMessage(payload);
            }
        }
    }, {passive: true});
})();