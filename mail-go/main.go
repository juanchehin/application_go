package main

import (
	"fmt"
	"log"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

func main() {

	fmt.Println("pasa -1")
	listenForMail()

}

func listenForMail() {

	server := mail.NewSMTPClient()

	fmt.Println("pasa 0")

	server.Host = "smtp.gmail.com"
	server.Port = 587

	server.Username = "sgmtucuman@gmail.com"
	server.Password = ""
	server.Encryption = mail.EncryptionSTARTTLS

	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	// server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	fmt.Println("pasa")
	// SMTP client
	smtpClient, err := server.Connect()

	if err != nil {
		log.Fatal(err)
	}

	email := mail.NewMSG()
	email.SetFrom("From Example <nube@example.com>").
		AddTo("chehin238@gmail.com").
		SetSubject("New Go Email")

	email.SetBody(mail.TextHTML, "Hello, <strong> World!</strong>!")

	// always check error after send
	if email.Error != nil {
		log.Fatal(email.Error)
	}

	// Call Send and pass the client
	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent")
	}

}
