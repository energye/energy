package crawling

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/pkgs/json"
)

// GMessageId devtools 全局消息id，唯一自动生成
var GMessageId int32 = 0

// CrawlingWindow 抓取窗口
type CrawlingWindow struct {
	Window     cef.IBrowserWindow
	ResultData chan []byte // 接收 devtools 返回的数据
}

// 维护所有窗口的集合
var windows = make(map[int]*CrawlingWindow)

// GetObjectId objectId 是要在其上下文中执行函数的对象的 ID。
// 如果函数要在全局上下文中执行，这个值应该是从 Runtime.enable 方法获取的全局上下文的 ID。
func GetObjectId(resultData chan []byte, chromium cef.IChromium) string {
	var dict = cef.DictionaryValueRef.New()
	dict.SetString("expression", "window")
	GMessageId = chromium.ExecuteDevToolsMethod(GMessageId, "Runtime.evaluate", dict)
	// 不去阻止ipc, 开启协程处理
	data := <-resultData
	fmt.Println("data", string(data))
	jData := json.NewJSON(data).JSONObject()
	result := jData.GetObjectByKey("result").GetObjectByKey("result")
	objectId := result.GetStringByKey("objectId")
	//GetDom(resultData, objectId, chromium, "#wrapper")
	return objectId
}

func GetDom(resultData chan []byte, objectId string, chromium cef.IChromium, selector string) {
	// {
	//    "id": 14,
	//    "sessionId": "996EFC980C7AEE2A82EB93463E39CBE9",
	//    "method": "Runtime.callFunctionOn",
	//    "params": {
	//        "functionDeclaration": "function() { return (function (f /* element */, ...args) { return f.apply(this, args) }).apply(this, arguments) }",
	//        "objectId": "6929699160585610419.1.1",
	//        "arguments": [
	//            {
	//                "value": null,
	//                "objectId": "6929699160585610419.1.4"
	//            },
	//            {
	//                "value": "#wrapper"
	//            }
	//        ]
	//    }
	// }

	sendMessage := `{"id":10,"method":"Runtime.evaluate","params":{"expression":"window"}}`
	sendMessage = `{
   "id": 14,
   "method": "Runtime.callFunctionOn",
   "params": {
       "functionDeclaration": "function() { return document.body.innerHTML }",
       "objectId": "` + objectId + `",
       "arguments": []
   }
}`
	fmt.Println("sendMessage:", sendMessage)
	success := chromium.SendDevToolsMessage(sendMessage)
	fmt.Println("success:", success)

	data := <-resultData
	fmt.Println("data", string(data))
}

//// EvalOptions for Page.Evaluate.
//type EvalOptions struct {
//	// If enabled the eval result will be a plain JSON value.
//	// If disabled the eval result will be a reference of a remote js object.
//	ByValue bool
//
//	AwaitPromise bool
//
//	// ThisObj represents the "this" object in the JS
//	ThisObj *methods.RuntimeRemoteObject
//
//	// JS function definition to execute.
//	JS string
//
//	// JSArgs represents the arguments that will be passed to JS.
//	// If an argument is [*proto.RuntimeRemoteObject] type, the corresponding remote object will be used.
//	// Or it will be passed as a plain JSON value.
//	// When an arg in the args is a *js.Function, the arg will be cached on the page's js context.
//	// When the arg.Name exists in the page's cache, it reuse the cache without sending
//	// the definition to the browser again.
//	// Useful when you need to eval a huge js expression many times.
//	JSArgs []interface{}
//
//	// Whether execution should be treated as initiated by user in the UI.
//	UserGesture bool
//}
