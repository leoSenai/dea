package utils

import (
	"api/models"
	"net/smtp"
)

func SendEmail(to, subject, body string) error {
	config := models.EmailConfig{
		SMTPServer:   "smtp.gmail.com",
		SMTPPort:     "587",
		SMTPUsername: "suportedeanaoresponda@gmail.com",
		SMTPPassword: "xhyd nrfx vjnc erth",
	}
	from := config.SMTPUsername
	message := "Subject: " + subject + "\r\n" +
		"From: " + from + "\r\n" +
		"To: " + to + "\r\n\r\n" +
		body

	auth := smtp.PlainAuth("", config.SMTPUsername, config.SMTPPassword, config.SMTPServer)
	err := smtp.SendMail(config.SMTPServer+":"+config.SMTPPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
