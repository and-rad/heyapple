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

package conv

import (
	"encoding/json"

	"github.com/and-rad/heyapple/internal/core"
)

type usdaNutrientData struct {
	Name string `json:"name"`
	Unit string `json:"unitName"`
	ID   int    `json:"id"`
}

type usdaNutrient struct {
	Data   usdaNutrientData `json:"nutrient"`
	Amount float32          `json:"amount"`
}

type usdaFood struct {
	Desc      string         `json:"description"`
	Nutrients []usdaNutrient `json:"foodNutrients"`
}

type usda struct {
	Foods []usdaFood `json:"SurveyFoods"`
}

func FromUSDA(data []byte) ([]Food, error) {
	usda := usda{}
	if err := json.Unmarshal(data, &usda); err != nil {
		return nil, ErrFromJSON
	}

	foods := []Food{}
	for i, inFood := range usda.Foods {
		outFood := Food{
			Name: inFood.Desc,
			Food: core.Food{
				ID: i + 1,
			},
		}

		for _, n := range inFood.Nutrients {
			switch n.Data.ID {
			case 1003:
				outFood.Protein = n.Amount
			}
		}

		foods = append(foods, outFood)
	}

	return foods, nil
}
