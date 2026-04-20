//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package types

const DefaultActionIdentifier = "DEFAULT_ACTION"

type TNotificationResponseEvent func(result Result)

// Response 通知响应
type Response struct {
	ID               string                 `json:"id,omitempty"`
	ActionIdentifier string                 `json:"actionIdentifier,omitempty"`
	CategoryID       string                 `json:"categoryIdentifier,omitempty"`
	Title            string                 `json:"title,omitempty"`
	Subtitle         string                 `json:"subtitle,omitempty"`
	Body             string                 `json:"body,omitempty"`
	UserText         string                 `json:"userText,omitempty"`
	UserInfo         map[string]interface{} `json:"userInfo,omitempty"`
}

// Result 通知结果
type Result struct {
	Response Response
	Error    error
}

// Options 通知配置
type Options struct {
	ID         string                 `json:"id"`
	Title      string                 `json:"title"`
	Subtitle   string                 `json:"subtitle,omitempty"`
	Body       string                 `json:"body,omitempty"`
	CategoryID string                 `json:"categoryId,omitempty"`
	Data       map[string]interface{} `json:"data,omitempty"`
}

// Category 通知类别
type Category struct {
	ID               string   `json:"id,omitempty"`
	Actions          []Action `json:"actions,omitempty"`
	HasReplyField    bool     `json:"hasReplyField,omitempty"`
	ReplyPlaceholder string   `json:"replyPlaceholder,omitempty"`
	ReplyButtonTitle string   `json:"replyButtonTitle,omitempty"`
}

// Action 通知操作按钮
type Action struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Destructive bool   `json:"destructive,omitempty"`
}

type INotification interface {
	// 初始化
	Initialize() error

	// 核心通知方法
	RequestNotificationAuthorization() (bool, error)
	CheckNotificationAuthorization() (bool, error)
	SendNotification(options Options) error
	SendNotificationWithActions(options Options) error

	// 类别管理
	RegisterNotificationCategory(category Category) error
	RemoveNotificationCategory(categoryID string) error

	// 通知管理
	RemoveAllPendingNotifications() error
	RemovePendingNotification(identifier string) error
	RemoveAllDeliveredNotifications() error
	RemoveDeliveredNotification(identifier string) error
	RemoveNotification(identifier string) error
}

type INotificationDarwin interface {
	SetOnNotificationResponse(callback TNotificationResponseEvent)
}
