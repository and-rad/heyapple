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

package app_test

import (
	"heyapple/internal/mock"
	"heyapple/pkg/app"
	"heyapple/pkg/core"
	"reflect"
	"testing"
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
			db:  mock.NewDB().WithRecipe(mock.Recipe1),
			rec: mock.Recipe1,
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
			db: mock.NewDB().WithRecipe(mock.Recipe1),
			rec: func() core.Recipe {
				r := mock.Recipe1
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
