package ipc

type ProcessMessage struct {
	Name string      `json:"n"`
	Data interface{} `json:"d"`
	Id   int         `json:"i"`
}
