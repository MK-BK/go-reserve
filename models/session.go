package models

import "gorm.io/gorm"

type SessionManage interface {
	List() ([]*Session, error)
	SendMessage(message *Message) error

	Create(session *Session) error
	Get(id string) (*SessionInfo, error)
	Delete(id string) error
}

type Session struct {
	gorm.Model
	Name   string
	ShopID string `json:"shopID" binding:"required"`
	UserID string `json:"userID" binding:"required"`
}

type SessionInfo struct {
	Session
	Message []Message
}
