package email

import (
	"errors"
	"heyapple/pkg/api/v1"
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
type Notifier struct {
	tr api.Translator
}

// NewNotifier returns a default-initialized Notifier.
func NewNotifier(tr api.Translator) *Notifier {
	return &Notifier{
		tr: tr,
	}
}

func (n *Notifier) Send(to string, msg app.Notification, data app.NotificationData) error {
	conf := getConfig()
	mail := &email.Email{
		To:      []string{to},
		From:    conf.from(),
		Subject: n.subject(msg, data),
		HTML:    n.message(msg, data),
		Headers: textproto.MIMEHeader{},
	}

	return n.send(mail, conf.server(), conf.auth())
}

func (n *Notifier) subject(msg app.Notification, data app.NotificationData) string {
	return ""
}

func (n *Notifier) message(msg app.Notification, data app.NotificationData) []byte {
	return []byte("")
}

func (n *Notifier) send(e *email.Email, server string, auth smtp.Auth) error {
	return nil //e.Send(server, auth)
}
