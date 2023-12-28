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

package app

import (
	"fmt"
	"io"
)

type Logger interface {
	Log(interface{})
	Warn(interface{})
	Error(interface{})
}

type Log struct {
	out io.Writer
}

func NewLog(wr io.Writer) *Log {
	return &Log{
		out: wr,
	}
}

func (l *Log) Log(i interface{}) {
	l.send(i, "Log")
}

func (l *Log) Warn(i interface{}) {
	l.send(i, "Warn")
}

func (l *Log) Error(i interface{}) {
	l.send(i, "Error")
}

func (l *Log) send(data interface{}, level string) {
	switch v := data.(type) {
	case string:
		fmt.Fprintf(l.out, "%-5s: %s\n", level, v)
	case error:
		fmt.Fprintf(l.out, "%-5s: %s\n", level, v.Error())
	default:
		fmt.Fprintf(l.out, "%-5s: %v\n", level, v)
	}
}
