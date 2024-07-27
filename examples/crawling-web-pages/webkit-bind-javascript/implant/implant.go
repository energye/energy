package implant

import (
	_ "embed"
)

//go:embed helper.js
var helperJS []byte

func HelperJS() string {
	// 自定义的内置本地函数
	result := string(helperJS) + `
(function () {
    // 在 JS 里的 dom 对象定义html字段对象名的set方法
    dom.__defineSetter__('html', function (e) {
        native function html();
        html(e);
    });
})();
`
	return result
}
