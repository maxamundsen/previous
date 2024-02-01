package snailmail

import (
	"gohttp/config"
	"net/smtp"
	"strings"
	"log"
)

type Email struct {
	Recipients []string
	Subject string
	Body string
}

func SendPlaintextMail(message Email) error {
	config := config.GetConfiguration()

	smtpHost := config.SmtpServer
	smtpPort := config.SmtpPort

	from := config.SmtpUsername
	password := config.SmtpPassword

	recipientString := strings.Join(message.Recipients, ",")

	smtpMessage := []byte("From: " + from + "\r\n" + "To: " + recipientString + "\r\n" + "Subject: " + message.Subject + "\r\n" + message.Body)

	var auth smtp.Auth = nil

	if config.SmtpRequireAuth {
		auth = smtp.PlainAuth("", from, password, smtpHost)
	}

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, message.Recipients, smtpMessage)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}