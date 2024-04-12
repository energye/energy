// url页面加载进度
ipc.on("open-url-process", function (value, windowId) {
    $("#loadProcess").html(value)
})

// 仅测试区分测试的功能类型
const testTypeDefault = 0
const testTypeUpload = 1
const testTypeDownload = 2

$(function () {
    // 默认加载出当前已创建的窗口，例如刷新页面后
    ipc.emit("window-infos", function (result) {
        console.log("window-infos list:", JSON.stringify(result))
        for (let i in result) {
            let data = result[i]
            let url = data.URL
            if (url === "") {
                url = defaultURL
            }
            $("#box").append(create(data.WindowId, url, data.Typ))
        }
    })


    let defaultURL = "https://gitee.com" // 给个默认地址

    // 创建页面的功能按钮
    let create = function (windowId, url, type) {
        let html = `
<div className="row" id="${windowId}">
    <span>${windowId}</span>
    <input id="url" style="width: 250px" value="${url}" readonly>
    <button id="show">打开</button>
    <button id="closeWindow">关闭</button>
    <button id="crawling">打开后-测试</button>
    <span id="loadProcess"> - </span>
</div>`
        let row = $(html)
        let showBtn = row.find("#show")
        let urlInp = row.find("#url")
        showBtn.click(function () {
            let url = urlInp.val()
            if (url === "") {
                return
            }
            ipc.emit("show", [windowId, url])
        })
        // 关闭这个窗口
        let closeWindowBtn = row.find("#closeWindow")
        closeWindowBtn.click(function () {
            ipc.emit("close-window", [windowId], function (result) {
                if (result) {
                    row.remove()
                }
            })
        })
        // 测试按钮
        let crawlingBtn = row.find("#crawling")
        // 抓取一些内容
        crawlingBtn.click(function () {
            ipc.emit("crawling", [windowId, type], function (result) {
                console.log("crawling-result:", result)
            })
        })

        return row
    }

    // 主窗口功能, 以下功能全部在主窗口运行

    $("#create").click(function () {
        // defaultURL 默认地址
        ipc.emit("create", [defaultURL, testTypeDefault], function (windowId) {
            console.log("create windowId:", windowId)
            if (windowId > 0) {
                $("#box").append(create(windowId, defaultURL, testTypeDefault))
            }
        })
    })
    // 关闭指定窗口
    ipc.on("close-window", function (windowId) {
        console.log("close-windowId:", windowId)
        $("#" + windowId).remove()
    })
    // 创建一个窗口
    ipc.on("create-window", function (windowId, url) {
        console.log("create-windowId:", windowId, "url:", url)
        $("#box").append(create(windowId, url, testTypeDefault))
    })
    // 窗口的加载进度
    ipc.on("window-loading-progress", function (windowId, progress) {
        $("#" + windowId).find("#loadProcess").html("Loading: " + progress)
    })

    // 上传文件

    $("#upload").click(function () {
        ipc.emit("upload-start-server", [testTypeUpload], function (url, windowId) {
            $("#box").append(create(windowId, url, testTypeUpload))
        })
    })

    // 下载文件

    $("#download").click(function () {
        ipc.emit("download-file", [testTypeDownload], function (url, windowId) {
            $("#box").append(create(windowId, url, testTypeDownload))
        })
    })
})
