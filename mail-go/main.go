package main

import (
	"fmt"
	"models"
)

func main() {
	fmt.Println("Hello, World!")

	mailChain := make(chan models.MailData)
}

func listenForMail() {
	// Funcion asincronica (se ejecuta en segundo plano)
	go func() {
		for {
			msg := <- 
		}
	}
}

func sendMsg() {
	
}