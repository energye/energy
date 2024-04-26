// url页面加载进度
$(function () {
    $("#execute").click(function () {
        ipc.emit("execute")
    })
   $("#JSGen").click(function () {
       let time = new Date().getTime()
       let rowHtml = `
<div class="row">
    ${time}
</div>`
       $("#box").append($(rowHtml))
   })
})
