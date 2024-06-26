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

package mem

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/core"
)

const (
	backupFile  = "db"
	storageFile = "db"
)

type access struct {
	Resource int `json:"res"`
	Perms    int `json:"perms"`
}

type backup struct {
	db  *DB
	log app.Logger

	lastBackup string
}

type backupData struct {
	Users        map[int]app.User          `json:"users"`
	Prefs        map[int]app.StoredPrefs   `json:"prefs"`
	Tokens       map[string]app.Token      `json:"tokens"`
	Food         map[int]core.Food         `json:"food"`
	Recipes      map[int]core.Recipe       `json:"recipes"`
	Instructions map[int]string            `json:"inst"`
	UserRec      map[int][]access          `json:"recaccess"`
	Entries      map[int][]core.DiaryEntry `json:"entries"`
	Aisles       aisleMap                  `json:"aisles"`
	Prices       priceMap                  `json:"prices"`
	Done         doneMap                   `json:"done"`
	UserID       int                       `json:"userid"`
	FoodID       int                       `json:"foodid"`
	RecID        int                       `json:"recid"`
}

func (b *backup) Run() {
	if err := b.save(); err != nil {
		b.log.Error(err)
	}

	if now := time.Now().Format("2006-01-02"); now != b.lastBackup {
		if err := b.backUp(); err != nil {
			b.log.Error(err)
		}
	}
}

func (b *backup) load() {
	b.db.mtx.Lock()
	defer b.db.mtx.Unlock()

	path := filepath.Join(getConfig().storageDir, storageFile+".json")
	if data, err := os.ReadFile(path); err == nil {
		var db backupData
		if err = json.Unmarshal(data, &db); err == nil {
			b.db.userID = db.UserID
			b.db.users = db.Users
			b.db.prefs = db.Prefs
			b.db.tokens = db.Tokens
			b.db.food = db.Food
			b.db.foodID = db.FoodID
			b.db.recipes = db.Recipes
			b.db.instructions = db.Instructions
			b.db.recID = db.RecID
			b.db.aisles = db.Aisles
			b.db.prices = db.Prices
			b.db.done = db.Done

			for k, v := range b.db.users {
				b.db.emails[v.Email] = k
			}

			for k, v := range db.UserRec {
				b.db.userRec[k] = map[int]int{}
				for _, a := range v {
					b.db.userRec[k][a.Resource] = a.Perms
					if _, ok := b.db.recUser[a.Resource]; !ok {
						b.db.recUser[a.Resource] = map[int]int{}
					}
					b.db.recUser[a.Resource][k] = a.Perms
				}
			}

			for k, v := range db.Entries {
				b.db.entries[k] = map[time.Time][]core.DiaryEntry{}
				b.db.days[k] = map[int]map[int][]core.DiaryDay{}
				for _, e := range v {
					day := e.Date.Truncate(time.Hour * 24)
					if _, ok := b.db.entries[k][day]; !ok {
						b.db.entries[k][day] = []core.DiaryEntry{}
					}
					b.db.entries[k][day] = append(b.db.entries[k][day], e)

					year := e.Date.Year()
					if _, ok := b.db.days[k][year]; !ok {
						b.db.days[k][year] = map[int][]core.DiaryDay{}
					}

					month := int(e.Date.Month())
					if _, ok := b.db.days[k][year][month]; !ok {
						b.db.days[k][year][month] = []core.DiaryDay{}
					}

					date := e.Date.Format("2006-01-02")
					amount := e.Food.Amount * 0.01
					food := b.db.food[e.Food.ID]

					exists := false
					for i := range b.db.days[k][year][month] {
						day := &b.db.days[k][year][month][i]
						if day.Date == date {
							exists = true
							day.KCal += amount * food.KCal
							day.Fat += amount * food.Fat
							day.FatSat += amount * food.FatSat
							day.FatO3 += amount * food.FatO3
							day.FatO6 += amount * food.FatO6
							day.Carbs += amount * food.Carbs
							day.Sugar += amount * food.Sugar
							day.Fructose += amount * food.Fructose
							day.Glucose += amount * food.Glucose
							day.Sucrose += amount * food.Sucrose
							day.Fiber += amount * food.Fiber
							day.Protein += amount * food.Protein
							day.Salt += amount * food.Salt

							day.Potassium += amount * food.Potassium
							day.Chlorine += amount * food.Chlorine
							day.Sodium += amount * food.Sodium
							day.Calcium += amount * food.Calcium
							day.Phosphorus += amount * food.Phosphorus
							day.Magnesium += amount * food.Magnesium
							day.Iron += amount * food.Iron
							day.Zinc += amount * food.Zinc
							day.Manganse += amount * food.Manganse
							day.Copper += amount * food.Copper
							day.Iodine += amount * food.Iodine
							day.Chromium += amount * food.Chromium
							day.Molybdenum += amount * food.Molybdenum
							day.Selenium += amount * food.Selenium

							day.VitA += amount * food.VitA
							day.VitB1 += amount * food.VitB1
							day.VitB2 += amount * food.VitB2
							day.VitB3 += amount * food.VitB3
							day.VitB5 += amount * food.VitB5
							day.VitB6 += amount * food.VitB6
							day.VitB7 += amount * food.VitB7
							day.VitB9 += amount * food.VitB9
							day.VitB12 += amount * food.VitB12
							day.VitC += amount * food.VitC
							day.VitD += amount * food.VitD
							day.VitE += amount * food.VitE
							day.VitK += amount * food.VitK
							break
						}
					}

					if !exists {
						b.db.days[k][year][month] = append(b.db.days[k][year][month], core.DiaryDay{
							Date:     date,
							KCal:     amount * food.KCal,
							Fat:      amount * food.Fat,
							FatSat:   amount * food.FatSat,
							FatO3:    amount * food.FatO3,
							FatO6:    amount * food.FatO6,
							Carbs:    amount * food.Carbs,
							Sugar:    amount * food.Sugar,
							Fructose: amount * food.Fructose,
							Glucose:  amount * food.Glucose,
							Sucrose:  amount * food.Sucrose,
							Fiber:    amount * food.Fiber,
							Protein:  amount * food.Protein,
							Salt:     amount * food.Salt,

							Potassium:  amount * food.Potassium,
							Chlorine:   amount * food.Chlorine,
							Sodium:     amount * food.Sodium,
							Calcium:    amount * food.Calcium,
							Phosphorus: amount * food.Phosphorus,
							Magnesium:  amount * food.Magnesium,
							Iron:       amount * food.Iron,
							Zinc:       amount * food.Zinc,
							Manganse:   amount * food.Manganse,
							Copper:     amount * food.Copper,
							Iodine:     amount * food.Iodine,
							Chromium:   amount * food.Chromium,
							Molybdenum: amount * food.Molybdenum,
							Selenium:   amount * food.Selenium,

							VitA:   amount * food.VitA,
							VitB1:  amount * food.VitB1,
							VitB2:  amount * food.VitB2,
							VitB3:  amount * food.VitB3,
							VitB5:  amount * food.VitB5,
							VitB6:  amount * food.VitB6,
							VitB7:  amount * food.VitB7,
							VitB9:  amount * food.VitB9,
							VitB12: amount * food.VitB12,
							VitC:   amount * food.VitC,
							VitD:   amount * food.VitD,
							VitE:   amount * food.VitE,
							VitK:   amount * food.VitK,
						})
					}
				}
			}
		}
	}
}

func (b *backup) save() error {
	b.db.mtx.RLock()
	defer b.db.mtx.RUnlock()

	dir := getConfig().storageDir
	if err := os.MkdirAll(dir, 0750); err != nil {
		return err
	}

	path := filepath.Join(dir, storageFile+".json")
	return os.WriteFile(path, b.bytes(), 0644)
}

func (b *backup) backUp() error {
	b.db.mtx.RLock()
	defer b.db.mtx.RUnlock()

	dir := getConfig().backupDir
	if err := os.MkdirAll(dir, 0750); err != nil {
		return err
	}

	now := time.Now().Format("2006-01-02")
	path := filepath.Join(dir, backupFile+now+".json")
	err := os.WriteFile(path, b.bytes(), 0644)

	if err == nil {
		b.lastBackup = now
	}

	return err
}

func (b *backup) bytes() []byte {
	recAccess := map[int][]access{}
	for k, v := range b.db.userRec {
		for r, p := range v {
			recAccess[k] = append(recAccess[k], access{Resource: r, Perms: p})
		}
	}

	entries := map[int][]core.DiaryEntry{}
	for id, days := range b.db.entries {
		for _, list := range days {
			entries[id] = append(entries[id], list...)
		}
		sort.Slice(entries[id], func(i, j int) bool {
			return entries[id][i].Date.Before(entries[id][j].Date)
		})
	}

	data, _ := json.Marshal(backupData{
		Users:        b.db.users,
		Prefs:        b.db.prefs,
		UserID:       b.db.userID,
		Food:         b.db.food,
		FoodID:       b.db.foodID,
		Recipes:      b.db.recipes,
		Instructions: b.db.instructions,
		Entries:      entries,
		Aisles:       b.db.aisles,
		Prices:       b.db.prices,
		Done:         b.db.done,
		RecID:        b.db.recID,
		Tokens:       b.db.tokens,
		UserRec:      recAccess,
	})

	return data
}
