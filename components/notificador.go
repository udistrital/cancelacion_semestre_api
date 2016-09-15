package components

import (
	"bytes"
	"log"
	"net/smtp"
)

type Notificador struct {
	From string
	To   string
}

func (n Notificador) EnviarPlano() {
  // Considerar https://godoc.org/github.com/astaxie/beego/utils#Email
  // Se está usando https://github.com/golang/go/wiki/SendingMail
	// Connect to the remote SMTP server.
	c, err := smtp.Dial("smtp.mail.yahoo.com:465")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	// Set the sender and recipient.
	c.Mail(n.From)
	c.Rcpt(n.To)
	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString("This is the email body.")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}

func (n Notificador) Enviar() {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"jorgenator2@yahoo.es",
		"",
		"smtp.mail.yahoo.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.mail.yahoo.com:465",
		auth,
		n.From,
		[]string{n.To},
		[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}
