package main

import (
	"fmt"
	"net/smtp"
)

type smtpServer struct {
	host string
	port string

}

func(s *smtpServer) serverName() string {
	return s.host + ":" + s.port
}

func main(){
	from := "sender email address"
	password := "password"
	to := "receiver email address"
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "507"}
	message := []byte ("Enter the message you want to send")
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("email sent!")
}