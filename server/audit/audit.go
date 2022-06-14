package audit

func ListLogs() ([]*AuditLog, error) {
	var logs []*AuditLog
	if err := _db.Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
