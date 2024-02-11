package snailmail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"webdawgengine/config"
	"log"
	"mime"
	"net/mail"
	"net/smtp"
	"strings"
	"time"
)

const (
	TYPE_TEXT = iota
	TYPE_HTML = iota
)

type Email struct {
	Recipients []string
	Subject    string
	Body       *bytes.Buffer
}

func SendMail(message Email, mailtype int) error {
	recipientString := strings.Join(message.Recipients, ",")
	from := mail.Address{config.GetSmtpDisplayFrom(), config.GetSmtpUsername()}

	header := make(map[string]string)
	header["To"] = recipientString
	header["From"] = from.String()
	header["Subject"] = mime.QEncoding.Encode("UTF-8", message.Subject)
	header["MIME-Version"] = "1.0"
	header["Content-Transfer-Encoding"] = "base64"
	header["Date"] = time.Now().Format(time.RFC1123)

	if mailtype == TYPE_HTML {
		header["Content-Type"] = "text/html; charset=\"utf-8\""
	} else {
		header["Content-Type"] = "text/plain; charset=\"utf-8\""
	}

	email := ""
	for k, v := range header {
		email += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	email += "\r\n" + base64.StdEncoding.EncodeToString(message.Body.Bytes())

	var auth smtp.Auth = nil

	if config.GetSmtpRequireAuth() {
		auth = smtp.PlainAuth("", config.GetSmtpUsername(), config.GetSmtpPassword(), config.GetSmtpServer())
	}

	err := smtp.SendMail(config.GetSmtpServer() + ":" + config.GetSmtpPort(), auth, config.GetSmtpUsername(), message.Recipients, []byte(email))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
