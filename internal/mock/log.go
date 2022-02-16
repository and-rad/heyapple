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

package mock

type Log struct {
	Message string
	Warning string
	Err     string
}

func NewLog() *Log {
	return &Log{}
}

func (l *Log) Log(i interface{}) {
	switch v := i.(type) {
	case string:
		l.Message = v
	case error:
		l.Message = v.Error()
	}
}

func (l *Log) Warn(i interface{}) {
	switch v := i.(type) {
	case string:
		l.Warning = v
	case error:
		l.Warning = v.Error()
	}
}

func (l *Log) Error(i interface{}) {
	switch v := i.(type) {
	case string:
		l.Err = v
	case error:
		l.Err = v.Error()
	}
}
