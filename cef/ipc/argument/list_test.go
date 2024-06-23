package argument

import (
	"testing"
)

func TestList(t *testing.T) {
	list := &List{
		Data: make([]string, 0),
	}
	json := list.JSON().JSONArray()
	json.Add("value1")
	json.Add("value2")
	json.Add("value3")
	if json.Size() != 3 {
		t.Fail()
	}
	if json.GetStringByIndex(1) != "value2" {
		t.Fail()
	}
	data := json.Bytes()
	if data == nil {
		t.Fail()
	}
}
