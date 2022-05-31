package server

import (
	"go-reserve/models"
)

type SessionManaer struct{}

var _ models.SessionManage = (*SessionManaer)(nil)

func NewSessionManager() *SessionManaer { return &SessionManaer{} }

func (m *SessionManaer) List() ([]*models.Session, error) {
	var sessions []*models.Session
	if err := _db.Find(&sessions).Error; err != nil {
		return nil, err
	}

	return sessions, nil
}

func (m *SessionManaer) Get(id string) (*models.SessionInfo, error) {
	var (
		messages []models.Message
		session  *models.Session
	)

	if err := _db.Find(&session, id).Error; err != nil {
		return nil, err
	}

	if err := _db.Where("session_id = ?", id).Find(&messages).Error; err != nil {
		return nil, err
	}

	return &models.SessionInfo{
		Session: *session,
		Message: messages,
	}, nil
}

func (m *SessionManaer) Create(session *models.Session) error {
	return _db.Create(&session).Error
}

func (m *SessionManaer) Delete(id string) error {
	return _db.Delete(&models.Session{}, id).Error
}

func (m *SessionManaer) SendMessage(message *models.Message) error {
	return _db.Create(message).Error
}
