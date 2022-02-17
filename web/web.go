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
)

var (
	App   = template.Must(template.ParseFS(Templates, "templates/web/app.html"))
	Home  = template.Must(template.ParseFS(Templates, "templates/web/home.html"))
	Login = template.Must(template.ParseFS(Templates, "templates/web/login.html"))
)
