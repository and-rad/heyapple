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

package web

import (
	"embed"
	"html/template"
)

var (
	//go:embed templates
	Templates embed.FS

	//go:embed static
	Static embed.FS

	//go:embed l10n
	L10n embed.FS
)

var (
	App     = template.Must(template.New("").Funcs(funcs).ParseFS(Templates, "templates/web/partial/*.html", "templates/web/app.html"))
	Confirm = template.Must(template.New("").Funcs(funcs).ParseFS(Templates, "templates/web/partial/*.html", "templates/web/confirm.html"))
	Home    = template.Must(template.New("").Funcs(funcs).ParseFS(Templates, "templates/web/partial/*.html", "templates/web/home.html"))
	Login   = template.Must(template.New("").Funcs(funcs).ParseFS(Templates, "templates/web/partial/*.html", "templates/web/login.html"))
)

var (
	MailRegister = template.Must(template.New("base").Funcs(funcs).ParseFS(Templates, "templates/email/base.html", "templates/email/register.html"))
	MailRename   = template.Must(template.New("base").Funcs(funcs).ParseFS(Templates, "templates/email/base.html", "templates/email/rename.html"))
	MailReset    = template.Must(template.New("base").Funcs(funcs).ParseFS(Templates, "templates/email/base.html", "templates/email/reset.html"))
)

var funcs = template.FuncMap{
	"all":  func(i ...interface{}) []interface{} { return i },
	"l10n": func(interface{}) string { return "" },
	"raw":  func(s string) template.HTML { return template.HTML(s) },
}
