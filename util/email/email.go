package mail

import (
	"net/smtp"
)

const (
	smtpHost = "smtp.gmail.com"
	smtpPort = "587"

	from     = "from@gmail.com"
	password = "<Email Password>"
)

func NewEmail(to []string) error {
	auth := smtp.PlainAuth("", from, password, smtpHost)
	message := []byte("This is a test email message.")

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}

	return nil
}
