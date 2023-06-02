package main

import (
	"fmt"
	"net/smtp"
)

func main() {

	// Sender data.
	from := "macy.mohr@ethereal.email"
	password := "A1YBH3vjZmbeRKhd9J"

	// Receiver email address.
	to := []string{
		"sender@example.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.ethereal.email"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

// https://ethereal.email/create
// https://ethereal.email/messages/64748ee3564f802bb62d6b31/2?tab=header
