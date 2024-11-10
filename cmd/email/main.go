package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"gopkg.in/gomail.v2"
)

func main() {
	fmt.Println(SendEmail("ssg0808@yandex.ru", "test", "aboba"))
	ReadEmails("imap.yandex.ru:993", "ssg0808@yandex.ru", "prvhxbylrwrmgnxo")
}

func SendEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "ssg0808@yandex.ru")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.yandex.ru", 587, "ssg0808@yandex.ru", "prvhxbylrwrmgnxo")

	return d.DialAndSend(m)
}

func ReadEmails(server, email, password string) {
	// Connect to the IMAP server
	c, err := client.DialTLS(server, &tls.Config{ServerName: "imap.yandex.ru"})
	if err != nil {
		log.Fatalf("dial error: %v", err)
	}
	defer c.Logout()

	// Login
	if err := c.Login(email, password); err != nil {
		log.Fatalf("login error: %v", err)
	}

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatalf("select error: %v", err)
	}

	// Get the last 10 messages
	seqset := new(imap.SeqSet)
	seqset.AddNum(uint32(mbox.Messages))

	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *imap.Message, 1)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, items, messages)
	}()

	for msg := range messages {
		if msg == nil {
			log.Println("Server didn't return message")
			continue
		}

		// Process the message
		r := msg.GetBody(section)
		if r == nil {
			log.Println("Server didn't return message body")
			continue
		}

		mr, err := mail.CreateReader(r)
		if err != nil {
			log.Println("Failed to create message reader:", err)
			continue
		}

		header := mr.Header
		if date, err := header.Date(); err == nil {
			log.Println("Date:", date)
		}
		if from, err := header.AddressList("From"); err == nil {
			log.Println("From:", from)
		}
		if subject, err := header.Subject(); err == nil {
			log.Println("Subject:", subject)
		}

		// Process each message part
		for {
			p, err := mr.NextPart()
			if err != nil {
				break
			}

			switch p.Header.(type) {
			case *mail.InlineHeader:
				// This is the message's text (can be plain-text or HTML)
				b, _ := io.ReadAll(p.Body)
				log.Println("Body:", string(b))
			}
		}
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}
}
