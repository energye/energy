// url页面加载进度
ipc.on("open-url-process", function (value, windowId) {
    $("#loadProcess").html(value)
})

$(function () {
    let defaultURL = "https://gitee.com"
    let create = function (windowId, url) {
        let html = `
<div className="row" id="${windowId}">
    <span>windowId: ${windowId}</span>
    <input id="url" style="width: 250px" value="${url}">
    <button id="show">打开</button>
    <button id="closeWindow">关闭</button>
    <button id="crawling">页面打开后-测试一下</button>
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
        let crawlingBtn = row.find("#crawling")
        // 抓取一些内容
        crawlingBtn.click(function () {
            ipc.emit("crawling", [windowId], function (result) {
                console.log("crawling-result:", result)
            })
        })
        return row
    }
    $("#create").click(function () {
        ipc.emit("create", [], function (windowId) {
            console.log("create windowId:", windowId)
            if (windowId > 0) {
                $("#box").append(create(windowId, defaultURL))
            }
        })
    })
    ipc.on("close-window", function (windowId) {
        console.log("close-windowId:", windowId)
        $("#" + windowId).remove()
    })
    ipc.on("create-window", function (windowId, url) {
        console.log("create-windowId:", windowId, "url:", url)
        $("#box").append(create(windowId, url))
    })
    ipc.on("window-loading-progress", function (windowId, progress) {
        $("#" + windowId).find("#loadProcess").html("Loading: " + progress)
    })
    ipc.emit("window-infos", [], function (result) {
        console.log("window-infos list:", JSON.stringify(result))
        for (let i in result) {
            let data = result[i]
            let url = data.URL
            if (url === "") {
                url = defaultURL
            }
            $("#box").append(create(data.WindowId, url))
        }
    })
})
