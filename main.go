package main

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func main() {
	message := `
	<p>Hello World<p>
	`

	host := "smtp.gmail.com"
	port := 25
	username := "zhuyuhan163@gmail.com"
	password := "cegzwvlkjfoelunc"

	mail := gomail.NewMessage()
	mail.SetHeader("From", username)
	mail.SetHeader("To", "2016573858@qq.com")
	mail.SetHeader("Subject", "Test")

	mail.SetBody("text/html", message)

	dial := gomail.NewDialer(
		host,
		port,
		username,
		password,
	)

	dial.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dial.DialAndSend(mail); err != nil {
		panic(err)
	}

}
