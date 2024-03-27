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
	"math/rand"
	"strconv"
	"strings"
)

// ChangeName is a command to assign a new user name to the
// user identified by ID. The name is randomly generated.
// If successful, Name contains the new name.
type ChangeName struct {
	ID   int
	Name string
}

func (c *ChangeName) Execute(db DB) error {
	if c.ID == 0 {
		return ErrNotFound
	}

	user, err := db.UserByID(c.ID)
	if err != nil {
		return err
	}

	name := userNameAdjective() + userNameNoun()
	names, err := db.UserNames(name)
	if err != nil {
		return err
	}

	if len(names) > 0 {
		next := 1
		for _, n := range names {
			suffix, _ := strings.CutPrefix(n, name)
			if suffix == "" {
				continue
			}

			num, err := strconv.Atoi(suffix)
			if err != nil {
				continue
			}

			next = max(next, num+1)
		}
		name += strconv.Itoa(next)
	}

	user.Name = name
	if err = db.SetUser(user); err == nil {
		c.Name = name
	}

	return err
}

// TODO: add more names
var (
	adjectives = []string{
		"Adamant", "Angry", "Annoying", "Artful",
		"Beautiful", "Bespectacled", "Boneheaded", "Bumpy",
		"Cantankerous", "Charming", "Crosseyed", "Cunning",
	}
	nouns = []string{
		"Apple", "Apricot", "Aubergine", "Avocado",
		"Bacon", "Baguette", "Banana", "Blueberry",
		"Carrot", "Cheese", "Chicken", "Croissant",
	}
)

func userNameAdjective() string {
	return adjectives[rand.Intn(len(adjectives))]
}

func userNameNoun() string {
	return nouns[rand.Intn(len(nouns))]
}
