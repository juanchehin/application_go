package main

import (
	"fmt"
	"mail-go/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func main() {
	fmt.Println("Hello, World!")

	mailChain := make(chan models.MailData)
}

func listenForMail() {
	// Funcion asincronica (se ejecuta en segundo plano)
	go func() {
		for {
			// msg := <- app.MailChan
			// msg := chan models.MailData
			msg := make(chan models.MailData)
			sendMsg(msg)
		}
	}
}

func sendMsg(m models.MailData) {
	server : mail.NewSMTOClient()
	server.Hoost = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).
		AddTo(m.To).
		AddCc(m.Subject).
		SetSubject("New Go Email")

	email.SetBody(mail.TextHTML, "Hello, <strong> World!</strong>!")

	// Call Send and pass the client
	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent")
	}
}