package implant

import (
	"os"
	"path/filepath"
)

var JS string

func init() {
	wd, _ := os.Getwd()
	// 读取 js 扩展文件
	// helper.js 来自 rod, 根据情况自己增加或修改
	path := filepath.Join(wd, "examples", "crawling-web-pages", "webkit-bind-javascript", "implant", "helper.js")
	data, _ := os.ReadFile(path)
	JS = string(data)
	// 自定义的内置本地函数
	JS += `
(function () {
    // 在 JS 里的 dom 对象定义html字段对象名的set方法
    dom.__defineSetter__('html', function (e) {
        native function html();
        html(e);
    });
})();
`
}
