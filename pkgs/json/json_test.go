package json

import (
	"testing"
)

func TestJSONObject(t *testing.T) {
	testObject := NewJSONObject(nil)
	testObject.Set("str", "字符串")
	testObject.Set("str1", "字符串1")
	testObject.Set("int1", 100)
	testObject.Set("int2", 200)
	testObject.Set("bool1", true)
	testObject.Set("bool2", false)
	t.Log("testObject", testObject.ToJSONString())
	t.Log("testObject-size:", testObject.Size(), "==6")
	testObject.RemoveByKey("bool1")
	t.Log("testObject.remove-size:", testObject.Size(), "==5")
	//
	test2Object := NewJSONObject(nil)
	test2Object.Set("key1", "key1value")
	test2Object.Set("key2", testObject)
	t.Log("test2Object-size:", test2Object.Size(), "==2")
	//t.Log("testObject", test2Object.JsonData().V)
	t.Log("test2Object:", test2Object.ToJSONString())
	test2Key2 := test2Object.GetObjectByKey("key2")
	test2Key2.Set("bool1", true)
	t.Log("test2Object.test2Key2:", test2Key2.ToJSONString())
	t.Log("test2Object.test2Key2:", testObject.ToJSONString())
	t.Log("test2Object.key2-testObject-size:", test2Key2.Size(), testObject.Size(), "=6")
	test2Key2.Set("bool5", true)
	t.Log("test2Object.test2Key2:", test2Key2.ToJSONString())
	t.Log("test2Object.test2Key2:", testObject.ToJSONString())
	t.Log("test2Object.key2-testObject-size:", test2Key2.Size(), testObject.Size(), "=7")
	t.Log("--------------------------------------------------- 1 ---------------------------------------------------------")
	test2Key2.RemoveByKey("bool1")
	t.Log("test2Object.test2Key2:", test2Key2.ToJSONString())
	t.Log("test2Object.test2Key2:", testObject.ToJSONString())
	t.Log("test2Object.key2-testObject-size:", test2Key2.Size(), testObject.Size(), "=6")
	test2Key2.RemoveByKey("bool5")
	t.Log("test2Object.test2Key2:", test2Key2.ToJSONString())
	t.Log("test2Object.test2Key2:", testObject.ToJSONString())
	t.Log("test2Object.key2-testObject-size:", test2Key2.Size(), testObject.Size(), "=5")
	t.Log("--------------------------------------------------- 2 ---------------------------------------------------------")
	t.Log("test2Object-key1:", test2Object.GetStringByKey("key1"), "=key1value")
	t.Log("--------------------------------------------------- 3 ---------------------------------------------------------")
	testArray := NewJSONArray(nil)
	testArray.Add("第1个值", 2222, 3333.33, true)
	testArray.Add("第5个值")
	t.Log("testArray-size:", testArray.Size(), "=5")
	t.Log("testArray-remove-3")
	testArray.RemoveByIndex(3)
	t.Log("testArray-size:", testArray.Size(), "=4")
	t.Log("testArray-add")
	testArray.Add("第5个新值")
	t.Log("testArray-size:", testArray.Size(), "=5")
	t.Log("testArray-3:", testArray.GetStringByIndex(3), "=第5个值")
	t.Log("--------------------------------------------------- 4 ---------------------------------------------------------")
	testArray.SetByIndex(3, true)
	t.Log("testArray-3:", testArray.GetBoolByIndex(3), "=true")
	t.Log("--------------------------------------------------- 5 ---------------------------------------------------------")
	testArray.SetByIndex(3, testObject)
	object3 := testArray.GetObjectByIndex(3)
	t.Log("testArray-3.IsObject:", object3.IsObject(), "=true")
	t.Log("testArray.3-object3-size:", object3.Size(), testObject.Size(), "=5")
	object3.Set("key7", "字符串7")
	t.Log("testArray.3-object3-size:", object3.Size(), testObject.Size(), "=6")
	t.Log("testArray.3-object3-key7:", object3.GetStringByKey("key7"), testObject.GetStringByKey("key7"), "=字符串7")
	t.Log("--------------------------------------------------- 6 ---------------------------------------------------------")
	//
	testObject.Set("array", testArray)
	t.Log("testObject", testObject.ToJSONString())
	t.Log("testObject-size:", testObject.Size())
	testObjectKeyArray := testObject.GetArrayByKey("array")
	t.Log("testObject.array-size:", testObjectKeyArray.Size(), testArray.Size())
	testObjectKeyArray.Add("添加了一个字符串")
	t.Log("testObject.array-size:", testObjectKeyArray.Size(), testArray.Size())
	//
	t.Log("--------------------------------------------------- 7 ---------------------------------------------------------")
	testArray.Add(test2Object)
	testArrayObject := testArray.GetByIndex(testArray.Size() - 1)
	t.Log("testArray-last-IsObject", testArrayObject.IsObject())
	t.Log("testArray-last-object-size:", testArrayObject.Size(), test2Object.Size(), "=2")
	testArrayObject.JSONObject().Set("addadd", "添加添加")
	t.Log("testArray-last-object-size:", testArrayObject.Size(), test2Object.Size(), "=3")
	t.Log("--------------------------------------------------- 8 ---------------------------------------------------------")
	t.Log("testObject", testObject.ToJSONString())
	t.Log("--------------------------------------------------- 9 ---------------------------------------------------------")
	type Stt struct {
		Name  string
		Age   int
		Items []string
	}
	var stt = &Stt{
		Name:  "TestName",
		Age:   333,
		Items: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	sttJSON := NewJSONObject(stt)
	t.Log("sttJSON", sttJSON.ToJSONString())
	t.Log("--------------------------------------------------- 10 ---------------------------------------------------------")
	var stts = make([]*Stt, 2)
	stts[0] = stt
	stts[1] = stt
	sttsArray := NewJSONArray(stts)
	t.Log("sttsArray", sttsArray.ToJSONString())
	t.Log("--------------------------------------------------- 11 ---------------------------------------------------------")
	var mapTT = make(map[string]interface{})
	mapTT["stt"] = stt
	mapTT["stts"] = stts
	mapTTJSON := NewJSONObject(mapTT)
	t.Log("mapTTJSON", mapTTJSON.ToJSONString())
	t.Log("--------------------------------------------------- 12 ---------------------------------------------------------")
	t.Log("mapTTJSON.stt.Name", mapTTJSON.GetObjectByKey("stt").GetStringByKey("Name"), "=TestName")
	t.Log("mapTTJSON.stt.Age", mapTTJSON.GetObjectByKey("stt").GetIntByKey("Age"), "=333")
	mapTTJSON.GetObjectByKey("stt").Set("Name", "张三")
	t.Log("mapTTJSON.stt.Name", mapTTJSON.GetObjectByKey("stt").GetStringByKey("Name"), "=张三")
	t.Log("--------------------------------------------------- end ---------------------------------------------------------")
}

func TestJSONArray(t *testing.T) {
	var jsonStr = "[\"stringValue\",345,true,23666644.66666666666,\"字符串？\",30344.66,{\"stringField\":\"stringFieldValue\",\"intField\":1000,\"arrayField\":[100,200,\"数组里的字符串\",66996.99],\"doubleField\":999991.102,\"booleanField\":true},[100,200,\"数组里的字符串\",66996.99,{\"stringField\":\"stringFieldValue\",\"intField\":1000,\"arrayField\":[100,200,\"数组里的字符串\",66996.99],\"doubleField\":999991.102,\"booleanField\":true},true,false],8888888889233,\"null\",\"undefined\"]"
	var jsonBytes = []byte(jsonStr)
	t.Log("jsonStr", jsonStr)
	t.Log("jsonBytes", jsonBytes)
	argsJSONString := NewJSONArray(jsonStr)
	t.Log("argsJSONString.1:", argsJSONString.GetByIndex(1).IsInt(), argsJSONString.GetByIndex(1).Int(), "=true =345")
	t.Log("argsJSONString.3:", argsJSONString.GetByIndex(3).IsFloat(), argsJSONString.GetByIndex(3).Int(), argsJSONString.GetByIndex(3).Float(), argsJSONString.GetByIndex(3).String())
	argsJSONBytes := NewJSON(jsonBytes).JSONArray()
	t.Log("argsJSONBytes.1:", argsJSONBytes.GetByIndex(1).IsInt(), argsJSONBytes.GetByIndex(1).Int(), "=true =345")
	t.Log("argsJSONBytes.3:", argsJSONBytes.GetByIndex(3).IsFloat(), argsJSONBytes.GetByIndex(3).Int(), argsJSONBytes.GetByIndex(3).Float(), argsJSONBytes.GetByIndex(3).String())
	t.Log("argsJSONBytes.6:", argsJSONBytes.GetByIndex(6).IsObject(), "=true")
	var mapObject = make(map[string]interface{})
	mapObject["k1"] = "子子"
	mapObject["k2"] = 55555
	mapObject["k3"] = 7777.999
	var arrayObject = []interface{}{"aaaaaa", "bbbbbb", 33333}
	jsonArray := NewJSONArray(nil)
	jsonArray.Add("字符串", 11111, true, mapObject, arrayObject)
	t.Log("jsonArray-size:", jsonArray.Size(), "=5")
	t.Log("jsonArray.1:", jsonArray.GetIntByIndex(1), "=11111")
	t.Log("jsonArray.bytes:", jsonArray.Bytes())
	t.Log("jsonArray.ToJSONString:", jsonArray.ToJSONString())
	mp := jsonArray.GetByIndex(3).JSONObject() //map
	t.Log("jsonArray.3.ToJSONString:", mp.ToJSONString())
	t.Log("jsonArray.3.k1:", mp.GetStringByKey("k1"), "=子子")
	t.Log("jsonArray.3.k2:", mp.GetIntByKey("k2"), "=55555")
	jsonArray.GetByIndex(3).JSONObject().Set("k2", "杨红岩")
	t.Log("jsonArray.3.k2:", mp.GetIntByKey("k2"), "=0")
	t.Log("jsonArray.3.k2:", jsonArray.GetByIndex(3).JSONObject().GetStringByKey("k2"), "=杨红岩")
	index3 := jsonArray.GetByIndex(3)
	t.Log("jsonArray.3.Data:", index3.Data())
	index3.JSONObject().Set("k2", "index.k2=杨红岩")
	t.Log("jsonArray.3.k2:", jsonArray.GetByIndex(3).JSONObject().GetStringByKey("k2"), "=index.k2=杨红岩")
	t.Log("jsonArray.4.1:", jsonArray.GetByIndex(4).JSONArray().GetStringByIndex(1), "=bbbbbb")
	jsonArray.GetByIndex(4).JSONArray().SetByIndex(1, "数组.4变字符串")
	t.Log("jsonArray.4.1:", jsonArray.GetByIndex(4).JSONArray().GetStringByIndex(1), "=数组.4变字符串")
	jsonArray.SetByIndex(4, "jsonArray.4=数组变字符串")
	t.Log("jsonArray.4.type:", jsonArray.GetByIndex(4).Type(), "=")
	t.Log("jsonArray.4:", jsonArray.GetByIndex(4).String(), "=jsonArray.4=数组变字符串")
	t.Log("jsonArray.ToJSONString:", jsonArray.ToJSONString())
	jsonArray.GetByIndex(4).SetValue(111111)
	t.Log("jsonArray.4.type:", jsonArray.GetByIndex(4).Type())
	t.Log("jsonArray.ToJSONString:", jsonArray.ToJSONString())
	jsonArray.GetByIndex(4).SetValue(3333.33)
	t.Log("jsonArray.4.type:", jsonArray.GetByIndex(4).Type())
	t.Log("jsonArray.ToJSONString:", jsonArray.ToJSONString())
}
