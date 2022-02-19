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

package app

// CreateRecipe is a command to create a new recipe in the
// food database. If successful, the new item id is stored
// in the command.
type CreateRecipe struct {
	ID int
}

func (c *CreateRecipe) Execute(db DB) error {
	if id, err := db.NewRecipe(); err != nil {
		return err
	} else {
		c.ID = id
	}

	return nil
}
