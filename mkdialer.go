package subscribemail

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

func MakeGMailDialer(username, passwd, message, subject, target string) {
	host := "smtp.gmail.com"
	port := 25
	mail := gomail.NewMessage()
	mail.SetHeader("From", username)
	mail.SetHeader("To", target)
	mail.SetHeader("Subject", subject)

	mail.SetBody("text/html", message)

	dial := gomail.NewDialer(
		host,
		port,
		username,
		passwd,
	)

	dial.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dial.DialAndSend(mail); err != nil {
		fmt.Println("Error!", err)
	}

}
