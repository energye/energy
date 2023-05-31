package src

import "fmt"

var ObjDemoVar *ObjDemo

type ObjDemo struct {
	NameField           string
	ObjDemoStringField  string
	ObjDemoBoolField    bool
	ObjDemoIntField     int32
	ObjDemoFloat64Field float64
	SubObjDemoField     *SubObjDemo
}

type SubObjDemo struct {
	SubString string
	SubBool   bool
}

func (m *ObjDemo) CallbackCount() {
	fmt.Println("CallbackCount")
}

func (m *SubObjDemo) SubObjFn1(aaa int32) {
	fmt.Println("SubObjFn1", aaa)
}
