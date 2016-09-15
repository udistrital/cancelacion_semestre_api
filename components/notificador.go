package components

import (
	"gopkg.in/gomail.v2"
)

type Notificador struct {
	From string
	To   string
}

func (n Notificador) Enviar() {
  m := gomail.NewMessage()
    m.SetAddressHeader("From", n.From, "testMail1")
    m.SetHeader("To",
        m.FormatAddress(n.To, "testMail1"),
        //m.FormatAddress("testmail2@yahoo.com", "testMail2"),
    )
    m.SetHeader("Subject", "mail")
    m.SetBody("text/plain", "Hello!")

    d := gomail.NewPlainDialer("smtp.mail.yahoo.com", 25, "jorgenator2@yahoo.es", "password")
    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }
}
