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

package email

import (
	"errors"
	"heyapple/internal/mock"
	"heyapple/pkg/app"
	"heyapple/web"
	"html/template"
	"net"
	"net/smtp"
	"net/textproto"
	"reflect"
	"testing"

	"github.com/jordan-wright/email"
)

var funcs = template.FuncMap{
	"l10n": func(interface{}) string { return "" },
}

func TestNotifier_Send(t *testing.T) {
	tmpRegister := web.MailRegister

	for idx, data := range []struct {
		to     string
		msg    app.Notification
		sender *mock.EmailNotifier

		email *email.Email
		err   error
	}{
		{ //00// force error on send
			to:     "a@a.a",
			msg:    app.RegisterNotification,
			sender: mock.NewEmailNotifier().WithError(mock.ErrDOS),
			err:    mock.ErrDOS,
		},
		{ //01// success
			to:     "a@a.a",
			msg:    app.RegisterNotification,
			sender: mock.NewEmailNotifier(),
			email: &email.Email{
				To:      []string{"a@a.a"},
				From:    getConfig().from(),
				Subject: "email.sub.1",
				HTML:    []byte("Hello from http://localhost:8080"),
				Headers: textproto.MIMEHeader{},
			},
		},
	} {
		web.MailRegister, _ = template.New("base").Funcs(funcs).Parse("Hello from {{ .domain }}")

		notifier := NewNotifier(mock.NewTranslator())
		notifier.sendFunc = data.sender.SendFunc

		err := notifier.Send(data.to, data.msg, nil)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.sender.Email, data.email) {
			t.Errorf("test case %d: email mismatch \nhave: %v\nwant: %v", idx, data.sender.Email, data.email)
		}
	}

	web.MailRegister = tmpRegister
}

func TestNotifier_subject(t *testing.T) {
	for idx, data := range []struct {
		msg  app.Notification
		data app.NotificationData
		tr   *mock.Translator

		out string
	}{
		{ //00// empty translation, no data
			tr:  mock.NewTranslator(),
			out: "email.sub.0",
		},
		{ //01// key found
			tr: &mock.Translator{Map: map[string]string{
				"email.sub.1": "Subject 1",
			}},
			msg:  app.RegisterNotification,
			data: app.NotificationData{"lang": "en"},
			out:  "Subject 1",
		},
	} {
		out := NewNotifier(data.tr).subject(data.msg, data.data)

		if out != data.out {
			t.Errorf("test case %d: subject mismatch \nhave: %v\nwant: %v", idx, out, data.out)
		}
	}
}

func TestNotifier_message(t *testing.T) {
	tmpRegister := web.MailRegister
	tmpRename := web.MailRename
	tmpReset := web.MailReset

	for idx, data := range []struct {
		msg  app.Notification
		data app.NotificationData
		tr   *mock.Translator
		tmpl string

		out string
	}{
		{ //00// invalid template
			msg:  app.RegisterNotification,
			tr:   mock.NewTranslator(),
			tmpl: `{{ .Foo "Bar" }}`,
			out:  "",
		},
		{ //01// data is passed to template
			msg:  app.RegisterNotification,
			data: app.NotificationData{"lang": "en", "token": "abcd"},
			tr:   mock.NewTranslator(),
			tmpl: `{{ .lang }} {{ .token }}`,
			out:  "en abcd",
		},
		{ //02// translate content
			msg:  app.RegisterNotification,
			tr:   &mock.Translator{Map: map[string]string{"hi": "Hello"}},
			tmpl: `{{ l10n "hi" }}`,
			out:  "Hello",
		},
		{ //03// notification types
			msg:  app.RenameNotification,
			tr:   mock.NewTranslator(),
			tmpl: `Rename`,
			out:  "Rename",
		},
		{ //04// notification types
			msg:  app.ResetNotification,
			tr:   mock.NewTranslator(),
			tmpl: `Reset`,
			out:  "Reset",
		},
	} {
		switch data.msg {
		case app.RegisterNotification:
			web.MailRegister, _ = template.New("base").Funcs(funcs).Parse(data.tmpl)
		case app.RenameNotification:
			web.MailRename, _ = template.New("base").Funcs(funcs).Parse(data.tmpl)
		case app.ResetNotification:
			web.MailReset, _ = template.New("base").Funcs(funcs).Parse(data.tmpl)
		}

		out := string(NewNotifier(data.tr).message(data.msg, data.data))

		if out != data.out {
			t.Errorf("test case %d: subject mismatch \nhave: %v\nwant: %v", idx, out, data.out)
		}
	}

	web.MailRegister = tmpRegister
	web.MailRename = tmpRename
	web.MailReset = tmpReset
}

func Test_send(t *testing.T) {
	for idx, data := range []struct {
		email  *email.Email
		server string
		auth   smtp.Auth
		err    error
	}{
		{ //00// missing From & To
			email: &email.Email{},
			err:   errors.New(""),
		},
		{ //01//
			email: &email.Email{From: "user@example.com", To: []string{"user@example.com"}},
			err:   &net.OpError{},
		},
	} {
		err := send(data.email, data.server, data.auth)

		if reflect.TypeOf(err) != reflect.TypeOf(data.err) {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, reflect.TypeOf(err), reflect.TypeOf(data.err))
		}
	}
}
