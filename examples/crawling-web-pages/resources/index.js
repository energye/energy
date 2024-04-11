// url页面加载进度
ipc.on("open-url-process", function (value, windowId) {
    $("#loadProcess").html(value)
})

$(function () {
    // 打开一个地址
    $("#openUrl").click(function () {
        var me = $(this)
        var url = me.parent().find("#url").val()
        if (url === "") {
            return
        }
        var window = me.parent().find("#windowId")
        var windowId = parseInt(window.val())
        ipc.emit("open-url-window", [url, windowId], function (windowId) {
            console.log("windowId:", windowId)
            window.val(windowId)
        })
    })
    // 关闭这个窗口
    $("#closeWindow").click(function () {
        var windowId = parseInt($(this).parent().find("#windowId").val())
        console.log("close-window id:", windowId)
        ipc.emit("close-window", [windowId])
    })
    //
    $("#crawling").click(function () {
        var windowId = parseInt($(this).parent().find("#windowId").val())
        console.log("crawling-window id:", windowId)
        ipc.emit("crawling", [windowId], function (result) {
            console.log("crawling-result:", result)
        })
    })
})
