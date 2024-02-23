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

package app_test

import (
	"reflect"
	"testing"

	"github.com/and-rad/heyapple/internal/app"
	"github.com/and-rad/heyapple/internal/core"
	"github.com/and-rad/heyapple/internal/mock"
)

func TestCreateRecipe_Execute(t *testing.T) {
	for idx, data := range []struct {
		db   *mock.DB
		name string

		err error
	}{
		{ //00//
			db:  mock.NewDB().WithError(mock.ErrDOS),
			err: mock.ErrDOS,
		},
		{ //01//
			db:   mock.NewDB().WithID(42),
			name: "Apple Pie",
		},
	} {
		cmd := &app.CreateRecipe{Name: data.name}
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if cmd.Recipe.ID != data.db.ID {
			t.Errorf("test case %d: id mismatch \nhave: %v\nwant: %v", idx, cmd.Recipe.ID, data.db.ID)
		}

		if cmd.Name != data.db.Name {
			t.Errorf("test case %d: name mismatch \nhave: %v\nwant: %v", idx, data.db.Name, cmd.Name)
		}
	}
}

func TestSaveRecipe_Execute(t *testing.T) {
	for idx, data := range []struct {
		data map[string]interface{}
		id   int
		db   *mock.DB

		rec core.Recipe
		err error
	}{
		{ //00// id missing or invalid
			db:  mock.NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// connection failed
			id:  1,
			db:  mock.NewDB().WithError(mock.ErrDOS),
			err: mock.ErrDOS,
		},
		{ //02// id not found
			id:  2,
			db:  mock.NewDB().WithRecipe(mock.Recipe1()),
			rec: mock.Recipe1(),
			err: app.ErrNotFound,
		},
		{ //03// success
			id: 1,
			data: map[string]interface{}{
				"name":     "Apple Pie",
				"size":     2,
				"preptime": 5,
				"cooktime": 30,
				"misctime": 45,
			},
			db: mock.NewDB().WithRecipe(mock.Recipe1()),
			rec: func() core.Recipe {
				r := mock.Recipe1()
				r.Name = "Apple Pie"
				r.Size = 2
				r.PrepTime = 5
				r.CookTime = 30
				r.MiscTime = 45
				return r
			}(),
		},
	} {
		cmd := &app.SaveRecipe{ID: data.id, Data: data.data}
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v \nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.rec, data.db.RecipeItem) {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, data.db.RecipeItem, data.rec)
		}
	}
}

func TestRecipeAccess_Fetch(t *testing.T) {
	for idx, data := range []struct {
		db   *mock.DB
		user int
		rec  int

		perms int
		err   error
	}{
		{ //00// id missing or invalid
			db:  mock.NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// connection failed
			db:   mock.NewDB().WithError(mock.ErrDOS),
			user: 1,
			rec:  2,
			err:  mock.ErrDOS,
		},
		{ //02// empty db
			db:    mock.NewDB(),
			user:  1,
			rec:   2,
			perms: app.PermNone,
		},
		{ //03// success
			db:    mock.NewDB().WithAccess(mock.Access{User: 1, Resource: 2, Perms: app.PermCreate | app.PermDelete}),
			user:  1,
			rec:   2,
			perms: app.PermCreate | app.PermDelete,
		},
	} {
		qry := &app.RecipeAccess{UserID: data.user, RecID: data.rec}
		err := qry.Fetch(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if qry.Permission != data.perms {
			t.Errorf("test case %d: permission mismatch \nhave: %v \nwant: %v", idx, qry.Permission, data.perms)
		}
	}
}

func TestSaveRecipeAccess_Execute(t *testing.T) {
	for idx, data := range []struct {
		db    *mock.DB
		user  int
		rec   int
		perms int

		err error
	}{
		{ //00// id missing or invalid
			db:  mock.NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// invalid permissions
			db:    mock.NewDB(),
			user:  1,
			rec:   1,
			perms: app.PermCreate - 1,
			err:   app.ErrPermission,
		},
		{ //02// connection failed
			db:    mock.NewDB().WithError(mock.ErrDOS),
			user:  1,
			rec:   2,
			perms: app.PermCreate,
			err:   mock.ErrDOS,
		},
		{ //03// empty db
			db:    mock.NewDB(),
			user:  1,
			rec:   2,
			perms: app.PermCreate | app.PermEdit,
		},
		{ //04// override existing permissions
			db:    mock.NewDB().WithAccess(mock.Access{User: 1, Resource: 2, Perms: app.PermCreate}),
			user:  1,
			rec:   2,
			perms: app.PermCreate | app.PermDelete,
		},
	} {
		cmd := &app.SaveRecipeAccess{UserID: data.user, RecID: data.rec, Permission: data.perms}
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if err == nil && data.db.Access.Perms != data.perms {
			t.Errorf("test case %d: permission mismatch \nhave: %v \nwant: %v", idx, data.db.Access.Perms, data.perms)
		}
	}
}

func TestRecipeAccess_HasPerms(t *testing.T) {
	for idx, data := range []struct {
		query app.RecipeAccess
		perms int

		ok bool
	}{
		{ //00// empty, default is "none"
			query: app.RecipeAccess{},
			perms: app.PermNone,
			ok:    true,
		},
		{ //01// exact match
			query: app.RecipeAccess{Permission: app.PermDelete},
			perms: app.PermDelete,
			ok:    true,
		},
		{ //02// sub match
			query: app.RecipeAccess{Permission: app.PermOwner},
			perms: app.PermDelete | app.PermEdit,
			ok:    true,
		},
		{ //03// no match
			query: app.RecipeAccess{Permission: app.PermRead},
			perms: app.PermEdit,
			ok:    false,
		},
	} {
		if ok := data.query.HasPerms(data.perms); ok != data.ok {
			t.Errorf("test case %d: result mismatch \nhave: %v\nwant: %v", idx, ok, data.ok)
		}
	}
}

func TestRecipeInstructions_Fetch(t *testing.T) {
	for idx, data := range []struct {
		db  *mock.DB
		rec int

		inst string
		err  error
	}{
		{ //00// id missing or invalid
			db:  mock.NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// connection failed
			db:  mock.NewDB().WithError(mock.ErrDOS),
			rec: 1,
			err: mock.ErrDOS,
		},
		{ //02// success
			db:   mock.NewDB().WithInstructions(1, "Cook it well"),
			rec:  1,
			inst: "Cook it well",
		},
		{ //03// no instructions for this recipe
			db:   mock.NewDB().WithInstructions(2, "Cook it well"),
			rec:  1,
			inst: "",
		},
	} {
		qry := &app.RecipeInstructions{RecID: data.rec}
		err := qry.Fetch(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if qry.Instructions != data.inst {
			t.Errorf("test case %d: text mismatch \nhave: %v \nwant: %v", idx, qry.Instructions, data.inst)
		}
	}
}

func TestSaveRecipeInstructions_Execute(t *testing.T) {
	for idx, data := range []struct {
		db  *mock.DB
		rec int
		in  string

		out string
		err error
	}{
		{ //00// id missing or invalid
			db:  mock.NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// connection failed
			db:  mock.NewDB().WithError(mock.ErrDOS),
			rec: 1,
			err: mock.ErrDOS,
		},
		{ //02// success setting instructions
			db:  mock.NewDB().WithInstructions(1, "Cook it well"),
			rec: 1,
			in:  "Eat it raw.",
		},
		{ //03// success deleting instructions
			db:  mock.NewDB().WithInstructions(2, "Cook it well"),
			rec: 2,
			in:  "",
		},
	} {
		cmd := &app.SaveRecipeInstructions{RecipeID: data.rec, Instructions: data.in}
		err := cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if out := data.db.Instructions.Instructions; out != data.in {
			t.Errorf("test case %d: text mismatch \nhave: %v \nwant: %v", idx, out, data.in)
		}
	}
}

func TestRecipes_Fetch(t *testing.T) {
	for idx, data := range []struct {
		db     *mock.DB
		filter core.Filter
		uid    int

		recs []core.Recipe
		err  error
	}{
		{ //00// connection failed
			db:  mock.NewDB().WithError(mock.ErrDOS),
			uid: 1,
			err: mock.ErrDOS,
		},
		{ //01// user doesn't exist
			db:  mock.NewDB(),
			err: app.ErrNotFound,
		},
		{ //02// empty db
			db:   mock.NewDB(),
			uid:  1,
			recs: []core.Recipe{},
		},
		{ //03// success, no filter
			db:   mock.NewDB().WithRecipes(mock.Recipe1(), mock.Recipe2()),
			uid:  1,
			recs: []core.Recipe{mock.Recipe1(), mock.Recipe2()},
		},
		{ //04// success, with filter
			db:     mock.NewDB().WithRecipes(mock.Recipe1()),
			filter: core.Filter{"kcal": 150},
			uid:    1,
			recs:   []core.Recipe{mock.Recipe1()},
		},
	} {
		qry := &app.Recipes{UserID: data.uid, Filter: data.filter}
		err := qry.Fetch(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(qry.Items, data.recs) {
			t.Errorf("test case %d: data mismatch \nhave: %#v\nwant: %#v", idx, qry.Items, data.recs)
		}

		if !reflect.DeepEqual(data.db.Filter, data.filter) {
			t.Errorf("test case %d: filter mismatch \nhave: %#v\nwant: %#v", idx, data.db.Filter, data.filter)
		}
	}
}

func TestRecipe_Fetch(t *testing.T) {
	for idx, data := range []struct {
		id int
		db *mock.DB

		rec core.Recipe
		err error
	}{
		{ //00// id missing or invalid
			db:  mock.NewDB(),
			err: app.ErrNotFound,
		},
		{ //01// connection failed
			id:  1,
			db:  mock.NewDB().WithError(mock.ErrDOS),
			rec: core.Recipe{ID: 1},
			err: mock.ErrDOS,
		},
		{ //02// empty db
			id:  1,
			db:  mock.NewDB(),
			rec: core.Recipe{ID: 1},
			err: app.ErrNotFound,
		},
		{ //03// id not found
			id:  1,
			db:  mock.NewDB().WithRecipe(mock.Recipe2()),
			rec: core.Recipe{ID: 1},
			err: app.ErrNotFound,
		},
		{ //04// success
			id:  1,
			db:  mock.NewDB().WithRecipe(mock.Recipe1()),
			rec: mock.Recipe1(),
		},
	} {
		qry := &app.Recipe{Item: core.Recipe{ID: data.rec.ID}}
		err := qry.Fetch(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v\nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(qry.Item, data.rec) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, qry.Item, data.rec)
		}
	}
}

func TestSaveIngredient_Execute(t *testing.T) {
	for idx, data := range []struct {
		cmd *app.SaveIngredient
		db  *mock.DB

		rec core.Recipe
		err error
	}{
		{ //00// connection failed
			db:  mock.NewDB().WithError(mock.ErrDOS),
			cmd: &app.SaveIngredient{},
			err: mock.ErrDOS,
		},
		{ //01// recipe not found
			db:  mock.NewDB(),
			cmd: &app.SaveIngredient{RecipeID: 1},
			err: app.ErrNotFound,
		},
		{ //02// remove ingredient success
			db: mock.NewDB().WithRecipe(core.Recipe{
				ID: 1, Name: "rec1", Items: []core.Ingredient{{ID: 2, Amount: 120}}, KCal: 123, Fat: 43},
			),
			cmd: &app.SaveIngredient{RecipeID: 1, IngredientID: 2, Amount: 0},
			rec: core.Recipe{ID: 1, Name: "rec1", Items: []core.Ingredient{}},
		},
		{ //03// ingredient doesn't exist
			db:  mock.NewDB().WithRecipe(mock.Recipe1()).WithFood(mock.Food2),
			cmd: &app.SaveIngredient{RecipeID: 1, IngredientID: 3, Amount: 350},
			rec: mock.Recipe1(),
		},
		{ //04// replace an ingredient's amount
			db:  mock.NewDB().WithRecipe(mock.Recipe1()).WithFoods(mock.Food2),
			cmd: &app.SaveIngredient{RecipeID: 1, IngredientID: 2, Amount: 200, Replace: true},
			rec: func() core.Recipe {
				r := mock.Recipe1()
				r.Items = []core.Ingredient{{ID: mock.Food2.ID, Amount: 200}}
				r.KCal = mock.Food2.KCal * 2
				r.Fat = mock.Food2.Fat * 2
				r.Carbs = mock.Food2.Carbs * 2
				r.Sugar = mock.Food2.Sugar * 2
				r.Fructose = mock.Food2.Fructose * 2
				r.Glucose = mock.Food2.Glucose * 2
				r.Sucrose = mock.Food2.Sucrose * 2
				r.Protein = mock.Food2.Protein * 2
				r.Fiber = mock.Food2.Fiber * 2
				r.Iron = mock.Food2.Iron * 2
				r.Zinc = mock.Food2.Zinc * 2
				r.Magnesium = mock.Food2.Magnesium * 2
				r.Chlorine = mock.Food2.Chlorine * 2
				r.Sodium = mock.Food2.Sodium * 2
				r.Calcium = mock.Food2.Calcium * 2
				r.Potassium = mock.Food2.Potassium * 2
				r.Phosphorus = mock.Food2.Phosphorus * 2
				r.Copper = mock.Food2.Copper * 2
				r.Iodine = mock.Food2.Iodine * 2
				r.Manganse = mock.Food2.Manganse * 2
				r.VitA = mock.Food2.VitA * 2
				r.VitB1 = mock.Food2.VitB1 * 2
				r.VitB2 = mock.Food2.VitB2 * 2
				r.VitB6 = mock.Food2.VitB6 * 2
				r.VitC = mock.Food2.VitC * 2
				r.VitE = mock.Food2.VitE * 2
				return r
			}(),
		},
		{ //05// add to an ingredient's amount
			db:  mock.NewDB().WithRecipe(mock.Recipe1()).WithFoods(mock.Food2),
			cmd: &app.SaveIngredient{RecipeID: 1, IngredientID: 2, Amount: 123},
			rec: func() core.Recipe {
				r := mock.Recipe1()
				r.Items = []core.Ingredient{{ID: mock.Food2.ID, Amount: 273}}
				r.KCal = mock.Food2.KCal * 2.73
				r.Fat = mock.Food2.Fat * 2.73
				r.Carbs = mock.Food2.Carbs * 2.73
				r.Sugar = mock.Food2.Sugar * 2.73
				r.Fructose = mock.Food2.Fructose * 2.73
				r.Glucose = mock.Food2.Glucose * 2.73
				r.Sucrose = mock.Food2.Sucrose * 2.73
				r.Protein = mock.Food2.Protein * 2.73
				r.Fiber = mock.Food2.Fiber * 2.73
				r.Iron = mock.Food2.Iron * 2.73
				r.Zinc = mock.Food2.Zinc * 2.73
				r.Magnesium = mock.Food2.Magnesium * 2.73
				r.Chlorine = mock.Food2.Chlorine * 2.73
				r.Sodium = mock.Food2.Sodium * 2.73
				r.Calcium = mock.Food2.Calcium * 2.73
				r.Potassium = mock.Food2.Potassium * 2.73
				r.Phosphorus = mock.Food2.Phosphorus * 2.73
				r.Copper = mock.Food2.Copper * 2.73
				r.Iodine = mock.Food2.Iodine * 2.73
				r.Manganse = mock.Food2.Manganse * 2.73
				r.VitA = mock.Food2.VitA * 2.73
				r.VitB1 = mock.Food2.VitB1 * 2.73
				r.VitB2 = mock.Food2.VitB2 * 2.73
				r.VitB6 = mock.Food2.VitB6 * 2.73
				r.VitC = mock.Food2.VitC * 2.73
				r.VitE = mock.Food2.VitE * 2.73
				return r
			}(),
		},
		{ //06// add an ingredient
			db:  mock.NewDB().WithRecipe(mock.Recipe1()).WithFoods(mock.Food1, mock.Food2),
			cmd: &app.SaveIngredient{RecipeID: 1, IngredientID: 1, Amount: 100},
			rec: func() core.Recipe {
				r := mock.Recipe1()
				r.Items = append(r.Items, core.Ingredient{ID: mock.Food1.ID, Amount: 100})
				r.KCal += mock.Food1.KCal
				r.Fat += mock.Food1.Fat
				r.Carbs += mock.Food1.Carbs
				r.Sugar += mock.Food1.Sugar
				r.Fructose += mock.Food1.Fructose
				r.Glucose += mock.Food1.Glucose
				r.Sucrose += mock.Food1.Sucrose
				r.Protein += mock.Food1.Protein
				r.Fiber += mock.Food1.Fiber
				r.Iron += mock.Food1.Iron
				r.Zinc += mock.Food1.Zinc
				r.Magnesium += mock.Food1.Magnesium
				r.Chlorine += mock.Food1.Chlorine
				r.Sodium += mock.Food1.Sodium
				r.Calcium += mock.Food1.Calcium
				r.Potassium += mock.Food1.Potassium
				r.Phosphorus += mock.Food1.Phosphorus
				r.Copper += mock.Food1.Copper
				r.Iodine += mock.Food1.Iodine
				r.Manganse += mock.Food1.Manganse
				r.VitA += mock.Food1.VitA
				r.VitB1 += mock.Food1.VitB1
				r.VitB2 += mock.Food1.VitB2
				r.VitB6 += mock.Food1.VitB6
				r.VitC += mock.Food1.VitC
				r.VitE += mock.Food1.VitE
				return r
			}(),
		},
		{ //07// handle contradicting flags
			db:  mock.NewDB().WithRecipe(mock.Recipe1()).WithFoods(mock.Food2, core.Food{ID: 12, Flags: core.FlagBeef}),
			cmd: &app.SaveIngredient{RecipeID: 1, IngredientID: 12, Amount: 100},
			rec: func() core.Recipe {
				r := mock.Recipe1()
				r.Items = append(r.Items, core.Ingredient{ID: 12, Amount: 100})
				r.Flags = core.FlagBeef
				return r
			}(),
		},
		{ //08// handle contradicting flags
			db:  mock.NewDB().WithRecipe(mock.Recipe1()).WithFoods(mock.Food2, core.Food{ID: 12, Flags: core.FlagDairy}),
			cmd: &app.SaveIngredient{RecipeID: 1, IngredientID: 12, Amount: 100},
			rec: func() core.Recipe {
				r := mock.Recipe1()
				r.Items = append(r.Items, core.Ingredient{ID: 12, Amount: 100})
				r.Flags = core.FlagVegetarian | core.FlagDairy
				return r
			}(),
		},
	} {
		err := data.cmd.Execute(data.db)

		if err != data.err {
			t.Errorf("test case %d: error mismatch \nhave: %v \nwant: %v", idx, err, data.err)
		}

		if !reflect.DeepEqual(data.rec, data.db.RecipeItem) {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, data.db.RecipeItem, data.rec)
		}
	}
}
