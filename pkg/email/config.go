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
	"fmt"
	"net/smtp"
	"os"
	"strconv"
)

const (
	envAddr     = "HEYAPPLE_EMAIL_ADDR"
	envFromAddr = "HEYAPPLE_EMAIL_FROMADDR"
	envFromName = "HEYAPPLE_EMAIL_FROMNAME"
	envHost     = "HEYAPPLE_EMAIL_HOST"
	envPass     = "HEYAPPLE_EMAIL_PASS"
	envPort     = "HEYAPPLE_EMAIL_PORT"
	envDomain   = "HEYAPPLE_APP_DOMAIN"
)

type config struct {
	domain   string
	addr     string
	fromAddr string
	fromName string
	host     string
	pass     string
	port     int
}

func getConfig() config {
	cfg := config{
		domain:   "http://localhost:8080",
		addr:     "user@example.com",
		fromAddr: "user@example.com",
		fromName: "User",
		host:     "example.com",
		pass:     "topsecret",
		port:     587,
	}

	if domain := os.Getenv(envDomain); domain != "" {
		cfg.domain = domain
	}
	if addr := os.Getenv(envAddr); addr != "" {
		cfg.addr = addr
	}
	if fromAddr := os.Getenv(envFromAddr); fromAddr != "" {
		cfg.fromAddr = fromAddr
	}
	if fromName := os.Getenv(envFromName); fromName != "" {
		cfg.fromName = fromName
	}
	if host := os.Getenv(envHost); host != "" {
		cfg.host = host
	}
	if pass := os.Getenv(envPass); pass != "" {
		cfg.pass = pass
	}
	if port, err := strconv.Atoi((os.Getenv(envPort))); err == nil {
		cfg.port = port
	}

	return cfg
}

func (c config) server() string {
	return fmt.Sprintf("%s:%v", c.host, c.port)
}

func (c config) auth() smtp.Auth {
	return smtp.PlainAuth("", c.addr, c.pass, c.host)
}

func (c config) from() string {
	return fmt.Sprintf("%s <%s>", c.fromName, c.fromAddr)
}
