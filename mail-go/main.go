package main

import (
	"fmt"
	"log"
	"mail-go/internal/config"
	"mail-go/models"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

var app config.AppConfig

func main() {
	fmt.Println("Hello, World!")

	mailChan := make(chan models.MailData) // Creo el canal
	app.MailChan = mailChan

	listenForMail()

	// app.MailChan <- msg
	// mailChan <- msg // Envio 'msg' al channel 'mailChan'

	defer close(app.MailChan) // Postergar hasta que se cierre main()
	listenForMail()

	msg := models.MailData{
		To:      "chehin238@gmail.com",
		From:    "sgmtucuman",
		Subject: "Asunto - Golang",
		Content: "",
	}

	app.MailChan <- msg
}

func listenForMail() {
	// mailChan := make(chan models.MailData) // Creo el canal
	// Funcion asincronica (se ejecuta en segundo plano)
	// Bucle infinito en para levantar el server
	go func() {
		for {
			msg := <-app.MailChan // Recibo desde 'mailChan' y asigno a 'msg'
			// msg := chan models.MailData
			// msg := make(chan models.MailData)
			fmt.Println("msg ..", msg)
			sendMsg(msg)
		}
	}()
}

func sendMsg(m models.MailData) {

	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		// errorLog.Println(err)
		log.Println(err)
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
