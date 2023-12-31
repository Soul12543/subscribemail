package subscribemail

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func MakeGMailDialer(username, passwd, subject, message, attachment, target string) error {
	host := "smtp.gmail.com"
	port := 25
	mail := gomail.NewMessage()
	mail.SetHeader("From", username)
	mail.SetHeader("To", target)
	mail.SetHeader("Subject", subject)

	mail.SetBody("text/html", message)

	if attachment != "" {
		mail.Attach(attachment)
	}

	dial := gomail.NewDialer(
		host,
		port,
		username,
		passwd,
	)

	dial.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dial.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}
