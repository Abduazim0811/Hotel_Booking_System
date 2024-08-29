package email

import (
	"log"

	"gopkg.in/gomail.v2"
)

func SendEmail(email, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "yusupovabduazim0811@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Registration Status")
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "yusupovabduazim0811@gmail.com", "slws nzfk namk eali")

	if err := d.DialAndSend(m); err != nil {
		log.Println("failed to send an email:", err)
		return err
	}

	return nil
}
