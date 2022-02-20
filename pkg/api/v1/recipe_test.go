package api_test

import (
	"heyapple/internal/mock"
	"heyapple/pkg/api/v1"
	"heyapple/pkg/core"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestNewRecipe(t *testing.T) {
	for idx, data := range []struct {
		db     *mock.DB
		in     url.Values
		out    string
		status int
	}{
		{ //00// missing input
			db:     mock.NewDB(),
			status: http.StatusBadRequest,
		},
		{ //01// connection failure
			in:     url.Values{"name": {"My Recipe"}},
			db:     mock.NewDB().WithError(mock.ErrDOS),
			status: http.StatusInternalServerError,
		},
		{ //02// success
			in:     url.Values{"name": {"My Recipe"}},
			db:     mock.NewDB().WithID(23),
			status: http.StatusCreated,
			out:    "23",
		},
	} {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data.in.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		res := httptest.NewRecorder()
		api.NewRecipe(data.db)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, body, data.out)
		}
	}
}

func TestSaveRecipe(t *testing.T) {
	for idx, data := range []struct {
		db     *mock.DB
		params httprouter.Params
		in     url.Values

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
		{ //03// connection failure
			db:     mock.NewDB().WithError(mock.ErrDOS),
			params: httprouter.Params{{Key: "id", Value: "42"}},
			status: http.StatusInternalServerError,
		},
		{ //04// item doesn't exist
			db:     mock.NewDB().WithRecipe(mock.Recipe1),
			params: httprouter.Params{{Key: "id", Value: "42"}},
			in:     url.Values{"size": {"9"}},
			status: http.StatusNotFound,
			rec:    mock.Recipe1,
		},
		{ //05// success
			db:     mock.NewDB().WithRecipe(mock.Recipe1),
			params: httprouter.Params{{Key: "id", Value: "1"}},
			in:     url.Values{"size": {"9"}},
			status: http.StatusNoContent,
			rec:    core.Recipe{ID: 1, Size: 9, Items: []core.Ingredient{}},
		},
		{ //06// ignore invalid values
			db:     mock.NewDB().WithRecipe(mock.Recipe1).WithFoods(mock.Food1, mock.Food2),
			params: httprouter.Params{{Key: "id", Value: "1"}},
			in:     url.Values{"item": {"1", "2", "34"}, "amount": {"100", "alot", "340"}},
			status: http.StatusNoContent,
			rec:    core.Recipe{ID: 1, Size: 1, Items: []core.Ingredient{{ID: 1, Amount: 100}}},
		},
		{ //07// array sizes don't match
			db:     mock.NewDB().WithRecipe(mock.Recipe1),
			params: httprouter.Params{{Key: "id", Value: "1"}},
			in:     url.Values{"item": {"1", "2"}, "amount": {"100", "250", "340"}},
			status: http.StatusBadRequest,
			rec:    mock.Recipe1,
		},
	} {
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(data.in.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		res := httptest.NewRecorder()
		api.SaveRecipe(data.db)(res, req, data.params)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}

		if !reflect.DeepEqual(data.db.RecipeItem, data.rec) {
			t.Errorf("test case %d: data mismatch \nhave: %v\nwant: %v", idx, data.db.RecipeItem, data.rec)
		}
	}
}
