////////////////////////////////////////////////////////////////////////
//
// Copyright (C) 2021-2022 The HeyApple Authors.
//
// Use of this source code is governed by the GNU Affero General
// Public License as published by the Free Software Foundation,
// either version 3 of the License, or any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.
//
////////////////////////////////////////////////////////////////////////

// Package email provides an implementation of the app.Notifier interface
// that sends notifications over standard e-mail.
package email

import (
	"bytes"
	"fmt"
	"heyapple/pkg/app"
	"heyapple/web"
	"html/template"
	"net/smtp"
	"net/textproto"

	"github.com/jordan-wright/email"
)

// Notifier provides messaging & notification capabilities via e-mail.
type Notifier struct {
	sendFunc func(e *email.Email, server string, auth smtp.Auth) error
	tr       app.Translator
}

// NewNotifier returns a default-initialized Notifier.
func NewNotifier(tr app.Translator) *Notifier {
	return &Notifier{
		sendFunc: send,
		tr:       tr,
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

	return n.sendFunc(mail, conf.server(), conf.auth())
}

func (n *Notifier) subject(msg app.Notification, data app.NotificationData) string {
	lang, ok := data["lang"].(string)
	if !ok {
		lang = n.tr.Default()
	}

	return n.tr.Translate(fmt.Sprintf("email.sub.%d", msg), lang)
}

func (n *Notifier) message(msg app.Notification, data app.NotificationData) []byte {
	lang, ok := data["lang"].(string)
	if !ok {
		lang = n.tr.Default()
	}

	var tpl *template.Template
	switch msg {
	case app.RegisterNotification:
		tpl = web.MailRegister
	case app.RenameNotification:
		tpl = web.MailRename
	case app.ResetNotification:
		tpl = web.MailReset
	}

	var buf bytes.Buffer
	if err := template.Must(tpl.Clone()).Funcs(template.FuncMap{
		"l10n": func(in interface{}) string { return n.tr.Translate(in, lang) },
	}).Execute(&buf, data); err != nil {
		return []byte("")
	}

	return buf.Bytes()
}

func send(e *email.Email, server string, auth smtp.Auth) error {
	return e.Send(server, auth)
}
