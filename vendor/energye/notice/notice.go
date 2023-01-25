package notice

import (
	"strconv"
	"time"
)

var uniqueID string

// Notification 表示可发送到操作系统的用户通知
type Notification struct {
	Title, Content string
	Icon           string
	Timeout        int32
}

// NewNotification 创建一个通知，可以传递给SendNotification
func NewNotification(title, content string) *Notification {
	return &Notification{Title: title, Content: content}
}

func UniqueID() string {
	if uniqueID != "" {
		return uniqueID
	}
	uniqueID = "missing-id-" + strconv.FormatInt(time.Now().Unix(), 10) // This is a fake unique - it just has to not be reused...
	return uniqueID
}
