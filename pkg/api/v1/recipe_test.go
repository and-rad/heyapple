package api_test

import (
	"heyapple/internal/mock"
	"heyapple/pkg/api/v1"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
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
