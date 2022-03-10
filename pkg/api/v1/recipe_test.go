package api_test

import (
	"fmt"
	"heyapple/internal/mock"
	"heyapple/pkg/api/v1"
	"heyapple/pkg/app"
	"heyapple/pkg/core"
	"heyapple/pkg/handler"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/and-rad/scs/v2"
	"github.com/julienschmidt/httprouter"
)

func TestNewRecipe(t *testing.T) {
	for idx, data := range []struct {
		db        *mock.DB
		in        url.Values
		setCookie bool

		access mock.Access
		out    string
		status int
	}{
		{ //00// missing input
			db:     mock.NewDB(),
			status: http.StatusBadRequest,
		},
		{ //01// no session
			in:     url.Values{"name": {"My Recipe"}},
			db:     mock.NewDB(),
			status: http.StatusUnauthorized,
		},
		{ //02// connection failure
			in:        url.Values{"name": {"My Recipe"}},
			db:        mock.NewDB().WithUser(mock.User1).WithError(mock.ErrDOS),
			setCookie: true,
			status:    http.StatusInternalServerError,
		},
		{ //03// partial success
			in:        url.Values{"name": {"My Recipe"}},
			db:        mock.NewDB().WithUser(mock.User1).WithID(mock.Recipe0.ID).WithError(nil, mock.ErrDOS),
			setCookie: true,
			status:    http.StatusAccepted,
			out:       strings.Replace(mock.Recipe0Json, `"name":""`, `"name":"My Recipe"`, 1),
		},
		{ //04// success
			in:        url.Values{"name": {"My Recipe"}},
			db:        mock.NewDB().WithUser(mock.User1).WithID(mock.Recipe0.ID),
			setCookie: true,
			status:    http.StatusCreated,
			access:    mock.Access{User: mock.User1.ID, Resource: mock.Recipe0.ID, Perms: app.PermOwner},
			out:       strings.Replace(mock.Recipe0Json, `"name":""`, `"name":"My Recipe"`, 1),
		},
	} {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data.in.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		res := httptest.NewRecorder()
		env := &handler.Environment{DB: data.db, Session: scs.New()}

		if data.setCookie {
			if ctx, err := env.Session.Load(req.Context(), "abc"); err == nil {
				req = req.WithContext(ctx)
				env.Session.Put(req.Context(), "id", data.db.User.ID)
			}
		}

		api.NewRecipe(env)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, body, data.out)
		}

		if acc := data.db.Access; acc != data.access {
			t.Errorf("test case %d: permission mismatch \nhave: %v \nwant: %v", idx, acc, data.access)
		}
	}
}

func TestSaveRecipe(t *testing.T) {
	for idx, data := range []struct {
		db        *mock.DB
		params    httprouter.Params
		in        url.Values
		setCookie bool

		rec    core.Recipe
		status int
	}{
		{ //00// missing mandatory form values
			db:     mock.NewDB(),
			status: http.StatusBadRequest,
		},
		{ //01// missing mandatory form values
			db:     mock.NewDB(),
			params: httprouter.Params{{Key: "name", Value: "12"}},
			status: http.StatusBadRequest,
		},
		{ //02// wrong data type for id
			db:     mock.NewDB(),
			params: httprouter.Params{{Key: "id", Value: "someone"}},
			status: http.StatusBadRequest,
		},
		{ //03// no session
			db:     mock.NewDB(),
			params: httprouter.Params{{Key: "id", Value: "42"}},
			status: http.StatusUnauthorized,
		},
		{ //04// connection failure
			db: mock.NewDB().
				WithUser(mock.User1).
				WithAccess(mock.Access{User: 1, Resource: 42, Perms: app.PermEdit}).
				WithError(nil, mock.ErrDOS),
			params:    httprouter.Params{{Key: "id", Value: "42"}},
			setCookie: true,
			status:    http.StatusInternalServerError,
		},
		{ //05// item doesn't exist
			db: mock.NewDB().
				WithUser(mock.User1).
				WithRecipe(mock.Recipe1).
				WithAccess(mock.Access{User: 1, Resource: 42, Perms: app.PermEdit}),
			params:    httprouter.Params{{Key: "id", Value: "42"}},
			in:        url.Values{"size": {"9"}},
			setCookie: true,
			status:    http.StatusNotFound,
			rec:       mock.Recipe1,
		},
		{ //06// success
			db: mock.NewDB().
				WithUser(mock.User1).
				WithRecipe(mock.Recipe1).
				WithAccess(mock.Access{User: 1, Resource: 1, Perms: app.PermEdit}),
			params:    httprouter.Params{{Key: "id", Value: "1"}},
			in:        url.Values{"size": {"9"}, "name": {"Banana Pie"}},
			setCookie: true,
			status:    http.StatusNoContent,
			rec: func() core.Recipe {
				r := mock.Recipe1
				r.Size = 9
				r.Name = "Banana Pie"
				return r
			}(),
		},
		{ //07// success
			db: mock.NewDB().
				WithUser(mock.User1).
				WithRecipe(mock.Recipe1).
				WithFoods(mock.Food1, mock.Food2).
				WithAccess(mock.Access{User: 1, Resource: 1, Perms: app.PermEdit}),
			params:    httprouter.Params{{Key: "id", Value: "1"}},
			in:        url.Values{"preptime": {"4"}, "cooktime": {"5"}, "misctime": {"6"}},
			setCookie: true,
			status:    http.StatusNoContent,
			rec: func() core.Recipe {
				r := mock.Recipe1
				r.PrepTime = 4
				r.CookTime = 5
				r.MiscTime = 6
				return r
			}()},
	} {
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(data.in.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		res := httptest.NewRecorder()
		env := &handler.Environment{DB: data.db, Session: scs.New()}

		if data.setCookie {
			if ctx, err := env.Session.Load(req.Context(), "abc"); err == nil {
				req = req.WithContext(ctx)
				env.Session.Put(req.Context(), "id", data.db.User.ID)
			}
		}

		api.SaveRecipe(env)(res, req, data.params)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if !reflect.DeepEqual(data.db.RecipeItem, data.rec) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.RecipeItem, data.rec)
		}
	}
}

func TestRecipes(t *testing.T) {
	for idx, data := range []struct {
		db        *mock.DB
		setCookie bool

		out    string
		status int
	}{
		{ //00// connection failure
			db:        mock.NewDB().WithError(mock.ErrDOS),
			status:    http.StatusInternalServerError,
			setCookie: true,
		},
		{ //01// anonymous user
			db:     mock.NewDB(),
			status: http.StatusInternalServerError,
		},
		{ //02// empty set
			db:        mock.NewDB(),
			setCookie: true,
			status:    http.StatusOK,
			out:       "[]",
		},
		{ //03// success
			setCookie: true,
			db:        mock.NewDB().WithRecipes(mock.Recipe1, mock.Recipe2),
			out:       fmt.Sprintf(`[%s,%s]`, mock.Recipe1Json, mock.Recipe2Json),
			status:    http.StatusOK,
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{DB: data.db, Session: scs.New()}

		if data.setCookie {
			if ctx, err := env.Session.Load(req.Context(), "abc"); err == nil {
				req = req.WithContext(ctx)
				env.Session.Put(req.Context(), "id", 1)
			}
		}

		api.Recipes(env)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, body, data.out)
		}
	}
}

func TestRecipe(t *testing.T) {
	for idx, data := range []struct {
		db        *mock.DB
		params    httprouter.Params
		setCookie bool

		out    string
		status int
	}{
		{ //00// missing id
			db:     mock.NewDB(),
			status: http.StatusBadRequest,
		},
		{ //01// missing id
			db:     mock.NewDB(),
			params: httprouter.Params{{Key: "name", Value: "12"}},
			status: http.StatusBadRequest,
		},
		{ //02// wrong data type for id
			db:     mock.NewDB(),
			params: httprouter.Params{{Key: "id", Value: "someone"}},
			status: http.StatusBadRequest,
		},
		{ //03// insufficient permission
			db:        mock.NewDB(),
			params:    httprouter.Params{{Key: "id", Value: "42"}},
			setCookie: true,
			status:    http.StatusUnauthorized,
		},
		{ //04// connection failure
			db:        mock.NewDB().WithAccess(mock.Access{1, 42, app.PermRead}).WithError(nil, mock.ErrDOS),
			params:    httprouter.Params{{Key: "id", Value: "42"}},
			setCookie: true,
			status:    http.StatusInternalServerError,
		},
		{ //05// item doesn't exist
			db:        mock.NewDB().WithAccess(mock.Access{User: 1, Resource: 42, Perms: app.PermRead}),
			params:    httprouter.Params{{Key: "id", Value: "42"}},
			setCookie: true,
			status:    http.StatusNotFound,
		},
		{ //06// success
			db:        mock.NewDB().WithRecipe(mock.Recipe1).WithAccess(mock.Access{1, 1, app.PermRead}),
			params:    httprouter.Params{{Key: "id", Value: "1"}},
			setCookie: true,
			status:    http.StatusOK,
			out:       mock.Recipe1Json,
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{DB: data.db, Session: scs.New()}

		if data.setCookie {
			if ctx, err := env.Session.Load(req.Context(), "abc"); err == nil {
				req = req.WithContext(ctx)
				env.Session.Put(req.Context(), "id", 1)
			}
		}

		api.Recipe(env)(res, req, data.params)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, body, data.out)
		}
	}
}

func TestRecipeOwner(t *testing.T) {
	for idx, data := range []struct {
		db        *mock.DB
		params    httprouter.Params
		setCookie bool

		out    string
		status int
	}{
		{ //00// missing id
			db:     mock.NewDB(),
			status: http.StatusBadRequest,
		},
		{ //01// missing id
			db:     mock.NewDB(),
			params: httprouter.Params{{Key: "name", Value: "12"}},
			status: http.StatusBadRequest,
		},
		{ //02// wrong data type for id
			db:     mock.NewDB(),
			params: httprouter.Params{{Key: "id", Value: "someone"}},
			status: http.StatusBadRequest,
		},
		{ //03// connection failure
			db:        mock.NewDB().WithError(mock.ErrDOS),
			params:    httprouter.Params{{Key: "id", Value: "42"}},
			setCookie: true,
			status:    http.StatusInternalServerError,
		},
		{ //04// item doesn't exist
			db:     mock.NewDB(),
			params: httprouter.Params{{Key: "id", Value: "0"}},
			status: http.StatusNotFound,
		},
		{ //05// success for owner
			db:        mock.NewDB().WithAccess(mock.Access{1, 2, app.PermOwner}),
			params:    httprouter.Params{{Key: "id", Value: "2"}},
			setCookie: true,
			status:    http.StatusOK,
			out:       `{"isowner":true,"owner":""}`,
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		env := &handler.Environment{DB: data.db, Session: scs.New()}

		if data.setCookie {
			if ctx, err := env.Session.Load(req.Context(), "abc"); err == nil {
				req = req.WithContext(ctx)
				env.Session.Put(req.Context(), "id", 1)
			}
		}

		api.RecipeOwner(env)(res, req, data.params)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, body, data.out)
		}
	}
}
