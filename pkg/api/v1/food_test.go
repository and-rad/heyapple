package api_test

import (
	"fmt"
	"heyapple/internal/mock"
	"heyapple/pkg/api/v1"
	"heyapple/pkg/core"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFoods(t *testing.T) {
	for idx, data := range []struct {
		db     *mock.DB
		out    string
		status int
	}{
		{ //00// connection failure
			db:     mock.NewDB().Fail(true),
			status: http.StatusInternalServerError,
		},
		{ //01// empty set
			db:     mock.NewDB(),
			status: http.StatusOK,
			out:    "[]",
		},
		{ //02// success
			db:     &mock.DB{FoodInfo: []core.Food{food1, food2}},
			status: http.StatusOK,
			out:    fmt.Sprintf(`[%s,%s]`, food1json, food2json),
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		api.Foods(data.db)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.out {
			t.Errorf("test case %d: data mismatch \nhave: %v \nwant: %v", idx, body, data.out)
		}
	}
}
