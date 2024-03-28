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
	ID   int    `json:"id"`
	Name string `json:"name"`
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

// We're trying to go for 10 entries per letter, which would
// give us about 67,000 unique combinations. Combined with
// number suffixes, this is easily enough for any number of
// users the app will ever realistically have.
var (
	adjectives = []string{
		"Able", "Adventurous", "Adorable", "Adroit", "Adaptable", "Annoying", "Auspicious", "Artful", "Adamant", "Angry",
		"Brave", "Bright", "Beautiful", "Bountiful", "Blissful", "Bad", "Brilliant", "Benevolent", "Bespectacled", "Bumpy",
		"Calm", "Creative", "Clever", "Cheerful", "Curious", "Courageous", "Charming", "Cantankerous", "Crosseyed", "Cunning",
		"Daring", "Dazzling", "Delightful", "Dynamic", "Determined", "Diligent",
		"Eager", "Elegant", "Enthusiastic", "Energetic", "Efficient", "Eloquent",
		"Fabulous", "Friendly", "Funny", "Fierce", "Flexible",
		"Gentle", "Generous", "Graceful", "Genuine", "Glorious",
		"Happy", "Helpful", "Honest", "Humble", "Healthy",
		"Innovative", "Intelligent", "Impressive", "Ingenious", "Inspiring",
		"Jolly", "Joyful", "Jubilant", "Just", "Jovial",
		"Kind", "Keen", "Knowledgeable", "Kooky", "Kaleidoscopic",
		"Lively", "Loyal", "Luminous", "Lucky", "Lovely",
		"Magnificent", "Majestic", "Mindful", "Magical", "Merry", "Marvelous",
		"Nice", "Noble", "Neat", "Nurturing", "Natural",
		"Optimistic", "Outstanding", "Optimal", "Original", "Observant",
		"Peaceful", "Passionate", "Patient", "Polite", "Positive",
		"Quick", "Quiet", "Quirky", "Qualified", "Quaint", "Quixotic",
		"Radiant", "Resilient", "Resourceful", "Reliable", "Remarkable",
		"Sincere", "Spirited", "Strong", "Stellar", "Sensible",
		"Talented", "Thoughtful", "Trustworthy", "Thrilling", "Tenacious",
		"Unique", "Understanding", "Upbeat", "Unwavering", "Uplifting",
		"Vibrant", "Valiant", "Vivacious", "Virtuous", "Versatile",
		"Witty", "Warm", "Wise", "Wonderful", "Welcoming",
		"Xenial", "Xenodochial", "Xylophonic", "Xenogenetic",
		"Youthful", "Yummy", "Yearning", "Yare",
		"Zesty", "Zealous", "Zany", "Zippy", "Zestful",
	}
	nouns = []string{
		"Apple", "Apricot", "Avocado", "Artichoke", "Asparagus", "Almond",
		"Banana", "Blueberry", "Blackberry", "Broccoli", "Bean", "Barley", "Beet",
		"Carrot", "Cucumber", "Cantaloupe", "Celery", "Cherry", "Cabbage", "Coconut", "Cauliflower",
		"Date", "Dragonfruit", "Durian", "Daikon", "Dill",
		"Eggplant", "Endive", "Escarole", "Edamame", "Enoki",
		"Fig", "Fennel", "Flax", "Feijoa",
		"Grape", "Guava", "Garlic", "Ginger", "Gooseberry",
		"Honeydew", "Huckleberry", "Honey",
		"Imbe",
		"Jackfruit",
		"Kiwi", "Kale", "Kumquat",
		"Lemon", "Lime", "Lettuce", "Lychee", "Leek",
		"Mango", "Melon", "Mulberry", "Mushroom", "Mustard",
		"Nectarine",
		"Orange", "Onion", "Okra", "Oats", "Olive",
		"Papaya", "Pear", "Pineapple", "Peach", "Pea", "Potato", "Plum",
		"Quince", "Quinoa",
		"Raspberry", "Rhubarb", "Radish",
		"Strawberry", "Spinach", "Squash", "Starfruit",
		"Tomato", "Tangerine", "Turnip", "Tamarillo", "Taro",
		"Ulluco", "Ugni",
		"Vanilla", "Vidalia",
		"Watermelon", "Wasabi",
		"Xigua", "Xylocarp", "Xylopia", "Xylitol",
		"Yam", "Yautia", "Yuzu", "Yacon",
		"Zucchini", "Zebu", "Ziziphus", "Zuiki",
	}
)

func userNameAdjective() string {
	return adjectives[rand.Intn(len(adjectives))]
}

func userNameNoun() string {
	return nouns[rand.Intn(len(nouns))]
}
