////////////////////////////////////////////////////////////////////////
//
// Copyright (C) 2021-2024 The HeyApple Authors.
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

package mock

import (
	"net/smtp"

	"github.com/and-rad/heyapple/internal/app"

	"github.com/jordan-wright/email"
)

type Notifier struct {
	To   string
	Msg  app.Notification
	Data app.NotificationData

	Err error
}

func NewNotifier() *Notifier {
	return &Notifier{}
}

func (n *Notifier) WithError(err error) *Notifier {
	n.Err = err
	return n
}

func (n *Notifier) Send(to string, msg app.Notification, data app.NotificationData) error {
	if n.Err != nil {
		return n.Err
	}

	n.To = to
	n.Msg = msg
	n.Data = data

	return nil
}

type EmailNotifier struct {
	Email  *email.Email
	Server string

	Err error
}

func NewEmailNotifier() *EmailNotifier {
	return &EmailNotifier{}
}

func (e *EmailNotifier) WithError(err error) *EmailNotifier {
	e.Err = err
	return e
}

func (e *EmailNotifier) SendFunc(em *email.Email, server string, auth smtp.Auth) error {
	if e.Err != nil {
		return e.Err
	}
	e.Email = em
	e.Server = server
	return nil
}
