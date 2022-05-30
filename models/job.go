package models

type JobManage interface {
	Create(*Job) error
	Get(id string) (*Job, error)
	List() ([]*Job, error)
	Update(id string, job *Job) error
	Delete(id string) error
}

type Job struct {
	ID        uint
	UserID    string
	ProductID string
	Status    string
}
