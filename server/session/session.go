package session

import (
	"go-reserve/common/db"
	"go-reserve/models"
)

var _db = db.GetDB()

type manage struct{}

var _ models.SessionManage = (*manage)(nil)

func NewManager() *manage { return &manage{} }

func (m *manage) ListSession() ([]*models.Session, error) {
	var sessions []*models.Session
	if err := _db.Find(&sessions).Error; err != nil {
		return nil, err
	}

	return sessions, nil
}

func (m *manage) GetSession(id string) (*models.Session, error) {
	var session *models.Session
	if err := _db.Find(&session, id).Error; err != nil {
		return nil, err
	}

	return session, nil
}

func (m *manage) DeleteSession(id string) error {
	return _db.Delete(&models.Session{}, id).Error
}

func (m *manage) CreateSession(session *models.Session) error {
	return _db.Create(&session).Error
}

func (m *manage) SendMessage(message *models.Message) error { return nil }
