package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"path/filepath"
)

const (
	smtpHost = "smtp.qq.com"
	smtpPort = "587"

	from     = "827301519@qq.com"
	password = "rbzfaqfuqfiwbecj"
)

func NewEmail(to []string) error {
	auth := smtp.PlainAuth("", from, password, smtpHost)

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	t, err := template.ParseFiles(filepath.Join(dir, "util", "email", "email.html"))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, struct {
		Subject   string
		Name      string
		PageTitle string
	}{
		Name:      "test",
		PageTitle: "title",
	}); err != nil {
		fmt.Println("Parse Templatet error:", err)
		return err
	}

	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, buf.Bytes()); err != nil {
		return err
	}

	return nil
}
