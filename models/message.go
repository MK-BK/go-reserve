package models

type Message struct {
	ID        string
	SessionID string
	Content   Content `gorm:"json"`
}

type Content interface {
	Marshal(value interface{}) ([]byte, error)
	Unmarshal(bytes []byte, value interface{}) error
}

type ContentText string

type ContentIcon uint64
