package authentication

import (
	"os"

	"gopkg.in/gomail.v2"
)

var emailCodeChannel = make(chan string)

func SendGoMail(from string, to string, subject string, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, from, os.Getenv("APP_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
