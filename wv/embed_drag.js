// console.log("postMessageWithAdditionalObjects:",window.chrome.webview.postMessageWithAdditionalObjects)
// if (window.chrome.webview.postMessageWithAdditionalObjects) {
//     window.chrome.webview.postMessageWithAdditionalObjects(`file:drop:1:2`, dragResult.files);
// }

(function () {
    const MT_DRAG_DROP_ENTER = 100;
    const MT_DRAG_DROP_LEAVE = MT_DRAG_DROP_ENTER + 1;
    const MT_DRAG_DROP_OVER = MT_DRAG_DROP_LEAVE + 1;

    const DragTypeNo = 0
    const DragTypeFile = DragTypeNo + 1; // file list
    const DragTypeData = DragTypeFile + 1;


    const docElement = document.documentElement;
    const formats = ['text/plain', 'Files'];
    let dragEnterCount = 0;
    let isDragging = false;

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

    // 1. 拖拽进入事件
    docElement.addEventListener('dragenter', (event) => {
        const format = isFormats(event);
        if (format == null) {
            return;
        }
        dragEnterCount++;
        if (dragEnterCount === 1 && !isDragging) {
            isDragging = true;
            const data = {type: format.type, x: event.clientX, y: event.clientY};
            const payload = makePayload(MT_DRAG_DROP_ENTER, "", data, 0);
            processMessage(payload);
        }
    }, {passive: true});

    // 2. 拖拽离开事件
    docElement.addEventListener('dragleave', (event) => {
        const format = isFormats(event);
        if (format == null) {
            return;
        }
        dragEnterCount--;
        if (dragEnterCount <= 0 && isDragging) {
            dragEnterCount = 0;
            isDragging = false;
            const data = {type: format.type, x: event.clientX, y: event.clientY};
            const payload = makePayload(MT_DRAG_DROP_LEAVE, "", data, 0);
            processMessage(payload);
        }
    }, {passive: true});

    // 3. 拖拽结束事件（包含 放置完成 和 未放置直接离开）
    docElement.addEventListener('drop', (event) => {
        const format = isFormats(event);
        if (format == null) {
            return;
        }
        if (isDragging) {
            dragEnterCount = 0;
            isDragging = false;
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