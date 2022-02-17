package email

import (
	"errors"
	"heyapple/pkg/app"
	"net/smtp"
	"net/textproto"

	"github.com/jordan-wright/email"
)

// Error definitions
var (
	ErrNoTemplate      = errors.New("no template")
	ErrParseTemplate   = errors.New("parse")
	ErrExecuteTemplate = errors.New("execute")
)

// Notifier provides messaging & notification capabilities via e-mail.
type Notifier struct{}

// NewNotifier returns a default-initialized Notifier.
func NewNotifier() *Notifier {
	return &Notifier{}
}

func (n *Notifier) Send(to string, msg app.Notification, data interface{}) error {
	conf := getConfig()
	mail := &email.Email{
		To:      []string{to},
		From:    conf.from(),
		Subject: "subject",
		HTML:    []byte(""),
		Headers: textproto.MIMEHeader{},
	}

	return n.send(mail, conf.server(), conf.auth())
}

func (n *Notifier) send(e *email.Email, server string, auth smtp.Auth) error {
	return nil //e.Send(server, auth)
}
