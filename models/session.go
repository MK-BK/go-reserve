package models

type SessionManage interface {
	CreateSession(session *Session) error
	ListSession() ([]*Session, error)
	GetSession(id string) (*Session, error)
	DeleteSession(id string) error
	SendMessage(message *Message) error
}

type Session struct {
	ID   uint
	Name string
}
