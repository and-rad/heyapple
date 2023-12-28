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

package email

import (
	"net/smtp"
	"os"
	"reflect"
	"testing"
)

type auth struct {
	identity, username, password string
	host                         string
}

func Test_getConfig(t *testing.T) {
	for idx, data := range []struct {
		env map[string]string
		cfg config
	}{
		{ //00// empty environment, default values
			cfg: config{
				domain:   "http://localhost:8080",
				addr:     "user@example.com",
				fromAddr: "user@example.com",
				fromName: "User",
				host:     "example.com",
				pass:     "topsecret",
				port:     587,
			},
		},
		{ //01// ignore other vars
			env: map[string]string{
				"PATH":      "/usr/bin",
				envFromName: "Dude Dudeson",
			},
			cfg: config{
				domain:   "http://localhost:8080",
				addr:     "user@example.com",
				fromAddr: "user@example.com",
				fromName: "Dude Dudeson",
				host:     "example.com",
				pass:     "topsecret",
				port:     587,
			},
		},
		{ //02// invalid type
			env: map[string]string{
				envPort: "Port",
			},
			cfg: config{
				domain:   "http://localhost:8080",
				addr:     "user@example.com",
				fromAddr: "user@example.com",
				fromName: "User",
				host:     "example.com",
				pass:     "topsecret",
				port:     587,
			},
		},
		{ //03// load all vars
			env: map[string]string{
				envDomain:   "http://here",
				envAddr:     "a@a.a",
				envFromAddr: "b@b.b",
				envFromName: "Slickback",
				envHost:     "a.com",
				envPass:     "secret",
				envPort:     "22",
			},
			cfg: config{
				domain:   "http://here",
				addr:     "a@a.a",
				fromAddr: "b@b.b",
				fromName: "Slickback",
				host:     "a.com",
				pass:     "secret",
				port:     22,
			},
		},
	} {
		os.Clearenv()
		for k, v := range data.env {
			os.Setenv(k, v)
			defer os.Unsetenv(k)
		}

		cfg := getConfig()

		if cfg != data.cfg {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, cfg, data.cfg)
		}
	}
}

func Test_config_server(t *testing.T) {
	for idx, data := range []struct {
		cfg config
		out string
	}{
		{ //00// empty config
			cfg: config{},
			out: ":0",
		},
		{ //01// default config
			cfg: getConfig(),
			out: "example.com:587",
		},
		{ //02// custom config
			cfg: config{host: "a.com", port: 22},
			out: "a.com:22",
		},
	} {
		if out := data.cfg.server(); out != data.out {
			t.Errorf("test case %d: result mismatch \nhave: %v\nwant: %v", idx, out, data.out)
		}
	}
}

func Test_config_auth(t *testing.T) {
	for idx, data := range []struct {
		cfg  config
		auth smtp.Auth
	}{
		{ //00// empty config
			cfg:  config{},
			auth: smtp.PlainAuth("", "", "", ""),
		},
		{ //01// default config
			cfg:  getConfig(),
			auth: smtp.PlainAuth("", "user@example.com", "topsecret", "example.com"),
		},
		{ //02// custom config
			cfg:  config{host: "a.com", addr: "a@a.a", pass: "nope"},
			auth: smtp.PlainAuth("", "a@a.a", "nope", "a.com"),
		},
	} {
		auth := data.cfg.auth()

		if !reflect.DeepEqual(auth, data.auth) {
			t.Errorf("test case %d: result mismatch \nhave: %#v\nwant: %#v", idx, auth, data.auth)
		}
	}
}

func Test_config_from(t *testing.T) {
	for idx, data := range []struct {
		cfg config
		out string
	}{
		{ //00// empty config
			cfg: config{},
			out: " <>",
		},
		{ //01// default config
			cfg: getConfig(),
			out: "User <user@example.com>",
		},
		{ //02// custom config
			cfg: config{fromAddr: "a@a.a", fromName: "A"},
			out: "A <a@a.a>",
		},
	} {
		if out := data.cfg.from(); out != data.out {
			t.Errorf("test case %d: result mismatch \nhave: %v\nwant: %v", idx, out, data.out)
		}
	}
}
